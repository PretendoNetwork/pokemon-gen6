// Package protocol implements the Account Management protocol
package protocol

import (
	"fmt"
	"slices"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

const (
	// ProtocolID is the protocol ID for the Account Management protocol
	ProtocolID = 0x19

	// MethodCreateAccount is the method ID for method CreateAccount
	MethodCreateAccount = 0x1

	// MethodDeleteAccount is the method ID for method DeleteAccount
	MethodDeleteAccount = 0x2

	// MethodDisableAccount is the method ID for method DisableAccount
	MethodDisableAccount = 0x3

	// MethodChangePassword is the method ID for method ChangePassword
	MethodChangePassword = 0x4

	// MethodTestCapability is the method ID for method TestCapability
	MethodTestCapability = 0x5

	// MethodGetName is the method ID for method GetName
	MethodGetName = 0x6

	// MethodGetAccountData is the method ID for method GetAccountData
	MethodGetAccountData = 0x7

	// MethodGetPrivateData is the method ID for method GetPrivateData
	MethodGetPrivateData = 0x8

	// MethodGetPublicData is the method ID for method GetPublicData
	MethodGetPublicData = 0x9

	// MethodGetMultiplePublicData is the method ID for method GetMultiplePublicData
	MethodGetMultiplePublicData = 0xA

	// MethodUpdateAccountName is the method ID for method UpdateAccountName
	MethodUpdateAccountName = 0xB

	// MethodUpdateAccountEmail is the method ID for method UpdateAccountEmail
	MethodUpdateAccountEmail = 0xC

	// MethodUpdateCustomData is the method ID for method UpdateCustomData
	MethodUpdateCustomData = 0xD

	// MethodFindByNameRegex is the method ID for method FindByNameRegex
	MethodFindByNameRegex = 0xE

	// MethodUpdateAccountExpiryDate is the method ID for method UpdateAccountExpiryDate
	MethodUpdateAccountExpiryDate = 0xF

	// MethodUpdateAccountEffectiveDate is the method ID for method UpdateAccountEffectiveDate
	MethodUpdateAccountEffectiveDate = 0x10

	// MethodUpdateStatus is the method ID for method UpdateStatus
	MethodUpdateStatus = 0x11

	// MethodGetStatus is the method ID for method GetStatus
	MethodGetStatus = 0x12

	// MethodGetLastConnectionStats is the method ID for method GetLastConnectionStats
	MethodGetLastConnectionStats = 0x13

	// MethodResetPassword is the method ID for method ResetPassword
	MethodResetPassword = 0x14

	// MethodCreateAccountWithCustomData is the method ID for method CreateAccountWithCustomData
	MethodCreateAccountWithCustomData = 0x15

	// MethodRetrieveAccount is the method ID for method RetrieveAccount
	MethodRetrieveAccount = 0x16

	// MethodUpdateAccount is the method ID for method UpdateAccount
	MethodUpdateAccount = 0x17

	// MethodChangePasswordByGuest is the method ID for method ChangePasswordByGuest
	MethodChangePasswordByGuest = 0x18

	// MethodFindByNameLike is the method ID for method FindByNameLike
	MethodFindByNameLike = 0x19

	// MethodCustomCreateAccount is the method ID for method CustomCreateAccount
	MethodCustomCreateAccount = 0x1A

	// MethodNintendoCreateAccount is the method ID for method NintendoCreateAccount
	MethodNintendoCreateAccount = 0x1B

	// MethodLookupOrCreateAccount is the method ID for method LookupOrCreateAccount
	MethodLookupOrCreateAccount = 0x1C

	// MethodDisconnectPrincipal is the method ID for method DisconnectPrincipal
	MethodDisconnectPrincipal = 0x1D

	// MethodDisconnectAllPrincipals is the method ID for method DisconnectAllPrincipals
	MethodDisconnectAllPrincipals = 0x1E
)

// Protocol stores all the RMC method handlers for the Account Management protocol and listens for requests
type Protocol struct {
	endpoint                    nex.EndpointInterface
	CreateAccount               func(err error, packet nex.PacketInterface, callID uint32, strPrincipalName *types.String, strKey *types.String, uiGroups *types.PrimitiveU32, strEmail *types.String) (*nex.RMCMessage, *nex.Error)
	DeleteAccount               func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *types.PID) (*nex.RMCMessage, *nex.Error)
	DisableAccount              func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *types.PID, dtUntil *types.DateTime, strMessage *types.String) (*nex.RMCMessage, *nex.Error)
	ChangePassword              func(err error, packet nex.PacketInterface, callID uint32, strNewKey *types.String) (*nex.RMCMessage, *nex.Error)
	TestCapability              func(err error, packet nex.PacketInterface, callID uint32, uiCapability *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	GetName                     func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *types.PID) (*nex.RMCMessage, *nex.Error)
	GetAccountData              func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	GetPrivateData              func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	GetPublicData               func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *types.PID) (*nex.RMCMessage, *nex.Error)
	GetMultiplePublicData       func(err error, packet nex.PacketInterface, callID uint32, lstPrincipals *types.List[*types.PID]) (*nex.RMCMessage, *nex.Error)
	UpdateAccountName           func(err error, packet nex.PacketInterface, callID uint32, strName *types.String) (*nex.RMCMessage, *nex.Error)
	UpdateAccountEmail          func(err error, packet nex.PacketInterface, callID uint32, strName *types.String) (*nex.RMCMessage, *nex.Error)
	UpdateCustomData            func(err error, packet nex.PacketInterface, callID uint32, oPublicData *types.AnyDataHolder, oPrivateData *types.AnyDataHolder) (*nex.RMCMessage, *nex.Error)
	FindByNameRegex             func(err error, packet nex.PacketInterface, callID uint32, uiGroups *types.PrimitiveU32, strRegex *types.String, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error)
	UpdateAccountExpiryDate     func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *types.PID, dtExpiry *types.DateTime, strExpiredMessage *types.String) (*nex.RMCMessage, *nex.Error)
	UpdateAccountEffectiveDate  func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *types.PID, dtEffectiveFrom *types.DateTime, strNotEffectiveMessage *types.String) (*nex.RMCMessage, *nex.Error)
	UpdateStatus                func(err error, packet nex.PacketInterface, callID uint32, strStatus *types.String) (*nex.RMCMessage, *nex.Error)
	GetStatus                   func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *types.PID) (*nex.RMCMessage, *nex.Error)
	GetLastConnectionStats      func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *types.PID) (*nex.RMCMessage, *nex.Error)
	ResetPassword               func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	CreateAccountWithCustomData func(err error, packet nex.PacketInterface, callID uint32, strPrincipalName *types.String, strKey *types.String, uiGroups *types.PrimitiveU32, strEmail *types.String, oPublicData *types.AnyDataHolder, oPrivateData *types.AnyDataHolder) (*nex.RMCMessage, *nex.Error)
	RetrieveAccount             func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	UpdateAccount               func(err error, packet nex.PacketInterface, callID uint32, strKey *types.String, strEmail *types.String, oPublicData *types.AnyDataHolder, oPrivateData *types.AnyDataHolder) (*nex.RMCMessage, *nex.Error)
	ChangePasswordByGuest       func(err error, packet nex.PacketInterface, callID uint32, strPrincipalName *types.String, strKey *types.String, strEmail *types.String) (*nex.RMCMessage, *nex.Error)
	FindByNameLike              func(err error, packet nex.PacketInterface, callID uint32, uiGroups *types.PrimitiveU32, strLike *types.String, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error)
	CustomCreateAccount         func(err error, packet nex.PacketInterface, callID uint32, strPrincipalName *types.String, strKey *types.String, uiGroups *types.PrimitiveU32, strEmail *types.String, oAuthData *types.AnyDataHolder) (*nex.RMCMessage, *nex.Error)
	NintendoCreateAccount       func(err error, packet nex.PacketInterface, callID uint32, strPrincipalName *types.String, strKey *types.String, uiGroups *types.PrimitiveU32, strEmail *types.String, oAuthData *types.AnyDataHolder) (*nex.RMCMessage, *nex.Error)
	LookupOrCreateAccount       func(err error, packet nex.PacketInterface, callID uint32, strPrincipalName *types.String, strKey *types.String, uiGroups *types.PrimitiveU32, strEmail *types.String, oAuthData *types.AnyDataHolder) (*nex.RMCMessage, *nex.Error)
	DisconnectPrincipal         func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *types.PID) (*nex.RMCMessage, *nex.Error)
	DisconnectAllPrincipals     func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	Patches                     nex.ServiceProtocol
	PatchedMethods              []uint32
}

// Interface implements the methods present on the Account Management Protocol struct
type Interface interface {
	Endpoint() nex.EndpointInterface
	SetEndpoint(endpoint nex.EndpointInterface)
	SetHandlerCreateAccount(handler func(err error, packet nex.PacketInterface, callID uint32, strPrincipalName *types.String, strKey *types.String, uiGroups *types.PrimitiveU32, strEmail *types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerDeleteAccount(handler func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *types.PID) (*nex.RMCMessage, *nex.Error))
	SetHandlerDisableAccount(handler func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *types.PID, dtUntil *types.DateTime, strMessage *types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerChangePassword(handler func(err error, packet nex.PacketInterface, callID uint32, strNewKey *types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerTestCapability(handler func(err error, packet nex.PacketInterface, callID uint32, uiCapability *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetName(handler func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *types.PID) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetAccountData(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetPrivateData(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetPublicData(handler func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *types.PID) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetMultiplePublicData(handler func(err error, packet nex.PacketInterface, callID uint32, lstPrincipals *types.List[*types.PID]) (*nex.RMCMessage, *nex.Error))
	SetHandlerUpdateAccountName(handler func(err error, packet nex.PacketInterface, callID uint32, strName *types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerUpdateAccountEmail(handler func(err error, packet nex.PacketInterface, callID uint32, strName *types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerUpdateCustomData(handler func(err error, packet nex.PacketInterface, callID uint32, oPublicData *types.AnyDataHolder, oPrivateData *types.AnyDataHolder) (*nex.RMCMessage, *nex.Error))
	SetHandlerFindByNameRegex(handler func(err error, packet nex.PacketInterface, callID uint32, uiGroups *types.PrimitiveU32, strRegex *types.String, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error))
	SetHandlerUpdateAccountExpiryDate(handler func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *types.PID, dtExpiry *types.DateTime, strExpiredMessage *types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerUpdateAccountEffectiveDate(handler func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *types.PID, dtEffectiveFrom *types.DateTime, strNotEffectiveMessage *types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerUpdateStatus(handler func(err error, packet nex.PacketInterface, callID uint32, strStatus *types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetStatus(handler func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *types.PID) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetLastConnectionStats(handler func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *types.PID) (*nex.RMCMessage, *nex.Error))
	SetHandlerResetPassword(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error))
	SetHandlerCreateAccountWithCustomData(handler func(err error, packet nex.PacketInterface, callID uint32, strPrincipalName *types.String, strKey *types.String, uiGroups *types.PrimitiveU32, strEmail *types.String, oPublicData *types.AnyDataHolder, oPrivateData *types.AnyDataHolder) (*nex.RMCMessage, *nex.Error))
	SetHandlerRetrieveAccount(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error))
	SetHandlerUpdateAccount(handler func(err error, packet nex.PacketInterface, callID uint32, strKey *types.String, strEmail *types.String, oPublicData *types.AnyDataHolder, oPrivateData *types.AnyDataHolder) (*nex.RMCMessage, *nex.Error))
	SetHandlerChangePasswordByGuest(handler func(err error, packet nex.PacketInterface, callID uint32, strPrincipalName *types.String, strKey *types.String, strEmail *types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerFindByNameLike(handler func(err error, packet nex.PacketInterface, callID uint32, uiGroups *types.PrimitiveU32, strLike *types.String, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error))
	SetHandlerCustomCreateAccount(handler func(err error, packet nex.PacketInterface, callID uint32, strPrincipalName *types.String, strKey *types.String, uiGroups *types.PrimitiveU32, strEmail *types.String, oAuthData *types.AnyDataHolder) (*nex.RMCMessage, *nex.Error))
	SetHandlerNintendoCreateAccount(handler func(err error, packet nex.PacketInterface, callID uint32, strPrincipalName *types.String, strKey *types.String, uiGroups *types.PrimitiveU32, strEmail *types.String, oAuthData *types.AnyDataHolder) (*nex.RMCMessage, *nex.Error))
	SetHandlerLookupOrCreateAccount(handler func(err error, packet nex.PacketInterface, callID uint32, strPrincipalName *types.String, strKey *types.String, uiGroups *types.PrimitiveU32, strEmail *types.String, oAuthData *types.AnyDataHolder) (*nex.RMCMessage, *nex.Error))
	SetHandlerDisconnectPrincipal(handler func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *types.PID) (*nex.RMCMessage, *nex.Error))
	SetHandlerDisconnectAllPrincipals(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error))
}

// Endpoint returns the endpoint implementing the protocol
func (protocol *Protocol) Endpoint() nex.EndpointInterface {
	return protocol.endpoint
}

// SetEndpoint sets the endpoint implementing the protocol
func (protocol *Protocol) SetEndpoint(endpoint nex.EndpointInterface) {
	protocol.endpoint = endpoint
}

// SetHandlerCreateAccount sets the handler for the CreateAccount method
func (protocol *Protocol) SetHandlerCreateAccount(handler func(err error, packet nex.PacketInterface, callID uint32, strPrincipalName *types.String, strKey *types.String, uiGroups *types.PrimitiveU32, strEmail *types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.CreateAccount = handler
}

// SetHandlerDeleteAccount sets the handler for the DeleteAccount method
func (protocol *Protocol) SetHandlerDeleteAccount(handler func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *types.PID) (*nex.RMCMessage, *nex.Error)) {
	protocol.DeleteAccount = handler
}

// SetHandlerDisableAccount sets the handler for the DisableAccount method
func (protocol *Protocol) SetHandlerDisableAccount(handler func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *types.PID, dtUntil *types.DateTime, strMessage *types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.DisableAccount = handler
}

// SetHandlerChangePassword sets the handler for the ChangePassword method
func (protocol *Protocol) SetHandlerChangePassword(handler func(err error, packet nex.PacketInterface, callID uint32, strNewKey *types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.ChangePassword = handler
}

// SetHandlerTestCapability sets the handler for the TestCapability method
func (protocol *Protocol) SetHandlerTestCapability(handler func(err error, packet nex.PacketInterface, callID uint32, uiCapability *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)) {
	protocol.TestCapability = handler
}

// SetHandlerGetName sets the handler for the GetName method
func (protocol *Protocol) SetHandlerGetName(handler func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *types.PID) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetName = handler
}

// SetHandlerGetAccountData sets the handler for the GetAccountData method
func (protocol *Protocol) SetHandlerGetAccountData(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetAccountData = handler
}

// SetHandlerGetPrivateData sets the handler for the GetPrivateData method
func (protocol *Protocol) SetHandlerGetPrivateData(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetPrivateData = handler
}

// SetHandlerGetPublicData sets the handler for the GetPublicData method
func (protocol *Protocol) SetHandlerGetPublicData(handler func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *types.PID) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetPublicData = handler
}

// SetHandlerGetMultiplePublicData sets the handler for the GetMultiplePublicData method
func (protocol *Protocol) SetHandlerGetMultiplePublicData(handler func(err error, packet nex.PacketInterface, callID uint32, lstPrincipals *types.List[*types.PID]) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetMultiplePublicData = handler
}

// SetHandlerUpdateAccountName sets the handler for the UpdateAccountName method
func (protocol *Protocol) SetHandlerUpdateAccountName(handler func(err error, packet nex.PacketInterface, callID uint32, strName *types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.UpdateAccountName = handler
}

// SetHandlerUpdateAccountEmail sets the handler for the UpdateAccountEmail method
func (protocol *Protocol) SetHandlerUpdateAccountEmail(handler func(err error, packet nex.PacketInterface, callID uint32, strName *types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.UpdateAccountEmail = handler
}

// SetHandlerUpdateCustomData sets the handler for the UpdateCustomData method
func (protocol *Protocol) SetHandlerUpdateCustomData(handler func(err error, packet nex.PacketInterface, callID uint32, oPublicData *types.AnyDataHolder, oPrivateData *types.AnyDataHolder) (*nex.RMCMessage, *nex.Error)) {
	protocol.UpdateCustomData = handler
}

// SetHandlerFindByNameRegex sets the handler for the FindByNameRegex method
func (protocol *Protocol) SetHandlerFindByNameRegex(handler func(err error, packet nex.PacketInterface, callID uint32, uiGroups *types.PrimitiveU32, strRegex *types.String, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error)) {
	protocol.FindByNameRegex = handler
}

// SetHandlerUpdateAccountExpiryDate sets the handler for the UpdateAccountExpiryDate method
func (protocol *Protocol) SetHandlerUpdateAccountExpiryDate(handler func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *types.PID, dtExpiry *types.DateTime, strExpiredMessage *types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.UpdateAccountExpiryDate = handler
}

// SetHandlerUpdateAccountEffectiveDate sets the handler for the UpdateAccountEffectiveDate method
func (protocol *Protocol) SetHandlerUpdateAccountEffectiveDate(handler func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *types.PID, dtEffectiveFrom *types.DateTime, strNotEffectiveMessage *types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.UpdateAccountEffectiveDate = handler
}

// SetHandlerUpdateStatus sets the handler for the UpdateStatus method
func (protocol *Protocol) SetHandlerUpdateStatus(handler func(err error, packet nex.PacketInterface, callID uint32, strStatus *types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.UpdateStatus = handler
}

// SetHandlerGetStatus sets the handler for the GetStatus method
func (protocol *Protocol) SetHandlerGetStatus(handler func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *types.PID) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetStatus = handler
}

// SetHandlerGetLastConnectionStats sets the handler for the GetLastConnectionStats method
func (protocol *Protocol) SetHandlerGetLastConnectionStats(handler func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *types.PID) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetLastConnectionStats = handler
}

// SetHandlerResetPassword sets the handler for the ResetPassword method
func (protocol *Protocol) SetHandlerResetPassword(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.ResetPassword = handler
}

// SetHandlerCreateAccountWithCustomData sets the handler for the CreateAccountWithCustomData method
func (protocol *Protocol) SetHandlerCreateAccountWithCustomData(handler func(err error, packet nex.PacketInterface, callID uint32, strPrincipalName *types.String, strKey *types.String, uiGroups *types.PrimitiveU32, strEmail *types.String, oPublicData *types.AnyDataHolder, oPrivateData *types.AnyDataHolder) (*nex.RMCMessage, *nex.Error)) {
	protocol.CreateAccountWithCustomData = handler
}

// SetHandlerRetrieveAccount sets the handler for the RetrieveAccount method
func (protocol *Protocol) SetHandlerRetrieveAccount(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.RetrieveAccount = handler
}

// SetHandlerUpdateAccount sets the handler for the UpdateAccount method
func (protocol *Protocol) SetHandlerUpdateAccount(handler func(err error, packet nex.PacketInterface, callID uint32, strKey *types.String, strEmail *types.String, oPublicData *types.AnyDataHolder, oPrivateData *types.AnyDataHolder) (*nex.RMCMessage, *nex.Error)) {
	protocol.UpdateAccount = handler
}

// SetHandlerChangePasswordByGuest sets the handler for the ChangePasswordByGuest method
func (protocol *Protocol) SetHandlerChangePasswordByGuest(handler func(err error, packet nex.PacketInterface, callID uint32, strPrincipalName *types.String, strKey *types.String, strEmail *types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.ChangePasswordByGuest = handler
}

// SetHandlerFindByNameLike sets the handler for the FindByNameLike method
func (protocol *Protocol) SetHandlerFindByNameLike(handler func(err error, packet nex.PacketInterface, callID uint32, uiGroups *types.PrimitiveU32, strLike *types.String, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error)) {
	protocol.FindByNameLike = handler
}

// SetHandlerCustomCreateAccount sets the handler for the CustomCreateAccount method
func (protocol *Protocol) SetHandlerCustomCreateAccount(handler func(err error, packet nex.PacketInterface, callID uint32, strPrincipalName *types.String, strKey *types.String, uiGroups *types.PrimitiveU32, strEmail *types.String, oAuthData *types.AnyDataHolder) (*nex.RMCMessage, *nex.Error)) {
	protocol.CustomCreateAccount = handler
}

// SetHandlerNintendoCreateAccount sets the handler for the NintendoCreateAccount method
func (protocol *Protocol) SetHandlerNintendoCreateAccount(handler func(err error, packet nex.PacketInterface, callID uint32, strPrincipalName *types.String, strKey *types.String, uiGroups *types.PrimitiveU32, strEmail *types.String, oAuthData *types.AnyDataHolder) (*nex.RMCMessage, *nex.Error)) {
	protocol.NintendoCreateAccount = handler
}

// SetHandlerLookupOrCreateAccount sets the handler for the LookupOrCreateAccount method
func (protocol *Protocol) SetHandlerLookupOrCreateAccount(handler func(err error, packet nex.PacketInterface, callID uint32, strPrincipalName *types.String, strKey *types.String, uiGroups *types.PrimitiveU32, strEmail *types.String, oAuthData *types.AnyDataHolder) (*nex.RMCMessage, *nex.Error)) {
	protocol.LookupOrCreateAccount = handler
}

// SetHandlerDisconnectPrincipal sets the handler for the DisconnectPrincipal method
func (protocol *Protocol) SetHandlerDisconnectPrincipal(handler func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *types.PID) (*nex.RMCMessage, *nex.Error)) {
	protocol.DisconnectPrincipal = handler
}

// SetHandlerDisconnectAllPrincipals sets the handler for the DisconnectAllPrincipals method
func (protocol *Protocol) SetHandlerDisconnectAllPrincipals(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.DisconnectAllPrincipals = handler
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	message := packet.RMCMessage()

	if !message.IsRequest || message.ProtocolID != ProtocolID {
		return
	}

	if protocol.Patches != nil && slices.Contains(protocol.PatchedMethods, message.MethodID) {
		protocol.Patches.HandlePacket(packet)
		return
	}

	switch message.MethodID {
	case MethodCreateAccount:
		protocol.handleCreateAccount(packet)
	case MethodDeleteAccount:
		protocol.handleDeleteAccount(packet)
	case MethodDisableAccount:
		protocol.handleDisableAccount(packet)
	case MethodChangePassword:
		protocol.handleChangePassword(packet)
	case MethodTestCapability:
		protocol.handleTestCapability(packet)
	case MethodGetName:
		protocol.handleGetName(packet)
	case MethodGetAccountData:
		protocol.handleGetAccountData(packet)
	case MethodGetPrivateData:
		protocol.handleGetPrivateData(packet)
	case MethodGetPublicData:
		protocol.handleGetPublicData(packet)
	case MethodGetMultiplePublicData:
		protocol.handleGetMultiplePublicData(packet)
	case MethodUpdateAccountName:
		protocol.handleUpdateAccountName(packet)
	case MethodUpdateAccountEmail:
		protocol.handleUpdateAccountEmail(packet)
	case MethodUpdateCustomData:
		protocol.handleUpdateCustomData(packet)
	case MethodFindByNameRegex:
		protocol.handleFindByNameRegex(packet)
	case MethodUpdateAccountExpiryDate:
		protocol.handleUpdateAccountExpiryDate(packet)
	case MethodUpdateAccountEffectiveDate:
		protocol.handleUpdateAccountEffectiveDate(packet)
	case MethodUpdateStatus:
		protocol.handleUpdateStatus(packet)
	case MethodGetStatus:
		protocol.handleGetStatus(packet)
	case MethodGetLastConnectionStats:
		protocol.handleGetLastConnectionStats(packet)
	case MethodResetPassword:
		protocol.handleResetPassword(packet)
	case MethodCreateAccountWithCustomData:
		protocol.handleCreateAccountWithCustomData(packet)
	case MethodRetrieveAccount:
		protocol.handleRetrieveAccount(packet)
	case MethodUpdateAccount:
		protocol.handleUpdateAccount(packet)
	case MethodChangePasswordByGuest:
		protocol.handleChangePasswordByGuest(packet)
	case MethodFindByNameLike:
		protocol.handleFindByNameLike(packet)
	case MethodCustomCreateAccount:
		protocol.handleCustomCreateAccount(packet)
	case MethodNintendoCreateAccount:
		protocol.handleNintendoCreateAccount(packet)
	case MethodLookupOrCreateAccount:
		protocol.handleLookupOrCreateAccount(packet)
	case MethodDisconnectPrincipal:
		protocol.handleDisconnectPrincipal(packet)
	case MethodDisconnectAllPrincipals:
		protocol.handleDisconnectAllPrincipals(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported AccountManagement method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new Account Management protocol
func NewProtocol() *Protocol {
	return &Protocol{}
}
