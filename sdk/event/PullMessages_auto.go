// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package event

import (
	"context"
	"github.com/juju/errors"
	"github.com/duwi2024/onvif"
	"github.com/duwi2024/onvif/sdk"
	"github.com/duwi2024/onvif/event"
	xsdOnvif "github.com/duwi2024/onvif/xsd/onvif"
)

// Call_PullMessages forwards the call to dev.CallMethod() then parses the payload of the reply as a PullMessagesResponse.
func Call_PullMessages(ctx context.Context, dev *onvif.Device, request event.PullMessages) (event.PullMessagesResponse, *xsdOnvif.Fault,error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			Fault *xsdOnvif.Fault
			PullMessagesResponse event.PullMessagesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.PullMessagesResponse, nil,errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "PullMessages")
		return reply.Body.PullMessagesResponse,reply.Body.Fault , errors.Annotate(err, "reply")
	}
}
