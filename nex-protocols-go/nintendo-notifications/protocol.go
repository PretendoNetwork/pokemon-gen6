// Package protocol implements the Nintendo Notfications protocol
package protocol

import nex "github.com/PretendoNetwork/nex-go/v2"

const (
	// ProtocolID is the protocol ID for the Nintendo Notifications protocol
	ProtocolID = 0x64

	// MethodProcessNintendoNotificationEvent1 is the method ID for the method ProcessNintendoNotificationEvent (1)
	MethodProcessNintendoNotificationEvent1 = 0x1

	// MethodProcessNintendoNotificationEvent2 is the method ID for the method ProcessNintendoNotificationEvent (2)
	MethodProcessNintendoNotificationEvent2 = 0x2
)

// Protocol handles the NintendoNotifications protocol
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

// NewProtocol returns a new Nintendo Notifications protocol
func NewProtocol() *Protocol {
	return &Protocol{}
}
