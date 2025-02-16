package nex_matchmake_extension_common

import (
	"github.com/PretendoNetwork/nex-go/v2/types"
	matchmaking_types "github.com/PretendoNetwork/nex-protocols-go/v2/match-making/types"
)

func CleanupSearchMatchmakeSession(matchmakeSession *matchmaking_types.MatchmakeSession) {
	// Cleanup VR
	matchmakeSession.Attributes.SetIndex(1, types.NewPrimitiveU32(0))

	// Cleanup participation count
	matchmakeSession.ParticipationCount.Value = 0
}
