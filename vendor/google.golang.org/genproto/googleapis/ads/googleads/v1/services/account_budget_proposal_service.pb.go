// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/account_budget_proposal_service.proto

package services // import "google.golang.org/genproto/googleapis/ads/googleads/v1/services"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import field_mask "google.golang.org/genproto/protobuf/field_mask"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Request message for
// [AccountBudgetProposalService.GetAccountBudgetProposal][google.ads.googleads.v1.services.AccountBudgetProposalService.GetAccountBudgetProposal].
type GetAccountBudgetProposalRequest struct {
	// The resource name of the account-level budget proposal to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAccountBudgetProposalRequest) Reset()         { *m = GetAccountBudgetProposalRequest{} }
func (m *GetAccountBudgetProposalRequest) String() string { return proto.CompactTextString(m) }
func (*GetAccountBudgetProposalRequest) ProtoMessage()    {}
func (*GetAccountBudgetProposalRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_budget_proposal_service_19af3128d079e096, []int{0}
}
func (m *GetAccountBudgetProposalRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAccountBudgetProposalRequest.Unmarshal(m, b)
}
func (m *GetAccountBudgetProposalRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAccountBudgetProposalRequest.Marshal(b, m, deterministic)
}
func (dst *GetAccountBudgetProposalRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAccountBudgetProposalRequest.Merge(dst, src)
}
func (m *GetAccountBudgetProposalRequest) XXX_Size() int {
	return xxx_messageInfo_GetAccountBudgetProposalRequest.Size(m)
}
func (m *GetAccountBudgetProposalRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAccountBudgetProposalRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAccountBudgetProposalRequest proto.InternalMessageInfo

func (m *GetAccountBudgetProposalRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for
// [AccountBudgetProposalService.MutateAccountBudgetProposal][google.ads.googleads.v1.services.AccountBudgetProposalService.MutateAccountBudgetProposal].
type MutateAccountBudgetProposalRequest struct {
	// The ID of the customer.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The operation to perform on an individual account-level budget proposal.
	Operation *AccountBudgetProposalOperation `protobuf:"bytes,2,opt,name=operation,proto3" json:"operation,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly         bool     `protobuf:"varint,3,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateAccountBudgetProposalRequest) Reset()         { *m = MutateAccountBudgetProposalRequest{} }
func (m *MutateAccountBudgetProposalRequest) String() string { return proto.CompactTextString(m) }
func (*MutateAccountBudgetProposalRequest) ProtoMessage()    {}
func (*MutateAccountBudgetProposalRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_budget_proposal_service_19af3128d079e096, []int{1}
}
func (m *MutateAccountBudgetProposalRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAccountBudgetProposalRequest.Unmarshal(m, b)
}
func (m *MutateAccountBudgetProposalRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAccountBudgetProposalRequest.Marshal(b, m, deterministic)
}
func (dst *MutateAccountBudgetProposalRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAccountBudgetProposalRequest.Merge(dst, src)
}
func (m *MutateAccountBudgetProposalRequest) XXX_Size() int {
	return xxx_messageInfo_MutateAccountBudgetProposalRequest.Size(m)
}
func (m *MutateAccountBudgetProposalRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAccountBudgetProposalRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAccountBudgetProposalRequest proto.InternalMessageInfo

func (m *MutateAccountBudgetProposalRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateAccountBudgetProposalRequest) GetOperation() *AccountBudgetProposalOperation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *MutateAccountBudgetProposalRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation to propose the creation of a new account-level budget or
// edit/end/remove an existing one.
type AccountBudgetProposalOperation struct {
	// FieldMask that determines which budget fields are modified.  While budgets
	// may be modified, proposals that propose such modifications are final.
	// Therefore, update operations are not supported for proposals.
	//
	// Proposals that modify budgets have the 'update' proposal type.  Specifying
	// a mask for any other proposal type is considered an error.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*AccountBudgetProposalOperation_Create
	//	*AccountBudgetProposalOperation_Remove
	Operation            isAccountBudgetProposalOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                                   `json:"-"`
	XXX_unrecognized     []byte                                     `json:"-"`
	XXX_sizecache        int32                                      `json:"-"`
}

func (m *AccountBudgetProposalOperation) Reset()         { *m = AccountBudgetProposalOperation{} }
func (m *AccountBudgetProposalOperation) String() string { return proto.CompactTextString(m) }
func (*AccountBudgetProposalOperation) ProtoMessage()    {}
func (*AccountBudgetProposalOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_budget_proposal_service_19af3128d079e096, []int{2}
}
func (m *AccountBudgetProposalOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountBudgetProposalOperation.Unmarshal(m, b)
}
func (m *AccountBudgetProposalOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountBudgetProposalOperation.Marshal(b, m, deterministic)
}
func (dst *AccountBudgetProposalOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountBudgetProposalOperation.Merge(dst, src)
}
func (m *AccountBudgetProposalOperation) XXX_Size() int {
	return xxx_messageInfo_AccountBudgetProposalOperation.Size(m)
}
func (m *AccountBudgetProposalOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountBudgetProposalOperation.DiscardUnknown(m)
}

var xxx_messageInfo_AccountBudgetProposalOperation proto.InternalMessageInfo

func (m *AccountBudgetProposalOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isAccountBudgetProposalOperation_Operation interface {
	isAccountBudgetProposalOperation_Operation()
}

type AccountBudgetProposalOperation_Create struct {
	Create *resources.AccountBudgetProposal `protobuf:"bytes,2,opt,name=create,proto3,oneof"`
}

type AccountBudgetProposalOperation_Remove struct {
	Remove string `protobuf:"bytes,1,opt,name=remove,proto3,oneof"`
}

func (*AccountBudgetProposalOperation_Create) isAccountBudgetProposalOperation_Operation() {}

func (*AccountBudgetProposalOperation_Remove) isAccountBudgetProposalOperation_Operation() {}

func (m *AccountBudgetProposalOperation) GetOperation() isAccountBudgetProposalOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *AccountBudgetProposalOperation) GetCreate() *resources.AccountBudgetProposal {
	if x, ok := m.GetOperation().(*AccountBudgetProposalOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *AccountBudgetProposalOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*AccountBudgetProposalOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*AccountBudgetProposalOperation) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _AccountBudgetProposalOperation_OneofMarshaler, _AccountBudgetProposalOperation_OneofUnmarshaler, _AccountBudgetProposalOperation_OneofSizer, []interface{}{
		(*AccountBudgetProposalOperation_Create)(nil),
		(*AccountBudgetProposalOperation_Remove)(nil),
	}
}

func _AccountBudgetProposalOperation_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*AccountBudgetProposalOperation)
	// operation
	switch x := m.Operation.(type) {
	case *AccountBudgetProposalOperation_Create:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Create); err != nil {
			return err
		}
	case *AccountBudgetProposalOperation_Remove:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Remove)
	case nil:
	default:
		return fmt.Errorf("AccountBudgetProposalOperation.Operation has unexpected type %T", x)
	}
	return nil
}

func _AccountBudgetProposalOperation_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*AccountBudgetProposalOperation)
	switch tag {
	case 2: // operation.create
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(resources.AccountBudgetProposal)
		err := b.DecodeMessage(msg)
		m.Operation = &AccountBudgetProposalOperation_Create{msg}
		return true, err
	case 1: // operation.remove
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Operation = &AccountBudgetProposalOperation_Remove{x}
		return true, err
	default:
		return false, nil
	}
}

func _AccountBudgetProposalOperation_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*AccountBudgetProposalOperation)
	// operation
	switch x := m.Operation.(type) {
	case *AccountBudgetProposalOperation_Create:
		s := proto.Size(x.Create)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AccountBudgetProposalOperation_Remove:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Remove)))
		n += len(x.Remove)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Response message for account-level budget mutate operations.
type MutateAccountBudgetProposalResponse struct {
	// The result of the mutate.
	Result               *MutateAccountBudgetProposalResult `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *MutateAccountBudgetProposalResponse) Reset()         { *m = MutateAccountBudgetProposalResponse{} }
func (m *MutateAccountBudgetProposalResponse) String() string { return proto.CompactTextString(m) }
func (*MutateAccountBudgetProposalResponse) ProtoMessage()    {}
func (*MutateAccountBudgetProposalResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_budget_proposal_service_19af3128d079e096, []int{3}
}
func (m *MutateAccountBudgetProposalResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAccountBudgetProposalResponse.Unmarshal(m, b)
}
func (m *MutateAccountBudgetProposalResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAccountBudgetProposalResponse.Marshal(b, m, deterministic)
}
func (dst *MutateAccountBudgetProposalResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAccountBudgetProposalResponse.Merge(dst, src)
}
func (m *MutateAccountBudgetProposalResponse) XXX_Size() int {
	return xxx_messageInfo_MutateAccountBudgetProposalResponse.Size(m)
}
func (m *MutateAccountBudgetProposalResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAccountBudgetProposalResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAccountBudgetProposalResponse proto.InternalMessageInfo

func (m *MutateAccountBudgetProposalResponse) GetResult() *MutateAccountBudgetProposalResult {
	if m != nil {
		return m.Result
	}
	return nil
}

// The result for the account budget proposal mutate.
type MutateAccountBudgetProposalResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateAccountBudgetProposalResult) Reset()         { *m = MutateAccountBudgetProposalResult{} }
func (m *MutateAccountBudgetProposalResult) String() string { return proto.CompactTextString(m) }
func (*MutateAccountBudgetProposalResult) ProtoMessage()    {}
func (*MutateAccountBudgetProposalResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_budget_proposal_service_19af3128d079e096, []int{4}
}
func (m *MutateAccountBudgetProposalResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAccountBudgetProposalResult.Unmarshal(m, b)
}
func (m *MutateAccountBudgetProposalResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAccountBudgetProposalResult.Marshal(b, m, deterministic)
}
func (dst *MutateAccountBudgetProposalResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAccountBudgetProposalResult.Merge(dst, src)
}
func (m *MutateAccountBudgetProposalResult) XXX_Size() int {
	return xxx_messageInfo_MutateAccountBudgetProposalResult.Size(m)
}
func (m *MutateAccountBudgetProposalResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAccountBudgetProposalResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAccountBudgetProposalResult proto.InternalMessageInfo

func (m *MutateAccountBudgetProposalResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetAccountBudgetProposalRequest)(nil), "google.ads.googleads.v1.services.GetAccountBudgetProposalRequest")
	proto.RegisterType((*MutateAccountBudgetProposalRequest)(nil), "google.ads.googleads.v1.services.MutateAccountBudgetProposalRequest")
	proto.RegisterType((*AccountBudgetProposalOperation)(nil), "google.ads.googleads.v1.services.AccountBudgetProposalOperation")
	proto.RegisterType((*MutateAccountBudgetProposalResponse)(nil), "google.ads.googleads.v1.services.MutateAccountBudgetProposalResponse")
	proto.RegisterType((*MutateAccountBudgetProposalResult)(nil), "google.ads.googleads.v1.services.MutateAccountBudgetProposalResult")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccountBudgetProposalServiceClient is the client API for AccountBudgetProposalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountBudgetProposalServiceClient interface {
	// Returns an account-level budget proposal in full detail.
	GetAccountBudgetProposal(ctx context.Context, in *GetAccountBudgetProposalRequest, opts ...grpc.CallOption) (*resources.AccountBudgetProposal, error)
	// Creates, updates, or removes account budget proposals.  Operation statuses
	// are returned.
	MutateAccountBudgetProposal(ctx context.Context, in *MutateAccountBudgetProposalRequest, opts ...grpc.CallOption) (*MutateAccountBudgetProposalResponse, error)
}

type accountBudgetProposalServiceClient struct {
	cc *grpc.ClientConn
}

func NewAccountBudgetProposalServiceClient(cc *grpc.ClientConn) AccountBudgetProposalServiceClient {
	return &accountBudgetProposalServiceClient{cc}
}

func (c *accountBudgetProposalServiceClient) GetAccountBudgetProposal(ctx context.Context, in *GetAccountBudgetProposalRequest, opts ...grpc.CallOption) (*resources.AccountBudgetProposal, error) {
	out := new(resources.AccountBudgetProposal)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.AccountBudgetProposalService/GetAccountBudgetProposal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountBudgetProposalServiceClient) MutateAccountBudgetProposal(ctx context.Context, in *MutateAccountBudgetProposalRequest, opts ...grpc.CallOption) (*MutateAccountBudgetProposalResponse, error) {
	out := new(MutateAccountBudgetProposalResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.AccountBudgetProposalService/MutateAccountBudgetProposal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountBudgetProposalServiceServer is the server API for AccountBudgetProposalService service.
type AccountBudgetProposalServiceServer interface {
	// Returns an account-level budget proposal in full detail.
	GetAccountBudgetProposal(context.Context, *GetAccountBudgetProposalRequest) (*resources.AccountBudgetProposal, error)
	// Creates, updates, or removes account budget proposals.  Operation statuses
	// are returned.
	MutateAccountBudgetProposal(context.Context, *MutateAccountBudgetProposalRequest) (*MutateAccountBudgetProposalResponse, error)
}

func RegisterAccountBudgetProposalServiceServer(s *grpc.Server, srv AccountBudgetProposalServiceServer) {
	s.RegisterService(&_AccountBudgetProposalService_serviceDesc, srv)
}

func _AccountBudgetProposalService_GetAccountBudgetProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountBudgetProposalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountBudgetProposalServiceServer).GetAccountBudgetProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.AccountBudgetProposalService/GetAccountBudgetProposal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountBudgetProposalServiceServer).GetAccountBudgetProposal(ctx, req.(*GetAccountBudgetProposalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountBudgetProposalService_MutateAccountBudgetProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAccountBudgetProposalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountBudgetProposalServiceServer).MutateAccountBudgetProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.AccountBudgetProposalService/MutateAccountBudgetProposal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountBudgetProposalServiceServer).MutateAccountBudgetProposal(ctx, req.(*MutateAccountBudgetProposalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccountBudgetProposalService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.AccountBudgetProposalService",
	HandlerType: (*AccountBudgetProposalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAccountBudgetProposal",
			Handler:    _AccountBudgetProposalService_GetAccountBudgetProposal_Handler,
		},
		{
			MethodName: "MutateAccountBudgetProposal",
			Handler:    _AccountBudgetProposalService_MutateAccountBudgetProposal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/account_budget_proposal_service.proto",
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/account_budget_proposal_service.proto", fileDescriptor_account_budget_proposal_service_19af3128d079e096)
}

var fileDescriptor_account_budget_proposal_service_19af3128d079e096 = []byte{
	// 629 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0x3f, 0x6f, 0xd3, 0x4e,
	0x18, 0xfe, 0xd9, 0x95, 0xa2, 0x5f, 0x2f, 0x65, 0xb9, 0x29, 0x0a, 0x15, 0x4d, 0xdd, 0x0e, 0x51,
	0x06, 0x5b, 0x09, 0x4b, 0xe5, 0xa8, 0x22, 0x0e, 0x90, 0x84, 0xa1, 0x34, 0x32, 0x52, 0x06, 0x88,
	0xb0, 0x2e, 0xf6, 0xd5, 0xb2, 0x6a, 0xfb, 0x8c, 0xef, 0x1c, 0xa9, 0xaa, 0xba, 0x74, 0x62, 0xe7,
	0x1b, 0x30, 0xf2, 0x51, 0xba, 0xb2, 0xc0, 0xce, 0xc4, 0x82, 0xc4, 0x27, 0x40, 0xf6, 0xdd, 0xa5,
	0xad, 0x94, 0xd8, 0x88, 0x6e, 0xaf, 0xcf, 0x8f, 0x9f, 0xe7, 0x7d, 0xde, 0x3f, 0x67, 0x30, 0xf2,
	0x09, 0xf1, 0x43, 0x6c, 0x20, 0x8f, 0x1a, 0x3c, 0xcc, 0xa3, 0x65, 0xd7, 0xa0, 0x38, 0x5d, 0x06,
	0x2e, 0xa6, 0x06, 0x72, 0x5d, 0x92, 0xc5, 0xcc, 0x59, 0x64, 0x9e, 0x8f, 0x99, 0x93, 0xa4, 0x24,
	0x21, 0x14, 0x85, 0x8e, 0x00, 0xe8, 0x49, 0x4a, 0x18, 0x81, 0x2d, 0xfe, 0xb1, 0x8e, 0x3c, 0xaa,
	0xaf, 0x78, 0xf4, 0x65, 0x57, 0x97, 0x3c, 0xcd, 0x67, 0x9b, 0x94, 0x52, 0x4c, 0x49, 0x96, 0x96,
	0x48, 0x71, 0x89, 0xe6, 0xae, 0x24, 0x48, 0x02, 0x03, 0xc5, 0x31, 0x61, 0x88, 0x05, 0x24, 0xa6,
	0xe2, 0xad, 0x48, 0xc0, 0x28, 0x9e, 0x16, 0xd9, 0x99, 0x71, 0x16, 0xe0, 0xd0, 0x73, 0x22, 0x44,
	0xcf, 0x39, 0x42, 0x1b, 0x81, 0xbd, 0x31, 0x66, 0x16, 0xd7, 0x18, 0x16, 0x12, 0x53, 0xa1, 0x60,
	0xe3, 0x0f, 0x19, 0xa6, 0x0c, 0x1e, 0x80, 0x47, 0x32, 0x1b, 0x27, 0x46, 0x11, 0x6e, 0x28, 0x2d,
	0xa5, 0xbd, 0x6d, 0xef, 0xc8, 0xc3, 0xd7, 0x28, 0xc2, 0xda, 0x8d, 0x02, 0xb4, 0x93, 0x8c, 0x21,
	0x86, 0x4b, 0xb9, 0xf6, 0x40, 0xdd, 0xcd, 0x28, 0x23, 0x11, 0x4e, 0x9d, 0xc0, 0x13, 0x4c, 0x40,
	0x1e, 0xbd, 0xf2, 0xe0, 0x7b, 0xb0, 0x4d, 0x12, 0x9c, 0x16, 0x2e, 0x1a, 0x6a, 0x4b, 0x69, 0xd7,
	0x7b, 0x03, 0xbd, 0xaa, 0x8c, 0xfa, 0x5a, 0xcd, 0x53, 0xc9, 0x63, 0xdf, 0x52, 0xe6, 0x66, 0x96,
	0x28, 0x0c, 0x3c, 0xc4, 0xb0, 0x43, 0xe2, 0xf0, 0xa2, 0xb1, 0xd5, 0x52, 0xda, 0xff, 0xdb, 0x3b,
	0xf2, 0xf0, 0x34, 0x0e, 0x2f, 0xb4, 0x6f, 0x0a, 0x78, 0x52, 0x4e, 0x09, 0xfb, 0xa0, 0x9e, 0x25,
	0x05, 0x4b, 0x5e, 0xcc, 0x82, 0xa5, 0xde, 0x6b, 0xca, 0x4c, 0x65, 0xbd, 0xf5, 0x51, 0x5e, 0xef,
	0x13, 0x44, 0xcf, 0x6d, 0xc0, 0xe1, 0x79, 0x0c, 0x6d, 0x50, 0x73, 0x53, 0x8c, 0x18, 0x16, 0x0e,
	0x8f, 0x36, 0x3a, 0x5c, 0x8d, 0xc1, 0x7a, 0x8b, 0x93, 0xff, 0x6c, 0xc1, 0x04, 0x1b, 0xa0, 0x96,
	0xe2, 0x88, 0x2c, 0x45, 0x7b, 0xf2, 0x37, 0xfc, 0x79, 0x58, 0xbf, 0x53, 0x52, 0xed, 0x5a, 0x01,
	0x07, 0xa5, 0x7d, 0xa2, 0x09, 0x89, 0x29, 0x86, 0xef, 0x72, 0x3a, 0x9a, 0x85, 0x4c, 0xa4, 0xf8,
	0xbc, 0xba, 0x09, 0xe5, 0xb4, 0x59, 0xc8, 0x6c, 0x41, 0xa9, 0x4d, 0xc0, 0x7e, 0x25, 0xf8, 0xaf,
	0xc6, 0xae, 0xf7, 0x6b, 0x0b, 0xec, 0xae, 0x25, 0x79, 0xc3, 0xb3, 0x82, 0xdf, 0x15, 0xd0, 0xd8,
	0x34, 0xe0, 0xd0, 0xaa, 0x36, 0x55, 0xb1, 0x1c, 0xcd, 0x7f, 0x6e, 0x9d, 0x36, 0xb8, 0xfe, 0xfa,
	0xe3, 0x93, 0x6a, 0xc2, 0xa3, 0x7c, 0xdd, 0x2f, 0xef, 0x59, 0x3d, 0x96, 0x0b, 0x41, 0x8d, 0x8e,
	0xdc, 0xff, 0xfb, 0x5f, 0x53, 0xa3, 0x73, 0x05, 0x7f, 0x2b, 0xe0, 0x71, 0x49, 0x1d, 0xe1, 0x8b,
	0x07, 0xf6, 0x8c, 0x3b, 0x7c, 0xf9, 0xd0, 0xce, 0x17, 0x03, 0xa5, 0x8d, 0x0a, 0xbb, 0x03, 0xad,
	0x9f, 0xdb, 0xbd, 0xf5, 0x77, 0x79, 0xe7, 0x3a, 0x38, 0xee, 0x5c, 0x6d, 0x70, 0x6b, 0x46, 0x85,
	0x82, 0xa9, 0x74, 0x86, 0x1f, 0x55, 0x70, 0xe8, 0x92, 0xa8, 0x32, 0xa9, 0xe1, 0x7e, 0xd9, 0x5c,
	0x4c, 0xf3, 0x05, 0x9d, 0x2a, 0x6f, 0x27, 0x82, 0xc6, 0x27, 0x21, 0x8a, 0x7d, 0x9d, 0xa4, 0xbe,
	0xe1, 0xe3, 0xb8, 0x58, 0x5f, 0x79, 0x1f, 0x27, 0x01, 0xdd, 0xfc, 0x23, 0xe8, 0xcb, 0xe0, 0xb3,
	0xba, 0x35, 0xb6, 0xac, 0x2f, 0x6a, 0x6b, 0xcc, 0x09, 0x2d, 0x8f, 0xea, 0x3c, 0xcc, 0xa3, 0x59,
	0x57, 0x17, 0xc2, 0xf4, 0x46, 0x42, 0xe6, 0x96, 0x47, 0xe7, 0x2b, 0xc8, 0x7c, 0xd6, 0x9d, 0x4b,
	0xc8, 0x4f, 0xf5, 0x90, 0x9f, 0x9b, 0xa6, 0xe5, 0x51, 0xd3, 0x5c, 0x81, 0x4c, 0x73, 0xd6, 0x35,
	0x4d, 0x09, 0x5b, 0xd4, 0x8a, 0x3c, 0x9f, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0x83, 0x78, 0xdf,
	0x87, 0xaf, 0x06, 0x00, 0x00,
}
