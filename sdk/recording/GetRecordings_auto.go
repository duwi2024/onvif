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

// Call_GetRecordings forwards the call to dev.CallMethod() then parses the payload of the reply as a GetRecordingsResponse.
func Call_GetRecordings(ctx context.Context, dev *onvif.Device, request recording.GetRecordings) (recording.GetRecordingsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetRecordingsResponse recording.GetRecordingsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetRecordingsResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetRecordings")
		return reply.Body.GetRecordingsResponse, errors.Annotate(err, "reply")
	}
}
