package nex

import (
	"github.com/PretendoNetwork/pokemon-gen6/globals"
	matchmake_extension "github.com/PretendoNetwork/nex-protocols-go/v2/matchmake-extension"
	common_matchmake_extension "github.com/PretendoNetwork/nex-protocols-common-go/v2/matchmake-extension"
	match_making "github.com/PretendoNetwork/nex-protocols-go/v2/match-making"
	common_match_making "github.com/PretendoNetwork/nex-protocols-common-go/v2/match-making"
	match_making_ext "github.com/PretendoNetwork/nex-protocols-go/v2/match-making-ext"
	common_match_making_ext "github.com/PretendoNetwork/nex-protocols-common-go/v2/match-making-ext"
	nat_traversal "github.com/PretendoNetwork/nex-protocols-go/v2/nat-traversal"
	common_nat_traversal "github.com/PretendoNetwork/nex-protocols-common-go/v2/nat-traversal"
	secure "github.com/PretendoNetwork/nex-protocols-go/v2/secure-connection"
	common_secure "github.com/PretendoNetwork/nex-protocols-common-go/v2/secure-connection"
	utility "github.com/PretendoNetwork/nex-protocols-go/v2/utility"
	common_utility "github.com/PretendoNetwork/nex-protocols-common-go/v2/utility"

	nex_matchmake_extension_common "github.com/PretendoNetwork/pokemon-gen6/nex/matchmake-extension/common"
)

func GenerateNEXUniqueID() uint64 {
	return 0
}

func registerCommonSecureServerProtocols() {
	secureProtocol := secure.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(secureProtocol)
	common_secure.NewCommonProtocol(secureProtocol)

	natTraversalProtocol := nat_traversal.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(natTraversalProtocol)
	common_nat_traversal.NewCommonProtocol(natTraversalProtocol)

	matchMakingProtocol := match_making.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(matchMakingProtocol)
	common_match_making.NewCommonProtocol(matchMakingProtocol)

	matchMakingExtProtocol := match_making_ext.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(matchMakingExtProtocol)
	common_match_making_ext.NewCommonProtocol(matchMakingExtProtocol)

	matchmakeExtensionProtocol := matchmake_extension.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(matchmakeExtensionProtocol)
	commonMatchmakeExtensionProtocol := common_matchmake_extension.NewCommonProtocol(matchmakeExtensionProtocol)

	commonMatchmakeExtensionProtocol.CleanupSearchMatchmakeSession = nex_matchmake_extension_common.CleanupSearchMatchmakeSession

	utilityProtocol := utility.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(utilityProtocol)
	commonUtilityProtocol := common_utility.NewCommonProtocol(utilityProtocol)

	commonUtilityProtocol.GenerateNEXUniqueID = GenerateNEXUniqueID
}
