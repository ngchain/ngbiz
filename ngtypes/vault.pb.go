// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: vault.proto

package ngtypes

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

type Vault struct {
	NetworkId            int32      `protobuf:"varint,1,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
	Height               uint64     `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	Timestamp            int64      `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	PrevVaultHash        []byte     `protobuf:"bytes,10,opt,name=prev_vault_hash,json=prevVaultHash,proto3" json:"prev_vault_hash,omitempty"`
	Sheet                *Sheet     `protobuf:"bytes,20,opt,name=sheet,proto3" json:"sheet,omitempty"`
	List                 *Account   `protobuf:"bytes,30,opt,name=list,proto3" json:"list,omitempty"`
	Delists              []*Account `protobuf:"bytes,31,rep,name=delists,proto3" json:"delists,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Vault) Reset()         { *m = Vault{} }
func (m *Vault) String() string { return proto.CompactTextString(m) }
func (*Vault) ProtoMessage()    {}
func (*Vault) Descriptor() ([]byte, []int) {
	return fileDescriptor_0adf1cc59b0dff3b, []int{0}
}
func (m *Vault) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Vault) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Vault.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Vault) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vault.Merge(m, src)
}
func (m *Vault) XXX_Size() int {
	return m.Size()
}
func (m *Vault) XXX_DiscardUnknown() {
	xxx_messageInfo_Vault.DiscardUnknown(m)
}

var xxx_messageInfo_Vault proto.InternalMessageInfo

func (m *Vault) GetNetworkId() int32 {
	if m != nil {
		return m.NetworkId
	}
	return 0
}

func (m *Vault) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Vault) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Vault) GetPrevVaultHash() []byte {
	if m != nil {
		return m.PrevVaultHash
	}
	return nil
}

func (m *Vault) GetSheet() *Sheet {
	if m != nil {
		return m.Sheet
	}
	return nil
}

func (m *Vault) GetList() *Account {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *Vault) GetDelists() []*Account {
	if m != nil {
		return m.Delists
	}
	return nil
}

func init() {
	proto.RegisterType((*Vault)(nil), "ngtypes.Vault")
}

func init() { proto.RegisterFile("vault.proto", fileDescriptor_0adf1cc59b0dff3b) }

var fileDescriptor_0adf1cc59b0dff3b = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x4b, 0x2c, 0xcd,
	0x29, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcf, 0x4b, 0x2f, 0xa9, 0x2c, 0x48, 0x2d,
	0x96, 0xe2, 0x2e, 0xce, 0x48, 0x4d, 0x85, 0x8a, 0x4a, 0xf1, 0x26, 0x26, 0x27, 0xe7, 0x97, 0xe6,
	0x41, 0xb9, 0x4a, 0x7f, 0x18, 0xb9, 0x58, 0xc3, 0x40, 0x9a, 0x84, 0x64, 0xb9, 0xb8, 0xf2, 0x52,
	0x4b, 0xca, 0xf3, 0x8b, 0xb2, 0xe3, 0x33, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58, 0x83, 0x38,
	0xa1, 0x22, 0x9e, 0x29, 0x42, 0x62, 0x5c, 0x6c, 0x19, 0xa9, 0x99, 0xe9, 0x19, 0x25, 0x12, 0x4c,
	0x0a, 0x8c, 0x1a, 0x2c, 0x41, 0x50, 0x9e, 0x90, 0x0c, 0x17, 0x67, 0x49, 0x66, 0x6e, 0x6a, 0x71,
	0x49, 0x62, 0x6e, 0x81, 0x04, 0xb3, 0x02, 0xa3, 0x06, 0x73, 0x10, 0x42, 0x40, 0x48, 0x8d, 0x8b,
	0xbf, 0xa0, 0x28, 0xb5, 0x2c, 0x1e, 0xec, 0xae, 0xf8, 0x8c, 0xc4, 0xe2, 0x0c, 0x09, 0x2e, 0x05,
	0x46, 0x0d, 0x9e, 0x20, 0x5e, 0x90, 0x30, 0xd8, 0x62, 0x8f, 0xc4, 0xe2, 0x0c, 0x21, 0x15, 0x2e,
	0x56, 0xb0, 0x23, 0x25, 0x44, 0x14, 0x18, 0x35, 0xb8, 0x8d, 0xf8, 0xf4, 0xa0, 0x6e, 0xd7, 0x0b,
	0x06, 0x89, 0x06, 0x41, 0x24, 0x85, 0x54, 0xb8, 0x58, 0x72, 0x32, 0x8b, 0x4b, 0x24, 0xe4, 0xc0,
	0x8a, 0x04, 0xe0, 0x8a, 0x1c, 0x21, 0x5e, 0x0a, 0x02, 0xcb, 0x0a, 0x69, 0x71, 0xb1, 0xa7, 0xa4,
	0x82, 0x58, 0xc5, 0x12, 0xf2, 0x0a, 0xcc, 0x58, 0x15, 0xc2, 0x14, 0x38, 0x09, 0x9c, 0x78, 0x24,
	0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x33, 0x1e, 0xcb, 0x31, 0x24, 0xb1,
	0x81, 0xc3, 0xc5, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xa5, 0xa0, 0x5a, 0xcf, 0x4b, 0x01, 0x00,
	0x00,
}

func (m *Vault) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Vault) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Vault) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Delists) > 0 {
		for iNdEx := len(m.Delists) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Delists[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintVault(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1
			i--
			dAtA[i] = 0xfa
		}
	}
	if m.List != nil {
		{
			size, err := m.List.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintVault(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xf2
	}
	if m.Sheet != nil {
		{
			size, err := m.Sheet.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintVault(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xa2
	}
	if len(m.PrevVaultHash) > 0 {
		i -= len(m.PrevVaultHash)
		copy(dAtA[i:], m.PrevVaultHash)
		i = encodeVarintVault(dAtA, i, uint64(len(m.PrevVaultHash)))
		i--
		dAtA[i] = 0x52
	}
	if m.Timestamp != 0 {
		i = encodeVarintVault(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x18
	}
	if m.Height != 0 {
		i = encodeVarintVault(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x10
	}
	if m.NetworkId != 0 {
		i = encodeVarintVault(dAtA, i, uint64(m.NetworkId))
		i--
		dAtA[i] = 0x8
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
func (m *Vault) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NetworkId != 0 {
		n += 1 + sovVault(uint64(m.NetworkId))
	}
	if m.Height != 0 {
		n += 1 + sovVault(uint64(m.Height))
	}
	if m.Timestamp != 0 {
		n += 1 + sovVault(uint64(m.Timestamp))
	}
	l = len(m.PrevVaultHash)
	if l > 0 {
		n += 1 + l + sovVault(uint64(l))
	}
	if m.Sheet != nil {
		l = m.Sheet.Size()
		n += 2 + l + sovVault(uint64(l))
	}
	if m.List != nil {
		l = m.List.Size()
		n += 2 + l + sovVault(uint64(l))
	}
	if len(m.Delists) > 0 {
		for _, e := range m.Delists {
			l = e.Size()
			n += 2 + l + sovVault(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovVault(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVault(x uint64) (n int) {
	return sovVault(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Vault) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Vault: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Vault: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NetworkId", wireType)
			}
			m.NetworkId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NetworkId |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrevVaultHash", wireType)
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
			m.PrevVaultHash = append(m.PrevVaultHash[:0], dAtA[iNdEx:postIndex]...)
			if m.PrevVaultHash == nil {
				m.PrevVaultHash = []byte{}
			}
			iNdEx = postIndex
		case 20:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sheet", wireType)
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
			if m.Sheet == nil {
				m.Sheet = &Sheet{}
			}
			if err := m.Sheet.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 30:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field List", wireType)
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
			if m.List == nil {
				m.List = &Account{}
			}
			if err := m.List.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 31:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Delists", wireType)
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
			m.Delists = append(m.Delists, &Account{})
			if err := m.Delists[len(m.Delists)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVault(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthVault
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthVault
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
