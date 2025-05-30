// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: cloudstack/management/direct/download/v1/download.gen.proto

package downloadv1

import (
	_ "github.com/walteh/cloudstack-proxy/gen/proto/golang/buf/validate"
	_ "github.com/walteh/cloudstack-proxy/gen/proto/golang/cloudstack/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ListTemplateDirectDownloadCertificatesRequest represents the parameters for list the uploaded certificates for direct download templates
type ListTemplateDirectDownloadCertificatesRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// list direct download certificate by ID
	Id *int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// the zone where certificates are uploaded
	ZoneId *int64 `protobuf:"varint,2,opt,name=zone_id,json=zoneId" json:"zone_id,omitempty"`
	// if set to true: include the hosts where the certificate is uploaded to
	ListHosts *bool `protobuf:"varint,3,opt,name=list_hosts,json=listHosts" json:"list_hosts,omitempty"`
	// List by keyword
	Keyword *string `protobuf:"bytes,4,opt,name=keyword" json:"keyword,omitempty"`
	Page *int32 `protobuf:"varint,5,opt,name=page" json:"page,omitempty"`
	PageSize *int32 `protobuf:"varint,6,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	ResponseType  *string `protobuf:"bytes,7,opt,name=response_type,json=responseType" json:"response_type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListTemplateDirectDownloadCertificatesRequest) Reset() {
	*x = ListTemplateDirectDownloadCertificatesRequest{}
	mi := &file_cloudstack_management_direct_download_v1_download_gen_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListTemplateDirectDownloadCertificatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTemplateDirectDownloadCertificatesRequest) ProtoMessage() {}

func (x *ListTemplateDirectDownloadCertificatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cloudstack_management_direct_download_v1_download_gen_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTemplateDirectDownloadCertificatesRequest.ProtoReflect.Descriptor instead.
func (*ListTemplateDirectDownloadCertificatesRequest) Descriptor() ([]byte, []int) {
	return file_cloudstack_management_direct_download_v1_download_gen_proto_rawDescGZIP(), []int{0}
}

func (x *ListTemplateDirectDownloadCertificatesRequest) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *ListTemplateDirectDownloadCertificatesRequest) GetZoneId() int64 {
	if x != nil && x.ZoneId != nil {
		return *x.ZoneId
	}
	return 0
}

func (x *ListTemplateDirectDownloadCertificatesRequest) GetListHosts() bool {
	if x != nil && x.ListHosts != nil {
		return *x.ListHosts
	}
	return false
}

func (x *ListTemplateDirectDownloadCertificatesRequest) GetKeyword() string {
	if x != nil && x.Keyword != nil {
		return *x.Keyword
	}
	return ""
}

func (x *ListTemplateDirectDownloadCertificatesRequest) GetPage() int32 {
	if x != nil && x.Page != nil {
		return *x.Page
	}
	return 0
}

func (x *ListTemplateDirectDownloadCertificatesRequest) GetPageSize() int32 {
	if x != nil && x.PageSize != nil {
		return *x.PageSize
	}
	return 0
}

func (x *ListTemplateDirectDownloadCertificatesRequest) GetResponseType() string {
	if x != nil && x.ResponseType != nil {
		return *x.ResponseType
	}
	return ""
}

// ListTemplateDirectDownloadCertificatesResponse represents the response from list the uploaded certificates for direct download templates
type ListTemplateDirectDownloadCertificatesResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The list of DirectDownloadCertificates
	Items []*DirectDownloadCertificate `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	// The total count of DirectDownloadCertificates
	TotalCount    *int32 `protobuf:"varint,2,opt,name=total_count,json=totalCount" json:"total_count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListTemplateDirectDownloadCertificatesResponse) Reset() {
	*x = ListTemplateDirectDownloadCertificatesResponse{}
	mi := &file_cloudstack_management_direct_download_v1_download_gen_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListTemplateDirectDownloadCertificatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTemplateDirectDownloadCertificatesResponse) ProtoMessage() {}

func (x *ListTemplateDirectDownloadCertificatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cloudstack_management_direct_download_v1_download_gen_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTemplateDirectDownloadCertificatesResponse.ProtoReflect.Descriptor instead.
func (*ListTemplateDirectDownloadCertificatesResponse) Descriptor() ([]byte, []int) {
	return file_cloudstack_management_direct_download_v1_download_gen_proto_rawDescGZIP(), []int{1}
}

func (x *ListTemplateDirectDownloadCertificatesResponse) GetItems() []*DirectDownloadCertificate {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListTemplateDirectDownloadCertificatesResponse) GetTotalCount() int32 {
	if x != nil && x.TotalCount != nil {
		return *x.TotalCount
	}
	return 0
}

// ProvisionTemplateDirectDownloadCertificateRequest represents the parameters for provisions a host with a direct download certificate
type ProvisionTemplateDirectDownloadCertificateRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// the id of the direct download certificate to provision
	Id *int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// the host to provision the certificate
	HostId *int64 `protobuf:"varint,2,opt,name=host_id,json=hostId" json:"host_id,omitempty"`
	ResponseType  *string `protobuf:"bytes,3,opt,name=response_type,json=responseType" json:"response_type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProvisionTemplateDirectDownloadCertificateRequest) Reset() {
	*x = ProvisionTemplateDirectDownloadCertificateRequest{}
	mi := &file_cloudstack_management_direct_download_v1_download_gen_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProvisionTemplateDirectDownloadCertificateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProvisionTemplateDirectDownloadCertificateRequest) ProtoMessage() {}

func (x *ProvisionTemplateDirectDownloadCertificateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cloudstack_management_direct_download_v1_download_gen_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProvisionTemplateDirectDownloadCertificateRequest.ProtoReflect.Descriptor instead.
func (*ProvisionTemplateDirectDownloadCertificateRequest) Descriptor() ([]byte, []int) {
	return file_cloudstack_management_direct_download_v1_download_gen_proto_rawDescGZIP(), []int{2}
}

func (x *ProvisionTemplateDirectDownloadCertificateRequest) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *ProvisionTemplateDirectDownloadCertificateRequest) GetHostId() int64 {
	if x != nil && x.HostId != nil {
		return *x.HostId
	}
	return 0
}

func (x *ProvisionTemplateDirectDownloadCertificateRequest) GetResponseType() string {
	if x != nil && x.ResponseType != nil {
		return *x.ResponseType
	}
	return ""
}

// ProvisionTemplateDirectDownloadCertificateResponse represents the response from provisions a host with a direct download certificate
type ProvisionTemplateDirectDownloadCertificateResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The Result
	Result        *Result `protobuf:"bytes,1,opt,name=result" json:"result,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProvisionTemplateDirectDownloadCertificateResponse) Reset() {
	*x = ProvisionTemplateDirectDownloadCertificateResponse{}
	mi := &file_cloudstack_management_direct_download_v1_download_gen_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProvisionTemplateDirectDownloadCertificateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProvisionTemplateDirectDownloadCertificateResponse) ProtoMessage() {}

func (x *ProvisionTemplateDirectDownloadCertificateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cloudstack_management_direct_download_v1_download_gen_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProvisionTemplateDirectDownloadCertificateResponse.ProtoReflect.Descriptor instead.
func (*ProvisionTemplateDirectDownloadCertificateResponse) Descriptor() ([]byte, []int) {
	return file_cloudstack_management_direct_download_v1_download_gen_proto_rawDescGZIP(), []int{3}
}

func (x *ProvisionTemplateDirectDownloadCertificateResponse) GetResult() *Result {
	if x != nil {
		return x.Result
	}
	return nil
}

// RevokeTemplateDirectDownloadCertificateRequest represents the parameters for revoke a direct download certificate from hosts in a zone
type RevokeTemplateDirectDownloadCertificateRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// id of the certificate
	CertificateId *int64 `protobuf:"varint,1,opt,name=certificate_id,json=certificateId" json:"certificate_id,omitempty"`
	// (optional) alias of the SSL certificate
	CertificateAlias *string `protobuf:"bytes,2,opt,name=certificate_alias,json=certificateAlias" json:"certificate_alias,omitempty"`
	// (optional) hypervisor type
	Hypervisor *string `protobuf:"bytes,3,opt,name=hypervisor" json:"hypervisor,omitempty"`
	// (optional) zone to revoke certificate
	ZoneId *int64 `protobuf:"varint,4,opt,name=zone_id,json=zoneId" json:"zone_id,omitempty"`
	// (optional) the host ID to revoke certificate
	HostId *int64 `protobuf:"varint,5,opt,name=host_id,json=hostId" json:"host_id,omitempty"`
	ResponseType  *string `protobuf:"bytes,6,opt,name=response_type,json=responseType" json:"response_type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RevokeTemplateDirectDownloadCertificateRequest) Reset() {
	*x = RevokeTemplateDirectDownloadCertificateRequest{}
	mi := &file_cloudstack_management_direct_download_v1_download_gen_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RevokeTemplateDirectDownloadCertificateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevokeTemplateDirectDownloadCertificateRequest) ProtoMessage() {}

func (x *RevokeTemplateDirectDownloadCertificateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cloudstack_management_direct_download_v1_download_gen_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevokeTemplateDirectDownloadCertificateRequest.ProtoReflect.Descriptor instead.
func (*RevokeTemplateDirectDownloadCertificateRequest) Descriptor() ([]byte, []int) {
	return file_cloudstack_management_direct_download_v1_download_gen_proto_rawDescGZIP(), []int{4}
}

func (x *RevokeTemplateDirectDownloadCertificateRequest) GetCertificateId() int64 {
	if x != nil && x.CertificateId != nil {
		return *x.CertificateId
	}
	return 0
}

func (x *RevokeTemplateDirectDownloadCertificateRequest) GetCertificateAlias() string {
	if x != nil && x.CertificateAlias != nil {
		return *x.CertificateAlias
	}
	return ""
}

func (x *RevokeTemplateDirectDownloadCertificateRequest) GetHypervisor() string {
	if x != nil && x.Hypervisor != nil {
		return *x.Hypervisor
	}
	return ""
}

func (x *RevokeTemplateDirectDownloadCertificateRequest) GetZoneId() int64 {
	if x != nil && x.ZoneId != nil {
		return *x.ZoneId
	}
	return 0
}

func (x *RevokeTemplateDirectDownloadCertificateRequest) GetHostId() int64 {
	if x != nil && x.HostId != nil {
		return *x.HostId
	}
	return 0
}

func (x *RevokeTemplateDirectDownloadCertificateRequest) GetResponseType() string {
	if x != nil && x.ResponseType != nil {
		return *x.ResponseType
	}
	return ""
}

// RevokeTemplateDirectDownloadCertificateResponse represents the response from revoke a direct download certificate from hosts in a zone
type RevokeTemplateDirectDownloadCertificateResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The Result
	Result        *Result `protobuf:"bytes,1,opt,name=result" json:"result,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RevokeTemplateDirectDownloadCertificateResponse) Reset() {
	*x = RevokeTemplateDirectDownloadCertificateResponse{}
	mi := &file_cloudstack_management_direct_download_v1_download_gen_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RevokeTemplateDirectDownloadCertificateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevokeTemplateDirectDownloadCertificateResponse) ProtoMessage() {}

func (x *RevokeTemplateDirectDownloadCertificateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cloudstack_management_direct_download_v1_download_gen_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevokeTemplateDirectDownloadCertificateResponse.ProtoReflect.Descriptor instead.
func (*RevokeTemplateDirectDownloadCertificateResponse) Descriptor() ([]byte, []int) {
	return file_cloudstack_management_direct_download_v1_download_gen_proto_rawDescGZIP(), []int{5}
}

func (x *RevokeTemplateDirectDownloadCertificateResponse) GetResult() *Result {
	if x != nil {
		return x.Result
	}
	return nil
}

// UploadTemplateDirectDownloadCertificateRequest represents the parameters for upload a certificate for https direct template download on kvm hosts
type UploadTemplateDirectDownloadCertificateRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// SSL certificate
	Certificate *string `protobuf:"bytes,1,opt,name=certificate" json:"certificate,omitempty"`
	// Name for the uploaded certificate
	Name *string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// Hypervisor type
	Hypervisor *string `protobuf:"bytes,3,opt,name=hypervisor" json:"hypervisor,omitempty"`
	// Zone to upload certificate
	ZoneId *int64 `protobuf:"varint,4,opt,name=zone_id,json=zoneId" json:"zone_id,omitempty"`
	// (optional) the host ID to upload certificate
	HostId *int64 `protobuf:"varint,5,opt,name=host_id,json=hostId" json:"host_id,omitempty"`
	ResponseType  *string `protobuf:"bytes,6,opt,name=response_type,json=responseType" json:"response_type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UploadTemplateDirectDownloadCertificateRequest) Reset() {
	*x = UploadTemplateDirectDownloadCertificateRequest{}
	mi := &file_cloudstack_management_direct_download_v1_download_gen_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadTemplateDirectDownloadCertificateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadTemplateDirectDownloadCertificateRequest) ProtoMessage() {}

func (x *UploadTemplateDirectDownloadCertificateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cloudstack_management_direct_download_v1_download_gen_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadTemplateDirectDownloadCertificateRequest.ProtoReflect.Descriptor instead.
func (*UploadTemplateDirectDownloadCertificateRequest) Descriptor() ([]byte, []int) {
	return file_cloudstack_management_direct_download_v1_download_gen_proto_rawDescGZIP(), []int{6}
}

func (x *UploadTemplateDirectDownloadCertificateRequest) GetCertificate() string {
	if x != nil && x.Certificate != nil {
		return *x.Certificate
	}
	return ""
}

func (x *UploadTemplateDirectDownloadCertificateRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *UploadTemplateDirectDownloadCertificateRequest) GetHypervisor() string {
	if x != nil && x.Hypervisor != nil {
		return *x.Hypervisor
	}
	return ""
}

func (x *UploadTemplateDirectDownloadCertificateRequest) GetZoneId() int64 {
	if x != nil && x.ZoneId != nil {
		return *x.ZoneId
	}
	return 0
}

func (x *UploadTemplateDirectDownloadCertificateRequest) GetHostId() int64 {
	if x != nil && x.HostId != nil {
		return *x.HostId
	}
	return 0
}

func (x *UploadTemplateDirectDownloadCertificateRequest) GetResponseType() string {
	if x != nil && x.ResponseType != nil {
		return *x.ResponseType
	}
	return ""
}

// UploadTemplateDirectDownloadCertificateResponse represents the response from upload a certificate for https direct template download on kvm hosts
type UploadTemplateDirectDownloadCertificateResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The Result
	Result        *Result `protobuf:"bytes,1,opt,name=result" json:"result,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UploadTemplateDirectDownloadCertificateResponse) Reset() {
	*x = UploadTemplateDirectDownloadCertificateResponse{}
	mi := &file_cloudstack_management_direct_download_v1_download_gen_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadTemplateDirectDownloadCertificateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadTemplateDirectDownloadCertificateResponse) ProtoMessage() {}

func (x *UploadTemplateDirectDownloadCertificateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cloudstack_management_direct_download_v1_download_gen_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadTemplateDirectDownloadCertificateResponse.ProtoReflect.Descriptor instead.
func (*UploadTemplateDirectDownloadCertificateResponse) Descriptor() ([]byte, []int) {
	return file_cloudstack_management_direct_download_v1_download_gen_proto_rawDescGZIP(), []int{7}
}

func (x *UploadTemplateDirectDownloadCertificateResponse) GetResult() *Result {
	if x != nil {
		return x.Result
	}
	return nil
}

// DirectDownloadCertificate represents a DirectDownloadCertificate Item
type DirectDownloadCertificate struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The ID of the DirectDownloadCertificate
	Id *string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// The name of the DirectDownloadCertificate
	Name *string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// The display name of the DirectDownloadCertificate
	DisplayName *string `protobuf:"bytes,3,opt,name=display_name,json=displayName" json:"display_name,omitempty"`
	// The description of the DirectDownloadCertificate
	Description *string `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
	// The date this entity was created
	Created       *string `protobuf:"bytes,5,opt,name=created" json:"created,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DirectDownloadCertificate) Reset() {
	*x = DirectDownloadCertificate{}
	mi := &file_cloudstack_management_direct_download_v1_download_gen_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DirectDownloadCertificate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DirectDownloadCertificate) ProtoMessage() {}

func (x *DirectDownloadCertificate) ProtoReflect() protoreflect.Message {
	mi := &file_cloudstack_management_direct_download_v1_download_gen_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DirectDownloadCertificate.ProtoReflect.Descriptor instead.
func (*DirectDownloadCertificate) Descriptor() ([]byte, []int) {
	return file_cloudstack_management_direct_download_v1_download_gen_proto_rawDescGZIP(), []int{8}
}

func (x *DirectDownloadCertificate) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *DirectDownloadCertificate) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *DirectDownloadCertificate) GetDisplayName() string {
	if x != nil && x.DisplayName != nil {
		return *x.DisplayName
	}
	return ""
}

func (x *DirectDownloadCertificate) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *DirectDownloadCertificate) GetCreated() string {
	if x != nil && x.Created != nil {
		return *x.Created
	}
	return ""
}

// Success represents a Success Operation Response
type Success struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// true if operation is executed successfully
	Success *bool `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	// any text associated with the success or failure
	DisplayText   *string `protobuf:"bytes,2,opt,name=display_text,json=displayText" json:"display_text,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Success) Reset() {
	*x = Success{}
	mi := &file_cloudstack_management_direct_download_v1_download_gen_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Success) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Success) ProtoMessage() {}

func (x *Success) ProtoReflect() protoreflect.Message {
	mi := &file_cloudstack_management_direct_download_v1_download_gen_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Success.ProtoReflect.Descriptor instead.
func (*Success) Descriptor() ([]byte, []int) {
	return file_cloudstack_management_direct_download_v1_download_gen_proto_rawDescGZIP(), []int{9}
}

func (x *Success) GetSuccess() bool {
	if x != nil && x.Success != nil {
		return *x.Success
	}
	return false
}

func (x *Success) GetDisplayText() string {
	if x != nil && x.DisplayText != nil {
		return *x.DisplayText
	}
	return ""
}

// Result represents a generic operation result
type Result struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Whether the operation was successful
	Success *bool `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	// Any text associated with the success or failure
	DisplayText *string `protobuf:"bytes,2,opt,name=display_text,json=displayText" json:"display_text,omitempty"`
	// The ID of the resource affected by the operation
	Id *string `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
	// The job ID for an async operation
	JobId *string `protobuf:"bytes,4,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	// The status of the job
	JobStatus     *string `protobuf:"bytes,5,opt,name=job_status,json=jobStatus" json:"job_status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Result) Reset() {
	*x = Result{}
	mi := &file_cloudstack_management_direct_download_v1_download_gen_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_cloudstack_management_direct_download_v1_download_gen_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_cloudstack_management_direct_download_v1_download_gen_proto_rawDescGZIP(), []int{10}
}

func (x *Result) GetSuccess() bool {
	if x != nil && x.Success != nil {
		return *x.Success
	}
	return false
}

func (x *Result) GetDisplayText() string {
	if x != nil && x.DisplayText != nil {
		return *x.DisplayText
	}
	return ""
}

func (x *Result) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *Result) GetJobId() string {
	if x != nil && x.JobId != nil {
		return *x.JobId
	}
	return ""
}

func (x *Result) GetJobStatus() string {
	if x != nil && x.JobStatus != nil {
		return *x.JobStatus
	}
	return ""
}

var File_cloudstack_management_direct_download_v1_download_gen_proto protoreflect.FileDescriptor

const file_cloudstack_management_direct_download_v1_download_gen_proto_rawDesc = "" +
	"\n" +
	";cloudstack/management/direct/download/v1/download.gen.proto\x12(cloudstack.management.direct.download.v1\x1a\x1bbuf/validate/validate.proto\x1a(cloudstack/annotations/annotations.proto\x1a google/protobuf/descriptor.proto\"\xee\x01\n" +
	"-ListTemplateDirectDownloadCertificatesRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12\x17\n" +
	"\azone_id\x18\x02 \x01(\x03R\x06zoneId\x12$\n" +
	"\n" +
	"list_hosts\x18\x03 \x01(\bB\x05\xaa\x01\x02\b\x01R\tlistHosts\x12\x18\n" +
	"\akeyword\x18\x04 \x01(\tR\akeyword\x12\x12\n" +
	"\x04page\x18\x05 \x01(\x05R\x04page\x12\x1b\n" +
	"\tpage_size\x18\x06 \x01(\x05R\bpageSize\x12#\n" +
	"\rresponse_type\x18\a \x01(\tR\fresponseType\"\xb3\x01\n" +
	".ListTemplateDirectDownloadCertificatesResponse\x12Y\n" +
	"\x05items\x18\x01 \x03(\v2C.cloudstack.management.direct.download.v1.DirectDownloadCertificateR\x05items\x12\x1f\n" +
	"\vtotal_count\x18\x02 \x01(\x05R\n" +
	"totalCount:\x05\xbaH\x02\b\x00\"\x91\x01\n" +
	"1ProvisionTemplateDirectDownloadCertificateRequest\x12\x16\n" +
	"\x02id\x18\x01 \x01(\x03B\x06\xbaH\x03\xc8\x01\x01R\x02id\x12\x1f\n" +
	"\ahost_id\x18\x02 \x01(\x03B\x06\xbaH\x03\xc8\x01\x01R\x06hostId\x12#\n" +
	"\rresponse_type\x18\x03 \x01(\tR\fresponseType\"~\n" +
	"2ProvisionTemplateDirectDownloadCertificateResponse\x12H\n" +
	"\x06result\x18\x01 \x01(\v20.cloudstack.management.direct.download.v1.ResultR\x06result\"\x8f\x02\n" +
	".RevokeTemplateDirectDownloadCertificateRequest\x12%\n" +
	"\x0ecertificate_id\x18\x01 \x01(\x03R\rcertificateId\x127\n" +
	"\x11certificate_alias\x18\x02 \x01(\tB\n" +
	"\xbaH\ar\x05\x10\x01\x18\xff\x01R\x10certificateAlias\x12\x1e\n" +
	"\n" +
	"hypervisor\x18\x03 \x01(\tR\n" +
	"hypervisor\x12\x1f\n" +
	"\azone_id\x18\x04 \x01(\x03B\x06\xbaH\x03\xc8\x01\x01R\x06zoneId\x12\x17\n" +
	"\ahost_id\x18\x05 \x01(\x03R\x06hostId\x12#\n" +
	"\rresponse_type\x18\x06 \x01(\tR\fresponseType\"{\n" +
	"/RevokeTemplateDirectDownloadCertificateResponse\x12H\n" +
	"\x06result\x18\x01 \x01(\v20.cloudstack.management.direct.download.v1.ResultR\x06result\"\x84\x02\n" +
	".UploadTemplateDirectDownloadCertificateRequest\x12(\n" +
	"\vcertificate\x18\x01 \x01(\tB\x06\xbaH\x03\xc8\x01\x01R\vcertificate\x12!\n" +
	"\x04name\x18\x02 \x01(\tB\r\xbaH\n" +
	"\xc8\x01\x01r\x05\x10\x01\x18\xff\x01R\x04name\x12&\n" +
	"\n" +
	"hypervisor\x18\x03 \x01(\tB\x06\xbaH\x03\xc8\x01\x01R\n" +
	"hypervisor\x12\x1f\n" +
	"\azone_id\x18\x04 \x01(\x03B\x06\xbaH\x03\xc8\x01\x01R\x06zoneId\x12\x17\n" +
	"\ahost_id\x18\x05 \x01(\x03R\x06hostId\x12#\n" +
	"\rresponse_type\x18\x06 \x01(\tR\fresponseType\"{\n" +
	"/UploadTemplateDirectDownloadCertificateResponse\x12H\n" +
	"\x06result\x18\x01 \x01(\v20.cloudstack.management.direct.download.v1.ResultR\x06result\"\xa8\x01\n" +
	"\x19DirectDownloadCertificate\x12\x18\n" +
	"\x02id\x18\x01 \x01(\tB\b\xbaH\x05r\x03\xb0\x01\x01R\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12!\n" +
	"\fdisplay_name\x18\x03 \x01(\tR\vdisplayName\x12 \n" +
	"\vdescription\x18\x04 \x01(\tR\vdescription\x12\x18\n" +
	"\acreated\x18\x05 \x01(\tR\acreated\"F\n" +
	"\aSuccess\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\x12!\n" +
	"\fdisplay_text\x18\x02 \x01(\tR\vdisplayText\"\x9f\x01\n" +
	"\x06Result\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\x12!\n" +
	"\fdisplay_text\x18\x02 \x01(\tR\vdisplayText\x12\x18\n" +
	"\x02id\x18\x03 \x01(\tB\b\xbaH\x05r\x03\xb0\x01\x01R\x02id\x12\x1f\n" +
	"\x06job_id\x18\x04 \x01(\tB\b\xbaH\x05r\x03\xb0\x01\x01R\x05jobId\x12\x1d\n" +
	"\n" +
	"job_status\x18\x05 \x01(\tR\tjobStatus2\xc8\a\n" +
	"\x0fDownloadService\x12\xe4\x01\n" +
	"&ListTemplateDirectDownloadCertificates\x12W.cloudstack.management.direct.download.v1.ListTemplateDirectDownloadCertificatesRequest\x1aX.cloudstack.management.direct.download.v1.ListTemplateDirectDownloadCertificatesResponse\"\a\xc2>\x04\xc2>\x01\x02\x12\xf0\x01\n" +
	"*ProvisionTemplateDirectDownloadCertificate\x12[.cloudstack.management.direct.download.v1.ProvisionTemplateDirectDownloadCertificateRequest\x1a\\.cloudstack.management.direct.download.v1.ProvisionTemplateDirectDownloadCertificateResponse\"\a\xc2>\x04\xc2>\x01\x02\x12\xe7\x01\n" +
	"'RevokeTemplateDirectDownloadCertificate\x12X.cloudstack.management.direct.download.v1.RevokeTemplateDirectDownloadCertificateRequest\x1aY.cloudstack.management.direct.download.v1.RevokeTemplateDirectDownloadCertificateResponse\"\a\xc2>\x04\xc2>\x01\x02\x12\xe7\x01\n" +
	"'UploadTemplateDirectDownloadCertificate\x12X.cloudstack.management.direct.download.v1.UploadTemplateDirectDownloadCertificateRequest\x1aY.cloudstack.management.direct.download.v1.UploadTemplateDirectDownloadCertificateResponse\"\a\xc2>\x04\xc2>\x01\x02\x1a\a\xc2>\x04\xc2>\x01\x02B\xee\x02\n" +
	",com.cloudstack.management.direct.download.v1B\x10DownloadGenProtoP\x01Zggithub.com/walteh/cloudstack-proxy/gen/proto/golang/cloudstack/management/direct/download/v1;downloadv1\xa2\x02\x04CMDD\xaa\x02(Cloudstack.Management.Direct.Download.V1\xca\x02(Cloudstack\\Management\\Direct\\Download\\V1\xe2\x024Cloudstack\\Management\\Direct\\Download\\V1\\GPBMetadata\xea\x02,Cloudstack::Management::Direct::Download::V1b\beditionsp\xe8\a"

var (
	file_cloudstack_management_direct_download_v1_download_gen_proto_rawDescOnce sync.Once
	file_cloudstack_management_direct_download_v1_download_gen_proto_rawDescData []byte
)

func file_cloudstack_management_direct_download_v1_download_gen_proto_rawDescGZIP() []byte {
	file_cloudstack_management_direct_download_v1_download_gen_proto_rawDescOnce.Do(func() {
		file_cloudstack_management_direct_download_v1_download_gen_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_cloudstack_management_direct_download_v1_download_gen_proto_rawDesc), len(file_cloudstack_management_direct_download_v1_download_gen_proto_rawDesc)))
	})
	return file_cloudstack_management_direct_download_v1_download_gen_proto_rawDescData
}

var file_cloudstack_management_direct_download_v1_download_gen_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_cloudstack_management_direct_download_v1_download_gen_proto_goTypes = []any{
	(*ListTemplateDirectDownloadCertificatesRequest)(nil),      // 0: cloudstack.management.direct.download.v1.ListTemplateDirectDownloadCertificatesRequest
	(*ListTemplateDirectDownloadCertificatesResponse)(nil),     // 1: cloudstack.management.direct.download.v1.ListTemplateDirectDownloadCertificatesResponse
	(*ProvisionTemplateDirectDownloadCertificateRequest)(nil),  // 2: cloudstack.management.direct.download.v1.ProvisionTemplateDirectDownloadCertificateRequest
	(*ProvisionTemplateDirectDownloadCertificateResponse)(nil), // 3: cloudstack.management.direct.download.v1.ProvisionTemplateDirectDownloadCertificateResponse
	(*RevokeTemplateDirectDownloadCertificateRequest)(nil),     // 4: cloudstack.management.direct.download.v1.RevokeTemplateDirectDownloadCertificateRequest
	(*RevokeTemplateDirectDownloadCertificateResponse)(nil),    // 5: cloudstack.management.direct.download.v1.RevokeTemplateDirectDownloadCertificateResponse
	(*UploadTemplateDirectDownloadCertificateRequest)(nil),     // 6: cloudstack.management.direct.download.v1.UploadTemplateDirectDownloadCertificateRequest
	(*UploadTemplateDirectDownloadCertificateResponse)(nil),    // 7: cloudstack.management.direct.download.v1.UploadTemplateDirectDownloadCertificateResponse
	(*DirectDownloadCertificate)(nil),                          // 8: cloudstack.management.direct.download.v1.DirectDownloadCertificate
	(*Success)(nil),                                            // 9: cloudstack.management.direct.download.v1.Success
	(*Result)(nil),                                             // 10: cloudstack.management.direct.download.v1.Result
}
var file_cloudstack_management_direct_download_v1_download_gen_proto_depIdxs = []int32{
	8,  // 0: cloudstack.management.direct.download.v1.ListTemplateDirectDownloadCertificatesResponse.items:type_name -> cloudstack.management.direct.download.v1.DirectDownloadCertificate
	10, // 1: cloudstack.management.direct.download.v1.ProvisionTemplateDirectDownloadCertificateResponse.result:type_name -> cloudstack.management.direct.download.v1.Result
	10, // 2: cloudstack.management.direct.download.v1.RevokeTemplateDirectDownloadCertificateResponse.result:type_name -> cloudstack.management.direct.download.v1.Result
	10, // 3: cloudstack.management.direct.download.v1.UploadTemplateDirectDownloadCertificateResponse.result:type_name -> cloudstack.management.direct.download.v1.Result
	0,  // 4: cloudstack.management.direct.download.v1.DownloadService.ListTemplateDirectDownloadCertificates:input_type -> cloudstack.management.direct.download.v1.ListTemplateDirectDownloadCertificatesRequest
	2,  // 5: cloudstack.management.direct.download.v1.DownloadService.ProvisionTemplateDirectDownloadCertificate:input_type -> cloudstack.management.direct.download.v1.ProvisionTemplateDirectDownloadCertificateRequest
	4,  // 6: cloudstack.management.direct.download.v1.DownloadService.RevokeTemplateDirectDownloadCertificate:input_type -> cloudstack.management.direct.download.v1.RevokeTemplateDirectDownloadCertificateRequest
	6,  // 7: cloudstack.management.direct.download.v1.DownloadService.UploadTemplateDirectDownloadCertificate:input_type -> cloudstack.management.direct.download.v1.UploadTemplateDirectDownloadCertificateRequest
	1,  // 8: cloudstack.management.direct.download.v1.DownloadService.ListTemplateDirectDownloadCertificates:output_type -> cloudstack.management.direct.download.v1.ListTemplateDirectDownloadCertificatesResponse
	3,  // 9: cloudstack.management.direct.download.v1.DownloadService.ProvisionTemplateDirectDownloadCertificate:output_type -> cloudstack.management.direct.download.v1.ProvisionTemplateDirectDownloadCertificateResponse
	5,  // 10: cloudstack.management.direct.download.v1.DownloadService.RevokeTemplateDirectDownloadCertificate:output_type -> cloudstack.management.direct.download.v1.RevokeTemplateDirectDownloadCertificateResponse
	7,  // 11: cloudstack.management.direct.download.v1.DownloadService.UploadTemplateDirectDownloadCertificate:output_type -> cloudstack.management.direct.download.v1.UploadTemplateDirectDownloadCertificateResponse
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_cloudstack_management_direct_download_v1_download_gen_proto_init() }
func file_cloudstack_management_direct_download_v1_download_gen_proto_init() {
	if File_cloudstack_management_direct_download_v1_download_gen_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_cloudstack_management_direct_download_v1_download_gen_proto_rawDesc), len(file_cloudstack_management_direct_download_v1_download_gen_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cloudstack_management_direct_download_v1_download_gen_proto_goTypes,
		DependencyIndexes: file_cloudstack_management_direct_download_v1_download_gen_proto_depIdxs,
		MessageInfos:      file_cloudstack_management_direct_download_v1_download_gen_proto_msgTypes,
	}.Build()
	File_cloudstack_management_direct_download_v1_download_gen_proto = out.File
	file_cloudstack_management_direct_download_v1_download_gen_proto_goTypes = nil
	file_cloudstack_management_direct_download_v1_download_gen_proto_depIdxs = nil
}
