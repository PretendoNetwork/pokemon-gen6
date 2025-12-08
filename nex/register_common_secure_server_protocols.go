package nex

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	common_globals "github.com/PretendoNetwork/nex-protocols-common-go/v2/globals"
	common_match_making "github.com/PretendoNetwork/nex-protocols-common-go/v2/match-making"
	common_match_making_ext "github.com/PretendoNetwork/nex-protocols-common-go/v2/match-making-ext"
	common_matchmake_extension "github.com/PretendoNetwork/nex-protocols-common-go/v2/matchmake-extension"
	common_message_delivery "github.com/PretendoNetwork/nex-protocols-common-go/v2/message-delivery"
	common_nat_traversal "github.com/PretendoNetwork/nex-protocols-common-go/v2/nat-traversal"
	common_secure "github.com/PretendoNetwork/nex-protocols-common-go/v2/secure-connection"
	common_utility "github.com/PretendoNetwork/nex-protocols-common-go/v2/utility"
	match_making "github.com/PretendoNetwork/nex-protocols-go/v2/match-making"
	match_making_ext "github.com/PretendoNetwork/nex-protocols-go/v2/match-making-ext"
	mm_types "github.com/PretendoNetwork/nex-protocols-go/v2/match-making/types"
	matchmake_extension "github.com/PretendoNetwork/nex-protocols-go/v2/matchmake-extension"
	message_delivery "github.com/PretendoNetwork/nex-protocols-go/v2/message-delivery"
	nat_traversal "github.com/PretendoNetwork/nex-protocols-go/v2/nat-traversal"
	secure "github.com/PretendoNetwork/nex-protocols-go/v2/secure-connection"
	utility "github.com/PretendoNetwork/nex-protocols-go/v2/utility"
	"github.com/PretendoNetwork/pokemon-gen6/database"
	database_utility "github.com/PretendoNetwork/pokemon-gen6/database/utility"
	"github.com/PretendoNetwork/pokemon-gen6/globals"

	nex_matchmake_extension_common "github.com/PretendoNetwork/pokemon-gen6/nex/matchmake-extension/common"
)

func registerCommonSecureServerProtocols() {
	secureProtocol := secure.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(secureProtocol)
	secure := common_secure.NewCommonProtocol(secureProtocol)
	secure.CreateReportDBRecord = func(pid types.PID, reportID types.UInt32, reportData types.QBuffer) error {
		return nil
	}

	secure.EnableInsecureRegister()

	globals.MatchmakingManager = common_globals.NewMatchmakingManager(globals.SecureEndpoint, database.Postgres)
	globals.MessagingManager = common_globals.NewMessagingManager(globals.SecureEndpoint, database.Postgres)

	natTraversalProtocol := nat_traversal.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(natTraversalProtocol)
	common_nat_traversal.NewCommonProtocol(natTraversalProtocol)

	utilityProtocol := utility.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(utilityProtocol)
	commonUtilityProtocol := common_utility.NewCommonProtocol(utilityProtocol)
	commonUtilityProtocol.GenerateNEXUniqueID = database_utility.GenerateNEXUniqueID

	matchMakingProtocol := match_making.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(matchMakingProtocol)
	common_match_making.NewCommonProtocol(matchMakingProtocol).SetManager(globals.MatchmakingManager)

	matchMakingExtProtocol := match_making_ext.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(matchMakingExtProtocol)
	common_match_making_ext.NewCommonProtocol(matchMakingExtProtocol).SetManager(globals.MatchmakingManager)

	matchmakeExtensionProtocol := matchmake_extension.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(matchmakeExtensionProtocol)
	commonMatchmakeExtensionProtocol := common_matchmake_extension.NewCommonProtocol(matchmakeExtensionProtocol)
	commonMatchmakeExtensionProtocol.SetManager(globals.MatchmakingManager)

	messageDeliveryProtocol := message_delivery.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(messageDeliveryProtocol)
	common_message_delivery.NewCommonProtocol(messageDeliveryProtocol).SetManager(globals.MessagingManager)

	commonMatchmakeExtensionProtocol.CleanupSearchMatchmakeSession = nex_matchmake_extension_common.CleanupSearchMatchmakeSession

	commonMatchmakeExtensionProtocol.OnAfterAutoMatchmakeWithSearchCriteriaPostpone = func(packet nex.PacketInterface, lstSearchCriteria types.List[mm_types.MatchmakeSessionSearchCriteria], anyGathering types.AnyObjectHolder[mm_types.GatheringInterface], strMessage types.String) {
		fmt.Println(anyGathering)
	}

	commonMatchmakeExtensionProtocol.CleanupMatchmakeSessionSearchCriterias = func(searchCriterias types.List[mm_types.MatchmakeSessionSearchCriteria]) {
		// Stubbed
	}
}
