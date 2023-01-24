// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: internal/kit_test/kafka/kafka_test.proto

package kafka

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type TestMsg struct {
	A                    string   `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
	B                    int64    `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestMsg) Reset()         { *m = TestMsg{} }
func (m *TestMsg) String() string { return proto.CompactTextString(m) }
func (*TestMsg) ProtoMessage()    {}
func (*TestMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_66689f4a742b44db, []int{0}
}
func (m *TestMsg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TestMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TestMsg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TestMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestMsg.Merge(m, src)
}
func (m *TestMsg) XXX_Size() int {
	return m.Size()
}
func (m *TestMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_TestMsg.DiscardUnknown(m)
}

var xxx_messageInfo_TestMsg proto.InternalMessageInfo

func (m *TestMsg) GetA() string {
	if m != nil {
		return m.A
	}
	return ""
}

func (m *TestMsg) GetB() int64 {
	if m != nil {
		return m.B
	}
	return 0
}

func init() {
	proto.RegisterType((*TestMsg)(nil), "kafka.TestMsg")
}

func init() {
	proto.RegisterFile("internal/kit_test/kafka/kafka_test.proto", fileDescriptor_66689f4a742b44db)
}

var fileDescriptor_66689f4a742b44db = []byte{
	// 123 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xc8, 0xcc, 0x2b, 0x49,
	0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0xcf, 0xce, 0x2c, 0x89, 0x2f, 0x49, 0x2d, 0x2e, 0xd1, 0xcf, 0x4e,
	0x4c, 0xcb, 0x4e, 0x84, 0x90, 0x60, 0x01, 0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21, 0x56, 0xb0,
	0x88, 0x92, 0x2a, 0x17, 0x7b, 0x48, 0x6a, 0x71, 0x89, 0x6f, 0x71, 0xba, 0x10, 0x0f, 0x17, 0x63,
	0xa2, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x63, 0x22, 0x88, 0x97, 0x24, 0xc1, 0xa4, 0xc0,
	0xa8, 0xc1, 0x1c, 0xc4, 0x98, 0xe4, 0x24, 0x70, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c,
	0x0f, 0x1e, 0xc9, 0x31, 0xce, 0x78, 0x2c, 0xc7, 0x90, 0xc4, 0x06, 0x36, 0xc6, 0x18, 0x10, 0x00,
	0x00, 0xff, 0xff, 0x2c, 0x84, 0x72, 0xa7, 0x72, 0x00, 0x00, 0x00,
}

func (m *TestMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TestMsg) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TestMsg) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.B != 0 {
		i = encodeVarintKafkaTest(dAtA, i, uint64(m.B))
		i--
		dAtA[i] = 0x10
	}
	if len(m.A) > 0 {
		i -= len(m.A)
		copy(dAtA[i:], m.A)
		i = encodeVarintKafkaTest(dAtA, i, uint64(len(m.A)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintKafkaTest(dAtA []byte, offset int, v uint64) int {
	offset -= sovKafkaTest(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TestMsg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.A)
	if l > 0 {
		n += 1 + l + sovKafkaTest(uint64(l))
	}
	if m.B != 0 {
		n += 1 + sovKafkaTest(uint64(m.B))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovKafkaTest(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozKafkaTest(x uint64) (n int) {
	return sovKafkaTest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TestMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKafkaTest
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
			return fmt.Errorf("proto: TestMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TestMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field A", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKafkaTest
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
				return ErrInvalidLengthKafkaTest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthKafkaTest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.A = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field B", wireType)
			}
			m.B = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKafkaTest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.B |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipKafkaTest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthKafkaTest
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipKafkaTest(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowKafkaTest
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
					return 0, ErrIntOverflowKafkaTest
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
					return 0, ErrIntOverflowKafkaTest
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
				return 0, ErrInvalidLengthKafkaTest
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupKafkaTest
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthKafkaTest
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthKafkaTest        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowKafkaTest          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupKafkaTest = fmt.Errorf("proto: unexpected end of group")
)
