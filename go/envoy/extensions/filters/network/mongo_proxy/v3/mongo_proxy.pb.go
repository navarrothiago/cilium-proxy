// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/network/mongo_proxy/v3/mongo_proxy.proto

package envoy_extensions_filters_network_mongo_proxy_v3

import (
	fmt "fmt"
	v3 "github.com/cilium/proxy/go/envoy/extensions/filters/common/fault/v3"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type MongoProxy struct {
	// The human readable prefix to use when emitting :ref:`statistics
	// <config_network_filters_mongo_proxy_stats>`.
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// The optional path to use for writing Mongo access logs. If not access log
	// path is specified no access logs will be written. Note that access log is
	// also gated :ref:`runtime <config_network_filters_mongo_proxy_runtime>`.
	AccessLog string `protobuf:"bytes,2,opt,name=access_log,json=accessLog,proto3" json:"access_log,omitempty"`
	// Inject a fixed delay before proxying a Mongo operation. Delays are
	// applied to the following MongoDB operations: Query, Insert, GetMore,
	// and KillCursors. Once an active delay is in progress, all incoming
	// data up until the timer event fires will be a part of the delay.
	Delay *v3.FaultDelay `protobuf:"bytes,3,opt,name=delay,proto3" json:"delay,omitempty"`
	// Flag to specify whether :ref:`dynamic metadata
	// <config_network_filters_mongo_proxy_dynamic_metadata>` should be emitted. Defaults to false.
	EmitDynamicMetadata  bool     `protobuf:"varint,4,opt,name=emit_dynamic_metadata,json=emitDynamicMetadata,proto3" json:"emit_dynamic_metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MongoProxy) Reset()         { *m = MongoProxy{} }
func (m *MongoProxy) String() string { return proto.CompactTextString(m) }
func (*MongoProxy) ProtoMessage()    {}
func (*MongoProxy) Descriptor() ([]byte, []int) {
	return fileDescriptor_60ec9b7c3b00e562, []int{0}
}

func (m *MongoProxy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MongoProxy.Unmarshal(m, b)
}
func (m *MongoProxy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MongoProxy.Marshal(b, m, deterministic)
}
func (m *MongoProxy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MongoProxy.Merge(m, src)
}
func (m *MongoProxy) XXX_Size() int {
	return xxx_messageInfo_MongoProxy.Size(m)
}
func (m *MongoProxy) XXX_DiscardUnknown() {
	xxx_messageInfo_MongoProxy.DiscardUnknown(m)
}

var xxx_messageInfo_MongoProxy proto.InternalMessageInfo

func (m *MongoProxy) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *MongoProxy) GetAccessLog() string {
	if m != nil {
		return m.AccessLog
	}
	return ""
}

func (m *MongoProxy) GetDelay() *v3.FaultDelay {
	if m != nil {
		return m.Delay
	}
	return nil
}

func (m *MongoProxy) GetEmitDynamicMetadata() bool {
	if m != nil {
		return m.EmitDynamicMetadata
	}
	return false
}

func init() {
	proto.RegisterType((*MongoProxy)(nil), "envoy.extensions.filters.network.mongo_proxy.v3.MongoProxy")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/network/mongo_proxy/v3/mongo_proxy.proto", fileDescriptor_60ec9b7c3b00e562)
}

var fileDescriptor_60ec9b7c3b00e562 = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0xe9, 0x9c, 0x7f, 0x96, 0x1d, 0x84, 0x8a, 0x58, 0x06, 0x4a, 0xf5, 0xd4, 0x53, 0x02,
	0xeb, 0x44, 0x10, 0x3d, 0x58, 0x86, 0x07, 0x71, 0x50, 0xfa, 0x05, 0x4a, 0x6c, 0xd3, 0x12, 0x6c,
	0xf3, 0x96, 0x34, 0xab, 0xeb, 0x37, 0xf0, 0xee, 0xcd, 0xef, 0xe3, 0x97, 0xf2, 0x24, 0x69, 0x2a,
	0xdd, 0xc1, 0x1d, 0xbc, 0xbd, 0xed, 0x2f, 0xcf, 0xf3, 0xbc, 0x79, 0x82, 0x1e, 0x98, 0x68, 0xa0,
	0x25, 0x6c, 0xa3, 0x98, 0xa8, 0x39, 0x88, 0x9a, 0x64, 0xbc, 0x50, 0x4c, 0xd6, 0x44, 0x30, 0xf5,
	0x06, 0xf2, 0x95, 0x94, 0x20, 0x72, 0x88, 0x2b, 0x09, 0x9b, 0x96, 0x34, 0xfe, 0xf6, 0x27, 0xae,
	0x24, 0x28, 0xb0, 0x49, 0x67, 0x81, 0x07, 0x0b, 0xdc, 0x5b, 0xe0, 0xde, 0x02, 0x6f, 0x6b, 0x1a,
	0x7f, 0xb6, 0xd8, 0x99, 0x99, 0x40, 0x59, 0x82, 0x20, 0x19, 0x5d, 0x17, 0x4a, 0x87, 0x75, 0x83,
	0x89, 0x99, 0x5d, 0xae, 0xd3, 0x8a, 0x12, 0x2a, 0x04, 0x28, 0xaa, 0x3a, 0x55, 0xc3, 0xa4, 0x96,
	0x73, 0x91, 0xf7, 0x47, 0xce, 0x1a, 0x5a, 0xf0, 0x94, 0x2a, 0x46, 0x7e, 0x07, 0x03, 0xae, 0x3e,
	0x46, 0x08, 0xad, 0xf4, 0x12, 0xa1, 0xde, 0xc1, 0xf6, 0xd0, 0xb4, 0x56, 0x54, 0xc5, 0x95, 0x64,
	0x19, 0xdf, 0x38, 0x96, 0x6b, 0x79, 0x93, 0xe0, 0xf0, 0x3b, 0x18, 0xcb, 0x91, 0x6b, 0x45, 0x48,
	0xb3, 0xb0, 0x43, 0xf6, 0x39, 0x42, 0x34, 0x49, 0x58, 0x5d, 0xc7, 0x05, 0xe4, 0xce, 0x48, 0x1f,
	0x8c, 0x26, 0xe6, 0xcf, 0x33, 0xe4, 0xf6, 0x13, 0xda, 0x4f, 0x59, 0x41, 0x5b, 0x67, 0xcf, 0xb5,
	0xbc, 0xe9, 0x7c, 0x81, 0x77, 0x56, 0x61, 0x6e, 0x86, 0xcd, 0x85, 0x1a, 0x1f, 0x3f, 0xea, 0x61,
	0xa9, 0xb5, 0x91, 0xb1, 0xb0, 0xe7, 0xe8, 0x94, 0x95, 0x5c, 0xc5, 0x69, 0x2b, 0x68, 0xc9, 0x93,
	0xb8, 0x64, 0x8a, 0xa6, 0x54, 0x51, 0x67, 0xec, 0x5a, 0xde, 0x51, 0x74, 0xa2, 0xe1, 0xd2, 0xb0,
	0x55, 0x8f, 0x6e, 0xef, 0x3e, 0xbf, 0xde, 0x2f, 0x6e, 0xd0, 0xb5, 0x89, 0x4d, 0x40, 0x64, 0x3c,
	0xef, 0x23, 0xff, 0x2e, 0x7f, 0x8e, 0x87, 0x1a, 0x82, 0x08, 0xdd, 0x73, 0x30, 0x2b, 0x1b, 0xfc,
	0xcf, 0x87, 0x0c, 0x8e, 0x07, 0xb3, 0x50, 0xf7, 0x1c, 0x5a, 0x2f, 0x07, 0x5d, 0xe1, 0xfe, 0x4f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xea, 0x16, 0x02, 0xc0, 0x58, 0x02, 0x00, 0x00,
}