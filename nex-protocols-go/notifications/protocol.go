// Package protocol implements the Notfications protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go/v2"
)

const (
	// ProtocolID is the protocol ID for the Nintendo Notifications protocol
	ProtocolID = 0xE

	// MethodProcessNotificationEvent is the method ID for the method ProcessNotificationEvent
	MethodProcessNotificationEvent = 0x1
)

// Protocol handles the Notifications protocol
type Protocol struct {
	endpoint       nex.EndpointInterface
	Patches        nex.ServiceProtocol
	PatchedMethods []uint32
}

// Endpoint returns the endpoint implementing the protocol
func (protocol *Protocol) Endpoint() nex.EndpointInterface {
	return protocol.endpoint
}

// SetEndpoint sets the endpoint implementing the protocol
func (protocol *Protocol) SetEndpoint(endpoint nex.EndpointInterface) {
	protocol.endpoint = endpoint
}

// NewProtocol returns a new Notifications protocol
func NewProtocol() *Protocol {
	return &Protocol{}
}
