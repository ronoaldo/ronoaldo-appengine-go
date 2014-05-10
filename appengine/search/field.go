// Copyright 2014 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package search

import (
	"fmt"
	"reflect"
)

// Field is a name/value pair. A search index's document can be loaded and
// saved as a sequence of Fields.
type Field struct {
	// Name is the field name. A valid field name matches /[A-Z][A-Za-z0-9_]*/.
	// A field name cannot be longer than 500 characters.
	Name string
	// Value is the field value. The valid types are:
	//  - string,
	//  - search.Atom,
	//  - search.HTML,
	//  - time.Time (stored with millisecond precision),
	//  - float64,
	//  - appengine.GeoPoint.
	Value interface{}
}

// FieldLoadSaver can be converted from and to a slice of Fields.
type FieldLoadSaver interface {
	Load([]Field) error
	Save() ([]Field, error)
}

// FieldList converts a []Field to implement FieldLoadSaver.
type FieldList []Field

// Load loads all of the provided fields into l.
// It does not first reset *l to an empty slice.
func (l *FieldList) Load(f []Field) error {
	*l = append(*l, f...)
	return nil
}

// Save returns all of l's fields as a slice of Fields.
func (l *FieldList) Save() ([]Field, error) {
	return *l, nil
}

// structFLS adapts a struct to be a FieldLoadSaver.
type structFLS struct {
	reflect.Value
}

func (s structFLS) Load(fields []Field) error {
	for _, field := range fields {
		f := s.FieldByName(field.Name)
		if !f.IsValid() {
			// TODO: continue but eventually return ErrFieldMismatch, similar to package datastore.
			continue
		}
		if !f.CanSet() {
			continue
		}
		v := reflect.ValueOf(field.Value)
		if ft, vt := f.Type(), v.Type(); ft != vt {
			return fmt.Errorf("search: type mismatch: %v for %v data", ft, vt)
		}
		f.Set(v)
	}
	return nil
}

func (s structFLS) Save() ([]Field, error) {
	fields := make([]Field, 0, s.NumField())
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		if !f.CanSet() {
			continue
		}
		fields = append(fields, Field{
			Name:  s.Type().Field(i).Name,
			Value: f.Interface(),
		})
	}
	return fields, nil
}

// newStructFLS returns a FieldLoadSaver for the struct pointer p.
func newStructFLS(p interface{}) (FieldLoadSaver, error) {
	v := reflect.ValueOf(p)
	if v.Kind() != reflect.Ptr || v.IsNil() || v.Elem().Kind() != reflect.Struct {
		return nil, ErrInvalidDocumentType
	}
	return structFLS{v.Elem()}, nil
}

// LoadStruct loads the fields from f to dst. dst must be a struct pointer.
func LoadStruct(dst interface{}, f []Field) error {
	x, err := newStructFLS(dst)
	if err != nil {
		return err
	}
	return x.Load(f)
}

// SaveStruct returns the fields from src as a slice of Field.
// src must be a struct pointer.
func SaveStruct(src interface{}) ([]Field, error) {
	x, err := newStructFLS(src)
	if err != nil {
		return nil, err
	}
	return x.Save()
}
