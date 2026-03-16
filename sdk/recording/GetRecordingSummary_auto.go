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
	xsdOnvif "github.com/duwi2024/onvif/xsd/onvif"
)

// Call_GetRecordingSummary forwards the call to dev.CallMethod() then parses the payload of the reply as a GetRecordingSummaryResponse.
func Call_GetRecordingSummary(ctx context.Context, dev *onvif.Device, request recording.GetRecordingSummary) (recording.GetRecordingSummaryResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			Fault *xsdOnvif.Fault
			GetRecordingSummaryResponse recording.GetRecordingSummaryResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetRecordingSummaryResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetRecordingSummary")
		return reply.Body.GetRecordingSummaryResponse, errors.Annotate(err, "reply")
	}
}
