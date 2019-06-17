// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/feed.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// The operator.
type FeedAttributeOperation_Operator int32

const (
	// Unspecified.
	FeedAttributeOperation_UNSPECIFIED FeedAttributeOperation_Operator = 0
	// Used for return value only. Represents value unknown in this version.
	FeedAttributeOperation_UNKNOWN FeedAttributeOperation_Operator = 1
	// Add the attribute to the existing attributes.
	FeedAttributeOperation_ADD FeedAttributeOperation_Operator = 2
)

var FeedAttributeOperation_Operator_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ADD",
}

var FeedAttributeOperation_Operator_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"ADD":         2,
}

func (x FeedAttributeOperation_Operator) String() string {
	return proto.EnumName(FeedAttributeOperation_Operator_name, int32(x))
}

func (FeedAttributeOperation_Operator) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_90b38030e9747922, []int{2, 0}
}

// A feed.
type Feed struct {
	// The resource name of the feed.
	// Feed resource names have the form:
	//
	// `customers/{customer_id}/feeds/{feed_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of the feed.
	// This field is read-only.
	Id *wrappers.Int64Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Name of the feed. Required.
	Name *wrappers.StringValue `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// The Feed's attributes. Required on CREATE.
	// Disallowed on UPDATE. Use attribute_operations to add new attributes.
	Attributes []*FeedAttribute `protobuf:"bytes,4,rep,name=attributes,proto3" json:"attributes,omitempty"`
	// The list of operations changing the feed attributes. Attributes can only
	// be added, not removed.
	AttributeOperations []*FeedAttributeOperation `protobuf:"bytes,9,rep,name=attribute_operations,json=attributeOperations,proto3" json:"attribute_operations,omitempty"`
	// Specifies who manages the FeedAttributes for the Feed.
	Origin enums.FeedOriginEnum_FeedOrigin `protobuf:"varint,5,opt,name=origin,proto3,enum=google.ads.googleads.v1.enums.FeedOriginEnum_FeedOrigin" json:"origin,omitempty"`
	// Status of the feed.
	// This field is read-only.
	Status enums.FeedStatusEnum_FeedStatus `protobuf:"varint,8,opt,name=status,proto3,enum=google.ads.googleads.v1.enums.FeedStatusEnum_FeedStatus" json:"status,omitempty"`
	// The system data for the Feed. This data specifies information for
	// generating the feed items of the system generated feed.
	//
	// Types that are valid to be assigned to SystemFeedGenerationData:
	//	*Feed_PlacesLocationFeedData_
	//	*Feed_AffiliateLocationFeedData_
	SystemFeedGenerationData isFeed_SystemFeedGenerationData `protobuf_oneof:"system_feed_generation_data"`
	XXX_NoUnkeyedLiteral     struct{}                        `json:"-"`
	XXX_unrecognized         []byte                          `json:"-"`
	XXX_sizecache            int32                           `json:"-"`
}

func (m *Feed) Reset()         { *m = Feed{} }
func (m *Feed) String() string { return proto.CompactTextString(m) }
func (*Feed) ProtoMessage()    {}
func (*Feed) Descriptor() ([]byte, []int) {
	return fileDescriptor_90b38030e9747922, []int{0}
}

func (m *Feed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Feed.Unmarshal(m, b)
}
func (m *Feed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Feed.Marshal(b, m, deterministic)
}
func (m *Feed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Feed.Merge(m, src)
}
func (m *Feed) XXX_Size() int {
	return xxx_messageInfo_Feed.Size(m)
}
func (m *Feed) XXX_DiscardUnknown() {
	xxx_messageInfo_Feed.DiscardUnknown(m)
}

var xxx_messageInfo_Feed proto.InternalMessageInfo

func (m *Feed) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *Feed) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Feed) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *Feed) GetAttributes() []*FeedAttribute {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *Feed) GetAttributeOperations() []*FeedAttributeOperation {
	if m != nil {
		return m.AttributeOperations
	}
	return nil
}

func (m *Feed) GetOrigin() enums.FeedOriginEnum_FeedOrigin {
	if m != nil {
		return m.Origin
	}
	return enums.FeedOriginEnum_UNSPECIFIED
}

func (m *Feed) GetStatus() enums.FeedStatusEnum_FeedStatus {
	if m != nil {
		return m.Status
	}
	return enums.FeedStatusEnum_UNSPECIFIED
}

type isFeed_SystemFeedGenerationData interface {
	isFeed_SystemFeedGenerationData()
}

type Feed_PlacesLocationFeedData_ struct {
	PlacesLocationFeedData *Feed_PlacesLocationFeedData `protobuf:"bytes,6,opt,name=places_location_feed_data,json=placesLocationFeedData,proto3,oneof"`
}

type Feed_AffiliateLocationFeedData_ struct {
	AffiliateLocationFeedData *Feed_AffiliateLocationFeedData `protobuf:"bytes,7,opt,name=affiliate_location_feed_data,json=affiliateLocationFeedData,proto3,oneof"`
}

func (*Feed_PlacesLocationFeedData_) isFeed_SystemFeedGenerationData() {}

func (*Feed_AffiliateLocationFeedData_) isFeed_SystemFeedGenerationData() {}

func (m *Feed) GetSystemFeedGenerationData() isFeed_SystemFeedGenerationData {
	if m != nil {
		return m.SystemFeedGenerationData
	}
	return nil
}

func (m *Feed) GetPlacesLocationFeedData() *Feed_PlacesLocationFeedData {
	if x, ok := m.GetSystemFeedGenerationData().(*Feed_PlacesLocationFeedData_); ok {
		return x.PlacesLocationFeedData
	}
	return nil
}

func (m *Feed) GetAffiliateLocationFeedData() *Feed_AffiliateLocationFeedData {
	if x, ok := m.GetSystemFeedGenerationData().(*Feed_AffiliateLocationFeedData_); ok {
		return x.AffiliateLocationFeedData
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Feed) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Feed_PlacesLocationFeedData_)(nil),
		(*Feed_AffiliateLocationFeedData_)(nil),
	}
}

// Data used to configure a location feed populated from Google My Business
// Locations.
type Feed_PlacesLocationFeedData struct {
	// Required authentication token (from OAuth API) for the email.
	// This field can only be specified in a create request. All its subfields
	// are not selectable.
	OauthInfo *Feed_PlacesLocationFeedData_OAuthInfo `protobuf:"bytes,1,opt,name=oauth_info,json=oauthInfo,proto3" json:"oauth_info,omitempty"`
	// Email address of a Google My Business account or email address of a
	// manager of the Google My Business account. Required.
	EmailAddress *wrappers.StringValue `protobuf:"bytes,2,opt,name=email_address,json=emailAddress,proto3" json:"email_address,omitempty"`
	// Plus page ID of the managed business whose locations should be used. If
	// this field is not set, then all businesses accessible by the user
	// (specified by email_address) are used.
	// This field is mutate-only and is not selectable.
	BusinessAccountId *wrappers.StringValue `protobuf:"bytes,10,opt,name=business_account_id,json=businessAccountId,proto3" json:"business_account_id,omitempty"`
	// Used to filter Google My Business listings by business name. If
	// business_name_filter is set, only listings with a matching business name
	// are candidates to be sync'd into FeedItems.
	BusinessNameFilter *wrappers.StringValue `protobuf:"bytes,4,opt,name=business_name_filter,json=businessNameFilter,proto3" json:"business_name_filter,omitempty"`
	// Used to filter Google My Business listings by categories. If entries
	// exist in category_filters, only listings that belong to any of the
	// categories are candidates to be sync'd into FeedItems. If no entries
	// exist in category_filters, then all listings are candidates for syncing.
	CategoryFilters []*wrappers.StringValue `protobuf:"bytes,5,rep,name=category_filters,json=categoryFilters,proto3" json:"category_filters,omitempty"`
	// Used to filter Google My Business listings by labels. If entries exist in
	// label_filters, only listings that has any of the labels set are
	// candidates to be synchronized into FeedItems. If no entries exist in
	// label_filters, then all listings are candidates for syncing.
	LabelFilters         []*wrappers.StringValue `protobuf:"bytes,6,rep,name=label_filters,json=labelFilters,proto3" json:"label_filters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *Feed_PlacesLocationFeedData) Reset()         { *m = Feed_PlacesLocationFeedData{} }
func (m *Feed_PlacesLocationFeedData) String() string { return proto.CompactTextString(m) }
func (*Feed_PlacesLocationFeedData) ProtoMessage()    {}
func (*Feed_PlacesLocationFeedData) Descriptor() ([]byte, []int) {
	return fileDescriptor_90b38030e9747922, []int{0, 0}
}

func (m *Feed_PlacesLocationFeedData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Feed_PlacesLocationFeedData.Unmarshal(m, b)
}
func (m *Feed_PlacesLocationFeedData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Feed_PlacesLocationFeedData.Marshal(b, m, deterministic)
}
func (m *Feed_PlacesLocationFeedData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Feed_PlacesLocationFeedData.Merge(m, src)
}
func (m *Feed_PlacesLocationFeedData) XXX_Size() int {
	return xxx_messageInfo_Feed_PlacesLocationFeedData.Size(m)
}
func (m *Feed_PlacesLocationFeedData) XXX_DiscardUnknown() {
	xxx_messageInfo_Feed_PlacesLocationFeedData.DiscardUnknown(m)
}

var xxx_messageInfo_Feed_PlacesLocationFeedData proto.InternalMessageInfo

func (m *Feed_PlacesLocationFeedData) GetOauthInfo() *Feed_PlacesLocationFeedData_OAuthInfo {
	if m != nil {
		return m.OauthInfo
	}
	return nil
}

func (m *Feed_PlacesLocationFeedData) GetEmailAddress() *wrappers.StringValue {
	if m != nil {
		return m.EmailAddress
	}
	return nil
}

func (m *Feed_PlacesLocationFeedData) GetBusinessAccountId() *wrappers.StringValue {
	if m != nil {
		return m.BusinessAccountId
	}
	return nil
}

func (m *Feed_PlacesLocationFeedData) GetBusinessNameFilter() *wrappers.StringValue {
	if m != nil {
		return m.BusinessNameFilter
	}
	return nil
}

func (m *Feed_PlacesLocationFeedData) GetCategoryFilters() []*wrappers.StringValue {
	if m != nil {
		return m.CategoryFilters
	}
	return nil
}

func (m *Feed_PlacesLocationFeedData) GetLabelFilters() []*wrappers.StringValue {
	if m != nil {
		return m.LabelFilters
	}
	return nil
}

// Data used for authorization using OAuth.
type Feed_PlacesLocationFeedData_OAuthInfo struct {
	// The HTTP method used to obtain authorization.
	HttpMethod *wrappers.StringValue `protobuf:"bytes,1,opt,name=http_method,json=httpMethod,proto3" json:"http_method,omitempty"`
	// The HTTP request URL used to obtain authorization.
	HttpRequestUrl *wrappers.StringValue `protobuf:"bytes,2,opt,name=http_request_url,json=httpRequestUrl,proto3" json:"http_request_url,omitempty"`
	// The HTTP authorization header used to obtain authorization.
	HttpAuthorizationHeader *wrappers.StringValue `protobuf:"bytes,3,opt,name=http_authorization_header,json=httpAuthorizationHeader,proto3" json:"http_authorization_header,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}              `json:"-"`
	XXX_unrecognized        []byte                `json:"-"`
	XXX_sizecache           int32                 `json:"-"`
}

func (m *Feed_PlacesLocationFeedData_OAuthInfo) Reset()         { *m = Feed_PlacesLocationFeedData_OAuthInfo{} }
func (m *Feed_PlacesLocationFeedData_OAuthInfo) String() string { return proto.CompactTextString(m) }
func (*Feed_PlacesLocationFeedData_OAuthInfo) ProtoMessage()    {}
func (*Feed_PlacesLocationFeedData_OAuthInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_90b38030e9747922, []int{0, 0, 0}
}

func (m *Feed_PlacesLocationFeedData_OAuthInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Feed_PlacesLocationFeedData_OAuthInfo.Unmarshal(m, b)
}
func (m *Feed_PlacesLocationFeedData_OAuthInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Feed_PlacesLocationFeedData_OAuthInfo.Marshal(b, m, deterministic)
}
func (m *Feed_PlacesLocationFeedData_OAuthInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Feed_PlacesLocationFeedData_OAuthInfo.Merge(m, src)
}
func (m *Feed_PlacesLocationFeedData_OAuthInfo) XXX_Size() int {
	return xxx_messageInfo_Feed_PlacesLocationFeedData_OAuthInfo.Size(m)
}
func (m *Feed_PlacesLocationFeedData_OAuthInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_Feed_PlacesLocationFeedData_OAuthInfo.DiscardUnknown(m)
}

var xxx_messageInfo_Feed_PlacesLocationFeedData_OAuthInfo proto.InternalMessageInfo

func (m *Feed_PlacesLocationFeedData_OAuthInfo) GetHttpMethod() *wrappers.StringValue {
	if m != nil {
		return m.HttpMethod
	}
	return nil
}

func (m *Feed_PlacesLocationFeedData_OAuthInfo) GetHttpRequestUrl() *wrappers.StringValue {
	if m != nil {
		return m.HttpRequestUrl
	}
	return nil
}

func (m *Feed_PlacesLocationFeedData_OAuthInfo) GetHttpAuthorizationHeader() *wrappers.StringValue {
	if m != nil {
		return m.HttpAuthorizationHeader
	}
	return nil
}

// Data used to configure an affiliate location feed populated with the
// specified chains.
type Feed_AffiliateLocationFeedData struct {
	// The list of chains that the affiliate location feed will sync the
	// locations from.
	ChainIds []*wrappers.Int64Value `protobuf:"bytes,1,rep,name=chain_ids,json=chainIds,proto3" json:"chain_ids,omitempty"`
	// The relationship the chains have with the advertiser.
	RelationshipType     enums.AffiliateLocationFeedRelationshipTypeEnum_AffiliateLocationFeedRelationshipType `protobuf:"varint,2,opt,name=relationship_type,json=relationshipType,proto3,enum=google.ads.googleads.v1.enums.AffiliateLocationFeedRelationshipTypeEnum_AffiliateLocationFeedRelationshipType" json:"relationship_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                                              `json:"-"`
	XXX_unrecognized     []byte                                                                                `json:"-"`
	XXX_sizecache        int32                                                                                 `json:"-"`
}

func (m *Feed_AffiliateLocationFeedData) Reset()         { *m = Feed_AffiliateLocationFeedData{} }
func (m *Feed_AffiliateLocationFeedData) String() string { return proto.CompactTextString(m) }
func (*Feed_AffiliateLocationFeedData) ProtoMessage()    {}
func (*Feed_AffiliateLocationFeedData) Descriptor() ([]byte, []int) {
	return fileDescriptor_90b38030e9747922, []int{0, 1}
}

func (m *Feed_AffiliateLocationFeedData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Feed_AffiliateLocationFeedData.Unmarshal(m, b)
}
func (m *Feed_AffiliateLocationFeedData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Feed_AffiliateLocationFeedData.Marshal(b, m, deterministic)
}
func (m *Feed_AffiliateLocationFeedData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Feed_AffiliateLocationFeedData.Merge(m, src)
}
func (m *Feed_AffiliateLocationFeedData) XXX_Size() int {
	return xxx_messageInfo_Feed_AffiliateLocationFeedData.Size(m)
}
func (m *Feed_AffiliateLocationFeedData) XXX_DiscardUnknown() {
	xxx_messageInfo_Feed_AffiliateLocationFeedData.DiscardUnknown(m)
}

var xxx_messageInfo_Feed_AffiliateLocationFeedData proto.InternalMessageInfo

func (m *Feed_AffiliateLocationFeedData) GetChainIds() []*wrappers.Int64Value {
	if m != nil {
		return m.ChainIds
	}
	return nil
}

func (m *Feed_AffiliateLocationFeedData) GetRelationshipType() enums.AffiliateLocationFeedRelationshipTypeEnum_AffiliateLocationFeedRelationshipType {
	if m != nil {
		return m.RelationshipType
	}
	return enums.AffiliateLocationFeedRelationshipTypeEnum_UNSPECIFIED
}

// FeedAttributes define the types of data expected to be present in a Feed. A
// single FeedAttribute specifies the expected type of the FeedItemAttributes
// with the same FeedAttributeId. Optionally, a FeedAttribute can be marked as
// being part of a FeedItem's unique key.
type FeedAttribute struct {
	// ID of the attribute.
	Id *wrappers.Int64Value `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The name of the attribute. Required.
	Name *wrappers.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Data type for feed attribute. Required.
	Type enums.FeedAttributeTypeEnum_FeedAttributeType `protobuf:"varint,3,opt,name=type,proto3,enum=google.ads.googleads.v1.enums.FeedAttributeTypeEnum_FeedAttributeType" json:"type,omitempty"`
	// Indicates that data corresponding to this attribute is part of a
	// FeedItem's unique key. It defaults to false if it is unspecified. Note
	// that a unique key is not required in a Feed's schema, in which case the
	// FeedItems must be referenced by their feed_item_id.
	IsPartOfKey          *wrappers.BoolValue `protobuf:"bytes,4,opt,name=is_part_of_key,json=isPartOfKey,proto3" json:"is_part_of_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *FeedAttribute) Reset()         { *m = FeedAttribute{} }
func (m *FeedAttribute) String() string { return proto.CompactTextString(m) }
func (*FeedAttribute) ProtoMessage()    {}
func (*FeedAttribute) Descriptor() ([]byte, []int) {
	return fileDescriptor_90b38030e9747922, []int{1}
}

func (m *FeedAttribute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedAttribute.Unmarshal(m, b)
}
func (m *FeedAttribute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedAttribute.Marshal(b, m, deterministic)
}
func (m *FeedAttribute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedAttribute.Merge(m, src)
}
func (m *FeedAttribute) XXX_Size() int {
	return xxx_messageInfo_FeedAttribute.Size(m)
}
func (m *FeedAttribute) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedAttribute.DiscardUnknown(m)
}

var xxx_messageInfo_FeedAttribute proto.InternalMessageInfo

func (m *FeedAttribute) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *FeedAttribute) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *FeedAttribute) GetType() enums.FeedAttributeTypeEnum_FeedAttributeType {
	if m != nil {
		return m.Type
	}
	return enums.FeedAttributeTypeEnum_UNSPECIFIED
}

func (m *FeedAttribute) GetIsPartOfKey() *wrappers.BoolValue {
	if m != nil {
		return m.IsPartOfKey
	}
	return nil
}

// Operation to be performed on a feed attribute list in a mutate.
type FeedAttributeOperation struct {
	// Type of list operation to perform.
	Operator FeedAttributeOperation_Operator `protobuf:"varint,1,opt,name=operator,proto3,enum=google.ads.googleads.v1.resources.FeedAttributeOperation_Operator" json:"operator,omitempty"`
	// The feed attribute being added to the list.
	Value                *FeedAttribute `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *FeedAttributeOperation) Reset()         { *m = FeedAttributeOperation{} }
func (m *FeedAttributeOperation) String() string { return proto.CompactTextString(m) }
func (*FeedAttributeOperation) ProtoMessage()    {}
func (*FeedAttributeOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_90b38030e9747922, []int{2}
}

func (m *FeedAttributeOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedAttributeOperation.Unmarshal(m, b)
}
func (m *FeedAttributeOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedAttributeOperation.Marshal(b, m, deterministic)
}
func (m *FeedAttributeOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedAttributeOperation.Merge(m, src)
}
func (m *FeedAttributeOperation) XXX_Size() int {
	return xxx_messageInfo_FeedAttributeOperation.Size(m)
}
func (m *FeedAttributeOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedAttributeOperation.DiscardUnknown(m)
}

var xxx_messageInfo_FeedAttributeOperation proto.InternalMessageInfo

func (m *FeedAttributeOperation) GetOperator() FeedAttributeOperation_Operator {
	if m != nil {
		return m.Operator
	}
	return FeedAttributeOperation_UNSPECIFIED
}

func (m *FeedAttributeOperation) GetValue() *FeedAttribute {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.resources.FeedAttributeOperation_Operator", FeedAttributeOperation_Operator_name, FeedAttributeOperation_Operator_value)
	proto.RegisterType((*Feed)(nil), "google.ads.googleads.v1.resources.Feed")
	proto.RegisterType((*Feed_PlacesLocationFeedData)(nil), "google.ads.googleads.v1.resources.Feed.PlacesLocationFeedData")
	proto.RegisterType((*Feed_PlacesLocationFeedData_OAuthInfo)(nil), "google.ads.googleads.v1.resources.Feed.PlacesLocationFeedData.OAuthInfo")
	proto.RegisterType((*Feed_AffiliateLocationFeedData)(nil), "google.ads.googleads.v1.resources.Feed.AffiliateLocationFeedData")
	proto.RegisterType((*FeedAttribute)(nil), "google.ads.googleads.v1.resources.FeedAttribute")
	proto.RegisterType((*FeedAttributeOperation)(nil), "google.ads.googleads.v1.resources.FeedAttributeOperation")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/feed.proto", fileDescriptor_90b38030e9747922)
}

var fileDescriptor_90b38030e9747922 = []byte{
	// 1031 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x96, 0xcf, 0x6f, 0xe3, 0x44,
	0x14, 0xc7, 0xd7, 0x49, 0xfa, 0x23, 0x93, 0x36, 0x9b, 0x9d, 0x5d, 0x15, 0x37, 0x5b, 0x50, 0xb7,
	0x68, 0xa5, 0x4a, 0x20, 0x67, 0x13, 0x10, 0x2c, 0x41, 0x80, 0x1c, 0xda, 0xb4, 0x61, 0x77, 0x93,
	0xc8, 0xa5, 0x05, 0xad, 0x2a, 0xac, 0x49, 0x3c, 0x71, 0x46, 0x38, 0x1e, 0x33, 0x33, 0x2e, 0x0a,
	0x5c, 0xb9, 0xf1, 0x27, 0x70, 0xe3, 0xc8, 0x9f, 0xc2, 0xff, 0xc0, 0x19, 0x09, 0x71, 0x5b, 0x89,
	0x33, 0xf2, 0x8c, 0x6d, 0x42, 0xf3, 0xa3, 0x2e, 0xdc, 0xc6, 0x33, 0xdf, 0xef, 0xc7, 0xcf, 0xef,
	0xe5, 0xbd, 0x09, 0x78, 0xdb, 0xa5, 0xd4, 0xf5, 0x70, 0x0d, 0x39, 0xbc, 0xa6, 0x96, 0xd1, 0xea,
	0xaa, 0x5e, 0x63, 0x98, 0xd3, 0x90, 0x0d, 0x31, 0xaf, 0x8d, 0x30, 0x76, 0x8c, 0x80, 0x51, 0x41,
	0xe1, 0x23, 0x25, 0x31, 0x90, 0xc3, 0x8d, 0x54, 0x6d, 0x5c, 0xd5, 0x8d, 0x54, 0x5d, 0x7d, 0xb1,
	0x0c, 0x88, 0xfd, 0x70, 0xc2, 0x6b, 0x68, 0x34, 0x22, 0x1e, 0x41, 0x02, 0xdb, 0x1e, 0x1d, 0x22,
	0x41, 0xa8, 0x6f, 0x47, 0x7c, 0x9b, 0x61, 0x4f, 0x3e, 0xf1, 0x31, 0x09, 0x6c, 0x31, 0x0d, 0xb0,
	0x7a, 0x63, 0xf5, 0xfd, 0xd5, 0x38, 0xe9, 0x45, 0x42, 0x30, 0x32, 0x08, 0x05, 0x9e, 0x35, 0xd6,
	0x32, 0x18, 0x29, 0x23, 0x2e, 0xf1, 0x6f, 0x61, 0xe0, 0x02, 0x89, 0x90, 0xc7, 0x86, 0xbd, 0xc4,
	0x10, 0x90, 0x1a, 0xf2, 0x7d, 0x2a, 0xd4, 0x07, 0xc4, 0xa7, 0x6f, 0xc4, 0xa7, 0xf2, 0x69, 0x10,
	0x8e, 0x6a, 0xdf, 0x32, 0x14, 0x04, 0x98, 0xc5, 0xe7, 0x07, 0xbf, 0x6f, 0x83, 0x42, 0x1b, 0x63,
	0x07, 0xbe, 0x09, 0xb6, 0x93, 0xec, 0xd9, 0x3e, 0x9a, 0x60, 0x5d, 0xdb, 0xd7, 0x0e, 0x8b, 0xd6,
	0x56, 0xb2, 0xd9, 0x45, 0x13, 0x0c, 0xdf, 0x02, 0x39, 0xe2, 0xe8, 0xb9, 0x7d, 0xed, 0xb0, 0xd4,
	0x78, 0x18, 0xa7, 0xde, 0x48, 0xd0, 0x46, 0xc7, 0x17, 0xef, 0xbd, 0x7b, 0x81, 0xbc, 0x10, 0x5b,
	0x39, 0xe2, 0xc0, 0x27, 0xa0, 0x20, 0x41, 0x79, 0x29, 0xdf, 0x9b, 0x93, 0x9f, 0x09, 0x46, 0x7c,
	0x57, 0xe9, 0xa5, 0x12, 0xf6, 0x01, 0x48, 0x93, 0xc8, 0xf5, 0xc2, 0x7e, 0xfe, 0xb0, 0xd4, 0x78,
	0x62, 0xdc, 0x58, 0x6c, 0x23, 0xfa, 0x00, 0x33, 0x31, 0x5a, 0x33, 0x0c, 0xe8, 0x81, 0x07, 0xff,
	0x94, 0x85, 0x06, 0x98, 0xa9, 0xe4, 0xe8, 0x45, 0xc9, 0xfe, 0xe0, 0xb6, 0xec, 0x5e, 0x42, 0xb0,
	0xee, 0xa3, 0xb9, 0x3d, 0x0e, 0xfb, 0x60, 0x5d, 0xd5, 0x52, 0x5f, 0xdb, 0xd7, 0x0e, 0xcb, 0x8d,
	0xa7, 0x4b, 0xf9, 0xb2, 0x98, 0x92, 0xdd, 0x93, 0x86, 0x63, 0x3f, 0x9c, 0xcc, 0x3c, 0x5a, 0x31,
	0x27, 0x22, 0xaa, 0x62, 0xeb, 0x9b, 0x99, 0x89, 0x67, 0xd2, 0x90, 0x12, 0xd5, 0xa3, 0x15, 0x73,
	0xe0, 0xf7, 0x60, 0x37, 0xf0, 0xd0, 0x10, 0xf3, 0x6b, 0xbf, 0x7c, 0x07, 0x09, 0xa4, 0xaf, 0xcb,
	0x52, 0x7d, 0x9c, 0x31, 0x2d, 0x46, 0x5f, 0x82, 0x9e, 0xc7, 0x9c, 0x68, 0xeb, 0x08, 0x09, 0x74,
	0x7a, 0xc7, 0xda, 0x09, 0x16, 0x9e, 0xc0, 0x1f, 0x34, 0xb0, 0xb7, 0xac, 0xf5, 0x64, 0x00, 0x1b,
	0x32, 0x00, 0x33, 0x6b, 0x00, 0x66, 0xc2, 0x5a, 0x10, 0xc3, 0x2e, 0x5a, 0x76, 0x58, 0xfd, 0x6d,
	0x0d, 0xec, 0x2c, 0x8e, 0x1d, 0xba, 0x00, 0x50, 0x14, 0x8a, 0xb1, 0x4d, 0xfc, 0x11, 0x95, 0x3d,
	0x50, 0x6a, 0x9c, 0xfe, 0xbf, 0x7c, 0x18, 0x3d, 0x33, 0x14, 0xe3, 0x8e, 0x3f, 0xa2, 0x56, 0x51,
	0xb2, 0xa3, 0x25, 0x34, 0xc1, 0x36, 0x9e, 0x20, 0xe2, 0xd9, 0xc8, 0x71, 0x18, 0xe6, 0x3c, 0xee,
	0xaa, 0xd5, 0x6d, 0xb2, 0x25, 0x2d, 0xa6, 0x72, 0xc0, 0xe7, 0xe0, 0xfe, 0x20, 0xe4, 0xc4, 0xc7,
	0x9c, 0xdb, 0x68, 0x38, 0xa4, 0xa1, 0x2f, 0x6c, 0xe2, 0xe8, 0x20, 0x03, 0xe8, 0x5e, 0x62, 0x34,
	0x95, 0xaf, 0xe3, 0xc0, 0x2e, 0x78, 0x90, 0xd2, 0xa2, 0x6e, 0xb4, 0x47, 0xc4, 0x13, 0x98, 0xe9,
	0x85, 0x0c, 0x38, 0x98, 0x38, 0xa3, 0x29, 0xd1, 0x96, 0x3e, 0x78, 0x02, 0x2a, 0x43, 0x24, 0xb0,
	0x4b, 0xd9, 0x34, 0x46, 0x71, 0x7d, 0x4d, 0xb6, 0xdd, 0x6a, 0xd6, 0xdd, 0xc4, 0xa5, 0x38, 0x3c,
	0xca, 0x94, 0x87, 0x06, 0xd8, 0x4b, 0x29, 0xeb, 0x19, 0x28, 0x5b, 0xd2, 0x12, 0x23, 0xaa, 0x7f,
	0x6a, 0xa0, 0x98, 0x56, 0x01, 0x7e, 0x04, 0x4a, 0x63, 0x21, 0x02, 0x7b, 0x82, 0xc5, 0x98, 0x3a,
	0x71, 0x91, 0x57, 0xe3, 0x40, 0x64, 0x78, 0x21, 0xf5, 0xb0, 0x0d, 0x2a, 0xd2, 0xce, 0xf0, 0x37,
	0x21, 0xe6, 0xc2, 0x0e, 0x99, 0x97, 0xa9, 0x78, 0xe5, 0xc8, 0x65, 0x29, 0xd3, 0x39, 0xf3, 0xe0,
	0x97, 0x60, 0x57, 0x72, 0xa2, 0x9f, 0x04, 0x65, 0xe4, 0x3b, 0xd5, 0x0b, 0x63, 0x8c, 0x1c, 0xcc,
	0x32, 0x0d, 0xcd, 0xd7, 0x22, 0xbb, 0x39, 0xeb, 0x3e, 0x95, 0xe6, 0xea, 0x2b, 0x0d, 0xec, 0x2e,
	0x6d, 0x0d, 0xf8, 0x14, 0x14, 0x87, 0x63, 0x44, 0x7c, 0x9b, 0x38, 0x5c, 0xd7, 0x64, 0x2e, 0x57,
	0xce, 0xf2, 0x4d, 0xa9, 0xee, 0x38, 0x1c, 0xfe, 0xa4, 0x81, 0x7b, 0x73, 0x37, 0xa4, 0xfc, 0xf6,
	0x72, 0xc3, 0xbf, 0x61, 0x32, 0x2d, 0x8c, 0xc7, 0x9a, 0x81, 0x7d, 0x3e, 0x0d, 0xb0, 0x1c, 0x5a,
	0x99, 0x94, 0x56, 0x85, 0x5d, 0xdb, 0x69, 0xbd, 0x0e, 0x1e, 0xf2, 0x29, 0x17, 0x78, 0xa2, 0xe6,
	0x89, 0x8b, 0xfd, 0x78, 0x30, 0xcb, 0xd1, 0x72, 0xf0, 0x63, 0x0e, 0x6c, 0xff, 0x6b, 0x98, 0xc7,
	0xb7, 0x99, 0x76, 0xbb, 0xdb, 0x2c, 0x97, 0xf9, 0x36, 0x7b, 0x09, 0x0a, 0x32, 0x3f, 0x79, 0x99,
	0x9f, 0x76, 0x86, 0xc9, 0x9d, 0x86, 0x96, 0xe6, 0x62, 0x6e, 0xd7, 0x92, 0x4c, 0xf8, 0x09, 0x28,
	0x13, 0x6e, 0x07, 0x88, 0x09, 0x9b, 0x8e, 0xec, 0xaf, 0xf1, 0x34, 0x6e, 0xd3, 0xea, 0x5c, 0x5c,
	0x2d, 0x4a, 0x3d, 0x15, 0x55, 0x89, 0xf0, 0x3e, 0x62, 0xa2, 0x37, 0x7a, 0x86, 0xa7, 0x07, 0xaf,
	0x34, 0xb0, 0xb3, 0xf8, 0x6a, 0x83, 0x5f, 0x81, 0x4d, 0x75, 0x53, 0x52, 0x26, 0x93, 0x53, 0x6e,
	0xb4, 0xfe, 0xf3, 0x3d, 0x69, 0xf4, 0x62, 0x92, 0x95, 0x32, 0x61, 0x1b, 0xac, 0x5d, 0x45, 0x01,
	0xc5, 0xa9, 0xbc, 0xfd, 0x05, 0xaf, 0xec, 0x07, 0x75, 0xb0, 0x99, 0xd0, 0xe1, 0x5d, 0x50, 0x3a,
	0xef, 0x9e, 0xf5, 0x8f, 0x3f, 0xed, 0xb4, 0x3b, 0xc7, 0x47, 0x95, 0x3b, 0xb0, 0x04, 0x36, 0xce,
	0xbb, 0xcf, 0xba, 0xbd, 0x2f, 0xba, 0x15, 0x0d, 0x6e, 0x80, 0xbc, 0x79, 0x74, 0x54, 0xc9, 0xb5,
	0xfe, 0xd2, 0xc0, 0xe3, 0x21, 0x9d, 0xdc, 0xfc, 0xc6, 0x56, 0x31, 0x7a, 0x65, 0x3f, 0x4a, 0x62,
	0x5f, 0x7b, 0xf9, 0x59, 0xac, 0x77, 0xa9, 0x87, 0x7c, 0xd7, 0xa0, 0xcc, 0xad, 0xb9, 0xd8, 0x97,
	0x29, 0x4e, 0xfe, 0xa3, 0x05, 0x84, 0xaf, 0xf8, 0xf3, 0xfa, 0x61, 0xba, 0xfa, 0x39, 0x97, 0x3f,
	0x31, 0xcd, 0x5f, 0x72, 0x8f, 0x4e, 0x14, 0xd2, 0x74, 0xb8, 0xa1, 0x96, 0xd1, 0xea, 0xa2, 0x6e,
	0x58, 0x89, 0xf2, 0xd7, 0x44, 0x73, 0x69, 0x3a, 0xfc, 0x32, 0xd5, 0x5c, 0x5e, 0xd4, 0x2f, 0x53,
	0xcd, 0x1f, 0xb9, 0xc7, 0xea, 0xa0, 0xd9, 0x34, 0x1d, 0xde, 0x6c, 0xa6, 0xaa, 0x66, 0xf3, 0xa2,
	0xde, 0x6c, 0xa6, 0xba, 0xc1, 0xba, 0x0c, 0xf6, 0x9d, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xcf,
	0xea, 0xcb, 0x6d, 0x68, 0x0b, 0x00, 0x00,
}
