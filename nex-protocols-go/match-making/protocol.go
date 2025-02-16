// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"
	"slices"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/v2/match-making/types"
)

const (
	// ProtocolID is the protocol ID for the Match Making protocol
	ProtocolID = 0x15

	// MethodRegisterGathering is the method ID for method RegisterGathering
	MethodRegisterGathering = 0x1

	// MethodUnregisterGathering is the method ID for method UnregisterGathering
	MethodUnregisterGathering = 0x2

	// MethodUnregisterGatherings is the method ID for method UnregisterGatherings
	MethodUnregisterGatherings = 0x3

	// MethodUpdateGathering is the method ID for method UpdateGathering
	MethodUpdateGathering = 0x4

	// MethodInvite is the method ID for method Invite
	MethodInvite = 0x5

	// MethodAcceptInvitation is the method ID for method AcceptInvitation
	MethodAcceptInvitation = 0x6

	// MethodDeclineInvitation is the method ID for method DeclineInvitation
	MethodDeclineInvitation = 0x7

	// MethodCancelInvitation is the method ID for method CancelInvitation
	MethodCancelInvitation = 0x8

	// MethodGetInvitationsSent is the method ID for method GetInvitationsSent
	MethodGetInvitationsSent = 0x9

	// MethodGetInvitationsReceived is the method ID for method GetInvitationsReceived
	MethodGetInvitationsReceived = 0xA

	// MethodParticipate is the method ID for method Participate
	MethodParticipate = 0xB

	// MethodCancelParticipation is the method ID for method CancelParticipation
	MethodCancelParticipation = 0xC

	// MethodGetParticipants is the method ID for method GetParticipants
	MethodGetParticipants = 0xD

	// MethodAddParticipants is the method ID for method AddParticipants
	MethodAddParticipants = 0xE

	// MethodGetDetailedParticipants is the method ID for method GetDetailedParticipants
	MethodGetDetailedParticipants = 0xF

	// MethodGetParticipantsURLs is the method ID for method GetParticipantsURLs
	MethodGetParticipantsURLs = 0x10

	// MethodFindByType is the method ID for method FindByType
	MethodFindByType = 0x11

	// MethodFindByDescription is the method ID for method FindByDescription
	MethodFindByDescription = 0x12

	// MethodFindByDescriptionRegex is the method ID for method FindByDescriptionRegex
	MethodFindByDescriptionRegex = 0x13

	// MethodFindByID is the method ID for method FindByID
	MethodFindByID = 0x14

	// MethodFindBySingleID is the method ID for method FindBySingleID
	MethodFindBySingleID = 0x15

	// MethodFindByOwner is the method ID for method FindByOwner
	MethodFindByOwner = 0x16

	// MethodFindByParticipants is the method ID for method FindByParticipants
	MethodFindByParticipants = 0x17

	// MethodFindInvitations is the method ID for method FindInvitations
	MethodFindInvitations = 0x18

	// MethodFindBySQLQuery is the method ID for method FindBySQLQuery
	MethodFindBySQLQuery = 0x19

	// MethodLaunchSession is the method ID for method LaunchSession
	MethodLaunchSession = 0x1A

	// MethodUpdateSessionURL is the method ID for method UpdateSessionURL
	MethodUpdateSessionURL = 0x1B

	// MethodGetSessionURL is the method ID for method GetSessionURL
	MethodGetSessionURL = 0x1C

	// MethodGetState is the method ID for method GetState
	MethodGetState = 0x1D

	// MethodSetState is the method ID for method SetState
	MethodSetState = 0x1E

	// MethodReportStats is the method ID for method ReportStats
	MethodReportStats = 0x1F

	// MethodGetStats is the method ID for method GetStats
	MethodGetStats = 0x20

	// MethodDeleteGathering is the method ID for method DeleteGathering
	MethodDeleteGathering = 0x21

	// MethodGetPendingDeletions is the method ID for method GetPendingDeletions
	MethodGetPendingDeletions = 0x22

	// MethodDeleteFromDeletions is the method ID for method DeleteFromDeletions
	MethodDeleteFromDeletions = 0x23

	// MethodMigrateGatheringOwnershipV1 is the method ID for method MigrateGatheringOwnershipV1
	MethodMigrateGatheringOwnershipV1 = 0x24

	// MethodFindByDescriptionLike is the method ID for method FindByDescriptionLike
	MethodFindByDescriptionLike = 0x25

	// MethodRegisterLocalURL is the method ID for method RegisterLocalURL
	MethodRegisterLocalURL = 0x26

	// MethodRegisterLocalURLs is the method ID for method RegisterLocalURLs
	MethodRegisterLocalURLs = 0x27

	// MethodUpdateSessionHostV1 is the method ID for method UpdateSessionHostV1
	MethodUpdateSessionHostV1 = 0x28

	// MethodGetSessionURLs is the method ID for method GetSessionURLs
	MethodGetSessionURLs = 0x29

	// MethodUpdateSessionHost is the method ID for method UpdateSessionHost
	MethodUpdateSessionHost = 0x2A

	// MethodUpdateGatheringOwnership is the method ID for method UpdateGatheringOwnership
	MethodUpdateGatheringOwnership = 0x2B

	// MethodMigrateGatheringOwnership is the method ID for method MigrateGatheringOwnership
	MethodMigrateGatheringOwnership = 0x2C
)

// Protocol stores all the RMC method handlers for the MatchMaking protocol and listens for requests
type Protocol struct {
	endpoint                    nex.EndpointInterface
	RegisterGathering           func(err error, packet nex.PacketInterface, callID uint32, anyGathering *types.AnyDataHolder) (*nex.RMCMessage, *nex.Error)
	UnregisterGathering         func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	UnregisterGatherings        func(err error, packet nex.PacketInterface, callID uint32, lstGatherings *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, *nex.Error)
	UpdateGathering             func(err error, packet nex.PacketInterface, callID uint32, anyGathering *types.AnyDataHolder) (*nex.RMCMessage, *nex.Error)
	Invite                      func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32, lstPrincipals *types.List[*types.PID], strMessage *types.String) (*nex.RMCMessage, *nex.Error)
	AcceptInvitation            func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32, strMessage *types.String) (*nex.RMCMessage, *nex.Error)
	DeclineInvitation           func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32, strMessage *types.String) (*nex.RMCMessage, *nex.Error)
	CancelInvitation            func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32, lstPrincipals *types.List[*types.PID], strMessage *types.String) (*nex.RMCMessage, *nex.Error)
	GetInvitationsSent          func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	GetInvitationsReceived      func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	Participate                 func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32, strMessage *types.String) (*nex.RMCMessage, *nex.Error)
	CancelParticipation         func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32, strMessage *types.String) (*nex.RMCMessage, *nex.Error)
	GetParticipants             func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	AddParticipants             func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32, lstPrincipals *types.List[*types.PID], strMessage *types.String) (*nex.RMCMessage, *nex.Error)
	GetDetailedParticipants     func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	GetParticipantsURLs         func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	FindByType                  func(err error, packet nex.PacketInterface, callID uint32, strType *types.String, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error)
	FindByDescription           func(err error, packet nex.PacketInterface, callID uint32, strDescription *types.String, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error)
	FindByDescriptionRegex      func(err error, packet nex.PacketInterface, callID uint32, strDescriptionRegex *types.String, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error)
	FindByID                    func(err error, packet nex.PacketInterface, callID uint32, lstID *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, *nex.Error)
	FindBySingleID              func(err error, packet nex.PacketInterface, callID uint32, id *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	FindByOwner                 func(err error, packet nex.PacketInterface, callID uint32, id *types.PID, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error)
	FindByParticipants          func(err error, packet nex.PacketInterface, callID uint32, pid *types.List[*types.PID]) (*nex.RMCMessage, *nex.Error)
	FindInvitations             func(err error, packet nex.PacketInterface, callID uint32, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error)
	FindBySQLQuery              func(err error, packet nex.PacketInterface, callID uint32, strQuery *types.String, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error)
	LaunchSession               func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32, strURL *types.String) (*nex.RMCMessage, *nex.Error)
	UpdateSessionURL            func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32, strURL *types.String) (*nex.RMCMessage, *nex.Error)
	GetSessionURL               func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	GetState                    func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	SetState                    func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32, uiNewState *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	ReportStats                 func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32, lstStats *types.List[*match_making_types.GatheringStats]) (*nex.RMCMessage, *nex.Error)
	GetStats                    func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32, lstParticipants *types.List[*types.PID], lstColumns *types.Buffer) (*nex.RMCMessage, *nex.Error)
	DeleteGathering             func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	GetPendingDeletions         func(err error, packet nex.PacketInterface, callID uint32, uiReason *types.PrimitiveU32, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error)
	DeleteFromDeletions         func(err error, packet nex.PacketInterface, callID uint32, lstDeletions *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, *nex.Error)
	MigrateGatheringOwnershipV1 func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, lstPotentialNewOwnersID *types.List[*types.PID]) (*nex.RMCMessage, *nex.Error)
	FindByDescriptionLike       func(err error, packet nex.PacketInterface, callID uint32, strDescriptionLike *types.String, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error)
	RegisterLocalURL            func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, url *types.StationURL) (*nex.RMCMessage, *nex.Error)
	RegisterLocalURLs           func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, lstURLs *types.List[*types.StationURL]) (*nex.RMCMessage, *nex.Error)
	UpdateSessionHostV1         func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	GetSessionURLs              func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	UpdateSessionHost           func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, isMigrateOwner *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error)
	UpdateGatheringOwnership    func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, participantsOnly *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error)
	MigrateGatheringOwnership   func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, lstPotentialNewOwnersID *types.List[*types.PID], participantsOnly *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error)
	Patches                     nex.ServiceProtocol
	PatchedMethods              []uint32
}

// Interface implements the methods present on the Match Making protocol struct
type Interface interface {
	Endpoint() nex.EndpointInterface
	SetEndpoint(endpoint nex.EndpointInterface)
	SetHandlerRegisterGathering(handler func(err error, packet nex.PacketInterface, callID uint32, anyGathering *types.AnyDataHolder) (*nex.RMCMessage, *nex.Error))
	SetHandlerUnregisterGathering(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error))
	SetHandlerUnregisterGatherings(handler func(err error, packet nex.PacketInterface, callID uint32, lstGatherings *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, *nex.Error))
	SetHandlerUpdateGathering(handler func(err error, packet nex.PacketInterface, callID uint32, anyGathering *types.AnyDataHolder) (*nex.RMCMessage, *nex.Error))
	SetHandlerInvite(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32, lstPrincipals *types.List[*types.PID], strMessage *types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerAcceptInvitation(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32, strMessage *types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerDeclineInvitation(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32, strMessage *types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerCancelInvitation(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32, lstPrincipals *types.List[*types.PID], strMessage *types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetInvitationsSent(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetInvitationsReceived(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error))
	SetHandlerParticipate(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32, strMessage *types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerCancelParticipation(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32, strMessage *types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetParticipants(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error))
	SetHandlerAddParticipants(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32, lstPrincipals *types.List[*types.PID], strMessage *types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetDetailedParticipants(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetParticipantsURLs(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error))
	SetHandlerFindByType(handler func(err error, packet nex.PacketInterface, callID uint32, strType *types.String, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error))
	SetHandlerFindByDescription(handler func(err error, packet nex.PacketInterface, callID uint32, strDescription *types.String, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error))
	SetHandlerFindByDescriptionRegex(handler func(err error, packet nex.PacketInterface, callID uint32, strDescriptionRegex *types.String, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error))
	SetHandlerFindByID(handler func(err error, packet nex.PacketInterface, callID uint32, lstID *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, *nex.Error))
	SetHandlerFindBySingleID(handler func(err error, packet nex.PacketInterface, callID uint32, id *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error))
	SetHandlerFindByOwner(handler func(err error, packet nex.PacketInterface, callID uint32, id *types.PID, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error))
	SetHandlerFindByParticipants(handler func(err error, packet nex.PacketInterface, callID uint32, pid *types.List[*types.PID]) (*nex.RMCMessage, *nex.Error))
	SetHandlerFindInvitations(handler func(err error, packet nex.PacketInterface, callID uint32, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error))
	SetHandlerFindBySQLQuery(handler func(err error, packet nex.PacketInterface, callID uint32, strQuery *types.String, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error))
	SetHandlerLaunchSession(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32, strURL *types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerUpdateSessionURL(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32, strURL *types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetSessionURL(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetState(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error))
	SetHandlerSetState(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32, uiNewState *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error))
	SetHandlerReportStats(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32, lstStats *types.List[*match_making_types.GatheringStats]) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetStats(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32, lstParticipants *types.List[*types.PID], lstColumns *types.Buffer) (*nex.RMCMessage, *nex.Error))
	SetHandlerDeleteGathering(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetPendingDeletions(handler func(err error, packet nex.PacketInterface, callID uint32, uiReason *types.PrimitiveU32, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error))
	SetHandlerDeleteFromDeletions(handler func(err error, packet nex.PacketInterface, callID uint32, lstDeletions *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, *nex.Error))
	SetHandlerMigrateGatheringOwnershipV1(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, lstPotentialNewOwnersID *types.List[*types.PID]) (*nex.RMCMessage, *nex.Error))
	SetHandlerFindByDescriptionLike(handler func(err error, packet nex.PacketInterface, callID uint32, strDescriptionLike *types.String, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error))
	SetHandlerRegisterLocalURL(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, url *types.StationURL) (*nex.RMCMessage, *nex.Error))
	SetHandlerRegisterLocalURLs(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, lstURLs *types.List[*types.StationURL]) (*nex.RMCMessage, *nex.Error))
	SetHandlerUpdateSessionHostV1(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetSessionURLs(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error))
	SetHandlerUpdateSessionHost(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, isMigrateOwner *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error))
	SetHandlerUpdateGatheringOwnership(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, participantsOnly *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error))
	SetHandlerMigrateGatheringOwnership(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, lstPotentialNewOwnersID *types.List[*types.PID], participantsOnly *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error))
}

// Endpoint returns the endpoint implementing the protocol
func (protocol *Protocol) Endpoint() nex.EndpointInterface {
	return protocol.endpoint
}

// SetEndpoint sets the endpoint implementing the protocol
func (protocol *Protocol) SetEndpoint(endpoint nex.EndpointInterface) {
	protocol.endpoint = endpoint
}

// SetHandlerRegisterGathering sets the handler for the RegisterGathering method
func (protocol *Protocol) SetHandlerRegisterGathering(handler func(err error, packet nex.PacketInterface, callID uint32, anyGathering *types.AnyDataHolder) (*nex.RMCMessage, *nex.Error)) {
	protocol.RegisterGathering = handler
}

// SetHandlerUnregisterGathering sets the handler for the UnregisterGathering method
func (protocol *Protocol) SetHandlerUnregisterGathering(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)) {
	protocol.UnregisterGathering = handler
}

// SetHandlerUnregisterGatherings sets the handler for the UnregisterGatherings method
func (protocol *Protocol) SetHandlerUnregisterGatherings(handler func(err error, packet nex.PacketInterface, callID uint32, lstGatherings *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, *nex.Error)) {
	protocol.UnregisterGatherings = handler
}

// SetHandlerUpdateGathering sets the handler for the UpdateGathering method
func (protocol *Protocol) SetHandlerUpdateGathering(handler func(err error, packet nex.PacketInterface, callID uint32, anyGathering *types.AnyDataHolder) (*nex.RMCMessage, *nex.Error)) {
	protocol.UpdateGathering = handler
}

// SetHandlerInvite sets the handler for the Invite method
func (protocol *Protocol) SetHandlerInvite(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32, lstPrincipals *types.List[*types.PID], strMessage *types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.Invite = handler
}

// SetHandlerAcceptInvitation sets the handler for the AcceptInvitation method
func (protocol *Protocol) SetHandlerAcceptInvitation(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32, strMessage *types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.AcceptInvitation = handler
}

// SetHandlerDeclineInvitation sets the handler for the DeclineInvitation method
func (protocol *Protocol) SetHandlerDeclineInvitation(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32, strMessage *types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.DeclineInvitation = handler
}

// SetHandlerCancelInvitation sets the handler for the CancelInvitation method
func (protocol *Protocol) SetHandlerCancelInvitation(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32, lstPrincipals *types.List[*types.PID], strMessage *types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.CancelInvitation = handler
}

// SetHandlerGetInvitationsSent sets the handler for the GetInvitationsSent method
func (protocol *Protocol) SetHandlerGetInvitationsSent(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetInvitationsSent = handler
}

// SetHandlerGetInvitationsReceived sets the handler for the GetInvitationsReceived method
func (protocol *Protocol) SetHandlerGetInvitationsReceived(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetInvitationsReceived = handler
}

// SetHandlerParticipate sets the handler for the Participate method
func (protocol *Protocol) SetHandlerParticipate(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32, strMessage *types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.Participate = handler
}

// SetHandlerCancelParticipation sets the handler for the CancelParticipation method
func (protocol *Protocol) SetHandlerCancelParticipation(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32, strMessage *types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.CancelParticipation = handler
}

// SetHandlerGetParticipants sets the handler for the GetParticipants method
func (protocol *Protocol) SetHandlerGetParticipants(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetParticipants = handler
}

// SetHandlerAddParticipants sets the handler for the AddParticipants method
func (protocol *Protocol) SetHandlerAddParticipants(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32, lstPrincipals *types.List[*types.PID], strMessage *types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.AddParticipants = handler
}

// SetHandlerGetDetailedParticipants sets the handler for the GetDetailedParticipants method
func (protocol *Protocol) SetHandlerGetDetailedParticipants(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetDetailedParticipants = handler
}

// SetHandlerGetParticipantsURLs sets the handler for the GetParticipantsURLs method
func (protocol *Protocol) SetHandlerGetParticipantsURLs(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetParticipantsURLs = handler
}

// SetHandlerFindByType sets the handler for the FindByType method
func (protocol *Protocol) SetHandlerFindByType(handler func(err error, packet nex.PacketInterface, callID uint32, strType *types.String, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error)) {
	protocol.FindByType = handler
}

// SetHandlerFindByDescription sets the handler for the FindByDescription method
func (protocol *Protocol) SetHandlerFindByDescription(handler func(err error, packet nex.PacketInterface, callID uint32, strDescription *types.String, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error)) {
	protocol.FindByDescription = handler
}

// SetHandlerFindByDescriptionRegex sets the handler for the FindByDescriptionRegex method
func (protocol *Protocol) SetHandlerFindByDescriptionRegex(handler func(err error, packet nex.PacketInterface, callID uint32, strDescriptionRegex *types.String, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error)) {
	protocol.FindByDescriptionRegex = handler
}

// SetHandlerFindByID sets the handler for the FindByID method
func (protocol *Protocol) SetHandlerFindByID(handler func(err error, packet nex.PacketInterface, callID uint32, lstID *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, *nex.Error)) {
	protocol.FindByID = handler
}

// SetHandlerFindBySingleID sets the handler for the FindBySingleID method
func (protocol *Protocol) SetHandlerFindBySingleID(handler func(err error, packet nex.PacketInterface, callID uint32, id *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)) {
	protocol.FindBySingleID = handler
}

// SetHandlerFindByOwner sets the handler for the FindByOwner method
func (protocol *Protocol) SetHandlerFindByOwner(handler func(err error, packet nex.PacketInterface, callID uint32, id *types.PID, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error)) {
	protocol.FindByOwner = handler
}

// SetHandlerFindByParticipants sets the handler for the FindByParticipants method
func (protocol *Protocol) SetHandlerFindByParticipants(handler func(err error, packet nex.PacketInterface, callID uint32, pid *types.List[*types.PID]) (*nex.RMCMessage, *nex.Error)) {
	protocol.FindByParticipants = handler
}

// SetHandlerFindInvitations sets the handler for the FindInvitations method
func (protocol *Protocol) SetHandlerFindInvitations(handler func(err error, packet nex.PacketInterface, callID uint32, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error)) {
	protocol.FindInvitations = handler
}

// SetHandlerFindBySQLQuery sets the handler for the FindBySQLQuery method
func (protocol *Protocol) SetHandlerFindBySQLQuery(handler func(err error, packet nex.PacketInterface, callID uint32, strQuery *types.String, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error)) {
	protocol.FindBySQLQuery = handler
}

// SetHandlerLaunchSession sets the handler for the LaunchSession method
func (protocol *Protocol) SetHandlerLaunchSession(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32, strURL *types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.LaunchSession = handler
}

// SetHandlerUpdateSessionURL sets the handler for the UpdateSessionURL method
func (protocol *Protocol) SetHandlerUpdateSessionURL(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32, strURL *types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.UpdateSessionURL = handler
}

// SetHandlerGetSessionURL sets the handler for the GetSessionURL method
func (protocol *Protocol) SetHandlerGetSessionURL(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetSessionURL = handler
}

// SetHandlerGetState sets the handler for the GetState method
func (protocol *Protocol) SetHandlerGetState(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetState = handler
}

// SetHandlerSetState sets the handler for the SetState method
func (protocol *Protocol) SetHandlerSetState(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32, uiNewState *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)) {
	protocol.SetState = handler
}

// SetHandlerReportStats sets the handler for the ReportStats method
func (protocol *Protocol) SetHandlerReportStats(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32, lstStats *types.List[*match_making_types.GatheringStats]) (*nex.RMCMessage, *nex.Error)) {
	protocol.ReportStats = handler
}

// SetHandlerGetStats sets the handler for the GetStats method
func (protocol *Protocol) SetHandlerGetStats(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32, lstParticipants *types.List[*types.PID], lstColumns *types.Buffer) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetStats = handler
}

// SetHandlerDeleteGathering sets the handler for the DeleteGathering method
func (protocol *Protocol) SetHandlerDeleteGathering(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)) {
	protocol.DeleteGathering = handler
}

// SetHandlerGetPendingDeletions sets the handler for the GetPendingDeletions method
func (protocol *Protocol) SetHandlerGetPendingDeletions(handler func(err error, packet nex.PacketInterface, callID uint32, uiReason *types.PrimitiveU32, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetPendingDeletions = handler
}

// SetHandlerDeleteFromDeletions sets the handler for the DeleteFromDeletions method
func (protocol *Protocol) SetHandlerDeleteFromDeletions(handler func(err error, packet nex.PacketInterface, callID uint32, lstDeletions *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, *nex.Error)) {
	protocol.DeleteFromDeletions = handler
}

// SetHandlerMigrateGatheringOwnershipV1 sets the handler for the MigrateGatheringOwnershipV1 method
func (protocol *Protocol) SetHandlerMigrateGatheringOwnershipV1(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, lstPotentialNewOwnersID *types.List[*types.PID]) (*nex.RMCMessage, *nex.Error)) {
	protocol.MigrateGatheringOwnershipV1 = handler
}

// SetHandlerFindByDescriptionLike sets the handler for the FindByDescriptionLike method
func (protocol *Protocol) SetHandlerFindByDescriptionLike(handler func(err error, packet nex.PacketInterface, callID uint32, strDescriptionLike *types.String, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error)) {
	protocol.FindByDescriptionLike = handler
}

// SetHandlerRegisterLocalURL sets the handler for the RegisterLocalURL method
func (protocol *Protocol) SetHandlerRegisterLocalURL(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, url *types.StationURL) (*nex.RMCMessage, *nex.Error)) {
	protocol.RegisterLocalURL = handler
}

// SetHandlerRegisterLocalURLs sets the handler for the RegisterLocalURLs method
func (protocol *Protocol) SetHandlerRegisterLocalURLs(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, lstURLs *types.List[*types.StationURL]) (*nex.RMCMessage, *nex.Error)) {
	protocol.RegisterLocalURLs = handler
}

// SetHandlerUpdateSessionHostV1 sets the handler for the UpdateSessionHostV1 method
func (protocol *Protocol) SetHandlerUpdateSessionHostV1(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)) {
	protocol.UpdateSessionHostV1 = handler
}

// SetHandlerGetSessionURLs sets the handler for the GetSessionURLs method
func (protocol *Protocol) SetHandlerGetSessionURLs(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetSessionURLs = handler
}

// SetHandlerUpdateSessionHost sets the handler for the UpdateSessionHost method
func (protocol *Protocol) SetHandlerUpdateSessionHost(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, isMigrateOwner *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error)) {
	protocol.UpdateSessionHost = handler
}

// SetHandlerUpdateGatheringOwnership sets the handler for the UpdateGatheringOwnership method
func (protocol *Protocol) SetHandlerUpdateGatheringOwnership(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, participantsOnly *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error)) {
	protocol.UpdateGatheringOwnership = handler
}

// SetHandlerMigrateGatheringOwnership sets the handler for the MigrateGatheringOwnership method
func (protocol *Protocol) SetHandlerMigrateGatheringOwnership(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, lstPotentialNewOwnersID *types.List[*types.PID], participantsOnly *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error)) {
	protocol.MigrateGatheringOwnership = handler
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
	case MethodRegisterGathering:
		protocol.handleRegisterGathering(packet)
	case MethodUnregisterGathering:
		protocol.handleUnregisterGathering(packet)
	case MethodUnregisterGatherings:
		protocol.handleUnregisterGatherings(packet)
	case MethodUpdateGathering:
		protocol.handleUpdateGathering(packet)
	case MethodInvite:
		protocol.handleInvite(packet)
	case MethodAcceptInvitation:
		protocol.handleAcceptInvitation(packet)
	case MethodDeclineInvitation:
		protocol.handleDeclineInvitation(packet)
	case MethodCancelInvitation:
		protocol.handleCancelInvitation(packet)
	case MethodGetInvitationsSent:
		protocol.handleGetInvitationsSent(packet)
	case MethodGetInvitationsReceived:
		protocol.handleGetInvitationsReceived(packet)
	case MethodParticipate:
		protocol.handleParticipate(packet)
	case MethodCancelParticipation:
		protocol.handleCancelParticipation(packet)
	case MethodGetParticipants:
		protocol.handleGetParticipants(packet)
	case MethodAddParticipants:
		protocol.handleAddParticipants(packet)
	case MethodGetDetailedParticipants:
		protocol.handleGetDetailedParticipants(packet)
	case MethodGetParticipantsURLs:
		protocol.handleGetParticipantsURLs(packet)
	case MethodFindByType:
		protocol.handleFindByType(packet)
	case MethodFindByDescription:
		protocol.handleFindByDescription(packet)
	case MethodFindByDescriptionRegex:
		protocol.handleFindByDescriptionRegex(packet)
	case MethodFindByID:
		protocol.handleFindByID(packet)
	case MethodFindBySingleID:
		protocol.handleFindBySingleID(packet)
	case MethodFindByOwner:
		protocol.handleFindByOwner(packet)
	case MethodFindByParticipants:
		protocol.handleFindByParticipants(packet)
	case MethodFindInvitations:
		protocol.handleFindInvitations(packet)
	case MethodFindBySQLQuery:
		protocol.handleFindBySQLQuery(packet)
	case MethodLaunchSession:
		protocol.handleLaunchSession(packet)
	case MethodUpdateSessionURL:
		protocol.handleUpdateSessionURL(packet)
	case MethodGetSessionURL:
		protocol.handleGetSessionURL(packet)
	case MethodGetState:
		protocol.handleGetState(packet)
	case MethodSetState:
		protocol.handleSetState(packet)
	case MethodReportStats:
		protocol.handleReportStats(packet)
	case MethodGetStats:
		protocol.handleGetStats(packet)
	case MethodDeleteGathering:
		protocol.handleDeleteGathering(packet)
	case MethodGetPendingDeletions:
		protocol.handleGetPendingDeletions(packet)
	case MethodDeleteFromDeletions:
		protocol.handleDeleteFromDeletions(packet)
	case MethodMigrateGatheringOwnershipV1:
		protocol.handleMigrateGatheringOwnershipV1(packet)
	case MethodFindByDescriptionLike:
		protocol.handleFindByDescriptionLike(packet)
	case MethodRegisterLocalURL:
		protocol.handleRegisterLocalURL(packet)
	case MethodRegisterLocalURLs:
		protocol.handleRegisterLocalURLs(packet)
	case MethodUpdateSessionHostV1:
		protocol.handleUpdateSessionHostV1(packet)
	case MethodGetSessionURLs:
		protocol.handleGetSessionURLs(packet)
	case MethodUpdateSessionHost:
		protocol.handleUpdateSessionHost(packet)
	case MethodUpdateGatheringOwnership:
		protocol.handleUpdateGatheringOwnership(packet)
	case MethodMigrateGatheringOwnership:
		protocol.handleMigrateGatheringOwnership(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported MatchMaking method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new Match Making protocol
func NewProtocol() *Protocol {
	return &Protocol{}
}
