package globals

import (
	"math"
	"time"

	"github.com/PretendoNetwork/nex-go/v2/types"
	common_globals "github.com/PretendoNetwork/nex-protocols-common-go/v2/globals"
	notification_types "github.com/PretendoNetwork/nex-protocols-go/v2/notifications/types"
	subscription_types "github.com/PretendoNetwork/nex-protocols-go/v2/subscription/types"
)

var Timeline SubscriptionDataTimeline
var DataTargets SubscriptionDataTargets

type SubscriptionDataHolder struct {
	CreationTime time.Time
	UpdatedTime  time.Time
	Data         subscription_types.SubscriptionData
	IsActive     bool
}

type SubscriptionTargets struct {
	PrincipalIDs types.List[types.PID]
}

type SubscriptionDataTimeline map[types.PID]SubscriptionDataHolder
type SubscriptionDataTargets map[types.PID]SubscriptionTargets

// THIS IS COMPLETELY ARBITARY
// PLEASE DO NOT ADD IT WHEN CONVERTING TO A DATABASE
const TimelineLength int = 100000

func (tm SubscriptionDataTimeline) ClearData(owner types.PID, targetsTable SubscriptionDataTargets) {
	// TODO: delete entry or set data to nil?
	delete(tm, owner)
	delete(targetsTable, owner)
}

func (tm SubscriptionDataTimeline) CreateData(owner types.PID, data subscription_types.SubscriptionData, targetsTable SubscriptionDataTargets) {
	// TODO: should create overwrite if existing? assuming no
	_, exists := tm[owner]
	if exists {
		return
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
		IsActive:     true,
	}

	targetsTable.CreateTargets(owner)
}

func (tm SubscriptionDataTimeline) GetData(target types.PID) SubscriptionDataHolder {
	_, exists := tm[target]
	if !exists {
		return SubscriptionDataHolder{}
	}

	return tm[target]
}

func (tm SubscriptionDataTimeline) HasData(owner types.PID) bool {
	_, exists := tm[owner]

	return exists
}

func (tm SubscriptionDataTimeline) UpdateData(owner types.PID, data subscription_types.SubscriptionData) {
	_, exists := tm[owner]
	if !exists {
		return
	}

	holder := tm[owner]
	holder.UpdatedTime = time.Now()
	holder.Data = data.Copy().(subscription_types.SubscriptionData)

	tm[owner] = holder
}

func (dt SubscriptionDataTargets) ClearTargets(owner types.PID) {
	delete(dt, owner)
}

func (dt SubscriptionDataTargets) CreateTargets(owner types.PID) {
	// TODO: should create overwrite if existing? assuming no
	_, exists := dt[owner]
	if exists {
		return
	}

	targets := SubscriptionTargets{
		PrincipalIDs: types.List[types.PID]{},
	}

	dt[owner] = targets
}

func (dt SubscriptionDataTargets) DeleteTargets(owner types.PID, targetsToDelete types.List[types.PID]) {
	// TODO: Implement
}

func (dt SubscriptionDataTargets) GetTargets(owner types.PID) types.List[types.PID] {
	_, exists := dt[owner]
	if !exists {
		return types.List[types.PID]{}
	}

	return dt[owner].PrincipalIDs
}

func (dt SubscriptionDataTargets) HasData(owner types.PID) bool {
	_, exists := dt[owner]

	if !exists {
		return false
	}

	return len(dt[owner].PrincipalIDs) > 0
}

func (dt SubscriptionDataTargets) ReplaceTargets(owner types.PID, newTargets types.List[types.PID]) {
	_, exists := dt[owner]
	if !exists {
		return
	}

	dt[owner] = SubscriptionTargets{newTargets.Copy().(types.List[types.PID])}
}

func HandleSubscriptionOfflineNotification(source types.PID) {
	handleSubscriptionNotification(source, 112000)
}

func HandleSubscriptionChangeNotification(source types.PID) {
	handleSubscriptionNotification(source, 112001)
}

// 112002 seems to trigger a GetActivePlayerSubscriptionData request from the client, notification sent from the secure account pid (2)
// 112003-112009 are either rare or do not exist

func handleSubscriptionNotification(source types.PID, notificationType uint32) {
	notificationEvent := notification_types.NewNotificationEvent()
	notificationEvent.PIDSource = source
	notificationEvent.Type = types.UInt32(notificationType)

	targets := GetOnlineFriendPIDs(uint32(source))

	// if DataTargets.HasData(source) {
	// 	for _, target := range DataTargets.GetTargets(source) {
	// 		targets = append(targets, uint64(target))
	// 	}
	// }

	common_globals.SendNotificationEvent(SecureEndpoint, notificationEvent, targets)
}
