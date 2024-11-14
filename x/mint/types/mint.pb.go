// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zrchain/mint/v1beta1/mint.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// Minter represents the minting state.
type Minter struct {
	// current annual inflation rate
	Inflation cosmossdk_io_math.LegacyDec `protobuf:"bytes,1,opt,name=inflation,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"inflation"`
	// current annual expected provisions
	AnnualProvisions cosmossdk_io_math.LegacyDec `protobuf:"bytes,2,opt,name=annual_provisions,json=annualProvisions,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"annual_provisions"`
}

func (m *Minter) Reset()         { *m = Minter{} }
func (m *Minter) String() string { return proto.CompactTextString(m) }
func (*Minter) ProtoMessage()    {}
func (*Minter) Descriptor() ([]byte, []int) {
	return fileDescriptor_008fb58650beebd1, []int{0}
}
func (m *Minter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Minter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Minter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Minter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Minter.Merge(m, src)
}
func (m *Minter) XXX_Size() int {
	return m.Size()
}
func (m *Minter) XXX_DiscardUnknown() {
	xxx_messageInfo_Minter.DiscardUnknown(m)
}

var xxx_messageInfo_Minter proto.InternalMessageInfo

// Params defines the parameters for the x/mint module.
type Params struct {
	// type of coin to mint
	MintDenom string `protobuf:"bytes,1,opt,name=mint_denom,json=mintDenom,proto3" json:"mint_denom,omitempty"`
	// maximum annual change in inflation rate
	InflationRateChange cosmossdk_io_math.LegacyDec `protobuf:"bytes,2,opt,name=inflation_rate_change,json=inflationRateChange,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"inflation_rate_change"`
	// maximum inflation rate
	InflationMax cosmossdk_io_math.LegacyDec `protobuf:"bytes,3,opt,name=inflation_max,json=inflationMax,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"inflation_max"`
	// minimum inflation rate
	InflationMin cosmossdk_io_math.LegacyDec `protobuf:"bytes,4,opt,name=inflation_min,json=inflationMin,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"inflation_min"`
	// goal of percent bonded atoms
	GoalBonded cosmossdk_io_math.LegacyDec `protobuf:"bytes,5,opt,name=goal_bonded,json=goalBonded,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"goal_bonded"`
	// expected blocks per year
	BlocksPerYear uint64 `protobuf:"varint,6,opt,name=blocks_per_year,json=blocksPerYear,proto3" json:"blocks_per_year,omitempty"`
	// goal of percent staking yield
	StakingYield cosmossdk_io_math.LegacyDec `protobuf:"bytes,7,opt,name=staking_yield,json=stakingYield,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"staking_yield"`
	// rewards burn rate
	BurnRate cosmossdk_io_math.LegacyDec `protobuf:"bytes,8,opt,name=burn_rate,json=burnRate,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"burn_rate"`
	// rewards protocol wallet rate
	ProtocolWalletRate cosmossdk_io_math.LegacyDec `protobuf:"bytes,9,opt,name=protocol_wallet_rate,json=protocolWalletRate,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"protocol_wallet_rate"`
	// address of protocol wallet
	ProtocolWalletAddress string `protobuf:"bytes,10,opt,name=protocol_wallet_address,json=protocolWalletAddress,proto3" json:"protocol_wallet_address,omitempty"`
	// reward excess retention rate
	RetentionRate cosmossdk_io_math.LegacyDec `protobuf:"bytes,11,opt,name=retention_rate,json=retentionRate,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"retention_rate"`
	// additional staking rewards rate
	AdditionalStakingRewards cosmossdk_io_math.LegacyDec `protobuf:"bytes,12,opt,name=additional_staking_rewards,json=additionalStakingRewards,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"additional_staking_rewards"`
	// additional MPC rewards rate
	AdditionalMpcRewards cosmossdk_io_math.LegacyDec `protobuf:"bytes,13,opt,name=additional_mpc_rewards,json=additionalMpcRewards,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"additional_mpc_rewards"`
	// additional burn rate
	AdditionalBurnRate cosmossdk_io_math.LegacyDec `protobuf:"bytes,14,opt,name=additional_burn_rate,json=additionalBurnRate,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"additional_burn_rate"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_008fb58650beebd1, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetMintDenom() string {
	if m != nil {
		return m.MintDenom
	}
	return ""
}

func (m *Params) GetBlocksPerYear() uint64 {
	if m != nil {
		return m.BlocksPerYear
	}
	return 0
}

func (m *Params) GetProtocolWalletAddress() string {
	if m != nil {
		return m.ProtocolWalletAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*Minter)(nil), "zrchain.mint.v1beta1.Minter")
	proto.RegisterType((*Params)(nil), "zrchain.mint.v1beta1.Params")
}

func init() { proto.RegisterFile("zrchain/mint/v1beta1/mint.proto", fileDescriptor_008fb58650beebd1) }

var fileDescriptor_008fb58650beebd1 = []byte{
	// 615 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0xc1, 0x4f, 0x13, 0x4f,
	0x14, 0xc7, 0xbb, 0xbf, 0x1f, 0x56, 0x3a, 0x50, 0x94, 0xb1, 0xe8, 0x88, 0xb1, 0x10, 0x0e, 0x86,
	0x60, 0xe8, 0x4a, 0x0c, 0x1c, 0xbc, 0x59, 0x89, 0x27, 0x09, 0xa4, 0x1c, 0x08, 0x18, 0xdd, 0xbc,
	0xce, 0x8e, 0xdb, 0xb1, 0xbb, 0x33, 0xcd, 0xcc, 0x14, 0xa8, 0x37, 0xaf, 0x9e, 0xfc, 0x33, 0x3c,
	0x72, 0xf0, 0xe2, 0x7f, 0xc0, 0x91, 0x78, 0x32, 0x1e, 0x88, 0x81, 0x03, 0xff, 0x86, 0xd9, 0x99,
	0xed, 0x16, 0x39, 0xa9, 0xeb, 0xa5, 0xe9, 0xbc, 0xf7, 0xf6, 0xf3, 0xfd, 0xce, 0xbe, 0x37, 0x3b,
	0x68, 0xee, 0x9d, 0xa2, 0x1d, 0xe0, 0xc2, 0x4f, 0xb8, 0x30, 0xfe, 0xfe, 0x4a, 0x9b, 0x19, 0x58,
	0xb1, 0x8b, 0x46, 0x4f, 0x49, 0x23, 0x71, 0x2d, 0x2b, 0x68, 0xd8, 0x58, 0x56, 0x30, 0x5b, 0x8b,
	0x64, 0x24, 0x6d, 0x81, 0x9f, 0xfe, 0x73, 0xb5, 0xb3, 0x77, 0xa9, 0xd4, 0x89, 0xd4, 0x81, 0x4b,
	0xb8, 0x45, 0x96, 0x9a, 0x86, 0x84, 0x0b, 0xe9, 0xdb, 0x5f, 0x17, 0x5a, 0xf8, 0xe2, 0xa1, 0xf2,
	0x06, 0x17, 0x86, 0x29, 0xbc, 0x89, 0x2a, 0x5c, 0xbc, 0x89, 0xc1, 0x70, 0x29, 0x88, 0x37, 0xef,
	0x2d, 0x56, 0x9a, 0x2b, 0xc7, 0xa7, 0x73, 0xa5, 0xef, 0xa7, 0x73, 0xf7, 0x1c, 0x46, 0x87, 0xdd,
	0x06, 0x97, 0x7e, 0x02, 0xa6, 0xd3, 0x78, 0xc1, 0x22, 0xa0, 0x83, 0x75, 0x46, 0xbf, 0x7e, 0x5e,
	0x46, 0x99, 0xca, 0x3a, 0xa3, 0xad, 0x11, 0x03, 0xbf, 0x46, 0xd3, 0x20, 0x44, 0x1f, 0xe2, 0xd4,
	0xcb, 0x3e, 0xd7, 0x5c, 0x0a, 0x4d, 0xfe, 0xfb, 0x5b, 0xf0, 0x4d, 0xc7, 0xda, 0xca, 0x51, 0x0b,
	0xef, 0x11, 0x2a, 0x6f, 0x81, 0x82, 0x44, 0xe3, 0xfb, 0x08, 0xa5, 0xaf, 0x26, 0x08, 0x99, 0x90,
	0x89, 0x33, 0xdf, 0xaa, 0xa4, 0x91, 0xf5, 0x34, 0x80, 0xdf, 0xa2, 0x99, 0xdc, 0x56, 0xa0, 0xc0,
	0xb0, 0x80, 0x76, 0x40, 0x44, 0x2c, 0x73, 0xb3, 0xf6, 0xc7, 0x6e, 0x3e, 0x5d, 0x1c, 0x2d, 0x79,
	0xad, 0x5b, 0x39, 0xb4, 0x05, 0x86, 0x3d, 0xb3, 0x48, 0xfc, 0x12, 0x55, 0x47, 0x5a, 0x09, 0x1c,
	0x92, 0xff, 0x0b, 0x69, 0x4c, 0xe6, 0xb0, 0x0d, 0x38, 0xbc, 0x02, 0xe7, 0x82, 0x8c, 0xfd, 0x2b,
	0x38, 0x17, 0x78, 0x07, 0x4d, 0x44, 0x12, 0xe2, 0xa0, 0x2d, 0x45, 0xc8, 0x42, 0x72, 0xad, 0x10,
	0x1a, 0xa5, 0xa8, 0xa6, 0x25, 0xe1, 0x07, 0xe8, 0x46, 0x3b, 0x96, 0xb4, 0xab, 0x83, 0x1e, 0x53,
	0xc1, 0x80, 0x81, 0x22, 0xe5, 0x79, 0x6f, 0x71, 0xac, 0x55, 0x75, 0xe1, 0x2d, 0xa6, 0x76, 0x19,
	0xa8, 0x74, 0x77, 0xda, 0x40, 0x97, 0x8b, 0x28, 0x18, 0x70, 0x16, 0x87, 0xe4, 0x7a, 0xb1, 0xdd,
	0x65, 0xb0, 0xdd, 0x94, 0x85, 0xb7, 0x51, 0xa5, 0xdd, 0x57, 0xae, 0xfd, 0x64, 0xbc, 0x10, 0x78,
	0x3c, 0x05, 0xa5, 0x2d, 0xc7, 0x1d, 0x54, 0xb3, 0xe7, 0x88, 0xca, 0x38, 0x38, 0x80, 0x38, 0x66,
	0xc6, 0xf1, 0x2b, 0x85, 0xf8, 0x78, 0xc8, 0xdc, 0xb1, 0x48, 0xab, 0xb4, 0x86, 0xee, 0x5c, 0x55,
	0x82, 0x30, 0x54, 0x4c, 0x6b, 0x82, 0xec, 0xb8, 0xcf, 0xfc, 0xfa, 0xd0, 0x53, 0x97, 0xc4, 0xaf,
	0xd0, 0x94, 0x62, 0x86, 0x89, 0x7c, 0xf4, 0xc9, 0x44, 0x21, 0x6f, 0xd5, 0x9c, 0x66, 0x6d, 0x19,
	0x34, 0x0b, 0x61, 0xc8, 0xd3, 0x35, 0xc4, 0xc1, 0xb0, 0x7b, 0x8a, 0x1d, 0x80, 0x0a, 0x35, 0x99,
	0x2c, 0x24, 0x45, 0x46, 0xe4, 0x6d, 0x07, 0x6e, 0x39, 0x2e, 0x8e, 0xd1, 0xed, 0x4b, 0xaa, 0x49,
	0x8f, 0xe6, 0x8a, 0xd5, 0x42, 0x8a, 0xb5, 0x11, 0x75, 0xa3, 0x47, 0x87, 0x6a, 0x1d, 0x74, 0x29,
	0x1e, 0x8c, 0x86, 0x68, 0xaa, 0x58, 0x93, 0x47, 0xcc, 0x66, 0x36, 0x4e, 0x4f, 0x1e, 0x7d, 0xb8,
	0x38, 0x5a, 0x7a, 0xb8, 0xc7, 0x84, 0x92, 0xb4, 0xbb, 0xfc, 0x5c, 0xf6, 0x45, 0x68, 0x4f, 0xa7,
	0x3f, 0xbc, 0x20, 0xf6, 0x57, 0xfd, 0x43, 0x77, 0x4b, 0xb8, 0x0f, 0x5f, 0x73, 0xf3, 0xf8, 0xac,
	0xee, 0x9d, 0x9c, 0xd5, 0xbd, 0x1f, 0x67, 0x75, 0xef, 0xe3, 0x79, 0xbd, 0x74, 0x72, 0x5e, 0x2f,
	0x7d, 0x3b, 0xaf, 0x97, 0xf6, 0x56, 0x23, 0x6e, 0x3a, 0xfd, 0x76, 0x83, 0xca, 0xc4, 0xff, 0x2d,
	0xa2, 0x19, 0xf4, 0x98, 0x6e, 0x97, 0xed, 0x18, 0x3d, 0xfe, 0x19, 0x00, 0x00, 0xff, 0xff, 0x42,
	0x0e, 0x00, 0x66, 0x94, 0x06, 0x00, 0x00,
}

func (m *Minter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Minter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Minter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.AnnualProvisions.Size()
		i -= size
		if _, err := m.AnnualProvisions.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.Inflation.Size()
		i -= size
		if _, err := m.Inflation.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.AdditionalBurnRate.Size()
		i -= size
		if _, err := m.AdditionalBurnRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x72
	{
		size := m.AdditionalMpcRewards.Size()
		i -= size
		if _, err := m.AdditionalMpcRewards.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x6a
	{
		size := m.AdditionalStakingRewards.Size()
		i -= size
		if _, err := m.AdditionalStakingRewards.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x62
	{
		size := m.RetentionRate.Size()
		i -= size
		if _, err := m.RetentionRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x5a
	if len(m.ProtocolWalletAddress) > 0 {
		i -= len(m.ProtocolWalletAddress)
		copy(dAtA[i:], m.ProtocolWalletAddress)
		i = encodeVarintMint(dAtA, i, uint64(len(m.ProtocolWalletAddress)))
		i--
		dAtA[i] = 0x52
	}
	{
		size := m.ProtocolWalletRate.Size()
		i -= size
		if _, err := m.ProtocolWalletRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	{
		size := m.BurnRate.Size()
		i -= size
		if _, err := m.BurnRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	{
		size := m.StakingYield.Size()
		i -= size
		if _, err := m.StakingYield.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if m.BlocksPerYear != 0 {
		i = encodeVarintMint(dAtA, i, uint64(m.BlocksPerYear))
		i--
		dAtA[i] = 0x30
	}
	{
		size := m.GoalBonded.Size()
		i -= size
		if _, err := m.GoalBonded.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.InflationMin.Size()
		i -= size
		if _, err := m.InflationMin.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.InflationMax.Size()
		i -= size
		if _, err := m.InflationMax.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.InflationRateChange.Size()
		i -= size
		if _, err := m.InflationRateChange.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.MintDenom) > 0 {
		i -= len(m.MintDenom)
		copy(dAtA[i:], m.MintDenom)
		i = encodeVarintMint(dAtA, i, uint64(len(m.MintDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMint(dAtA []byte, offset int, v uint64) int {
	offset -= sovMint(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Minter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Inflation.Size()
	n += 1 + l + sovMint(uint64(l))
	l = m.AnnualProvisions.Size()
	n += 1 + l + sovMint(uint64(l))
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MintDenom)
	if l > 0 {
		n += 1 + l + sovMint(uint64(l))
	}
	l = m.InflationRateChange.Size()
	n += 1 + l + sovMint(uint64(l))
	l = m.InflationMax.Size()
	n += 1 + l + sovMint(uint64(l))
	l = m.InflationMin.Size()
	n += 1 + l + sovMint(uint64(l))
	l = m.GoalBonded.Size()
	n += 1 + l + sovMint(uint64(l))
	if m.BlocksPerYear != 0 {
		n += 1 + sovMint(uint64(m.BlocksPerYear))
	}
	l = m.StakingYield.Size()
	n += 1 + l + sovMint(uint64(l))
	l = m.BurnRate.Size()
	n += 1 + l + sovMint(uint64(l))
	l = m.ProtocolWalletRate.Size()
	n += 1 + l + sovMint(uint64(l))
	l = len(m.ProtocolWalletAddress)
	if l > 0 {
		n += 1 + l + sovMint(uint64(l))
	}
	l = m.RetentionRate.Size()
	n += 1 + l + sovMint(uint64(l))
	l = m.AdditionalStakingRewards.Size()
	n += 1 + l + sovMint(uint64(l))
	l = m.AdditionalMpcRewards.Size()
	n += 1 + l + sovMint(uint64(l))
	l = m.AdditionalBurnRate.Size()
	n += 1 + l + sovMint(uint64(l))
	return n
}

func sovMint(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMint(x uint64) (n int) {
	return sovMint(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Minter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMint
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
			return fmt.Errorf("proto: Minter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Minter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Inflation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Inflation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AnnualProvisions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AnnualProvisions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMint
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMint
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MintDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InflationRateChange", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InflationRateChange.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InflationMax", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InflationMax.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InflationMin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InflationMin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GoalBonded", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.GoalBonded.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlocksPerYear", wireType)
			}
			m.BlocksPerYear = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlocksPerYear |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakingYield", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.StakingYield.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BurnRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BurnRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProtocolWalletRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ProtocolWalletRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProtocolWalletAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProtocolWalletAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RetentionRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RetentionRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdditionalStakingRewards", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AdditionalStakingRewards.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdditionalMpcRewards", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AdditionalMpcRewards.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdditionalBurnRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AdditionalBurnRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMint
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
func skipMint(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMint
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
					return 0, ErrIntOverflowMint
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
					return 0, ErrIntOverflowMint
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
				return 0, ErrInvalidLengthMint
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMint
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMint
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMint        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMint          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMint = fmt.Errorf("proto: unexpected end of group")
)
