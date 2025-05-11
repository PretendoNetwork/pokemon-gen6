package nex

import (
	datastore "github.com/PretendoNetwork/nex-protocols-go/v2/datastore"
	message_delivery "github.com/PretendoNetwork/nex-protocols-go/v2/message-delivery"
	rating "github.com/PretendoNetwork/nex-protocols-go/v2/rating"
	subscription "github.com/PretendoNetwork/nex-protocols-go/v2/subscription"
	"github.com/PretendoNetwork/pokemon-gen6/globals"

	nex_message_delivery "github.com/PretendoNetwork/pokemon-gen6/nex/message-delivery"
	nex_rating "github.com/PretendoNetwork/pokemon-gen6/nex/rating"
	nex_subscription "github.com/PretendoNetwork/pokemon-gen6/nex/subscription"
	nex_subscription_custom_handlers "github.com/PretendoNetwork/pokemon-gen6/nex/subscription/custom_handlers"
)

func registerSecureServerNEXProtocols() {
	datastoreProtocol := datastore.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(datastoreProtocol)

	messageDeliveryProtocol := message_delivery.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(messageDeliveryProtocol)

	messageDeliveryProtocol.DeliverMessage = nex_message_delivery.DeliverMessage

	subscriptionProtocol := subscription.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(subscriptionProtocol)

	subscriptionProtocol.GetFriendSubscriptionData = nex_subscription.GetFriendSubscriptionData
	subscriptionProtocol.GetTargetSubscriptionData = nex_subscription.GetTargetSubscriptionData
	subscriptionProtocol.GetActivePlayerSubscriptionData = nex_subscription.GetActivePlayerSubscriptionData
	subscriptionProtocol.GetSubscriptionData = nex_subscription.GetSubscriptionData
	subscriptionProtocol.ReplaceTargetAndGetSubscriptionData = nex_subscription.ReplaceTargetAndGetSubscriptionData
	subscriptionProtocol.GetPrivacyLevels = nex_subscription.GetPrivacyLevels

	patches := nex_subscription_custom_handlers.NewProtocol()
	patches.CreateMySubscriptionData = nex_subscription.CreateMySubscriptionData
	patches.UpdateMySubscriptionData = nex_subscription.UpdateMySubscriptionData

	subscriptionProtocol.Patches = patches
	subscriptionProtocol.PatchedMethods = []uint32{1, 2}

	ratingProtocol := rating.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(ratingProtocol)

	ratingProtocol.Unk1 = nex_rating.Unk1
	ratingProtocol.Unk2 = nex_rating.Unk2
}
