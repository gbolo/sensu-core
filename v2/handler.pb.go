// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: handler.proto

package v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import bytes "bytes"

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

// A Handler is a handler specification.
type Handler struct {
	// Metadata contains the name, namespace, labels and annotations of the handler
	ObjectMeta `protobuf:"bytes,1,opt,name=metadata,embedded=metadata" json:"metadata"`
	// Type is the handler type, i.e. pipe.
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	// Mutator is the handler event data mutator.
	Mutator string `protobuf:"bytes,3,opt,name=mutator,proto3" json:"mutator,omitempty"`
	// Command is the command to be executed for a pipe handler.
	Command string `protobuf:"bytes,4,opt,name=command,proto3" json:"command,omitempty"`
	// Timeout is the handler timeout in seconds.
	Timeout uint32 `protobuf:"varint,5,opt,name=timeout,proto3" json:"timeout"`
	// Socket contains configuration for a TCP or UDP handler.
	Socket *HandlerSocket `protobuf:"bytes,6,opt,name=socket" json:"socket,omitempty"`
	// Handlers is a list of handlers for a handler set.
	Handlers []string `protobuf:"bytes,7,rep,name=handlers" json:"handlers"`
	// Filters is a list of filters name to evaluate before executing this handler
	Filters []string `protobuf:"bytes,8,rep,name=filters" json:"filters"`
	// EnvVars is a list of environment variables to use with command execution
	EnvVars []string `protobuf:"bytes,9,rep,name=env_vars,json=envVars" json:"env_vars"`
	// RuntimeAssets are a list of assets required to execute a handler.
	RuntimeAssets        []string `protobuf:"bytes,13,rep,name=runtime_assets,json=runtimeAssets" json:"runtime_assets"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Handler) Reset()         { *m = Handler{} }
func (m *Handler) String() string { return proto.CompactTextString(m) }
func (*Handler) ProtoMessage()    {}
func (*Handler) Descriptor() ([]byte, []int) {
	return fileDescriptor_handler_dead5ddb763948ac, []int{0}
}
func (m *Handler) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Handler) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Handler.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Handler) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Handler.Merge(dst, src)
}
func (m *Handler) XXX_Size() int {
	return m.Size()
}
func (m *Handler) XXX_DiscardUnknown() {
	xxx_messageInfo_Handler.DiscardUnknown(m)
}

var xxx_messageInfo_Handler proto.InternalMessageInfo

func (m *Handler) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Handler) GetMutator() string {
	if m != nil {
		return m.Mutator
	}
	return ""
}

func (m *Handler) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *Handler) GetTimeout() uint32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *Handler) GetSocket() *HandlerSocket {
	if m != nil {
		return m.Socket
	}
	return nil
}

func (m *Handler) GetHandlers() []string {
	if m != nil {
		return m.Handlers
	}
	return nil
}

func (m *Handler) GetFilters() []string {
	if m != nil {
		return m.Filters
	}
	return nil
}

func (m *Handler) GetEnvVars() []string {
	if m != nil {
		return m.EnvVars
	}
	return nil
}

func (m *Handler) GetRuntimeAssets() []string {
	if m != nil {
		return m.RuntimeAssets
	}
	return nil
}

// HandlerSocket contains configuration for a TCP or UDP handler.
type HandlerSocket struct {
	// Host is the socket peer address.
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	// Port is the socket peer port.
	Port                 uint32   `protobuf:"varint,2,opt,name=port,proto3" json:"port"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HandlerSocket) Reset()         { *m = HandlerSocket{} }
func (m *HandlerSocket) String() string { return proto.CompactTextString(m) }
func (*HandlerSocket) ProtoMessage()    {}
func (*HandlerSocket) Descriptor() ([]byte, []int) {
	return fileDescriptor_handler_dead5ddb763948ac, []int{1}
}
func (m *HandlerSocket) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HandlerSocket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HandlerSocket.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *HandlerSocket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HandlerSocket.Merge(dst, src)
}
func (m *HandlerSocket) XXX_Size() int {
	return m.Size()
}
func (m *HandlerSocket) XXX_DiscardUnknown() {
	xxx_messageInfo_HandlerSocket.DiscardUnknown(m)
}

var xxx_messageInfo_HandlerSocket proto.InternalMessageInfo

func (m *HandlerSocket) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *HandlerSocket) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func init() {
	proto.RegisterType((*Handler)(nil), "sensu.core.v2.Handler")
	proto.RegisterType((*HandlerSocket)(nil), "sensu.core.v2.HandlerSocket")
}
func (this *Handler) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Handler)
	if !ok {
		that2, ok := that.(Handler)
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
	if !this.ObjectMeta.Equal(&that1.ObjectMeta) {
		return false
	}
	if this.Type != that1.Type {
		return false
	}
	if this.Mutator != that1.Mutator {
		return false
	}
	if this.Command != that1.Command {
		return false
	}
	if this.Timeout != that1.Timeout {
		return false
	}
	if !this.Socket.Equal(that1.Socket) {
		return false
	}
	if len(this.Handlers) != len(that1.Handlers) {
		return false
	}
	for i := range this.Handlers {
		if this.Handlers[i] != that1.Handlers[i] {
			return false
		}
	}
	if len(this.Filters) != len(that1.Filters) {
		return false
	}
	for i := range this.Filters {
		if this.Filters[i] != that1.Filters[i] {
			return false
		}
	}
	if len(this.EnvVars) != len(that1.EnvVars) {
		return false
	}
	for i := range this.EnvVars {
		if this.EnvVars[i] != that1.EnvVars[i] {
			return false
		}
	}
	if len(this.RuntimeAssets) != len(that1.RuntimeAssets) {
		return false
	}
	for i := range this.RuntimeAssets {
		if this.RuntimeAssets[i] != that1.RuntimeAssets[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *HandlerSocket) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HandlerSocket)
	if !ok {
		that2, ok := that.(HandlerSocket)
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
	if this.Host != that1.Host {
		return false
	}
	if this.Port != that1.Port {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (m *Handler) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Handler) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintHandler(dAtA, i, uint64(m.ObjectMeta.Size()))
	n1, err := m.ObjectMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if len(m.Type) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintHandler(dAtA, i, uint64(len(m.Type)))
		i += copy(dAtA[i:], m.Type)
	}
	if len(m.Mutator) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintHandler(dAtA, i, uint64(len(m.Mutator)))
		i += copy(dAtA[i:], m.Mutator)
	}
	if len(m.Command) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintHandler(dAtA, i, uint64(len(m.Command)))
		i += copy(dAtA[i:], m.Command)
	}
	if m.Timeout != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintHandler(dAtA, i, uint64(m.Timeout))
	}
	if m.Socket != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintHandler(dAtA, i, uint64(m.Socket.Size()))
		n2, err := m.Socket.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.Handlers) > 0 {
		for _, s := range m.Handlers {
			dAtA[i] = 0x3a
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
	if len(m.Filters) > 0 {
		for _, s := range m.Filters {
			dAtA[i] = 0x42
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
	if len(m.EnvVars) > 0 {
		for _, s := range m.EnvVars {
			dAtA[i] = 0x4a
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
	if len(m.RuntimeAssets) > 0 {
		for _, s := range m.RuntimeAssets {
			dAtA[i] = 0x6a
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
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *HandlerSocket) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HandlerSocket) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Host) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintHandler(dAtA, i, uint64(len(m.Host)))
		i += copy(dAtA[i:], m.Host)
	}
	if m.Port != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintHandler(dAtA, i, uint64(m.Port))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintHandler(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedHandler(r randyHandler, easy bool) *Handler {
	this := &Handler{}
	v1 := NewPopulatedObjectMeta(r, easy)
	this.ObjectMeta = *v1
	this.Type = string(randStringHandler(r))
	this.Mutator = string(randStringHandler(r))
	this.Command = string(randStringHandler(r))
	this.Timeout = uint32(r.Uint32())
	if r.Intn(10) != 0 {
		this.Socket = NewPopulatedHandlerSocket(r, easy)
	}
	v2 := r.Intn(10)
	this.Handlers = make([]string, v2)
	for i := 0; i < v2; i++ {
		this.Handlers[i] = string(randStringHandler(r))
	}
	v3 := r.Intn(10)
	this.Filters = make([]string, v3)
	for i := 0; i < v3; i++ {
		this.Filters[i] = string(randStringHandler(r))
	}
	v4 := r.Intn(10)
	this.EnvVars = make([]string, v4)
	for i := 0; i < v4; i++ {
		this.EnvVars[i] = string(randStringHandler(r))
	}
	v5 := r.Intn(10)
	this.RuntimeAssets = make([]string, v5)
	for i := 0; i < v5; i++ {
		this.RuntimeAssets[i] = string(randStringHandler(r))
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedHandler(r, 14)
	}
	return this
}

func NewPopulatedHandlerSocket(r randyHandler, easy bool) *HandlerSocket {
	this := &HandlerSocket{}
	this.Host = string(randStringHandler(r))
	this.Port = uint32(r.Uint32())
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedHandler(r, 3)
	}
	return this
}

type randyHandler interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneHandler(r randyHandler) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringHandler(r randyHandler) string {
	v6 := r.Intn(100)
	tmps := make([]rune, v6)
	for i := 0; i < v6; i++ {
		tmps[i] = randUTF8RuneHandler(r)
	}
	return string(tmps)
}
func randUnrecognizedHandler(r randyHandler, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldHandler(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldHandler(dAtA []byte, r randyHandler, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateHandler(dAtA, uint64(key))
		v7 := r.Int63()
		if r.Intn(2) == 0 {
			v7 *= -1
		}
		dAtA = encodeVarintPopulateHandler(dAtA, uint64(v7))
	case 1:
		dAtA = encodeVarintPopulateHandler(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateHandler(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateHandler(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateHandler(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateHandler(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Handler) Size() (n int) {
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovHandler(uint64(l))
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovHandler(uint64(l))
	}
	l = len(m.Mutator)
	if l > 0 {
		n += 1 + l + sovHandler(uint64(l))
	}
	l = len(m.Command)
	if l > 0 {
		n += 1 + l + sovHandler(uint64(l))
	}
	if m.Timeout != 0 {
		n += 1 + sovHandler(uint64(m.Timeout))
	}
	if m.Socket != nil {
		l = m.Socket.Size()
		n += 1 + l + sovHandler(uint64(l))
	}
	if len(m.Handlers) > 0 {
		for _, s := range m.Handlers {
			l = len(s)
			n += 1 + l + sovHandler(uint64(l))
		}
	}
	if len(m.Filters) > 0 {
		for _, s := range m.Filters {
			l = len(s)
			n += 1 + l + sovHandler(uint64(l))
		}
	}
	if len(m.EnvVars) > 0 {
		for _, s := range m.EnvVars {
			l = len(s)
			n += 1 + l + sovHandler(uint64(l))
		}
	}
	if len(m.RuntimeAssets) > 0 {
		for _, s := range m.RuntimeAssets {
			l = len(s)
			n += 1 + l + sovHandler(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *HandlerSocket) Size() (n int) {
	var l int
	_ = l
	l = len(m.Host)
	if l > 0 {
		n += 1 + l + sovHandler(uint64(l))
	}
	if m.Port != 0 {
		n += 1 + sovHandler(uint64(m.Port))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovHandler(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozHandler(x uint64) (n int) {
	return sovHandler(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Handler) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHandler
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
			return fmt.Errorf("proto: Handler: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Handler: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandler
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
				return ErrInvalidLengthHandler
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandler
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
				return ErrInvalidLengthHandler
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mutator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandler
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
				return ErrInvalidLengthHandler
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Mutator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Command", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandler
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
				return ErrInvalidLengthHandler
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Command = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timeout", wireType)
			}
			m.Timeout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandler
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timeout |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Socket", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandler
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
				return ErrInvalidLengthHandler
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Socket == nil {
				m.Socket = &HandlerSocket{}
			}
			if err := m.Socket.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Handlers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandler
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
				return ErrInvalidLengthHandler
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Handlers = append(m.Handlers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Filters", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandler
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
				return ErrInvalidLengthHandler
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Filters = append(m.Filters, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnvVars", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandler
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
				return ErrInvalidLengthHandler
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EnvVars = append(m.EnvVars, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RuntimeAssets", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandler
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
				return ErrInvalidLengthHandler
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RuntimeAssets = append(m.RuntimeAssets, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHandler(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHandler
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
func (m *HandlerSocket) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHandler
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
			return fmt.Errorf("proto: HandlerSocket: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HandlerSocket: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Host", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandler
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
				return ErrInvalidLengthHandler
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Host = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
			}
			m.Port = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandler
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Port |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipHandler(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHandler
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
func skipHandler(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHandler
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
					return 0, ErrIntOverflowHandler
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
					return 0, ErrIntOverflowHandler
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
				return 0, ErrInvalidLengthHandler
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowHandler
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
				next, err := skipHandler(dAtA[start:])
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
	ErrInvalidLengthHandler = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHandler   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("handler.proto", fileDescriptor_handler_dead5ddb763948ac) }

var fileDescriptor_handler_dead5ddb763948ac = []byte{
	// 418 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x52, 0x41, 0x6e, 0xd4, 0x30,
	0x14, 0xad, 0x3b, 0xc3, 0x24, 0x71, 0x09, 0x0b, 0x8b, 0x85, 0xa9, 0xaa, 0x24, 0xaa, 0x84, 0xc8,
	0x86, 0x54, 0x1a, 0x56, 0xb0, 0x6b, 0x56, 0x48, 0x08, 0x21, 0x19, 0x89, 0x05, 0x9b, 0xca, 0xc9,
	0xb8, 0x33, 0x03, 0x4d, 0x3c, 0xb2, 0x7f, 0x22, 0x71, 0x13, 0x8e, 0xc0, 0x09, 0x10, 0x47, 0x98,
	0x65, 0x4f, 0x10, 0x41, 0xd8, 0xe5, 0x04, 0x2c, 0x51, 0x7e, 0x92, 0xa9, 0xda, 0x4d, 0xf4, 0xde,
	0xfb, 0xef, 0x3b, 0xff, 0xf9, 0x9b, 0xfa, 0x1b, 0x59, 0xae, 0x6e, 0x94, 0x49, 0x76, 0x46, 0x83,
	0x66, 0xbe, 0x55, 0xa5, 0xad, 0x92, 0x5c, 0x1b, 0x95, 0xd4, 0xcb, 0xd3, 0x97, 0xeb, 0x2d, 0x6c,
	0xaa, 0x2c, 0xc9, 0x75, 0x71, 0xb1, 0xd6, 0x6b, 0x7d, 0x81, 0xae, 0xac, 0xba, 0x46, 0x86, 0x04,
	0xd1, 0xd0, 0x7d, 0x4a, 0x0b, 0x05, 0x72, 0xc0, 0xe7, 0x3f, 0x67, 0xd4, 0x79, 0x3b, 0x9c, 0xcd,
	0xde, 0x51, 0xb7, 0xaf, 0xac, 0x24, 0x48, 0x4e, 0x22, 0x12, 0x9f, 0x2c, 0x9f, 0x25, 0xf7, 0x7e,
	0x94, 0x7c, 0xc8, 0xbe, 0xa8, 0x1c, 0xde, 0x2b, 0x90, 0xe9, 0xd3, 0x7d, 0x13, 0x1e, 0xdd, 0x36,
	0x21, 0xe9, 0x9a, 0xf0, 0xd0, 0x26, 0x0e, 0x88, 0x31, 0x3a, 0x87, 0x6f, 0x3b, 0xc5, 0x8f, 0x23,
	0x12, 0x7b, 0x02, 0x31, 0xe3, 0xd4, 0x29, 0x2a, 0x90, 0xa0, 0x0d, 0x9f, 0xa1, 0x3c, 0xd1, 0xbe,
	0x92, 0xeb, 0xa2, 0x90, 0xe5, 0x8a, 0xcf, 0x87, 0xca, 0x48, 0xd9, 0x73, 0xea, 0xc0, 0xb6, 0x50,
	0xba, 0x02, 0xfe, 0x28, 0x22, 0xb1, 0x9f, 0x9e, 0x74, 0x4d, 0x38, 0x49, 0x62, 0x02, 0xec, 0x0d,
	0x5d, 0x58, 0x9d, 0x7f, 0x55, 0xc0, 0x17, 0x38, 0xf9, 0xd9, 0x83, 0xc9, 0xc7, 0x8c, 0x1f, 0xd1,
	0x93, 0xce, 0xf7, 0x4d, 0x48, 0xc4, 0xd8, 0xc1, 0x62, 0xea, 0x8e, 0xd7, 0x6b, 0xb9, 0x13, 0xcd,
	0x62, 0x2f, 0x7d, 0xdc, 0x87, 0x9a, 0x34, 0x71, 0x40, 0xfd, 0x30, 0xd7, 0xdb, 0x1b, 0xe8, 0x8d,
	0x2e, 0x1a, 0x71, 0x98, 0x51, 0x12, 0x13, 0x60, 0x2f, 0xa8, 0xab, 0xca, 0xfa, 0xaa, 0x96, 0xc6,
	0x72, 0xef, 0xee, 0xc0, 0x49, 0x13, 0x8e, 0x2a, 0xeb, 0x4f, 0xd2, 0x58, 0xf6, 0x9a, 0x3e, 0x31,
	0x55, 0xd9, 0x67, 0xb8, 0x92, 0xd6, 0x2a, 0xb0, 0xdc, 0x47, 0x3b, 0xeb, 0x9a, 0xf0, 0x41, 0x45,
	0xf8, 0x23, 0xbf, 0x44, 0x7a, 0x7e, 0x49, 0xfd, 0x7b, 0x99, 0xfa, 0x0b, 0xdf, 0x68, 0x0b, 0xb8,
	0x39, 0x4f, 0x20, 0x66, 0x67, 0x74, 0xbe, 0xd3, 0x06, 0x70, 0x09, 0x7e, 0xea, 0x76, 0x4d, 0x88,
	0x5c, 0xe0, 0x37, 0x8d, 0xfe, 0xfd, 0x09, 0xc8, 0x8f, 0x36, 0x20, 0xbf, 0xda, 0x80, 0xec, 0xdb,
	0x80, 0xdc, 0xb6, 0x01, 0xf9, 0xdd, 0x06, 0xe4, 0xfb, 0xdf, 0xe0, 0xe8, 0xf3, 0x71, 0xbd, 0xcc,
	0x16, 0xf8, 0x48, 0x5e, 0xfd, 0x0f, 0x00, 0x00, 0xff, 0xff, 0xb7, 0xe9, 0x64, 0x8a, 0x7f, 0x02,
	0x00, 0x00,
}
