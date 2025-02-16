package nex_matchmake_extension_common

import (
	"fmt"

	matchmaking_types "github.com/PretendoNetwork/nex-protocols-go/v2/match-making/types"
)

func CleanupSearchMatchmakeSession(matchmakeSession *matchmaking_types.MatchmakeSession) {
	fmt.Println(matchmakeSession)
}