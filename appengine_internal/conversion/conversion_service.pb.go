// Code generated by protoc-gen-go from "conversion_service.proto"
// DO NOT EDIT!

package appengine

import proto "code.google.com/p/goprotobuf/proto"
import "math"

// Reference proto and math imports to suppress error if they are not otherwise used.
var _ = proto.GetString
var _ = math.Inf

type ConversionServiceError_ErrorCode int32

const (
	ConversionServiceError_OK                     ConversionServiceError_ErrorCode = 0
	ConversionServiceError_TIMEOUT                ConversionServiceError_ErrorCode = 1
	ConversionServiceError_TRANSIENT_ERROR        ConversionServiceError_ErrorCode = 2
	ConversionServiceError_INTERNAL_ERROR         ConversionServiceError_ErrorCode = 3
	ConversionServiceError_UNSUPPORTED_CONVERSION ConversionServiceError_ErrorCode = 4
	ConversionServiceError_CONVERSION_TOO_LARGE   ConversionServiceError_ErrorCode = 5
	ConversionServiceError_TOO_MANY_CONVERSIONS   ConversionServiceError_ErrorCode = 6
	ConversionServiceError_INVALID_REQUEST        ConversionServiceError_ErrorCode = 7
)

var ConversionServiceError_ErrorCode_name = map[int32]string{
	0: "OK",
	1: "TIMEOUT",
	2: "TRANSIENT_ERROR",
	3: "INTERNAL_ERROR",
	4: "UNSUPPORTED_CONVERSION",
	5: "CONVERSION_TOO_LARGE",
	6: "TOO_MANY_CONVERSIONS",
	7: "INVALID_REQUEST",
}
var ConversionServiceError_ErrorCode_value = map[string]int32{
	"OK":                     0,
	"TIMEOUT":                1,
	"TRANSIENT_ERROR":        2,
	"INTERNAL_ERROR":         3,
	"UNSUPPORTED_CONVERSION": 4,
	"CONVERSION_TOO_LARGE":   5,
	"TOO_MANY_CONVERSIONS":   6,
	"INVALID_REQUEST":        7,
}

// NewConversionServiceError_ErrorCode is deprecated. Use x.Enum() instead.
func NewConversionServiceError_ErrorCode(x ConversionServiceError_ErrorCode) *ConversionServiceError_ErrorCode {
	e := ConversionServiceError_ErrorCode(x)
	return &e
}
func (x ConversionServiceError_ErrorCode) Enum() *ConversionServiceError_ErrorCode {
	p := new(ConversionServiceError_ErrorCode)
	*p = x
	return p
}
func (x ConversionServiceError_ErrorCode) String() string {
	return proto.EnumName(ConversionServiceError_ErrorCode_name, int32(x))
}

type ConversionServiceError struct {
	XXX_unrecognized []byte `json:"-"`
}

func (this *ConversionServiceError) Reset()         { *this = ConversionServiceError{} }
func (this *ConversionServiceError) String() string { return proto.CompactTextString(this) }

type AssetInfo struct {
	Name             *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Data             []byte  `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
	MimeType         *string `protobuf:"bytes,3,opt,name=mime_type" json:"mime_type,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *AssetInfo) Reset()         { *this = AssetInfo{} }
func (this *AssetInfo) String() string { return proto.CompactTextString(this) }

type DocumentInfo struct {
	Asset            []*AssetInfo `protobuf:"bytes,1,rep,name=asset" json:"asset,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (this *DocumentInfo) Reset()         { *this = DocumentInfo{} }
func (this *DocumentInfo) String() string { return proto.CompactTextString(this) }

type ConversionInput struct {
	Input            *DocumentInfo              `protobuf:"bytes,1,req,name=input" json:"input,omitempty"`
	OutputMimeType   *string                    `protobuf:"bytes,2,req,name=output_mime_type" json:"output_mime_type,omitempty"`
	Flag             []*ConversionInput_AuxData `protobuf:"bytes,3,rep,name=flag" json:"flag,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (this *ConversionInput) Reset()         { *this = ConversionInput{} }
func (this *ConversionInput) String() string { return proto.CompactTextString(this) }

type ConversionInput_AuxData struct {
	Key              *string `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	Value            *string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *ConversionInput_AuxData) Reset()         { *this = ConversionInput_AuxData{} }
func (this *ConversionInput_AuxData) String() string { return proto.CompactTextString(this) }

type ConversionOutput struct {
	ErrorCode        *ConversionServiceError_ErrorCode `protobuf:"varint,1,req,name=error_code,enum=appengine.ConversionServiceError_ErrorCode" json:"error_code,omitempty"`
	Output           *DocumentInfo                     `protobuf:"bytes,2,opt,name=output" json:"output,omitempty"`
	XXX_unrecognized []byte                            `json:"-"`
}

func (this *ConversionOutput) Reset()         { *this = ConversionOutput{} }
func (this *ConversionOutput) String() string { return proto.CompactTextString(this) }

type ConversionRequest struct {
	Conversion       []*ConversionInput `protobuf:"bytes,1,rep,name=conversion" json:"conversion,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (this *ConversionRequest) Reset()         { *this = ConversionRequest{} }
func (this *ConversionRequest) String() string { return proto.CompactTextString(this) }

type ConversionResponse struct {
	Result           []*ConversionOutput `protobuf:"bytes,1,rep,name=result" json:"result,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (this *ConversionResponse) Reset()         { *this = ConversionResponse{} }
func (this *ConversionResponse) String() string { return proto.CompactTextString(this) }

func init() {
	proto.RegisterEnum("appengine.ConversionServiceError_ErrorCode", ConversionServiceError_ErrorCode_name, ConversionServiceError_ErrorCode_value)
}
