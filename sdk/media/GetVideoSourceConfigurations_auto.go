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

// Call_GetVideoSourceConfigurations forwards the call to dev.CallMethod() then parses the payload of the reply as a GetVideoSourceConfigurationsResponse.
func Call_GetVideoSourceConfigurations(ctx context.Context, dev *onvif.Device, request media.GetVideoSourceConfigurations) (media.GetVideoSourceConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			Fault *xsdOnvif.Fault
			GetVideoSourceConfigurationsResponse media.GetVideoSourceConfigurationsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetVideoSourceConfigurationsResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetVideoSourceConfigurations")
		return reply.Body.GetVideoSourceConfigurationsResponse, errors.Annotate(err, "reply")
	}
}
