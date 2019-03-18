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
	//	*SslConfig_Sds
	SslSecrets isSslConfig_SslSecrets `protobuf_oneof:"ssl_secrets"`
	// optional. the SNI domains that should be considered for TLS connections
	SniDomains []string `protobuf:"bytes,3,rep,name=sni_domains,json=sniDomains,proto3" json:"sni_domains,omitempty"`
	// Verify that the Subject Alternative Name in the peer certificate is one of the specified values.
	// note that a root_ca must be provided if this option is used.
	VerifySubjectAltName []string `protobuf:"bytes,5,rep,name=verify_subject_alt_name,json=verifySubjectAltName,proto3" json:"verify_subject_alt_name,omitempty"`
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
type SslConfig_Sds struct {
	Sds *SDSConfig `protobuf:"bytes,4,opt,name=sds,proto3,oneof"`
}

func (*SslConfig_SecretRef) isSslConfig_SslSecrets() {}
func (*SslConfig_SslFiles) isSslConfig_SslSecrets()  {}
func (*SslConfig_Sds) isSslConfig_SslSecrets()       {}

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

func (m *SslConfig) GetSds() *SDSConfig {
	if x, ok := m.GetSslSecrets().(*SslConfig_Sds); ok {
		return x.Sds
	}
	return nil
}

func (m *SslConfig) GetSniDomains() []string {
	if m != nil {
		return m.SniDomains
	}
	return nil
}

func (m *SslConfig) GetVerifySubjectAltName() []string {
	if m != nil {
		return m.VerifySubjectAltName
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*SslConfig) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _SslConfig_OneofMarshaler, _SslConfig_OneofUnmarshaler, _SslConfig_OneofSizer, []interface{}{
		(*SslConfig_SecretRef)(nil),
		(*SslConfig_SslFiles)(nil),
		(*SslConfig_Sds)(nil),
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
	case *SslConfig_Sds:
		_ = b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Sds); err != nil {
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
	case 4: // ssl_secrets.sds
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SDSConfig)
		err := b.DecodeMessage(msg)
		m.SslSecrets = &SslConfig_Sds{msg}
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
	case *SslConfig_Sds:
		s := proto.Size(x.Sds)
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
	//	*UpstreamSslConfig_Sds
	SslSecrets isUpstreamSslConfig_SslSecrets `protobuf_oneof:"ssl_secrets"`
	// optional. the SNI domains that should be considered for TLS connections
	Sni string `protobuf:"bytes,3,opt,name=sni,proto3" json:"sni,omitempty"`
	// Verify that the Subject Alternative Name in the peer certificate is one of the specified values.
	// note that a root_ca must be provided if this option is used.
	VerifySubjectAltName []string `protobuf:"bytes,5,rep,name=verify_subject_alt_name,json=verifySubjectAltName,proto3" json:"verify_subject_alt_name,omitempty"`
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
type UpstreamSslConfig_Sds struct {
	Sds *SDSConfig `protobuf:"bytes,4,opt,name=sds,proto3,oneof"`
}

func (*UpstreamSslConfig_SecretRef) isUpstreamSslConfig_SslSecrets() {}
func (*UpstreamSslConfig_SslFiles) isUpstreamSslConfig_SslSecrets()  {}
func (*UpstreamSslConfig_Sds) isUpstreamSslConfig_SslSecrets()       {}

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

func (m *UpstreamSslConfig) GetSds() *SDSConfig {
	if x, ok := m.GetSslSecrets().(*UpstreamSslConfig_Sds); ok {
		return x.Sds
	}
	return nil
}

func (m *UpstreamSslConfig) GetSni() string {
	if m != nil {
		return m.Sni
	}
	return ""
}

func (m *UpstreamSslConfig) GetVerifySubjectAltName() []string {
	if m != nil {
		return m.VerifySubjectAltName
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*UpstreamSslConfig) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _UpstreamSslConfig_OneofMarshaler, _UpstreamSslConfig_OneofUnmarshaler, _UpstreamSslConfig_OneofSizer, []interface{}{
		(*UpstreamSslConfig_SecretRef)(nil),
		(*UpstreamSslConfig_SslFiles)(nil),
		(*UpstreamSslConfig_Sds)(nil),
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
	case *UpstreamSslConfig_Sds:
		_ = b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Sds); err != nil {
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
	case 4: // ssl_secrets.sds
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SDSConfig)
		err := b.DecodeMessage(msg)
		m.SslSecrets = &UpstreamSslConfig_Sds{msg}
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
	case *UpstreamSslConfig_Sds:
		s := proto.Size(x.Sds)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type SDSConfig struct {
	TargetUri              string           `protobuf:"bytes,1,opt,name=TargetUri,proto3" json:"TargetUri,omitempty"`
	CallCredentials        *CallCredentials `protobuf:"bytes,2,opt,name=call_credentials,json=callCredentials,proto3" json:"call_credentials,omitempty"`
	CertificatesSecretName string           `protobuf:"bytes,3,opt,name=certificates_secret_name,json=certificatesSecretName,proto3" json:"certificates_secret_name,omitempty"`
	ValidationContextName  string           `protobuf:"bytes,4,opt,name=validation_context_name,json=validationContextName,proto3" json:"validation_context_name,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}         `json:"-"`
	XXX_unrecognized       []byte           `json:"-"`
	XXX_sizecache          int32            `json:"-"`
}

func (m *SDSConfig) Reset()         { *m = SDSConfig{} }
func (m *SDSConfig) String() string { return proto.CompactTextString(m) }
func (*SDSConfig) ProtoMessage()    {}
func (*SDSConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4a65e8067d81add, []int{3}
}
func (m *SDSConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SDSConfig.Unmarshal(m, b)
}
func (m *SDSConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SDSConfig.Marshal(b, m, deterministic)
}
func (m *SDSConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SDSConfig.Merge(m, src)
}
func (m *SDSConfig) XXX_Size() int {
	return xxx_messageInfo_SDSConfig.Size(m)
}
func (m *SDSConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_SDSConfig.DiscardUnknown(m)
}

var xxx_messageInfo_SDSConfig proto.InternalMessageInfo

func (m *SDSConfig) GetTargetUri() string {
	if m != nil {
		return m.TargetUri
	}
	return ""
}

func (m *SDSConfig) GetCallCredentials() *CallCredentials {
	if m != nil {
		return m.CallCredentials
	}
	return nil
}

func (m *SDSConfig) GetCertificatesSecretName() string {
	if m != nil {
		return m.CertificatesSecretName
	}
	return ""
}

func (m *SDSConfig) GetValidationContextName() string {
	if m != nil {
		return m.ValidationContextName
	}
	return ""
}

type CallCredentials struct {
	FileCredentialSource *CallCredentials_FileCredentialSource `protobuf:"bytes,1,opt,name=file_credential_source,json=fileCredentialSource,proto3" json:"file_credential_source,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *CallCredentials) Reset()         { *m = CallCredentials{} }
func (m *CallCredentials) String() string { return proto.CompactTextString(m) }
func (*CallCredentials) ProtoMessage()    {}
func (*CallCredentials) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4a65e8067d81add, []int{4}
}
func (m *CallCredentials) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallCredentials.Unmarshal(m, b)
}
func (m *CallCredentials) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallCredentials.Marshal(b, m, deterministic)
}
func (m *CallCredentials) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallCredentials.Merge(m, src)
}
func (m *CallCredentials) XXX_Size() int {
	return xxx_messageInfo_CallCredentials.Size(m)
}
func (m *CallCredentials) XXX_DiscardUnknown() {
	xxx_messageInfo_CallCredentials.DiscardUnknown(m)
}

var xxx_messageInfo_CallCredentials proto.InternalMessageInfo

func (m *CallCredentials) GetFileCredentialSource() *CallCredentials_FileCredentialSource {
	if m != nil {
		return m.FileCredentialSource
	}
	return nil
}

type CallCredentials_FileCredentialSource struct {
	TokenFileName        string   `protobuf:"bytes,1,opt,name=TokenFileName,proto3" json:"TokenFileName,omitempty"`
	Header               string   `protobuf:"bytes,2,opt,name=Header,proto3" json:"Header,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CallCredentials_FileCredentialSource) Reset()         { *m = CallCredentials_FileCredentialSource{} }
func (m *CallCredentials_FileCredentialSource) String() string { return proto.CompactTextString(m) }
func (*CallCredentials_FileCredentialSource) ProtoMessage()    {}
func (*CallCredentials_FileCredentialSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4a65e8067d81add, []int{4, 0}
}
func (m *CallCredentials_FileCredentialSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallCredentials_FileCredentialSource.Unmarshal(m, b)
}
func (m *CallCredentials_FileCredentialSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallCredentials_FileCredentialSource.Marshal(b, m, deterministic)
}
func (m *CallCredentials_FileCredentialSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallCredentials_FileCredentialSource.Merge(m, src)
}
func (m *CallCredentials_FileCredentialSource) XXX_Size() int {
	return xxx_messageInfo_CallCredentials_FileCredentialSource.Size(m)
}
func (m *CallCredentials_FileCredentialSource) XXX_DiscardUnknown() {
	xxx_messageInfo_CallCredentials_FileCredentialSource.DiscardUnknown(m)
}

var xxx_messageInfo_CallCredentials_FileCredentialSource proto.InternalMessageInfo

func (m *CallCredentials_FileCredentialSource) GetTokenFileName() string {
	if m != nil {
		return m.TokenFileName
	}
	return ""
}

func (m *CallCredentials_FileCredentialSource) GetHeader() string {
	if m != nil {
		return m.Header
	}
	return ""
}

func init() {
	proto.RegisterType((*SslConfig)(nil), "gloo.solo.io.SslConfig")
	proto.RegisterType((*SSLFiles)(nil), "gloo.solo.io.SSLFiles")
	proto.RegisterType((*UpstreamSslConfig)(nil), "gloo.solo.io.UpstreamSslConfig")
	proto.RegisterType((*SDSConfig)(nil), "gloo.solo.io.SDSConfig")
	proto.RegisterType((*CallCredentials)(nil), "gloo.solo.io.CallCredentials")
	proto.RegisterType((*CallCredentials_FileCredentialSource)(nil), "gloo.solo.io.CallCredentials.FileCredentialSource")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/ssl.proto", fileDescriptor_c4a65e8067d81add)
}

var fileDescriptor_c4a65e8067d81add = []byte{
	// 581 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x54, 0xcf, 0x4e, 0x14, 0x4f,
	0x10, 0x66, 0x59, 0x7e, 0xfc, 0x98, 0x42, 0x02, 0x76, 0xd6, 0xdd, 0x81, 0xf8, 0x87, 0x6c, 0x3c,
	0x90, 0xa8, 0xb3, 0x8a, 0x81, 0x18, 0x6f, 0x32, 0xc4, 0x6c, 0xa2, 0xf1, 0x30, 0x03, 0x31, 0xf1,
	0xd2, 0x69, 0x7a, 0x6b, 0x86, 0x76, 0x7b, 0xa7, 0x37, 0x5d, 0x0d, 0x91, 0x67, 0xf0, 0xe6, 0x53,
	0xf8, 0x3a, 0xbe, 0x82, 0x89, 0xef, 0x61, 0x7a, 0x7a, 0xd6, 0x05, 0x42, 0x8c, 0x07, 0x2f, 0x9e,
	0xa6, 0xaa, 0xbe, 0xfa, 0xa6, 0xbf, 0xaf, 0xaa, 0xd3, 0xb0, 0x5f, 0x2a, 0x77, 0x7a, 0x76, 0x92,
	0x48, 0x33, 0x19, 0x90, 0xd1, 0xe6, 0x89, 0x32, 0x83, 0x52, 0x1b, 0x33, 0x98, 0x5a, 0xf3, 0x11,
	0xa5, 0xa3, 0x90, 0x89, 0xa9, 0x1a, 0x9c, 0x3f, 0x1b, 0x10, 0xe9, 0x64, 0x6a, 0x8d, 0x33, 0xec,
	0x96, 0x2f, 0x27, 0x9e, 0x91, 0x28, 0xb3, 0xd5, 0x29, 0x4d, 0x69, 0x6a, 0x60, 0xe0, 0xa3, 0xd0,
	0xb3, 0xf5, 0xf8, 0x86, 0x7f, 0xd7, 0xdf, 0xb1, 0x72, 0xb3, 0x3f, 0x5a, 0x2c, 0x42, 0x77, 0xff,
	0xcb, 0x22, 0x44, 0x39, 0xe9, 0xd4, 0x54, 0x85, 0x2a, 0xd9, 0x4b, 0x00, 0x42, 0x69, 0xd1, 0x71,
	0x8b, 0x45, 0xdc, 0xda, 0x6e, 0xed, 0xac, 0xee, 0x6e, 0x26, 0xd2, 0x58, 0x9c, 0x1d, 0x9a, 0x64,
	0x48, 0xe6, 0xcc, 0x4a, 0xcc, 0xb0, 0x18, 0x2e, 0x64, 0x51, 0x68, 0xcf, 0xb0, 0x60, 0x7b, 0x10,
	0x11, 0x69, 0x5e, 0x28, 0x8d, 0x14, 0x2f, 0xd6, 0xd4, 0x6e, 0x72, 0x59, 0x6f, 0x92, 0xe7, 0x6f,
	0x5f, 0x7b, 0x74, 0xb8, 0x90, 0xad, 0x10, 0xe9, 0x3a, 0x66, 0x8f, 0xa0, 0x4d, 0x23, 0x8a, 0x97,
	0x6a, 0x42, 0xef, 0x1a, 0xe1, 0x30, 0x0f, 0xc2, 0x86, 0x0b, 0x99, 0xef, 0x62, 0x0f, 0x60, 0x95,
	0x2a, 0xc5, 0x47, 0x66, 0x22, 0x54, 0x45, 0x71, 0x7b, 0xbb, 0xbd, 0x13, 0x65, 0x40, 0x95, 0x3a,
	0x0c, 0x15, 0xb6, 0x07, 0xbd, 0x73, 0xb4, 0xaa, 0xb8, 0xe0, 0x74, 0x76, 0xe2, 0x27, 0xc9, 0x85,
	0x76, 0xbc, 0x12, 0x13, 0x8c, 0xff, 0xab, 0x9b, 0x3b, 0x01, 0xce, 0x03, 0xfa, 0x4a, 0xbb, 0x77,
	0x62, 0x82, 0x07, 0x6b, 0xb0, 0xea, 0xb5, 0x07, 0x33, 0xd4, 0x7f, 0x0f, 0x2b, 0x33, 0xad, 0x6c,
	0x13, 0x56, 0x9c, 0x26, 0x2e, 0xd1, 0xba, 0x7a, 0x20, 0x51, 0xf6, 0xbf, 0xd3, 0x94, 0xa2, 0x75,
	0xac, 0x07, 0x3e, 0xe4, 0x63, 0xbc, 0xa8, 0xfd, 0x46, 0xd9, 0xb2, 0xd3, 0xf4, 0x06, 0x2f, 0x3c,
	0x60, 0x8d, 0x71, 0x5c, 0x8a, 0xb8, 0x1d, 0x00, 0x9f, 0xa6, 0xa2, 0xff, 0x79, 0x11, 0x6e, 0x1f,
	0x4f, 0xc9, 0x59, 0x14, 0x93, 0x7f, 0x67, 0xea, 0x1b, 0xd0, 0xa6, 0x4a, 0x35, 0x56, 0x7c, 0xf8,
	0x97, 0xc6, 0xfc, 0xa3, 0x05, 0xd1, 0xaf, 0xc3, 0xd8, 0x5d, 0x88, 0x8e, 0x84, 0x2d, 0xd1, 0x1d,
	0x5b, 0xd5, 0x4c, 0x7a, 0x5e, 0x60, 0x43, 0xd8, 0x90, 0x42, 0x6b, 0x2e, 0x2d, 0x8e, 0xb0, 0x72,
	0x4a, 0xe8, 0x99, 0xdd, 0x7b, 0x57, 0xd5, 0xa7, 0x42, 0xeb, 0x74, 0xde, 0x94, 0xad, 0xcb, 0xab,
	0x05, 0xf6, 0x02, 0x62, 0xbf, 0x4c, 0x55, 0x28, 0x29, 0x1c, 0x52, 0xa3, 0x26, 0x88, 0x0f, 0x16,
	0xbb, 0x97, 0xf1, 0xbc, 0x86, 0xbd, 0x7c, 0xb6, 0x0f, 0xbd, 0x73, 0xa1, 0xd5, 0x48, 0x38, 0x65,
	0x2a, 0x2e, 0x4d, 0xe5, 0xf0, 0x53, 0x43, 0x5c, 0xaa, 0x89, 0x77, 0xe6, 0x70, 0x1a, 0x50, 0xcf,
	0xeb, 0x7f, 0x6b, 0xc1, 0xfa, 0x35, 0x59, 0xec, 0x14, 0xba, 0x7e, 0x67, 0x97, 0xfc, 0xf0, 0xb0,
	0xe1, 0x66, 0xff, 0xbb, 0xbf, 0x75, 0x95, 0xf8, 0x2d, 0xce, 0xf3, 0x3c, 0xdc, 0x8d, 0x4e, 0x71,
	0x43, 0x75, 0xeb, 0x08, 0x3a, 0x37, 0x75, 0xb3, 0x87, 0xb0, 0x76, 0x64, 0xc6, 0x58, 0x79, 0xd0,
	0xcb, 0x6c, 0x66, 0x7e, 0xb5, 0xc8, 0xba, 0xb0, 0x3c, 0x44, 0x31, 0x42, 0x3b, 0xbb, 0xe2, 0x21,
	0x3b, 0xd8, 0xff, 0xfa, 0xfd, 0x7e, 0xeb, 0xc3, 0xd3, 0x3f, 0x7b, 0xc7, 0xa6, 0xe3, 0xb2, 0x79,
	0x79, 0x4e, 0x96, 0xeb, 0x67, 0xe7, 0xf9, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x16, 0x3b, 0x6c,
	0xc8, 0x02, 0x05, 0x00, 0x00,
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
	if len(this.VerifySubjectAltName) != len(that1.VerifySubjectAltName) {
		return false
	}
	for i := range this.VerifySubjectAltName {
		if this.VerifySubjectAltName[i] != that1.VerifySubjectAltName[i] {
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
func (this *SslConfig_Sds) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SslConfig_Sds)
	if !ok {
		that2, ok := that.(SslConfig_Sds)
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
	if !this.Sds.Equal(that1.Sds) {
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
	if len(this.VerifySubjectAltName) != len(that1.VerifySubjectAltName) {
		return false
	}
	for i := range this.VerifySubjectAltName {
		if this.VerifySubjectAltName[i] != that1.VerifySubjectAltName[i] {
			return false
		}
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
func (this *UpstreamSslConfig_Sds) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpstreamSslConfig_Sds)
	if !ok {
		that2, ok := that.(UpstreamSslConfig_Sds)
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
	if !this.Sds.Equal(that1.Sds) {
		return false
	}
	return true
}
func (this *SDSConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SDSConfig)
	if !ok {
		that2, ok := that.(SDSConfig)
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
	if this.TargetUri != that1.TargetUri {
		return false
	}
	if !this.CallCredentials.Equal(that1.CallCredentials) {
		return false
	}
	if this.CertificatesSecretName != that1.CertificatesSecretName {
		return false
	}
	if this.ValidationContextName != that1.ValidationContextName {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *CallCredentials) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CallCredentials)
	if !ok {
		that2, ok := that.(CallCredentials)
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
	if !this.FileCredentialSource.Equal(that1.FileCredentialSource) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *CallCredentials_FileCredentialSource) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CallCredentials_FileCredentialSource)
	if !ok {
		that2, ok := that.(CallCredentials_FileCredentialSource)
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
	if this.TokenFileName != that1.TokenFileName {
		return false
	}
	if this.Header != that1.Header {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
