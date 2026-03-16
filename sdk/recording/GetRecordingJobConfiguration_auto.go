// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package recording

import (
	"context"
	"github.com/juju/errors"
	"github.com/duwi2024/onvif"
	"github.com/duwi2024/onvif/sdk"
	"github.com/duwi2024/onvif/recording"
)

// Call_GetRecordingJobConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a GetRecordingJobConfigurationResponse.
func Call_GetRecordingJobConfiguration(ctx context.Context, dev *onvif.Device, request recording.GetRecordingJobConfiguration) (recording.GetRecordingJobConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetRecordingJobConfigurationResponse recording.GetRecordingJobConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetRecordingJobConfigurationResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetRecordingJobConfiguration")
		return reply.Body.GetRecordingJobConfigurationResponse, errors.Annotate(err, "reply")
	}
}
