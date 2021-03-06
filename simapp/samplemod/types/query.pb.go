// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: samplemod/query.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_datachainlab_cross_x_core_account_types "github.com/datachainlab/cross/x/core/account/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type QueryCounterRequest struct {
	Account github_com_datachainlab_cross_x_core_account_types.AccountID `protobuf:"bytes,1,opt,name=account,proto3,casttype=github.com/datachainlab/cross/x/core/account/types.AccountID" json:"account,omitempty"`
}

func (m *QueryCounterRequest) Reset()         { *m = QueryCounterRequest{} }
func (m *QueryCounterRequest) String() string { return proto.CompactTextString(m) }
func (*QueryCounterRequest) ProtoMessage()    {}
func (*QueryCounterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_416f9a1090665dae, []int{0}
}
func (m *QueryCounterRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCounterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCounterRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCounterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCounterRequest.Merge(m, src)
}
func (m *QueryCounterRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryCounterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCounterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCounterRequest proto.InternalMessageInfo

type QueryCounterResponse struct {
	Value uint64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *QueryCounterResponse) Reset()         { *m = QueryCounterResponse{} }
func (m *QueryCounterResponse) String() string { return proto.CompactTextString(m) }
func (*QueryCounterResponse) ProtoMessage()    {}
func (*QueryCounterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_416f9a1090665dae, []int{1}
}
func (m *QueryCounterResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCounterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCounterResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCounterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCounterResponse.Merge(m, src)
}
func (m *QueryCounterResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryCounterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCounterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCounterResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*QueryCounterRequest)(nil), "samplemod.types.QueryCounterRequest")
	proto.RegisterType((*QueryCounterResponse)(nil), "samplemod.types.QueryCounterResponse")
}

func init() { proto.RegisterFile("samplemod/query.proto", fileDescriptor_416f9a1090665dae) }

var fileDescriptor_416f9a1090665dae = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xc1, 0x4a, 0x33, 0x31,
	0x14, 0x85, 0x67, 0x7e, 0xfe, 0x5a, 0x08, 0x82, 0x30, 0x56, 0x28, 0x83, 0xc4, 0x52, 0x14, 0x5c,
	0x48, 0x2e, 0xa8, 0x4b, 0x17, 0x5a, 0xdd, 0xb8, 0x74, 0x96, 0xdd, 0xa5, 0x69, 0x98, 0x0e, 0xcc,
	0xe4, 0xa6, 0x93, 0x8c, 0xb4, 0x4b, 0x7d, 0x02, 0xc1, 0x97, 0xea, 0xb2, 0xe0, 0xc6, 0x95, 0x68,
	0xeb, 0x53, 0xb8, 0x92, 0x26, 0x55, 0xb1, 0x88, 0xee, 0x72, 0x0f, 0xdf, 0x3d, 0x27, 0xf7, 0x90,
	0x2d, 0xc3, 0x0b, 0x9d, 0xcb, 0x02, 0xfb, 0x30, 0xac, 0x64, 0x39, 0x66, 0xba, 0x44, 0x8b, 0xd1,
	0xc6, 0xa7, 0xcc, 0xec, 0x58, 0x4b, 0x13, 0x6f, 0xa7, 0x88, 0x69, 0x2e, 0x81, 0xeb, 0x0c, 0xb8,
	0x52, 0x68, 0xb9, 0xcd, 0x50, 0x19, 0x8f, 0xc7, 0x8d, 0x14, 0x53, 0x74, 0x4f, 0x58, 0xbc, 0xbc,
	0xda, 0x1e, 0x92, 0xcd, 0xab, 0x85, 0xe7, 0x39, 0x56, 0xca, 0xca, 0x32, 0x91, 0xc3, 0x4a, 0x1a,
	0x1b, 0x75, 0x49, 0x9d, 0x0b, 0xb1, 0xd0, 0x9a, 0x61, 0x2b, 0xdc, 0x5f, 0xef, 0x9c, 0xbe, 0x3d,
	0xed, 0x9c, 0xa4, 0x99, 0x1d, 0x54, 0x3d, 0x26, 0xb0, 0x80, 0x3e, 0xb7, 0x5c, 0x0c, 0x78, 0xa6,
	0x72, 0xde, 0x03, 0x51, 0xa2, 0x31, 0x30, 0x02, 0x81, 0xa5, 0x84, 0xe5, 0x1e, 0xb8, 0x3f, 0xb1,
	0x33, 0x3f, 0x5d, 0x5e, 0x24, 0x1f, 0x86, 0xed, 0x03, 0xd2, 0xf8, 0x1e, 0x69, 0x34, 0x2a, 0x23,
	0xa3, 0x06, 0xa9, 0x5d, 0xf3, 0xbc, 0x92, 0x2e, 0xf1, 0x7f, 0xe2, 0x87, 0xc3, 0x9b, 0x90, 0xd4,
	0x1c, 0x1e, 0x8d, 0x48, 0x7d, 0xb9, 0x12, 0xed, 0xb2, 0x95, 0xdb, 0xd9, 0x0f, 0x47, 0xc4, 0x7b,
	0x7f, 0x50, 0x3e, 0xb7, 0xdd, 0xba, 0x7d, 0x78, 0xbd, 0xff, 0x17, 0x47, 0x4d, 0xf8, 0xea, 0xd9,
	0xe1, 0x20, 0x3c, 0xd9, 0x49, 0x26, 0x2f, 0x34, 0x98, 0xcc, 0x68, 0x38, 0x9d, 0xd1, 0xf0, 0x79,
	0x46, 0xc3, 0xbb, 0x39, 0x0d, 0xa6, 0x73, 0x1a, 0x3c, 0xce, 0x69, 0xd0, 0x3d, 0xfe, 0xbd, 0x16,
	0x93, 0x15, 0x5c, 0xeb, 0x55, 0xef, 0xde, 0x9a, 0xeb, 0xff, 0xe8, 0x3d, 0x00, 0x00, 0xff, 0xff,
	0xcf, 0x9e, 0x47, 0x65, 0xdd, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	Counter(ctx context.Context, in *QueryCounterRequest, opts ...grpc.CallOption) (*QueryCounterResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Counter(ctx context.Context, in *QueryCounterRequest, opts ...grpc.CallOption) (*QueryCounterResponse, error) {
	out := new(QueryCounterResponse)
	err := c.cc.Invoke(ctx, "/samplemod.types.Query/Counter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	Counter(context.Context, *QueryCounterRequest) (*QueryCounterResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Counter(ctx context.Context, req *QueryCounterRequest) (*QueryCounterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Counter not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Counter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCounterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Counter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/samplemod.types.Query/Counter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Counter(ctx, req.(*QueryCounterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "samplemod.types.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Counter",
			Handler:    _Query_Counter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "samplemod/query.proto",
}

func (m *QueryCounterRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCounterRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCounterRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Account) > 0 {
		i -= len(m.Account)
		copy(dAtA[i:], m.Account)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Account)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryCounterResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCounterResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCounterResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Value != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Value))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryCounterRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Account)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryCounterResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value != 0 {
		n += 1 + sovQuery(uint64(m.Value))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryCounterRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryCounterRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCounterRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Account = append(m.Account[:0], dAtA[iNdEx:postIndex]...)
			if m.Account == nil {
				m.Account = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryCounterResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryCounterResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCounterResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			m.Value = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Value |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
