// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/service/auth.proto

package service

import (
	context "context"
	fmt "fmt"
	_ "github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/admin"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type OAuth2MetadataRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OAuth2MetadataRequest) Reset()         { *m = OAuth2MetadataRequest{} }
func (m *OAuth2MetadataRequest) String() string { return proto.CompactTextString(m) }
func (*OAuth2MetadataRequest) ProtoMessage()    {}
func (*OAuth2MetadataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6eee4a0c193ab842, []int{0}
}

func (m *OAuth2MetadataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OAuth2MetadataRequest.Unmarshal(m, b)
}
func (m *OAuth2MetadataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OAuth2MetadataRequest.Marshal(b, m, deterministic)
}
func (m *OAuth2MetadataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OAuth2MetadataRequest.Merge(m, src)
}
func (m *OAuth2MetadataRequest) XXX_Size() int {
	return xxx_messageInfo_OAuth2MetadataRequest.Size(m)
}
func (m *OAuth2MetadataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OAuth2MetadataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OAuth2MetadataRequest proto.InternalMessageInfo

// OAuth2MetadataResponse defines an RFC-Compliant response for /.well-known/oauth-authorization-server metadata
// as defined in https://tools.ietf.org/html/rfc8414
type OAuth2MetadataResponse struct {
	// Defines the issuer string in all JWT tokens this server issues. The issuer can be admin itself or an external
	// issuer.
	Issuer string `protobuf:"bytes,1,opt,name=issuer,proto3" json:"issuer,omitempty"`
	// URL of the authorization server's authorization endpoint [RFC6749]. This is REQUIRED unless no grant types are
	// supported that use the authorization endpoint.
	AuthorizationEndpoint string `protobuf:"bytes,2,opt,name=authorization_endpoint,json=authorizationEndpoint,proto3" json:"authorization_endpoint,omitempty"`
	// URL of the authorization server's token endpoint [RFC6749].
	TokenEndpoint string `protobuf:"bytes,3,opt,name=token_endpoint,json=tokenEndpoint,proto3" json:"token_endpoint,omitempty"`
	// Array containing a list of the OAuth 2.0 "response_type" values that this authorization server supports.
	ResponseTypesSupported []string `protobuf:"bytes,4,rep,name=response_types_supported,json=responseTypesSupported,proto3" json:"response_types_supported,omitempty"`
	// JSON array containing a list of the OAuth 2.0 [RFC6749] "scope" values that this authorization server supports.
	ScopesSupported []string `protobuf:"bytes,5,rep,name=scopes_supported,json=scopesSupported,proto3" json:"scopes_supported,omitempty"`
	// JSON array containing a list of client authentication methods supported by this token endpoint.
	TokenEndpointAuthMethodsSupported []string `protobuf:"bytes,6,rep,name=token_endpoint_auth_methods_supported,json=tokenEndpointAuthMethodsSupported,proto3" json:"token_endpoint_auth_methods_supported,omitempty"`
	// URL of the authorization server's JWK Set [JWK] document. The referenced document contains the signing key(s) the
	// client uses to validate signatures from the authorization server.
	JwksUri string `protobuf:"bytes,7,opt,name=jwks_uri,json=jwksUri,proto3" json:"jwks_uri,omitempty"`
	// JSON array containing a list of Proof Key for Code Exchange (PKCE) [RFC7636] code challenge methods supported by
	// this authorization server.
	CodeChallengeMethodsSupported []string `protobuf:"bytes,8,rep,name=code_challenge_methods_supported,json=codeChallengeMethodsSupported,proto3" json:"code_challenge_methods_supported,omitempty"`
	// JSON array containing a list of the OAuth 2.0 grant type values that this authorization server supports.
	GrantTypesSupported  []string `protobuf:"bytes,9,rep,name=grant_types_supported,json=grantTypesSupported,proto3" json:"grant_types_supported,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OAuth2MetadataResponse) Reset()         { *m = OAuth2MetadataResponse{} }
func (m *OAuth2MetadataResponse) String() string { return proto.CompactTextString(m) }
func (*OAuth2MetadataResponse) ProtoMessage()    {}
func (*OAuth2MetadataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6eee4a0c193ab842, []int{1}
}

func (m *OAuth2MetadataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OAuth2MetadataResponse.Unmarshal(m, b)
}
func (m *OAuth2MetadataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OAuth2MetadataResponse.Marshal(b, m, deterministic)
}
func (m *OAuth2MetadataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OAuth2MetadataResponse.Merge(m, src)
}
func (m *OAuth2MetadataResponse) XXX_Size() int {
	return xxx_messageInfo_OAuth2MetadataResponse.Size(m)
}
func (m *OAuth2MetadataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OAuth2MetadataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OAuth2MetadataResponse proto.InternalMessageInfo

func (m *OAuth2MetadataResponse) GetIssuer() string {
	if m != nil {
		return m.Issuer
	}
	return ""
}

func (m *OAuth2MetadataResponse) GetAuthorizationEndpoint() string {
	if m != nil {
		return m.AuthorizationEndpoint
	}
	return ""
}

func (m *OAuth2MetadataResponse) GetTokenEndpoint() string {
	if m != nil {
		return m.TokenEndpoint
	}
	return ""
}

func (m *OAuth2MetadataResponse) GetResponseTypesSupported() []string {
	if m != nil {
		return m.ResponseTypesSupported
	}
	return nil
}

func (m *OAuth2MetadataResponse) GetScopesSupported() []string {
	if m != nil {
		return m.ScopesSupported
	}
	return nil
}

func (m *OAuth2MetadataResponse) GetTokenEndpointAuthMethodsSupported() []string {
	if m != nil {
		return m.TokenEndpointAuthMethodsSupported
	}
	return nil
}

func (m *OAuth2MetadataResponse) GetJwksUri() string {
	if m != nil {
		return m.JwksUri
	}
	return ""
}

func (m *OAuth2MetadataResponse) GetCodeChallengeMethodsSupported() []string {
	if m != nil {
		return m.CodeChallengeMethodsSupported
	}
	return nil
}

func (m *OAuth2MetadataResponse) GetGrantTypesSupported() []string {
	if m != nil {
		return m.GrantTypesSupported
	}
	return nil
}

type PublicClientAuthConfigRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublicClientAuthConfigRequest) Reset()         { *m = PublicClientAuthConfigRequest{} }
func (m *PublicClientAuthConfigRequest) String() string { return proto.CompactTextString(m) }
func (*PublicClientAuthConfigRequest) ProtoMessage()    {}
func (*PublicClientAuthConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6eee4a0c193ab842, []int{2}
}

func (m *PublicClientAuthConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublicClientAuthConfigRequest.Unmarshal(m, b)
}
func (m *PublicClientAuthConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublicClientAuthConfigRequest.Marshal(b, m, deterministic)
}
func (m *PublicClientAuthConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublicClientAuthConfigRequest.Merge(m, src)
}
func (m *PublicClientAuthConfigRequest) XXX_Size() int {
	return xxx_messageInfo_PublicClientAuthConfigRequest.Size(m)
}
func (m *PublicClientAuthConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PublicClientAuthConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PublicClientAuthConfigRequest proto.InternalMessageInfo

// FlyteClientResponse encapsulates public information that flyte clients (CLIs... etc.) can use to authenticate users.
type PublicClientAuthConfigResponse struct {
	// client_id to use when initiating OAuth2 authorization requests.
	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	// redirect uri to use when initiating OAuth2 authorization requests.
	RedirectUri string `protobuf:"bytes,2,opt,name=redirect_uri,json=redirectUri,proto3" json:"redirect_uri,omitempty"`
	// scopes to request when initiating OAuth2 authorization requests.
	Scopes []string `protobuf:"bytes,3,rep,name=scopes,proto3" json:"scopes,omitempty"`
	// Authorization Header to use when passing Access Tokens to the server. If not provided, the client should use the
	// default http `Authorization` header.
	AuthorizationMetadataKey string   `protobuf:"bytes,4,opt,name=authorization_metadata_key,json=authorizationMetadataKey,proto3" json:"authorization_metadata_key,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *PublicClientAuthConfigResponse) Reset()         { *m = PublicClientAuthConfigResponse{} }
func (m *PublicClientAuthConfigResponse) String() string { return proto.CompactTextString(m) }
func (*PublicClientAuthConfigResponse) ProtoMessage()    {}
func (*PublicClientAuthConfigResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6eee4a0c193ab842, []int{3}
}

func (m *PublicClientAuthConfigResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublicClientAuthConfigResponse.Unmarshal(m, b)
}
func (m *PublicClientAuthConfigResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublicClientAuthConfigResponse.Marshal(b, m, deterministic)
}
func (m *PublicClientAuthConfigResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublicClientAuthConfigResponse.Merge(m, src)
}
func (m *PublicClientAuthConfigResponse) XXX_Size() int {
	return xxx_messageInfo_PublicClientAuthConfigResponse.Size(m)
}
func (m *PublicClientAuthConfigResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PublicClientAuthConfigResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PublicClientAuthConfigResponse proto.InternalMessageInfo

func (m *PublicClientAuthConfigResponse) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *PublicClientAuthConfigResponse) GetRedirectUri() string {
	if m != nil {
		return m.RedirectUri
	}
	return ""
}

func (m *PublicClientAuthConfigResponse) GetScopes() []string {
	if m != nil {
		return m.Scopes
	}
	return nil
}

func (m *PublicClientAuthConfigResponse) GetAuthorizationMetadataKey() string {
	if m != nil {
		return m.AuthorizationMetadataKey
	}
	return ""
}

func init() {
	proto.RegisterType((*OAuth2MetadataRequest)(nil), "flyteidl.service.OAuth2MetadataRequest")
	proto.RegisterType((*OAuth2MetadataResponse)(nil), "flyteidl.service.OAuth2MetadataResponse")
	proto.RegisterType((*PublicClientAuthConfigRequest)(nil), "flyteidl.service.PublicClientAuthConfigRequest")
	proto.RegisterType((*PublicClientAuthConfigResponse)(nil), "flyteidl.service.PublicClientAuthConfigResponse")
}

func init() { proto.RegisterFile("flyteidl/service/auth.proto", fileDescriptor_6eee4a0c193ab842) }

var fileDescriptor_6eee4a0c193ab842 = []byte{
	// 805 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x95, 0xcf, 0x8f, 0xdb, 0x44,
	0x14, 0xc7, 0x95, 0xa6, 0xa4, 0x9b, 0xe1, 0x57, 0x71, 0x95, 0xd4, 0xeb, 0xed, 0x16, 0x37, 0x55,
	0xb5, 0xa9, 0x44, 0x6c, 0x58, 0x84, 0x00, 0x89, 0xcb, 0xb2, 0x42, 0x2b, 0x04, 0x85, 0x2a, 0x2d,
	0x12, 0xea, 0xc5, 0x9a, 0x8c, 0x5f, 0xec, 0x69, 0xec, 0x19, 0x33, 0x33, 0x4e, 0x08, 0x47, 0xce,
	0x9c, 0x96, 0xff, 0x85, 0x7f, 0x84, 0x33, 0x37, 0x8e, 0x5c, 0xb9, 0x71, 0x40, 0x9e, 0x19, 0x27,
	0x6b, 0x27, 0xcb, 0x8f, 0xd3, 0xca, 0xf3, 0xf9, 0xce, 0x7b, 0x6f, 0xde, 0xfb, 0xee, 0x0b, 0x3a,
	0x9a, 0x67, 0x6b, 0x05, 0x34, 0xce, 0x42, 0x09, 0x62, 0x49, 0x09, 0x84, 0xb8, 0x54, 0x69, 0x50,
	0x08, 0xae, 0xb8, 0x73, 0xbb, 0x86, 0x81, 0x85, 0xde, 0xbd, 0x84, 0xf3, 0x24, 0x83, 0x10, 0x17,
	0x34, 0xc4, 0x8c, 0x71, 0x85, 0x15, 0xe5, 0x4c, 0x1a, 0xbd, 0x77, 0x6f, 0x13, 0x0c, 0xc7, 0x39,
	0x65, 0x61, 0x21, 0xf8, 0x4b, 0x20, 0xca, 0xd2, 0x60, 0x3f, 0x8d, 0x62, 0x9e, 0x63, 0xca, 0x22,
	0xac, 0x94, 0xa0, 0xb3, 0x52, 0x41, 0x1d, 0xed, 0xb0, 0xa5, 0x57, 0x58, 0x2e, 0x2c, 0x3a, 0x6e,
	0xa1, 0x15, 0x17, 0x8b, 0x79, 0xc6, 0x57, 0x16, 0x8f, 0xaf, 0xc1, 0xbb, 0x39, 0xfc, 0x96, 0x32,
	0xc3, 0x25, 0x23, 0x69, 0x54, 0x64, 0x98, 0x59, 0x85, 0xd7, 0x52, 0xc0, 0x12, 0x58, 0xfd, 0xa2,
	0xfb, 0x6d, 0xf6, 0x3d, 0x90, 0xb2, 0x6a, 0x88, 0xe5, 0x27, 0x2d, 0x9e, 0x63, 0x45, 0x52, 0x3c,
	0xcb, 0x20, 0x12, 0x20, 0x79, 0x29, 0x08, 0x58, 0xe1, 0xc3, 0x96, 0x90, 0xf1, 0x18, 0xa2, 0x76,
	0xb4, 0x87, 0x7b, 0xfa, 0xb1, 0x23, 0x6a, 0x8f, 0x60, 0x09, 0x42, 0x6e, 0xe9, 0x51, 0x8b, 0x12,
	0x9e, 0xe7, 0x1b, 0xf8, 0x8e, 0xfe, 0x43, 0x26, 0x09, 0xb0, 0x89, 0x5c, 0xe1, 0x24, 0x01, 0x11,
	0xf2, 0x42, 0xcf, 0x77, 0x77, 0xd6, 0xa3, 0xbb, 0x68, 0xf0, 0xf5, 0x59, 0xa9, 0xd2, 0xd3, 0x27,
	0xa0, 0x70, 0x8c, 0x15, 0x9e, 0xc2, 0x77, 0x25, 0x48, 0x35, 0xfa, 0xad, 0x8b, 0x86, 0x6d, 0x22,
	0x0b, 0xce, 0x24, 0x38, 0x43, 0xd4, 0xa3, 0x52, 0x96, 0x20, 0xdc, 0x8e, 0xdf, 0x19, 0xf7, 0xa7,
	0xf6, 0xcb, 0xf9, 0x00, 0x0d, 0x2b, 0xd7, 0x71, 0x41, 0x7f, 0xd0, 0x39, 0x22, 0x60, 0x71, 0xc1,
	0x29, 0x53, 0xee, 0x0d, 0xad, 0x1b, 0x34, 0xe8, 0x67, 0x16, 0x3a, 0x8f, 0xd0, 0x1b, 0x8a, 0x2f,
	0xe0, 0x8a, 0xbc, 0xab, 0xe5, 0xaf, 0xeb, 0xd3, 0x8d, 0xec, 0x23, 0xe4, 0x0a, 0x5b, 0x41, 0xa4,
	0xd6, 0x05, 0xc8, 0x48, 0x96, 0x45, 0xc1, 0x85, 0x82, 0xd8, 0xbd, 0xe9, 0x77, 0xc7, 0xfd, 0xe9,
	0xb0, 0xe6, 0xcf, 0x2b, 0xfc, 0xac, 0xa6, 0xce, 0x63, 0x74, 0x5b, 0x12, 0xde, 0xbc, 0xf1, 0x8a,
	0xbe, 0xf1, 0xa6, 0x39, 0xdf, 0x4a, 0x9f, 0xa2, 0x47, 0xcd, 0x5a, 0xa2, 0xaa, 0xe6, 0x28, 0x07,
	0x95, 0xf2, 0xf8, 0xea, 0xfd, 0x9e, 0xbe, 0xff, 0xa0, 0x51, 0x62, 0xd5, 0xad, 0x27, 0x46, 0xb9,
	0x8d, 0x78, 0x88, 0x0e, 0x5e, 0xae, 0x16, 0x32, 0x2a, 0x05, 0x75, 0x6f, 0xe9, 0x77, 0xdd, 0xaa,
	0xbe, 0xbf, 0x11, 0xd4, 0xb9, 0x40, 0x3e, 0xa9, 0x1c, 0x42, 0x52, 0x9c, 0x65, 0xc0, 0x12, 0xd8,
	0x93, 0xe7, 0x40, 0xe7, 0x39, 0xae, 0x74, 0xe7, 0xb5, 0x6c, 0x27, 0xc7, 0x29, 0x1a, 0x24, 0x02,
	0x33, 0xb5, 0xd3, 0x97, 0xbe, 0xbe, 0x7d, 0x47, 0xc3, 0x66, 0x53, 0x46, 0x6f, 0xa3, 0xe3, 0xa7,
	0xe5, 0x2c, 0xa3, 0xe4, 0x3c, 0xa3, 0x60, 0x6a, 0x3f, 0xe7, 0x6c, 0x4e, 0x93, 0xda, 0x00, 0xbf,
	0x74, 0xd0, 0xfd, 0xeb, 0x14, 0xd6, 0x08, 0x47, 0xa8, 0x4f, 0x34, 0x8b, 0x68, 0x6c, 0xbd, 0x70,
	0x60, 0x0e, 0x3e, 0x8f, 0x9d, 0x07, 0xe8, 0x35, 0x01, 0x31, 0x15, 0xd5, 0x6e, 0xa8, 0x1e, 0x6f,
	0x3c, 0xf0, 0x6a, 0x7d, 0x56, 0x35, 0x60, 0x88, 0x7a, 0x66, 0x00, 0x6e, 0x57, 0x17, 0x6a, 0xbf,
	0x9c, 0x4f, 0x90, 0xd7, 0x34, 0x52, 0x6e, 0x2d, 0x18, 0x2d, 0x60, 0xed, 0xde, 0xd4, 0x81, 0xdc,
	0x86, 0xa2, 0xf6, 0xe8, 0x17, 0xb0, 0x3e, 0xfd, 0xab, 0x8b, 0xee, 0xd8, 0x51, 0xe8, 0xb3, 0x67,
	0x66, 0xe9, 0x39, 0x7f, 0x76, 0xd0, 0x5b, 0x17, 0xa0, 0x9a, 0xa6, 0x76, 0x4e, 0x82, 0xf6, 0x76,
	0x0c, 0xf6, 0xfe, 0x43, 0x78, 0xe3, 0x7f, 0x17, 0x9a, 0xb6, 0x8c, 0x7e, 0xea, 0x5c, 0x9e, 0xbd,
	0xf0, 0xbe, 0x9d, 0x82, 0x12, 0x14, 0x96, 0x20, 0x7d, 0xa3, 0xf3, 0x1b, 0x15, 0xfb, 0x55, 0x10,
	0x10, 0x7e, 0xfd, 0xb4, 0xc0, 0x7f, 0x9e, 0x52, 0xe9, 0xd7, 0xbe, 0xf3, 0xa9, 0xf4, 0x31, 0xe3,
	0x6c, 0x9d, 0xf3, 0x52, 0x66, 0x6b, 0x1f, 0x13, 0x02, 0x52, 0xd2, 0x59, 0x06, 0xc1, 0x8f, 0xbf,
	0xfe, 0xfe, 0xf3, 0x8d, 0xc7, 0xce, 0x49, 0x18, 0xac, 0x20, 0xcb, 0x26, 0x0b, 0xc6, 0x57, 0x2c,
	0xe4, 0x55, 0xf0, 0x49, 0x23, 0xc3, 0xc4, 0x64, 0x70, 0xfe, 0xe8, 0xa0, 0xc1, 0x05, 0xa8, 0xab,
	0xb3, 0x34, 0x73, 0x74, 0xc2, 0xdd, 0x27, 0xfd, 0xa3, 0x27, 0xbc, 0x77, 0xff, 0xfb, 0x05, 0xdb,
	0x8b, 0xe2, 0xf2, 0xec, 0x2b, 0xef, 0xcb, 0x6d, 0x2b, 0x0a, 0x2d, 0x37, 0xfb, 0xda, 0x37, 0x5e,
	0xf1, 0x29, 0x9b, 0xf3, 0xff, 0xf9, 0xfc, 0x43, 0xe7, 0x6e, 0x48, 0x74, 0xa6, 0x70, 0xf9, 0x5e,
	0xa8, 0xa3, 0x45, 0x26, 0xda, 0xa7, 0x1f, 0xbf, 0xf8, 0x30, 0xa1, 0x2a, 0x2d, 0x67, 0x01, 0xe1,
	0xb9, 0x41, 0x5c, 0x24, 0xe1, 0x66, 0x65, 0x26, 0xc0, 0xc2, 0x62, 0x36, 0x49, 0x78, 0xd8, 0xfe,
	0xcd, 0x9c, 0xf5, 0xf4, 0x4e, 0x7c, 0xff, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x77, 0xc8, 0x78,
	0xb9, 0x4e, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthMetadataServiceClient is the client API for AuthMetadataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthMetadataServiceClient interface {
	// Anonymously accessible. Retrieves local or external oauth authorization server metadata.
	GetOAuth2Metadata(ctx context.Context, in *OAuth2MetadataRequest, opts ...grpc.CallOption) (*OAuth2MetadataResponse, error)
	// Anonymously accessible. Retrieves the client information clients should use when initiating OAuth2 authorization
	// requests.
	GetPublicClientConfig(ctx context.Context, in *PublicClientAuthConfigRequest, opts ...grpc.CallOption) (*PublicClientAuthConfigResponse, error)
}

type authMetadataServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthMetadataServiceClient(cc *grpc.ClientConn) AuthMetadataServiceClient {
	return &authMetadataServiceClient{cc}
}

func (c *authMetadataServiceClient) GetOAuth2Metadata(ctx context.Context, in *OAuth2MetadataRequest, opts ...grpc.CallOption) (*OAuth2MetadataResponse, error) {
	out := new(OAuth2MetadataResponse)
	err := c.cc.Invoke(ctx, "/flyteidl.service.AuthMetadataService/GetOAuth2Metadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authMetadataServiceClient) GetPublicClientConfig(ctx context.Context, in *PublicClientAuthConfigRequest, opts ...grpc.CallOption) (*PublicClientAuthConfigResponse, error) {
	out := new(PublicClientAuthConfigResponse)
	err := c.cc.Invoke(ctx, "/flyteidl.service.AuthMetadataService/GetPublicClientConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthMetadataServiceServer is the server API for AuthMetadataService service.
type AuthMetadataServiceServer interface {
	// Anonymously accessible. Retrieves local or external oauth authorization server metadata.
	GetOAuth2Metadata(context.Context, *OAuth2MetadataRequest) (*OAuth2MetadataResponse, error)
	// Anonymously accessible. Retrieves the client information clients should use when initiating OAuth2 authorization
	// requests.
	GetPublicClientConfig(context.Context, *PublicClientAuthConfigRequest) (*PublicClientAuthConfigResponse, error)
}

// UnimplementedAuthMetadataServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAuthMetadataServiceServer struct {
}

func (*UnimplementedAuthMetadataServiceServer) GetOAuth2Metadata(ctx context.Context, req *OAuth2MetadataRequest) (*OAuth2MetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOAuth2Metadata not implemented")
}
func (*UnimplementedAuthMetadataServiceServer) GetPublicClientConfig(ctx context.Context, req *PublicClientAuthConfigRequest) (*PublicClientAuthConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPublicClientConfig not implemented")
}

func RegisterAuthMetadataServiceServer(s *grpc.Server, srv AuthMetadataServiceServer) {
	s.RegisterService(&_AuthMetadataService_serviceDesc, srv)
}

func _AuthMetadataService_GetOAuth2Metadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OAuth2MetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthMetadataServiceServer).GetOAuth2Metadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flyteidl.service.AuthMetadataService/GetOAuth2Metadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthMetadataServiceServer).GetOAuth2Metadata(ctx, req.(*OAuth2MetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthMetadataService_GetPublicClientConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublicClientAuthConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthMetadataServiceServer).GetPublicClientConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flyteidl.service.AuthMetadataService/GetPublicClientConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthMetadataServiceServer).GetPublicClientConfig(ctx, req.(*PublicClientAuthConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthMetadataService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "flyteidl.service.AuthMetadataService",
	HandlerType: (*AuthMetadataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOAuth2Metadata",
			Handler:    _AuthMetadataService_GetOAuth2Metadata_Handler,
		},
		{
			MethodName: "GetPublicClientConfig",
			Handler:    _AuthMetadataService_GetPublicClientConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "flyteidl/service/auth.proto",
}