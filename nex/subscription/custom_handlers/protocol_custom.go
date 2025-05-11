package nex_subscription_custom_handlers

import (
	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	subscription_types "github.com/PretendoNetwork/pokemon-gen6/nex/subscription/types"
)

const (
	ProtocolID = 0x75

	// MethodCreateMySubscriptionData is the method ID for the method CreateMySubscriptionDataID
	MethodCreateMySubscriptionData = 0x1

	// MethodUpdateMySubscriptionData is the method ID for the method UpdateMySubscriptionData
	MethodUpdateMySubscriptionData = 0x2
)

// Protocol handles the Subscription nex protocol
type Protocol struct {
	endpoint                 nex.EndpointInterface
	CreateMySubscriptionData func(err error, packet nex.PacketInterface, callID uint32, unk types.UInt32, param subscription_types.SubscriptionData) (*nex.RMCMessage, *nex.Error)
	UpdateMySubscriptionData func(err error, packet nex.PacketInterface, callID uint32, param subscription_types.SubscriptionData) (*nex.RMCMessage, *nex.Error)
}

// Endpoint returns the endpoint implementing the protocol
func (protocol *Protocol) Endpoint() nex.EndpointInterface {
	return protocol.endpoint
}

// SetEndpoint sets the endpoint implementing the protocol
func (protocol *Protocol) SetEndpoint(endpoint nex.EndpointInterface) {
	protocol.endpoint = endpoint
}

// SetHandlerCreateMySubscriptionData sets the handler for the CreateMySubscriptionData method
func (protocol *Protocol) SetHandlerCreateMySubscriptionData(handler func(err error, packet nex.PacketInterface, callID uint32, unk types.UInt32, param subscription_types.SubscriptionData) (*nex.RMCMessage, *nex.Error)) {
	protocol.CreateMySubscriptionData = handler
}

// SetHandlerUpdateMySubscriptionData sets the handler for the UpdateMySubscriptionData method
func (protocol *Protocol) SetHandlerUpdateMySubscriptionData(handler func(err error, packet nex.PacketInterface, callID uint32, param subscription_types.SubscriptionData) (*nex.RMCMessage, *nex.Error)) {
	protocol.UpdateMySubscriptionData = handler
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	message := packet.RMCMessage()

	switch message.MethodID {
	case MethodCreateMySubscriptionData:
		protocol.handleCreateMySubscriptionData(packet)
	case MethodUpdateMySubscriptionData:
		protocol.handleUpdateMySubscriptionData(packet)
	}
}

// NewProtocol returns a new Protocol
func NewProtocol() *Protocol {
	return &Protocol{}
}
