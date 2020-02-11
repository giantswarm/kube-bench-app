// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/billing_setup_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Request message for
// [BillingSetupService.GetBillingSetup][google.ads.googleads.v2.services.BillingSetupService.GetBillingSetup].
type GetBillingSetupRequest struct {
	// The resource name of the billing setup to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBillingSetupRequest) Reset()         { *m = GetBillingSetupRequest{} }
func (m *GetBillingSetupRequest) String() string { return proto.CompactTextString(m) }
func (*GetBillingSetupRequest) ProtoMessage()    {}
func (*GetBillingSetupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4870990e51228b9d, []int{0}
}

func (m *GetBillingSetupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBillingSetupRequest.Unmarshal(m, b)
}
func (m *GetBillingSetupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBillingSetupRequest.Marshal(b, m, deterministic)
}
func (m *GetBillingSetupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBillingSetupRequest.Merge(m, src)
}
func (m *GetBillingSetupRequest) XXX_Size() int {
	return xxx_messageInfo_GetBillingSetupRequest.Size(m)
}
func (m *GetBillingSetupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBillingSetupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBillingSetupRequest proto.InternalMessageInfo

func (m *GetBillingSetupRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for billing setup mutate operations.
type MutateBillingSetupRequest struct {
	// Id of the customer to apply the billing setup mutate operation to.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The operation to perform.
	Operation            *BillingSetupOperation `protobuf:"bytes,2,opt,name=operation,proto3" json:"operation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *MutateBillingSetupRequest) Reset()         { *m = MutateBillingSetupRequest{} }
func (m *MutateBillingSetupRequest) String() string { return proto.CompactTextString(m) }
func (*MutateBillingSetupRequest) ProtoMessage()    {}
func (*MutateBillingSetupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4870990e51228b9d, []int{1}
}

func (m *MutateBillingSetupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateBillingSetupRequest.Unmarshal(m, b)
}
func (m *MutateBillingSetupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateBillingSetupRequest.Marshal(b, m, deterministic)
}
func (m *MutateBillingSetupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateBillingSetupRequest.Merge(m, src)
}
func (m *MutateBillingSetupRequest) XXX_Size() int {
	return xxx_messageInfo_MutateBillingSetupRequest.Size(m)
}
func (m *MutateBillingSetupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateBillingSetupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateBillingSetupRequest proto.InternalMessageInfo

func (m *MutateBillingSetupRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateBillingSetupRequest) GetOperation() *BillingSetupOperation {
	if m != nil {
		return m.Operation
	}
	return nil
}

// A single operation on a billing setup, which describes the cancellation of an
// existing billing setup.
type BillingSetupOperation struct {
	// Only one of these operations can be set. "Update" operations are not
	// supported.
	//
	// Types that are valid to be assigned to Operation:
	//	*BillingSetupOperation_Create
	//	*BillingSetupOperation_Remove
	Operation            isBillingSetupOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *BillingSetupOperation) Reset()         { *m = BillingSetupOperation{} }
func (m *BillingSetupOperation) String() string { return proto.CompactTextString(m) }
func (*BillingSetupOperation) ProtoMessage()    {}
func (*BillingSetupOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_4870990e51228b9d, []int{2}
}

func (m *BillingSetupOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BillingSetupOperation.Unmarshal(m, b)
}
func (m *BillingSetupOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BillingSetupOperation.Marshal(b, m, deterministic)
}
func (m *BillingSetupOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BillingSetupOperation.Merge(m, src)
}
func (m *BillingSetupOperation) XXX_Size() int {
	return xxx_messageInfo_BillingSetupOperation.Size(m)
}
func (m *BillingSetupOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_BillingSetupOperation.DiscardUnknown(m)
}

var xxx_messageInfo_BillingSetupOperation proto.InternalMessageInfo

type isBillingSetupOperation_Operation interface {
	isBillingSetupOperation_Operation()
}

type BillingSetupOperation_Create struct {
	Create *resources.BillingSetup `protobuf:"bytes,2,opt,name=create,proto3,oneof"`
}

type BillingSetupOperation_Remove struct {
	Remove string `protobuf:"bytes,1,opt,name=remove,proto3,oneof"`
}

func (*BillingSetupOperation_Create) isBillingSetupOperation_Operation() {}

func (*BillingSetupOperation_Remove) isBillingSetupOperation_Operation() {}

func (m *BillingSetupOperation) GetOperation() isBillingSetupOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *BillingSetupOperation) GetCreate() *resources.BillingSetup {
	if x, ok := m.GetOperation().(*BillingSetupOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *BillingSetupOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*BillingSetupOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*BillingSetupOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*BillingSetupOperation_Create)(nil),
		(*BillingSetupOperation_Remove)(nil),
	}
}

// Response message for a billing setup operation.
type MutateBillingSetupResponse struct {
	// A result that identifies the resource affected by the mutate request.
	Result               *MutateBillingSetupResult `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *MutateBillingSetupResponse) Reset()         { *m = MutateBillingSetupResponse{} }
func (m *MutateBillingSetupResponse) String() string { return proto.CompactTextString(m) }
func (*MutateBillingSetupResponse) ProtoMessage()    {}
func (*MutateBillingSetupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4870990e51228b9d, []int{3}
}

func (m *MutateBillingSetupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateBillingSetupResponse.Unmarshal(m, b)
}
func (m *MutateBillingSetupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateBillingSetupResponse.Marshal(b, m, deterministic)
}
func (m *MutateBillingSetupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateBillingSetupResponse.Merge(m, src)
}
func (m *MutateBillingSetupResponse) XXX_Size() int {
	return xxx_messageInfo_MutateBillingSetupResponse.Size(m)
}
func (m *MutateBillingSetupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateBillingSetupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateBillingSetupResponse proto.InternalMessageInfo

func (m *MutateBillingSetupResponse) GetResult() *MutateBillingSetupResult {
	if m != nil {
		return m.Result
	}
	return nil
}

// Result for a single billing setup mutate.
type MutateBillingSetupResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateBillingSetupResult) Reset()         { *m = MutateBillingSetupResult{} }
func (m *MutateBillingSetupResult) String() string { return proto.CompactTextString(m) }
func (*MutateBillingSetupResult) ProtoMessage()    {}
func (*MutateBillingSetupResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_4870990e51228b9d, []int{4}
}

func (m *MutateBillingSetupResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateBillingSetupResult.Unmarshal(m, b)
}
func (m *MutateBillingSetupResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateBillingSetupResult.Marshal(b, m, deterministic)
}
func (m *MutateBillingSetupResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateBillingSetupResult.Merge(m, src)
}
func (m *MutateBillingSetupResult) XXX_Size() int {
	return xxx_messageInfo_MutateBillingSetupResult.Size(m)
}
func (m *MutateBillingSetupResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateBillingSetupResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateBillingSetupResult proto.InternalMessageInfo

func (m *MutateBillingSetupResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetBillingSetupRequest)(nil), "google.ads.googleads.v2.services.GetBillingSetupRequest")
	proto.RegisterType((*MutateBillingSetupRequest)(nil), "google.ads.googleads.v2.services.MutateBillingSetupRequest")
	proto.RegisterType((*BillingSetupOperation)(nil), "google.ads.googleads.v2.services.BillingSetupOperation")
	proto.RegisterType((*MutateBillingSetupResponse)(nil), "google.ads.googleads.v2.services.MutateBillingSetupResponse")
	proto.RegisterType((*MutateBillingSetupResult)(nil), "google.ads.googleads.v2.services.MutateBillingSetupResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/billing_setup_service.proto", fileDescriptor_4870990e51228b9d)
}

var fileDescriptor_4870990e51228b9d = []byte{
	// 567 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4f, 0x6f, 0xd3, 0x3e,
	0x18, 0xfe, 0xb9, 0x93, 0x2a, 0xcd, 0xfd, 0x21, 0x24, 0x23, 0x20, 0x14, 0x24, 0xaa, 0xb0, 0xc3,
	0xd4, 0x83, 0x2d, 0x82, 0x50, 0x91, 0xb7, 0x0a, 0xa5, 0x97, 0x6e, 0x07, 0x60, 0xca, 0x44, 0x0f,
	0xa8, 0x52, 0x95, 0x25, 0x56, 0x14, 0x29, 0x89, 0x43, 0xec, 0xf4, 0x32, 0xed, 0xc2, 0x8d, 0xf3,
	0xbe, 0x01, 0x47, 0xee, 0x9c, 0xf8, 0x06, 0x3b, 0x70, 0xe1, 0x2b, 0x70, 0xe2, 0x4b, 0x80, 0x12,
	0xdb, 0xfd, 0xc3, 0x52, 0x15, 0x76, 0x7b, 0xfd, 0xbe, 0xef, 0xf3, 0x3c, 0xef, 0x1f, 0xdb, 0xf0,
	0x30, 0xe2, 0x3c, 0x4a, 0x18, 0xf1, 0x43, 0x41, 0x94, 0x59, 0x59, 0x73, 0x87, 0x08, 0x56, 0xcc,
	0xe3, 0x80, 0x09, 0x72, 0x16, 0x27, 0x49, 0x9c, 0x45, 0x33, 0xc1, 0x64, 0x99, 0xcf, 0xb4, 0x1b,
	0xe7, 0x05, 0x97, 0x1c, 0xf5, 0x14, 0x04, 0xfb, 0xa1, 0xc0, 0x0b, 0x34, 0x9e, 0x3b, 0xd8, 0xa0,
	0xbb, 0xcf, 0x37, 0xf1, 0x17, 0x4c, 0xf0, 0xb2, 0xb8, 0x26, 0xa0, 0x88, 0xbb, 0x8f, 0x0c, 0x2c,
	0x8f, 0x89, 0x9f, 0x65, 0x5c, 0xfa, 0x32, 0xe6, 0x99, 0xd0, 0xd1, 0xfb, 0x2b, 0xd1, 0x20, 0x89,
	0x59, 0x26, 0x55, 0xc0, 0x1e, 0xc2, 0x7b, 0x63, 0x26, 0x47, 0x8a, 0xf0, 0xb4, 0xe2, 0xf3, 0xd8,
	0xfb, 0x92, 0x09, 0x89, 0x9e, 0xc0, 0x5b, 0x46, 0x71, 0x96, 0xf9, 0x29, 0xb3, 0x40, 0x0f, 0xec,
	0xef, 0x7a, 0xff, 0x1b, 0xe7, 0x6b, 0x3f, 0x65, 0xf6, 0x25, 0x80, 0x0f, 0x5e, 0x95, 0xd2, 0x97,
	0xac, 0x89, 0xe2, 0x31, 0xec, 0x04, 0xa5, 0x90, 0x3c, 0x65, 0xc5, 0x2c, 0x0e, 0x35, 0x01, 0x34,
	0xae, 0xe3, 0x10, 0xbd, 0x85, 0xbb, 0x3c, 0x67, 0x45, 0x5d, 0xaa, 0xd5, 0xea, 0x81, 0xfd, 0x8e,
	0x33, 0xc0, 0xdb, 0x26, 0x84, 0x57, 0xa5, 0xde, 0x18, 0xb8, 0xb7, 0x64, 0xb2, 0x3f, 0x02, 0x78,
	0xb7, 0x31, 0x09, 0x1d, 0xc3, 0x76, 0x50, 0x30, 0x5f, 0x32, 0xad, 0x46, 0x36, 0xaa, 0x2d, 0xa6,
	0xbd, 0x26, 0x77, 0xf4, 0x9f, 0xa7, 0x09, 0x90, 0x05, 0xdb, 0x05, 0x4b, 0xf9, 0x5c, 0x0f, 0xa6,
	0x8a, 0xa8, 0xf3, 0xa8, 0xb3, 0xd2, 0x95, 0x9d, 0xc3, 0x6e, 0xd3, 0x80, 0x44, 0xce, 0x33, 0xc1,
	0x90, 0x57, 0x91, 0x88, 0x32, 0x91, 0x35, 0x49, 0xc7, 0xa1, 0xdb, 0xbb, 0x6f, 0x64, 0x2b, 0x13,
	0xe9, 0x69, 0x26, 0xfb, 0x25, 0xb4, 0x36, 0xe5, 0xfc, 0xd5, 0x52, 0x9d, 0xaf, 0x3b, 0xf0, 0xce,
	0x2a, 0xf6, 0x54, 0x49, 0xa3, 0x2f, 0x00, 0xde, 0xfe, 0xe3, 0xb2, 0xa0, 0x17, 0xdb, 0x0b, 0x6e,
	0xbe, 0x5f, 0xdd, 0x7f, 0x1d, 0xbd, 0x3d, 0xf8, 0xf0, 0xfd, 0xc7, 0x65, 0xeb, 0x29, 0x22, 0xd5,
	0x63, 0x38, 0x5f, 0x6b, 0x63, 0x68, 0xee, 0x94, 0x20, 0x7d, 0xf3, 0x3a, 0x6a, 0x90, 0x20, 0xfd,
	0x0b, 0xf4, 0x0d, 0x40, 0x74, 0x7d, 0x22, 0xe8, 0xe0, 0x66, 0xb3, 0x56, 0xd5, 0x1f, 0xde, 0x70,
	0x51, 0xf5, 0xda, 0xed, 0x61, 0xdd, 0xca, 0xc0, 0x76, 0xaa, 0x56, 0x96, 0xb5, 0x9f, 0xaf, 0xbc,
	0x96, 0x61, 0xff, 0x62, 0xbd, 0x13, 0x9a, 0xd6, 0x7c, 0x14, 0xf4, 0xbb, 0x0f, 0xaf, 0x5c, 0x6b,
	0xa9, 0xa9, 0xad, 0x3c, 0x16, 0x38, 0xe0, 0xe9, 0xe8, 0x17, 0x80, 0x7b, 0x01, 0x4f, 0xb7, 0xd6,
	0x37, 0xb2, 0x1a, 0x76, 0x7c, 0x52, 0x7d, 0x0a, 0x27, 0xe0, 0xdd, 0x91, 0x46, 0x47, 0x3c, 0xf1,
	0xb3, 0x08, 0xf3, 0x22, 0x22, 0x11, 0xcb, 0xea, 0x2f, 0x83, 0x2c, 0xf5, 0x36, 0xff, 0x81, 0x07,
	0xc6, 0xf8, 0xd4, 0xda, 0x19, 0xbb, 0xee, 0xe7, 0x56, 0x6f, 0xac, 0x08, 0xdd, 0x50, 0x60, 0x65,
	0x56, 0xd6, 0xc4, 0xc1, 0x5a, 0x58, 0x5c, 0x99, 0x94, 0xa9, 0x1b, 0x8a, 0xe9, 0x22, 0x65, 0x3a,
	0x71, 0xa6, 0x26, 0xe5, 0x67, 0x6b, 0x4f, 0xf9, 0x29, 0x75, 0x43, 0x41, 0xe9, 0x22, 0x89, 0xd2,
	0x89, 0x43, 0xa9, 0x49, 0x3b, 0x6b, 0xd7, 0x75, 0x3e, 0xfb, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x91,
	0x28, 0x5a, 0x53, 0xaa, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BillingSetupServiceClient is the client API for BillingSetupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BillingSetupServiceClient interface {
	// Returns a billing setup.
	GetBillingSetup(ctx context.Context, in *GetBillingSetupRequest, opts ...grpc.CallOption) (*resources.BillingSetup, error)
	// Creates a billing setup, or cancels an existing billing setup.
	MutateBillingSetup(ctx context.Context, in *MutateBillingSetupRequest, opts ...grpc.CallOption) (*MutateBillingSetupResponse, error)
}

type billingSetupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBillingSetupServiceClient(cc grpc.ClientConnInterface) BillingSetupServiceClient {
	return &billingSetupServiceClient{cc}
}

func (c *billingSetupServiceClient) GetBillingSetup(ctx context.Context, in *GetBillingSetupRequest, opts ...grpc.CallOption) (*resources.BillingSetup, error) {
	out := new(resources.BillingSetup)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.BillingSetupService/GetBillingSetup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingSetupServiceClient) MutateBillingSetup(ctx context.Context, in *MutateBillingSetupRequest, opts ...grpc.CallOption) (*MutateBillingSetupResponse, error) {
	out := new(MutateBillingSetupResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.BillingSetupService/MutateBillingSetup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BillingSetupServiceServer is the server API for BillingSetupService service.
type BillingSetupServiceServer interface {
	// Returns a billing setup.
	GetBillingSetup(context.Context, *GetBillingSetupRequest) (*resources.BillingSetup, error)
	// Creates a billing setup, or cancels an existing billing setup.
	MutateBillingSetup(context.Context, *MutateBillingSetupRequest) (*MutateBillingSetupResponse, error)
}

// UnimplementedBillingSetupServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBillingSetupServiceServer struct {
}

func (*UnimplementedBillingSetupServiceServer) GetBillingSetup(ctx context.Context, req *GetBillingSetupRequest) (*resources.BillingSetup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBillingSetup not implemented")
}
func (*UnimplementedBillingSetupServiceServer) MutateBillingSetup(ctx context.Context, req *MutateBillingSetupRequest) (*MutateBillingSetupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateBillingSetup not implemented")
}

func RegisterBillingSetupServiceServer(s *grpc.Server, srv BillingSetupServiceServer) {
	s.RegisterService(&_BillingSetupService_serviceDesc, srv)
}

func _BillingSetupService_GetBillingSetup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBillingSetupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingSetupServiceServer).GetBillingSetup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.BillingSetupService/GetBillingSetup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingSetupServiceServer).GetBillingSetup(ctx, req.(*GetBillingSetupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BillingSetupService_MutateBillingSetup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateBillingSetupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingSetupServiceServer).MutateBillingSetup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.BillingSetupService/MutateBillingSetup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingSetupServiceServer).MutateBillingSetup(ctx, req.(*MutateBillingSetupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BillingSetupService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.BillingSetupService",
	HandlerType: (*BillingSetupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBillingSetup",
			Handler:    _BillingSetupService_GetBillingSetup_Handler,
		},
		{
			MethodName: "MutateBillingSetup",
			Handler:    _BillingSetupService_MutateBillingSetup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/billing_setup_service.proto",
}
