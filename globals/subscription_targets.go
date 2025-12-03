package globals

import (
	"github.com/PretendoNetwork/nex-go/v2/types"
)

type SubscriptionTargetsHolder struct {
	PrincipalIDs types.List[types.PID]
}

type SubscriptionDataTargets map[types.PID]SubscriptionTargetsHolder

func (dt SubscriptionDataTargets) ClearTargets(owner types.PID) {
	delete(dt, owner)
}

func (dt SubscriptionDataTargets) CreateTargets(owner types.PID) {
	// TODO: should create overwrite if existing? assuming no
	_, exists := dt[owner]
	if exists {
		return
	}

	targets := SubscriptionTargetsHolder{
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

	dt[owner] = SubscriptionTargetsHolder{newTargets.Copy().(types.List[types.PID])}
}
