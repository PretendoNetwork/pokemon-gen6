package globals

import (
	"math"
	"time"

	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	subscription_types "github.com/PretendoNetwork/nex-protocols-go/v2/subscription/types"
)

// THIS IS COMPLETELY ARBITARY
// PLEASE DO NOT ADD IT WHEN CONVERTING TO A DATABASE
const TimelineLength int = 100000

type SubscriptionDataHolder struct {
	CreationTime time.Time
	UpdatedTime  time.Time
	Data         subscription_types.SubscriptionData
}

type SubscriptionDataTimeline map[types.PID]SubscriptionDataHolder

func (tm SubscriptionDataTimeline) ClearData(owner types.PID, targetsTable SubscriptionDataTargets) {
	delete(tm, owner)
	delete(targetsTable, owner)

	HandleSubscriptionClearNotification(owner, true)
}

func (tm SubscriptionDataTimeline) CreateData(owner types.PID, data subscription_types.SubscriptionData, targetsTable SubscriptionDataTargets) *nex.Error {
	// TODO: should createdata overwrite if existing? assuming no
	_, exists := tm[owner]
	if exists {
		return nex.NewError(nex.ResultCodes.Core.InvalidArgument, "Subscription already exists for this pid")
	}

	// again, completely arbitrary, do not add when converting to database
	if len(tm) >= TimelineLength {
		oldest := int64(math.MaxInt64)
		oldestKey := types.PID(0)
		for key, value := range tm {
			if value.CreationTime.Unix() < oldest {
				oldest = value.CreationTime.Unix()
				oldestKey = key
			}
		}

		if oldestKey == types.PID(0) {
			// thats not ok
			panic("unaccounted for scenario, timeline length is over or equalling 10000 and the oldest key has no assigned owner")
		}

		delete(tm, oldestKey)
	}

	tm[owner] = SubscriptionDataHolder{
		CreationTime: time.Now(),
		UpdatedTime:  time.Now(),
		Data:         data.Copy().(subscription_types.SubscriptionData),
	}

	targetsTable.CreateTargets(owner)

	HandleSubscriptionChangeNotification(owner, true)

	return nil
}

func (tm SubscriptionDataTimeline) GetData(target types.PID) (*SubscriptionDataHolder, *nex.Error) {
	data, exists := tm[target]
	if !exists {
		return nil, nex.NewError(nex.ResultCodes.Core.InvalidArgument, "Subscription target does not exist")
	}

	return &data, nil
}

func (tm SubscriptionDataTimeline) HasData(owner types.PID) bool {
	_, exists := tm[owner]

	return exists
}

func (tm SubscriptionDataTimeline) UpdateData(owner types.PID, data subscription_types.SubscriptionData) *nex.Error {
	oldData, exists := tm[owner]
	if !exists {
		return nex.NewError(nex.ResultCodes.Core.InvalidArgument, "Subscription target does not exist")
	}

	holder := oldData
	holder.UpdatedTime = time.Now()
	holder.Data = data.Copy().(subscription_types.SubscriptionData)

	tm[owner] = holder

	HandleSubscriptionChangeNotification(owner, false)

	return nil
}
