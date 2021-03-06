// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lookout/sdk/service_analyzer.proto

package pb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type EventResponse struct {
	AnalyzerVersion string     `protobuf:"bytes,1,opt,name=analyzer_version,json=analyzerVersion,proto3" json:"analyzer_version,omitempty"`
	Comments        []*Comment `protobuf:"bytes,2,rep,name=comments,proto3" json:"comments,omitempty"`
}

func (m *EventResponse) Reset()         { *m = EventResponse{} }
func (m *EventResponse) String() string { return proto.CompactTextString(m) }
func (*EventResponse) ProtoMessage()    {}
func (*EventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_analyzer_7c2ea649f74307b7, []int{0}
}
func (m *EventResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *EventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventResponse.Merge(dst, src)
}
func (m *EventResponse) XXX_Size() int {
	return m.Size()
}
func (m *EventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EventResponse proto.InternalMessageInfo

// Comment is a comment on a commit or changeset.
type Comment struct {
	// File this comment belongs to. If empty, it is a global comment.
	File string `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
	// Line this comment refers to. If 0 (and file is set), it is a
	// file-level comment. Line is expressed aqs a 1-based index.
	Line int32 `protobuf:"varint,2,opt,name=line,proto3" json:"line,omitempty"`
	// Text of the comment.
	Text string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	// Confidence in the comment. It should be an integer between 0 and 100.
	Confidence uint32 `protobuf:"varint,4,opt,name=confidence,proto3" json:"confidence,omitempty"`
}

func (m *Comment) Reset()         { *m = Comment{} }
func (m *Comment) String() string { return proto.CompactTextString(m) }
func (*Comment) ProtoMessage()    {}
func (*Comment) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_analyzer_7c2ea649f74307b7, []int{1}
}
func (m *Comment) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Comment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Comment.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Comment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Comment.Merge(dst, src)
}
func (m *Comment) XXX_Size() int {
	return m.Size()
}
func (m *Comment) XXX_DiscardUnknown() {
	xxx_messageInfo_Comment.DiscardUnknown(m)
}

var xxx_messageInfo_Comment proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EventResponse)(nil), "pb.EventResponse")
	proto.RegisterType((*Comment)(nil), "pb.Comment")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AnalyzerClient is the client API for Analyzer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AnalyzerClient interface {
	NotifyReviewEvent(ctx context.Context, in *ReviewEvent, opts ...grpc.CallOption) (*EventResponse, error)
	NotifyPushEvent(ctx context.Context, in *PushEvent, opts ...grpc.CallOption) (*EventResponse, error)
}

type analyzerClient struct {
	cc *grpc.ClientConn
}

func NewAnalyzerClient(cc *grpc.ClientConn) AnalyzerClient {
	return &analyzerClient{cc}
}

func (c *analyzerClient) NotifyReviewEvent(ctx context.Context, in *ReviewEvent, opts ...grpc.CallOption) (*EventResponse, error) {
	out := new(EventResponse)
	err := c.cc.Invoke(ctx, "/pb.Analyzer/NotifyReviewEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyzerClient) NotifyPushEvent(ctx context.Context, in *PushEvent, opts ...grpc.CallOption) (*EventResponse, error) {
	out := new(EventResponse)
	err := c.cc.Invoke(ctx, "/pb.Analyzer/NotifyPushEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AnalyzerServer is the server API for Analyzer service.
type AnalyzerServer interface {
	NotifyReviewEvent(context.Context, *ReviewEvent) (*EventResponse, error)
	NotifyPushEvent(context.Context, *PushEvent) (*EventResponse, error)
}

func RegisterAnalyzerServer(s *grpc.Server, srv AnalyzerServer) {
	s.RegisterService(&_Analyzer_serviceDesc, srv)
}

func _Analyzer_NotifyReviewEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReviewEvent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyzerServer).NotifyReviewEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Analyzer/NotifyReviewEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyzerServer).NotifyReviewEvent(ctx, req.(*ReviewEvent))
	}
	return interceptor(ctx, in, info, handler)
}

func _Analyzer_NotifyPushEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushEvent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyzerServer).NotifyPushEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Analyzer/NotifyPushEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyzerServer).NotifyPushEvent(ctx, req.(*PushEvent))
	}
	return interceptor(ctx, in, info, handler)
}

var _Analyzer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Analyzer",
	HandlerType: (*AnalyzerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NotifyReviewEvent",
			Handler:    _Analyzer_NotifyReviewEvent_Handler,
		},
		{
			MethodName: "NotifyPushEvent",
			Handler:    _Analyzer_NotifyPushEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lookout/sdk/service_analyzer.proto",
}

func (m *EventResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.AnalyzerVersion) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintServiceAnalyzer(dAtA, i, uint64(len(m.AnalyzerVersion)))
		i += copy(dAtA[i:], m.AnalyzerVersion)
	}
	if len(m.Comments) > 0 {
		for _, msg := range m.Comments {
			dAtA[i] = 0x12
			i++
			i = encodeVarintServiceAnalyzer(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *Comment) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Comment) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.File) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintServiceAnalyzer(dAtA, i, uint64(len(m.File)))
		i += copy(dAtA[i:], m.File)
	}
	if m.Line != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintServiceAnalyzer(dAtA, i, uint64(m.Line))
	}
	if len(m.Text) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintServiceAnalyzer(dAtA, i, uint64(len(m.Text)))
		i += copy(dAtA[i:], m.Text)
	}
	if m.Confidence != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintServiceAnalyzer(dAtA, i, uint64(m.Confidence))
	}
	return i, nil
}

func encodeVarintServiceAnalyzer(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *EventResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AnalyzerVersion)
	if l > 0 {
		n += 1 + l + sovServiceAnalyzer(uint64(l))
	}
	if len(m.Comments) > 0 {
		for _, e := range m.Comments {
			l = e.Size()
			n += 1 + l + sovServiceAnalyzer(uint64(l))
		}
	}
	return n
}

func (m *Comment) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.File)
	if l > 0 {
		n += 1 + l + sovServiceAnalyzer(uint64(l))
	}
	if m.Line != 0 {
		n += 1 + sovServiceAnalyzer(uint64(m.Line))
	}
	l = len(m.Text)
	if l > 0 {
		n += 1 + l + sovServiceAnalyzer(uint64(l))
	}
	if m.Confidence != 0 {
		n += 1 + sovServiceAnalyzer(uint64(m.Confidence))
	}
	return n
}

func sovServiceAnalyzer(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozServiceAnalyzer(x uint64) (n int) {
	return sovServiceAnalyzer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowServiceAnalyzer
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EventResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AnalyzerVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceAnalyzer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthServiceAnalyzer
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AnalyzerVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Comments", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceAnalyzer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthServiceAnalyzer
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Comments = append(m.Comments, &Comment{})
			if err := m.Comments[len(m.Comments)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipServiceAnalyzer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthServiceAnalyzer
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
func (m *Comment) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowServiceAnalyzer
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Comment: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Comment: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field File", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceAnalyzer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthServiceAnalyzer
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.File = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Line", wireType)
			}
			m.Line = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceAnalyzer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Line |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Text", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceAnalyzer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthServiceAnalyzer
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Text = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Confidence", wireType)
			}
			m.Confidence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceAnalyzer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Confidence |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipServiceAnalyzer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthServiceAnalyzer
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
func skipServiceAnalyzer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowServiceAnalyzer
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
					return 0, ErrIntOverflowServiceAnalyzer
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowServiceAnalyzer
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthServiceAnalyzer
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowServiceAnalyzer
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipServiceAnalyzer(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthServiceAnalyzer = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowServiceAnalyzer   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("lookout/sdk/service_analyzer.proto", fileDescriptor_service_analyzer_7c2ea649f74307b7)
}

var fileDescriptor_service_analyzer_7c2ea649f74307b7 = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xb1, 0x6e, 0xfa, 0x30,
	0x10, 0xc6, 0x63, 0xe0, 0xff, 0x2f, 0x35, 0x42, 0x14, 0x2f, 0x8d, 0x50, 0x65, 0x45, 0x2c, 0x4d,
	0x87, 0x26, 0x12, 0x0c, 0x9d, 0xdb, 0xaa, 0x6b, 0x55, 0x65, 0xe8, 0x8a, 0x48, 0xb8, 0x80, 0x45,
	0xb0, 0xa3, 0xd8, 0x49, 0xa1, 0x4f, 0xd1, 0xc7, 0x62, 0x64, 0xec, 0xd8, 0xc2, 0x8b, 0x54, 0xb6,
	0x01, 0x51, 0xa9, 0xdb, 0xf7, 0xfd, 0xee, 0xbe, 0xf3, 0x9d, 0x71, 0x3f, 0x13, 0x62, 0x2e, 0x4a,
	0x15, 0xca, 0xc9, 0x3c, 0x94, 0x50, 0x54, 0x2c, 0x81, 0xd1, 0x98, 0x8f, 0xb3, 0xd5, 0x3b, 0x14,
	0x41, 0x5e, 0x08, 0x25, 0x48, 0x2d, 0x8f, 0x7b, 0xb7, 0x53, 0xa6, 0x66, 0x65, 0x1c, 0x24, 0x62,
	0x11, 0x4e, 0xc5, 0x54, 0x84, 0xa6, 0x14, 0x97, 0xa9, 0x71, 0xc6, 0x18, 0x65, 0x23, 0xbd, 0xcb,
	0xd3, 0xb1, 0x50, 0x01, 0x57, 0xb6, 0xd0, 0x4f, 0x70, 0xfb, 0x49, 0xdb, 0x08, 0x64, 0x2e, 0xb8,
	0x04, 0x72, 0x83, 0x2f, 0x0e, 0xcf, 0x8d, 0x2a, 0x28, 0x24, 0x13, 0xdc, 0x45, 0x1e, 0xf2, 0xcf,
	0xa3, 0xce, 0x81, 0xbf, 0x5a, 0x4c, 0xae, 0x71, 0x33, 0x11, 0x8b, 0x05, 0x70, 0x25, 0xdd, 0x9a,
	0x57, 0xf7, 0x5b, 0x83, 0x56, 0x90, 0xc7, 0xc1, 0xa3, 0x65, 0xd1, 0xb1, 0xd8, 0x07, 0x7c, 0xb6,
	0x87, 0x84, 0xe0, 0x46, 0xca, 0x32, 0xd8, 0x8f, 0x34, 0x5a, 0xb3, 0x8c, 0x71, 0x70, 0x6b, 0x1e,
	0xf2, 0xff, 0x45, 0x46, 0x6b, 0xa6, 0x60, 0xa9, 0xdc, 0xba, 0xed, 0xd3, 0x9a, 0x50, 0x8c, 0x13,
	0xc1, 0x53, 0x36, 0x01, 0x9e, 0x80, 0xdb, 0xf0, 0x90, 0xdf, 0x8e, 0x4e, 0xc8, 0x60, 0x89, 0x9b,
	0xf7, 0xfb, 0x15, 0xc9, 0x1d, 0xee, 0x3e, 0x0b, 0xc5, 0xd2, 0x55, 0x04, 0x15, 0x83, 0x37, 0x73,
	0x23, 0xe9, 0xe8, 0xf5, 0x4e, 0x40, 0xaf, 0xab, 0xc1, 0xef, 0xfb, 0x87, 0xb8, 0x63, 0x83, 0x2f,
	0xa5, 0x9c, 0xd9, 0x58, 0x5b, 0x77, 0x1d, 0xed, 0x1f, 0xa1, 0x87, 0xab, 0xf5, 0x37, 0x75, 0xd6,
	0x5b, 0x8a, 0x36, 0x5b, 0x8a, 0xbe, 0xb6, 0x14, 0x7d, 0xec, 0xa8, 0xb3, 0xd9, 0x51, 0xe7, 0x73,
	0x47, 0x9d, 0xf8, 0xbf, 0xf9, 0xea, 0xe1, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7c, 0x76, 0xb2,
	0xd1, 0xdc, 0x01, 0x00, 0x00,
}
