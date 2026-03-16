package recording

import (
	"github.com/duwi2024/onvif/xsd"
)

// GetServiceCapabilities action
type GetServiceCapabilities struct {
	XMLName string `xml:"trc:GetServiceCapabilities"`
}

// GetServiceCapabilitiesResponse type
type GetServiceCapabilitiesResponse struct {
	Capabilities Capabilities
}

type Capabilities struct { //tev
	DynamicRecordings xsd.Boolean `xml:"DynamicRecordings"`
}

type GetRecordings struct {
	XMLName string `xml:"trc:GetRecordings"`
}

type GetRecordingsResponse struct {
	RecordingItem []RecordingItemEntity
}

type RecordingItemEntity struct {
	RecordingToken string `xml:"RecordingToken"`
}

type GetRecordingJobs struct {
	XMLName string `xml:"trc:GetRecordingJobs"`
}

type GetRecordingJobsResponse struct {
	JobItem []RecordingJobItem
}

type RecordingJobItem struct {
	JobToken string `xml:"JobToken"`
}

type GetRecordingJobConfiguration struct {
	XMLName  string `xml:"trc:GetRecordingJobConfiguration"`
	JobToken string `xml:"JobToken"`
}

type GetRecordingJobConfigurationResponse struct {
	JobConfiguration RecordingJobConfiguration
}

type RecordingJobConfiguration struct {
	ScheduleToken  string
	RecordingToken string
}

type GetRecordingConfiguration struct {
}

type GetRecordingConfigurationResponse struct {
}

type GetRecordingSummary struct {
	XMLName string `xml:"trc:GetRecordingSummary"`
}

type GetRecordingSummaryResponse struct {
}
