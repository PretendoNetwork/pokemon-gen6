// Package protocol implements the Matchmake Extension protocol
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
	// ProtocolID is the protocol ID for the Matchmake Extension protocol
	ProtocolID = 0x6D

	// MethodCloseParticipation is the method ID for method CloseParticipation
	MethodCloseParticipation = 0x1

	// MethodOpenParticipation is the method ID for method OpenParticipation
	MethodOpenParticipation = 0x2

	// MethodAutoMatchmakePostpone is the method ID for method AutoMatchmakePostpone
	MethodAutoMatchmakePostpone = 0x3

	// MethodBrowseMatchmakeSession is the method ID for method BrowseMatchmakeSession
	MethodBrowseMatchmakeSession = 0x4

	// MethodBrowseMatchmakeSessionWithHostURLs is the method ID for method BrowseMatchmakeSessionWithHostURLs
	MethodBrowseMatchmakeSessionWithHostURLs = 0x5

	// MethodCreateMatchmakeSession is the method ID for method CreateMatchmakeSession
	MethodCreateMatchmakeSession = 0x6

	// MethodJoinMatchmakeSession is the method ID for method JoinMatchmakeSession
	MethodJoinMatchmakeSession = 0x7

	// MethodModifyCurrentGameAttribute is the method ID for method ModifyCurrentGameAttribute
	MethodModifyCurrentGameAttribute = 0x8

	// MethodUpdateNotificationData is the method ID for method UpdateNotificationData
	MethodUpdateNotificationData = 0x9

	// MethodGetFriendNotificationData is the method ID for method GetFriendNotificationData
	MethodGetFriendNotificationData = 0xA

	// MethodUpdateApplicationBuffer is the method ID for method UpdateApplicationBuffer
	MethodUpdateApplicationBuffer = 0xB

	// MethodUpdateMatchmakeSessionAttribute is the method ID for method UpdateMatchmakeSessionAttribute
	MethodUpdateMatchmakeSessionAttribute = 0xC

	// MethodGetlstFriendNotificationData is the method ID for method GetlstFriendNotificationData
	MethodGetlstFriendNotificationData = 0xD

	// MethodUpdateMatchmakeSession is the method ID for method UpdateMatchmakeSession
	MethodUpdateMatchmakeSession = 0xE

	// MethodAutoMatchmakeWithSearchCriteriaPostpone is the method ID for method AutoMatchmakeWithSearchCriteriaPostpone
	MethodAutoMatchmakeWithSearchCriteriaPostpone = 0xF

	// MethodGetPlayingSession is the method ID for method GetPlayingSession
	MethodGetPlayingSession = 0x10

	// MethodCreateCommunity is the method ID for method CreateCommunity
	MethodCreateCommunity = 0x11

	// MethodUpdateCommunity is the method ID for method UpdateCommunity
	MethodUpdateCommunity = 0x12

	// MethodJoinCommunity is the method ID for method JoinCommunity
	MethodJoinCommunity = 0x13

	// MethodFindCommunityByGatheringID is the method ID for method FindCommunityByGatheringID
	MethodFindCommunityByGatheringID = 0x14

	// MethodFindOfficialCommunity is the method ID for method FindOfficialCommunity
	MethodFindOfficialCommunity = 0x15

	// MethodFindCommunityByParticipant is the method ID for method FindCommunityByParticipant
	MethodFindCommunityByParticipant = 0x16

	// MethodUpdatePrivacySetting is the method ID for method UpdatePrivacySetting
	MethodUpdatePrivacySetting = 0x17

	// MethodGetMyBlockList is the method ID for method GetMyBlockList
	MethodGetMyBlockList = 0x18

	// MethodAddToBlockList is the method ID for method AddToBlockList
	MethodAddToBlockList = 0x19

	// MethodRemoveFromBlockList is the method ID for method RemoveFromBlockList
	MethodRemoveFromBlockList = 0x1A

	// MethodClearMyBlockList is the method ID for method ClearMyBlockList
	MethodClearMyBlockList = 0x1B

	// MethodReportViolation is the method ID for method ReportViolation
	MethodReportViolation = 0x1C

	// MethodIsViolationUser is the method ID for method IsViolationUser
	MethodIsViolationUser = 0x1D

	// MethodJoinMatchmakeSessionEx is the method ID for method JoinMatchmakeSessionEx
	MethodJoinMatchmakeSessionEx = 0x1E

	// MethodGetSimplePlayingSession is the method ID for method GetSimplePlayingSession
	MethodGetSimplePlayingSession = 0x1F

	// MethodGetSimpleCommunity is the method ID for method GetSimpleCommunity
	MethodGetSimpleCommunity = 0x20

	// MethodAutoMatchmakeWithGatheringIDPostpone is the method ID for method AutoMatchmakeWithGatheringIDPostpone
	MethodAutoMatchmakeWithGatheringIDPostpone = 0x21

	// MethodUpdateProgressScore is the method ID for method UpdateProgressScore
	MethodUpdateProgressScore = 0x22

	// MethodDebugNotifyEvent is the method ID for method DebugNotifyEvent
	MethodDebugNotifyEvent = 0x23

	// MethodGenerateMatchmakeSessionSystemPassword is the method ID for method GenerateMatchmakeSessionSystemPassword
	MethodGenerateMatchmakeSessionSystemPassword = 0x24

	// MethodClearMatchmakeSessionSystemPassword is the method ID for method ClearMatchmakeSessionSystemPassword
	MethodClearMatchmakeSessionSystemPassword = 0x25

	// MethodCreateMatchmakeSessionWithParam is the method ID for method CreateMatchmakeSessionWithParam
	MethodCreateMatchmakeSessionWithParam = 0x26

	// MethodJoinMatchmakeSessionWithParam is the method ID for method JoinMatchmakeSessionWithParam
	MethodJoinMatchmakeSessionWithParam = 0x27

	// MethodAutoMatchmakeWithParamPostpone is the method ID for method AutoMatchmakeWithParamPostpone
	MethodAutoMatchmakeWithParamPostpone = 0x28

	// MethodFindMatchmakeSessionByGatheringIDDetail is the method ID for method FindMatchmakeSessionByGatheringIDDetail
	MethodFindMatchmakeSessionByGatheringIDDetail = 0x29

	// MethodBrowseMatchmakeSessionNoHolder is the method ID for method BrowseMatchmakeSessionNoHolder
	MethodBrowseMatchmakeSessionNoHolder = 0x2A

	// MethodBrowseMatchmakeSessionWithHostURLsNoHolder is the method ID for method BrowseMatchmakeSessionWithHostURLsNoHolder
	MethodBrowseMatchmakeSessionWithHostURLsNoHolder = 0x2B

	// MethodUpdateMatchmakeSessionPart is the method ID for method UpdateMatchmakeSessionPart
	MethodUpdateMatchmakeSessionPart = 0x2C

	// MethodRequestMatchmaking is the method ID for method RequestMatchmaking
	MethodRequestMatchmaking = 0x2D

	// MethodWithdrawMatchmaking is the method ID for method WithdrawMatchmaking
	MethodWithdrawMatchmaking = 0x2E

	// MethodWithdrawMatchmakingAll is the method ID for method WithdrawMatchmakingAll
	MethodWithdrawMatchmakingAll = 0x2F

	// MethodFindMatchmakeSessionByGatheringID is the method ID for method FindMatchmakeSessionByGatheringID
	MethodFindMatchmakeSessionByGatheringID = 0x30

	// MethodFindMatchmakeSessionBySingleGatheringID is the method ID for method FindMatchmakeSessionBySingleGatheringID
	MethodFindMatchmakeSessionBySingleGatheringID = 0x31

	// MethodFindMatchmakeSessionByOwner is the method ID for method FindMatchmakeSessionByOwner
	MethodFindMatchmakeSessionByOwner = 0x32

	// MethodFindMatchmakeSessionByParticipant is the method ID for method FindMatchmakeSessionByParticipant
	MethodFindMatchmakeSessionByParticipant = 0x33

	// MethodBrowseMatchmakeSessionNoHolderNoResultRange is the method ID for method BrowseMatchmakeSessionNoHolderNoResultRange
	MethodBrowseMatchmakeSessionNoHolderNoResultRange = 0x34

	// MethodBrowseMatchmakeSessionWithHostURLsNoHolderNoResultRange is the method ID for method BrowseMatchmakeSessionWithHostURLsNoHolderNoResultRange
	MethodBrowseMatchmakeSessionWithHostURLsNoHolderNoResultRange = 0x35

	// MethodFindCommunityByOwner is the method ID for method FindCommunityByOwner
	MethodFindCommunityByOwner = 0x36
)

// Protocol stores all the RMC method handlers for the Matchmake Extension protocol and listens for requests
type Protocol struct {
	endpoint                                                nex.EndpointInterface
	CloseParticipation                                      func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	OpenParticipation                                       func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	AutoMatchmakePostpone                                   func(err error, packet nex.PacketInterface, callID uint32, anyGathering *types.AnyDataHolder, message *types.String) (*nex.RMCMessage, *nex.Error)
	BrowseMatchmakeSession                                  func(err error, packet nex.PacketInterface, callID uint32, searchCriteria *match_making_types.MatchmakeSessionSearchCriteria, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error)
	BrowseMatchmakeSessionWithHostURLs                      func(err error, packet nex.PacketInterface, callID uint32, searchCriteria *match_making_types.MatchmakeSessionSearchCriteria, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error)
	CreateMatchmakeSession                                  func(err error, packet nex.PacketInterface, callID uint32, anyGathering *types.AnyDataHolder, message *types.String, participationCount *types.PrimitiveU16) (*nex.RMCMessage, *nex.Error)
	JoinMatchmakeSession                                    func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, strMessage *types.String) (*nex.RMCMessage, *nex.Error)
	ModifyCurrentGameAttribute                              func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, attribIndex *types.PrimitiveU32, newValue *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	UpdateNotificationData                                  func(err error, packet nex.PacketInterface, callID uint32, uiType *types.PrimitiveU32, uiParam1 *types.PrimitiveU32, uiParam2 *types.PrimitiveU32, strParam *types.String) (*nex.RMCMessage, *nex.Error)
	GetFriendNotificationData                               func(err error, packet nex.PacketInterface, callID uint32, uiType *types.PrimitiveS32) (*nex.RMCMessage, *nex.Error)
	UpdateApplicationBuffer                                 func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, applicationBuffer *types.Buffer) (*nex.RMCMessage, *nex.Error)
	UpdateMatchmakeSessionAttribute                         func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, attribs *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, *nex.Error)
	GetlstFriendNotificationData                            func(err error, packet nex.PacketInterface, callID uint32, lstTypes *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, *nex.Error)
	UpdateMatchmakeSession                                  func(err error, packet nex.PacketInterface, callID uint32, anyGathering *types.AnyDataHolder) (*nex.RMCMessage, *nex.Error)
	AutoMatchmakeWithSearchCriteriaPostpone                 func(err error, packet nex.PacketInterface, callID uint32, lstSearchCriteria *types.List[*match_making_types.MatchmakeSessionSearchCriteria], anyGathering *types.AnyDataHolder, strMessage *types.String) (*nex.RMCMessage, *nex.Error)
	GetPlayingSession                                       func(err error, packet nex.PacketInterface, callID uint32, lstPID *types.List[*types.PID]) (*nex.RMCMessage, *nex.Error)
	CreateCommunity                                         func(err error, packet nex.PacketInterface, callID uint32, community *match_making_types.PersistentGathering, strMessage *types.String) (*nex.RMCMessage, *nex.Error)
	UpdateCommunity                                         func(err error, packet nex.PacketInterface, callID uint32, community *match_making_types.PersistentGathering) (*nex.RMCMessage, *nex.Error)
	JoinCommunity                                           func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, strMessage *types.String, strPassword *types.String) (*nex.RMCMessage, *nex.Error)
	FindCommunityByGatheringID                              func(err error, packet nex.PacketInterface, callID uint32, lstGID *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, *nex.Error)
	FindOfficialCommunity                                   func(err error, packet nex.PacketInterface, callID uint32, isAvailableOnly *types.PrimitiveBool, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error)
	FindCommunityByParticipant                              func(err error, packet nex.PacketInterface, callID uint32, pid *types.PID, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error)
	UpdatePrivacySetting                                    func(err error, packet nex.PacketInterface, callID uint32, onlineStatus *types.PrimitiveBool, participationCommunity *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error)
	GetMyBlockList                                          func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	AddToBlockList                                          func(err error, packet nex.PacketInterface, callID uint32, lstPrincipalID *types.List[*types.PID]) (*nex.RMCMessage, *nex.Error)
	RemoveFromBlockList                                     func(err error, packet nex.PacketInterface, callID uint32, lstPrincipalID *types.List[*types.PID]) (*nex.RMCMessage, *nex.Error)
	ClearMyBlockList                                        func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	ReportViolation                                         func(err error, packet nex.PacketInterface, callID uint32, pid *types.PID, userName *types.String, violationCode *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	IsViolationUser                                         func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	JoinMatchmakeSessionEx                                  func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, strMessage *types.String, dontCareMyBlockList *types.PrimitiveBool, participationCount *types.PrimitiveU16) (*nex.RMCMessage, *nex.Error)
	GetSimplePlayingSession                                 func(err error, packet nex.PacketInterface, callID uint32, listPID *types.List[*types.PID], includeLoginUser *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error)
	GetSimpleCommunity                                      func(err error, packet nex.PacketInterface, callID uint32, gatheringIDList *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, *nex.Error)
	AutoMatchmakeWithGatheringIDPostpone                    func(err error, packet nex.PacketInterface, callID uint32, lstGID *types.List[*types.PrimitiveU32], anyGathering *types.AnyDataHolder, strMessage *types.String) (*nex.RMCMessage, *nex.Error)
	UpdateProgressScore                                     func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, progressScore *types.PrimitiveU8) (*nex.RMCMessage, *nex.Error)
	DebugNotifyEvent                                        func(err error, packet nex.PacketInterface, callID uint32, pid *types.PID, mainType *types.PrimitiveU32, subType *types.PrimitiveU32, param1 *types.PrimitiveU64, param2 *types.PrimitiveU64, stringParam *types.String) (*nex.RMCMessage, *nex.Error)
	GenerateMatchmakeSessionSystemPassword                  func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	ClearMatchmakeSessionSystemPassword                     func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	CreateMatchmakeSessionWithParam                         func(err error, packet nex.PacketInterface, callID uint32, createMatchmakeSessionParam *match_making_types.CreateMatchmakeSessionParam) (*nex.RMCMessage, *nex.Error)
	JoinMatchmakeSessionWithParam                           func(err error, packet nex.PacketInterface, callID uint32, joinMatchmakeSessionParam *match_making_types.JoinMatchmakeSessionParam) (*nex.RMCMessage, *nex.Error)
	AutoMatchmakeWithParamPostpone                          func(err error, packet nex.PacketInterface, callID uint32, autoMatchmakeParam *match_making_types.AutoMatchmakeParam) (*nex.RMCMessage, *nex.Error)
	FindMatchmakeSessionByGatheringIDDetail                 func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	BrowseMatchmakeSessionNoHolder                          func(err error, packet nex.PacketInterface, callID uint32, searchCriteria *match_making_types.MatchmakeSessionSearchCriteria, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error)
	BrowseMatchmakeSessionWithHostURLsNoHolder              func(err error, packet nex.PacketInterface, callID uint32, searchCriteria *match_making_types.MatchmakeSessionSearchCriteria, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error)
	UpdateMatchmakeSessionPart                              func(err error, packet nex.PacketInterface, callID uint32, updateMatchmakeSessionParam *match_making_types.UpdateMatchmakeSessionParam) (*nex.RMCMessage, *nex.Error)
	RequestMatchmaking                                      func(err error, packet nex.PacketInterface, callID uint32, autoMatchmakeParam *match_making_types.AutoMatchmakeParam) (*nex.RMCMessage, *nex.Error)
	WithdrawMatchmaking                                     func(err error, packet nex.PacketInterface, callID uint32, requestID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)
	WithdrawMatchmakingAll                                  func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	FindMatchmakeSessionByGatheringID                       func(err error, packet nex.PacketInterface, callID uint32, lstGID *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, *nex.Error)
	FindMatchmakeSessionBySingleGatheringID                 func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	FindMatchmakeSessionByOwner                             func(err error, packet nex.PacketInterface, callID uint32, id *types.PrimitiveU32, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error)
	FindMatchmakeSessionByParticipant                       func(err error, packet nex.PacketInterface, callID uint32, param *match_making_types.FindMatchmakeSessionByParticipantParam) (*nex.RMCMessage, *nex.Error)
	BrowseMatchmakeSessionNoHolderNoResultRange             func(err error, packet nex.PacketInterface, callID uint32, searchCriteria *match_making_types.MatchmakeSessionSearchCriteria) (*nex.RMCMessage, *nex.Error)
	BrowseMatchmakeSessionWithHostURLsNoHolderNoResultRange func(err error, packet nex.PacketInterface, callID uint32, searchCriteria *match_making_types.MatchmakeSessionSearchCriteria) (*nex.RMCMessage, *nex.Error)
	FindCommunityByOwner                                    func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error) // TODO - Unknown request/response format
	Patches                                                 nex.ServiceProtocol
	PatchedMethods                                          []uint32
}

// Interface implements the methods present on the Matchmake Extension protocol struct
type Interface interface {
	Endpoint() nex.EndpointInterface
	SetEndpoint(endpoint nex.EndpointInterface)
	SetHandlerCloseParticipation(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error))
	SetHandlerOpenParticipation(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error))
	SetHandlerAutoMatchmakePostpone(handler func(err error, packet nex.PacketInterface, callID uint32, anyGathering *types.AnyDataHolder, message *types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerBrowseMatchmakeSession(handler func(err error, packet nex.PacketInterface, callID uint32, searchCriteria *match_making_types.MatchmakeSessionSearchCriteria, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error))
	SetHandlerBrowseMatchmakeSessionWithHostURLs(handler func(err error, packet nex.PacketInterface, callID uint32, searchCriteria *match_making_types.MatchmakeSessionSearchCriteria, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error))
	SetHandlerCreateMatchmakeSession(handler func(err error, packet nex.PacketInterface, callID uint32, anyGathering *types.AnyDataHolder, message *types.String, participationCount *types.PrimitiveU16) (*nex.RMCMessage, *nex.Error))
	SetHandlerJoinMatchmakeSession(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, strMessage *types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerModifyCurrentGameAttribute(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, attribIndex *types.PrimitiveU32, newValue *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error))
	SetHandlerUpdateNotificationData(handler func(err error, packet nex.PacketInterface, callID uint32, uiType *types.PrimitiveU32, uiParam1 *types.PrimitiveU32, uiParam2 *types.PrimitiveU32, strParam *types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetFriendNotificationData(handler func(err error, packet nex.PacketInterface, callID uint32, uiType *types.PrimitiveS32) (*nex.RMCMessage, *nex.Error))
	SetHandlerUpdateApplicationBuffer(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, applicationBuffer *types.Buffer) (*nex.RMCMessage, *nex.Error))
	SetHandlerUpdateMatchmakeSessionAttribute(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, attribs *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetlstFriendNotificationData(handler func(err error, packet nex.PacketInterface, callID uint32, lstTypes *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, *nex.Error))
	SetHandlerUpdateMatchmakeSession(handler func(err error, packet nex.PacketInterface, callID uint32, anyGathering *types.AnyDataHolder) (*nex.RMCMessage, *nex.Error))
	SetHandlerAutoMatchmakeWithSearchCriteriaPostpone(handler func(err error, packet nex.PacketInterface, callID uint32, lstSearchCriteria *types.List[*match_making_types.MatchmakeSessionSearchCriteria], anyGathering *types.AnyDataHolder, strMessage *types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetPlayingSession(handler func(err error, packet nex.PacketInterface, callID uint32, lstPID *types.List[*types.PID]) (*nex.RMCMessage, *nex.Error))
	SetHandlerCreateCommunity(handler func(err error, packet nex.PacketInterface, callID uint32, community *match_making_types.PersistentGathering, strMessage *types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerUpdateCommunity(handler func(err error, packet nex.PacketInterface, callID uint32, community *match_making_types.PersistentGathering) (*nex.RMCMessage, *nex.Error))
	SetHandlerJoinCommunity(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, strMessage *types.String, strPassword *types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerFindCommunityByGatheringID(handler func(err error, packet nex.PacketInterface, callID uint32, lstGID *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, *nex.Error))
	SetHandlerFindOfficialCommunity(handler func(err error, packet nex.PacketInterface, callID uint32, isAvailableOnly *types.PrimitiveBool, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error))
	SetHandlerFindCommunityByParticipant(handler func(err error, packet nex.PacketInterface, callID uint32, pid *types.PID, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error))
	SetHandlerUpdatePrivacySetting(handler func(err error, packet nex.PacketInterface, callID uint32, onlineStatus *types.PrimitiveBool, participationCommunity *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetMyBlockList(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error))
	SetHandlerAddToBlockList(handler func(err error, packet nex.PacketInterface, callID uint32, lstPrincipalID *types.List[*types.PID]) (*nex.RMCMessage, *nex.Error))
	SetHandlerRemoveFromBlockList(handler func(err error, packet nex.PacketInterface, callID uint32, lstPrincipalID *types.List[*types.PID]) (*nex.RMCMessage, *nex.Error))
	SetHandlerClearMyBlockList(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error))
	SetHandlerReportViolation(handler func(err error, packet nex.PacketInterface, callID uint32, pid *types.PID, userName *types.String, violationCode *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error))
	SetHandlerIsViolationUser(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error))
	SetHandlerJoinMatchmakeSessionEx(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, strMessage *types.String, dontCareMyBlockList *types.PrimitiveBool, participationCount *types.PrimitiveU16) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetSimplePlayingSession(handler func(err error, packet nex.PacketInterface, callID uint32, listPID *types.List[*types.PID], includeLoginUser *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetSimpleCommunity(handler func(err error, packet nex.PacketInterface, callID uint32, gatheringIDList *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, *nex.Error))
	SetHandlerAutoMatchmakeWithGatheringIDPostpone(handler func(err error, packet nex.PacketInterface, callID uint32, lstGID *types.List[*types.PrimitiveU32], anyGathering *types.AnyDataHolder, strMessage *types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerUpdateProgressScore(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, progressScore *types.PrimitiveU8) (*nex.RMCMessage, *nex.Error))
	SetHandlerDebugNotifyEvent(handler func(err error, packet nex.PacketInterface, callID uint32, pid *types.PID, mainType *types.PrimitiveU32, subType *types.PrimitiveU32, param1 *types.PrimitiveU64, param2 *types.PrimitiveU64, stringParam *types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerGenerateMatchmakeSessionSystemPassword(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error))
	SetHandlerClearMatchmakeSessionSystemPassword(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error))
	SetHandlerCreateMatchmakeSessionWithParam(handler func(err error, packet nex.PacketInterface, callID uint32, createMatchmakeSessionParam *match_making_types.CreateMatchmakeSessionParam) (*nex.RMCMessage, *nex.Error))
	SetHandlerJoinMatchmakeSessionWithParam(handler func(err error, packet nex.PacketInterface, callID uint32, joinMatchmakeSessionParam *match_making_types.JoinMatchmakeSessionParam) (*nex.RMCMessage, *nex.Error))
	SetHandlerAutoMatchmakeWithParamPostpone(handler func(err error, packet nex.PacketInterface, callID uint32, autoMatchmakeParam *match_making_types.AutoMatchmakeParam) (*nex.RMCMessage, *nex.Error))
	SetHandlerFindMatchmakeSessionByGatheringIDDetail(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error))
	SetHandlerBrowseMatchmakeSessionNoHolder(handler func(err error, packet nex.PacketInterface, callID uint32, searchCriteria *match_making_types.MatchmakeSessionSearchCriteria, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error))
	SetHandlerBrowseMatchmakeSessionWithHostURLsNoHolder(handler func(err error, packet nex.PacketInterface, callID uint32, searchCriteria *match_making_types.MatchmakeSessionSearchCriteria, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error))
	SetHandlerUpdateMatchmakeSessionPart(handler func(err error, packet nex.PacketInterface, callID uint32, updateMatchmakeSessionParam *match_making_types.UpdateMatchmakeSessionParam) (*nex.RMCMessage, *nex.Error))
	SetHandlerRequestMatchmaking(handler func(err error, packet nex.PacketInterface, callID uint32, autoMatchmakeParam *match_making_types.AutoMatchmakeParam) (*nex.RMCMessage, *nex.Error))
	SetHandlerWithdrawMatchmaking(handler func(err error, packet nex.PacketInterface, callID uint32, requestID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error))
	SetHandlerWithdrawMatchmakingAll(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error))
	SetHandlerFindMatchmakeSessionByGatheringID(handler func(err error, packet nex.PacketInterface, callID uint32, lstGID *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, *nex.Error))
	SetHandlerFindMatchmakeSessionBySingleGatheringID(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error))
	SetHandlerFindMatchmakeSessionByOwner(handler func(err error, packet nex.PacketInterface, callID uint32, id *types.PrimitiveU32, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error))
	SetHandlerFindMatchmakeSessionByParticipant(handler func(err error, packet nex.PacketInterface, callID uint32, param *match_making_types.FindMatchmakeSessionByParticipantParam) (*nex.RMCMessage, *nex.Error))
	SetHandlerBrowseMatchmakeSessionNoHolderNoResultRange(handler func(err error, packet nex.PacketInterface, callID uint32, searchCriteria *match_making_types.MatchmakeSessionSearchCriteria) (*nex.RMCMessage, *nex.Error))
	SetHandlerBrowseMatchmakeSessionWithHostURLsNoHolderNoResultRange(handler func(err error, packet nex.PacketInterface, callID uint32, searchCriteria *match_making_types.MatchmakeSessionSearchCriteria) (*nex.RMCMessage, *nex.Error))
	SetHandlerFindCommunityByOwner(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error))
}

// Endpoint returns the endpoint implementing the protocol
func (protocol *Protocol) Endpoint() nex.EndpointInterface {
	return protocol.endpoint
}

// SetEndpoint sets the endpoint implementing the protocol
func (protocol *Protocol) SetEndpoint(endpoint nex.EndpointInterface) {
	protocol.endpoint = endpoint
}

// SetHandlerCloseParticipation sets the handler for the CloseParticipation method
func (protocol *Protocol) SetHandlerCloseParticipation(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)) {
	protocol.CloseParticipation = handler
}

// SetHandlerOpenParticipation sets the handler for the OpenParticipation method
func (protocol *Protocol) SetHandlerOpenParticipation(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)) {
	protocol.OpenParticipation = handler
}

// SetHandlerAutoMatchmakePostpone sets the handler for the AutoMatchmakePostpone method
func (protocol *Protocol) SetHandlerAutoMatchmakePostpone(handler func(err error, packet nex.PacketInterface, callID uint32, anyGathering *types.AnyDataHolder, message *types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.AutoMatchmakePostpone = handler
}

// SetHandlerBrowseMatchmakeSession sets the handler for the BrowseMatchmakeSession method
func (protocol *Protocol) SetHandlerBrowseMatchmakeSession(handler func(err error, packet nex.PacketInterface, callID uint32, searchCriteria *match_making_types.MatchmakeSessionSearchCriteria, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error)) {
	protocol.BrowseMatchmakeSession = handler
}

// SetHandlerBrowseMatchmakeSessionWithHostURLs sets the handler for the BrowseMatchmakeSessionWithHostURLs method
func (protocol *Protocol) SetHandlerBrowseMatchmakeSessionWithHostURLs(handler func(err error, packet nex.PacketInterface, callID uint32, searchCriteria *match_making_types.MatchmakeSessionSearchCriteria, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error)) {
	protocol.BrowseMatchmakeSessionWithHostURLs = handler
}

// SetHandlerCreateMatchmakeSession sets the handler for the CreateMatchmakeSession method
func (protocol *Protocol) SetHandlerCreateMatchmakeSession(handler func(err error, packet nex.PacketInterface, callID uint32, anyGathering *types.AnyDataHolder, message *types.String, participationCount *types.PrimitiveU16) (*nex.RMCMessage, *nex.Error)) {
	protocol.CreateMatchmakeSession = handler
}

// SetHandlerJoinMatchmakeSession sets the handler for the JoinMatchmakeSession method
func (protocol *Protocol) SetHandlerJoinMatchmakeSession(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, strMessage *types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.JoinMatchmakeSession = handler
}

// SetHandlerModifyCurrentGameAttribute sets the handler for the ModifyCurrentGameAttribute method
func (protocol *Protocol) SetHandlerModifyCurrentGameAttribute(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, attribIndex *types.PrimitiveU32, newValue *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)) {
	protocol.ModifyCurrentGameAttribute = handler
}

// SetHandlerUpdateNotificationData sets the handler for the UpdateNotificationData method
func (protocol *Protocol) SetHandlerUpdateNotificationData(handler func(err error, packet nex.PacketInterface, callID uint32, uiType *types.PrimitiveU32, uiParam1 *types.PrimitiveU32, uiParam2 *types.PrimitiveU32, strParam *types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.UpdateNotificationData = handler
}

// SetHandlerGetFriendNotificationData sets the handler for the GetFriendNotificationData method
func (protocol *Protocol) SetHandlerGetFriendNotificationData(handler func(err error, packet nex.PacketInterface, callID uint32, uiType *types.PrimitiveS32) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetFriendNotificationData = handler
}

// SetHandlerUpdateApplicationBuffer sets the handler for the UpdateApplicationBuffer method
func (protocol *Protocol) SetHandlerUpdateApplicationBuffer(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, applicationBuffer *types.Buffer) (*nex.RMCMessage, *nex.Error)) {
	protocol.UpdateApplicationBuffer = handler
}

// SetHandlerUpdateMatchmakeSessionAttribute sets the handler for the UpdateMatchmakeSessionAttribute method
func (protocol *Protocol) SetHandlerUpdateMatchmakeSessionAttribute(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, attribs *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, *nex.Error)) {
	protocol.UpdateMatchmakeSessionAttribute = handler
}

// SetHandlerGetlstFriendNotificationData sets the handler for the GetlstFriendNotificationData method
func (protocol *Protocol) SetHandlerGetlstFriendNotificationData(handler func(err error, packet nex.PacketInterface, callID uint32, lstTypes *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetlstFriendNotificationData = handler
}

// SetHandlerUpdateMatchmakeSession sets the handler for the UpdateMatchmakeSession method
func (protocol *Protocol) SetHandlerUpdateMatchmakeSession(handler func(err error, packet nex.PacketInterface, callID uint32, anyGathering *types.AnyDataHolder) (*nex.RMCMessage, *nex.Error)) {
	protocol.UpdateMatchmakeSession = handler
}

// SetHandlerAutoMatchmakeWithSearchCriteriaPostpone sets the handler for the AutoMatchmakeWithSearchCriteriaPostpone method
func (protocol *Protocol) SetHandlerAutoMatchmakeWithSearchCriteriaPostpone(handler func(err error, packet nex.PacketInterface, callID uint32, lstSearchCriteria *types.List[*match_making_types.MatchmakeSessionSearchCriteria], anyGathering *types.AnyDataHolder, strMessage *types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.AutoMatchmakeWithSearchCriteriaPostpone = handler
}

// SetHandlerGetPlayingSession sets the handler for the GetPlayingSession method
func (protocol *Protocol) SetHandlerGetPlayingSession(handler func(err error, packet nex.PacketInterface, callID uint32, lstPID *types.List[*types.PID]) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetPlayingSession = handler
}

// SetHandlerCreateCommunity sets the handler for the CreateCommunity method
func (protocol *Protocol) SetHandlerCreateCommunity(handler func(err error, packet nex.PacketInterface, callID uint32, community *match_making_types.PersistentGathering, strMessage *types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.CreateCommunity = handler
}

// SetHandlerUpdateCommunity sets the handler for the UpdateCommunity method
func (protocol *Protocol) SetHandlerUpdateCommunity(handler func(err error, packet nex.PacketInterface, callID uint32, community *match_making_types.PersistentGathering) (*nex.RMCMessage, *nex.Error)) {
	protocol.UpdateCommunity = handler
}

// SetHandlerJoinCommunity sets the handler for the JoinCommunity method
func (protocol *Protocol) SetHandlerJoinCommunity(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, strMessage *types.String, strPassword *types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.JoinCommunity = handler
}

// SetHandlerFindCommunityByGatheringID sets the handler for the FindCommunityByGatheringID method
func (protocol *Protocol) SetHandlerFindCommunityByGatheringID(handler func(err error, packet nex.PacketInterface, callID uint32, lstGID *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, *nex.Error)) {
	protocol.FindCommunityByGatheringID = handler
}

// SetHandlerFindOfficialCommunity sets the handler for the FindOfficialCommunity method
func (protocol *Protocol) SetHandlerFindOfficialCommunity(handler func(err error, packet nex.PacketInterface, callID uint32, isAvailableOnly *types.PrimitiveBool, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error)) {
	protocol.FindOfficialCommunity = handler
}

// SetHandlerFindCommunityByParticipant sets the handler for the FindCommunityByParticipant method
func (protocol *Protocol) SetHandlerFindCommunityByParticipant(handler func(err error, packet nex.PacketInterface, callID uint32, pid *types.PID, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error)) {
	protocol.FindCommunityByParticipant = handler
}

// SetHandlerUpdatePrivacySetting sets the handler for the UpdatePrivacySetting method
func (protocol *Protocol) SetHandlerUpdatePrivacySetting(handler func(err error, packet nex.PacketInterface, callID uint32, onlineStatus *types.PrimitiveBool, participationCommunity *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error)) {
	protocol.UpdatePrivacySetting = handler
}

// SetHandlerGetMyBlockList sets the handler for the GetMyBlockList method
func (protocol *Protocol) SetHandlerGetMyBlockList(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetMyBlockList = handler
}

// SetHandlerAddToBlockList sets the handler for the AddToBlockList method
func (protocol *Protocol) SetHandlerAddToBlockList(handler func(err error, packet nex.PacketInterface, callID uint32, lstPrincipalID *types.List[*types.PID]) (*nex.RMCMessage, *nex.Error)) {
	protocol.AddToBlockList = handler
}

// SetHandlerRemoveFromBlockList sets the handler for the RemoveFromBlockList method
func (protocol *Protocol) SetHandlerRemoveFromBlockList(handler func(err error, packet nex.PacketInterface, callID uint32, lstPrincipalID *types.List[*types.PID]) (*nex.RMCMessage, *nex.Error)) {
	protocol.RemoveFromBlockList = handler
}

// SetHandlerClearMyBlockList sets the handler for the ClearMyBlockList method
func (protocol *Protocol) SetHandlerClearMyBlockList(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.ClearMyBlockList = handler
}

// SetHandlerReportViolation sets the handler for the ReportViolation method
func (protocol *Protocol) SetHandlerReportViolation(handler func(err error, packet nex.PacketInterface, callID uint32, pid *types.PID, userName *types.String, violationCode *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)) {
	protocol.ReportViolation = handler
}

// SetHandlerIsViolationUser sets the handler for the IsViolationUser method
func (protocol *Protocol) SetHandlerIsViolationUser(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.IsViolationUser = handler
}

// SetHandlerJoinMatchmakeSessionEx sets the handler for the JoinMatchmakeSessionEx method
func (protocol *Protocol) SetHandlerJoinMatchmakeSessionEx(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, strMessage *types.String, dontCareMyBlockList *types.PrimitiveBool, participationCount *types.PrimitiveU16) (*nex.RMCMessage, *nex.Error)) {
	protocol.JoinMatchmakeSessionEx = handler
}

// SetHandlerGetSimplePlayingSession sets the handler for the GetSimplePlayingSession method
func (protocol *Protocol) SetHandlerGetSimplePlayingSession(handler func(err error, packet nex.PacketInterface, callID uint32, listPID *types.List[*types.PID], includeLoginUser *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetSimplePlayingSession = handler
}

// SetHandlerGetSimpleCommunity sets the handler for the GetSimpleCommunity method
func (protocol *Protocol) SetHandlerGetSimpleCommunity(handler func(err error, packet nex.PacketInterface, callID uint32, gatheringIDList *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetSimpleCommunity = handler
}

// SetHandlerAutoMatchmakeWithGatheringIDPostpone sets the handler for the AutoMatchmakeWithGatheringIDPostpone method
func (protocol *Protocol) SetHandlerAutoMatchmakeWithGatheringIDPostpone(handler func(err error, packet nex.PacketInterface, callID uint32, lstGID *types.List[*types.PrimitiveU32], anyGathering *types.AnyDataHolder, strMessage *types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.AutoMatchmakeWithGatheringIDPostpone = handler
}

// SetHandlerUpdateProgressScore sets the handler for the UpdateProgressScore method
func (protocol *Protocol) SetHandlerUpdateProgressScore(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, progressScore *types.PrimitiveU8) (*nex.RMCMessage, *nex.Error)) {
	protocol.UpdateProgressScore = handler
}

// SetHandlerDebugNotifyEvent sets the handler for the DebugNotifyEvent method
func (protocol *Protocol) SetHandlerDebugNotifyEvent(handler func(err error, packet nex.PacketInterface, callID uint32, pid *types.PID, mainType *types.PrimitiveU32, subType *types.PrimitiveU32, param1 *types.PrimitiveU64, param2 *types.PrimitiveU64, stringParam *types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.DebugNotifyEvent = handler
}

// SetHandlerGenerateMatchmakeSessionSystemPassword sets the handler for the GenerateMatchmakeSessionSystemPassword method
func (protocol *Protocol) SetHandlerGenerateMatchmakeSessionSystemPassword(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)) {
	protocol.GenerateMatchmakeSessionSystemPassword = handler
}

// SetHandlerClearMatchmakeSessionSystemPassword sets the handler for the ClearMatchmakeSessionSystemPassword method
func (protocol *Protocol) SetHandlerClearMatchmakeSessionSystemPassword(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)) {
	protocol.ClearMatchmakeSessionSystemPassword = handler
}

// SetHandlerCreateMatchmakeSessionWithParam sets the handler for the CreateMatchmakeSessionWithParam method
func (protocol *Protocol) SetHandlerCreateMatchmakeSessionWithParam(handler func(err error, packet nex.PacketInterface, callID uint32, createMatchmakeSessionParam *match_making_types.CreateMatchmakeSessionParam) (*nex.RMCMessage, *nex.Error)) {
	protocol.CreateMatchmakeSessionWithParam = handler
}

// SetHandlerJoinMatchmakeSessionWithParam sets the handler for the JoinMatchmakeSessionWithParam method
func (protocol *Protocol) SetHandlerJoinMatchmakeSessionWithParam(handler func(err error, packet nex.PacketInterface, callID uint32, joinMatchmakeSessionParam *match_making_types.JoinMatchmakeSessionParam) (*nex.RMCMessage, *nex.Error)) {
	protocol.JoinMatchmakeSessionWithParam = handler
}

// SetHandlerAutoMatchmakeWithParamPostpone sets the handler for the AutoMatchmakeWithParamPostpone method
func (protocol *Protocol) SetHandlerAutoMatchmakeWithParamPostpone(handler func(err error, packet nex.PacketInterface, callID uint32, autoMatchmakeParam *match_making_types.AutoMatchmakeParam) (*nex.RMCMessage, *nex.Error)) {
	protocol.AutoMatchmakeWithParamPostpone = handler
}

// SetHandlerFindMatchmakeSessionByGatheringIDDetail sets the handler for the FindMatchmakeSessionByGatheringIDDetail method
func (protocol *Protocol) SetHandlerFindMatchmakeSessionByGatheringIDDetail(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)) {
	protocol.FindMatchmakeSessionByGatheringIDDetail = handler
}

// SetHandlerBrowseMatchmakeSessionNoHolder sets the handler for the BrowseMatchmakeSessionNoHolder method
func (protocol *Protocol) SetHandlerBrowseMatchmakeSessionNoHolder(handler func(err error, packet nex.PacketInterface, callID uint32, searchCriteria *match_making_types.MatchmakeSessionSearchCriteria, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error)) {
	protocol.BrowseMatchmakeSessionNoHolder = handler
}

// SetHandlerBrowseMatchmakeSessionWithHostURLsNoHolder sets the handler for the BrowseMatchmakeSessionWithHostURLsNoHolder method
func (protocol *Protocol) SetHandlerBrowseMatchmakeSessionWithHostURLsNoHolder(handler func(err error, packet nex.PacketInterface, callID uint32, searchCriteria *match_making_types.MatchmakeSessionSearchCriteria, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error)) {
	protocol.BrowseMatchmakeSessionWithHostURLsNoHolder = handler
}

// SetHandlerUpdateMatchmakeSessionPart sets the handler for the UpdateMatchmakeSessionPart method
func (protocol *Protocol) SetHandlerUpdateMatchmakeSessionPart(handler func(err error, packet nex.PacketInterface, callID uint32, updateMatchmakeSessionParam *match_making_types.UpdateMatchmakeSessionParam) (*nex.RMCMessage, *nex.Error)) {
	protocol.UpdateMatchmakeSessionPart = handler
}

// SetHandlerRequestMatchmaking sets the handler for the RequestMatchmaking method
func (protocol *Protocol) SetHandlerRequestMatchmaking(handler func(err error, packet nex.PacketInterface, callID uint32, autoMatchmakeParam *match_making_types.AutoMatchmakeParam) (*nex.RMCMessage, *nex.Error)) {
	protocol.RequestMatchmaking = handler
}

// SetHandlerWithdrawMatchmaking sets the handler for the WithdrawMatchmaking method
func (protocol *Protocol) SetHandlerWithdrawMatchmaking(handler func(err error, packet nex.PacketInterface, callID uint32, requestID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)) {
	protocol.WithdrawMatchmaking = handler
}

// SetHandlerWithdrawMatchmakingAll sets the handler for the WithdrawMatchmakingAll method
func (protocol *Protocol) SetHandlerWithdrawMatchmakingAll(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.WithdrawMatchmakingAll = handler
}

// SetHandlerFindMatchmakeSessionByGatheringID sets the handler for the FindMatchmakeSessionByGatheringID method
func (protocol *Protocol) SetHandlerFindMatchmakeSessionByGatheringID(handler func(err error, packet nex.PacketInterface, callID uint32, lstGID *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, *nex.Error)) {
	protocol.FindMatchmakeSessionByGatheringID = handler
}

// SetHandlerFindMatchmakeSessionBySingleGatheringID sets the handler for the FindMatchmakeSessionBySingleGatheringID method
func (protocol *Protocol) SetHandlerFindMatchmakeSessionBySingleGatheringID(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)) {
	protocol.FindMatchmakeSessionBySingleGatheringID = handler
}

// SetHandlerFindMatchmakeSessionByOwner sets the handler for the FindMatchmakeSessionByOwner method
func (protocol *Protocol) SetHandlerFindMatchmakeSessionByOwner(handler func(err error, packet nex.PacketInterface, callID uint32, id *types.PrimitiveU32, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error)) {
	protocol.FindMatchmakeSessionByOwner = handler
}

// SetHandlerFindMatchmakeSessionByParticipant sets the handler for the FindMatchmakeSessionByParticipant method
func (protocol *Protocol) SetHandlerFindMatchmakeSessionByParticipant(handler func(err error, packet nex.PacketInterface, callID uint32, param *match_making_types.FindMatchmakeSessionByParticipantParam) (*nex.RMCMessage, *nex.Error)) {
	protocol.FindMatchmakeSessionByParticipant = handler
}

// SetHandlerBrowseMatchmakeSessionNoHolderNoResultRange sets the handler for the BrowseMatchmakeSessionNoHolderNoResultRange method
func (protocol *Protocol) SetHandlerBrowseMatchmakeSessionNoHolderNoResultRange(handler func(err error, packet nex.PacketInterface, callID uint32, searchCriteria *match_making_types.MatchmakeSessionSearchCriteria) (*nex.RMCMessage, *nex.Error)) {
	protocol.BrowseMatchmakeSessionNoHolderNoResultRange = handler
}

// SetHandlerBrowseMatchmakeSessionWithHostURLsNoHolderNoResultRange sets the handler for the BrowseMatchmakeSessionWithHostURLsNoHolderNoResultRange method
func (protocol *Protocol) SetHandlerBrowseMatchmakeSessionWithHostURLsNoHolderNoResultRange(handler func(err error, packet nex.PacketInterface, callID uint32, searchCriteria *match_making_types.MatchmakeSessionSearchCriteria) (*nex.RMCMessage, *nex.Error)) {
	protocol.BrowseMatchmakeSessionWithHostURLsNoHolderNoResultRange = handler
}

// SetHandlerFindCommunityByOwner sets the handler for the FindCommunityByOwner method
func (protocol *Protocol) SetHandlerFindCommunityByOwner(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error)) {
	protocol.FindCommunityByOwner = handler
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
	case MethodCloseParticipation:
		protocol.handleCloseParticipation(packet)
	case MethodOpenParticipation:
		protocol.handleOpenParticipation(packet)
	case MethodAutoMatchmakePostpone:
		protocol.handleAutoMatchmakePostpone(packet)
	case MethodBrowseMatchmakeSession:
		protocol.handleBrowseMatchmakeSession(packet)
	case MethodBrowseMatchmakeSessionWithHostURLs:
		protocol.handleBrowseMatchmakeSessionWithHostURLs(packet)
	case MethodCreateMatchmakeSession:
		protocol.handleCreateMatchmakeSession(packet)
	case MethodJoinMatchmakeSession:
		protocol.handleJoinMatchmakeSession(packet)
	case MethodModifyCurrentGameAttribute:
		protocol.handleModifyCurrentGameAttribute(packet)
	case MethodUpdateNotificationData:
		protocol.handleUpdateNotificationData(packet)
	case MethodGetFriendNotificationData:
		protocol.handleGetFriendNotificationData(packet)
	case MethodUpdateApplicationBuffer:
		protocol.handleUpdateApplicationBuffer(packet)
	case MethodUpdateMatchmakeSessionAttribute:
		protocol.handleUpdateMatchmakeSessionAttribute(packet)
	case MethodGetlstFriendNotificationData:
		protocol.handleGetlstFriendNotificationData(packet)
	case MethodUpdateMatchmakeSession:
		protocol.handleUpdateMatchmakeSession(packet)
	case MethodAutoMatchmakeWithSearchCriteriaPostpone:
		protocol.handleAutoMatchmakeWithSearchCriteriaPostpone(packet)
	case MethodGetPlayingSession:
		protocol.handleGetPlayingSession(packet)
	case MethodCreateCommunity:
		protocol.handleCreateCommunity(packet)
	case MethodUpdateCommunity:
		protocol.handleUpdateCommunity(packet)
	case MethodJoinCommunity:
		protocol.handleJoinCommunity(packet)
	case MethodFindCommunityByGatheringID:
		protocol.handleFindCommunityByGatheringID(packet)
	case MethodFindOfficialCommunity:
		protocol.handleFindOfficialCommunity(packet)
	case MethodFindCommunityByParticipant:
		protocol.handleFindCommunityByParticipant(packet)
	case MethodUpdatePrivacySetting:
		protocol.handleUpdatePrivacySetting(packet)
	case MethodGetMyBlockList:
		protocol.handleGetMyBlockList(packet)
	case MethodAddToBlockList:
		protocol.handleAddToBlockList(packet)
	case MethodRemoveFromBlockList:
		protocol.handleRemoveFromBlockList(packet)
	case MethodClearMyBlockList:
		protocol.handleClearMyBlockList(packet)
	case MethodReportViolation:
		protocol.handleReportViolation(packet)
	case MethodIsViolationUser:
		protocol.handleIsViolationUser(packet)
	case MethodJoinMatchmakeSessionEx:
		protocol.handleJoinMatchmakeSessionEx(packet)
	case MethodGetSimplePlayingSession:
		protocol.handleGetSimplePlayingSession(packet)
	case MethodGetSimpleCommunity:
		protocol.handleGetSimpleCommunity(packet)
	case MethodAutoMatchmakeWithGatheringIDPostpone:
		protocol.handleAutoMatchmakeWithGatheringIDPostpone(packet)
	case MethodUpdateProgressScore:
		protocol.handleUpdateProgressScore(packet)
	case MethodDebugNotifyEvent:
		protocol.handleDebugNotifyEvent(packet)
	case MethodGenerateMatchmakeSessionSystemPassword:
		protocol.handleGenerateMatchmakeSessionSystemPassword(packet)
	case MethodClearMatchmakeSessionSystemPassword:
		protocol.handleClearMatchmakeSessionSystemPassword(packet)
	case MethodCreateMatchmakeSessionWithParam:
		protocol.handleCreateMatchmakeSessionWithParam(packet)
	case MethodJoinMatchmakeSessionWithParam:
		protocol.handleJoinMatchmakeSessionWithParam(packet)
	case MethodAutoMatchmakeWithParamPostpone:
		protocol.handleAutoMatchmakeWithParamPostpone(packet)
	case MethodFindMatchmakeSessionByGatheringIDDetail:
		protocol.handleFindMatchmakeSessionByGatheringIDDetail(packet)
	case MethodBrowseMatchmakeSessionNoHolder:
		protocol.handleBrowseMatchmakeSessionNoHolder(packet)
	case MethodBrowseMatchmakeSessionWithHostURLsNoHolder:
		protocol.handleBrowseMatchmakeSessionWithHostURLsNoHolder(packet)
	case MethodUpdateMatchmakeSessionPart:
		protocol.handleUpdateMatchmakeSessionPart(packet)
	case MethodRequestMatchmaking:
		protocol.handleRequestMatchmaking(packet)
	case MethodWithdrawMatchmaking:
		protocol.handleWithdrawMatchmaking(packet)
	case MethodWithdrawMatchmakingAll:
		protocol.handleWithdrawMatchmakingAll(packet)
	case MethodFindMatchmakeSessionByGatheringID:
		protocol.handleFindMatchmakeSessionByGatheringID(packet)
	case MethodFindMatchmakeSessionBySingleGatheringID:
		protocol.handleFindMatchmakeSessionBySingleGatheringID(packet)
	case MethodFindMatchmakeSessionByOwner:
		protocol.handleFindMatchmakeSessionByOwner(packet)
	case MethodFindMatchmakeSessionByParticipant:
		protocol.handleFindMatchmakeSessionByParticipant(packet)
	case MethodBrowseMatchmakeSessionNoHolderNoResultRange:
		protocol.handleBrowseMatchmakeSessionNoHolderNoResultRange(packet)
	case MethodBrowseMatchmakeSessionWithHostURLsNoHolderNoResultRange:
		protocol.handleBrowseMatchmakeSessionWithHostURLsNoHolderNoResultRange(packet)
	case MethodFindCommunityByOwner:
		protocol.handleFindCommunityByOwner(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported Matchmake Extension method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new Matchmake Extension protocol
func NewProtocol() *Protocol {
	return &Protocol{}
}
