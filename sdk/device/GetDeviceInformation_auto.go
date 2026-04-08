// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package device

import (
	"context"
	"github.com/duwi2024/onvif"
	"github.com/duwi2024/onvif/device"
	"github.com/duwi2024/onvif/sdk"
	xsdOnvif "github.com/duwi2024/onvif/xsd/onvif"
	"github.com/juju/errors"
)

// Call_GetDeviceInformation forwards the call to dev.CallMethod() then parses the payload of the reply as a GetDeviceInformationResponse.
func Call_GetDeviceInformation(ctx context.Context, dev *onvif.Device, request device.GetDeviceInformation) (device.GetDeviceInformationResponse, *xsdOnvif.Fault, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			Fault                        *xsdOnvif.Fault
			GetDeviceInformationResponse device.GetDeviceInformationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetDeviceInformationResponse, nil, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetDeviceInformation")
		return reply.Body.GetDeviceInformationResponse, reply.Body.Fault, errors.Annotate(err, "reply")
	}
}
