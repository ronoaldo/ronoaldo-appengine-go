// Copyright 2012 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

/*
Package search provides a client for App Engine's search service.

Indexes contains documents, and a document's contents are a mapping from case-
sensitive field names to values. In Go, documents are represented by struct
pointers, and the valid types for a struct's fields are:
  - string,
  - search.Atom,
  - search.HTML,
  - time.Time (stored with millisecond precision),
  - float64,
  - appengine.GeoPoint.

Documents can also be represented by any type implementing the FieldLoadSaver
interface.

Example code:

	type Doc struct {
		Author   string
		Comment  string
		Creation time.Time
	}

	index, err := search.Open("comments")
	if err != nil {
		return err
	}
	newID, err := index.Put(c, "", &Doc{
		Author:   "gopher",
		Comment:  "the truth of the matter",
		Creation: time.Now(),
	})
	if err != nil {
		return err
	}

Searching an index for a query will result in an iterator. As with an iterator
from package datastore, pass a destination struct to Next to decode the next
result. Next will return Done when the iterator is exhausted.

	for t := index.Search(c, "Comment:truth", nil); ; {
		var doc Doc
		id, err := t.Next(&doc)
		if err == search.Done {
			break
		}
		if err != nil {
			return err
		}
		fmt.Fprintf(w, "%s -> %#v\n", id, doc)
	}

Call List to iterate over documents.

	for t := index.List(c, nil); ; {
		var doc Doc
		id, err := t.Next(&doc)
		if err == search.Done {
			break
		}
		if err != nil {
			return err
		}
		fmt.Fprintf(w, "%s -> %#v\n", id, doc)
	}

A single document can also be retrieved by its ID. Pass a destination struct
to Get to hold the resulting document.

	var doc Doc
	err := index.Get(c, id, &doc)
	if err != nil {
		return err
	}

Queries are expressed as strings, plus some optional parameters. The query
language is described at
https://developers.google.com/appengine/docs/python/search/query_strings
*/
package search

// TODO: let Put specify the document language: "en", "fr", etc. Also: order_id?? storage??
// TODO: Index.GetAll (or Iterator.GetAll)?
// TODO: struct <-> protobuf tests.
// TODO: enforce Python's MIN_NUMBER_VALUE and MIN_DATE (which would disallow a zero
// time.Time)? _MAXIMUM_STRING_LENGTH?

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"appengine"
	"appengine_internal"
	"code.google.com/p/goprotobuf/proto"

	pb "appengine_internal/search"
)

var (
	// ErrInvalidDocumentType is returned when methods like Put, Get or Next
	// are passed a dst or src argument of invalid type.
	ErrInvalidDocumentType = errors.New("search: invalid document type")

	// ErrNoSuchDocument is returned when no document was found for a given ID.
	ErrNoSuchDocument = errors.New("search: no such document")
)

// Atom is a document field whose contents are indexed as a single indivisible
// string.
type Atom string

// HTML is a document field whose contents are indexed as HTML. Only text nodes
// are indexed: "foo<b>bar" will be treated as "foobar".
type HTML string

// validIndexNameOrDocID is the Go equivalent of Python's
// _ValidateVisiblePrintableAsciiNotReserved.
func validIndexNameOrDocID(s string) bool {
	if strings.HasPrefix(s, "!") {
		return false
	}
	for _, c := range s {
		if c < 0x21 || 0x7f <= c {
			return false
		}
	}
	return true
}

var fieldNameRE = regexp.MustCompile(`^[A-Z][A-Za-z0-9_]*$`)

// validFieldName is the Go equivalent of Python's _CheckFieldName.
func validFieldName(s string) bool {
	return len(s) <= 500 && fieldNameRE.MatchString(s)
}

// Index is an index of documents.
type Index struct {
	spec pb.IndexSpec
}

// Open opens the index with the given name. The index is created if it does
// not already exist.
//
// The name is a human-readable ASCII string. It must contain no whitespace
// characters and not start with "!".
func Open(name string) (*Index, error) {
	if !validIndexNameOrDocID(name) {
		return nil, fmt.Errorf("search: invalid index name %q", name)
	}
	return &Index{
		spec: pb.IndexSpec{
			Name: &name,
		},
	}, nil
}

// Put saves src to the index. If id is empty, a new ID is allocated by the
// service and returned. If id is not empty, any existing index entry for that
// ID is replaced.
//
// The ID is a human-readable ASCII string. It must contain no whitespace
// characters and not start with "!".
//
// src must be a non-nil struct pointer or implement the FieldLoadSaver
// interface.
func (x *Index) Put(c appengine.Context, id string, src interface{}) (string, error) {
	fields, err := saveFields(src)
	if err != nil {
		return "", err
	}
	d := &pb.Document{
		Field: fields,
	}
	if id != "" {
		if !validIndexNameOrDocID(id) {
			return "", fmt.Errorf("search: invalid ID %q", id)
		}
		d.Id = proto.String(id)
	}
	req := &pb.IndexDocumentRequest{
		Params: &pb.IndexDocumentParams{
			Document:  []*pb.Document{d},
			IndexSpec: &x.spec,
		},
	}
	res := &pb.IndexDocumentResponse{}
	if err := c.Call("search", "IndexDocument", req, res, nil); err != nil {
		return "", err
	}
	if len(res.Status) != 1 || len(res.DocId) != 1 {
		return "", fmt.Errorf("search: internal error: wrong number of results (%d Statuses, %d DocIDs)",
			len(res.Status), len(res.DocId))
	}
	if s := res.Status[0]; s.GetCode() != pb.SearchServiceError_OK {
		return "", fmt.Errorf("search: %s: %s", s.GetCode(), s.GetErrorDetail())
	}
	return res.DocId[0], nil
}

// Get loads the document with the given ID into dst.
//
// The ID is a human-readable ASCII string. It must be non-empty, contain no
// whitespace characters and not start with "!".
//
// dst must be a non-nil struct pointer or implement the FieldLoadSaver
// interface.
func (x *Index) Get(c appengine.Context, id string, dst interface{}) error {
	if id == "" || !validIndexNameOrDocID(id) {
		return fmt.Errorf("search: invalid ID %q", id)
	}
	req := &pb.ListDocumentsRequest{
		Params: &pb.ListDocumentsParams{
			IndexSpec:  &x.spec,
			StartDocId: proto.String(id),
			Limit:      proto.Int32(1),
		},
	}
	res := &pb.ListDocumentsResponse{}
	if err := c.Call("search", "ListDocuments", req, res, nil); err != nil {
		return err
	}
	if res.Status == nil || res.Status.GetCode() != pb.SearchServiceError_OK {
		return fmt.Errorf("search: %s: %s", res.Status.GetCode(), res.Status.GetErrorDetail())
	}
	if len(res.Document) != 1 || res.Document[0].GetId() != id {
		return ErrNoSuchDocument
	}
	return loadFields(dst, res.Document[0].Field)
}

// Delete deletes a document from the index.
func (x *Index) Delete(c appengine.Context, id string) error {
	req := &pb.DeleteDocumentRequest{
		Params: &pb.DeleteDocumentParams{
			DocId:     []string{id},
			IndexSpec: &x.spec,
		},
	}
	res := &pb.DeleteDocumentResponse{}
	if err := c.Call("search", "DeleteDocument", req, res, nil); err != nil {
		return err
	}
	if len(res.Status) != 1 {
		return fmt.Errorf("search: internal error: wrong number of results (%d)", len(res.Status))
	}
	if s := res.Status[0]; s.GetCode() != pb.SearchServiceError_OK {
		return fmt.Errorf("search: %s: %s", s.GetCode(), s.GetErrorDetail())
	}
	return nil
}

// List lists all of the documents in an index. The documents are returned in
// increasing ID order.
func (x *Index) List(c appengine.Context, opts *ListOptions) *Iterator {
	t := &Iterator{
		c:             c,
		index:         x,
		count:         -1,
		listInclusive: true,
		more:          moreList,
	}
	if opts != nil {
		if opts.StartID != "" {
			t.listStartID = opts.StartID
		}
	}
	return t
}

func moreList(t *Iterator) error {
	req := &pb.ListDocumentsRequest{
		Params: &pb.ListDocumentsParams{
			IndexSpec: &t.index.spec,
		},
	}
	if t.listStartID != "" {
		req.Params.StartDocId = &t.listStartID
		req.Params.IncludeStartDoc = &t.listInclusive
	}

	res := &pb.ListDocumentsResponse{}
	if err := t.c.Call("search", "ListDocuments", req, res, nil); err != nil {
		return err
	}
	if res.Status == nil || res.Status.GetCode() != pb.SearchServiceError_OK {
		return fmt.Errorf("search: %s: %s", res.Status.GetCode(), res.Status.GetErrorDetail())
	}
	t.listRes = res.Document
	t.listStartID, t.listInclusive, t.more = "", false, nil
	if len(res.Document) != 0 {
		if id := res.Document[len(res.Document)-1].GetId(); id != "" {
			t.listStartID, t.more = id, moreList
		}
	}
	return nil
}

// ListOptions are the options for listing documents in an index. Passing a nil
// *ListOptions is equivalent to using the default values.
type ListOptions struct {
	// StartID is the inclusive lower bound for the ID of the returned
	// documents. The zero value means all documents will be returned.
	StartID string

	// TODO: limit, idsOnly, maybe others.
}

// Search searches the index for the given query.
func (x *Index) Search(c appengine.Context, query string, opts *SearchOptions) *Iterator {
	return &Iterator{
		c:           c,
		index:       x,
		searchQuery: query,
		more:        moreSearch,
	}
}

func moreSearch(t *Iterator) error {
	req := &pb.SearchRequest{
		Params: &pb.SearchParams{
			IndexSpec:  &t.index.spec,
			Query:      &t.searchQuery,
			CursorType: pb.SearchParams_SINGLE.Enum(),
		},
	}
	if t.searchCursor != nil {
		req.Params.Cursor = t.searchCursor
	}
	res := &pb.SearchResponse{}
	if err := t.c.Call("search", "Search", req, res, nil); err != nil {
		return err
	}
	if res.Status == nil || res.Status.GetCode() != pb.SearchServiceError_OK {
		return fmt.Errorf("search: %s: %s", res.Status.GetCode(), res.Status.GetErrorDetail())
	}
	t.searchRes = res.Result
	t.count = int(*res.MatchedCount)
	if res.Cursor != nil {
		t.searchCursor, t.more = res.Cursor, moreSearch
	} else {
		t.searchCursor, t.more = nil, nil
	}
	return nil
}

// SearchOptions are the options for searching an index. Passing a nil
// *SearchOptions is equivalent to using the default values.
//
// There are currently no options. Future versions may introduce some.
type SearchOptions struct {
	// TODO: limit, cursor, offset, idsOnly, maybe others.
}

// Iterator is the result of searching an index for a query or listing an
// index.
type Iterator struct {
	c     appengine.Context
	index *Index
	err   error

	listRes       []*pb.Document
	listStartID   string
	listInclusive bool

	searchRes    []*pb.SearchResult
	searchQuery  string
	searchCursor *string

	more func(*Iterator) error

	count int
}

// Done is returned when a query iteration has completed.
var Done = errors.New("search: query has no more results")

// Count returns an approximation of the number of documents matched by the
// query. It is only valid to call for iterators returned by Search.
func (t *Iterator) Count() int { return t.count }

// Next returns the ID of the next result. When there are no more results,
// Done is returned as the error.
//
// dst must be a non-nil struct pointer, implement the FieldLoadSaver
// interface, or be a nil interface value. If a non-nil dst is provided, it
// will be filled with the indexed fields.
func (t *Iterator) Next(dst interface{}) (string, error) {
	if t.err == nil && len(t.listRes)+len(t.searchRes) == 0 && t.more != nil {
		t.err = t.more(t)
	}
	if t.err != nil {
		return "", t.err
	}

	var doc *pb.Document
	switch {
	case len(t.listRes) != 0:
		doc = t.listRes[0]
		t.listRes = t.listRes[1:]
	case len(t.searchRes) != 0:
		doc = t.searchRes[0].Document
		t.searchRes = t.searchRes[1:]
	default:
		return "", Done
	}
	if doc == nil {
		return "", errors.New("search: internal error: no document returned")
	}
	if dst != nil {
		if err := loadFields(dst, doc.Field); err != nil {
			return "", err
		}
	}
	return doc.GetId(), nil
}

// saveFields converts from a struct pointer or FieldLoadSaver to protobufs.
func saveFields(src interface{}) ([]*pb.Field, error) {
	var err error
	var fields []Field
	if x, ok := src.(FieldLoadSaver); ok {
		fields, err = x.Save()
	} else {
		fields, err = SaveStruct(src)
	}
	if err != nil {
		return nil, err
	}
	return fieldsToProto(fields)
}

func fieldsToProto(src []Field) ([]*pb.Field, error) {
	// Maps to catch duplicate time or numeric fields.
	timeFields, numericFields := make(map[string]bool), make(map[string]bool)
	dst := make([]*pb.Field, 0, len(src))
	for _, f := range src {
		if !validFieldName(f.Name) {
			return nil, fmt.Errorf("search: invalid field name %q", f.Name)
		}
		fieldValue := &pb.FieldValue{}
		switch x := f.Value.(type) {
		case string:
			fieldValue.Type = pb.FieldValue_TEXT.Enum()
			fieldValue.StringValue = proto.String(x)
		case Atom:
			fieldValue.Type = pb.FieldValue_ATOM.Enum()
			fieldValue.StringValue = proto.String(string(x))
		case HTML:
			fieldValue.Type = pb.FieldValue_HTML.Enum()
			fieldValue.StringValue = proto.String(string(x))
		case time.Time:
			if timeFields[f.Name] {
				return nil, fmt.Errorf("search: duplicate time field %q", f.Name)
			}
			timeFields[f.Name] = true
			fieldValue.Type = pb.FieldValue_DATE.Enum()
			fieldValue.StringValue = proto.String(strconv.FormatInt(x.UnixNano()/1e6, 10))
		case float64:
			if numericFields[f.Name] {
				return nil, fmt.Errorf("search: duplicate numeric field %q", f.Name)
			}
			numericFields[f.Name] = true
			fieldValue.Type = pb.FieldValue_NUMBER.Enum()
			fieldValue.StringValue = proto.String(strconv.FormatFloat(x, 'e', -1, 64))
		case appengine.GeoPoint:
			if !x.Valid() {
				return nil, fmt.Errorf(
					"search: GeoPoint field %q with invalid value %v",
					f.Name, x)
			}
			fieldValue.Type = pb.FieldValue_GEO.Enum()
			fieldValue.Geo = &pb.FieldValue_Geo{
				Lat: proto.Float64(x.Lat),
				Lng: proto.Float64(x.Lng),
			}
		default:
			return nil, fmt.Errorf("search: unsupported field type: %v", reflect.TypeOf(f.Value))
		}
		if p := fieldValue.StringValue; p != nil && !utf8.ValidString(*p) {
			return nil, fmt.Errorf("search: %q field is invalid UTF-8: %q", f.Name, *p)
		}
		dst = append(dst, &pb.Field{
			Name:  proto.String(f.Name),
			Value: fieldValue,
		})
	}
	return dst, nil
}

// loadFields converts from protobufs to a struct pointer or FieldLoadSaver.
func loadFields(dst interface{}, src []*pb.Field) (err error) {
	fields, err := protoToFields(src)
	if err != nil {
		return err
	}
	if x, ok := dst.(FieldLoadSaver); ok {
		return x.Load(fields)
	}
	return LoadStruct(dst, fields)
}

func protoToFields(fields []*pb.Field) ([]Field, error) {
	dst := make([]Field, 0, len(fields))
	for _, field := range fields {
		fieldValue := field.GetValue()
		f := Field{
			Name: field.GetName(),
		}
		switch fieldValue.GetType() {
		case pb.FieldValue_TEXT:
			f.Value = fieldValue.GetStringValue()
		case pb.FieldValue_ATOM:
			f.Value = Atom(fieldValue.GetStringValue())
		case pb.FieldValue_HTML:
			f.Value = HTML(fieldValue.GetStringValue())
		case pb.FieldValue_DATE:
			sv := fieldValue.GetStringValue()
			millis, err := strconv.ParseInt(sv, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("search: internal error: bad time.Time encoding %q: %v", sv, err)
			}
			f.Value = time.Unix(0, millis*1e6)
		case pb.FieldValue_NUMBER:
			sv := fieldValue.GetStringValue()
			x, err := strconv.ParseFloat(sv, 64)
			if err != nil {
				return nil, err
			}
			f.Value = x
		case pb.FieldValue_GEO:
			geoValue := fieldValue.GetGeo()
			geoPoint := appengine.GeoPoint{geoValue.GetLat(), geoValue.GetLng()}
			if !geoPoint.Valid() {
				return nil, fmt.Errorf("search: internal error: invalid GeoPoint encoding: %v", geoPoint)
			}
			f.Value = geoPoint
		default:
			return nil, fmt.Errorf("search: internal error: unknown data type %s", fieldValue.GetType())
		}
		dst = append(dst, f)
	}
	return dst, nil
}

func namespaceMod(m appengine_internal.ProtoMessage, namespace string) {
	set := func(s **string) {
		if *s == nil {
			*s = &namespace
		}
	}
	switch m := m.(type) {
	case *pb.IndexDocumentRequest:
		set(&m.Params.IndexSpec.Namespace)
	case *pb.ListDocumentsRequest:
		set(&m.Params.IndexSpec.Namespace)
	case *pb.DeleteDocumentRequest:
		set(&m.Params.IndexSpec.Namespace)
	case *pb.SearchRequest:
		set(&m.Params.IndexSpec.Namespace)
	}
}

func init() {
	appengine_internal.RegisterErrorCodeMap("search", pb.SearchServiceError_ErrorCode_name)
	appengine_internal.NamespaceMods["search"] = namespaceMod
}
