// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/errors/campaign_experiment_error.proto

package errors

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

// Enum describing possible campaign experiment errors.
type CampaignExperimentErrorEnum_CampaignExperimentError int32

const (
	// Enum unspecified.
	CampaignExperimentErrorEnum_UNSPECIFIED CampaignExperimentErrorEnum_CampaignExperimentError = 0
	// The received error code is not known in this version.
	CampaignExperimentErrorEnum_UNKNOWN CampaignExperimentErrorEnum_CampaignExperimentError = 1
	// An active campaign or experiment with this name already exists.
	CampaignExperimentErrorEnum_DUPLICATE_NAME CampaignExperimentErrorEnum_CampaignExperimentError = 2
	// Experiment cannot be updated from the current state to the
	// requested target state. For example, an experiment can only graduate
	// if its status is ENABLED.
	CampaignExperimentErrorEnum_INVALID_TRANSITION CampaignExperimentErrorEnum_CampaignExperimentError = 3
	// Cannot create an experiment from a campaign using an explicitly shared
	// budget.
	CampaignExperimentErrorEnum_CANNOT_CREATE_EXPERIMENT_WITH_SHARED_BUDGET CampaignExperimentErrorEnum_CampaignExperimentError = 4
	// Cannot create an experiment for a removed base campaign.
	CampaignExperimentErrorEnum_CANNOT_CREATE_EXPERIMENT_FOR_REMOVED_BASE_CAMPAIGN CampaignExperimentErrorEnum_CampaignExperimentError = 5
	// Cannot create an experiment from a draft, which has a status other than
	// proposed.
	CampaignExperimentErrorEnum_CANNOT_CREATE_EXPERIMENT_FOR_NON_PROPOSED_DRAFT CampaignExperimentErrorEnum_CampaignExperimentError = 6
	// This customer is not allowed to create an experiment.
	CampaignExperimentErrorEnum_CUSTOMER_CANNOT_CREATE_EXPERIMENT CampaignExperimentErrorEnum_CampaignExperimentError = 7
	// This campaign is not allowed to create an experiment.
	CampaignExperimentErrorEnum_CAMPAIGN_CANNOT_CREATE_EXPERIMENT CampaignExperimentErrorEnum_CampaignExperimentError = 8
	// Trying to set an experiment duration which overlaps with another
	// experiment.
	CampaignExperimentErrorEnum_EXPERIMENT_DURATIONS_MUST_NOT_OVERLAP CampaignExperimentErrorEnum_CampaignExperimentError = 9
	// All non-removed experiments must start and end within their campaign's
	// duration.
	CampaignExperimentErrorEnum_EXPERIMENT_DURATION_MUST_BE_WITHIN_CAMPAIGN_DURATION CampaignExperimentErrorEnum_CampaignExperimentError = 10
	// The experiment cannot be modified because its status is in a terminal
	// state, such as REMOVED.
	CampaignExperimentErrorEnum_CANNOT_MUTATE_EXPERIMENT_DUE_TO_STATUS CampaignExperimentErrorEnum_CampaignExperimentError = 11
)

var CampaignExperimentErrorEnum_CampaignExperimentError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "DUPLICATE_NAME",
	3:  "INVALID_TRANSITION",
	4:  "CANNOT_CREATE_EXPERIMENT_WITH_SHARED_BUDGET",
	5:  "CANNOT_CREATE_EXPERIMENT_FOR_REMOVED_BASE_CAMPAIGN",
	6:  "CANNOT_CREATE_EXPERIMENT_FOR_NON_PROPOSED_DRAFT",
	7:  "CUSTOMER_CANNOT_CREATE_EXPERIMENT",
	8:  "CAMPAIGN_CANNOT_CREATE_EXPERIMENT",
	9:  "EXPERIMENT_DURATIONS_MUST_NOT_OVERLAP",
	10: "EXPERIMENT_DURATION_MUST_BE_WITHIN_CAMPAIGN_DURATION",
	11: "CANNOT_MUTATE_EXPERIMENT_DUE_TO_STATUS",
}

var CampaignExperimentErrorEnum_CampaignExperimentError_value = map[string]int32{
	"UNSPECIFIED":        0,
	"UNKNOWN":            1,
	"DUPLICATE_NAME":     2,
	"INVALID_TRANSITION": 3,
	"CANNOT_CREATE_EXPERIMENT_WITH_SHARED_BUDGET":          4,
	"CANNOT_CREATE_EXPERIMENT_FOR_REMOVED_BASE_CAMPAIGN":   5,
	"CANNOT_CREATE_EXPERIMENT_FOR_NON_PROPOSED_DRAFT":      6,
	"CUSTOMER_CANNOT_CREATE_EXPERIMENT":                    7,
	"CAMPAIGN_CANNOT_CREATE_EXPERIMENT":                    8,
	"EXPERIMENT_DURATIONS_MUST_NOT_OVERLAP":                9,
	"EXPERIMENT_DURATION_MUST_BE_WITHIN_CAMPAIGN_DURATION": 10,
	"CANNOT_MUTATE_EXPERIMENT_DUE_TO_STATUS":               11,
}

func (x CampaignExperimentErrorEnum_CampaignExperimentError) String() string {
	return proto.EnumName(CampaignExperimentErrorEnum_CampaignExperimentError_name, int32(x))
}

func (CampaignExperimentErrorEnum_CampaignExperimentError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_56a23521961f4d1f, []int{0, 0}
}

// Container for enum describing possible campaign experiment errors.
type CampaignExperimentErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CampaignExperimentErrorEnum) Reset()         { *m = CampaignExperimentErrorEnum{} }
func (m *CampaignExperimentErrorEnum) String() string { return proto.CompactTextString(m) }
func (*CampaignExperimentErrorEnum) ProtoMessage()    {}
func (*CampaignExperimentErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_56a23521961f4d1f, []int{0}
}

func (m *CampaignExperimentErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignExperimentErrorEnum.Unmarshal(m, b)
}
func (m *CampaignExperimentErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignExperimentErrorEnum.Marshal(b, m, deterministic)
}
func (m *CampaignExperimentErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignExperimentErrorEnum.Merge(m, src)
}
func (m *CampaignExperimentErrorEnum) XXX_Size() int {
	return xxx_messageInfo_CampaignExperimentErrorEnum.Size(m)
}
func (m *CampaignExperimentErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignExperimentErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignExperimentErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.errors.CampaignExperimentErrorEnum_CampaignExperimentError", CampaignExperimentErrorEnum_CampaignExperimentError_name, CampaignExperimentErrorEnum_CampaignExperimentError_value)
	proto.RegisterType((*CampaignExperimentErrorEnum)(nil), "google.ads.googleads.v1.errors.CampaignExperimentErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/errors/campaign_experiment_error.proto", fileDescriptor_56a23521961f4d1f)
}

var fileDescriptor_56a23521961f4d1f = []byte{
	// 502 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xd1, 0x8b, 0xd3, 0x30,
	0x1c, 0xc7, 0xdd, 0x76, 0xde, 0x69, 0x06, 0x1a, 0xf2, 0xa0, 0xa0, 0xc7, 0x81, 0x83, 0x13, 0x54,
	0x68, 0x99, 0x27, 0x22, 0x15, 0x84, 0xac, 0xfd, 0x6d, 0x17, 0x5c, 0x93, 0x92, 0xa6, 0x3d, 0x91,
	0x41, 0xa8, 0xb7, 0x52, 0x06, 0xb7, 0x76, 0xb4, 0xf3, 0xf0, 0xd1, 0xbf, 0xc5, 0x47, 0xff, 0x14,
	0xff, 0x0a, 0x9f, 0xfd, 0x07, 0x7c, 0x95, 0x36, 0x5b, 0x11, 0xb1, 0xf7, 0xd4, 0x1f, 0xcd, 0xe7,
	0x93, 0x6f, 0x12, 0xbe, 0xe8, 0x5d, 0x56, 0x14, 0xd9, 0x55, 0x6a, 0x27, 0xcb, 0xca, 0x36, 0x63,
	0x3d, 0x5d, 0x8f, 0xed, 0xb4, 0x2c, 0x8b, 0xb2, 0xb2, 0x2f, 0x93, 0xf5, 0x26, 0x59, 0x65, 0xb9,
	0x4e, 0xbf, 0x6c, 0xd2, 0x72, 0xb5, 0x4e, 0xf3, 0xad, 0x6e, 0x96, 0xac, 0x4d, 0x59, 0x6c, 0x0b,
	0x72, 0x62, 0x24, 0x2b, 0x59, 0x56, 0x56, 0xeb, 0x5b, 0xd7, 0x63, 0xcb, 0xf8, 0x8f, 0x8e, 0xf7,
	0xfb, 0x6f, 0x56, 0x76, 0x92, 0xe7, 0xc5, 0x36, 0xd9, 0xae, 0x8a, 0xbc, 0x32, 0xf6, 0xe8, 0xeb,
	0x01, 0x7a, 0xec, 0xee, 0x12, 0xa0, 0x0d, 0x80, 0x5a, 0x85, 0xfc, 0xf3, 0x7a, 0xf4, 0x73, 0x80,
	0x1e, 0x76, 0xac, 0x93, 0xfb, 0x68, 0x18, 0xf1, 0x30, 0x00, 0x97, 0x4d, 0x19, 0x78, 0xf8, 0x16,
	0x19, 0xa2, 0xa3, 0x88, 0xbf, 0xe7, 0xe2, 0x82, 0xe3, 0x1e, 0x21, 0xe8, 0x9e, 0x17, 0x05, 0x73,
	0xe6, 0x52, 0x05, 0x9a, 0x53, 0x1f, 0x70, 0x9f, 0x3c, 0x40, 0x84, 0xf1, 0x98, 0xce, 0x99, 0xa7,
	0x95, 0xa4, 0x3c, 0x64, 0x8a, 0x09, 0x8e, 0x07, 0xc4, 0x46, 0x2f, 0x5c, 0xca, 0xb9, 0x50, 0xda,
	0x95, 0x50, 0xf3, 0xf0, 0x21, 0x00, 0xc9, 0x7c, 0xe0, 0x4a, 0x5f, 0x30, 0x75, 0xae, 0xc3, 0x73,
	0x2a, 0xc1, 0xd3, 0x93, 0xc8, 0x9b, 0x81, 0xc2, 0x07, 0xe4, 0x35, 0x7a, 0xd9, 0x29, 0x4c, 0x85,
	0xd4, 0x12, 0x7c, 0x11, 0xd7, 0x02, 0x0d, 0x41, 0xbb, 0xd4, 0x0f, 0x28, 0x9b, 0x71, 0x7c, 0x9b,
	0x9c, 0x21, 0xfb, 0x46, 0x8f, 0x0b, 0xae, 0x03, 0x29, 0x02, 0x11, 0x82, 0xa7, 0x3d, 0x49, 0xa7,
	0x0a, 0x1f, 0x92, 0x53, 0xf4, 0xc4, 0x8d, 0x42, 0x25, 0x7c, 0x90, 0xba, 0xcb, 0xc6, 0x47, 0x0d,
	0xb6, 0x4b, 0xea, 0xc6, 0xee, 0x90, 0x67, 0xe8, 0xf4, 0xaf, 0x50, 0x2f, 0x92, 0xb4, 0x7e, 0x84,
	0x50, 0xfb, 0x51, 0xa8, 0x74, 0x2d, 0x89, 0x18, 0xe4, 0x9c, 0x06, 0xf8, 0x2e, 0x79, 0x83, 0x5e,
	0xfd, 0x07, 0x35, 0xe4, 0x04, 0x9a, 0x97, 0x61, 0xbc, 0xbd, 0x5e, 0x0b, 0x60, 0x44, 0x9e, 0xa3,
	0xa7, 0xbb, 0x23, 0xf8, 0x91, 0xfa, 0xe7, 0x9e, 0x5e, 0x04, 0x5a, 0x09, 0x1d, 0x2a, 0xaa, 0xa2,
	0x10, 0x0f, 0x27, 0xbf, 0x7b, 0x68, 0x74, 0x59, 0xac, 0xad, 0x9b, 0x7b, 0x34, 0x39, 0xee, 0xa8,
	0x41, 0x50, 0xf7, 0x28, 0xe8, 0x7d, 0xf4, 0x76, 0x7e, 0x56, 0x5c, 0x25, 0x79, 0x66, 0x15, 0x65,
	0x66, 0x67, 0x69, 0xde, 0xb4, 0x6c, 0xdf, 0xeb, 0xcd, 0xaa, 0xea, 0xaa, 0xf9, 0x5b, 0xf3, 0xf9,
	0xd6, 0x1f, 0xcc, 0x28, 0xfd, 0xde, 0x3f, 0x99, 0x99, 0xcd, 0xe8, 0xb2, 0xb2, 0xcc, 0x58, 0x4f,
	0xf1, 0xd8, 0x6a, 0x22, 0xab, 0x1f, 0x7b, 0x60, 0x41, 0x97, 0xd5, 0xa2, 0x05, 0x16, 0xf1, 0x78,
	0x61, 0x80, 0x5f, 0xfd, 0x91, 0xf9, 0xeb, 0x38, 0x74, 0x59, 0x39, 0x4e, 0x8b, 0x38, 0x4e, 0x3c,
	0x76, 0x1c, 0x03, 0x7d, 0x3a, 0x6c, 0x4e, 0x77, 0xf6, 0x27, 0x00, 0x00, 0xff, 0xff, 0x81, 0x55,
	0xf7, 0xc3, 0x83, 0x03, 0x00, 0x00,
}
