// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: transaction.proto

package ngtypes

import (
	bytes "bytes"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

type TxHeader struct {
	Version      int32    `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Type         int32    `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	Convener     uint64   `protobuf:"varint,10,opt,name=convener,proto3" json:"convener,omitempty"`
	Participants [][]byte `protobuf:"bytes,11,rep,name=participants,proto3" json:"participants,omitempty"`
	Fee          []byte   `protobuf:"bytes,12,opt,name=fee,proto3" json:"fee,omitempty"`
	Values       [][]byte `protobuf:"bytes,13,rep,name=values,proto3" json:"values,omitempty"`
	Nonce        uint64   `protobuf:"varint,14,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// extension
	Extra []byte `protobuf:"bytes,20,opt,name=extra,proto3" json:"extra,omitempty"`
}

func (m *TxHeader) Reset()      { *m = TxHeader{} }
func (*TxHeader) ProtoMessage() {}
func (*TxHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{0}
}
func (m *TxHeader) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TxHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TxHeader.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TxHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxHeader.Merge(m, src)
}
func (m *TxHeader) XXX_Size() int {
	return m.Size()
}
func (m *TxHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_TxHeader.DiscardUnknown(m)
}

var xxx_messageInfo_TxHeader proto.InternalMessageInfo

func (m *TxHeader) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *TxHeader) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *TxHeader) GetConvener() uint64 {
	if m != nil {
		return m.Convener
	}
	return 0
}

func (m *TxHeader) GetParticipants() [][]byte {
	if m != nil {
		return m.Participants
	}
	return nil
}

func (m *TxHeader) GetFee() []byte {
	if m != nil {
		return m.Fee
	}
	return nil
}

func (m *TxHeader) GetValues() [][]byte {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *TxHeader) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *TxHeader) GetExtra() []byte {
	if m != nil {
		return m.Extra
	}
	return nil
}

type Transaction struct {
	Header *TxHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// sign
	HeaderHash []byte `protobuf:"bytes,2,opt,name=header_hash,json=headerHash,proto3" json:"header_hash,omitempty"`
	R          []byte `protobuf:"bytes,10,opt,name=r,proto3" json:"r,omitempty"`
	S          []byte `protobuf:"bytes,11,opt,name=s,proto3" json:"s,omitempty"`
}

func (m *Transaction) Reset()      { *m = Transaction{} }
func (*Transaction) ProtoMessage() {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{1}
}
func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(m, src)
}
func (m *Transaction) XXX_Size() int {
	return m.Size()
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetHeader() *TxHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Transaction) GetHeaderHash() []byte {
	if m != nil {
		return m.HeaderHash
	}
	return nil
}

func (m *Transaction) GetR() []byte {
	if m != nil {
		return m.R
	}
	return nil
}

func (m *Transaction) GetS() []byte {
	if m != nil {
		return m.S
	}
	return nil
}

func init() {
	proto.RegisterType((*TxHeader)(nil), "ngtypes.TxHeader")
	proto.RegisterType((*Transaction)(nil), "ngtypes.Transaction")
}

func init() { proto.RegisterFile("transaction.proto", fileDescriptor_2cc4e03d2c28c490) }

var fileDescriptor_2cc4e03d2c28c490 = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x85, 0x73, 0xe9, 0xaf, 0x6e, 0x0c, 0xa2, 0x56, 0x85, 0x2c, 0x06, 0x13, 0x75, 0x0a, 0x4b,
	0x07, 0xe0, 0x09, 0x98, 0x3a, 0x5b, 0xdd, 0x91, 0x09, 0x86, 0x44, 0x42, 0x4e, 0x64, 0xbb, 0x55,
	0xd9, 0x78, 0x04, 0x1e, 0x83, 0x47, 0x61, 0x42, 0x1d, 0x3b, 0x52, 0x77, 0x61, 0xec, 0x23, 0xa0,
	0x38, 0x29, 0x88, 0xed, 0x7c, 0x47, 0x47, 0xb6, 0x3e, 0x5d, 0x1c, 0x39, 0x23, 0xb5, 0x95, 0x99,
	0x2b, 0x4a, 0x3d, 0xad, 0x4c, 0xe9, 0x4a, 0x3a, 0xd0, 0x4f, 0xee, 0xa5, 0x52, 0x76, 0xf2, 0x09,
	0x38, 0x9c, 0xaf, 0x66, 0x4a, 0x3e, 0x28, 0x43, 0x19, 0x0e, 0x96, 0xca, 0xd8, 0xa2, 0xd4, 0x0c,
	0x12, 0x48, 0x7b, 0xe2, 0x80, 0x94, 0x62, 0xb7, 0xde, 0xb3, 0xa3, 0x50, 0x87, 0x4c, 0xcf, 0x71,
	0x98, 0x95, 0x7a, 0xa9, 0xb4, 0x32, 0x0c, 0x13, 0x48, 0xbb, 0xe2, 0x97, 0xe9, 0x04, 0x49, 0x25,
	0x8d, 0x2b, 0xb2, 0xa2, 0x92, 0xda, 0x59, 0x16, 0x27, 0x9d, 0x94, 0x88, 0x7f, 0x1d, 0x3d, 0xc5,
	0xce, 0xa3, 0x52, 0x8c, 0x24, 0x90, 0x12, 0x51, 0x47, 0x7a, 0x86, 0xfd, 0xa5, 0x7c, 0x5e, 0x28,
	0xcb, 0x8e, 0xc3, 0xbe, 0x25, 0x3a, 0xc6, 0x9e, 0x2e, 0x75, 0xa6, 0xd8, 0x49, 0xf8, 0xa6, 0x81,
	0xba, 0x55, 0x2b, 0x67, 0x24, 0x1b, 0x87, 0x17, 0x1a, 0x98, 0x2c, 0x30, 0x9e, 0xff, 0xe9, 0xd2,
	0x4b, 0xec, 0xe7, 0x41, 0x2e, 0x18, 0xc5, 0x57, 0xa3, 0x69, 0x6b, 0x3e, 0x3d, 0x58, 0x8b, 0x76,
	0x40, 0x2f, 0x30, 0x6e, 0xd2, 0x5d, 0x2e, 0x6d, 0x1e, 0x54, 0x89, 0xc0, 0xa6, 0x9a, 0x49, 0x9b,
	0x53, 0x82, 0xd0, 0x98, 0x12, 0x01, 0xa6, 0xa6, 0xda, 0x2b, 0x90, 0xbd, 0xbd, 0x59, 0x6f, 0x79,
	0xb4, 0xd9, 0xf2, 0x68, 0xbf, 0xe5, 0xf0, 0xea, 0x39, 0xbc, 0x7b, 0x0e, 0x1f, 0x9e, 0xc3, 0xda,
	0x73, 0xf8, 0xf2, 0x1c, 0xbe, 0x3d, 0x8f, 0xf6, 0x9e, 0xc3, 0xdb, 0x8e, 0x47, 0xeb, 0x1d, 0x8f,
	0x36, 0x3b, 0x1e, 0xdd, 0xf7, 0xc3, 0x35, 0xae, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xbe, 0xb5,
	0x88, 0xf2, 0xa2, 0x01, 0x00, 0x00,
}

func (this *TxHeader) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TxHeader)
	if !ok {
		that2, ok := that.(TxHeader)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Version != that1.Version {
		return false
	}
	if this.Type != that1.Type {
		return false
	}
	if this.Convener != that1.Convener {
		return false
	}
	if len(this.Participants) != len(that1.Participants) {
		return false
	}
	for i := range this.Participants {
		if !bytes.Equal(this.Participants[i], that1.Participants[i]) {
			return false
		}
	}
	if !bytes.Equal(this.Fee, that1.Fee) {
		return false
	}
	if len(this.Values) != len(that1.Values) {
		return false
	}
	for i := range this.Values {
		if !bytes.Equal(this.Values[i], that1.Values[i]) {
			return false
		}
	}
	if this.Nonce != that1.Nonce {
		return false
	}
	if !bytes.Equal(this.Extra, that1.Extra) {
		return false
	}
	return true
}
func (this *Transaction) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Transaction)
	if !ok {
		that2, ok := that.(Transaction)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Header.Equal(that1.Header) {
		return false
	}
	if !bytes.Equal(this.HeaderHash, that1.HeaderHash) {
		return false
	}
	if !bytes.Equal(this.R, that1.R) {
		return false
	}
	if !bytes.Equal(this.S, that1.S) {
		return false
	}
	return true
}
func (this *TxHeader) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 12)
	s = append(s, "&ngtypes.TxHeader{")
	s = append(s, "Version: "+fmt.Sprintf("%#v", this.Version)+",\n")
	s = append(s, "Type: "+fmt.Sprintf("%#v", this.Type)+",\n")
	s = append(s, "Convener: "+fmt.Sprintf("%#v", this.Convener)+",\n")
	s = append(s, "Participants: "+fmt.Sprintf("%#v", this.Participants)+",\n")
	s = append(s, "Fee: "+fmt.Sprintf("%#v", this.Fee)+",\n")
	s = append(s, "Values: "+fmt.Sprintf("%#v", this.Values)+",\n")
	s = append(s, "Nonce: "+fmt.Sprintf("%#v", this.Nonce)+",\n")
	s = append(s, "Extra: "+fmt.Sprintf("%#v", this.Extra)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Transaction) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&ngtypes.Transaction{")
	if this.Header != nil {
		s = append(s, "Header: "+fmt.Sprintf("%#v", this.Header)+",\n")
	}
	s = append(s, "HeaderHash: "+fmt.Sprintf("%#v", this.HeaderHash)+",\n")
	s = append(s, "R: "+fmt.Sprintf("%#v", this.R)+",\n")
	s = append(s, "S: "+fmt.Sprintf("%#v", this.S)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringTransaction(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *TxHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TxHeader) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TxHeader) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Extra) > 0 {
		i -= len(m.Extra)
		copy(dAtA[i:], m.Extra)
		i = encodeVarintTransaction(dAtA, i, uint64(len(m.Extra)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xa2
	}
	if m.Nonce != 0 {
		i = encodeVarintTransaction(dAtA, i, uint64(m.Nonce))
		i--
		dAtA[i] = 0x70
	}
	if len(m.Values) > 0 {
		for iNdEx := len(m.Values) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Values[iNdEx])
			copy(dAtA[i:], m.Values[iNdEx])
			i = encodeVarintTransaction(dAtA, i, uint64(len(m.Values[iNdEx])))
			i--
			dAtA[i] = 0x6a
		}
	}
	if len(m.Fee) > 0 {
		i -= len(m.Fee)
		copy(dAtA[i:], m.Fee)
		i = encodeVarintTransaction(dAtA, i, uint64(len(m.Fee)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.Participants) > 0 {
		for iNdEx := len(m.Participants) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Participants[iNdEx])
			copy(dAtA[i:], m.Participants[iNdEx])
			i = encodeVarintTransaction(dAtA, i, uint64(len(m.Participants[iNdEx])))
			i--
			dAtA[i] = 0x5a
		}
	}
	if m.Convener != 0 {
		i = encodeVarintTransaction(dAtA, i, uint64(m.Convener))
		i--
		dAtA[i] = 0x50
	}
	if m.Type != 0 {
		i = encodeVarintTransaction(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x10
	}
	if m.Version != 0 {
		i = encodeVarintTransaction(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Transaction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Transaction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Transaction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.S) > 0 {
		i -= len(m.S)
		copy(dAtA[i:], m.S)
		i = encodeVarintTransaction(dAtA, i, uint64(len(m.S)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.R) > 0 {
		i -= len(m.R)
		copy(dAtA[i:], m.R)
		i = encodeVarintTransaction(dAtA, i, uint64(len(m.R)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.HeaderHash) > 0 {
		i -= len(m.HeaderHash)
		copy(dAtA[i:], m.HeaderHash)
		i = encodeVarintTransaction(dAtA, i, uint64(len(m.HeaderHash)))
		i--
		dAtA[i] = 0x12
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTransaction(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTransaction(dAtA []byte, offset int, v uint64) int {
	offset -= sovTransaction(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TxHeader) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Version != 0 {
		n += 1 + sovTransaction(uint64(m.Version))
	}
	if m.Type != 0 {
		n += 1 + sovTransaction(uint64(m.Type))
	}
	if m.Convener != 0 {
		n += 1 + sovTransaction(uint64(m.Convener))
	}
	if len(m.Participants) > 0 {
		for _, b := range m.Participants {
			l = len(b)
			n += 1 + l + sovTransaction(uint64(l))
		}
	}
	l = len(m.Fee)
	if l > 0 {
		n += 1 + l + sovTransaction(uint64(l))
	}
	if len(m.Values) > 0 {
		for _, b := range m.Values {
			l = len(b)
			n += 1 + l + sovTransaction(uint64(l))
		}
	}
	if m.Nonce != 0 {
		n += 1 + sovTransaction(uint64(m.Nonce))
	}
	l = len(m.Extra)
	if l > 0 {
		n += 2 + l + sovTransaction(uint64(l))
	}
	return n
}

func (m *Transaction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovTransaction(uint64(l))
	}
	l = len(m.HeaderHash)
	if l > 0 {
		n += 1 + l + sovTransaction(uint64(l))
	}
	l = len(m.R)
	if l > 0 {
		n += 1 + l + sovTransaction(uint64(l))
	}
	l = len(m.S)
	if l > 0 {
		n += 1 + l + sovTransaction(uint64(l))
	}
	return n
}

func sovTransaction(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTransaction(x uint64) (n int) {
	return sovTransaction(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *TxHeader) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&TxHeader{`,
		`Version:` + fmt.Sprintf("%v", this.Version) + `,`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`Convener:` + fmt.Sprintf("%v", this.Convener) + `,`,
		`Participants:` + fmt.Sprintf("%v", this.Participants) + `,`,
		`Fee:` + fmt.Sprintf("%v", this.Fee) + `,`,
		`Values:` + fmt.Sprintf("%v", this.Values) + `,`,
		`Nonce:` + fmt.Sprintf("%v", this.Nonce) + `,`,
		`Extra:` + fmt.Sprintf("%v", this.Extra) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Transaction) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Transaction{`,
		`Header:` + strings.Replace(this.Header.String(), "TxHeader", "TxHeader", 1) + `,`,
		`HeaderHash:` + fmt.Sprintf("%v", this.HeaderHash) + `,`,
		`R:` + fmt.Sprintf("%v", this.R) + `,`,
		`S:` + fmt.Sprintf("%v", this.S) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringTransaction(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *TxHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTransaction
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
			return fmt.Errorf("proto: TxHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TxHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Convener", wireType)
			}
			m.Convener = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Convener |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Participants", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
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
				return ErrInvalidLengthTransaction
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTransaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Participants = append(m.Participants, make([]byte, postIndex-iNdEx))
			copy(m.Participants[len(m.Participants)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
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
				return ErrInvalidLengthTransaction
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTransaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Fee = append(m.Fee[:0], dAtA[iNdEx:postIndex]...)
			if m.Fee == nil {
				m.Fee = []byte{}
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Values", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
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
				return ErrInvalidLengthTransaction
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTransaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Values = append(m.Values, make([]byte, postIndex-iNdEx))
			copy(m.Values[len(m.Values)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 20:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Extra", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
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
				return ErrInvalidLengthTransaction
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTransaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Extra = append(m.Extra[:0], dAtA[iNdEx:postIndex]...)
			if m.Extra == nil {
				m.Extra = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTransaction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTransaction
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTransaction
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
func (m *Transaction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTransaction
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
			return fmt.Errorf("proto: Transaction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Transaction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
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
				return ErrInvalidLengthTransaction
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTransaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &TxHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HeaderHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
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
				return ErrInvalidLengthTransaction
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTransaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HeaderHash = append(m.HeaderHash[:0], dAtA[iNdEx:postIndex]...)
			if m.HeaderHash == nil {
				m.HeaderHash = []byte{}
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field R", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
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
				return ErrInvalidLengthTransaction
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTransaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.R = append(m.R[:0], dAtA[iNdEx:postIndex]...)
			if m.R == nil {
				m.R = []byte{}
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field S", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
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
				return ErrInvalidLengthTransaction
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTransaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.S = append(m.S[:0], dAtA[iNdEx:postIndex]...)
			if m.S == nil {
				m.S = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTransaction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTransaction
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTransaction
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
func skipTransaction(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTransaction
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
					return 0, ErrIntOverflowTransaction
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
					return 0, ErrIntOverflowTransaction
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
				return 0, ErrInvalidLengthTransaction
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTransaction
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTransaction
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTransaction        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTransaction          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTransaction = fmt.Errorf("proto: unexpected end of group")
)
