// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zrchain/identity/workspace.proto

package types

import (
	fmt "fmt"
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

// defines workspace attributes
type Workspace struct {
	Address         string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Creator         string   `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	Owners          []string `protobuf:"bytes,3,rep,name=owners,proto3" json:"owners,omitempty"`
	ChildWorkspaces []string `protobuf:"bytes,4,rep,name=child_workspaces,json=childWorkspaces,proto3" json:"child_workspaces,omitempty"`
	// Optional ID of the policy to be applied to every *admin* operation.
	// If not specified, the default policy is used.
	//
	// Admin operations are:
	// - zrchain.identity.Msg.AddWorkspaceOwner
	// - zrchain.identity.Msg.RemoveWorkspaceOwner
	// - zrchain.identity.Msg.AppendChildWorkspace
	// - zrchain.identity.Msg.NewChildWorkspace
	//
	// The default policy is to allow any operation when at least one of its
	// owner approves it.
	AdminPolicyId uint64 `protobuf:"varint,5,opt,name=admin_policy_id,json=adminPolicyId,proto3" json:"admin_policy_id,omitempty"`
	// Optional ID of the policy to be applied to every *sign* operation.
	// If not specified, the default policy is used.
	//
	// Sign operations are:
	// - zrchain.treasury.Msg.NewKeyRequest
	// - zrchain.treasury.Msg.NewSignatureRequest
	// - zrchain.treasury.Msg.NewWalletRequest
	// - zrchain.treasury.Msg.NewSignTransactionRequest
	//
	// The default policy is to allow any operation when at least one of its
	// owner approves it.
	SignPolicyId uint64 `protobuf:"varint,6,opt,name=sign_policy_id,json=signPolicyId,proto3" json:"sign_policy_id,omitempty"`
	// Optional alias field to be used in various situations
	// e.g. wallet type for zr sign workspaces
	Alias string `protobuf:"bytes,7,opt,name=alias,proto3" json:"alias,omitempty"`
}

func (m *Workspace) Reset()         { *m = Workspace{} }
func (m *Workspace) String() string { return proto.CompactTextString(m) }
func (*Workspace) ProtoMessage()    {}
func (*Workspace) Descriptor() ([]byte, []int) {
	return fileDescriptor_2784590776081ccd, []int{0}
}
func (m *Workspace) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Workspace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Workspace.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Workspace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Workspace.Merge(m, src)
}
func (m *Workspace) XXX_Size() int {
	return m.Size()
}
func (m *Workspace) XXX_DiscardUnknown() {
	xxx_messageInfo_Workspace.DiscardUnknown(m)
}

var xxx_messageInfo_Workspace proto.InternalMessageInfo

func (m *Workspace) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Workspace) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Workspace) GetOwners() []string {
	if m != nil {
		return m.Owners
	}
	return nil
}

func (m *Workspace) GetChildWorkspaces() []string {
	if m != nil {
		return m.ChildWorkspaces
	}
	return nil
}

func (m *Workspace) GetAdminPolicyId() uint64 {
	if m != nil {
		return m.AdminPolicyId
	}
	return 0
}

func (m *Workspace) GetSignPolicyId() uint64 {
	if m != nil {
		return m.SignPolicyId
	}
	return 0
}

func (m *Workspace) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

func init() {
	proto.RegisterType((*Workspace)(nil), "zrchain.identity.Workspace")
}

func init() { proto.RegisterFile("zrchain/identity/workspace.proto", fileDescriptor_2784590776081ccd) }

var fileDescriptor_2784590776081ccd = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x4f, 0x4b, 0xfb, 0x30,
	0x1c, 0xc6, 0x97, 0xdf, 0xfe, 0xb1, 0xf0, 0xd3, 0x8d, 0x20, 0x92, 0x53, 0x28, 0x22, 0x52, 0x0f,
	0xb6, 0x07, 0x41, 0xf0, 0xea, 0x41, 0xf0, 0x26, 0xf5, 0x20, 0xec, 0x52, 0xb2, 0x24, 0xac, 0x61,
	0x5d, 0x52, 0x92, 0xcc, 0x59, 0x5f, 0x85, 0x2f, 0xcb, 0xe3, 0x8e, 0x1e, 0xa5, 0x05, 0x5f, 0x87,
	0x2c, 0x6b, 0xab, 0xc7, 0xe7, 0xf3, 0xfd, 0xc0, 0xf7, 0xe1, 0x81, 0xc1, 0x9b, 0x61, 0x19, 0x95,
	0x2a, 0x96, 0x5c, 0x28, 0x27, 0x5d, 0x19, 0x6f, 0xb5, 0x59, 0xd9, 0x82, 0x32, 0x11, 0x15, 0x46,
	0x3b, 0x8d, 0x66, 0x8d, 0x11, 0xb5, 0xc6, 0xd9, 0x37, 0x80, 0x93, 0xe7, 0xd6, 0x42, 0x18, 0x8e,
	0x29, 0xe7, 0x46, 0x58, 0x8b, 0x41, 0x00, 0xc2, 0x49, 0xd2, 0xc6, 0xfd, 0x85, 0x19, 0x41, 0x9d,
	0x36, 0xf8, 0xdf, 0xe1, 0xd2, 0x44, 0x74, 0x0a, 0x47, 0x7a, 0xab, 0x84, 0xb1, 0xb8, 0x1f, 0xf4,
	0xc3, 0x49, 0xd2, 0x24, 0x74, 0x09, 0x67, 0x2c, 0x93, 0x39, 0x4f, 0xbb, 0x12, 0x16, 0x0f, 0xbc,
	0x31, 0xf5, 0xbc, 0xfb, 0x6a, 0xd1, 0x05, 0x9c, 0x52, 0xbe, 0x96, 0x2a, 0x2d, 0x74, 0x2e, 0x59,
	0x99, 0x4a, 0x8e, 0x87, 0x01, 0x08, 0x07, 0xc9, 0x91, 0xc7, 0x8f, 0x9e, 0x3e, 0x70, 0x74, 0x0e,
	0x8f, 0xad, 0x5c, 0xfe, 0xd5, 0x46, 0x5e, 0xfb, 0xbf, 0xa7, 0x9d, 0x75, 0x02, 0x87, 0x34, 0x97,
	0xd4, 0xe2, 0xb1, 0x2f, 0x7a, 0x08, 0x77, 0x4f, 0x1f, 0x15, 0x01, 0xbb, 0x8a, 0x80, 0xaf, 0x8a,
	0x80, 0xf7, 0x9a, 0xf4, 0x76, 0x35, 0xe9, 0x7d, 0xd6, 0xa4, 0x37, 0xbf, 0x5d, 0x4a, 0x97, 0x6d,
	0x16, 0x11, 0xd3, 0xeb, 0x78, 0x2e, 0x94, 0xd1, 0x6c, 0x75, 0x75, 0xaf, 0x37, 0x8a, 0x53, 0x27,
	0xb5, 0x8a, 0xdb, 0x51, 0x5f, 0x6e, 0xe2, 0xd7, 0xdf, 0x65, 0x5d, 0x59, 0x08, 0xbb, 0x18, 0xf9,
	0x59, 0xaf, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x61, 0x3d, 0xdf, 0x37, 0x7a, 0x01, 0x00, 0x00,
}

func (m *Workspace) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Workspace) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Workspace) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Alias) > 0 {
		i -= len(m.Alias)
		copy(dAtA[i:], m.Alias)
		i = encodeVarintWorkspace(dAtA, i, uint64(len(m.Alias)))
		i--
		dAtA[i] = 0x3a
	}
	if m.SignPolicyId != 0 {
		i = encodeVarintWorkspace(dAtA, i, uint64(m.SignPolicyId))
		i--
		dAtA[i] = 0x30
	}
	if m.AdminPolicyId != 0 {
		i = encodeVarintWorkspace(dAtA, i, uint64(m.AdminPolicyId))
		i--
		dAtA[i] = 0x28
	}
	if len(m.ChildWorkspaces) > 0 {
		for iNdEx := len(m.ChildWorkspaces) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ChildWorkspaces[iNdEx])
			copy(dAtA[i:], m.ChildWorkspaces[iNdEx])
			i = encodeVarintWorkspace(dAtA, i, uint64(len(m.ChildWorkspaces[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Owners) > 0 {
		for iNdEx := len(m.Owners) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Owners[iNdEx])
			copy(dAtA[i:], m.Owners[iNdEx])
			i = encodeVarintWorkspace(dAtA, i, uint64(len(m.Owners[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintWorkspace(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintWorkspace(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintWorkspace(dAtA []byte, offset int, v uint64) int {
	offset -= sovWorkspace(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Workspace) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovWorkspace(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovWorkspace(uint64(l))
	}
	if len(m.Owners) > 0 {
		for _, s := range m.Owners {
			l = len(s)
			n += 1 + l + sovWorkspace(uint64(l))
		}
	}
	if len(m.ChildWorkspaces) > 0 {
		for _, s := range m.ChildWorkspaces {
			l = len(s)
			n += 1 + l + sovWorkspace(uint64(l))
		}
	}
	if m.AdminPolicyId != 0 {
		n += 1 + sovWorkspace(uint64(m.AdminPolicyId))
	}
	if m.SignPolicyId != 0 {
		n += 1 + sovWorkspace(uint64(m.SignPolicyId))
	}
	l = len(m.Alias)
	if l > 0 {
		n += 1 + l + sovWorkspace(uint64(l))
	}
	return n
}

func sovWorkspace(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozWorkspace(x uint64) (n int) {
	return sovWorkspace(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Workspace) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWorkspace
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
			return fmt.Errorf("proto: Workspace: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Workspace: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkspace
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
				return ErrInvalidLengthWorkspace
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWorkspace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkspace
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
				return ErrInvalidLengthWorkspace
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWorkspace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owners", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkspace
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
				return ErrInvalidLengthWorkspace
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWorkspace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owners = append(m.Owners, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChildWorkspaces", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkspace
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
				return ErrInvalidLengthWorkspace
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWorkspace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChildWorkspaces = append(m.ChildWorkspaces, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdminPolicyId", wireType)
			}
			m.AdminPolicyId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkspace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AdminPolicyId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignPolicyId", wireType)
			}
			m.SignPolicyId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkspace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SignPolicyId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Alias", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkspace
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
				return ErrInvalidLengthWorkspace
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWorkspace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Alias = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWorkspace(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWorkspace
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
func skipWorkspace(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWorkspace
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
					return 0, ErrIntOverflowWorkspace
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
					return 0, ErrIntOverflowWorkspace
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
				return 0, ErrInvalidLengthWorkspace
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupWorkspace
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthWorkspace
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthWorkspace        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWorkspace          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupWorkspace = fmt.Errorf("proto: unexpected end of group")
)
