// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pkg/mq/kafka/kafka.proto

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

type KafkaMessage struct {
	TxId                 string   `protobuf:"bytes,1,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
	TraceId              string   `protobuf:"bytes,2,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	EventType            int32    `protobuf:"varint,3,opt,name=event_type,json=eventType,proto3" json:"event_type,omitempty"`
	Message              []byte   `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KafkaMessage) Reset()         { *m = KafkaMessage{} }
func (m *KafkaMessage) String() string { return proto.CompactTextString(m) }
func (*KafkaMessage) ProtoMessage()    {}
func (*KafkaMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6dca8070e0eac55, []int{0}
}
func (m *KafkaMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *KafkaMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_KafkaMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *KafkaMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KafkaMessage.Merge(m, src)
}
func (m *KafkaMessage) XXX_Size() int {
	return m.Size()
}
func (m *KafkaMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_KafkaMessage.DiscardUnknown(m)
}

var xxx_messageInfo_KafkaMessage proto.InternalMessageInfo

func (m *KafkaMessage) GetTxId() string {
	if m != nil {
		return m.TxId
	}
	return ""
}

func (m *KafkaMessage) GetTraceId() string {
	if m != nil {
		return m.TraceId
	}
	return ""
}

func (m *KafkaMessage) GetEventType() int32 {
	if m != nil {
		return m.EventType
	}
	return 0
}

func (m *KafkaMessage) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

func init() {
	proto.RegisterType((*KafkaMessage)(nil), "kafka.KafkaMessage")
}

func init() { proto.RegisterFile("pkg/mq/kafka/kafka.proto", fileDescriptor_c6dca8070e0eac55) }

var fileDescriptor_c6dca8070e0eac55 = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x28, 0xc8, 0x4e, 0xd7,
	0xcf, 0x2d, 0xd4, 0xcf, 0x4e, 0x4c, 0xcb, 0x4e, 0x84, 0x90, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9,
	0x42, 0xac, 0x60, 0x8e, 0x52, 0x39, 0x17, 0x8f, 0x37, 0x88, 0xe1, 0x9b, 0x5a, 0x5c, 0x9c, 0x98,
	0x9e, 0x2a, 0x24, 0xcc, 0xc5, 0x5a, 0x52, 0x11, 0x9f, 0x99, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1,
	0x19, 0xc4, 0x52, 0x52, 0xe1, 0x99, 0x22, 0x24, 0xc9, 0xc5, 0x51, 0x52, 0x94, 0x98, 0x9c, 0x0a,
	0x12, 0x67, 0x02, 0x8b, 0xb3, 0x83, 0xf9, 0x9e, 0x29, 0x42, 0xb2, 0x5c, 0x5c, 0xa9, 0x65, 0xa9,
	0x79, 0x25, 0xf1, 0x25, 0x95, 0x05, 0xa9, 0x12, 0xcc, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x9c, 0x60,
	0x91, 0x90, 0xca, 0x82, 0x54, 0x21, 0x09, 0x2e, 0xf6, 0x5c, 0x88, 0xc9, 0x12, 0x2c, 0x0a, 0x8c,
	0x1a, 0x3c, 0x41, 0x30, 0xae, 0x93, 0xc0, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e,
	0x78, 0x24, 0xc7, 0x38, 0xe3, 0xb1, 0x1c, 0x43, 0x12, 0x1b, 0xd8, 0x61, 0xc6, 0x80, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x4f, 0xad, 0x24, 0x0e, 0xb4, 0x00, 0x00, 0x00,
}

func (m *KafkaMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KafkaMessage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *KafkaMessage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintKafka(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0x22
	}
	if m.EventType != 0 {
		i = encodeVarintKafka(dAtA, i, uint64(m.EventType))
		i--
		dAtA[i] = 0x18
	}
	if len(m.TraceId) > 0 {
		i -= len(m.TraceId)
		copy(dAtA[i:], m.TraceId)
		i = encodeVarintKafka(dAtA, i, uint64(len(m.TraceId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.TxId) > 0 {
		i -= len(m.TxId)
		copy(dAtA[i:], m.TxId)
		i = encodeVarintKafka(dAtA, i, uint64(len(m.TxId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintKafka(dAtA []byte, offset int, v uint64) int {
	offset -= sovKafka(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *KafkaMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TxId)
	if l > 0 {
		n += 1 + l + sovKafka(uint64(l))
	}
	l = len(m.TraceId)
	if l > 0 {
		n += 1 + l + sovKafka(uint64(l))
	}
	if m.EventType != 0 {
		n += 1 + sovKafka(uint64(m.EventType))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovKafka(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovKafka(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozKafka(x uint64) (n int) {
	return sovKafka(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *KafkaMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKafka
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
			return fmt.Errorf("proto: KafkaMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KafkaMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKafka
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
				return ErrInvalidLengthKafka
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthKafka
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TraceId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKafka
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
				return ErrInvalidLengthKafka
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthKafka
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TraceId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EventType", wireType)
			}
			m.EventType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKafka
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EventType |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKafka
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
				return ErrInvalidLengthKafka
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthKafka
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = append(m.Message[:0], dAtA[iNdEx:postIndex]...)
			if m.Message == nil {
				m.Message = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipKafka(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthKafka
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
func skipKafka(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowKafka
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
					return 0, ErrIntOverflowKafka
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
					return 0, ErrIntOverflowKafka
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
				return 0, ErrInvalidLengthKafka
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupKafka
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthKafka
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthKafka        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowKafka          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupKafka = fmt.Errorf("proto: unexpected end of group")
)
