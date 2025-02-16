// Package protocol implements the ServiceItemWiiSportsClub protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	service_item "github.com/PretendoNetwork/nex-protocols-go/v2/service-item"
	service_item_wii_sports_club_types "github.com/PretendoNetwork/nex-protocols-go/v2/service-item/wii-sports-club/types"
	"golang.org/x/exp/slices"
)

const (
	// ProtocolID is the Protocol ID for the Service Item (Wii Sports Club) protocol
	ProtocolID = 0x77

	// MethodHello is the method ID for the Hello method
	MethodHello = 0x1

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

	// MethodGetNotice is the method ID for the GetNotice method
	MethodGetNotice = 0x10

	// MethodUpdateAndGetTicketInfo is the method ID for the UpdateAndGetTicketInfo method
	MethodUpdateAndGetTicketInfo = 0x11

	// MethodLoadUserInfo is the method ID for the LoadUserInfo method
	MethodLoadUserInfo = 0x12

	// MethodSaveUserInfo is the method ID for the SaveUserInfo method
	MethodSaveUserInfo = 0x13

	// MethodStartChallenge is the method ID for the StartChallenge method
	MethodStartChallenge = 0x14

	// MethodEndChallenge is the method ID for the EndChallenge method
	MethodEndChallenge = 0x15

	// MethodRequestTicketRestoration is the method ID for the RequestTicketRestoration method
	MethodRequestTicketRestoration = 0x16
)

var patchedMethods = []uint32{
	MethodHello,
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
	MethodGetNotice,
	MethodUpdateAndGetTicketInfo,
	MethodLoadUserInfo,
	MethodSaveUserInfo,
	MethodStartChallenge,
	MethodEndChallenge,
	MethodRequestTicketRestoration,
}

type serviceItemProtocol = service_item.Protocol

// Protocol stores all the RMC method handlers for the Service Item (Wii Sports Club) protocol and listens for requests
// Embeds the Service Item protocol
type Protocol struct {
	endpoint nex.EndpointInterface
	serviceItemProtocol
	Hello                       func(err error, packet nex.PacketInterface, callID uint32, name *types.String) (*nex.RMCMessage, *nex.Error)
	HTTPGetRequest              func(err error, packet nex.PacketInterface, callID uint32, url *service_item_wii_sports_club_types.ServiceItemHTTPGetParam) (*nex.RMCMessage, *nex.Error)
	HTTPGetResponse             func(err error, packet nex.PacketInterface, callID uint32, requestID *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	PurchaseServiceItemRequest  func(err error, packet nex.PacketInterface, callID uint32, purchaseServiceItemParam *service_item_wii_sports_club_types.ServiceItemPurchaseServiceItemParam) (*nex.RMCMessage, *nex.Error)
	PurchaseServiceItemResponse func(err error, packet nex.PacketInterface, callID uint32, requestID *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	ListServiceItemRequest      func(err error, packet nex.PacketInterface, callID uint32, listServiceItemParam *service_item_wii_sports_club_types.ServiceItemListServiceItemParam) (*nex.RMCMessage, *nex.Error)
	ListServiceItemResponse     func(err error, packet nex.PacketInterface, callID uint32, requestID *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	GetBalanceRequest           func(err error, packet nex.PacketInterface, callID uint32, getBalanceParam *service_item_wii_sports_club_types.ServiceItemGetBalanceParam) (*nex.RMCMessage, *nex.Error)
	GetBalanceResponse          func(err error, packet nex.PacketInterface, callID uint32, requestID *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	GetPrepurchaseInfoRequest   func(err error, packet nex.PacketInterface, callID uint32, getPrepurchaseInfoParam *service_item_wii_sports_club_types.ServiceItemGetPrepurchaseInfoParam) (*nex.RMCMessage, *nex.Error)
	GetPrepurchaseInfoResponse  func(err error, packet nex.PacketInterface, callID uint32, requestID *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	GetServiceItemRightRequest  func(err error, packet nex.PacketInterface, callID uint32, getServiceItemRightParam *service_item_wii_sports_club_types.ServiceItemGetServiceItemRightParam) (*nex.RMCMessage, *nex.Error)
	GetServiceItemRightResponse func(err error, packet nex.PacketInterface, callID uint32, requestID *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	GetPurchaseHistoryRequest   func(err error, packet nex.PacketInterface, callID uint32, getPurchaseHistoryParam *service_item_wii_sports_club_types.ServiceItemGetPurchaseHistoryParam) (*nex.RMCMessage, *nex.Error)
	GetPurchaseHistoryResponse  func(err error, packet nex.PacketInterface, callID uint32, requestID *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	GetNotice                   func(err error, packet nex.PacketInterface, callID uint32, getNoticeParam *service_item_wii_sports_club_types.ServiceItemGetNoticeParam) (*nex.RMCMessage, *nex.Error)
	UpdateAndGetTicketInfo      func(err error, packet nex.PacketInterface, callID uint32, forceRetrieveFromEShop *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error)
	LoadUserInfo                func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	SaveUserInfo                func(err error, packet nex.PacketInterface, callID uint32, userInfo *service_item_wii_sports_club_types.ServiceItemUserInfo) (*nex.RMCMessage, *nex.Error)
	StartChallenge              func(err error, packet nex.PacketInterface, callID uint32, startChallengeParam *service_item_wii_sports_club_types.ServiceItemStartChallengeParam) (*nex.RMCMessage, *nex.Error)
	EndChallenge                func(err error, packet nex.PacketInterface, callID uint32, endChallengeParam *service_item_wii_sports_club_types.ServiceItemEndChallengeParam) (*nex.RMCMessage, *nex.Error)
	RequestTicketRestoration    func(err error, packet nex.PacketInterface, callID uint32, requestTicketRestorationParam *service_item_wii_sports_club_types.ServiceItemRequestTicketRestorationParam) (*nex.RMCMessage, *nex.Error)
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
	case MethodHello:
		protocol.handleHello(packet)
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
	case MethodGetNotice:
		protocol.handleGetNotice(packet)
	case MethodUpdateAndGetTicketInfo:
		protocol.handleUpdateAndGetTicketInfo(packet)
	case MethodLoadUserInfo:
		protocol.handleLoadUserInfo(packet)
	case MethodSaveUserInfo:
		protocol.handleSaveUserInfo(packet)
	case MethodStartChallenge:
		protocol.handleStartChallenge(packet)
	case MethodEndChallenge:
		protocol.handleEndChallenge(packet)
	case MethodRequestTicketRestoration:
		protocol.handleRequestTicketRestoration(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported Service Item (Wii Sports Club) method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new ServiceItemWiiSportsClub protocol
func NewProtocol(endpoint nex.EndpointInterface) *Protocol {
	protocol := &Protocol{endpoint: endpoint}
	protocol.serviceItemProtocol.SetEndpoint(endpoint)

	return protocol
}
