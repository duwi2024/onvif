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

// Call_CreatePullPointSubscription forwards the call to dev.CallMethod() then parses the payload of the reply as a CreatePullPointSubscriptionResponse.
func Call_CreatePullPointSubscription(ctx context.Context, dev *onvif.Device, request event.CreatePullPointSubscription) (event.CreatePullPointSubscriptionResponse, *xsdOnvif.Fault,error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			Fault *xsdOnvif.Fault
			CreatePullPointSubscriptionResponse event.CreatePullPointSubscriptionResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.CreatePullPointSubscriptionResponse, nil,errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "CreatePullPointSubscription")
		return reply.Body.CreatePullPointSubscriptionResponse,reply.Body.Fault , errors.Annotate(err, "reply")
	}
}
