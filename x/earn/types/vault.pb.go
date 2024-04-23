// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kava/earn/v1beta1/vault.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

// AllowedVault is a vault that is allowed to be created. These can be
// modified via parameter governance.
type AllowedVault struct {
	// Denom is the only supported denomination of the vault for deposits and withdrawals.
	Denom string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	// VaultStrategy is the strategy used for this vault.
	Strategies StrategyTypes `protobuf:"varint,2,rep,packed,name=strategies,proto3,enum=kava.earn.v1beta1.StrategyType,castrepeated=StrategyTypes" json:"strategies,omitempty"`
	// IsPrivateVault is true if the vault only allows depositors contained in
	// AllowedDepositors.
	IsPrivateVault bool `protobuf:"varint,3,opt,name=is_private_vault,json=isPrivateVault,proto3" json:"is_private_vault,omitempty"`
	// AllowedDepositors is a list of addresses that are allowed to deposit to
	// this vault if IsPrivateVault is true. Addresses not contained in this list
	// are not allowed to deposit into this vault. If IsPrivateVault is false,
	// this should be empty and ignored.
	AllowedDepositors []github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,4,rep,name=allowed_depositors,json=allowedDepositors,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"allowed_depositors,omitempty"`
}

func (m *AllowedVault) Reset()         { *m = AllowedVault{} }
func (m *AllowedVault) String() string { return proto.CompactTextString(m) }
func (*AllowedVault) ProtoMessage()    {}
func (*AllowedVault) Descriptor() ([]byte, []int) {
	return fileDescriptor_884eb89509fbdc04, []int{0}
}
func (m *AllowedVault) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AllowedVault) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AllowedVault.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AllowedVault) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllowedVault.Merge(m, src)
}
func (m *AllowedVault) XXX_Size() int {
	return m.Size()
}
func (m *AllowedVault) XXX_DiscardUnknown() {
	xxx_messageInfo_AllowedVault.DiscardUnknown(m)
}

var xxx_messageInfo_AllowedVault proto.InternalMessageInfo

func (m *AllowedVault) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *AllowedVault) GetStrategies() StrategyTypes {
	if m != nil {
		return m.Strategies
	}
	return nil
}

func (m *AllowedVault) GetIsPrivateVault() bool {
	if m != nil {
		return m.IsPrivateVault
	}
	return false
}

func (m *AllowedVault) GetAllowedDepositors() []github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.AllowedDepositors
	}
	return nil
}

// VaultRecord is the state of a vault.
type VaultRecord struct {
	// TotalShares is the total distributed number of shares in the vault.
	TotalShares VaultShare `protobuf:"bytes,1,opt,name=total_shares,json=totalShares,proto3" json:"total_shares"`
}

func (m *VaultRecord) Reset()         { *m = VaultRecord{} }
func (m *VaultRecord) String() string { return proto.CompactTextString(m) }
func (*VaultRecord) ProtoMessage()    {}
func (*VaultRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_884eb89509fbdc04, []int{1}
}
func (m *VaultRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VaultRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VaultRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VaultRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VaultRecord.Merge(m, src)
}
func (m *VaultRecord) XXX_Size() int {
	return m.Size()
}
func (m *VaultRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_VaultRecord.DiscardUnknown(m)
}

var xxx_messageInfo_VaultRecord proto.InternalMessageInfo

func (m *VaultRecord) GetTotalShares() VaultShare {
	if m != nil {
		return m.TotalShares
	}
	return VaultShare{}
}

// VaultShareRecord defines the vault shares owned by a depositor.
type VaultShareRecord struct {
	// Depositor represents the owner of the shares
	Depositor github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=depositor,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"depositor,omitempty"`
	// Shares represent the vault shares owned by the depositor.
	Shares VaultShares `protobuf:"bytes,2,rep,name=shares,proto3,castrepeated=VaultShares" json:"shares"`
}

func (m *VaultShareRecord) Reset()         { *m = VaultShareRecord{} }
func (m *VaultShareRecord) String() string { return proto.CompactTextString(m) }
func (*VaultShareRecord) ProtoMessage()    {}
func (*VaultShareRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_884eb89509fbdc04, []int{2}
}
func (m *VaultShareRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VaultShareRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VaultShareRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VaultShareRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VaultShareRecord.Merge(m, src)
}
func (m *VaultShareRecord) XXX_Size() int {
	return m.Size()
}
func (m *VaultShareRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_VaultShareRecord.DiscardUnknown(m)
}

var xxx_messageInfo_VaultShareRecord proto.InternalMessageInfo

func (m *VaultShareRecord) GetDepositor() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Depositor
	}
	return nil
}

func (m *VaultShareRecord) GetShares() VaultShares {
	if m != nil {
		return m.Shares
	}
	return nil
}

// VaultShare defines shares of a vault owned by a depositor.
type VaultShare struct {
	Denom  string                                 `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	Amount github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"amount"`
}

func (m *VaultShare) Reset()      { *m = VaultShare{} }
func (*VaultShare) ProtoMessage() {}
func (*VaultShare) Descriptor() ([]byte, []int) {
	return fileDescriptor_884eb89509fbdc04, []int{3}
}
func (m *VaultShare) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VaultShare) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VaultShare.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VaultShare) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VaultShare.Merge(m, src)
}
func (m *VaultShare) XXX_Size() int {
	return m.Size()
}
func (m *VaultShare) XXX_DiscardUnknown() {
	xxx_messageInfo_VaultShare.DiscardUnknown(m)
}

var xxx_messageInfo_VaultShare proto.InternalMessageInfo

func (m *VaultShare) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func init() {
	proto.RegisterType((*AllowedVault)(nil), "kava.earn.v1beta1.AllowedVault")
	proto.RegisterType((*VaultRecord)(nil), "kava.earn.v1beta1.VaultRecord")
	proto.RegisterType((*VaultShareRecord)(nil), "kava.earn.v1beta1.VaultShareRecord")
	proto.RegisterType((*VaultShare)(nil), "kava.earn.v1beta1.VaultShare")
}

func init() { proto.RegisterFile("kava/earn/v1beta1/vault.proto", fileDescriptor_884eb89509fbdc04) }

var fileDescriptor_884eb89509fbdc04 = []byte{
	// 492 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x3f, 0x6f, 0xd3, 0x40,
	0x14, 0xb7, 0x93, 0x10, 0xd1, 0x4b, 0xa8, 0x1a, 0xb7, 0x83, 0xa9, 0x54, 0xdb, 0xca, 0x00, 0x5e,
	0x6c, 0xb7, 0x65, 0x43, 0x2c, 0xb1, 0x22, 0x84, 0x98, 0xd0, 0xb5, 0x30, 0xb0, 0x44, 0x17, 0xfb,
	0x70, 0xac, 0x3a, 0xb9, 0xc8, 0xef, 0x92, 0x92, 0x85, 0xcf, 0xc0, 0x88, 0xc4, 0xc2, 0xdc, 0xb9,
	0x9f, 0x01, 0x75, 0xac, 0x3a, 0x21, 0x86, 0x14, 0x25, 0xdf, 0x82, 0x09, 0xdd, 0x1f, 0x35, 0x91,
	0x02, 0x82, 0x81, 0xc9, 0xbe, 0xdf, 0x7b, 0xef, 0xf7, 0xe7, 0xd9, 0x87, 0x0e, 0xce, 0xc8, 0x94,
	0x44, 0x94, 0x94, 0xa3, 0x68, 0x7a, 0xd4, 0xa7, 0x9c, 0x1c, 0x45, 0x53, 0x32, 0x29, 0x78, 0x38,
	0x2e, 0x19, 0x67, 0x56, 0x4b, 0x94, 0x43, 0x51, 0x0e, 0x75, 0x79, 0xff, 0x61, 0xc2, 0x60, 0xc8,
	0xa0, 0x27, 0x1b, 0x22, 0x75, 0x50, 0xdd, 0xfb, 0x7b, 0x19, 0xcb, 0x98, 0xc2, 0xc5, 0x9b, 0x46,
	0xbd, 0x4d, 0x09, 0xe0, 0x25, 0xe1, 0x34, 0x9b, 0xa9, 0x8e, 0xf6, 0xe7, 0x0a, 0x6a, 0x76, 0x8a,
	0x82, 0x9d, 0xd3, 0xf4, 0x8d, 0x10, 0xb7, 0xf6, 0xd0, 0xbd, 0x94, 0x8e, 0xd8, 0xd0, 0x36, 0x3d,
	0xd3, 0xdf, 0xc2, 0xea, 0x60, 0x61, 0x84, 0xf4, 0x60, 0x4e, 0xc1, 0xae, 0x78, 0x55, 0x7f, 0xfb,
	0xd8, 0x0d, 0x37, 0x1c, 0x86, 0x27, 0x9a, 0xfd, 0x74, 0x36, 0xa6, 0x71, 0xeb, 0xe2, 0xd6, 0x7d,
	0xb0, 0x8e, 0x00, 0x5e, 0x63, 0xb1, 0x7c, 0xb4, 0x93, 0x8b, 0x2c, 0xf9, 0x94, 0x70, 0xda, 0x93,
	0xd1, 0xed, 0xaa, 0x67, 0xfa, 0xf7, 0xf1, 0x76, 0x0e, 0xaf, 0x14, 0xac, 0x3c, 0x9d, 0x23, 0x8b,
	0x28, 0x8f, 0xbd, 0x94, 0x8e, 0x19, 0xe4, 0x9c, 0x95, 0x60, 0xd7, 0xbc, 0xaa, 0xdf, 0x8c, 0x5f,
	0xfc, 0x9c, 0xbb, 0x41, 0x96, 0xf3, 0xc1, 0xa4, 0x1f, 0x26, 0x6c, 0xa8, 0xb7, 0xa2, 0x1f, 0x01,
	0xa4, 0x67, 0x11, 0x17, 0xca, 0x61, 0x27, 0x49, 0x3a, 0x69, 0x5a, 0x52, 0x80, 0x9b, 0xcb, 0x60,
	0x57, 0xef, 0x4e, 0x23, 0xf1, 0x8c, 0x53, 0xc0, 0x2d, 0xad, 0xd1, 0xbd, 0x93, 0x68, 0xbf, 0x46,
	0x0d, 0xe9, 0x00, 0xd3, 0x84, 0x95, 0xa9, 0xf5, 0x1c, 0x35, 0x39, 0xe3, 0xa4, 0xe8, 0xc1, 0x80,
	0x94, 0x14, 0xe4, 0x8a, 0x1a, 0xc7, 0x07, 0xbf, 0xd9, 0x83, 0x9c, 0x3a, 0x11, 0x5d, 0x71, 0xed,
	0x6a, 0xee, 0x1a, 0xb8, 0x21, 0x07, 0x25, 0x02, 0xed, 0xaf, 0x26, 0xda, 0x59, 0x75, 0x68, 0xf2,
	0x77, 0x68, 0xeb, 0x2e, 0x9c, 0x64, 0xfe, 0x9f, 0xd9, 0x56, 0xd4, 0xd6, 0x4b, 0x54, 0xd7, 0xf6,
	0xc5, 0x67, 0xfc, 0xab, 0xfd, 0x5d, 0x61, 0xff, 0xe2, 0xd6, 0x6d, 0xac, 0x30, 0xc0, 0x9a, 0xa1,
	0xfd, 0x01, 0xa1, 0x15, 0xfc, 0x87, 0x5f, 0xe7, 0x14, 0xd5, 0xc9, 0x90, 0x4d, 0x46, 0xdc, 0xae,
	0x08, 0x38, 0x7e, 0x26, 0x08, 0xbf, 0xcf, 0xdd, 0x47, 0xff, 0x10, 0xac, 0x4b, 0x93, 0x9b, 0xcb,
	0x00, 0xe9, 0x44, 0x5d, 0x9a, 0x60, 0xcd, 0xf5, 0xb4, 0xf6, 0xe9, 0x8b, 0x6b, 0xc4, 0x9d, 0xab,
	0x85, 0x63, 0x5e, 0x2f, 0x1c, 0xf3, 0xc7, 0xc2, 0x31, 0x3f, 0x2e, 0x1d, 0xe3, 0x7a, 0xe9, 0x18,
	0xdf, 0x96, 0x8e, 0xf1, 0xf6, 0xf1, 0x1a, 0xfb, 0x61, 0x56, 0x90, 0x3e, 0x44, 0x87, 0x59, 0x90,
	0x0c, 0x48, 0x3e, 0x8a, 0xde, 0xab, 0x1b, 0x21, 0x25, 0xfa, 0x75, 0x79, 0x0f, 0x9e, 0xfc, 0x0a,
	0x00, 0x00, 0xff, 0xff, 0x3f, 0x6e, 0x41, 0xab, 0x8e, 0x03, 0x00, 0x00,
}

func (m *AllowedVault) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AllowedVault) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AllowedVault) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AllowedDepositors) > 0 {
		for iNdEx := len(m.AllowedDepositors) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AllowedDepositors[iNdEx])
			copy(dAtA[i:], m.AllowedDepositors[iNdEx])
			i = encodeVarintVault(dAtA, i, uint64(len(m.AllowedDepositors[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if m.IsPrivateVault {
		i--
		if m.IsPrivateVault {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.Strategies) > 0 {
		dAtA2 := make([]byte, len(m.Strategies)*10)
		var j1 int
		for _, num := range m.Strategies {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintVault(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintVault(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *VaultRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VaultRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VaultRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.TotalShares.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintVault(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *VaultShareRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VaultShareRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VaultShareRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Shares) > 0 {
		for iNdEx := len(m.Shares) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Shares[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintVault(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Depositor) > 0 {
		i -= len(m.Depositor)
		copy(dAtA[i:], m.Depositor)
		i = encodeVarintVault(dAtA, i, uint64(len(m.Depositor)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *VaultShare) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VaultShare) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VaultShare) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintVault(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintVault(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintVault(dAtA []byte, offset int, v uint64) int {
	offset -= sovVault(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AllowedVault) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovVault(uint64(l))
	}
	if len(m.Strategies) > 0 {
		l = 0
		for _, e := range m.Strategies {
			l += sovVault(uint64(e))
		}
		n += 1 + sovVault(uint64(l)) + l
	}
	if m.IsPrivateVault {
		n += 2
	}
	if len(m.AllowedDepositors) > 0 {
		for _, b := range m.AllowedDepositors {
			l = len(b)
			n += 1 + l + sovVault(uint64(l))
		}
	}
	return n
}

func (m *VaultRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.TotalShares.Size()
	n += 1 + l + sovVault(uint64(l))
	return n
}

func (m *VaultShareRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Depositor)
	if l > 0 {
		n += 1 + l + sovVault(uint64(l))
	}
	if len(m.Shares) > 0 {
		for _, e := range m.Shares {
			l = e.Size()
			n += 1 + l + sovVault(uint64(l))
		}
	}
	return n
}

func (m *VaultShare) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovVault(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovVault(uint64(l))
	return n
}

func sovVault(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVault(x uint64) (n int) {
	return sovVault(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AllowedVault) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVault
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
			return fmt.Errorf("proto: AllowedVault: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AllowedVault: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVault
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
				return ErrInvalidLengthVault
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType == 0 {
				var v StrategyType
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowVault
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= StrategyType(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Strategies = append(m.Strategies, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowVault
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthVault
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthVault
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				if elementCount != 0 && len(m.Strategies) == 0 {
					m.Strategies = make([]StrategyType, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v StrategyType
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowVault
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= StrategyType(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Strategies = append(m.Strategies, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Strategies", wireType)
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsPrivateVault", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsPrivateVault = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowedDepositors", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVault
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
				return ErrInvalidLengthVault
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthVault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AllowedDepositors = append(m.AllowedDepositors, make([]byte, postIndex-iNdEx))
			copy(m.AllowedDepositors[len(m.AllowedDepositors)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVault(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVault
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
func (m *VaultRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVault
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
			return fmt.Errorf("proto: VaultRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VaultRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalShares", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVault
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
				return ErrInvalidLengthVault
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalShares.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVault(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVault
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
func (m *VaultShareRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVault
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
			return fmt.Errorf("proto: VaultShareRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VaultShareRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Depositor", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVault
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
				return ErrInvalidLengthVault
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthVault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Depositor = append(m.Depositor[:0], dAtA[iNdEx:postIndex]...)
			if m.Depositor == nil {
				m.Depositor = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Shares", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVault
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
				return ErrInvalidLengthVault
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Shares = append(m.Shares, VaultShare{})
			if err := m.Shares[len(m.Shares)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVault(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVault
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
func (m *VaultShare) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVault
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
			return fmt.Errorf("proto: VaultShare: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VaultShare: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVault
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
				return ErrInvalidLengthVault
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVault
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
				return ErrInvalidLengthVault
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVault(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVault
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
func skipVault(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVault
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
					return 0, ErrIntOverflowVault
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
					return 0, ErrIntOverflowVault
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
				return 0, ErrInvalidLengthVault
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVault
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVault
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVault        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVault          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVault = fmt.Errorf("proto: unexpected end of group")
)
