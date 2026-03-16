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

// Call_AddEventBroker forwards the call to dev.CallMethod() then parses the payload of the reply as a AddEventBrokerResponse.
func Call_AddEventBroker(ctx context.Context, dev *onvif.Device, request event.AddEventBroker) (event.AddEventBrokerResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			Fault *xsdOnvif.Fault
			AddEventBrokerResponse event.AddEventBrokerResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.AddEventBrokerResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "AddEventBroker")
		return reply.Body.AddEventBrokerResponse, errors.Annotate(err, "reply")
	}
}
