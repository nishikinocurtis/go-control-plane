// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.1
// source: envoy/type/v3/ratelimit_strategy.proto

package typev3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/cncf/xds/go/xds/annotations/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Choose between allow all and deny all.
type RateLimitStrategy_BlanketRule int32

const (
	RateLimitStrategy_ALLOW_ALL RateLimitStrategy_BlanketRule = 0
	RateLimitStrategy_DENY_ALL  RateLimitStrategy_BlanketRule = 1
)

// Enum value maps for RateLimitStrategy_BlanketRule.
var (
	RateLimitStrategy_BlanketRule_name = map[int32]string{
		0: "ALLOW_ALL",
		1: "DENY_ALL",
	}
	RateLimitStrategy_BlanketRule_value = map[string]int32{
		"ALLOW_ALL": 0,
		"DENY_ALL":  1,
	}
)

func (x RateLimitStrategy_BlanketRule) Enum() *RateLimitStrategy_BlanketRule {
	p := new(RateLimitStrategy_BlanketRule)
	*p = x
	return p
}

func (x RateLimitStrategy_BlanketRule) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RateLimitStrategy_BlanketRule) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_type_v3_ratelimit_strategy_proto_enumTypes[0].Descriptor()
}

func (RateLimitStrategy_BlanketRule) Type() protoreflect.EnumType {
	return &file_envoy_type_v3_ratelimit_strategy_proto_enumTypes[0]
}

func (x RateLimitStrategy_BlanketRule) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RateLimitStrategy_BlanketRule.Descriptor instead.
func (RateLimitStrategy_BlanketRule) EnumDescriptor() ([]byte, []int) {
	return file_envoy_type_v3_ratelimit_strategy_proto_rawDescGZIP(), []int{0, 0}
}

type RateLimitStrategy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Strategy:
	//	*RateLimitStrategy_BlanketRule_
	//	*RateLimitStrategy_RequestsPerTimeUnit_
	//	*RateLimitStrategy_TokenBucket
	Strategy isRateLimitStrategy_Strategy `protobuf_oneof:"strategy"`
}

func (x *RateLimitStrategy) Reset() {
	*x = RateLimitStrategy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_type_v3_ratelimit_strategy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateLimitStrategy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateLimitStrategy) ProtoMessage() {}

func (x *RateLimitStrategy) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_type_v3_ratelimit_strategy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateLimitStrategy.ProtoReflect.Descriptor instead.
func (*RateLimitStrategy) Descriptor() ([]byte, []int) {
	return file_envoy_type_v3_ratelimit_strategy_proto_rawDescGZIP(), []int{0}
}

func (m *RateLimitStrategy) GetStrategy() isRateLimitStrategy_Strategy {
	if m != nil {
		return m.Strategy
	}
	return nil
}

func (x *RateLimitStrategy) GetBlanketRule() RateLimitStrategy_BlanketRule {
	if x, ok := x.GetStrategy().(*RateLimitStrategy_BlanketRule_); ok {
		return x.BlanketRule
	}
	return RateLimitStrategy_ALLOW_ALL
}

func (x *RateLimitStrategy) GetRequestsPerTimeUnit() *RateLimitStrategy_RequestsPerTimeUnit {
	if x, ok := x.GetStrategy().(*RateLimitStrategy_RequestsPerTimeUnit_); ok {
		return x.RequestsPerTimeUnit
	}
	return nil
}

func (x *RateLimitStrategy) GetTokenBucket() *TokenBucket {
	if x, ok := x.GetStrategy().(*RateLimitStrategy_TokenBucket); ok {
		return x.TokenBucket
	}
	return nil
}

type isRateLimitStrategy_Strategy interface {
	isRateLimitStrategy_Strategy()
}

type RateLimitStrategy_BlanketRule_ struct {
	// Allow or Deny the requests.
	// If unset, allow all.
	BlanketRule RateLimitStrategy_BlanketRule `protobuf:"varint,1,opt,name=blanket_rule,json=blanketRule,proto3,enum=envoy.type.v3.RateLimitStrategy_BlanketRule,oneof"`
}

type RateLimitStrategy_RequestsPerTimeUnit_ struct {
	// Best-effort limit of the number of requests per time unit, f.e. requests per second.
	// Does not prescribe any specific rate limiting algorithm, see :ref:`RequestsPerTimeUnit
	// <envoy_v3_api_msg_type.v3.RateLimitStrategy.RequestsPerTimeUnit>` for details.
	RequestsPerTimeUnit *RateLimitStrategy_RequestsPerTimeUnit `protobuf:"bytes,2,opt,name=requests_per_time_unit,json=requestsPerTimeUnit,proto3,oneof"`
}

type RateLimitStrategy_TokenBucket struct {
	// Limit the requests by consuming tokens from the Token Bucket.
	// Allow the same number of requests as the number of tokens available in
	// the token bucket.
	TokenBucket *TokenBucket `protobuf:"bytes,3,opt,name=token_bucket,json=tokenBucket,proto3,oneof"`
}

func (*RateLimitStrategy_BlanketRule_) isRateLimitStrategy_Strategy() {}

func (*RateLimitStrategy_RequestsPerTimeUnit_) isRateLimitStrategy_Strategy() {}

func (*RateLimitStrategy_TokenBucket) isRateLimitStrategy_Strategy() {}

// Best-effort limit of the number of requests per time unit.
//
// Allows to specify the desired requests per second (RPS, QPS), requests per minute (QPM, RPM),
// etc., without specifying a rate limiting algorithm implementation.
//
// ``RequestsPerTimeUnit`` strategy does not demand any specific rate limiting algorithm to be
// used (in contrast to the :ref:`TokenBucket <envoy_v3_api_msg_type.v3.TokenBucket>`,
// for example). It implies that the implementation details of rate limiting algorithm are
// irrelevant as long as the configured number of "requests per time unit" is achieved.
//
// Note that the ``TokenBucket`` is still a valid implementation of the ``RequestsPerTimeUnit``
// strategy, and may be chosen to enforce the rate limit. However, there's no guarantee it will be
// the ``TokenBucket`` in particular, and not the Leaky Bucket, the Sliding Window, or any other
// rate limiting algorithm that fulfills the requirements.
type RateLimitStrategy_RequestsPerTimeUnit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The desired number of requests per :ref:`time_unit
	// <envoy_v3_api_field_type.v3.RateLimitStrategy.RequestsPerTimeUnit.time_unit>` to allow.
	// If set to ``0``, deny all (equivalent to ``BlanketRule.DENY_ALL``).
	//
	// .. note::
	//   Note that the algorithm implementation determines the course of action for the requests
	//   over the limit. As long as the ``requests_per_time_unit`` converges on the desired value,
	//   it's allowed to treat this field as a soft-limit: allow bursts, redistribute the allowance
	//   over time, etc.
	//
	RequestsPerTimeUnit uint64 `protobuf:"varint,1,opt,name=requests_per_time_unit,json=requestsPerTimeUnit,proto3" json:"requests_per_time_unit,omitempty"`
	// The unit of time. Ignored when :ref:`requests_per_time_unit
	// <envoy_v3_api_field_type.v3.RateLimitStrategy.RequestsPerTimeUnit.requests_per_time_unit>`
	// is ``0`` (deny all).
	TimeUnit RateLimitUnit `protobuf:"varint,2,opt,name=time_unit,json=timeUnit,proto3,enum=envoy.type.v3.RateLimitUnit" json:"time_unit,omitempty"`
}

func (x *RateLimitStrategy_RequestsPerTimeUnit) Reset() {
	*x = RateLimitStrategy_RequestsPerTimeUnit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_type_v3_ratelimit_strategy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateLimitStrategy_RequestsPerTimeUnit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateLimitStrategy_RequestsPerTimeUnit) ProtoMessage() {}

func (x *RateLimitStrategy_RequestsPerTimeUnit) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_type_v3_ratelimit_strategy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateLimitStrategy_RequestsPerTimeUnit.ProtoReflect.Descriptor instead.
func (*RateLimitStrategy_RequestsPerTimeUnit) Descriptor() ([]byte, []int) {
	return file_envoy_type_v3_ratelimit_strategy_proto_rawDescGZIP(), []int{0, 0}
}

func (x *RateLimitStrategy_RequestsPerTimeUnit) GetRequestsPerTimeUnit() uint64 {
	if x != nil {
		return x.RequestsPerTimeUnit
	}
	return 0
}

func (x *RateLimitStrategy_RequestsPerTimeUnit) GetTimeUnit() RateLimitUnit {
	if x != nil {
		return x.TimeUnit
	}
	return RateLimitUnit_UNKNOWN
}

var File_envoy_type_v3_ratelimit_strategy_proto protoreflect.FileDescriptor

var file_envoy_type_v3_ratelimit_strategy_proto_rawDesc = []byte{
	0x0a, 0x26, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x33, 0x2f,
	0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x33, 0x1a, 0x22, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x5f, 0x75, 0x6e, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x5f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x78,
	0x64, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76,
	0x33, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d,
	0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xed, 0x03, 0x0a, 0x11, 0x52, 0x61, 0x74, 0x65, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x5b, 0x0a, 0x0c,
	0x62, 0x6c, 0x61, 0x6e, 0x6b, 0x65, 0x74, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x76, 0x33, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x53, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x67, 0x79, 0x2e, 0x42, 0x6c, 0x61, 0x6e, 0x6b, 0x65, 0x74, 0x52, 0x75, 0x6c, 0x65,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x48, 0x00, 0x52, 0x0b, 0x62, 0x6c,
	0x61, 0x6e, 0x6b, 0x65, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x6b, 0x0a, 0x16, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75,
	0x6e, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x73, 0x50, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x6e, 0x69, 0x74, 0x48,
	0x00, 0x52, 0x13, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x50, 0x65, 0x72, 0x54, 0x69,
	0x6d, 0x65, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x3f, 0x0a, 0x0c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f,
	0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x48, 0x00, 0x52, 0x0b, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x1a, 0x8f, 0x01, 0x0a, 0x13, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x73, 0x50, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x6e, 0x69, 0x74, 0x12,
	0x33, 0x0a, 0x16, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x13, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x50, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65,
	0x55, 0x6e, 0x69, 0x74, 0x12, 0x43, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x6e, 0x69,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x55, 0x6e, 0x69, 0x74, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52,
	0x08, 0x74, 0x69, 0x6d, 0x65, 0x55, 0x6e, 0x69, 0x74, 0x22, 0x2a, 0x0a, 0x0b, 0x42, 0x6c, 0x61,
	0x6e, 0x6b, 0x65, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x4c, 0x4c, 0x4f,
	0x57, 0x5f, 0x41, 0x4c, 0x4c, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x45, 0x4e, 0x59, 0x5f,
	0x41, 0x4c, 0x4c, 0x10, 0x01, 0x42, 0x0f, 0x0a, 0x08, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67,
	0x79, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x42, 0x84, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10,
	0x02, 0xd2, 0xc6, 0xa4, 0xe1, 0x06, 0x02, 0x08, 0x01, 0x0a, 0x1b, 0x69, 0x6f, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x76, 0x33, 0x42, 0x16, 0x52, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x2f, 0x76, 0x33, 0x3b, 0x74, 0x79, 0x70, 0x65, 0x76, 0x33, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_type_v3_ratelimit_strategy_proto_rawDescOnce sync.Once
	file_envoy_type_v3_ratelimit_strategy_proto_rawDescData = file_envoy_type_v3_ratelimit_strategy_proto_rawDesc
)

func file_envoy_type_v3_ratelimit_strategy_proto_rawDescGZIP() []byte {
	file_envoy_type_v3_ratelimit_strategy_proto_rawDescOnce.Do(func() {
		file_envoy_type_v3_ratelimit_strategy_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_type_v3_ratelimit_strategy_proto_rawDescData)
	})
	return file_envoy_type_v3_ratelimit_strategy_proto_rawDescData
}

var file_envoy_type_v3_ratelimit_strategy_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_envoy_type_v3_ratelimit_strategy_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_type_v3_ratelimit_strategy_proto_goTypes = []interface{}{
	(RateLimitStrategy_BlanketRule)(0),            // 0: envoy.type.v3.RateLimitStrategy.BlanketRule
	(*RateLimitStrategy)(nil),                     // 1: envoy.type.v3.RateLimitStrategy
	(*RateLimitStrategy_RequestsPerTimeUnit)(nil), // 2: envoy.type.v3.RateLimitStrategy.RequestsPerTimeUnit
	(*TokenBucket)(nil),                           // 3: envoy.type.v3.TokenBucket
	(RateLimitUnit)(0),                            // 4: envoy.type.v3.RateLimitUnit
}
var file_envoy_type_v3_ratelimit_strategy_proto_depIdxs = []int32{
	0, // 0: envoy.type.v3.RateLimitStrategy.blanket_rule:type_name -> envoy.type.v3.RateLimitStrategy.BlanketRule
	2, // 1: envoy.type.v3.RateLimitStrategy.requests_per_time_unit:type_name -> envoy.type.v3.RateLimitStrategy.RequestsPerTimeUnit
	3, // 2: envoy.type.v3.RateLimitStrategy.token_bucket:type_name -> envoy.type.v3.TokenBucket
	4, // 3: envoy.type.v3.RateLimitStrategy.RequestsPerTimeUnit.time_unit:type_name -> envoy.type.v3.RateLimitUnit
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_envoy_type_v3_ratelimit_strategy_proto_init() }
func file_envoy_type_v3_ratelimit_strategy_proto_init() {
	if File_envoy_type_v3_ratelimit_strategy_proto != nil {
		return
	}
	file_envoy_type_v3_ratelimit_unit_proto_init()
	file_envoy_type_v3_token_bucket_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_envoy_type_v3_ratelimit_strategy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateLimitStrategy); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_type_v3_ratelimit_strategy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateLimitStrategy_RequestsPerTimeUnit); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_envoy_type_v3_ratelimit_strategy_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*RateLimitStrategy_BlanketRule_)(nil),
		(*RateLimitStrategy_RequestsPerTimeUnit_)(nil),
		(*RateLimitStrategy_TokenBucket)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_type_v3_ratelimit_strategy_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_type_v3_ratelimit_strategy_proto_goTypes,
		DependencyIndexes: file_envoy_type_v3_ratelimit_strategy_proto_depIdxs,
		EnumInfos:         file_envoy_type_v3_ratelimit_strategy_proto_enumTypes,
		MessageInfos:      file_envoy_type_v3_ratelimit_strategy_proto_msgTypes,
	}.Build()
	File_envoy_type_v3_ratelimit_strategy_proto = out.File
	file_envoy_type_v3_ratelimit_strategy_proto_rawDesc = nil
	file_envoy_type_v3_ratelimit_strategy_proto_goTypes = nil
	file_envoy_type_v3_ratelimit_strategy_proto_depIdxs = nil
}
