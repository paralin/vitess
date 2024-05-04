// Code generated by protoc-gen-go-lite. DO NOT EDIT.
// protoc-gen-go-lite version: v0.6.1
// source: github.com/dolthub/vitess/proto/dolthub/vt/vtrpc.proto

package vtrpc

import (
	fmt "fmt"
	protobuf_go_lite "github.com/aperturerobotics/protobuf-go-lite"
	json "github.com/aperturerobotics/protobuf-go-lite/json"
	io "io"
	strconv "strconv"
	strings "strings"
)

//
//Copyright 2019 The Vitess Authors.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// This file contains useful data structures for RPCs in Vitess.

// Code represents canonical error codes. The names, numbers and comments
// must match the ones defined by grpc:
// https://godoc.org/google.golang.org/grpc/codes.
type Code int32

const (
	// OK is returned on success.
	Code_OK Code = 0
	// CANCELED indicates the operation was cancelled (typically by the caller).
	Code_CANCELED Code = 1
	// UNKNOWN error. An example of where this error may be returned is
	// if a Status value received from another address space belongs to
	// an error-space that is not known in this address space. Also
	// errors raised by APIs that do not return enough error information
	// may be converted to this error.
	Code_UNKNOWN Code = 2
	// INVALID_ARGUMENT indicates client specified an invalid argument.
	// Note that this differs from FAILED_PRECONDITION. It indicates arguments
	// that are problematic regardless of the state of the system
	// (e.g., a malformed file name).
	Code_INVALID_ARGUMENT Code = 3
	// DEADLINE_EXCEEDED means operation expired before completion.
	// For operations that change the state of the system, this error may be
	// returned even if the operation has completed successfully. For
	// example, a successful response from a server could have been delayed
	// long enough for the deadline to expire.
	Code_DEADLINE_EXCEEDED Code = 4
	// NOT_FOUND means some requested entity (e.g., file or directory) was
	// not found.
	Code_NOT_FOUND Code = 5
	// ALREADY_EXISTS means an attempt to create an entity failed because one
	// already exists.
	Code_ALREADY_EXISTS Code = 6
	// PERMISSION_DENIED indicates the caller does not have permission to
	// execute the specified operation. It must not be used for rejections
	// caused by exhausting some resource (use RESOURCE_EXHAUSTED
	// instead for those errors).  It must not be
	// used if the caller cannot be identified (use Unauthenticated
	// instead for those errors).
	Code_PERMISSION_DENIED Code = 7
	// UNAUTHENTICATED indicates the request does not have valid
	// authentication credentials for the operation.
	Code_UNAUTHENTICATED Code = 16
	// RESOURCE_EXHAUSTED indicates some resource has been exhausted, perhaps
	// a per-user quota, or perhaps the entire file system is out of space.
	Code_RESOURCE_EXHAUSTED Code = 8
	// FAILED_PRECONDITION indicates operation was rejected because the
	// system is not in a state required for the operation's execution.
	// For example, directory to be deleted may be non-empty, an rmdir
	// operation is applied to a non-directory, etc.
	//
	// A litmus test that may help a service implementor in deciding
	// between FAILED_PRECONDITION, ABORTED, and UNAVAILABLE:
	//
	//	(a) Use UNAVAILABLE if the client can retry just the failing call.
	//	(b) Use ABORTED if the client should retry at a higher-level
	//	    (e.g., restarting a read-modify-write sequence).
	//	(c) Use FAILED_PRECONDITION if the client should not retry until
	//	    the system state has been explicitly fixed.  E.g., if an "rmdir"
	//	    fails because the directory is non-empty, FAILED_PRECONDITION
	//	    should be returned since the client should not retry unless
	//	    they have first fixed up the directory by deleting files from it.
	//	(d) Use FAILED_PRECONDITION if the client performs conditional
	//	    REST Get/Update/Delete on a resource and the resource on the
	//	    server does not match the condition. E.g., conflicting
	//	    read-modify-write on the same resource.
	Code_FAILED_PRECONDITION Code = 9
	// ABORTED indicates the operation was aborted, typically due to a
	// concurrency issue like sequencer check failures, transaction aborts,
	// etc.
	//
	// See litmus test above for deciding between FAILED_PRECONDITION,
	// ABORTED, and UNAVAILABLE.
	Code_ABORTED Code = 10
	// OUT_OF_RANGE means operation was attempted past the valid range.
	// E.g., seeking or reading past end of file.
	//
	// Unlike INVALID_ARGUMENT, this error indicates a problem that may
	// be fixed if the system state changes. For example, a 32-bit file
	// system will generate INVALID_ARGUMENT if asked to read at an
	// offset that is not in the range [0,2^32-1], but it will generate
	// OUT_OF_RANGE if asked to read from an offset past the current
	// file size.
	//
	// There is a fair bit of overlap between FAILED_PRECONDITION and
	// OUT_OF_RANGE.  We recommend using OUT_OF_RANGE (the more specific
	// error) when it applies so that callers who are iterating through
	// a space can easily look for an OUT_OF_RANGE error to detect when
	// they are done.
	Code_OUT_OF_RANGE Code = 11
	// UNIMPLEMENTED indicates operation is not implemented or not
	// supported/enabled in this service.
	Code_UNIMPLEMENTED Code = 12
	// INTERNAL errors. Means some invariants expected by underlying
	// system has been broken.  If you see one of these errors,
	// something is very broken.
	Code_INTERNAL Code = 13
	// UNAVAILABLE indicates the service is currently unavailable.
	// This is a most likely a transient condition and may be corrected
	// by retrying with a backoff.
	//
	// See litmus test above for deciding between FAILED_PRECONDITION,
	// ABORTED, and UNAVAILABLE.
	Code_UNAVAILABLE Code = 14
	// DATA_LOSS indicates unrecoverable data loss or corruption.
	Code_DATA_LOSS Code = 15
)

// Enum value maps for Code.
var (
	Code_name = map[int32]string{
		0:  "OK",
		1:  "CANCELED",
		2:  "UNKNOWN",
		3:  "INVALID_ARGUMENT",
		4:  "DEADLINE_EXCEEDED",
		5:  "NOT_FOUND",
		6:  "ALREADY_EXISTS",
		7:  "PERMISSION_DENIED",
		16: "UNAUTHENTICATED",
		8:  "RESOURCE_EXHAUSTED",
		9:  "FAILED_PRECONDITION",
		10: "ABORTED",
		11: "OUT_OF_RANGE",
		12: "UNIMPLEMENTED",
		13: "INTERNAL",
		14: "UNAVAILABLE",
		15: "DATA_LOSS",
	}
	Code_value = map[string]int32{
		"OK":                  0,
		"CANCELED":            1,
		"UNKNOWN":             2,
		"INVALID_ARGUMENT":    3,
		"DEADLINE_EXCEEDED":   4,
		"NOT_FOUND":           5,
		"ALREADY_EXISTS":      6,
		"PERMISSION_DENIED":   7,
		"UNAUTHENTICATED":     16,
		"RESOURCE_EXHAUSTED":  8,
		"FAILED_PRECONDITION": 9,
		"ABORTED":             10,
		"OUT_OF_RANGE":        11,
		"UNIMPLEMENTED":       12,
		"INTERNAL":            13,
		"UNAVAILABLE":         14,
		"DATA_LOSS":           15,
	}
)

func (x Code) Enum() *Code {
	p := new(Code)
	*p = x
	return p
}

func (x Code) String() string {
	name, valid := Code_name[int32(x)]
	if valid {
		return name
	}
	return strconv.Itoa(int(x))
}

// LegacyErrorCode is the enum values for Errors. This type is deprecated.
// Use Code instead. Background: In the initial design, we thought
// that we may end up with a different list of canonical error codes
// than the ones defined by grpc. In hindsight, we realize that
// the grpc error codes are fairly generic and mostly sufficient.
// In order to avoid confusion, this type will be deprecated in
// favor of the new Code that matches exactly what grpc defines.
// Some names below have a _LEGACY suffix. This is to prevent
// name collisions with Code.
type LegacyErrorCode int32

const (
	// SUCCESS_LEGACY is returned from a successful call.
	LegacyErrorCode_SUCCESS_LEGACY LegacyErrorCode = 0
	// CANCELLED_LEGACY means that the context was cancelled (and noticed in the app layer,
	// as opposed to the RPC layer).
	LegacyErrorCode_CANCELLED_LEGACY LegacyErrorCode = 1
	// UNKNOWN_ERROR_LEGACY includes:
	//  1. MySQL error codes that we don't explicitly handle.
	//  2. MySQL response that wasn't as expected. For example, we might expect a MySQL
	//     timestamp to be returned in a particular way, but it wasn't.
	//  3. Anything else that doesn't fall into a different bucket.
	LegacyErrorCode_UNKNOWN_ERROR_LEGACY LegacyErrorCode = 2
	// BAD_INPUT_LEGACY is returned when an end-user either sends SQL that couldn't be parsed correctly,
	// or tries a query that isn't supported by Vitess.
	LegacyErrorCode_BAD_INPUT_LEGACY LegacyErrorCode = 3
	// DEADLINE_EXCEEDED_LEGACY is returned when an action is taking longer than a given timeout.
	LegacyErrorCode_DEADLINE_EXCEEDED_LEGACY LegacyErrorCode = 4
	// INTEGRITY_ERROR_LEGACY is returned on integrity error from MySQL, usually due to
	// duplicate primary keys.
	LegacyErrorCode_INTEGRITY_ERROR_LEGACY LegacyErrorCode = 5
	// PERMISSION_DENIED_LEGACY errors are returned when a user requests access to something
	// that they don't have permissions for.
	LegacyErrorCode_PERMISSION_DENIED_LEGACY LegacyErrorCode = 6
	// RESOURCE_EXHAUSTED_LEGACY is returned when a query exceeds its quota in some dimension
	// and can't be completed due to that. Queries that return RESOURCE_EXHAUSTED
	// should not be retried, as it could be detrimental to the server's health.
	// Examples of errors that will cause the RESOURCE_EXHAUSTED code:
	//  1. TxPoolFull: this is retried server-side, and is only returned as an error
	//     if the server-side retries failed.
	//  2. Query is killed due to it taking too long.
	LegacyErrorCode_RESOURCE_EXHAUSTED_LEGACY LegacyErrorCode = 7
	// QUERY_NOT_SERVED_LEGACY means that a query could not be served right now.
	// Client can interpret it as: "the tablet that you sent this query to cannot
	// serve the query right now, try a different tablet or try again later."
	// This could be due to various reasons: QueryService is not serving, should
	// not be serving, wrong shard, wrong tablet type, blacklisted table, etc.
	// Clients that receive this error should usually retry the query, but after taking
	// the appropriate steps to make sure that the query will get sent to the correct
	// tablet.
	LegacyErrorCode_QUERY_NOT_SERVED_LEGACY LegacyErrorCode = 8
	// NOT_IN_TX_LEGACY means that we're not currently in a transaction, but we should be.
	LegacyErrorCode_NOT_IN_TX_LEGACY LegacyErrorCode = 9
	// INTERNAL_ERROR_LEGACY means some invariants expected by underlying
	// system has been broken.  If you see one of these errors,
	// something is very broken.
	LegacyErrorCode_INTERNAL_ERROR_LEGACY LegacyErrorCode = 10
	// TRANSIENT_ERROR_LEGACY is used for when there is some error that we expect we can
	// recover from automatically - often due to a resource limit temporarily being
	// reached. Retrying this error, with an exponential backoff, should succeed.
	// Clients should be able to successfully retry the query on the same backends.
	// Examples of things that can trigger this error:
	// 1. Query has been throttled
	// 2. VtGate could have request backlog
	LegacyErrorCode_TRANSIENT_ERROR_LEGACY LegacyErrorCode = 11
	// UNAUTHENTICATED_LEGACY errors are returned when a user requests access to something,
	// and we're unable to verify the user's authentication.
	LegacyErrorCode_UNAUTHENTICATED_LEGACY LegacyErrorCode = 12
)

// Enum value maps for LegacyErrorCode.
var (
	LegacyErrorCode_name = map[int32]string{
		0:  "SUCCESS_LEGACY",
		1:  "CANCELLED_LEGACY",
		2:  "UNKNOWN_ERROR_LEGACY",
		3:  "BAD_INPUT_LEGACY",
		4:  "DEADLINE_EXCEEDED_LEGACY",
		5:  "INTEGRITY_ERROR_LEGACY",
		6:  "PERMISSION_DENIED_LEGACY",
		7:  "RESOURCE_EXHAUSTED_LEGACY",
		8:  "QUERY_NOT_SERVED_LEGACY",
		9:  "NOT_IN_TX_LEGACY",
		10: "INTERNAL_ERROR_LEGACY",
		11: "TRANSIENT_ERROR_LEGACY",
		12: "UNAUTHENTICATED_LEGACY",
	}
	LegacyErrorCode_value = map[string]int32{
		"SUCCESS_LEGACY":            0,
		"CANCELLED_LEGACY":          1,
		"UNKNOWN_ERROR_LEGACY":      2,
		"BAD_INPUT_LEGACY":          3,
		"DEADLINE_EXCEEDED_LEGACY":  4,
		"INTEGRITY_ERROR_LEGACY":    5,
		"PERMISSION_DENIED_LEGACY":  6,
		"RESOURCE_EXHAUSTED_LEGACY": 7,
		"QUERY_NOT_SERVED_LEGACY":   8,
		"NOT_IN_TX_LEGACY":          9,
		"INTERNAL_ERROR_LEGACY":     10,
		"TRANSIENT_ERROR_LEGACY":    11,
		"UNAUTHENTICATED_LEGACY":    12,
	}
)

func (x LegacyErrorCode) Enum() *LegacyErrorCode {
	p := new(LegacyErrorCode)
	*p = x
	return p
}

func (x LegacyErrorCode) String() string {
	name, valid := LegacyErrorCode_name[int32(x)]
	if valid {
		return name
	}
	return strconv.Itoa(int(x))
}

// CallerID is passed along RPCs to identify the originating client
// for a request. It is not meant to be secure, but only
// informational.  The client can put whatever info they want in these
// fields, and they will be trusted by the servers. The fields will
// just be used for logging purposes, and to easily find a client.
// VtGate propagates it to VtTablet, and VtTablet may use this
// information for monitoring purposes, to display on dashboards, or
// for blacklisting purposes.
type CallerID struct {
	unknownFields []byte
	// principal is the effective user identifier. It is usually filled in
	// with whoever made the request to the appserver, if the request
	// came from an automated job or another system component.
	// If the request comes directly from the Internet, or if the Vitess client
	// takes action on its own accord, it is okay for this field to be absent.
	Principal string `protobuf:"bytes,1,opt,name=principal,proto3" json:"principal,omitempty"`
	// component describes the running process of the effective caller.
	// It can for instance be the hostname:port of the servlet initiating the
	// database call, or the container engine ID used by the servlet.
	Component string `protobuf:"bytes,2,opt,name=component,proto3" json:"component,omitempty"`
	// subcomponent describes a component inisde the immediate caller which
	// is responsible for generating is request. Suggested values are a
	// servlet name or an API endpoint name.
	Subcomponent string `protobuf:"bytes,3,opt,name=subcomponent,proto3" json:"subcomponent,omitempty"`
}

func (x *CallerID) Reset() {
	*x = CallerID{}
}

func (*CallerID) ProtoMessage() {}

func (x *CallerID) GetPrincipal() string {
	if x != nil {
		return x.Principal
	}
	return ""
}

func (x *CallerID) GetComponent() string {
	if x != nil {
		return x.Component
	}
	return ""
}

func (x *CallerID) GetSubcomponent() string {
	if x != nil {
		return x.Subcomponent
	}
	return ""
}

// RPCError is an application-level error structure returned by
// VtTablet (and passed along by VtGate if appropriate).
// We use this so the clients don't have to parse the error messages,
// but instead can depend on the value of the code.
type RPCError struct {
	unknownFields []byte
	LegacyCode    LegacyErrorCode `protobuf:"varint,1,opt,name=legacy_code,json=legacyCode,proto3" json:"legacyCode,omitempty"`
	Message       string          `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Code          Code            `protobuf:"varint,3,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *RPCError) Reset() {
	*x = RPCError{}
}

func (*RPCError) ProtoMessage() {}

func (x *RPCError) GetLegacyCode() LegacyErrorCode {
	if x != nil {
		return x.LegacyCode
	}
	return LegacyErrorCode_SUCCESS_LEGACY
}

func (x *RPCError) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *RPCError) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_OK
}

func (m *CallerID) CloneVT() *CallerID {
	if m == nil {
		return (*CallerID)(nil)
	}
	r := new(CallerID)
	r.Principal = m.Principal
	r.Component = m.Component
	r.Subcomponent = m.Subcomponent
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *CallerID) CloneMessageVT() protobuf_go_lite.CloneMessage {
	return m.CloneVT()
}

func (m *RPCError) CloneVT() *RPCError {
	if m == nil {
		return (*RPCError)(nil)
	}
	r := new(RPCError)
	r.LegacyCode = m.LegacyCode
	r.Message = m.Message
	r.Code = m.Code
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *RPCError) CloneMessageVT() protobuf_go_lite.CloneMessage {
	return m.CloneVT()
}

func (this *CallerID) EqualVT(that *CallerID) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Principal != that.Principal {
		return false
	}
	if this.Component != that.Component {
		return false
	}
	if this.Subcomponent != that.Subcomponent {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *CallerID) EqualMessageVT(thatMsg any) bool {
	that, ok := thatMsg.(*CallerID)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *RPCError) EqualVT(that *RPCError) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.LegacyCode != that.LegacyCode {
		return false
	}
	if this.Message != that.Message {
		return false
	}
	if this.Code != that.Code {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *RPCError) EqualMessageVT(thatMsg any) bool {
	that, ok := thatMsg.(*RPCError)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}

// MarshalProtoJSON marshals the Code to JSON.
func (x Code) MarshalProtoJSON(s *json.MarshalState) {
	s.WriteEnumString(int32(x), Code_name)
}

// MarshalText marshals the Code to text.
func (x Code) MarshalText() ([]byte, error) {
	return []byte(json.GetEnumString(int32(x), Code_name)), nil
}

// MarshalJSON marshals the Code to JSON.
func (x Code) MarshalJSON() ([]byte, error) {
	return json.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the Code from JSON.
func (x *Code) UnmarshalProtoJSON(s *json.UnmarshalState) {
	v := s.ReadEnum(Code_value)
	if err := s.Err(); err != nil {
		s.SetErrorf("could not read Code enum: %v", err)
		return
	}
	*x = Code(v)
}

// UnmarshalText unmarshals the Code from text.
func (x *Code) UnmarshalText(b []byte) error {
	i, err := json.ParseEnumString(string(b), Code_value)
	if err != nil {
		return err
	}
	*x = Code(i)
	return nil
}

// UnmarshalJSON unmarshals the Code from JSON.
func (x *Code) UnmarshalJSON(b []byte) error {
	return json.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the LegacyErrorCode to JSON.
func (x LegacyErrorCode) MarshalProtoJSON(s *json.MarshalState) {
	s.WriteEnumString(int32(x), LegacyErrorCode_name)
}

// MarshalText marshals the LegacyErrorCode to text.
func (x LegacyErrorCode) MarshalText() ([]byte, error) {
	return []byte(json.GetEnumString(int32(x), LegacyErrorCode_name)), nil
}

// MarshalJSON marshals the LegacyErrorCode to JSON.
func (x LegacyErrorCode) MarshalJSON() ([]byte, error) {
	return json.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the LegacyErrorCode from JSON.
func (x *LegacyErrorCode) UnmarshalProtoJSON(s *json.UnmarshalState) {
	v := s.ReadEnum(LegacyErrorCode_value)
	if err := s.Err(); err != nil {
		s.SetErrorf("could not read LegacyErrorCode enum: %v", err)
		return
	}
	*x = LegacyErrorCode(v)
}

// UnmarshalText unmarshals the LegacyErrorCode from text.
func (x *LegacyErrorCode) UnmarshalText(b []byte) error {
	i, err := json.ParseEnumString(string(b), LegacyErrorCode_value)
	if err != nil {
		return err
	}
	*x = LegacyErrorCode(i)
	return nil
}

// UnmarshalJSON unmarshals the LegacyErrorCode from JSON.
func (x *LegacyErrorCode) UnmarshalJSON(b []byte) error {
	return json.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the CallerID message to JSON.
func (x *CallerID) MarshalProtoJSON(s *json.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.Principal != "" || s.HasField("principal") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("principal")
		s.WriteString(x.Principal)
	}
	if x.Component != "" || s.HasField("component") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("component")
		s.WriteString(x.Component)
	}
	if x.Subcomponent != "" || s.HasField("subcomponent") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("subcomponent")
		s.WriteString(x.Subcomponent)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the CallerID to JSON.
func (x *CallerID) MarshalJSON() ([]byte, error) {
	return json.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the CallerID message from JSON.
func (x *CallerID) UnmarshalProtoJSON(s *json.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.Skip() // ignore unknown field
		case "principal":
			s.AddField("principal")
			x.Principal = s.ReadString()
		case "component":
			s.AddField("component")
			x.Component = s.ReadString()
		case "subcomponent":
			s.AddField("subcomponent")
			x.Subcomponent = s.ReadString()
		}
	})
}

// UnmarshalJSON unmarshals the CallerID from JSON.
func (x *CallerID) UnmarshalJSON(b []byte) error {
	return json.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the RPCError message to JSON.
func (x *RPCError) MarshalProtoJSON(s *json.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.LegacyCode != 0 || s.HasField("legacyCode") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("legacyCode")
		x.LegacyCode.MarshalProtoJSON(s)
	}
	if x.Message != "" || s.HasField("message") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("message")
		s.WriteString(x.Message)
	}
	if x.Code != 0 || s.HasField("code") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("code")
		x.Code.MarshalProtoJSON(s)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the RPCError to JSON.
func (x *RPCError) MarshalJSON() ([]byte, error) {
	return json.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the RPCError message from JSON.
func (x *RPCError) UnmarshalProtoJSON(s *json.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.Skip() // ignore unknown field
		case "legacy_code", "legacyCode":
			s.AddField("legacy_code")
			x.LegacyCode.UnmarshalProtoJSON(s)
		case "message":
			s.AddField("message")
			x.Message = s.ReadString()
		case "code":
			s.AddField("code")
			x.Code.UnmarshalProtoJSON(s)
		}
	})
}

// UnmarshalJSON unmarshals the RPCError from JSON.
func (x *RPCError) UnmarshalJSON(b []byte) error {
	return json.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

func (m *CallerID) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CallerID) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *CallerID) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if len(m.Subcomponent) > 0 {
		i -= len(m.Subcomponent)
		copy(dAtA[i:], m.Subcomponent)
		i = protobuf_go_lite.EncodeVarint(dAtA, i, uint64(len(m.Subcomponent)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Component) > 0 {
		i -= len(m.Component)
		copy(dAtA[i:], m.Component)
		i = protobuf_go_lite.EncodeVarint(dAtA, i, uint64(len(m.Component)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Principal) > 0 {
		i -= len(m.Principal)
		copy(dAtA[i:], m.Principal)
		i = protobuf_go_lite.EncodeVarint(dAtA, i, uint64(len(m.Principal)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RPCError) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RPCError) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *RPCError) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if m.Code != 0 {
		i = protobuf_go_lite.EncodeVarint(dAtA, i, uint64(m.Code))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = protobuf_go_lite.EncodeVarint(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0x12
	}
	if m.LegacyCode != 0 {
		i = protobuf_go_lite.EncodeVarint(dAtA, i, uint64(m.LegacyCode))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *CallerID) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Principal)
	if l > 0 {
		n += 1 + l + protobuf_go_lite.SizeOfVarint(uint64(l))
	}
	l = len(m.Component)
	if l > 0 {
		n += 1 + l + protobuf_go_lite.SizeOfVarint(uint64(l))
	}
	l = len(m.Subcomponent)
	if l > 0 {
		n += 1 + l + protobuf_go_lite.SizeOfVarint(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *RPCError) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LegacyCode != 0 {
		n += 1 + protobuf_go_lite.SizeOfVarint(uint64(m.LegacyCode))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + protobuf_go_lite.SizeOfVarint(uint64(l))
	}
	if m.Code != 0 {
		n += 1 + protobuf_go_lite.SizeOfVarint(uint64(m.Code))
	}
	n += len(m.unknownFields)
	return n
}

func (x Code) MarshalProtoText() string {
	return x.String()
}
func (x LegacyErrorCode) MarshalProtoText() string {
	return x.String()
}
func (x *CallerID) MarshalProtoText() string {
	var sb strings.Builder
	sb.WriteString("CallerID { ")
	if x.Principal != "" {
		sb.WriteString(" principal: ")
		sb.WriteString(strconv.Quote(x.Principal))
	}
	if x.Component != "" {
		sb.WriteString(" component: ")
		sb.WriteString(strconv.Quote(x.Component))
	}
	if x.Subcomponent != "" {
		sb.WriteString(" subcomponent: ")
		sb.WriteString(strconv.Quote(x.Subcomponent))
	}
	sb.WriteString("}")
	return sb.String()
}
func (x *CallerID) String() string {
	return x.MarshalProtoText()
}
func (x *RPCError) MarshalProtoText() string {
	var sb strings.Builder
	sb.WriteString("RPCError { ")
	if x.LegacyCode != 0 {
		sb.WriteString(" legacy_code: ")
		sb.WriteString(LegacyErrorCode(x.LegacyCode).String())
	}
	if x.Message != "" {
		sb.WriteString(" message: ")
		sb.WriteString(strconv.Quote(x.Message))
	}
	if x.Code != 0 {
		sb.WriteString(" code: ")
		sb.WriteString(Code(x.Code).String())
	}
	sb.WriteString("}")
	return sb.String()
}
func (x *RPCError) String() string {
	return x.MarshalProtoText()
}
func (m *CallerID) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return protobuf_go_lite.ErrIntOverflow
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CallerID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CallerID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Principal", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protobuf_go_lite.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return protobuf_go_lite.ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return protobuf_go_lite.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Principal = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Component", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protobuf_go_lite.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return protobuf_go_lite.ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return protobuf_go_lite.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Component = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subcomponent", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protobuf_go_lite.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return protobuf_go_lite.ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return protobuf_go_lite.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subcomponent = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := protobuf_go_lite.Skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return protobuf_go_lite.ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RPCError) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return protobuf_go_lite.ErrIntOverflow
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RPCError: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RPCError: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LegacyCode", wireType)
			}
			m.LegacyCode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protobuf_go_lite.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LegacyCode |= LegacyErrorCode(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protobuf_go_lite.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return protobuf_go_lite.ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return protobuf_go_lite.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protobuf_go_lite.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= Code(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := protobuf_go_lite.Skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return protobuf_go_lite.ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
