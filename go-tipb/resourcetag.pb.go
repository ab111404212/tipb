// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: resourcetag.proto

package tipb

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
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

type ResourceGroupTagLabel int32

const (
	ResourceGroupTagLabel_ResourceGroupTagLabelUnknown ResourceGroupTagLabel = 0
	ResourceGroupTagLabel_ResourceGroupTagLabelRow     ResourceGroupTagLabel = 1
	ResourceGroupTagLabel_ResourceGroupTagLabelIndex   ResourceGroupTagLabel = 2
)

var ResourceGroupTagLabel_name = map[int32]string{
	0: "ResourceGroupTagLabelUnknown",
	1: "ResourceGroupTagLabelRow",
	2: "ResourceGroupTagLabelIndex",
}

var ResourceGroupTagLabel_value = map[string]int32{
	"ResourceGroupTagLabelUnknown": 0,
	"ResourceGroupTagLabelRow":     1,
	"ResourceGroupTagLabelIndex":   2,
}

func (x ResourceGroupTagLabel) Enum() *ResourceGroupTagLabel {
	p := new(ResourceGroupTagLabel)
	*p = x
	return p
}

func (x ResourceGroupTagLabel) String() string {
	return proto.EnumName(ResourceGroupTagLabel_name, int32(x))
}

func (x *ResourceGroupTagLabel) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ResourceGroupTagLabel_value, data, "ResourceGroupTagLabel")
	if err != nil {
		return err
	}
	*x = ResourceGroupTagLabel(value)
	return nil
}

func (ResourceGroupTagLabel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5c1d3c692714f1ff, []int{0}
}

type ResourceGroupTag struct {
	SqlDigest  []byte `protobuf:"bytes,1,opt,name=sql_digest,json=sqlDigest" json:"sql_digest,omitempty"`
	PlanDigest []byte `protobuf:"bytes,2,opt,name=plan_digest,json=planDigest" json:"plan_digest,omitempty"`
	// Use to label the handling kv type of the request.
	// This is for TiKV resource_metering to collect execution information by the key label.
	Label                *ResourceGroupTagLabel `protobuf:"varint,3,opt,name=label,enum=tipb.ResourceGroupTagLabel" json:"label,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ResourceGroupTag) Reset()         { *m = ResourceGroupTag{} }
func (m *ResourceGroupTag) String() string { return proto.CompactTextString(m) }
func (*ResourceGroupTag) ProtoMessage()    {}
func (*ResourceGroupTag) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c1d3c692714f1ff, []int{0}
}
func (m *ResourceGroupTag) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResourceGroupTag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResourceGroupTag.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResourceGroupTag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceGroupTag.Merge(m, src)
}
func (m *ResourceGroupTag) XXX_Size() int {
	return m.Size()
}
func (m *ResourceGroupTag) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceGroupTag.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceGroupTag proto.InternalMessageInfo

func (m *ResourceGroupTag) GetSqlDigest() []byte {
	if m != nil {
		return m.SqlDigest
	}
	return nil
}

func (m *ResourceGroupTag) GetPlanDigest() []byte {
	if m != nil {
		return m.PlanDigest
	}
	return nil
}

func (m *ResourceGroupTag) GetLabel() ResourceGroupTagLabel {
	if m != nil && m.Label != nil {
		return *m.Label
	}
	return ResourceGroupTagLabel_ResourceGroupTagLabelUnknown
}

func init() {
	proto.RegisterEnum("tipb.ResourceGroupTagLabel", ResourceGroupTagLabel_name, ResourceGroupTagLabel_value)
	proto.RegisterType((*ResourceGroupTag)(nil), "tipb.ResourceGroupTag")
}

func init() { proto.RegisterFile("resourcetag.proto", fileDescriptor_5c1d3c692714f1ff) }

var fileDescriptor_5c1d3c692714f1ff = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x4a, 0x2d, 0xce,
	0x2f, 0x2d, 0x4a, 0x4e, 0x2d, 0x49, 0x4c, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x29,
	0xc9, 0x2c, 0x48, 0x92, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x0b, 0xe8, 0x83, 0x58, 0x10, 0x39,
	0xa5, 0x56, 0x46, 0x2e, 0x81, 0x20, 0xa8, 0x0e, 0xf7, 0xa2, 0xfc, 0xd2, 0x82, 0x90, 0xc4, 0x74,
	0x21, 0x59, 0x2e, 0xae, 0xe2, 0xc2, 0x9c, 0xf8, 0x94, 0xcc, 0xf4, 0xd4, 0xe2, 0x12, 0x09, 0x46,
	0x05, 0x46, 0x0d, 0x9e, 0x20, 0xce, 0xe2, 0xc2, 0x1c, 0x17, 0xb0, 0x80, 0x90, 0x3c, 0x17, 0x77,
	0x41, 0x4e, 0x62, 0x1e, 0x4c, 0x9e, 0x09, 0x2c, 0xcf, 0x05, 0x12, 0x82, 0x2a, 0x30, 0xe4, 0x62,
	0xcd, 0x49, 0x4c, 0x4a, 0xcd, 0x91, 0x60, 0x56, 0x60, 0xd4, 0xe0, 0x33, 0x92, 0xd6, 0x03, 0x39,
	0x40, 0x0f, 0xdd, 0x1a, 0x1f, 0x90, 0x92, 0x20, 0x88, 0x4a, 0xad, 0x72, 0x2e, 0x51, 0xac, 0xf2,
	0x42, 0x0a, 0x5c, 0x32, 0x58, 0x25, 0x42, 0xf3, 0xb2, 0xf3, 0xf2, 0xcb, 0xf3, 0x04, 0x18, 0x84,
	0x64, 0xb8, 0x24, 0xb0, 0x1b, 0x9d, 0x5f, 0x2e, 0xc0, 0x28, 0x24, 0xc7, 0x25, 0x85, 0x55, 0xd6,
	0x33, 0x2f, 0x25, 0xb5, 0x42, 0x80, 0xc9, 0x49, 0xf3, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4,
	0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf1, 0x58, 0x8e, 0x81, 0x4b, 0x34, 0x39, 0x3f, 0x57, 0xaf,
	0x20, 0x33, 0x2f, 0x3d, 0x39, 0xb1, 0x40, 0xaf, 0x24, 0x33, 0x25, 0x09, 0xec, 0xfc, 0x00, 0x46,
	0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4d, 0xfa, 0xe1, 0xd6, 0x5b, 0x01, 0x00, 0x00,
}

func (m *ResourceGroupTag) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResourceGroupTag) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResourceGroupTag) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Label != nil {
		i = encodeVarintResourcetag(dAtA, i, uint64(*m.Label))
		i--
		dAtA[i] = 0x18
	}
	if m.PlanDigest != nil {
		i -= len(m.PlanDigest)
		copy(dAtA[i:], m.PlanDigest)
		i = encodeVarintResourcetag(dAtA, i, uint64(len(m.PlanDigest)))
		i--
		dAtA[i] = 0x12
	}
	if m.SqlDigest != nil {
		i -= len(m.SqlDigest)
		copy(dAtA[i:], m.SqlDigest)
		i = encodeVarintResourcetag(dAtA, i, uint64(len(m.SqlDigest)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintResourcetag(dAtA []byte, offset int, v uint64) int {
	offset -= sovResourcetag(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ResourceGroupTag) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SqlDigest != nil {
		l = len(m.SqlDigest)
		n += 1 + l + sovResourcetag(uint64(l))
	}
	if m.PlanDigest != nil {
		l = len(m.PlanDigest)
		n += 1 + l + sovResourcetag(uint64(l))
	}
	if m.Label != nil {
		n += 1 + sovResourcetag(uint64(*m.Label))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovResourcetag(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozResourcetag(x uint64) (n int) {
	return sovResourcetag(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ResourceGroupTag) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResourcetag
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
			return fmt.Errorf("proto: ResourceGroupTag: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResourceGroupTag: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SqlDigest", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourcetag
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
				return ErrInvalidLengthResourcetag
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthResourcetag
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SqlDigest = append(m.SqlDigest[:0], dAtA[iNdEx:postIndex]...)
			if m.SqlDigest == nil {
				m.SqlDigest = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlanDigest", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourcetag
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
				return ErrInvalidLengthResourcetag
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthResourcetag
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PlanDigest = append(m.PlanDigest[:0], dAtA[iNdEx:postIndex]...)
			if m.PlanDigest == nil {
				m.PlanDigest = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Label", wireType)
			}
			var v ResourceGroupTagLabel
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourcetag
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= ResourceGroupTagLabel(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Label = &v
		default:
			iNdEx = preIndex
			skippy, err := skipResourcetag(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthResourcetag
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
func skipResourcetag(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowResourcetag
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
					return 0, ErrIntOverflowResourcetag
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
					return 0, ErrIntOverflowResourcetag
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
				return 0, ErrInvalidLengthResourcetag
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupResourcetag
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthResourcetag
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthResourcetag        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowResourcetag          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupResourcetag = fmt.Errorf("proto: unexpected end of group")
)
