// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zrchain/policy/policy.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	proto "github.com/cosmos/gogoproto/proto"
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

// Policy defines a policy that can be used to validate a transaction
type Policy struct {
	Creator string     `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id      uint64     `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Name    string     `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Policy  *types.Any `protobuf:"bytes,4,opt,name=policy,proto3" json:"policy,omitempty"`
	Btl     uint64     `protobuf:"varint,5,opt,name=btl,proto3" json:"btl,omitempty"`
}

func (m *Policy) Reset()         { *m = Policy{} }
func (m *Policy) String() string { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()    {}
func (*Policy) Descriptor() ([]byte, []int) {
	return fileDescriptor_53e400df02d23a81, []int{0}
}
func (m *Policy) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Policy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Policy.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Policy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Policy.Merge(m, src)
}
func (m *Policy) XXX_Size() int {
	return m.Size()
}
func (m *Policy) XXX_DiscardUnknown() {
	xxx_messageInfo_Policy.DiscardUnknown(m)
}

var xxx_messageInfo_Policy proto.InternalMessageInfo

func (m *Policy) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Policy) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Policy) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Policy) GetPolicy() *types.Any {
	if m != nil {
		return m.Policy
	}
	return nil
}

func (m *Policy) GetBtl() uint64 {
	if m != nil {
		return m.Btl
	}
	return 0
}

// BoolparserPolicy defines a policy that can be used to validate a transaction
type BoolparserPolicy struct {
	// Definition of the policy, eg.
	// "t1 + t2 + t3 > 1"
	Definition   string               `protobuf:"bytes,1,opt,name=definition,proto3" json:"definition,omitempty"`
	Participants []*PolicyParticipant `protobuf:"bytes,2,rep,name=participants,proto3" json:"participants,omitempty"`
}

func (m *BoolparserPolicy) Reset()         { *m = BoolparserPolicy{} }
func (m *BoolparserPolicy) String() string { return proto.CompactTextString(m) }
func (*BoolparserPolicy) ProtoMessage()    {}
func (*BoolparserPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_53e400df02d23a81, []int{1}
}
func (m *BoolparserPolicy) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BoolparserPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BoolparserPolicy.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BoolparserPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoolparserPolicy.Merge(m, src)
}
func (m *BoolparserPolicy) XXX_Size() int {
	return m.Size()
}
func (m *BoolparserPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_BoolparserPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_BoolparserPolicy proto.InternalMessageInfo

func (m *BoolparserPolicy) GetDefinition() string {
	if m != nil {
		return m.Definition
	}
	return ""
}

func (m *BoolparserPolicy) GetParticipants() []*PolicyParticipant {
	if m != nil {
		return m.Participants
	}
	return nil
}

// PolicyParticipant defines a participant in a policy
type PolicyParticipant struct {
	Abbreviation string `protobuf:"bytes,1,opt,name=abbreviation,proto3" json:"abbreviation,omitempty"`
	Address      string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *PolicyParticipant) Reset()         { *m = PolicyParticipant{} }
func (m *PolicyParticipant) String() string { return proto.CompactTextString(m) }
func (*PolicyParticipant) ProtoMessage()    {}
func (*PolicyParticipant) Descriptor() ([]byte, []int) {
	return fileDescriptor_53e400df02d23a81, []int{2}
}
func (m *PolicyParticipant) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PolicyParticipant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PolicyParticipant.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PolicyParticipant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PolicyParticipant.Merge(m, src)
}
func (m *PolicyParticipant) XXX_Size() int {
	return m.Size()
}
func (m *PolicyParticipant) XXX_DiscardUnknown() {
	xxx_messageInfo_PolicyParticipant.DiscardUnknown(m)
}

var xxx_messageInfo_PolicyParticipant proto.InternalMessageInfo

func (m *PolicyParticipant) GetAbbreviation() string {
	if m != nil {
		return m.Abbreviation
	}
	return ""
}

func (m *PolicyParticipant) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*Policy)(nil), "zrchain.policy.Policy")
	proto.RegisterType((*BoolparserPolicy)(nil), "zrchain.policy.BoolparserPolicy")
	proto.RegisterType((*PolicyParticipant)(nil), "zrchain.policy.PolicyParticipant")
}

func init() { proto.RegisterFile("zrchain/policy/policy.proto", fileDescriptor_53e400df02d23a81) }

var fileDescriptor_53e400df02d23a81 = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x4f, 0x4f, 0xfa, 0x30,
	0x18, 0xc7, 0xe9, 0xe0, 0xc7, 0x2f, 0x14, 0x42, 0xb0, 0xf1, 0x50, 0x35, 0x69, 0xe6, 0x4e, 0x3b,
	0x68, 0x97, 0xa0, 0x89, 0x67, 0x49, 0xf4, 0x0c, 0x3b, 0x72, 0xeb, 0xb6, 0x02, 0x8d, 0xa3, 0x5d,
	0xba, 0x42, 0x9c, 0x2f, 0xc0, 0xb3, 0x2f, 0xcb, 0x23, 0x47, 0x8f, 0x06, 0xde, 0x88, 0xa1, 0xdb,
	0x22, 0x8b, 0xa7, 0x3d, 0x7f, 0x3e, 0x7b, 0xbe, 0x4f, 0xbf, 0x0f, 0xbc, 0x7a, 0xd3, 0xf1, 0x8a,
	0x09, 0x19, 0x64, 0x2a, 0x15, 0x71, 0x51, 0x7d, 0x68, 0xa6, 0x95, 0x51, 0x68, 0x58, 0x35, 0x69,
	0x59, 0xbd, 0xbc, 0x58, 0x2a, 0xb5, 0x4c, 0x79, 0x60, 0xbb, 0xd1, 0x66, 0x11, 0x30, 0x59, 0xa1,
	0xde, 0x3b, 0x80, 0xdd, 0xa9, 0xa5, 0x10, 0x86, 0xff, 0x63, 0xcd, 0x99, 0x51, 0x1a, 0x03, 0x17,
	0xf8, 0xbd, 0xb0, 0x4e, 0xd1, 0x10, 0x3a, 0x22, 0xc1, 0x8e, 0x0b, 0xfc, 0x4e, 0xe8, 0x88, 0x04,
	0x21, 0xd8, 0x91, 0x6c, 0xcd, 0x71, 0xdb, 0x62, 0x36, 0x46, 0x37, 0xb0, 0x5b, 0xaa, 0xe1, 0x8e,
	0x0b, 0xfc, 0xfe, 0xf8, 0x9c, 0x96, 0xa2, 0xb4, 0x16, 0xa5, 0x8f, 0xb2, 0x08, 0x2b, 0x06, 0x8d,
	0x60, 0x3b, 0x32, 0x29, 0xfe, 0x67, 0x47, 0x1e, 0x43, 0xaf, 0x80, 0xa3, 0x89, 0x52, 0x69, 0xc6,
	0x74, 0xce, 0x75, 0xb5, 0x11, 0x81, 0x30, 0xe1, 0x0b, 0x21, 0x85, 0x11, 0x4a, 0x56, 0x4b, 0x9d,
	0x54, 0xd0, 0x13, 0x1c, 0x64, 0x4c, 0x1b, 0x11, 0x8b, 0x8c, 0x49, 0x93, 0x63, 0xc7, 0x6d, 0xfb,
	0xfd, 0xf1, 0x35, 0x6d, 0x3e, 0x9f, 0x96, 0xd3, 0xa6, 0xbf, 0x64, 0xd8, 0xf8, 0xcd, 0x9b, 0xc1,
	0xb3, 0x3f, 0x08, 0xf2, 0xe0, 0x80, 0x45, 0x91, 0xe6, 0x5b, 0xc1, 0x4e, 0xd4, 0x1b, 0xb5, 0xa3,
	0x63, 0x2c, 0x49, 0x34, 0xcf, 0x73, 0x6b, 0x4e, 0x2f, 0xac, 0xd3, 0xc9, 0xec, 0x73, 0x4f, 0xc0,
	0x6e, 0x4f, 0xc0, 0xf7, 0x9e, 0x80, 0x8f, 0x03, 0x69, 0xed, 0x0e, 0xa4, 0xf5, 0x75, 0x20, 0xad,
	0xf9, 0xc3, 0x52, 0x98, 0xd5, 0x26, 0xa2, 0xb1, 0x5a, 0x07, 0x73, 0x2e, 0xb5, 0x8a, 0x5f, 0x6e,
	0x9f, 0xd5, 0x46, 0x26, 0x76, 0x64, 0x50, 0x9f, 0x75, 0x7b, 0x1f, 0xbc, 0xd6, 0xb7, 0x35, 0x45,
	0xc6, 0xf3, 0xa8, 0x6b, 0x8d, 0xbc, 0xfb, 0x09, 0x00, 0x00, 0xff, 0xff, 0xaa, 0x45, 0x26, 0xf4,
	0xfa, 0x01, 0x00, 0x00,
}

func (m *Policy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Policy) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Policy) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Btl != 0 {
		i = encodeVarintPolicy(dAtA, i, uint64(m.Btl))
		i--
		dAtA[i] = 0x28
	}
	if m.Policy != nil {
		{
			size, err := m.Policy.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPolicy(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintPolicy(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Id != 0 {
		i = encodeVarintPolicy(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintPolicy(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BoolparserPolicy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BoolparserPolicy) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BoolparserPolicy) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Participants) > 0 {
		for iNdEx := len(m.Participants) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Participants[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPolicy(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Definition) > 0 {
		i -= len(m.Definition)
		copy(dAtA[i:], m.Definition)
		i = encodeVarintPolicy(dAtA, i, uint64(len(m.Definition)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PolicyParticipant) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PolicyParticipant) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PolicyParticipant) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintPolicy(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Abbreviation) > 0 {
		i -= len(m.Abbreviation)
		copy(dAtA[i:], m.Abbreviation)
		i = encodeVarintPolicy(dAtA, i, uint64(len(m.Abbreviation)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPolicy(dAtA []byte, offset int, v uint64) int {
	offset -= sovPolicy(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Policy) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovPolicy(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovPolicy(uint64(m.Id))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovPolicy(uint64(l))
	}
	if m.Policy != nil {
		l = m.Policy.Size()
		n += 1 + l + sovPolicy(uint64(l))
	}
	if m.Btl != 0 {
		n += 1 + sovPolicy(uint64(m.Btl))
	}
	return n
}

func (m *BoolparserPolicy) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Definition)
	if l > 0 {
		n += 1 + l + sovPolicy(uint64(l))
	}
	if len(m.Participants) > 0 {
		for _, e := range m.Participants {
			l = e.Size()
			n += 1 + l + sovPolicy(uint64(l))
		}
	}
	return n
}

func (m *PolicyParticipant) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Abbreviation)
	if l > 0 {
		n += 1 + l + sovPolicy(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovPolicy(uint64(l))
	}
	return n
}

func sovPolicy(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPolicy(x uint64) (n int) {
	return sovPolicy(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Policy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPolicy
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
			return fmt.Errorf("proto: Policy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Policy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolicy
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
				return ErrInvalidLengthPolicy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPolicy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolicy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolicy
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
				return ErrInvalidLengthPolicy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPolicy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Policy", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolicy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPolicy
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPolicy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Policy == nil {
				m.Policy = &types.Any{}
			}
			if err := m.Policy.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Btl", wireType)
			}
			m.Btl = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolicy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Btl |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPolicy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPolicy
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
func (m *BoolparserPolicy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPolicy
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
			return fmt.Errorf("proto: BoolparserPolicy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BoolparserPolicy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Definition", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolicy
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
				return ErrInvalidLengthPolicy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPolicy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Definition = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Participants", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolicy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPolicy
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPolicy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Participants = append(m.Participants, &PolicyParticipant{})
			if err := m.Participants[len(m.Participants)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPolicy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPolicy
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
func (m *PolicyParticipant) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPolicy
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
			return fmt.Errorf("proto: PolicyParticipant: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PolicyParticipant: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Abbreviation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolicy
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
				return ErrInvalidLengthPolicy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPolicy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Abbreviation = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolicy
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
				return ErrInvalidLengthPolicy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPolicy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPolicy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPolicy
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
func skipPolicy(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPolicy
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
					return 0, ErrIntOverflowPolicy
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
					return 0, ErrIntOverflowPolicy
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
				return 0, ErrInvalidLengthPolicy
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPolicy
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPolicy
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPolicy        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPolicy          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPolicy = fmt.Errorf("proto: unexpected end of group")
)
