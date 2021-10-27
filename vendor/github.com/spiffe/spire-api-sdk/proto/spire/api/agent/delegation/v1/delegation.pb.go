// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: spire/api/agent/delegation/v1/delegation.proto

package delegationv1

import (
	proto "github.com/golang/protobuf/proto"
	types "github.com/spiffe/spire-api-sdk/proto/spire/api/types"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type FetchX509SVIDsRequest_Operation int32

const (
	FetchX509SVIDsRequest_UNSPEC FetchX509SVIDsRequest_Operation = 0
	// Add a new workload to the watched set.
	FetchX509SVIDsRequest_ADD FetchX509SVIDsRequest_Operation = 1
	// Delete a workload from the watched set.
	FetchX509SVIDsRequest_DEL FetchX509SVIDsRequest_Operation = 2
)

// Enum value maps for FetchX509SVIDsRequest_Operation.
var (
	FetchX509SVIDsRequest_Operation_name = map[int32]string{
		0: "UNSPEC",
		1: "ADD",
		2: "DEL",
	}
	FetchX509SVIDsRequest_Operation_value = map[string]int32{
		"UNSPEC": 0,
		"ADD":    1,
		"DEL":    2,
	}
)

func (x FetchX509SVIDsRequest_Operation) Enum() *FetchX509SVIDsRequest_Operation {
	p := new(FetchX509SVIDsRequest_Operation)
	*p = x
	return p
}

func (x FetchX509SVIDsRequest_Operation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FetchX509SVIDsRequest_Operation) Descriptor() protoreflect.EnumDescriptor {
	return file_spire_api_agent_delegation_v1_delegation_proto_enumTypes[0].Descriptor()
}

func (FetchX509SVIDsRequest_Operation) Type() protoreflect.EnumType {
	return &file_spire_api_agent_delegation_v1_delegation_proto_enumTypes[0]
}

func (x FetchX509SVIDsRequest_Operation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FetchX509SVIDsRequest_Operation.Descriptor instead.
func (FetchX509SVIDsRequest_Operation) EnumDescriptor() ([]byte, []int) {
	return file_spire_api_agent_delegation_v1_delegation_proto_rawDescGZIP(), []int{1, 0}
}

// X.509 SPIFFE Verifiable Identity Document with the private key.
type X509SVIDWithKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X509Svid      *types.X509SVID `protobuf:"bytes,1,opt,name=x509_svid,json=x509Svid,proto3" json:"x509_svid,omitempty"`
	X509SvidKey   []byte          `protobuf:"bytes,2,opt,name=x509_svid_key,json=x509SvidKey,proto3" json:"x509_svid_key,omitempty"`
	FederatesWith []string        `protobuf:"bytes,3,rep,name=federates_with,json=federatesWith,proto3" json:"federates_with,omitempty"`
}

func (x *X509SVIDWithKey) Reset() {
	*x = X509SVIDWithKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_api_agent_delegation_v1_delegation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *X509SVIDWithKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*X509SVIDWithKey) ProtoMessage() {}

func (x *X509SVIDWithKey) ProtoReflect() protoreflect.Message {
	mi := &file_spire_api_agent_delegation_v1_delegation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use X509SVIDWithKey.ProtoReflect.Descriptor instead.
func (*X509SVIDWithKey) Descriptor() ([]byte, []int) {
	return file_spire_api_agent_delegation_v1_delegation_proto_rawDescGZIP(), []int{0}
}

func (x *X509SVIDWithKey) GetX509Svid() *types.X509SVID {
	if x != nil {
		return x.X509Svid
	}
	return nil
}

func (x *X509SVIDWithKey) GetX509SvidKey() []byte {
	if x != nil {
		return x.X509SvidKey
	}
	return nil
}

func (x *X509SVIDWithKey) GetFederatesWith() []string {
	if x != nil {
		return x.FederatesWith
	}
	return nil
}

// FetchX509SVIDsRequest is used by clients to modify the set of workloads that
// it receives SVIDs for.
type FetchX509SVIDsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Operation to perform. Add/delete a workload to/from the watched set.
	Operation FetchX509SVIDsRequest_Operation `protobuf:"varint,1,opt,name=operation,proto3,enum=spire.api.agent.delegation.v1.FetchX509SVIDsRequest_Operation" json:"operation,omitempty"`
	// Workload's identifier. The meaning of this field depends on the
	// operation:
	// - ADD: The server will include this id in the messages containing updates
	// for this workload.
	// - DEL: The id of the workload to remove from the watched set.
	Id uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	// Selectors of the workload to add to the watched set. Only used when
	// operation is ADD.
	Selectors []*types.Selector `protobuf:"bytes,3,rep,name=selectors,proto3" json:"selectors,omitempty"`
}

func (x *FetchX509SVIDsRequest) Reset() {
	*x = FetchX509SVIDsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_api_agent_delegation_v1_delegation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchX509SVIDsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchX509SVIDsRequest) ProtoMessage() {}

func (x *FetchX509SVIDsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spire_api_agent_delegation_v1_delegation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchX509SVIDsRequest.ProtoReflect.Descriptor instead.
func (*FetchX509SVIDsRequest) Descriptor() ([]byte, []int) {
	return file_spire_api_agent_delegation_v1_delegation_proto_rawDescGZIP(), []int{1}
}

func (x *FetchX509SVIDsRequest) GetOperation() FetchX509SVIDsRequest_Operation {
	if x != nil {
		return x.Operation
	}
	return FetchX509SVIDsRequest_UNSPEC
}

func (x *FetchX509SVIDsRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FetchX509SVIDsRequest) GetSelectors() []*types.Selector {
	if x != nil {
		return x.Selectors
	}
	return nil
}

type FetchX509SVIDsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Id of the workload. This value corresponds to the one used when adding
	// the workload to the watched set.
	Id        uint64             `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	X509Svids []*X509SVIDWithKey `protobuf:"bytes,2,rep,name=x509_svids,json=x509Svids,proto3" json:"x509_svids,omitempty"`
}

func (x *FetchX509SVIDsResponse) Reset() {
	*x = FetchX509SVIDsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_api_agent_delegation_v1_delegation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchX509SVIDsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchX509SVIDsResponse) ProtoMessage() {}

func (x *FetchX509SVIDsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spire_api_agent_delegation_v1_delegation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchX509SVIDsResponse.ProtoReflect.Descriptor instead.
func (*FetchX509SVIDsResponse) Descriptor() ([]byte, []int) {
	return file_spire_api_agent_delegation_v1_delegation_proto_rawDescGZIP(), []int{2}
}

func (x *FetchX509SVIDsResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FetchX509SVIDsResponse) GetX509Svids() []*X509SVIDWithKey {
	if x != nil {
		return x.X509Svids
	}
	return nil
}

type FetchX509BundlesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FetchX509BundlesRequest) Reset() {
	*x = FetchX509BundlesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_api_agent_delegation_v1_delegation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchX509BundlesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchX509BundlesRequest) ProtoMessage() {}

func (x *FetchX509BundlesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spire_api_agent_delegation_v1_delegation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchX509BundlesRequest.ProtoReflect.Descriptor instead.
func (*FetchX509BundlesRequest) Descriptor() ([]byte, []int) {
	return file_spire_api_agent_delegation_v1_delegation_proto_rawDescGZIP(), []int{3}
}

// FetchX509BundlesResponse represents the bundle for a given trust domain. An
// empty bundle indicates that it was removed.
type FetchX509BundlesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrustDomainName string `protobuf:"bytes,1,opt,name=trust_domain_name,json=trustDomainName,proto3" json:"trust_domain_name,omitempty"`
	Bundle          []byte `protobuf:"bytes,2,opt,name=bundle,proto3" json:"bundle,omitempty"`
}

func (x *FetchX509BundlesResponse) Reset() {
	*x = FetchX509BundlesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_api_agent_delegation_v1_delegation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchX509BundlesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchX509BundlesResponse) ProtoMessage() {}

func (x *FetchX509BundlesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spire_api_agent_delegation_v1_delegation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchX509BundlesResponse.ProtoReflect.Descriptor instead.
func (*FetchX509BundlesResponse) Descriptor() ([]byte, []int) {
	return file_spire_api_agent_delegation_v1_delegation_proto_rawDescGZIP(), []int{4}
}

func (x *FetchX509BundlesResponse) GetTrustDomainName() string {
	if x != nil {
		return x.TrustDomainName
	}
	return ""
}

func (x *FetchX509BundlesResponse) GetBundle() []byte {
	if x != nil {
		return x.Bundle
	}
	return nil
}

var File_spire_api_agent_delegation_v1_delegation_proto protoreflect.FileDescriptor

var file_spire_api_agent_delegation_v1_delegation_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f,
	0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1d, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x2e, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a,
	0x1e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2f, 0x78, 0x35, 0x30, 0x39, 0x73, 0x76, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x94, 0x01, 0x0a, 0x0f, 0x58, 0x35, 0x30, 0x39, 0x53, 0x56, 0x49, 0x44, 0x57, 0x69, 0x74, 0x68,
	0x4b, 0x65, 0x79, 0x12, 0x36, 0x0a, 0x09, 0x78, 0x35, 0x30, 0x39, 0x5f, 0x73, 0x76, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x58, 0x35, 0x30, 0x39, 0x53, 0x56, 0x49,
	0x44, 0x52, 0x08, 0x78, 0x35, 0x30, 0x39, 0x53, 0x76, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x78,
	0x35, 0x30, 0x39, 0x5f, 0x73, 0x76, 0x69, 0x64, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0b, 0x78, 0x35, 0x30, 0x39, 0x53, 0x76, 0x69, 0x64, 0x4b, 0x65, 0x79, 0x12,
	0x25, 0x0a, 0x0e, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x73, 0x5f, 0x77, 0x69, 0x74,
	0x68, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x73, 0x57, 0x69, 0x74, 0x68, 0x22, 0xe9, 0x01, 0x0a, 0x15, 0x46, 0x65, 0x74, 0x63, 0x68,
	0x58, 0x35, 0x30, 0x39, 0x53, 0x56, 0x49, 0x44, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x5c, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x3e, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x58, 0x35, 0x30, 0x39, 0x53, 0x56, 0x49,
	0x44, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x37,
	0x0a, 0x09, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x09, 0x73, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x22, 0x29, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x10, 0x00,
	0x12, 0x07, 0x0a, 0x03, 0x41, 0x44, 0x44, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x44, 0x45, 0x4c,
	0x10, 0x02, 0x22, 0x77, 0x0a, 0x16, 0x46, 0x65, 0x74, 0x63, 0x68, 0x58, 0x35, 0x30, 0x39, 0x53,
	0x56, 0x49, 0x44, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x4d, 0x0a, 0x0a,
	0x78, 0x35, 0x30, 0x39, 0x5f, 0x73, 0x76, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2e, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x2e, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x58, 0x35, 0x30, 0x39, 0x53, 0x56, 0x49, 0x44, 0x57, 0x69, 0x74, 0x68, 0x4b, 0x65, 0x79,
	0x52, 0x09, 0x78, 0x35, 0x30, 0x39, 0x53, 0x76, 0x69, 0x64, 0x73, 0x22, 0x19, 0x0a, 0x17, 0x46,
	0x65, 0x74, 0x63, 0x68, 0x58, 0x35, 0x30, 0x39, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x5e, 0x0a, 0x18, 0x46, 0x65, 0x74, 0x63, 0x68, 0x58,
	0x35, 0x30, 0x39, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x74, 0x72, 0x75, 0x73, 0x74, 0x5f, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74,
	0x72, 0x75, 0x73, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06,
	0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x32, 0x98, 0x02, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x67,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x81, 0x01, 0x0a, 0x0e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x58,
	0x35, 0x30, 0x39, 0x53, 0x56, 0x49, 0x44, 0x73, 0x12, 0x34, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x64, 0x65, 0x6c, 0x65, 0x67,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x58, 0x35,
	0x30, 0x39, 0x53, 0x56, 0x49, 0x44, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35,
	0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x2e, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x46,
	0x65, 0x74, 0x63, 0x68, 0x58, 0x35, 0x30, 0x39, 0x53, 0x56, 0x49, 0x44, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x12, 0x85, 0x01, 0x0a, 0x10, 0x46, 0x65,
	0x74, 0x63, 0x68, 0x58, 0x35, 0x30, 0x39, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x12, 0x36,
	0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x2e, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x46,
	0x65, 0x74, 0x63, 0x68, 0x58, 0x35, 0x30, 0x39, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x58, 0x35, 0x30, 0x39,
	0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30,
	0x01, 0x42, 0x52, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x70, 0x69, 0x66, 0x66, 0x65, 0x2f, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2d, 0x61, 0x70, 0x69,
	0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x70, 0x69, 0x72, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x67,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spire_api_agent_delegation_v1_delegation_proto_rawDescOnce sync.Once
	file_spire_api_agent_delegation_v1_delegation_proto_rawDescData = file_spire_api_agent_delegation_v1_delegation_proto_rawDesc
)

func file_spire_api_agent_delegation_v1_delegation_proto_rawDescGZIP() []byte {
	file_spire_api_agent_delegation_v1_delegation_proto_rawDescOnce.Do(func() {
		file_spire_api_agent_delegation_v1_delegation_proto_rawDescData = protoimpl.X.CompressGZIP(file_spire_api_agent_delegation_v1_delegation_proto_rawDescData)
	})
	return file_spire_api_agent_delegation_v1_delegation_proto_rawDescData
}

var file_spire_api_agent_delegation_v1_delegation_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_spire_api_agent_delegation_v1_delegation_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_spire_api_agent_delegation_v1_delegation_proto_goTypes = []interface{}{
	(FetchX509SVIDsRequest_Operation)(0), // 0: spire.api.agent.delegation.v1.FetchX509SVIDsRequest.Operation
	(*X509SVIDWithKey)(nil),              // 1: spire.api.agent.delegation.v1.X509SVIDWithKey
	(*FetchX509SVIDsRequest)(nil),        // 2: spire.api.agent.delegation.v1.FetchX509SVIDsRequest
	(*FetchX509SVIDsResponse)(nil),       // 3: spire.api.agent.delegation.v1.FetchX509SVIDsResponse
	(*FetchX509BundlesRequest)(nil),      // 4: spire.api.agent.delegation.v1.FetchX509BundlesRequest
	(*FetchX509BundlesResponse)(nil),     // 5: spire.api.agent.delegation.v1.FetchX509BundlesResponse
	(*types.X509SVID)(nil),               // 6: spire.api.types.X509SVID
	(*types.Selector)(nil),               // 7: spire.api.types.Selector
}
var file_spire_api_agent_delegation_v1_delegation_proto_depIdxs = []int32{
	6, // 0: spire.api.agent.delegation.v1.X509SVIDWithKey.x509_svid:type_name -> spire.api.types.X509SVID
	0, // 1: spire.api.agent.delegation.v1.FetchX509SVIDsRequest.operation:type_name -> spire.api.agent.delegation.v1.FetchX509SVIDsRequest.Operation
	7, // 2: spire.api.agent.delegation.v1.FetchX509SVIDsRequest.selectors:type_name -> spire.api.types.Selector
	1, // 3: spire.api.agent.delegation.v1.FetchX509SVIDsResponse.x509_svids:type_name -> spire.api.agent.delegation.v1.X509SVIDWithKey
	2, // 4: spire.api.agent.delegation.v1.Delegation.FetchX509SVIDs:input_type -> spire.api.agent.delegation.v1.FetchX509SVIDsRequest
	4, // 5: spire.api.agent.delegation.v1.Delegation.FetchX509Bundles:input_type -> spire.api.agent.delegation.v1.FetchX509BundlesRequest
	3, // 6: spire.api.agent.delegation.v1.Delegation.FetchX509SVIDs:output_type -> spire.api.agent.delegation.v1.FetchX509SVIDsResponse
	5, // 7: spire.api.agent.delegation.v1.Delegation.FetchX509Bundles:output_type -> spire.api.agent.delegation.v1.FetchX509BundlesResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_spire_api_agent_delegation_v1_delegation_proto_init() }
func file_spire_api_agent_delegation_v1_delegation_proto_init() {
	if File_spire_api_agent_delegation_v1_delegation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spire_api_agent_delegation_v1_delegation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*X509SVIDWithKey); i {
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
		file_spire_api_agent_delegation_v1_delegation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchX509SVIDsRequest); i {
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
		file_spire_api_agent_delegation_v1_delegation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchX509SVIDsResponse); i {
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
		file_spire_api_agent_delegation_v1_delegation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchX509BundlesRequest); i {
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
		file_spire_api_agent_delegation_v1_delegation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchX509BundlesResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spire_api_agent_delegation_v1_delegation_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spire_api_agent_delegation_v1_delegation_proto_goTypes,
		DependencyIndexes: file_spire_api_agent_delegation_v1_delegation_proto_depIdxs,
		EnumInfos:         file_spire_api_agent_delegation_v1_delegation_proto_enumTypes,
		MessageInfos:      file_spire_api_agent_delegation_v1_delegation_proto_msgTypes,
	}.Build()
	File_spire_api_agent_delegation_v1_delegation_proto = out.File
	file_spire_api_agent_delegation_v1_delegation_proto_rawDesc = nil
	file_spire_api_agent_delegation_v1_delegation_proto_goTypes = nil
	file_spire_api_agent_delegation_v1_delegation_proto_depIdxs = nil
}