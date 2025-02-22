package nex

import (
	"github.com/PretendoNetwork/pokemon-gen6/globals"
	datastore "github.com/PretendoNetwork/nex-protocols-go/v2/datastore"
	message_delivery "github.com/PretendoNetwork/nex-protocols-go/v2/message-delivery"
	subscription "github.com/PretendoNetwork/nex-protocols-go/v2/subscription"
	rating "github.com/PretendoNetwork/nex-protocols-go/v2/rating"

	nex_message_delivery "github.com/PretendoNetwork/pokemon-gen6/nex/message-delivery"
	nex_subscription "github.com/PretendoNetwork/pokemon-gen6/nex/subscription"
	nex_rating "github.com/PretendoNetwork/pokemon-gen6/nex/rating"
)

func registerSecureServerNEXProtocols() {
	datastoreProtocol := datastore.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(datastoreProtocol)

	messageDeliveryProtocol := message_delivery.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(messageDeliveryProtocol)

	messageDeliveryProtocol.DeliverMessage = nex_message_delivery.DeliverMessage

	subscriptionProtocol := subscription.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(subscriptionProtocol)

	subscriptionProtocol.CreateMySubscriptionData = nex_subscription.CreateMySubscriptionData
	subscriptionProtocol.UpdateMySubscriptionData = nex_subscription.UpdateMySubscriptionData
	subscriptionProtocol.GetFriendSubscriptionData = nex_subscription.GetFriendSubscriptionData
	subscriptionProtocol.GetTargetSubscriptionData = nex_subscription.GetTargetSubscriptionData
	subscriptionProtocol.GetActivePlayerSubscriptionData = nex_subscription.GetActivePlayerSubscriptionData
	subscriptionProtocol.GetSubscriptionData = nex_subscription.GetSubscriptionData
	subscriptionProtocol.ReplaceTargetAndGetSubscriptionData = nex_subscription.ReplaceTargetAndGetSubscriptionData
	subscriptionProtocol.GetPrivacyLevels = nex_subscription.GetPrivacyLevels

	ratingProtocol := rating.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(ratingProtocol)

	ratingProtocol.Unk1 = nex_rating.Unk1
	ratingProtocol.Unk2 = nex_rating.Unk2
}

