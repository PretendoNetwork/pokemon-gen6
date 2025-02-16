// Package protocol implements the ServiceItemTeamKirbyClashDeluxe protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	service_item "github.com/PretendoNetwork/nex-protocols-go/v2/service-item"
	service_item_team_kirby_clash_deluxe_types "github.com/PretendoNetwork/nex-protocols-go/v2/service-item/team-kirby-clash-deluxe/types"
	"golang.org/x/exp/slices"
)

const (
	// ProtocolID is the Protocol ID for the Service Item (Team Kirby Clash Deluxe) protocol
	ProtocolID = 0x77

	// MethodGetEnvironment is the method ID for the GetEnvironment method
	MethodGetEnvironment = 0x1

	// MethodHTTPGetRequest is the method ID for the HTTPGetRequest method
	MethodHTTPGetRequest = 0x2

	// MethodHTTPGetResponse is the method ID for the HTTPGetResponse method
	MethodHTTPGetResponse = 0x3

	// MethodPurchaseServiceItemRequest is the method ID for the PurchaseServiceItemRequest method
	MethodPurchaseServiceItemRequest = 0x4

	// MethodPurchaseServiceItemResponse is the method ID for the PurchaseServiceItemResponse method
	MethodPurchaseServiceItemResponse = 0x5

	// MethodListServiceItemRequest is the method ID for the ListServiceItemRequest method
	MethodListServiceItemRequest = 0x6

	// MethodListServiceItemResponse is the method ID for the ListServiceItemResponse method
	MethodListServiceItemResponse = 0x7

	// MethodGetBalanceRequest is the method ID for the GetBalanceRequest method
	MethodGetBalanceRequest = 0x8

	// MethodGetBalanceResponse is the method ID for the GetBalanceResponse method
	MethodGetBalanceResponse = 0x9

	// MethodGetPrepurchaseInfoRequest is the method ID for the GetPrepurchaseInfoRequest method
	MethodGetPrepurchaseInfoRequest = 0xA

	// MethodGetPrepurchaseInfoResponse is the method ID for the GetPrepurchaseInfoResponse method
	MethodGetPrepurchaseInfoResponse = 0xB

	// MethodGetServiceItemRightRequest is the method ID for the GetServiceItemRightRequest method
	MethodGetServiceItemRightRequest = 0xC

	// MethodGetServiceItemRightResponse is the method ID for the GetServiceItemRightResponse method
	MethodGetServiceItemRightResponse = 0xD

	// MethodGetPurchaseHistoryRequest is the method ID for the GetPurchaseHistoryRequest method
	MethodGetPurchaseHistoryRequest = 0xE

	// MethodGetPurchaseHistoryResponse is the method ID for the GetPurchaseHistoryResponse method
	MethodGetPurchaseHistoryResponse = 0xF

	// MethodPostRightBinaryByAccount is the method ID for the PostRightBinaryByAccount method
	MethodPostRightBinaryByAccount = 0x10

	// MethodUseServiceItemByAccountRequest is the method ID for the UseServiceItemByAccountRequest method
	MethodUseServiceItemByAccountRequest = 0x11

	// MethodUseServiceItemByAccountResponse is the method ID for the UseServiceItemByAccountResponse method
	MethodUseServiceItemByAccountResponse = 0x12

	// MethodAcquireServiceItemByAccount is the method ID for the AcquireServiceItemByAccount method
	MethodAcquireServiceItemByAccount = 0x13

	// MethodGetSupportID is the method ID for the GetSupportID method
	MethodGetSupportID = 0x14

	// MethodGetLawMessageRequest is the method ID for the GetLawMessageRequest method
	MethodGetLawMessageRequest = 0x15

	// MethodGetLawMessageResponse is the method ID for the GetLawMessageResponse method
	MethodGetLawMessageResponse = 0x16
)

var patchedMethods = []uint32{
	MethodGetEnvironment,
	MethodHTTPGetRequest,
	MethodHTTPGetResponse,
	MethodPurchaseServiceItemRequest,
	MethodPurchaseServiceItemResponse,
	MethodListServiceItemRequest,
	MethodListServiceItemResponse,
	MethodGetBalanceRequest,
	MethodGetBalanceResponse,
	MethodGetPrepurchaseInfoRequest,
	MethodGetPrepurchaseInfoResponse,
	MethodGetServiceItemRightRequest,
	MethodGetServiceItemRightResponse,
	MethodGetPurchaseHistoryRequest,
	MethodGetPurchaseHistoryResponse,
	MethodPostRightBinaryByAccount,
	MethodUseServiceItemByAccountRequest,
	MethodUseServiceItemByAccountResponse,
	MethodAcquireServiceItemByAccount,
	MethodGetSupportID,
	MethodGetLawMessageRequest,
	MethodGetLawMessageResponse,
}

type serviceItemProtocol = service_item.Protocol

// Protocol stores all the RMC method handlers for the Service Item (Team Kirby Clash Deluxe) protocol and listens for requests
// Embeds the Service Item protocol
type Protocol struct {
	endpoint nex.EndpointInterface
	serviceItemProtocol
	GetEnvironment                  func(err error, packet nex.PacketInterface, callID uint32, uniqueID *types.String, platform *types.PrimitiveU8) (*nex.RMCMessage, *nex.Error)
	HTTPGetRequest                  func(err error, packet nex.PacketInterface, callID uint32, url *service_item_team_kirby_clash_deluxe_types.ServiceItemHTTPGetParam) (*nex.RMCMessage, *nex.Error)
	HTTPGetResponse                 func(err error, packet nex.PacketInterface, callID uint32, requestID *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	PurchaseServiceItemRequest      func(err error, packet nex.PacketInterface, callID uint32, purchaseServiceItemParam *service_item_team_kirby_clash_deluxe_types.ServiceItemPurchaseServiceItemParam) (*nex.RMCMessage, *nex.Error)
	PurchaseServiceItemResponse     func(err error, packet nex.PacketInterface, callID uint32, requestID *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	ListServiceItemRequest          func(err error, packet nex.PacketInterface, callID uint32, listServiceItemParam *service_item_team_kirby_clash_deluxe_types.ServiceItemListServiceItemParam) (*nex.RMCMessage, *nex.Error)
	ListServiceItemResponse         func(err error, packet nex.PacketInterface, callID uint32, requestID *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	GetBalanceRequest               func(err error, packet nex.PacketInterface, callID uint32, getBalanceParam *service_item_team_kirby_clash_deluxe_types.ServiceItemGetBalanceParam) (*nex.RMCMessage, *nex.Error)
	GetBalanceResponse              func(err error, packet nex.PacketInterface, callID uint32, requestID *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	GetPrepurchaseInfoRequest       func(err error, packet nex.PacketInterface, callID uint32, getPrepurchaseInfoParam *service_item_team_kirby_clash_deluxe_types.ServiceItemGetPrepurchaseInfoParam) (*nex.RMCMessage, *nex.Error)
	GetPrepurchaseInfoResponse      func(err error, packet nex.PacketInterface, callID uint32, requestID *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	GetServiceItemRightRequest      func(err error, packet nex.PacketInterface, callID uint32, getServiceItemRightParam *service_item_team_kirby_clash_deluxe_types.ServiceItemGetServiceItemRightParam, withoutRightBinary *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error)
	GetServiceItemRightResponse     func(err error, packet nex.PacketInterface, callID uint32, requestID *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	GetPurchaseHistoryRequest       func(err error, packet nex.PacketInterface, callID uint32, getPurchaseHistoryParam *service_item_team_kirby_clash_deluxe_types.ServiceItemGetPurchaseHistoryParam) (*nex.RMCMessage, *nex.Error)
	GetPurchaseHistoryResponse      func(err error, packet nex.PacketInterface, callID uint32, requestID *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	PostRightBinaryByAccount        func(err error, packet nex.PacketInterface, callID uint32, postRightBinaryByAccountParam *service_item_team_kirby_clash_deluxe_types.ServiceItemPostRightBinaryByAccountParam) (*nex.RMCMessage, *nex.Error)
	UseServiceItemByAccountRequest  func(err error, packet nex.PacketInterface, callID uint32, useServiceItemByAccountParam *service_item_team_kirby_clash_deluxe_types.ServiceItemUseServiceItemByAccountParam) (*nex.RMCMessage, *nex.Error)
	UseServiceItemByAccountResponse func(err error, packet nex.PacketInterface, callID uint32, requestID *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	AcquireServiceItemByAccount     func(err error, packet nex.PacketInterface, callID uint32, acquireServiceItemByAccountParam *service_item_team_kirby_clash_deluxe_types.ServiceItemAcquireServiceItemByAccountParam) (*nex.RMCMessage, *nex.Error)
	GetSupportID                    func(err error, packet nex.PacketInterface, callID uint32, getSuppordIDParam *service_item_team_kirby_clash_deluxe_types.ServiceItemGetSupportIDParam) (*nex.RMCMessage, *nex.Error)
	GetLawMessageRequest            func(err error, packet nex.PacketInterface, callID uint32, getLawMessageParam *service_item_team_kirby_clash_deluxe_types.ServiceItemGetLawMessageParam) (*nex.RMCMessage, *nex.Error)
	GetLawMessageResponse           func(err error, packet nex.PacketInterface, callID uint32, requestID *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	message := packet.RMCMessage()

	if !message.IsRequest || message.ProtocolID != ProtocolID {
		return
	}

	if !slices.Contains(patchedMethods, message.MethodID) {
		protocol.serviceItemProtocol.HandlePacket(packet)
		return
	}

	switch message.MethodID {
	case MethodGetEnvironment:
		protocol.handleGetEnvironment(packet)
	case MethodHTTPGetRequest:
		protocol.handleHTTPGetRequest(packet)
	case MethodHTTPGetResponse:
		protocol.handleHTTPGetResponse(packet)
	case MethodPurchaseServiceItemRequest:
		protocol.handlePurchaseServiceItemRequest(packet)
	case MethodPurchaseServiceItemResponse:
		protocol.handlePurchaseServiceItemResponse(packet)
	case MethodListServiceItemRequest:
		protocol.handleListServiceItemRequest(packet)
	case MethodListServiceItemResponse:
		protocol.handleListServiceItemResponse(packet)
	case MethodGetBalanceRequest:
		protocol.handleGetBalanceRequest(packet)
	case MethodGetBalanceResponse:
		protocol.handleGetBalanceResponse(packet)
	case MethodGetPrepurchaseInfoRequest:
		protocol.handleGetPrepurchaseInfoRequest(packet)
	case MethodGetPrepurchaseInfoResponse:
		protocol.handleGetPrepurchaseInfoResponse(packet)
	case MethodGetServiceItemRightRequest:
		protocol.handleGetServiceItemRightRequest(packet)
	case MethodGetServiceItemRightResponse:
		protocol.handleGetServiceItemRightResponse(packet)
	case MethodGetPurchaseHistoryRequest:
		protocol.handleGetPurchaseHistoryRequest(packet)
	case MethodGetPurchaseHistoryResponse:
		protocol.handleGetPurchaseHistoryResponse(packet)
	case MethodPostRightBinaryByAccount:
		protocol.handlePostRightBinaryByAccount(packet)
	case MethodUseServiceItemByAccountRequest:
		protocol.handleUseServiceItemByAccountRequest(packet)
	case MethodUseServiceItemByAccountResponse:
		protocol.handleUseServiceItemByAccountResponse(packet)
	case MethodAcquireServiceItemByAccount:
		protocol.handleAcquireServiceItemByAccount(packet)
	case MethodGetSupportID:
		protocol.handleGetSupportID(packet)
	case MethodGetLawMessageRequest:
		protocol.handleGetLawMessageRequest(packet)
	case MethodGetLawMessageResponse:
		protocol.handleGetLawMessageResponse(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported Service Item (Team Kirby Clash Deluxe) method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new ServiceItemTeamKirbyClashDeluxe protocol
func NewProtocol(endpoint nex.EndpointInterface) *Protocol {
	protocol := &Protocol{endpoint: endpoint}
	protocol.serviceItemProtocol.SetEndpoint(endpoint)

	return protocol
}
