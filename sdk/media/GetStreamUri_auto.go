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

// Call_GetStreamUri forwards the call to dev.CallMethod() then parses the payload of the reply as a GetStreamUriResponse.
func Call_GetStreamUri(ctx context.Context, dev *onvif.Device, request media.GetStreamUri) (media.GetStreamUriResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			Fault *xsdOnvif.Fault
			GetStreamUriResponse media.GetStreamUriResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetStreamUriResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetStreamUri")
		return reply.Body.GetStreamUriResponse, errors.Annotate(err, "reply")
	}
}
