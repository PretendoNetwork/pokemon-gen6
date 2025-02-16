package matchmake_extension

import (
	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/v2/match-making/types"
	matchmake_extension "github.com/PretendoNetwork/nex-protocols-go/v2/matchmake-extension"

	common_globals "github.com/PretendoNetwork/nex-protocols-common-go/v2/globals"
)

type CommonProtocol struct {
	endpoint                                         nex.EndpointInterface
	protocol                                         matchmake_extension.Interface
	CleanupSearchMatchmakeSession                    func(matchmakeSession *match_making_types.MatchmakeSession)
	GameSpecificMatchmakeSessionSearchCriteriaChecks func(searchCriteria *match_making_types.MatchmakeSessionSearchCriteria, matchmakeSession *match_making_types.MatchmakeSession) bool
	OnAfterOpenParticipation                         func(packet nex.PacketInterface, gid *types.PrimitiveU32)
	OnAfterCloseParticipation                        func(packet nex.PacketInterface, gid *types.PrimitiveU32)
	OnAfterCreateMatchmakeSession                    func(packet nex.PacketInterface, anyGathering *types.AnyDataHolder, message *types.String, participationCount *types.PrimitiveU16)
	OnAfterGetSimplePlayingSession                   func(packet nex.PacketInterface, listPID *types.List[*types.PID], includeLoginUser *types.PrimitiveBool)
	OnAfterAutoMatchmakePostpone                     func(packet nex.PacketInterface, anyGathering *types.AnyDataHolder, message *types.String)
	OnAfterAutoMatchmakeWithParamPostpone            func(packet nex.PacketInterface, autoMatchmakeParam *match_making_types.AutoMatchmakeParam)
	OnAfterAutoMatchmakeWithSearchCriteriaPostpone   func(packet nex.PacketInterface, lstSearchCriteria *types.List[*match_making_types.MatchmakeSessionSearchCriteria], anyGathering *types.AnyDataHolder, strMessage *types.String)
	OnAfterUpdateProgressScore                       func(packet nex.PacketInterface, gid *types.PrimitiveU32, progressScore *types.PrimitiveU8)
	OnAfterCreateMatchmakeSessionWithParam           func(packet nex.PacketInterface, createMatchmakeSessionParam *match_making_types.CreateMatchmakeSessionParam)
	OnAfterUpdateApplicationBuffer                   func(packet nex.PacketInterface, gid *types.PrimitiveU32, applicationBuffer *types.Buffer)
	OnAfterJoinMatchmakeSession                      func(packet nex.PacketInterface, gid *types.PrimitiveU32, strMessage *types.String)
	OnAfterJoinMatchmakeSessionWithParam             func(packet nex.PacketInterface, joinMatchmakeSessionParam *match_making_types.JoinMatchmakeSessionParam)
	OnAfterModifyCurrentGameAttribute                func(packet nex.PacketInterface, gid *types.PrimitiveU32, attribIndex *types.PrimitiveU32, newValue *types.PrimitiveU32)
	OnAfterBrowseMatchmakeSession                    func(packet nex.PacketInterface, searchCriteria *match_making_types.MatchmakeSessionSearchCriteria, resultRange *types.ResultRange)
	OnAfterJoinMatchmakeSessionEx                    func(packet nex.PacketInterface, gid *types.PrimitiveU32, strMessage *types.String, dontCareMyBlockList *types.PrimitiveBool, participationCount *types.PrimitiveU16)
}

// GetUserFriendPIDs sets the GetUserFriendPIDs handler function
func (commonProtocol *CommonProtocol) GetUserFriendPIDs(handler func(pid uint32) []uint32) {
	common_globals.GetUserFriendPIDsHandler = handler
}

// NewCommonProtocol returns a new CommonProtocol
func NewCommonProtocol(protocol matchmake_extension.Interface) *CommonProtocol {
	commonProtocol := &CommonProtocol{
		endpoint: protocol.Endpoint(),
		protocol: protocol,
	}

	protocol.SetHandlerOpenParticipation(commonProtocol.openParticipation)
	protocol.SetHandlerCloseParticipation(commonProtocol.closeParticipation)
	protocol.SetHandlerCreateMatchmakeSession(commonProtocol.createMatchmakeSession)
	protocol.SetHandlerGetSimplePlayingSession(commonProtocol.getSimplePlayingSession)
	protocol.SetHandlerAutoMatchmakePostpone(commonProtocol.autoMatchmakePostpone)
	protocol.SetHandlerAutoMatchmakeWithParamPostpone(commonProtocol.autoMatchmakeWithParamPostpone)
	protocol.SetHandlerAutoMatchmakeWithSearchCriteriaPostpone(commonProtocol.autoMatchmakeWithSearchCriteriaPostpone)
	protocol.SetHandlerUpdateProgressScore(commonProtocol.updateProgressScore)
	protocol.SetHandlerCreateMatchmakeSessionWithParam(commonProtocol.createMatchmakeSessionWithParam)
	protocol.SetHandlerUpdateApplicationBuffer(commonProtocol.updateApplicationBuffer)
	protocol.SetHandlerJoinMatchmakeSession(commonProtocol.joinMatchmakeSession)
	protocol.SetHandlerJoinMatchmakeSessionWithParam(commonProtocol.joinMatchmakeSessionWithParam)
	protocol.SetHandlerModifyCurrentGameAttribute(commonProtocol.modifyCurrentGameAttribute)
	protocol.SetHandlerBrowseMatchmakeSession(commonProtocol.browseMatchmakeSession)
	protocol.SetHandlerJoinMatchmakeSessionEx(commonProtocol.joinMatchmakeSessionEx)

	return commonProtocol
}
