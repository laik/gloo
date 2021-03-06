// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/ssl.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// SslConfig contains the options necessary to configure a virtual host or listener to use TLS
type SslConfig struct {
	// Types that are valid to be assigned to SslSecrets:
	//	*SslConfig_SecretRef
	//	*SslConfig_SslFiles
	SslSecrets isSslConfig_SslSecrets `protobuf_oneof:"ssl_secrets"`
	// optional. the SNI domains that should be considered for TLS connections
	SniDomains           []string `protobuf:"bytes,3,rep,name=sni_domains,json=sniDomains,proto3" json:"sni_domains,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SslConfig) Reset()         { *m = SslConfig{} }
func (m *SslConfig) String() string { return proto.CompactTextString(m) }
func (*SslConfig) ProtoMessage()    {}
func (*SslConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4a65e8067d81add, []int{0}
}
func (m *SslConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SslConfig.Unmarshal(m, b)
}
func (m *SslConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SslConfig.Marshal(b, m, deterministic)
}
func (m *SslConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SslConfig.Merge(m, src)
}
func (m *SslConfig) XXX_Size() int {
	return xxx_messageInfo_SslConfig.Size(m)
}
func (m *SslConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_SslConfig.DiscardUnknown(m)
}

var xxx_messageInfo_SslConfig proto.InternalMessageInfo

type isSslConfig_SslSecrets interface {
	isSslConfig_SslSecrets()
	Equal(interface{}) bool
}

type SslConfig_SecretRef struct {
	SecretRef *core.ResourceRef `protobuf:"bytes,1,opt,name=secret_ref,json=secretRef,proto3,oneof"`
}
type SslConfig_SslFiles struct {
	SslFiles *SSLFiles `protobuf:"bytes,2,opt,name=ssl_files,json=sslFiles,proto3,oneof"`
}

func (*SslConfig_SecretRef) isSslConfig_SslSecrets() {}
func (*SslConfig_SslFiles) isSslConfig_SslSecrets()  {}

func (m *SslConfig) GetSslSecrets() isSslConfig_SslSecrets {
	if m != nil {
		return m.SslSecrets
	}
	return nil
}

func (m *SslConfig) GetSecretRef() *core.ResourceRef {
	if x, ok := m.GetSslSecrets().(*SslConfig_SecretRef); ok {
		return x.SecretRef
	}
	return nil
}

func (m *SslConfig) GetSslFiles() *SSLFiles {
	if x, ok := m.GetSslSecrets().(*SslConfig_SslFiles); ok {
		return x.SslFiles
	}
	return nil
}

func (m *SslConfig) GetSniDomains() []string {
	if m != nil {
		return m.SniDomains
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*SslConfig) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _SslConfig_OneofMarshaler, _SslConfig_OneofUnmarshaler, _SslConfig_OneofSizer, []interface{}{
		(*SslConfig_SecretRef)(nil),
		(*SslConfig_SslFiles)(nil),
	}
}

func _SslConfig_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*SslConfig)
	// ssl_secrets
	switch x := m.SslSecrets.(type) {
	case *SslConfig_SecretRef:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SecretRef); err != nil {
			return err
		}
	case *SslConfig_SslFiles:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SslFiles); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("SslConfig.SslSecrets has unexpected type %T", x)
	}
	return nil
}

func _SslConfig_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*SslConfig)
	switch tag {
	case 1: // ssl_secrets.secret_ref
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(core.ResourceRef)
		err := b.DecodeMessage(msg)
		m.SslSecrets = &SslConfig_SecretRef{msg}
		return true, err
	case 2: // ssl_secrets.ssl_files
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SSLFiles)
		err := b.DecodeMessage(msg)
		m.SslSecrets = &SslConfig_SslFiles{msg}
		return true, err
	default:
		return false, nil
	}
}

func _SslConfig_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*SslConfig)
	// ssl_secrets
	switch x := m.SslSecrets.(type) {
	case *SslConfig_SecretRef:
		s := proto.Size(x.SecretRef)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *SslConfig_SslFiles:
		s := proto.Size(x.SslFiles)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// SSLFiles reference paths to certificates which can be read by the proxy off of its local filesystem
type SSLFiles struct {
	TlsCert string `protobuf:"bytes,1,opt,name=tls_cert,json=tlsCert,proto3" json:"tls_cert,omitempty"`
	TlsKey  string `protobuf:"bytes,2,opt,name=tls_key,json=tlsKey,proto3" json:"tls_key,omitempty"`
	// for client cert validation. optional
	RootCa               string   `protobuf:"bytes,3,opt,name=root_ca,json=rootCa,proto3" json:"root_ca,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SSLFiles) Reset()         { *m = SSLFiles{} }
func (m *SSLFiles) String() string { return proto.CompactTextString(m) }
func (*SSLFiles) ProtoMessage()    {}
func (*SSLFiles) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4a65e8067d81add, []int{1}
}
func (m *SSLFiles) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SSLFiles.Unmarshal(m, b)
}
func (m *SSLFiles) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SSLFiles.Marshal(b, m, deterministic)
}
func (m *SSLFiles) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SSLFiles.Merge(m, src)
}
func (m *SSLFiles) XXX_Size() int {
	return xxx_messageInfo_SSLFiles.Size(m)
}
func (m *SSLFiles) XXX_DiscardUnknown() {
	xxx_messageInfo_SSLFiles.DiscardUnknown(m)
}

var xxx_messageInfo_SSLFiles proto.InternalMessageInfo

func (m *SSLFiles) GetTlsCert() string {
	if m != nil {
		return m.TlsCert
	}
	return ""
}

func (m *SSLFiles) GetTlsKey() string {
	if m != nil {
		return m.TlsKey
	}
	return ""
}

func (m *SSLFiles) GetRootCa() string {
	if m != nil {
		return m.RootCa
	}
	return ""
}

// SslConfig contains the options necessary to configure a virtual host or listener to use TLS
type UpstreamSslConfig struct {
	// Types that are valid to be assigned to SslSecrets:
	//	*UpstreamSslConfig_SecretRef
	//	*UpstreamSslConfig_SslFiles
	SslSecrets isUpstreamSslConfig_SslSecrets `protobuf_oneof:"ssl_secrets"`
	// optional. the SNI domains that should be considered for TLS connections
	Sni                  string   `protobuf:"bytes,3,opt,name=sni,proto3" json:"sni,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpstreamSslConfig) Reset()         { *m = UpstreamSslConfig{} }
func (m *UpstreamSslConfig) String() string { return proto.CompactTextString(m) }
func (*UpstreamSslConfig) ProtoMessage()    {}
func (*UpstreamSslConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4a65e8067d81add, []int{2}
}
func (m *UpstreamSslConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpstreamSslConfig.Unmarshal(m, b)
}
func (m *UpstreamSslConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpstreamSslConfig.Marshal(b, m, deterministic)
}
func (m *UpstreamSslConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpstreamSslConfig.Merge(m, src)
}
func (m *UpstreamSslConfig) XXX_Size() int {
	return xxx_messageInfo_UpstreamSslConfig.Size(m)
}
func (m *UpstreamSslConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_UpstreamSslConfig.DiscardUnknown(m)
}

var xxx_messageInfo_UpstreamSslConfig proto.InternalMessageInfo

type isUpstreamSslConfig_SslSecrets interface {
	isUpstreamSslConfig_SslSecrets()
	Equal(interface{}) bool
}

type UpstreamSslConfig_SecretRef struct {
	SecretRef *core.ResourceRef `protobuf:"bytes,1,opt,name=secret_ref,json=secretRef,proto3,oneof"`
}
type UpstreamSslConfig_SslFiles struct {
	SslFiles *SSLFiles `protobuf:"bytes,2,opt,name=ssl_files,json=sslFiles,proto3,oneof"`
}

func (*UpstreamSslConfig_SecretRef) isUpstreamSslConfig_SslSecrets() {}
func (*UpstreamSslConfig_SslFiles) isUpstreamSslConfig_SslSecrets()  {}

func (m *UpstreamSslConfig) GetSslSecrets() isUpstreamSslConfig_SslSecrets {
	if m != nil {
		return m.SslSecrets
	}
	return nil
}

func (m *UpstreamSslConfig) GetSecretRef() *core.ResourceRef {
	if x, ok := m.GetSslSecrets().(*UpstreamSslConfig_SecretRef); ok {
		return x.SecretRef
	}
	return nil
}

func (m *UpstreamSslConfig) GetSslFiles() *SSLFiles {
	if x, ok := m.GetSslSecrets().(*UpstreamSslConfig_SslFiles); ok {
		return x.SslFiles
	}
	return nil
}

func (m *UpstreamSslConfig) GetSni() string {
	if m != nil {
		return m.Sni
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*UpstreamSslConfig) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _UpstreamSslConfig_OneofMarshaler, _UpstreamSslConfig_OneofUnmarshaler, _UpstreamSslConfig_OneofSizer, []interface{}{
		(*UpstreamSslConfig_SecretRef)(nil),
		(*UpstreamSslConfig_SslFiles)(nil),
	}
}

func _UpstreamSslConfig_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*UpstreamSslConfig)
	// ssl_secrets
	switch x := m.SslSecrets.(type) {
	case *UpstreamSslConfig_SecretRef:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SecretRef); err != nil {
			return err
		}
	case *UpstreamSslConfig_SslFiles:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SslFiles); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("UpstreamSslConfig.SslSecrets has unexpected type %T", x)
	}
	return nil
}

func _UpstreamSslConfig_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*UpstreamSslConfig)
	switch tag {
	case 1: // ssl_secrets.secret_ref
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(core.ResourceRef)
		err := b.DecodeMessage(msg)
		m.SslSecrets = &UpstreamSslConfig_SecretRef{msg}
		return true, err
	case 2: // ssl_secrets.ssl_files
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SSLFiles)
		err := b.DecodeMessage(msg)
		m.SslSecrets = &UpstreamSslConfig_SslFiles{msg}
		return true, err
	default:
		return false, nil
	}
}

func _UpstreamSslConfig_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*UpstreamSslConfig)
	// ssl_secrets
	switch x := m.SslSecrets.(type) {
	case *UpstreamSslConfig_SecretRef:
		s := proto.Size(x.SecretRef)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *UpstreamSslConfig_SslFiles:
		s := proto.Size(x.SslFiles)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*SslConfig)(nil), "gloo.solo.io.SslConfig")
	proto.RegisterType((*SSLFiles)(nil), "gloo.solo.io.SSLFiles")
	proto.RegisterType((*UpstreamSslConfig)(nil), "gloo.solo.io.UpstreamSslConfig")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/ssl.proto", fileDescriptor_c4a65e8067d81add)
}

var fileDescriptor_c4a65e8067d81add = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x52, 0xcf, 0x4b, 0xeb, 0x40,
	0x10, 0x6e, 0x5e, 0xa0, 0xaf, 0xd9, 0xbe, 0x07, 0x1a, 0x44, 0xd3, 0x1e, 0xb4, 0xf4, 0xd4, 0x83,
	0x6e, 0xfc, 0x81, 0x3d, 0x78, 0x6c, 0x45, 0x0a, 0x7a, 0xda, 0x22, 0x82, 0x97, 0x90, 0xc6, 0x49,
	0x5c, 0xbb, 0xcd, 0x84, 0x9d, 0xad, 0xd0, 0xbf, 0x48, 0x4f, 0xfe, 0x51, 0xfe, 0x25, 0xb2, 0xd9,
	0x16, 0x44, 0x3c, 0x78, 0xf3, 0xb4, 0xdf, 0xcc, 0x37, 0xf3, 0xcd, 0xc7, 0xcc, 0xb2, 0x61, 0x21,
	0xcd, 0xe3, 0x72, 0xc6, 0x33, 0x5c, 0xc4, 0x84, 0x0a, 0x8f, 0x24, 0xc6, 0x85, 0x42, 0x8c, 0x2b,
	0x8d, 0x4f, 0x90, 0x19, 0x72, 0x51, 0x5a, 0xc9, 0xf8, 0xf9, 0x24, 0x26, 0x52, 0xbc, 0xd2, 0x68,
	0x30, 0xfc, 0x67, 0xd3, 0xdc, 0x76, 0x70, 0x89, 0xdd, 0x9d, 0x02, 0x0b, 0xac, 0x89, 0xd8, 0x22,
	0x57, 0xd3, 0x3d, 0xfc, 0x46, 0xbb, 0x7e, 0xe7, 0xd2, 0x6c, 0x14, 0x35, 0xe4, 0xae, 0xba, 0xff,
	0xe6, 0xb1, 0x60, 0x4a, 0x6a, 0x8c, 0x65, 0x2e, 0x8b, 0xf0, 0x82, 0x31, 0x82, 0x4c, 0x83, 0x49,
	0x34, 0xe4, 0x91, 0xd7, 0xf3, 0x06, 0xed, 0xd3, 0x0e, 0xcf, 0x50, 0xc3, 0x66, 0x28, 0x17, 0x40,
	0xb8, 0xd4, 0x19, 0x08, 0xc8, 0x27, 0x0d, 0x11, 0xb8, 0x72, 0x01, 0x79, 0x78, 0xce, 0x02, 0x22,
	0x95, 0xe4, 0x52, 0x01, 0x45, 0x7f, 0xea, 0xd6, 0x5d, 0xfe, 0xd9, 0x2f, 0x9f, 0x4e, 0x6f, 0xae,
	0x2c, 0x3b, 0x69, 0x88, 0x16, 0x91, 0xaa, 0x71, 0x78, 0xc0, 0xda, 0x54, 0xca, 0xe4, 0x01, 0x17,
	0xa9, 0x2c, 0x29, 0xf2, 0x7b, 0xfe, 0x20, 0x10, 0x8c, 0x4a, 0x79, 0xe9, 0x32, 0xa3, 0xff, 0xac,
	0x6d, 0x75, 0xdd, 0x20, 0xea, 0xdf, 0xb1, 0xd6, 0x46, 0x27, 0xec, 0xb0, 0x96, 0x51, 0x94, 0x64,
	0xa0, 0x4d, 0x6d, 0x36, 0x10, 0x7f, 0x8d, 0xa2, 0x31, 0x68, 0x13, 0xee, 0x31, 0x0b, 0x93, 0x39,
	0xac, 0x6a, 0x2f, 0x81, 0x68, 0x1a, 0x45, 0xd7, 0xb0, 0xb2, 0x84, 0x46, 0x34, 0x49, 0x96, 0x46,
	0xbe, 0x23, 0x6c, 0x38, 0x4e, 0xfb, 0x2f, 0x1e, 0xdb, 0xbe, 0xad, 0xc8, 0x68, 0x48, 0x17, 0xbf,
	0xba, 0x91, 0x2d, 0xe6, 0x53, 0x29, 0xd7, 0xee, 0x2c, 0xfc, 0xb2, 0x82, 0xd1, 0xf0, 0xf5, 0x7d,
	0xdf, 0xbb, 0x3f, 0xfe, 0xd9, 0x1f, 0xaa, 0xe6, 0xc5, 0xfa, 0xea, 0xb3, 0x66, 0x7d, 0xf2, 0xb3,
	0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4b, 0x0b, 0x95, 0xa0, 0x7e, 0x02, 0x00, 0x00,
}

func (this *SslConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SslConfig)
	if !ok {
		that2, ok := that.(SslConfig)
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
	if that1.SslSecrets == nil {
		if this.SslSecrets != nil {
			return false
		}
	} else if this.SslSecrets == nil {
		return false
	} else if !this.SslSecrets.Equal(that1.SslSecrets) {
		return false
	}
	if len(this.SniDomains) != len(that1.SniDomains) {
		return false
	}
	for i := range this.SniDomains {
		if this.SniDomains[i] != that1.SniDomains[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *SslConfig_SecretRef) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SslConfig_SecretRef)
	if !ok {
		that2, ok := that.(SslConfig_SecretRef)
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
	if !this.SecretRef.Equal(that1.SecretRef) {
		return false
	}
	return true
}
func (this *SslConfig_SslFiles) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SslConfig_SslFiles)
	if !ok {
		that2, ok := that.(SslConfig_SslFiles)
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
	if !this.SslFiles.Equal(that1.SslFiles) {
		return false
	}
	return true
}
func (this *SSLFiles) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SSLFiles)
	if !ok {
		that2, ok := that.(SSLFiles)
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
	if this.TlsCert != that1.TlsCert {
		return false
	}
	if this.TlsKey != that1.TlsKey {
		return false
	}
	if this.RootCa != that1.RootCa {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *UpstreamSslConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpstreamSslConfig)
	if !ok {
		that2, ok := that.(UpstreamSslConfig)
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
	if that1.SslSecrets == nil {
		if this.SslSecrets != nil {
			return false
		}
	} else if this.SslSecrets == nil {
		return false
	} else if !this.SslSecrets.Equal(that1.SslSecrets) {
		return false
	}
	if this.Sni != that1.Sni {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *UpstreamSslConfig_SecretRef) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpstreamSslConfig_SecretRef)
	if !ok {
		that2, ok := that.(UpstreamSslConfig_SecretRef)
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
	if !this.SecretRef.Equal(that1.SecretRef) {
		return false
	}
	return true
}
func (this *UpstreamSslConfig_SslFiles) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpstreamSslConfig_SslFiles)
	if !ok {
		that2, ok := that.(UpstreamSslConfig_SslFiles)
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
	if !this.SslFiles.Equal(that1.SslFiles) {
		return false
	}
	return true
}
