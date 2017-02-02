// Code generated by protoc-gen-gogo.
// source: types.proto
// DO NOT EDIT!

/*
	Package types is a generated protocol buffer package.

	It is generated from these files:
		types.proto

	It has these top-level messages:
		Posting
		PostingList
*/
package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Posting_ValType int32

const (
	Posting_STRING   Posting_ValType = 0
	Posting_BINARY   Posting_ValType = 1
	Posting_INT32    Posting_ValType = 2
	Posting_FLOAT    Posting_ValType = 3
	Posting_BOOL     Posting_ValType = 4
	Posting_DATE     Posting_ValType = 5
	Posting_DATETIME Posting_ValType = 6
	Posting_GEO      Posting_ValType = 7
	Posting_UID      Posting_ValType = 8
)

var Posting_ValType_name = map[int32]string{
	0: "STRING",
	1: "BINARY",
	2: "INT32",
	3: "FLOAT",
	4: "BOOL",
	5: "DATE",
	6: "DATETIME",
	7: "GEO",
	8: "UID",
}
var Posting_ValType_value = map[string]int32{
	"STRING":   0,
	"BINARY":   1,
	"INT32":    2,
	"FLOAT":    3,
	"BOOL":     4,
	"DATE":     5,
	"DATETIME": 6,
	"GEO":      7,
	"UID":      8,
}

func (x Posting_ValType) String() string {
	return proto.EnumName(Posting_ValType_name, int32(x))
}
func (Posting_ValType) EnumDescriptor() ([]byte, []int) { return fileDescriptorTypes, []int{0, 0} }

type Posting struct {
	Uid     uint64          `protobuf:"fixed64,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Value   []byte          `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	ValType Posting_ValType `protobuf:"varint,3,opt,name=val_type,json=valType,proto3,enum=types.Posting_ValType" json:"val_type,omitempty"`
	Lang    string          `protobuf:"bytes,4,opt,name=lang,proto3" json:"lang,omitempty"`
	Label   string          `protobuf:"bytes,5,opt,name=label,proto3" json:"label,omitempty"`
	Commit  uint64          `protobuf:"varint,6,opt,name=commit,proto3" json:"commit,omitempty"`
	// TODO: op is only used temporarily. See if we can remove it from here.
	Op uint32 `protobuf:"varint,12,opt,name=op,proto3" json:"op,omitempty"`
}

func (m *Posting) Reset()                    { *m = Posting{} }
func (m *Posting) String() string            { return proto.CompactTextString(m) }
func (*Posting) ProtoMessage()               {}
func (*Posting) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{0} }

func (m *Posting) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *Posting) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Posting) GetValType() Posting_ValType {
	if m != nil {
		return m.ValType
	}
	return Posting_STRING
}

func (m *Posting) GetLang() string {
	if m != nil {
		return m.Lang
	}
	return ""
}

func (m *Posting) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *Posting) GetCommit() uint64 {
	if m != nil {
		return m.Commit
	}
	return 0
}

func (m *Posting) GetOp() uint32 {
	if m != nil {
		return m.Op
	}
	return 0
}

type PostingList struct {
	Postings []*Posting `protobuf:"bytes,1,rep,name=postings" json:"postings,omitempty"`
	Checksum []byte     `protobuf:"bytes,2,opt,name=checksum,proto3" json:"checksum,omitempty"`
	Commit   uint64     `protobuf:"varint,3,opt,name=commit,proto3" json:"commit,omitempty"`
	Langs    []string   `protobuf:"bytes,4,rep,name=langs" json:"langs,omitempty"`
}

func (m *PostingList) Reset()                    { *m = PostingList{} }
func (m *PostingList) String() string            { return proto.CompactTextString(m) }
func (*PostingList) ProtoMessage()               {}
func (*PostingList) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{1} }

func (m *PostingList) GetPostings() []*Posting {
	if m != nil {
		return m.Postings
	}
	return nil
}

func (m *PostingList) GetChecksum() []byte {
	if m != nil {
		return m.Checksum
	}
	return nil
}

func (m *PostingList) GetCommit() uint64 {
	if m != nil {
		return m.Commit
	}
	return 0
}

func (m *PostingList) GetLangs() []string {
	if m != nil {
		return m.Langs
	}
	return nil
}

func init() {
	proto.RegisterType((*Posting)(nil), "types.Posting")
	proto.RegisterType((*PostingList)(nil), "types.PostingList")
	proto.RegisterEnum("types.Posting_ValType", Posting_ValType_name, Posting_ValType_value)
}
func (m *Posting) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Posting) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Uid != 0 {
		dAtA[i] = 0x9
		i++
		i = encodeFixed64Types(dAtA, i, uint64(m.Uid))
	}
	if len(m.Value) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Value)))
		i += copy(dAtA[i:], m.Value)
	}
	if m.ValType != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintTypes(dAtA, i, uint64(m.ValType))
	}
	if len(m.Lang) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Lang)))
		i += copy(dAtA[i:], m.Lang)
	}
	if len(m.Label) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Label)))
		i += copy(dAtA[i:], m.Label)
	}
	if m.Commit != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintTypes(dAtA, i, uint64(m.Commit))
	}
	if m.Op != 0 {
		dAtA[i] = 0x60
		i++
		i = encodeVarintTypes(dAtA, i, uint64(m.Op))
	}
	return i, nil
}

func (m *PostingList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PostingList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Postings) > 0 {
		for _, msg := range m.Postings {
			dAtA[i] = 0xa
			i++
			i = encodeVarintTypes(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Checksum) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Checksum)))
		i += copy(dAtA[i:], m.Checksum)
	}
	if m.Commit != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintTypes(dAtA, i, uint64(m.Commit))
	}
	if len(m.Langs) > 0 {
		for _, s := range m.Langs {
			dAtA[i] = 0x22
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func encodeFixed64Types(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Types(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Posting) Size() (n int) {
	var l int
	_ = l
	if m.Uid != 0 {
		n += 9
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.ValType != 0 {
		n += 1 + sovTypes(uint64(m.ValType))
	}
	l = len(m.Lang)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Label)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Commit != 0 {
		n += 1 + sovTypes(uint64(m.Commit))
	}
	if m.Op != 0 {
		n += 1 + sovTypes(uint64(m.Op))
	}
	return n
}

func (m *PostingList) Size() (n int) {
	var l int
	_ = l
	if len(m.Postings) > 0 {
		for _, e := range m.Postings {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	l = len(m.Checksum)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Commit != 0 {
		n += 1 + sovTypes(uint64(m.Commit))
	}
	if len(m.Langs) > 0 {
		for _, s := range m.Langs {
			l = len(s)
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func sovTypes(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Posting) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Posting: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Posting: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			m.Uid = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			m.Uid = uint64(dAtA[iNdEx-8])
			m.Uid |= uint64(dAtA[iNdEx-7]) << 8
			m.Uid |= uint64(dAtA[iNdEx-6]) << 16
			m.Uid |= uint64(dAtA[iNdEx-5]) << 24
			m.Uid |= uint64(dAtA[iNdEx-4]) << 32
			m.Uid |= uint64(dAtA[iNdEx-3]) << 40
			m.Uid |= uint64(dAtA[iNdEx-2]) << 48
			m.Uid |= uint64(dAtA[iNdEx-1]) << 56
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValType", wireType)
			}
			m.ValType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ValType |= (Posting_ValType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lang", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Lang = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Label", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Label = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Commit", wireType)
			}
			m.Commit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Commit |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Op", wireType)
			}
			m.Op = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Op |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
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
func (m *PostingList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: PostingList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PostingList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Postings", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Postings = append(m.Postings, &Posting{})
			if err := m.Postings[len(m.Postings)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Checksum", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Checksum = append(m.Checksum[:0], dAtA[iNdEx:postIndex]...)
			if m.Checksum == nil {
				m.Checksum = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Commit", wireType)
			}
			m.Commit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Commit |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Langs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Langs = append(m.Langs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTypes
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
				next, err := skipTypes(dAtA[start:])
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
	ErrInvalidLengthTypes = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("types.proto", fileDescriptorTypes) }

var fileDescriptorTypes = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x91, 0xcf, 0x6a, 0xea, 0x40,
	0x18, 0xc5, 0x9d, 0xfc, 0xf7, 0xd3, 0x2b, 0xc3, 0xc7, 0x45, 0x86, 0xbb, 0x08, 0x83, 0xab, 0xe1,
	0x2e, 0x84, 0xab, 0x4f, 0xa0, 0xe8, 0x95, 0x80, 0x35, 0x65, 0x9a, 0x16, 0xba, 0x2a, 0xd1, 0x06,
	0x1b, 0x8c, 0x26, 0x34, 0x51, 0x70, 0xdd, 0x97, 0xe8, 0xbe, 0x2f, 0xd3, 0x65, 0x1f, 0xa1, 0xd8,
	0x17, 0x29, 0x93, 0x04, 0xa9, 0xbb, 0xdf, 0x6f, 0x06, 0xce, 0x70, 0xce, 0x40, 0xab, 0x38, 0x66,
	0x51, 0xde, 0xcf, 0x9e, 0xd3, 0x22, 0x45, 0xb3, 0x94, 0xde, 0x9b, 0x06, 0xf6, 0x75, 0x9a, 0x17,
	0xf1, 0x6e, 0x8d, 0x14, 0xf4, 0x7d, 0xfc, 0xc8, 0x08, 0x27, 0xc2, 0x92, 0x0a, 0xf1, 0x37, 0x98,
	0x87, 0x30, 0xd9, 0x47, 0x4c, 0xe3, 0x44, 0xb4, 0x65, 0x25, 0xf8, 0x0f, 0x9c, 0x43, 0x98, 0x3c,
	0xa8, 0x00, 0xa6, 0x73, 0x22, 0x3a, 0x83, 0x6e, 0xbf, 0x8a, 0xae, 0x93, 0xfa, 0x77, 0x61, 0x12,
	0x1c, 0xb3, 0x48, 0xda, 0x87, 0x0a, 0x10, 0xc1, 0x48, 0xc2, 0xdd, 0x9a, 0x19, 0x9c, 0x88, 0xa6,
	0x2c, 0x59, 0x85, 0x27, 0xe1, 0x32, 0x4a, 0x98, 0x59, 0x1e, 0x56, 0x82, 0x5d, 0xb0, 0x56, 0xe9,
	0x76, 0x1b, 0x17, 0xcc, 0xe2, 0x44, 0x18, 0xb2, 0x36, 0xec, 0x80, 0x96, 0x66, 0xac, 0xcd, 0x89,
	0xf8, 0x25, 0xb5, 0x34, 0xeb, 0x6d, 0xc0, 0xae, 0x5f, 0x41, 0x00, 0xeb, 0x26, 0x90, 0xde, 0x62,
	0x46, 0x1b, 0x8a, 0xc7, 0xde, 0x62, 0x24, 0xef, 0x29, 0xc1, 0x26, 0x98, 0xde, 0x22, 0x18, 0x0e,
	0xa8, 0xa6, 0xf0, 0xff, 0xdc, 0x1f, 0x05, 0x54, 0x47, 0x07, 0x8c, 0xb1, 0xef, 0xcf, 0xa9, 0xa1,
	0x68, 0x32, 0x0a, 0xa6, 0xd4, 0xc4, 0x36, 0x38, 0x8a, 0x02, 0xef, 0x6a, 0x4a, 0x2d, 0xb4, 0x41,
	0x9f, 0x4d, 0x7d, 0x6a, 0x2b, 0xb8, 0xf5, 0x26, 0xd4, 0xe9, 0xbd, 0x10, 0x68, 0xd5, 0xdd, 0xe6,
	0x71, 0x5e, 0xe0, 0x5f, 0x70, 0xb2, 0x4a, 0x73, 0x46, 0xb8, 0x2e, 0x5a, 0x83, 0xce, 0xe5, 0x02,
	0xf2, 0x7c, 0x8f, 0x7f, 0xc0, 0x59, 0x3d, 0x45, 0xab, 0x4d, 0xbe, 0xdf, 0xd6, 0x33, 0x9e, 0xfd,
	0x47, 0x59, 0xfd, 0xa2, 0x6c, 0x39, 0x8d, 0x0a, 0x37, 0xb8, 0x5e, 0x4d, 0xb3, 0x5b, 0xe7, 0x63,
	0xfa, 0x7e, 0x72, 0xc9, 0xc7, 0xc9, 0x25, 0x9f, 0x27, 0x97, 0xbc, 0x7e, 0xb9, 0x8d, 0xa5, 0x55,
	0xfe, 0xe5, 0xf0, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x17, 0x9f, 0x56, 0xc6, 0xda, 0x01, 0x00, 0x00,
}
