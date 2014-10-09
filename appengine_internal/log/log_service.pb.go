// Code generated by protoc-gen-go.
// DO NOT EDIT!

/*
Package appengine is a generated protocol buffer package.

It is generated from these files:
	appengine_internal/log

It has these top-level messages:
	LogServiceError
	UserAppLogLine
	UserAppLogGroup
	FlushRequest
	SetStatusRequest
	LogOffset
	LogLine
	RequestLog
	LogModuleVersion
	LogReadRequest
	LogReadResponse
	LogUsageRecord
	LogUsageRequest
	LogUsageResponse
*/
package appengine

import proto "code.google.com/p/goprotobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type LogServiceError_ErrorCode int32

const (
	LogServiceError_OK              LogServiceError_ErrorCode = 0
	LogServiceError_INVALID_REQUEST LogServiceError_ErrorCode = 1
	LogServiceError_STORAGE_ERROR   LogServiceError_ErrorCode = 2
)

var LogServiceError_ErrorCode_name = map[int32]string{
	0: "OK",
	1: "INVALID_REQUEST",
	2: "STORAGE_ERROR",
}
var LogServiceError_ErrorCode_value = map[string]int32{
	"OK":              0,
	"INVALID_REQUEST": 1,
	"STORAGE_ERROR":   2,
}

func (x LogServiceError_ErrorCode) Enum() *LogServiceError_ErrorCode {
	p := new(LogServiceError_ErrorCode)
	*p = x
	return p
}
func (x LogServiceError_ErrorCode) String() string {
	return proto.EnumName(LogServiceError_ErrorCode_name, int32(x))
}
func (x *LogServiceError_ErrorCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(LogServiceError_ErrorCode_value, data, "LogServiceError_ErrorCode")
	if err != nil {
		return err
	}
	*x = LogServiceError_ErrorCode(value)
	return nil
}

type LogServiceError struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *LogServiceError) Reset()         { *m = LogServiceError{} }
func (m *LogServiceError) String() string { return proto.CompactTextString(m) }
func (*LogServiceError) ProtoMessage()    {}

type UserAppLogLine struct {
	TimestampUsec    *int64  `protobuf:"varint,1,req,name=timestamp_usec" json:"timestamp_usec,omitempty"`
	Level            *int64  `protobuf:"varint,2,req,name=level" json:"level,omitempty"`
	Message          *string `protobuf:"bytes,3,req,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *UserAppLogLine) Reset()         { *m = UserAppLogLine{} }
func (m *UserAppLogLine) String() string { return proto.CompactTextString(m) }
func (*UserAppLogLine) ProtoMessage()    {}

func (m *UserAppLogLine) GetTimestampUsec() int64 {
	if m != nil && m.TimestampUsec != nil {
		return *m.TimestampUsec
	}
	return 0
}

func (m *UserAppLogLine) GetLevel() int64 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

func (m *UserAppLogLine) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

type UserAppLogGroup struct {
	LogLine          []*UserAppLogLine `protobuf:"bytes,2,rep,name=log_line" json:"log_line,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *UserAppLogGroup) Reset()         { *m = UserAppLogGroup{} }
func (m *UserAppLogGroup) String() string { return proto.CompactTextString(m) }
func (*UserAppLogGroup) ProtoMessage()    {}

func (m *UserAppLogGroup) GetLogLine() []*UserAppLogLine {
	if m != nil {
		return m.LogLine
	}
	return nil
}

type FlushRequest struct {
	Logs             []byte `protobuf:"bytes,1,opt,name=logs" json:"logs,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *FlushRequest) Reset()         { *m = FlushRequest{} }
func (m *FlushRequest) String() string { return proto.CompactTextString(m) }
func (*FlushRequest) ProtoMessage()    {}

func (m *FlushRequest) GetLogs() []byte {
	if m != nil {
		return m.Logs
	}
	return nil
}

type SetStatusRequest struct {
	Status           *string `protobuf:"bytes,1,req,name=status" json:"status,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SetStatusRequest) Reset()         { *m = SetStatusRequest{} }
func (m *SetStatusRequest) String() string { return proto.CompactTextString(m) }
func (*SetStatusRequest) ProtoMessage()    {}

func (m *SetStatusRequest) GetStatus() string {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return ""
}

type LogOffset struct {
	RequestId        []byte `protobuf:"bytes,1,opt,name=request_id" json:"request_id,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *LogOffset) Reset()         { *m = LogOffset{} }
func (m *LogOffset) String() string { return proto.CompactTextString(m) }
func (*LogOffset) ProtoMessage()    {}

func (m *LogOffset) GetRequestId() []byte {
	if m != nil {
		return m.RequestId
	}
	return nil
}

type LogLine struct {
	Time             *int64  `protobuf:"varint,1,req,name=time" json:"time,omitempty"`
	Level            *int32  `protobuf:"varint,2,req,name=level" json:"level,omitempty"`
	LogMessage       *string `protobuf:"bytes,3,req,name=log_message" json:"log_message,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *LogLine) Reset()         { *m = LogLine{} }
func (m *LogLine) String() string { return proto.CompactTextString(m) }
func (*LogLine) ProtoMessage()    {}

func (m *LogLine) GetTime() int64 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

func (m *LogLine) GetLevel() int32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

func (m *LogLine) GetLogMessage() string {
	if m != nil && m.LogMessage != nil {
		return *m.LogMessage
	}
	return ""
}

type RequestLog struct {
	AppId                   *string    `protobuf:"bytes,1,req,name=app_id" json:"app_id,omitempty"`
	ModuleId                *string    `protobuf:"bytes,37,opt,name=module_id,def=default" json:"module_id,omitempty"`
	VersionId               *string    `protobuf:"bytes,2,req,name=version_id" json:"version_id,omitempty"`
	RequestId               []byte     `protobuf:"bytes,3,req,name=request_id" json:"request_id,omitempty"`
	Offset                  *LogOffset `protobuf:"bytes,35,opt,name=offset" json:"offset,omitempty"`
	Ip                      *string    `protobuf:"bytes,4,req,name=ip" json:"ip,omitempty"`
	Nickname                *string    `protobuf:"bytes,5,opt,name=nickname" json:"nickname,omitempty"`
	StartTime               *int64     `protobuf:"varint,6,req,name=start_time" json:"start_time,omitempty"`
	EndTime                 *int64     `protobuf:"varint,7,req,name=end_time" json:"end_time,omitempty"`
	Latency                 *int64     `protobuf:"varint,8,req,name=latency" json:"latency,omitempty"`
	Mcycles                 *int64     `protobuf:"varint,9,req,name=mcycles" json:"mcycles,omitempty"`
	Method                  *string    `protobuf:"bytes,10,req,name=method" json:"method,omitempty"`
	Resource                *string    `protobuf:"bytes,11,req,name=resource" json:"resource,omitempty"`
	HttpVersion             *string    `protobuf:"bytes,12,req,name=http_version" json:"http_version,omitempty"`
	Status                  *int32     `protobuf:"varint,13,req,name=status" json:"status,omitempty"`
	ResponseSize            *int64     `protobuf:"varint,14,req,name=response_size" json:"response_size,omitempty"`
	Referrer                *string    `protobuf:"bytes,15,opt,name=referrer" json:"referrer,omitempty"`
	UserAgent               *string    `protobuf:"bytes,16,opt,name=user_agent" json:"user_agent,omitempty"`
	UrlMapEntry             *string    `protobuf:"bytes,17,req,name=url_map_entry" json:"url_map_entry,omitempty"`
	Combined                *string    `protobuf:"bytes,18,req,name=combined" json:"combined,omitempty"`
	ApiMcycles              *int64     `protobuf:"varint,19,opt,name=api_mcycles" json:"api_mcycles,omitempty"`
	Host                    *string    `protobuf:"bytes,20,opt,name=host" json:"host,omitempty"`
	Cost                    *float64   `protobuf:"fixed64,21,opt,name=cost" json:"cost,omitempty"`
	TaskQueueName           *string    `protobuf:"bytes,22,opt,name=task_queue_name" json:"task_queue_name,omitempty"`
	TaskName                *string    `protobuf:"bytes,23,opt,name=task_name" json:"task_name,omitempty"`
	WasLoadingRequest       *bool      `protobuf:"varint,24,opt,name=was_loading_request" json:"was_loading_request,omitempty"`
	PendingTime             *int64     `protobuf:"varint,25,opt,name=pending_time" json:"pending_time,omitempty"`
	ReplicaIndex            *int32     `protobuf:"varint,26,opt,name=replica_index,def=-1" json:"replica_index,omitempty"`
	Finished                *bool      `protobuf:"varint,27,opt,name=finished,def=1" json:"finished,omitempty"`
	CloneKey                []byte     `protobuf:"bytes,28,opt,name=clone_key" json:"clone_key,omitempty"`
	Line                    []*LogLine `protobuf:"bytes,29,rep,name=line" json:"line,omitempty"`
	LinesIncomplete         *bool      `protobuf:"varint,36,opt,name=lines_incomplete" json:"lines_incomplete,omitempty"`
	AppEngineRelease        []byte     `protobuf:"bytes,38,opt,name=app_engine_release" json:"app_engine_release,omitempty"`
	TraceId                 *string    `protobuf:"bytes,39,opt,name=trace_id" json:"trace_id,omitempty"`
	ExitReason              *int32     `protobuf:"varint,30,opt,name=exit_reason" json:"exit_reason,omitempty"`
	WasThrottledForTime     *bool      `protobuf:"varint,31,opt,name=was_throttled_for_time" json:"was_throttled_for_time,omitempty"`
	WasThrottledForRequests *bool      `protobuf:"varint,32,opt,name=was_throttled_for_requests" json:"was_throttled_for_requests,omitempty"`
	ThrottledTime           *int64     `protobuf:"varint,33,opt,name=throttled_time" json:"throttled_time,omitempty"`
	ServerName              []byte     `protobuf:"bytes,34,opt,name=server_name" json:"server_name,omitempty"`
	XXX_unrecognized        []byte     `json:"-"`
}

func (m *RequestLog) Reset()         { *m = RequestLog{} }
func (m *RequestLog) String() string { return proto.CompactTextString(m) }
func (*RequestLog) ProtoMessage()    {}

const Default_RequestLog_ModuleId string = "default"
const Default_RequestLog_ReplicaIndex int32 = -1
const Default_RequestLog_Finished bool = true

func (m *RequestLog) GetAppId() string {
	if m != nil && m.AppId != nil {
		return *m.AppId
	}
	return ""
}

func (m *RequestLog) GetModuleId() string {
	if m != nil && m.ModuleId != nil {
		return *m.ModuleId
	}
	return Default_RequestLog_ModuleId
}

func (m *RequestLog) GetVersionId() string {
	if m != nil && m.VersionId != nil {
		return *m.VersionId
	}
	return ""
}

func (m *RequestLog) GetRequestId() []byte {
	if m != nil {
		return m.RequestId
	}
	return nil
}

func (m *RequestLog) GetOffset() *LogOffset {
	if m != nil {
		return m.Offset
	}
	return nil
}

func (m *RequestLog) GetIp() string {
	if m != nil && m.Ip != nil {
		return *m.Ip
	}
	return ""
}

func (m *RequestLog) GetNickname() string {
	if m != nil && m.Nickname != nil {
		return *m.Nickname
	}
	return ""
}

func (m *RequestLog) GetStartTime() int64 {
	if m != nil && m.StartTime != nil {
		return *m.StartTime
	}
	return 0
}

func (m *RequestLog) GetEndTime() int64 {
	if m != nil && m.EndTime != nil {
		return *m.EndTime
	}
	return 0
}

func (m *RequestLog) GetLatency() int64 {
	if m != nil && m.Latency != nil {
		return *m.Latency
	}
	return 0
}

func (m *RequestLog) GetMcycles() int64 {
	if m != nil && m.Mcycles != nil {
		return *m.Mcycles
	}
	return 0
}

func (m *RequestLog) GetMethod() string {
	if m != nil && m.Method != nil {
		return *m.Method
	}
	return ""
}

func (m *RequestLog) GetResource() string {
	if m != nil && m.Resource != nil {
		return *m.Resource
	}
	return ""
}

func (m *RequestLog) GetHttpVersion() string {
	if m != nil && m.HttpVersion != nil {
		return *m.HttpVersion
	}
	return ""
}

func (m *RequestLog) GetStatus() int32 {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return 0
}

func (m *RequestLog) GetResponseSize() int64 {
	if m != nil && m.ResponseSize != nil {
		return *m.ResponseSize
	}
	return 0
}

func (m *RequestLog) GetReferrer() string {
	if m != nil && m.Referrer != nil {
		return *m.Referrer
	}
	return ""
}

func (m *RequestLog) GetUserAgent() string {
	if m != nil && m.UserAgent != nil {
		return *m.UserAgent
	}
	return ""
}

func (m *RequestLog) GetUrlMapEntry() string {
	if m != nil && m.UrlMapEntry != nil {
		return *m.UrlMapEntry
	}
	return ""
}

func (m *RequestLog) GetCombined() string {
	if m != nil && m.Combined != nil {
		return *m.Combined
	}
	return ""
}

func (m *RequestLog) GetApiMcycles() int64 {
	if m != nil && m.ApiMcycles != nil {
		return *m.ApiMcycles
	}
	return 0
}

func (m *RequestLog) GetHost() string {
	if m != nil && m.Host != nil {
		return *m.Host
	}
	return ""
}

func (m *RequestLog) GetCost() float64 {
	if m != nil && m.Cost != nil {
		return *m.Cost
	}
	return 0
}

func (m *RequestLog) GetTaskQueueName() string {
	if m != nil && m.TaskQueueName != nil {
		return *m.TaskQueueName
	}
	return ""
}

func (m *RequestLog) GetTaskName() string {
	if m != nil && m.TaskName != nil {
		return *m.TaskName
	}
	return ""
}

func (m *RequestLog) GetWasLoadingRequest() bool {
	if m != nil && m.WasLoadingRequest != nil {
		return *m.WasLoadingRequest
	}
	return false
}

func (m *RequestLog) GetPendingTime() int64 {
	if m != nil && m.PendingTime != nil {
		return *m.PendingTime
	}
	return 0
}

func (m *RequestLog) GetReplicaIndex() int32 {
	if m != nil && m.ReplicaIndex != nil {
		return *m.ReplicaIndex
	}
	return Default_RequestLog_ReplicaIndex
}

func (m *RequestLog) GetFinished() bool {
	if m != nil && m.Finished != nil {
		return *m.Finished
	}
	return Default_RequestLog_Finished
}

func (m *RequestLog) GetCloneKey() []byte {
	if m != nil {
		return m.CloneKey
	}
	return nil
}

func (m *RequestLog) GetLine() []*LogLine {
	if m != nil {
		return m.Line
	}
	return nil
}

func (m *RequestLog) GetLinesIncomplete() bool {
	if m != nil && m.LinesIncomplete != nil {
		return *m.LinesIncomplete
	}
	return false
}

func (m *RequestLog) GetAppEngineRelease() []byte {
	if m != nil {
		return m.AppEngineRelease
	}
	return nil
}

func (m *RequestLog) GetTraceId() string {
	if m != nil && m.TraceId != nil {
		return *m.TraceId
	}
	return ""
}

func (m *RequestLog) GetExitReason() int32 {
	if m != nil && m.ExitReason != nil {
		return *m.ExitReason
	}
	return 0
}

func (m *RequestLog) GetWasThrottledForTime() bool {
	if m != nil && m.WasThrottledForTime != nil {
		return *m.WasThrottledForTime
	}
	return false
}

func (m *RequestLog) GetWasThrottledForRequests() bool {
	if m != nil && m.WasThrottledForRequests != nil {
		return *m.WasThrottledForRequests
	}
	return false
}

func (m *RequestLog) GetThrottledTime() int64 {
	if m != nil && m.ThrottledTime != nil {
		return *m.ThrottledTime
	}
	return 0
}

func (m *RequestLog) GetServerName() []byte {
	if m != nil {
		return m.ServerName
	}
	return nil
}

type LogModuleVersion struct {
	ModuleId         *string `protobuf:"bytes,1,opt,name=module_id,def=default" json:"module_id,omitempty"`
	VersionId        *string `protobuf:"bytes,2,opt,name=version_id" json:"version_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *LogModuleVersion) Reset()         { *m = LogModuleVersion{} }
func (m *LogModuleVersion) String() string { return proto.CompactTextString(m) }
func (*LogModuleVersion) ProtoMessage()    {}

const Default_LogModuleVersion_ModuleId string = "default"

func (m *LogModuleVersion) GetModuleId() string {
	if m != nil && m.ModuleId != nil {
		return *m.ModuleId
	}
	return Default_LogModuleVersion_ModuleId
}

func (m *LogModuleVersion) GetVersionId() string {
	if m != nil && m.VersionId != nil {
		return *m.VersionId
	}
	return ""
}

type LogReadRequest struct {
	AppId             *string             `protobuf:"bytes,1,req,name=app_id" json:"app_id,omitempty"`
	VersionId         []string            `protobuf:"bytes,2,rep,name=version_id" json:"version_id,omitempty"`
	ModuleVersion     []*LogModuleVersion `protobuf:"bytes,19,rep,name=module_version" json:"module_version,omitempty"`
	StartTime         *int64              `protobuf:"varint,3,opt,name=start_time" json:"start_time,omitempty"`
	EndTime           *int64              `protobuf:"varint,4,opt,name=end_time" json:"end_time,omitempty"`
	Offset            *LogOffset          `protobuf:"bytes,5,opt,name=offset" json:"offset,omitempty"`
	RequestId         [][]byte            `protobuf:"bytes,6,rep,name=request_id" json:"request_id,omitempty"`
	MinimumLogLevel   *int32              `protobuf:"varint,7,opt,name=minimum_log_level" json:"minimum_log_level,omitempty"`
	IncludeIncomplete *bool               `protobuf:"varint,8,opt,name=include_incomplete" json:"include_incomplete,omitempty"`
	Count             *int64              `protobuf:"varint,9,opt,name=count" json:"count,omitempty"`
	CombinedLogRegex  *string             `protobuf:"bytes,14,opt,name=combined_log_regex" json:"combined_log_regex,omitempty"`
	HostRegex         *string             `protobuf:"bytes,15,opt,name=host_regex" json:"host_regex,omitempty"`
	ReplicaIndex      *int32              `protobuf:"varint,16,opt,name=replica_index" json:"replica_index,omitempty"`
	IncludeAppLogs    *bool               `protobuf:"varint,10,opt,name=include_app_logs" json:"include_app_logs,omitempty"`
	AppLogsPerRequest *int32              `protobuf:"varint,17,opt,name=app_logs_per_request" json:"app_logs_per_request,omitempty"`
	IncludeHost       *bool               `protobuf:"varint,11,opt,name=include_host" json:"include_host,omitempty"`
	IncludeAll        *bool               `protobuf:"varint,12,opt,name=include_all" json:"include_all,omitempty"`
	CacheIterator     *bool               `protobuf:"varint,13,opt,name=cache_iterator" json:"cache_iterator,omitempty"`
	NumShards         *int32              `protobuf:"varint,18,opt,name=num_shards" json:"num_shards,omitempty"`
	XXX_unrecognized  []byte              `json:"-"`
}

func (m *LogReadRequest) Reset()         { *m = LogReadRequest{} }
func (m *LogReadRequest) String() string { return proto.CompactTextString(m) }
func (*LogReadRequest) ProtoMessage()    {}

func (m *LogReadRequest) GetAppId() string {
	if m != nil && m.AppId != nil {
		return *m.AppId
	}
	return ""
}

func (m *LogReadRequest) GetVersionId() []string {
	if m != nil {
		return m.VersionId
	}
	return nil
}

func (m *LogReadRequest) GetModuleVersion() []*LogModuleVersion {
	if m != nil {
		return m.ModuleVersion
	}
	return nil
}

func (m *LogReadRequest) GetStartTime() int64 {
	if m != nil && m.StartTime != nil {
		return *m.StartTime
	}
	return 0
}

func (m *LogReadRequest) GetEndTime() int64 {
	if m != nil && m.EndTime != nil {
		return *m.EndTime
	}
	return 0
}

func (m *LogReadRequest) GetOffset() *LogOffset {
	if m != nil {
		return m.Offset
	}
	return nil
}

func (m *LogReadRequest) GetRequestId() [][]byte {
	if m != nil {
		return m.RequestId
	}
	return nil
}

func (m *LogReadRequest) GetMinimumLogLevel() int32 {
	if m != nil && m.MinimumLogLevel != nil {
		return *m.MinimumLogLevel
	}
	return 0
}

func (m *LogReadRequest) GetIncludeIncomplete() bool {
	if m != nil && m.IncludeIncomplete != nil {
		return *m.IncludeIncomplete
	}
	return false
}

func (m *LogReadRequest) GetCount() int64 {
	if m != nil && m.Count != nil {
		return *m.Count
	}
	return 0
}

func (m *LogReadRequest) GetCombinedLogRegex() string {
	if m != nil && m.CombinedLogRegex != nil {
		return *m.CombinedLogRegex
	}
	return ""
}

func (m *LogReadRequest) GetHostRegex() string {
	if m != nil && m.HostRegex != nil {
		return *m.HostRegex
	}
	return ""
}

func (m *LogReadRequest) GetReplicaIndex() int32 {
	if m != nil && m.ReplicaIndex != nil {
		return *m.ReplicaIndex
	}
	return 0
}

func (m *LogReadRequest) GetIncludeAppLogs() bool {
	if m != nil && m.IncludeAppLogs != nil {
		return *m.IncludeAppLogs
	}
	return false
}

func (m *LogReadRequest) GetAppLogsPerRequest() int32 {
	if m != nil && m.AppLogsPerRequest != nil {
		return *m.AppLogsPerRequest
	}
	return 0
}

func (m *LogReadRequest) GetIncludeHost() bool {
	if m != nil && m.IncludeHost != nil {
		return *m.IncludeHost
	}
	return false
}

func (m *LogReadRequest) GetIncludeAll() bool {
	if m != nil && m.IncludeAll != nil {
		return *m.IncludeAll
	}
	return false
}

func (m *LogReadRequest) GetCacheIterator() bool {
	if m != nil && m.CacheIterator != nil {
		return *m.CacheIterator
	}
	return false
}

func (m *LogReadRequest) GetNumShards() int32 {
	if m != nil && m.NumShards != nil {
		return *m.NumShards
	}
	return 0
}

type LogReadResponse struct {
	Log              []*RequestLog `protobuf:"bytes,1,rep,name=log" json:"log,omitempty"`
	Offset           *LogOffset    `protobuf:"bytes,2,opt,name=offset" json:"offset,omitempty"`
	LastEndTime      *int64        `protobuf:"varint,3,opt,name=last_end_time" json:"last_end_time,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *LogReadResponse) Reset()         { *m = LogReadResponse{} }
func (m *LogReadResponse) String() string { return proto.CompactTextString(m) }
func (*LogReadResponse) ProtoMessage()    {}

func (m *LogReadResponse) GetLog() []*RequestLog {
	if m != nil {
		return m.Log
	}
	return nil
}

func (m *LogReadResponse) GetOffset() *LogOffset {
	if m != nil {
		return m.Offset
	}
	return nil
}

func (m *LogReadResponse) GetLastEndTime() int64 {
	if m != nil && m.LastEndTime != nil {
		return *m.LastEndTime
	}
	return 0
}

type LogUsageRecord struct {
	VersionId        *string `protobuf:"bytes,1,opt,name=version_id" json:"version_id,omitempty"`
	StartTime        *int32  `protobuf:"varint,2,opt,name=start_time" json:"start_time,omitempty"`
	EndTime          *int32  `protobuf:"varint,3,opt,name=end_time" json:"end_time,omitempty"`
	Count            *int64  `protobuf:"varint,4,opt,name=count" json:"count,omitempty"`
	TotalSize        *int64  `protobuf:"varint,5,opt,name=total_size" json:"total_size,omitempty"`
	Records          *int32  `protobuf:"varint,6,opt,name=records" json:"records,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *LogUsageRecord) Reset()         { *m = LogUsageRecord{} }
func (m *LogUsageRecord) String() string { return proto.CompactTextString(m) }
func (*LogUsageRecord) ProtoMessage()    {}

func (m *LogUsageRecord) GetVersionId() string {
	if m != nil && m.VersionId != nil {
		return *m.VersionId
	}
	return ""
}

func (m *LogUsageRecord) GetStartTime() int32 {
	if m != nil && m.StartTime != nil {
		return *m.StartTime
	}
	return 0
}

func (m *LogUsageRecord) GetEndTime() int32 {
	if m != nil && m.EndTime != nil {
		return *m.EndTime
	}
	return 0
}

func (m *LogUsageRecord) GetCount() int64 {
	if m != nil && m.Count != nil {
		return *m.Count
	}
	return 0
}

func (m *LogUsageRecord) GetTotalSize() int64 {
	if m != nil && m.TotalSize != nil {
		return *m.TotalSize
	}
	return 0
}

func (m *LogUsageRecord) GetRecords() int32 {
	if m != nil && m.Records != nil {
		return *m.Records
	}
	return 0
}

type LogUsageRequest struct {
	AppId            *string  `protobuf:"bytes,1,req,name=app_id" json:"app_id,omitempty"`
	VersionId        []string `protobuf:"bytes,2,rep,name=version_id" json:"version_id,omitempty"`
	StartTime        *int32   `protobuf:"varint,3,opt,name=start_time" json:"start_time,omitempty"`
	EndTime          *int32   `protobuf:"varint,4,opt,name=end_time" json:"end_time,omitempty"`
	ResolutionHours  *uint32  `protobuf:"varint,5,opt,name=resolution_hours,def=1" json:"resolution_hours,omitempty"`
	CombineVersions  *bool    `protobuf:"varint,6,opt,name=combine_versions" json:"combine_versions,omitempty"`
	UsageVersion     *int32   `protobuf:"varint,7,opt,name=usage_version" json:"usage_version,omitempty"`
	VersionsOnly     *bool    `protobuf:"varint,8,opt,name=versions_only" json:"versions_only,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *LogUsageRequest) Reset()         { *m = LogUsageRequest{} }
func (m *LogUsageRequest) String() string { return proto.CompactTextString(m) }
func (*LogUsageRequest) ProtoMessage()    {}

const Default_LogUsageRequest_ResolutionHours uint32 = 1

func (m *LogUsageRequest) GetAppId() string {
	if m != nil && m.AppId != nil {
		return *m.AppId
	}
	return ""
}

func (m *LogUsageRequest) GetVersionId() []string {
	if m != nil {
		return m.VersionId
	}
	return nil
}

func (m *LogUsageRequest) GetStartTime() int32 {
	if m != nil && m.StartTime != nil {
		return *m.StartTime
	}
	return 0
}

func (m *LogUsageRequest) GetEndTime() int32 {
	if m != nil && m.EndTime != nil {
		return *m.EndTime
	}
	return 0
}

func (m *LogUsageRequest) GetResolutionHours() uint32 {
	if m != nil && m.ResolutionHours != nil {
		return *m.ResolutionHours
	}
	return Default_LogUsageRequest_ResolutionHours
}

func (m *LogUsageRequest) GetCombineVersions() bool {
	if m != nil && m.CombineVersions != nil {
		return *m.CombineVersions
	}
	return false
}

func (m *LogUsageRequest) GetUsageVersion() int32 {
	if m != nil && m.UsageVersion != nil {
		return *m.UsageVersion
	}
	return 0
}

func (m *LogUsageRequest) GetVersionsOnly() bool {
	if m != nil && m.VersionsOnly != nil {
		return *m.VersionsOnly
	}
	return false
}

type LogUsageResponse struct {
	Usage            []*LogUsageRecord `protobuf:"bytes,1,rep,name=usage" json:"usage,omitempty"`
	Summary          *LogUsageRecord   `protobuf:"bytes,2,opt,name=summary" json:"summary,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *LogUsageResponse) Reset()         { *m = LogUsageResponse{} }
func (m *LogUsageResponse) String() string { return proto.CompactTextString(m) }
func (*LogUsageResponse) ProtoMessage()    {}

func (m *LogUsageResponse) GetUsage() []*LogUsageRecord {
	if m != nil {
		return m.Usage
	}
	return nil
}

func (m *LogUsageResponse) GetSummary() *LogUsageRecord {
	if m != nil {
		return m.Summary
	}
	return nil
}

func init() {
	proto.RegisterEnum("appengine.LogServiceError_ErrorCode", LogServiceError_ErrorCode_name, LogServiceError_ErrorCode_value)
}
