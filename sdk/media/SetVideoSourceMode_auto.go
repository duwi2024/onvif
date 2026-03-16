// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package media

import (
	"context"
	"github.com/juju/errors"
	"github.com/duwi2024/onvif"
	"github.com/duwi2024/onvif/sdk"
	"github.com/duwi2024/onvif/media"
	xsdOnvif "github.com/duwi2024/onvif/xsd/onvif"
)

// Call_SetVideoSourceMode forwards the call to dev.CallMethod() then parses the payload of the reply as a SetVideoSourceModeResponse.
func Call_SetVideoSourceMode(ctx context.Context, dev *onvif.Device, request media.SetVideoSourceMode) (media.SetVideoSourceModeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			Fault *xsdOnvif.Fault
			SetVideoSourceModeResponse media.SetVideoSourceModeResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetVideoSourceModeResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "SetVideoSourceMode")
		return reply.Body.SetVideoSourceModeResponse, errors.Annotate(err, "reply")
	}
}
