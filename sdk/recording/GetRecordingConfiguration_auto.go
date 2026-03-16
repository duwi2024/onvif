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

// Call_GetRecordingConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a GetRecordingConfigurationResponse.
func Call_GetRecordingConfiguration(ctx context.Context, dev *onvif.Device, request recording.GetRecordingConfiguration) (recording.GetRecordingConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetRecordingConfigurationResponse recording.GetRecordingConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetRecordingConfigurationResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetRecordingConfiguration")
		return reply.Body.GetRecordingConfigurationResponse, errors.Annotate(err, "reply")
	}
}
