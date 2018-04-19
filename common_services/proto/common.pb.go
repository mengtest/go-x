// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: common.proto

/*
	Package proto is a generated protocol buffer package.

	It is generated from these files:
		common.proto

	It has these top-level messages:
		EnumLogin
		MsgLogin
		MsgLoginResult
		MsgVerify
		MsgKick
		MsgForwardC
		MsgForwardS
*/
package proto

import proto1 "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type MsgTypeCmd int32

const (
	MsgTypeCmd_UNSPECIFIED MsgTypeCmd = 0
	MsgTypeCmd_Login       MsgTypeCmd = 1
	MsgTypeCmd_Verify      MsgTypeCmd = 2
	MsgTypeCmd_Kick        MsgTypeCmd = 3
	MsgTypeCmd_ForwardC    MsgTypeCmd = 4
	MsgTypeCmd_ForwardS    MsgTypeCmd = 5
)

var MsgTypeCmd_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "Login",
	2: "Verify",
	3: "Kick",
	4: "ForwardC",
	5: "ForwardS",
}
var MsgTypeCmd_value = map[string]int32{
	"UNSPECIFIED": 0,
	"Login":       1,
	"Verify":      2,
	"Kick":        3,
	"ForwardC":    4,
	"ForwardS":    5,
}

func (x MsgTypeCmd) String() string {
	return proto1.EnumName(MsgTypeCmd_name, int32(x))
}
func (MsgTypeCmd) EnumDescriptor() ([]byte, []int) { return fileDescriptorCommon, []int{0} }

// / Login (C->S)
type LoginMode int32

const (
	LoginMode_Default      LoginMode = 0
	LoginMode_CUSTOM_BEGIN LoginMode = 100
)

var LoginMode_name = map[int32]string{
	0:   "Default",
	100: "CUSTOM_BEGIN",
}
var LoginMode_value = map[string]int32{
	"Default":      0,
	"CUSTOM_BEGIN": 100,
}

func (x LoginMode) String() string {
	return proto1.EnumName(LoginMode_name, int32(x))
}
func (LoginMode) EnumDescriptor() ([]byte, []int) { return fileDescriptorCommon, []int{1} }

type EnumLogin_Error int32

const (
	EnumLogin_NoErr           EnumLogin_Error = 0
	EnumLogin_ErrPassword     EnumLogin_Error = 1
	EnumLogin_ErrAccount      EnumLogin_Error = 2
	EnumLogin_ErrPlatformSide EnumLogin_Error = 3
	EnumLogin_ErrMode         EnumLogin_Error = 4
	EnumLogin_ErrDB           EnumLogin_Error = 5
	EnumLogin_ErrGateway      EnumLogin_Error = 6
	EnumLogin_ErrParams       EnumLogin_Error = 7
)

var EnumLogin_Error_name = map[int32]string{
	0: "NoErr",
	1: "ErrPassword",
	2: "ErrAccount",
	3: "ErrPlatformSide",
	4: "ErrMode",
	5: "ErrDB",
	6: "ErrGateway",
	7: "ErrParams",
}
var EnumLogin_Error_value = map[string]int32{
	"NoErr":           0,
	"ErrPassword":     1,
	"ErrAccount":      2,
	"ErrPlatformSide": 3,
	"ErrMode":         4,
	"ErrDB":           5,
	"ErrGateway":      6,
	"ErrParams":       7,
}

func (x EnumLogin_Error) String() string {
	return proto1.EnumName(EnumLogin_Error_name, int32(x))
}
func (EnumLogin_Error) EnumDescriptor() ([]byte, []int) { return fileDescriptorCommon, []int{0, 0} }

type EnumLogin struct {
}

func (m *EnumLogin) Reset()                    { *m = EnumLogin{} }
func (m *EnumLogin) String() string            { return proto1.CompactTextString(m) }
func (*EnumLogin) ProtoMessage()               {}
func (*EnumLogin) Descriptor() ([]byte, []int) { return fileDescriptorCommon, []int{0} }

type MsgLogin struct {
	Account  string    `protobuf:"bytes,1,opt,name=Account,proto3" json:"Account,omitempty"`
	Password string    `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	Mode     LoginMode `protobuf:"varint,3,opt,name=Mode,proto3,enum=proto.LoginMode" json:"Mode,omitempty"`
	Userdata []byte    `protobuf:"bytes,4,opt,name=Userdata,proto3" json:"Userdata,omitempty"`
}

func (m *MsgLogin) Reset()                    { *m = MsgLogin{} }
func (m *MsgLogin) String() string            { return proto1.CompactTextString(m) }
func (*MsgLogin) ProtoMessage()               {}
func (*MsgLogin) Descriptor() ([]byte, []int) { return fileDescriptorCommon, []int{1} }

func (m *MsgLogin) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *MsgLogin) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *MsgLogin) GetMode() LoginMode {
	if m != nil {
		return m.Mode
	}
	return LoginMode_Default
}

func (m *MsgLogin) GetUserdata() []byte {
	if m != nil {
		return m.Userdata
	}
	return nil
}

type MsgLoginResult struct {
	Err     EnumLogin_Error `protobuf:"varint,1,opt,name=Err,proto3,enum=proto.EnumLogin_Error" json:"Err,omitempty"`
	Token   string          `protobuf:"bytes,2,opt,name=Token,proto3" json:"Token,omitempty"`
	Address string          `protobuf:"bytes,3,opt,name=Address,proto3" json:"Address,omitempty"`
}

func (m *MsgLoginResult) Reset()                    { *m = MsgLoginResult{} }
func (m *MsgLoginResult) String() string            { return proto1.CompactTextString(m) }
func (*MsgLoginResult) ProtoMessage()               {}
func (*MsgLoginResult) Descriptor() ([]byte, []int) { return fileDescriptorCommon, []int{2} }

func (m *MsgLoginResult) GetErr() EnumLogin_Error {
	if m != nil {
		return m.Err
	}
	return EnumLogin_NoErr
}

func (m *MsgLoginResult) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *MsgLoginResult) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

// / Verify (C->S; S->S)
type MsgVerify struct {
	Account string `protobuf:"bytes,1,opt,name=Account,proto3" json:"Account,omitempty"`
	Token   string `protobuf:"bytes,2,opt,name=Token,proto3" json:"Token,omitempty"`
}

func (m *MsgVerify) Reset()                    { *m = MsgVerify{} }
func (m *MsgVerify) String() string            { return proto1.CompactTextString(m) }
func (*MsgVerify) ProtoMessage()               {}
func (*MsgVerify) Descriptor() ([]byte, []int) { return fileDescriptorCommon, []int{3} }

func (m *MsgVerify) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *MsgVerify) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// / KickPlayer (S->S)
type MsgKick struct {
	UID uint64 `protobuf:"varint,1,opt,name=UID,proto3" json:"UID,omitempty"`
}

func (m *MsgKick) Reset()                    { *m = MsgKick{} }
func (m *MsgKick) String() string            { return proto1.CompactTextString(m) }
func (*MsgKick) ProtoMessage()               {}
func (*MsgKick) Descriptor() ([]byte, []int) { return fileDescriptorCommon, []int{4} }

func (m *MsgKick) GetUID() uint64 {
	if m != nil {
		return m.UID
	}
	return 0
}

// / ForwardC (C->S)
type MsgForwardC struct {
	Type uint32 `protobuf:"varint,1,opt,name=Type,proto3" json:"Type,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (m *MsgForwardC) Reset()                    { *m = MsgForwardC{} }
func (m *MsgForwardC) String() string            { return proto1.CompactTextString(m) }
func (*MsgForwardC) ProtoMessage()               {}
func (*MsgForwardC) Descriptor() ([]byte, []int) { return fileDescriptorCommon, []int{5} }

func (m *MsgForwardC) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *MsgForwardC) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// / ForwardS (S->S)
// /    若Type为0且Id不为空:投递给该Id服务器
// /    若Type为1且Id不为空:投递给某Type所有服务器（除了该Id服务器外）
// /    若Type为1且Id为空，则投递给某Type的一个服务器
type MsgForwardS struct {
	Type uint32 `protobuf:"varint,1,opt,name=Type,proto3" json:"Type,omitempty"`
	Id   string `protobuf:"bytes,2,opt,name=Id,proto3" json:"Id,omitempty"`
	Data []byte `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (m *MsgForwardS) Reset()                    { *m = MsgForwardS{} }
func (m *MsgForwardS) String() string            { return proto1.CompactTextString(m) }
func (*MsgForwardS) ProtoMessage()               {}
func (*MsgForwardS) Descriptor() ([]byte, []int) { return fileDescriptorCommon, []int{6} }

func (m *MsgForwardS) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *MsgForwardS) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MsgForwardS) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto1.RegisterType((*EnumLogin)(nil), "proto.EnumLogin")
	proto1.RegisterType((*MsgLogin)(nil), "proto.MsgLogin")
	proto1.RegisterType((*MsgLoginResult)(nil), "proto.MsgLoginResult")
	proto1.RegisterType((*MsgVerify)(nil), "proto.MsgVerify")
	proto1.RegisterType((*MsgKick)(nil), "proto.MsgKick")
	proto1.RegisterType((*MsgForwardC)(nil), "proto.MsgForwardC")
	proto1.RegisterType((*MsgForwardS)(nil), "proto.MsgForwardS")
	proto1.RegisterEnum("proto.MsgTypeCmd", MsgTypeCmd_name, MsgTypeCmd_value)
	proto1.RegisterEnum("proto.LoginMode", LoginMode_name, LoginMode_value)
	proto1.RegisterEnum("proto.EnumLogin_Error", EnumLogin_Error_name, EnumLogin_Error_value)
}
func (m *EnumLogin) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EnumLogin) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *MsgLogin) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgLogin) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Account) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Account)))
		i += copy(dAtA[i:], m.Account)
	}
	if len(m.Password) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Password)))
		i += copy(dAtA[i:], m.Password)
	}
	if m.Mode != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintCommon(dAtA, i, uint64(m.Mode))
	}
	if len(m.Userdata) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Userdata)))
		i += copy(dAtA[i:], m.Userdata)
	}
	return i, nil
}

func (m *MsgLoginResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgLoginResult) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Err != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintCommon(dAtA, i, uint64(m.Err))
	}
	if len(m.Token) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Token)))
		i += copy(dAtA[i:], m.Token)
	}
	if len(m.Address) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Address)))
		i += copy(dAtA[i:], m.Address)
	}
	return i, nil
}

func (m *MsgVerify) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgVerify) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Account) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Account)))
		i += copy(dAtA[i:], m.Account)
	}
	if len(m.Token) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Token)))
		i += copy(dAtA[i:], m.Token)
	}
	return i, nil
}

func (m *MsgKick) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgKick) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.UID != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintCommon(dAtA, i, uint64(m.UID))
	}
	return i, nil
}

func (m *MsgForwardC) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgForwardC) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintCommon(dAtA, i, uint64(m.Type))
	}
	if len(m.Data) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Data)))
		i += copy(dAtA[i:], m.Data)
	}
	return i, nil
}

func (m *MsgForwardS) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgForwardS) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintCommon(dAtA, i, uint64(m.Type))
	}
	if len(m.Id) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if len(m.Data) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Data)))
		i += copy(dAtA[i:], m.Data)
	}
	return i, nil
}

func encodeVarintCommon(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *EnumLogin) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *MsgLogin) Size() (n int) {
	var l int
	_ = l
	l = len(m.Account)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	l = len(m.Password)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	if m.Mode != 0 {
		n += 1 + sovCommon(uint64(m.Mode))
	}
	l = len(m.Userdata)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	return n
}

func (m *MsgLoginResult) Size() (n int) {
	var l int
	_ = l
	if m.Err != 0 {
		n += 1 + sovCommon(uint64(m.Err))
	}
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	return n
}

func (m *MsgVerify) Size() (n int) {
	var l int
	_ = l
	l = len(m.Account)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	return n
}

func (m *MsgKick) Size() (n int) {
	var l int
	_ = l
	if m.UID != 0 {
		n += 1 + sovCommon(uint64(m.UID))
	}
	return n
}

func (m *MsgForwardC) Size() (n int) {
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovCommon(uint64(m.Type))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	return n
}

func (m *MsgForwardS) Size() (n int) {
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovCommon(uint64(m.Type))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	return n
}

func sovCommon(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCommon(x uint64) (n int) {
	return sovCommon(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EnumLogin) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
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
			return fmt.Errorf("proto: EnumLogin: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EnumLogin: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCommon
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
func (m *MsgLogin) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
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
			return fmt.Errorf("proto: MsgLogin: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgLogin: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Account = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Password", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Password = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mode", wireType)
			}
			m.Mode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Mode |= (LoginMode(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Userdata", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Userdata = append(m.Userdata[:0], dAtA[iNdEx:postIndex]...)
			if m.Userdata == nil {
				m.Userdata = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCommon
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
func (m *MsgLoginResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
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
			return fmt.Errorf("proto: MsgLoginResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgLoginResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Err", wireType)
			}
			m.Err = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Err |= (EnumLogin_Error(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCommon
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
func (m *MsgVerify) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
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
			return fmt.Errorf("proto: MsgVerify: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgVerify: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Account = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCommon
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
func (m *MsgKick) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
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
			return fmt.Errorf("proto: MsgKick: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgKick: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UID", wireType)
			}
			m.UID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UID |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCommon
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
func (m *MsgForwardC) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
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
			return fmt.Errorf("proto: MsgForwardC: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgForwardC: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCommon
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
func (m *MsgForwardS) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
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
			return fmt.Errorf("proto: MsgForwardS: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgForwardS: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCommon
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
func skipCommon(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCommon
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
					return 0, ErrIntOverflowCommon
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
					return 0, ErrIntOverflowCommon
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
				return 0, ErrInvalidLengthCommon
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCommon
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
				next, err := skipCommon(dAtA[start:])
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
	ErrInvalidLengthCommon = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCommon   = fmt.Errorf("proto: integer overflow")
)

func init() { proto1.RegisterFile("common.proto", fileDescriptorCommon) }

var fileDescriptorCommon = []byte{
	// 499 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x3f, 0x8f, 0xda, 0x4e,
	0x10, 0x65, 0xfd, 0x07, 0xf0, 0xc0, 0x71, 0xa3, 0xfd, 0xfd, 0x14, 0xa1, 0x44, 0x42, 0xc8, 0x4a,
	0x81, 0x28, 0x28, 0x2e, 0x4a, 0x95, 0x2a, 0xc0, 0xde, 0xc9, 0x4a, 0x4c, 0x4e, 0x06, 0x52, 0xa4,
	0x89, 0x1c, 0x6c, 0x2c, 0x72, 0x98, 0x3d, 0xcd, 0x1a, 0x21, 0xaa, 0x48, 0xa9, 0x53, 0xe4, 0x63,
	0xa5, 0xcc, 0x47, 0x88, 0xc8, 0x17, 0x89, 0x76, 0x01, 0xe7, 0x8a, 0x53, 0x2a, 0xcf, 0xf3, 0xce,
	0xbc, 0x7d, 0xef, 0xed, 0x40, 0x73, 0x21, 0xf3, 0x5c, 0x6e, 0x06, 0xf7, 0x24, 0x0b, 0xc9, 0x5d,
	0xf3, 0xf1, 0xbf, 0x31, 0xf0, 0xc4, 0x66, 0x9b, 0xbf, 0x95, 0xd9, 0x6a, 0xe3, 0x7f, 0x01, 0x57,
	0x10, 0x49, 0xe2, 0x1e, 0xb8, 0x13, 0x29, 0x88, 0xb0, 0xc2, 0x2f, 0xa1, 0x21, 0x88, 0x6e, 0x63,
	0xa5, 0x76, 0x92, 0x12, 0x64, 0xbc, 0x05, 0x20, 0x88, 0x5e, 0x2f, 0x16, 0x72, 0xbb, 0x29, 0xd0,
	0xe2, 0xff, 0xc1, 0xa5, 0x6e, 0x58, 0xc7, 0xc5, 0x52, 0x52, 0x3e, 0x5d, 0x25, 0x29, 0xda, 0xbc,
	0x01, 0x35, 0x41, 0x14, 0xca, 0x24, 0x45, 0x47, 0xb3, 0x09, 0xa2, 0xf1, 0x10, 0xdd, 0xd3, 0xf0,
	0x4d, 0x5c, 0xa4, 0xbb, 0x78, 0x8f, 0x55, 0x7e, 0x01, 0x9e, 0x61, 0xa7, 0x38, 0x57, 0x58, 0xf3,
	0xbf, 0x32, 0xa8, 0x87, 0x2a, 0x33, 0x6a, 0x78, 0x1b, 0x6a, 0xa7, 0x5b, 0xda, 0xac, 0xcb, 0x7a,
	0x5e, 0x74, 0x86, 0xfc, 0x29, 0xd4, 0xcf, 0x82, 0xda, 0x96, 0x39, 0x2a, 0x31, 0x7f, 0x0e, 0x8e,
	0xbe, 0xb6, 0x6d, 0x77, 0x59, 0xaf, 0x75, 0x85, 0x47, 0xbb, 0x03, 0xc3, 0xa8, 0xff, 0x47, 0xe6,
	0x54, 0x33, 0xcc, 0x55, 0x4a, 0x49, 0x5c, 0xc4, 0x6d, 0xa7, 0xcb, 0x7a, 0xcd, 0xa8, 0xc4, 0xfe,
	0x67, 0x68, 0x9d, 0x35, 0x44, 0xa9, 0xda, 0xae, 0x0b, 0xde, 0x03, 0x5b, 0x10, 0x19, 0x15, 0xad,
	0xab, 0x27, 0x27, 0xca, 0x32, 0xb6, 0x81, 0xc9, 0x2c, 0xd2, 0x2d, 0xfc, 0x7f, 0x70, 0x67, 0xf2,
	0x2e, 0xdd, 0x9c, 0x64, 0x1d, 0x81, 0x71, 0x92, 0x24, 0x94, 0x2a, 0x65, 0x64, 0x69, 0x27, 0x47,
	0xe8, 0xbf, 0x02, 0x2f, 0x54, 0xd9, 0xfb, 0x94, 0x56, 0xcb, 0xfd, 0x3f, 0x0c, 0x3f, 0x4a, 0xeb,
	0x3f, 0x83, 0x5a, 0xa8, 0xb2, 0x37, 0xab, 0xc5, 0x1d, 0x47, 0xb0, 0xe7, 0xc1, 0xd8, 0x8c, 0x39,
	0x91, 0x2e, 0xfd, 0x97, 0xd0, 0x08, 0x55, 0x76, 0x2d, 0x69, 0x17, 0x53, 0x32, 0xe2, 0x1c, 0x9c,
	0xd9, 0xfe, 0x3e, 0x35, 0x1d, 0x17, 0x91, 0xa9, 0xf5, 0xbf, 0xb1, 0x0e, 0xc0, 0x32, 0x01, 0x98,
	0xda, 0x17, 0x0f, 0xc7, 0xa6, 0x8f, 0x8e, 0xb5, 0xc0, 0x0a, 0xce, 0xb9, 0x5b, 0x41, 0x52, 0xd2,
	0xd8, 0x7f, 0x69, 0xfa, 0x1f, 0x00, 0x42, 0x95, 0xe9, 0xf6, 0x51, 0x9e, 0xe8, 0x1d, 0x9a, 0x4f,
	0xa6, 0xb7, 0x62, 0x14, 0x5c, 0x07, 0x62, 0x8c, 0x15, 0xbd, 0x11, 0x26, 0x3a, 0x64, 0x1c, 0xa0,
	0x7a, 0xb4, 0x8f, 0x16, 0xaf, 0x83, 0xa3, 0xdd, 0xa0, 0xcd, 0x9b, 0x50, 0x3f, 0x4b, 0x47, 0xe7,
	0x01, 0x9a, 0xa2, 0xdb, 0xef, 0x83, 0x57, 0x3e, 0xa7, 0x5e, 0xb4, 0x71, 0xba, 0x8c, 0xb7, 0xeb,
	0x02, 0x2b, 0x1c, 0xa1, 0x39, 0x9a, 0x4f, 0x67, 0xef, 0xc2, 0x8f, 0x43, 0x71, 0x13, 0x4c, 0x30,
	0x19, 0xe2, 0x8f, 0x43, 0x87, 0xfd, 0x3c, 0x74, 0xd8, 0xaf, 0x43, 0x87, 0x7d, 0xff, 0xdd, 0xa9,
	0x7c, 0xaa, 0x9a, 0xd7, 0x7b, 0xf1, 0x27, 0x00, 0x00, 0xff, 0xff, 0x21, 0xc2, 0x68, 0xdc, 0x0f,
	0x03, 0x00, 0x00,
}