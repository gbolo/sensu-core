// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/sensu/core/v2/tls.proto

package v2

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

// TLSOptions holds TLS options that are used across the varying Sensu
// components
type TLSOptions struct {
	CertFile             string   `protobuf:"bytes,1,opt,name=cert_file,json=certFile,proto3" json:"cert_file,omitempty"`
	KeyFile              string   `protobuf:"bytes,2,opt,name=key_file,json=keyFile,proto3" json:"key_file,omitempty"`
	TrustedCAFile        string   `protobuf:"bytes,3,opt,name=trusted_ca_file,json=trustedCaFile,proto3" json:"trusted_ca_file,omitempty"`
	InsecureSkipVerify   bool     `protobuf:"varint,4,opt,name=insecure_skip_verify,json=insecureSkipVerify,proto3" json:"insecure_skip_verify"`
	ClientAuthType       bool     `protobuf:"varint,5,opt,name=client_auth_type,json=clientAuthType,proto3" json:"client_auth_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TLSOptions) Reset()         { *m = TLSOptions{} }
func (m *TLSOptions) String() string { return proto.CompactTextString(m) }
func (*TLSOptions) ProtoMessage()    {}
func (*TLSOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e2792dbd3e1e1ba, []int{0}
}
func (m *TLSOptions) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TLSOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TLSOptions.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TLSOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TLSOptions.Merge(m, src)
}
func (m *TLSOptions) XXX_Size() int {
	return m.Size()
}
func (m *TLSOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_TLSOptions.DiscardUnknown(m)
}

var xxx_messageInfo_TLSOptions proto.InternalMessageInfo

func (m *TLSOptions) GetCertFile() string {
	if m != nil {
		return m.CertFile
	}
	return ""
}

func (m *TLSOptions) GetKeyFile() string {
	if m != nil {
		return m.KeyFile
	}
	return ""
}

func (m *TLSOptions) GetTrustedCAFile() string {
	if m != nil {
		return m.TrustedCAFile
	}
	return ""
}

func (m *TLSOptions) GetInsecureSkipVerify() bool {
	if m != nil {
		return m.InsecureSkipVerify
	}
	return false
}

func (m *TLSOptions) GetClientAuthType() bool {
	if m != nil {
		return m.ClientAuthType
	}
	return false
}

func init() {
	proto.RegisterType((*TLSOptions)(nil), "sensu.core.v2.TLSOptions")
}

func init() {
	proto.RegisterFile("github.com/sensu/core/v2/tls.proto", fileDescriptor_9e2792dbd3e1e1ba)
}

var fileDescriptor_9e2792dbd3e1e1ba = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x3d, 0x4b, 0xc3, 0x40,
	0x1c, 0x87, 0xbd, 0xfa, 0xd6, 0x1e, 0xd4, 0x97, 0xe0, 0x10, 0x15, 0xae, 0xa5, 0x53, 0xa7, 0x04,
	0x53, 0x17, 0x37, 0xdb, 0x82, 0x83, 0x08, 0x42, 0x5a, 0x1c, 0x5c, 0x42, 0x1b, 0xff, 0x6d, 0x8e,
	0xc4, 0xdc, 0x71, 0x2f, 0x81, 0x7c, 0x13, 0x3f, 0x82, 0x1f, 0xc1, 0x8f, 0xe0, 0xe8, 0x27, 0x28,
	0x7a, 0x6e, 0xee, 0x82, 0xa3, 0xe4, 0xa2, 0xe0, 0xe0, 0x76, 0x3c, 0xcf, 0xf3, 0xff, 0x0d, 0x87,
	0x7b, 0x4b, 0xaa, 0x12, 0x3d, 0xf7, 0x62, 0x76, 0xef, 0x4b, 0xc8, 0xa5, 0xf6, 0x63, 0x26, 0xc0,
	0x2f, 0x02, 0x5f, 0x65, 0xd2, 0xe3, 0x82, 0x29, 0xe6, 0xb4, 0xad, 0xf0, 0x2a, 0xe1, 0x15, 0xc1,
	0xd1, 0xe9, 0x9f, 0x93, 0x25, 0x5b, 0x32, 0xdf, 0x56, 0x73, 0xbd, 0x38, 0x2f, 0x4e, 0xbc, 0x81,
	0x17, 0x58, 0x68, 0x99, 0x7d, 0xd5, 0x23, 0xbd, 0x4f, 0x84, 0xf1, 0xf4, 0x6a, 0x72, 0xcd, 0x15,
	0x65, 0xb9, 0x74, 0x8e, 0x71, 0x2b, 0x06, 0xa1, 0xa2, 0x05, 0xcd, 0xc0, 0x45, 0x5d, 0xd4, 0x6f,
	0x85, 0xcd, 0x0a, 0x5c, 0xd0, 0x0c, 0x9c, 0x43, 0xdc, 0x4c, 0xa1, 0xac, 0x5d, 0xc3, 0xba, 0xed,
	0x14, 0x4a, 0xab, 0xce, 0xf0, 0xae, 0x12, 0x5a, 0x2a, 0xb8, 0x8b, 0xe2, 0x59, 0x5d, 0xac, 0x57,
	0xc5, 0x68, 0xdf, 0xac, 0x3a, 0xed, 0x69, 0xad, 0xc6, 0xc3, 0xaa, 0x0d, 0xdb, 0x3f, 0xe5, 0x78,
	0x66, 0x4f, 0x2f, 0xf1, 0x01, 0xcd, 0x25, 0xc4, 0x5a, 0x40, 0x24, 0x53, 0xca, 0xa3, 0x02, 0x04,
	0x5d, 0x94, 0xee, 0x46, 0x17, 0xf5, 0x9b, 0x23, 0xf7, 0x63, 0xd5, 0xf9, 0xd7, 0x87, 0xce, 0x2f,
	0x9d, 0xa4, 0x94, 0xdf, 0x58, 0xe6, 0xf4, 0xf1, 0x5e, 0x9c, 0x51, 0xc8, 0x55, 0x34, 0xd3, 0x2a,
	0x89, 0x54, 0xc9, 0xc1, 0xdd, 0xac, 0x76, 0xc2, 0x9d, 0x9a, 0x0f, 0xb5, 0x4a, 0xa6, 0x25, 0x87,
	0x51, 0xf7, 0xeb, 0x8d, 0xa0, 0x47, 0x43, 0xd0, 0x93, 0x21, 0xe8, 0xd9, 0x10, 0xf4, 0x62, 0x08,
	0x7a, 0x35, 0x04, 0x3d, 0xbc, 0x93, 0xb5, 0xdb, 0x46, 0x11, 0xcc, 0xb7, 0xec, 0x07, 0x0d, 0xbe,
	0x03, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x10, 0xd7, 0xc4, 0x8b, 0x01, 0x00, 0x00,
}

func (this *TLSOptions) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TLSOptions)
	if !ok {
		that2, ok := that.(TLSOptions)
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
	if this.CertFile != that1.CertFile {
		return false
	}
	if this.KeyFile != that1.KeyFile {
		return false
	}
	if this.TrustedCAFile != that1.TrustedCAFile {
		return false
	}
	if this.InsecureSkipVerify != that1.InsecureSkipVerify {
		return false
	}
	if this.ClientAuthType != that1.ClientAuthType {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (m *TLSOptions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TLSOptions) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TLSOptions) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.ClientAuthType {
		i--
		if m.ClientAuthType {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if m.InsecureSkipVerify {
		i--
		if m.InsecureSkipVerify {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.TrustedCAFile) > 0 {
		i -= len(m.TrustedCAFile)
		copy(dAtA[i:], m.TrustedCAFile)
		i = encodeVarintTls(dAtA, i, uint64(len(m.TrustedCAFile)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.KeyFile) > 0 {
		i -= len(m.KeyFile)
		copy(dAtA[i:], m.KeyFile)
		i = encodeVarintTls(dAtA, i, uint64(len(m.KeyFile)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.CertFile) > 0 {
		i -= len(m.CertFile)
		copy(dAtA[i:], m.CertFile)
		i = encodeVarintTls(dAtA, i, uint64(len(m.CertFile)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTls(dAtA []byte, offset int, v uint64) int {
	offset -= sovTls(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedTLSOptions(r randyTls, easy bool) *TLSOptions {
	this := &TLSOptions{}
	this.CertFile = string(randStringTls(r))
	this.KeyFile = string(randStringTls(r))
	this.TrustedCAFile = string(randStringTls(r))
	this.InsecureSkipVerify = bool(bool(r.Intn(2) == 0))
	this.ClientAuthType = bool(bool(r.Intn(2) == 0))
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTls(r, 6)
	}
	return this
}

type randyTls interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneTls(r randyTls) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringTls(r randyTls) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneTls(r)
	}
	return string(tmps)
}
func randUnrecognizedTls(r randyTls, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldTls(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldTls(dAtA []byte, r randyTls, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateTls(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulateTls(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulateTls(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateTls(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateTls(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateTls(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateTls(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *TLSOptions) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CertFile)
	if l > 0 {
		n += 1 + l + sovTls(uint64(l))
	}
	l = len(m.KeyFile)
	if l > 0 {
		n += 1 + l + sovTls(uint64(l))
	}
	l = len(m.TrustedCAFile)
	if l > 0 {
		n += 1 + l + sovTls(uint64(l))
	}
	if m.InsecureSkipVerify {
		n += 2
	}
	if m.ClientAuthType {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTls(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTls(x uint64) (n int) {
	return sovTls(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TLSOptions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTls
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
			return fmt.Errorf("proto: TLSOptions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TLSOptions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CertFile", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTls
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
				return ErrInvalidLengthTls
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTls
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CertFile = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyFile", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTls
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
				return ErrInvalidLengthTls
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTls
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KeyFile = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TrustedCAFile", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTls
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
				return ErrInvalidLengthTls
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTls
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TrustedCAFile = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InsecureSkipVerify", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTls
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
			m.InsecureSkipVerify = bool(v != 0)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientAuthType", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTls
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
			m.ClientAuthType = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipTls(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTls
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
func skipTls(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTls
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
					return 0, ErrIntOverflowTls
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
					return 0, ErrIntOverflowTls
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
				return 0, ErrInvalidLengthTls
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTls
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTls
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTls        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTls          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTls = fmt.Errorf("proto: unexpected end of group")
)
