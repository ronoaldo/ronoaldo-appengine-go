// Code generated by protoc-gen-go from "xmpp_service.proto"
// DO NOT EDIT!

package appengine

import proto "code.google.com/p/goprotobuf/proto"
import "math"

// Reference proto and math imports to suppress error if they are not otherwise used.
var _ = proto.GetString
var _ = math.Inf

type XmppServiceError_ErrorCode int32

const (
	XmppServiceError_UNSPECIFIED_ERROR    XmppServiceError_ErrorCode = 1
	XmppServiceError_INVALID_JID          XmppServiceError_ErrorCode = 2
	XmppServiceError_NO_BODY              XmppServiceError_ErrorCode = 3
	XmppServiceError_INVALID_XML          XmppServiceError_ErrorCode = 4
	XmppServiceError_INVALID_TYPE         XmppServiceError_ErrorCode = 5
	XmppServiceError_INVALID_SHOW         XmppServiceError_ErrorCode = 6
	XmppServiceError_EXCEEDED_MAX_SIZE    XmppServiceError_ErrorCode = 7
	XmppServiceError_APPID_ALIAS_REQUIRED XmppServiceError_ErrorCode = 8
)

var XmppServiceError_ErrorCode_name = map[int32]string{
	1: "UNSPECIFIED_ERROR",
	2: "INVALID_JID",
	3: "NO_BODY",
	4: "INVALID_XML",
	5: "INVALID_TYPE",
	6: "INVALID_SHOW",
	7: "EXCEEDED_MAX_SIZE",
	8: "APPID_ALIAS_REQUIRED",
}
var XmppServiceError_ErrorCode_value = map[string]int32{
	"UNSPECIFIED_ERROR":    1,
	"INVALID_JID":          2,
	"NO_BODY":              3,
	"INVALID_XML":          4,
	"INVALID_TYPE":         5,
	"INVALID_SHOW":         6,
	"EXCEEDED_MAX_SIZE":    7,
	"APPID_ALIAS_REQUIRED": 8,
}

// NewXmppServiceError_ErrorCode is deprecated. Use x.Enum() instead.
func NewXmppServiceError_ErrorCode(x XmppServiceError_ErrorCode) *XmppServiceError_ErrorCode {
	e := XmppServiceError_ErrorCode(x)
	return &e
}
func (x XmppServiceError_ErrorCode) Enum() *XmppServiceError_ErrorCode {
	p := new(XmppServiceError_ErrorCode)
	*p = x
	return p
}
func (x XmppServiceError_ErrorCode) String() string {
	return proto.EnumName(XmppServiceError_ErrorCode_name, int32(x))
}

type PresenceResponse_SHOW int32

const (
	PresenceResponse_NORMAL         PresenceResponse_SHOW = 0
	PresenceResponse_AWAY           PresenceResponse_SHOW = 1
	PresenceResponse_DO_NOT_DISTURB PresenceResponse_SHOW = 2
	PresenceResponse_CHAT           PresenceResponse_SHOW = 3
	PresenceResponse_EXTENDED_AWAY  PresenceResponse_SHOW = 4
)

var PresenceResponse_SHOW_name = map[int32]string{
	0: "NORMAL",
	1: "AWAY",
	2: "DO_NOT_DISTURB",
	3: "CHAT",
	4: "EXTENDED_AWAY",
}
var PresenceResponse_SHOW_value = map[string]int32{
	"NORMAL":         0,
	"AWAY":           1,
	"DO_NOT_DISTURB": 2,
	"CHAT":           3,
	"EXTENDED_AWAY":  4,
}

// NewPresenceResponse_SHOW is deprecated. Use x.Enum() instead.
func NewPresenceResponse_SHOW(x PresenceResponse_SHOW) *PresenceResponse_SHOW {
	e := PresenceResponse_SHOW(x)
	return &e
}
func (x PresenceResponse_SHOW) Enum() *PresenceResponse_SHOW {
	p := new(PresenceResponse_SHOW)
	*p = x
	return p
}
func (x PresenceResponse_SHOW) String() string {
	return proto.EnumName(PresenceResponse_SHOW_name, int32(x))
}

type XmppMessageResponse_XmppMessageStatus int32

const (
	XmppMessageResponse_NO_ERROR    XmppMessageResponse_XmppMessageStatus = 0
	XmppMessageResponse_INVALID_JID XmppMessageResponse_XmppMessageStatus = 1
	XmppMessageResponse_OTHER_ERROR XmppMessageResponse_XmppMessageStatus = 2
)

var XmppMessageResponse_XmppMessageStatus_name = map[int32]string{
	0: "NO_ERROR",
	1: "INVALID_JID",
	2: "OTHER_ERROR",
}
var XmppMessageResponse_XmppMessageStatus_value = map[string]int32{
	"NO_ERROR":    0,
	"INVALID_JID": 1,
	"OTHER_ERROR": 2,
}

// NewXmppMessageResponse_XmppMessageStatus is deprecated. Use x.Enum() instead.
func NewXmppMessageResponse_XmppMessageStatus(x XmppMessageResponse_XmppMessageStatus) *XmppMessageResponse_XmppMessageStatus {
	e := XmppMessageResponse_XmppMessageStatus(x)
	return &e
}
func (x XmppMessageResponse_XmppMessageStatus) Enum() *XmppMessageResponse_XmppMessageStatus {
	p := new(XmppMessageResponse_XmppMessageStatus)
	*p = x
	return p
}
func (x XmppMessageResponse_XmppMessageStatus) String() string {
	return proto.EnumName(XmppMessageResponse_XmppMessageStatus_name, int32(x))
}

type XmppServiceError struct {
	XXX_unrecognized []byte `json:"-"`
}

func (this *XmppServiceError) Reset()         { *this = XmppServiceError{} }
func (this *XmppServiceError) String() string { return proto.CompactTextString(this) }

type PresenceRequest struct {
	Jid              *string `protobuf:"bytes,1,req,name=jid" json:"jid,omitempty"`
	FromJid          *string `protobuf:"bytes,2,opt,name=from_jid" json:"from_jid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *PresenceRequest) Reset()         { *this = PresenceRequest{} }
func (this *PresenceRequest) String() string { return proto.CompactTextString(this) }

type PresenceResponse struct {
	IsAvailable      *bool                  `protobuf:"varint,1,req,name=is_available" json:"is_available,omitempty"`
	Presence         *PresenceResponse_SHOW `protobuf:"varint,2,opt,name=presence,enum=appengine.PresenceResponse_SHOW" json:"presence,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (this *PresenceResponse) Reset()         { *this = PresenceResponse{} }
func (this *PresenceResponse) String() string { return proto.CompactTextString(this) }

type XmppMessageRequest struct {
	Jid              []string `protobuf:"bytes,1,rep,name=jid" json:"jid,omitempty"`
	Body             *string  `protobuf:"bytes,2,req,name=body" json:"body,omitempty"`
	RawXml           *bool    `protobuf:"varint,3,opt,name=raw_xml,def=0" json:"raw_xml,omitempty"`
	Type             *string  `protobuf:"bytes,4,opt,name=type,def=chat" json:"type,omitempty"`
	FromJid          *string  `protobuf:"bytes,5,opt,name=from_jid" json:"from_jid,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (this *XmppMessageRequest) Reset()         { *this = XmppMessageRequest{} }
func (this *XmppMessageRequest) String() string { return proto.CompactTextString(this) }

const Default_XmppMessageRequest_RawXml bool = false
const Default_XmppMessageRequest_Type string = "chat"

type XmppMessageResponse struct {
	Status           []XmppMessageResponse_XmppMessageStatus `protobuf:"varint,1,rep,name=status,enum=appengine.XmppMessageResponse_XmppMessageStatus" json:"status,omitempty"`
	XXX_unrecognized []byte                                  `json:"-"`
}

func (this *XmppMessageResponse) Reset()         { *this = XmppMessageResponse{} }
func (this *XmppMessageResponse) String() string { return proto.CompactTextString(this) }

type XmppSendPresenceRequest struct {
	Jid              *string `protobuf:"bytes,1,req,name=jid" json:"jid,omitempty"`
	Type             *string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Show             *string `protobuf:"bytes,3,opt,name=show" json:"show,omitempty"`
	Status           *string `protobuf:"bytes,4,opt,name=status" json:"status,omitempty"`
	FromJid          *string `protobuf:"bytes,5,opt,name=from_jid" json:"from_jid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *XmppSendPresenceRequest) Reset()         { *this = XmppSendPresenceRequest{} }
func (this *XmppSendPresenceRequest) String() string { return proto.CompactTextString(this) }

type XmppSendPresenceResponse struct {
	XXX_unrecognized []byte `json:"-"`
}

func (this *XmppSendPresenceResponse) Reset()         { *this = XmppSendPresenceResponse{} }
func (this *XmppSendPresenceResponse) String() string { return proto.CompactTextString(this) }

type XmppInviteRequest struct {
	Jid              *string `protobuf:"bytes,1,req,name=jid" json:"jid,omitempty"`
	FromJid          *string `protobuf:"bytes,2,opt,name=from_jid" json:"from_jid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *XmppInviteRequest) Reset()         { *this = XmppInviteRequest{} }
func (this *XmppInviteRequest) String() string { return proto.CompactTextString(this) }

type XmppInviteResponse struct {
	XXX_unrecognized []byte `json:"-"`
}

func (this *XmppInviteResponse) Reset()         { *this = XmppInviteResponse{} }
func (this *XmppInviteResponse) String() string { return proto.CompactTextString(this) }

func init() {
	proto.RegisterEnum("appengine.XmppServiceError_ErrorCode", XmppServiceError_ErrorCode_name, XmppServiceError_ErrorCode_value)
	proto.RegisterEnum("appengine.PresenceResponse_SHOW", PresenceResponse_SHOW_name, PresenceResponse_SHOW_value)
	proto.RegisterEnum("appengine.XmppMessageResponse_XmppMessageStatus", XmppMessageResponse_XmppMessageStatus_name, XmppMessageResponse_XmppMessageStatus_value)
}
