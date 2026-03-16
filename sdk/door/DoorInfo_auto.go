// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package door

import (
	"context"
	"github.com/juju/errors"
	"github.com/duwi2024/onvif"
	"github.com/duwi2024/onvif/sdk"
	"github.com/duwi2024/onvif/door"
)

// Call_DoorInfo forwards the call to dev.CallMethod() then parses the payload of the reply as a DoorInfoResponse.
func Call_DoorInfo(ctx context.Context, dev *onvif.Device, request door.DoorInfo) (door.DoorInfoResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			DoorInfoResponse door.DoorInfoResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.DoorInfoResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "DoorInfo")
		return reply.Body.DoorInfoResponse, errors.Annotate(err, "reply")
	}
}
