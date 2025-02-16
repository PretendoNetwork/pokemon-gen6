// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"
	"slices"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

const (
	// ProtocolID is the protocol ID for the DataStore protocol
	ProtocolID = 0x73

	// MethodPrepareGetObjectV1 is the method ID for the method PrepareGetObjectV1
	MethodPrepareGetObjectV1 = 0x1

	// MethodPreparePostObjectV1 is the method ID for the method PreparePostObjectV1
	MethodPreparePostObjectV1 = 0x2

	// MethodCompletePostObjectV1 is the method ID for the method CompletePostObjectV1
	MethodCompletePostObjectV1 = 0x3

	// MethodDeleteObject is the method ID for the method DeleteObject
	MethodDeleteObject = 0x4

	// MethodDeleteObjects is the method ID for the method DeleteObjects
	MethodDeleteObjects = 0x5

	// MethodChangeMetaV1 is the method ID for the method ChangeMetaV1
	MethodChangeMetaV1 = 0x6

	// MethodChangeMetasV1 is the method ID for the method ChangeMetasV1
	MethodChangeMetasV1 = 0x7

	// MethodGetMeta is the method ID for the method GetMeta
	MethodGetMeta = 0x8

	// MethodGetMetas is the method ID for the method GetMetas
	MethodGetMetas = 0x9

	// MethodPrepareUpdateObject is the method ID for the method PrepareUpdateObject
	MethodPrepareUpdateObject = 0xA

	// MethodCompleteUpdateObject is the method ID for the method CompleteUpdateObject
	MethodCompleteUpdateObject = 0xB

	// MethodSearchObject is the method ID for the method SearchObject
	MethodSearchObject = 0xC

	// MethodGetNotificationURL is the method ID for the method GetNotificationURL
	MethodGetNotificationURL = 0xD

	// MethodGetNewArrivedNotificationsV1 is the method ID for the method GetNewArrivedNotificationsV1
	MethodGetNewArrivedNotificationsV1 = 0xE

	// MethodRateObject is the method ID for the method RateObject
	MethodRateObject = 0xF

	// MethodGetRating is the method ID for the method GetRating
	MethodGetRating = 0x10

	// MethodGetRatings is the method ID for the method GetRatings
	MethodGetRatings = 0x11

	// MethodResetRating is the method ID for the method ResetRating
	MethodResetRating = 0x12

	// MethodResetRatings is the method ID for the method ResetRatings
	MethodResetRatings = 0x13

	// MethodGetSpecificMetaV1 is the method ID for the method GetSpecificMetaV1
	MethodGetSpecificMetaV1 = 0x14

	// MethodPostMetaBinary is the method ID for the method PostMetaBinary
	MethodPostMetaBinary = 0x15

	// MethodTouchObject is the method ID for the method TouchObject
	MethodTouchObject = 0x16

	// MethodGetRatingWithLog is the method ID for the method GetRatingWithLog
	MethodGetRatingWithLog = 0x17

	// MethodPreparePostObject is the method ID for the method PreparePostObject
	MethodPreparePostObject = 0x18

	// MethodPrepareGetObject is the method ID for the method PrepareGetObject
	MethodPrepareGetObject = 0x19

	// MethodCompletePostObject is the method ID for the method CompletePostObject
	MethodCompletePostObject = 0x1A

	// MethodGetNewArrivedNotifications is the method ID for the method GetNewArrivedNotifications
	MethodGetNewArrivedNotifications = 0x1B

	// MethodGetSpecificMeta is the method ID for the method GetSpecificMeta
	MethodGetSpecificMeta = 0x1C

	// MethodGetPersistenceInfo is the method ID for the method GetPersistenceInfo
	MethodGetPersistenceInfo = 0x1D

	// MethodGetPersistenceInfos is the method ID for the method GetPersistenceInfos
	MethodGetPersistenceInfos = 0x1E

	// MethodPerpetuateObject is the method ID for the method PerpetuateObject
	MethodPerpetuateObject = 0x1F

	// MethodUnperpetuateObject is the method ID for the method UnperpetuateObject
	MethodUnperpetuateObject = 0x20

	// MethodPrepareGetObjectOrMetaBinary is the method ID for the method PrepareGetObjectOrMetaBinary
	MethodPrepareGetObjectOrMetaBinary = 0x21

	// MethodGetPasswordInfo is the method ID for the method GetPasswordInfo
	MethodGetPasswordInfo = 0x22

	// MethodGetPasswordInfos is the method ID for the method GetPasswordInfos
	MethodGetPasswordInfos = 0x23

	// MethodGetMetasMultipleParam is the method ID for the method GetMetasMultipleParam
	MethodGetMetasMultipleParam = 0x24

	// MethodCompletePostObjects is the method ID for the method CompletePostObjects
	MethodCompletePostObjects = 0x25

	// MethodChangeMeta is the method ID for the method ChangeMeta
	MethodChangeMeta = 0x26

	// MethodChangeMetas is the method ID for the method ChangeMetas
	MethodChangeMetas = 0x27

	// MethodRateObjects is the method ID for the method RateObjects
	MethodRateObjects = 0x28

	// MethodPostMetaBinaryWithDataID is the method ID for the method PostMetaBinaryWithDataID
	MethodPostMetaBinaryWithDataID = 0x29

	// MethodPostMetaBinariesWithDataID is the method ID for the method PostMetaBinariesWithDataID
	MethodPostMetaBinariesWithDataID = 0x2A

	// MethodRateObjectWithPosting is the method ID for the method RateObjectWithPosting
	MethodRateObjectWithPosting = 0x2B

	// MethodRateObjectsWithPosting is the method ID for the method RateObjectsWithPosting
	MethodRateObjectsWithPosting = 0x2C

	// MethodGetObjectInfos is the method ID for the method GetObjectInfos
	MethodGetObjectInfos = 0x2D

	// MethodSearchObjectLight is the method ID for the method SearchObjectLight
	MethodSearchObjectLight = 0x2E
)

// Protocol stores all the RMC method handlers for the DataStore protocol and listens for requests
type Protocol struct {
	endpoint                     nex.EndpointInterface
	PrepareGetObjectV1           func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStorePrepareGetParamV1) (*nex.RMCMessage, *nex.Error)
	PreparePostObjectV1          func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStorePreparePostParamV1) (*nex.RMCMessage, *nex.Error)
	CompletePostObjectV1         func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreCompletePostParamV1) (*nex.RMCMessage, *nex.Error)
	DeleteObject                 func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreDeleteParam) (*nex.RMCMessage, *nex.Error)
	DeleteObjects                func(err error, packet nex.PacketInterface, callID uint32, params *types.List[*datastore_types.DataStoreDeleteParam], transactional *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error)
	ChangeMetaV1                 func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreChangeMetaParamV1) (*nex.RMCMessage, *nex.Error)
	ChangeMetasV1                func(err error, packet nex.PacketInterface, callID uint32, dataIDs *types.List[*types.PrimitiveU64], params *types.List[*datastore_types.DataStoreChangeMetaParamV1], transactional *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error)
	GetMeta                      func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreGetMetaParam) (*nex.RMCMessage, *nex.Error)
	GetMetas                     func(err error, packet nex.PacketInterface, callID uint32, dataIDs *types.List[*types.PrimitiveU64], param *datastore_types.DataStoreGetMetaParam) (*nex.RMCMessage, *nex.Error)
	PrepareUpdateObject          func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStorePrepareUpdateParam) (*nex.RMCMessage, *nex.Error)
	CompleteUpdateObject         func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreCompleteUpdateParam) (*nex.RMCMessage, *nex.Error)
	SearchObject                 func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreSearchParam) (*nex.RMCMessage, *nex.Error)
	GetNotificationURL           func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreGetNotificationURLParam) (*nex.RMCMessage, *nex.Error)
	GetNewArrivedNotificationsV1 func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreGetNewArrivedNotificationsParam) (*nex.RMCMessage, *nex.Error)
	RateObject                   func(err error, packet nex.PacketInterface, callID uint32, target *datastore_types.DataStoreRatingTarget, param *datastore_types.DataStoreRateObjectParam, fetchRatings *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error)
	GetRating                    func(err error, packet nex.PacketInterface, callID uint32, target *datastore_types.DataStoreRatingTarget, accessPassword *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)
	GetRatings                   func(err error, packet nex.PacketInterface, callID uint32, dataIDs *types.List[*types.PrimitiveU64], accessPassword *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)
	ResetRating                  func(err error, packet nex.PacketInterface, callID uint32, target *datastore_types.DataStoreRatingTarget, accessPassword *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)
	ResetRatings                 func(err error, packet nex.PacketInterface, callID uint32, target *datastore_types.DataStoreRatingTarget, transactional *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error)
	GetSpecificMetaV1            func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreGetSpecificMetaParamV1) (*nex.RMCMessage, *nex.Error)
	PostMetaBinary               func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStorePreparePostParam) (*nex.RMCMessage, *nex.Error)
	TouchObject                  func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreTouchObjectParam) (*nex.RMCMessage, *nex.Error)
	GetRatingWithLog             func(err error, packet nex.PacketInterface, callID uint32, target *datastore_types.DataStoreRatingTarget, accessPassword *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)
	PreparePostObject            func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStorePreparePostParam) (*nex.RMCMessage, *nex.Error)
	PrepareGetObject             func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStorePrepareGetParam) (*nex.RMCMessage, *nex.Error)
	CompletePostObject           func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreCompletePostParam) (*nex.RMCMessage, *nex.Error)
	GetNewArrivedNotifications   func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreGetNewArrivedNotificationsParam) (*nex.RMCMessage, *nex.Error)
	GetSpecificMeta              func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreGetSpecificMetaParam) (*nex.RMCMessage, *nex.Error)
	GetPersistenceInfo           func(err error, packet nex.PacketInterface, callID uint32, ownerID *types.PID, persistenceSlotID *types.PrimitiveU16) (*nex.RMCMessage, *nex.Error)
	GetPersistenceInfos          func(err error, packet nex.PacketInterface, callID uint32, ownerID *types.PID, persistenceSlotIDs *types.List[*types.PrimitiveU16]) (*nex.RMCMessage, *nex.Error)
	PerpetuateObject             func(err error, packet nex.PacketInterface, callID uint32, persistenceSlotID *types.PrimitiveU16, dataID *types.PrimitiveU64, deleteLastObject *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error)
	UnperpetuateObject           func(err error, packet nex.PacketInterface, callID uint32, persistenceSlotID *types.PrimitiveU16, deleteLastObject *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error)
	PrepareGetObjectOrMetaBinary func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStorePrepareGetParam) (*nex.RMCMessage, *nex.Error)
	GetPasswordInfo              func(err error, packet nex.PacketInterface, callID uint32, dataID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)
	GetPasswordInfos             func(err error, packet nex.PacketInterface, callID uint32, dataIDs *types.List[*types.PrimitiveU64]) (*nex.RMCMessage, *nex.Error)
	GetMetasMultipleParam        func(err error, packet nex.PacketInterface, callID uint32, params *types.List[*datastore_types.DataStoreGetMetaParam]) (*nex.RMCMessage, *nex.Error)
	CompletePostObjects          func(err error, packet nex.PacketInterface, callID uint32, dataIDs *types.List[*types.PrimitiveU64]) (*nex.RMCMessage, *nex.Error)
	ChangeMeta                   func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreChangeMetaParam) (*nex.RMCMessage, *nex.Error)
	ChangeMetas                  func(err error, packet nex.PacketInterface, callID uint32, dataIDs *types.List[*types.PrimitiveU64], params *types.List[*datastore_types.DataStoreChangeMetaParam], transactional *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error)
	RateObjects                  func(err error, packet nex.PacketInterface, callID uint32, targets *types.List[*datastore_types.DataStoreRatingTarget], params *types.List[*datastore_types.DataStoreRateObjectParam], transactional *types.PrimitiveBool, fetchRatings *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error)
	PostMetaBinaryWithDataID     func(err error, packet nex.PacketInterface, callID uint32, dataID *types.PrimitiveU64, param *datastore_types.DataStorePreparePostParam) (*nex.RMCMessage, *nex.Error)
	PostMetaBinariesWithDataID   func(err error, packet nex.PacketInterface, callID uint32, dataIDs *types.List[*types.PrimitiveU64], params *types.List[*datastore_types.DataStorePreparePostParam], transactional *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error)
	RateObjectWithPosting        func(err error, packet nex.PacketInterface, callID uint32, target *datastore_types.DataStoreRatingTarget, rateParam *datastore_types.DataStoreRateObjectParam, postParam *datastore_types.DataStorePreparePostParam, fetchRatings *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error)
	RateObjectsWithPosting       func(err error, packet nex.PacketInterface, callID uint32, targets *types.List[*datastore_types.DataStoreRatingTarget], rateParams *types.List[*datastore_types.DataStoreRateObjectParam], postParams *types.List[*datastore_types.DataStorePreparePostParam], transactional *types.PrimitiveBool, fetchRatings *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error)
	GetObjectInfos               func(err error, packet nex.PacketInterface, callID uint32, dataIDs *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)
	SearchObjectLight            func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreSearchParam) (*nex.RMCMessage, *nex.Error)
	Patches                      nex.ServiceProtocol
	PatchedMethods               []uint32
}

// Interface implements the methods present on the DataStore Protocol struct
type Interface interface {
	Endpoint() nex.EndpointInterface
	SetEndpoint(endpoint nex.EndpointInterface)
	SetHandlerPrepareGetObjectV1(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStorePrepareGetParamV1) (*nex.RMCMessage, *nex.Error))
	SetHandlerPreparePostObjectV1(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStorePreparePostParamV1) (*nex.RMCMessage, *nex.Error))
	SetHandlerCompletePostObjectV1(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreCompletePostParamV1) (*nex.RMCMessage, *nex.Error))
	SetHandlerDeleteObject(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreDeleteParam) (*nex.RMCMessage, *nex.Error))
	SetHandlerDeleteObjects(handler func(err error, packet nex.PacketInterface, callID uint32, params *types.List[*datastore_types.DataStoreDeleteParam], transactional *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error))
	SetHandlerChangeMetaV1(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreChangeMetaParamV1) (*nex.RMCMessage, *nex.Error))
	SetHandlerChangeMetasV1(handler func(err error, packet nex.PacketInterface, callID uint32, dataIDs *types.List[*types.PrimitiveU64], params *types.List[*datastore_types.DataStoreChangeMetaParamV1], transactional *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetMeta(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreGetMetaParam) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetMetas(handler func(err error, packet nex.PacketInterface, callID uint32, dataIDs *types.List[*types.PrimitiveU64], param *datastore_types.DataStoreGetMetaParam) (*nex.RMCMessage, *nex.Error))
	SetHandlerPrepareUpdateObject(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStorePrepareUpdateParam) (*nex.RMCMessage, *nex.Error))
	SetHandlerCompleteUpdateObject(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreCompleteUpdateParam) (*nex.RMCMessage, *nex.Error))
	SetHandlerSearchObject(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreSearchParam) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetNotificationURL(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreGetNotificationURLParam) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetNewArrivedNotificationsV1(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreGetNewArrivedNotificationsParam) (*nex.RMCMessage, *nex.Error))
	SetHandlerRateObject(handler func(err error, packet nex.PacketInterface, callID uint32, target *datastore_types.DataStoreRatingTarget, param *datastore_types.DataStoreRateObjectParam, fetchRatings *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetRating(handler func(err error, packet nex.PacketInterface, callID uint32, target *datastore_types.DataStoreRatingTarget, accessPassword *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetRatings(handler func(err error, packet nex.PacketInterface, callID uint32, dataIDs *types.List[*types.PrimitiveU64], accessPassword *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error))
	SetHandlerResetRating(handler func(err error, packet nex.PacketInterface, callID uint32, target *datastore_types.DataStoreRatingTarget, accessPassword *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error))
	SetHandlerResetRatings(handler func(err error, packet nex.PacketInterface, callID uint32, target *datastore_types.DataStoreRatingTarget, transactional *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetSpecificMetaV1(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreGetSpecificMetaParamV1) (*nex.RMCMessage, *nex.Error))
	SetHandlerPostMetaBinary(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStorePreparePostParam) (*nex.RMCMessage, *nex.Error))
	SetHandlerTouchObject(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreTouchObjectParam) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetRatingWithLog(handler func(err error, packet nex.PacketInterface, callID uint32, target *datastore_types.DataStoreRatingTarget, accessPassword *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error))
	SetHandlerPreparePostObject(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStorePreparePostParam) (*nex.RMCMessage, *nex.Error))
	SetHandlerPrepareGetObject(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStorePrepareGetParam) (*nex.RMCMessage, *nex.Error))
	SetHandlerCompletePostObject(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreCompletePostParam) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetNewArrivedNotifications(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreGetNewArrivedNotificationsParam) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetSpecificMeta(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreGetSpecificMetaParam) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetPersistenceInfo(handler func(err error, packet nex.PacketInterface, callID uint32, ownerID *types.PID, persistenceSlotID *types.PrimitiveU16) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetPersistenceInfos(handler func(err error, packet nex.PacketInterface, callID uint32, ownerID *types.PID, persistenceSlotIDs *types.List[*types.PrimitiveU16]) (*nex.RMCMessage, *nex.Error))
	SetHandlerPerpetuateObject(handler func(err error, packet nex.PacketInterface, callID uint32, persistenceSlotID *types.PrimitiveU16, dataID *types.PrimitiveU64, deleteLastObject *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error))
	SetHandlerUnperpetuateObject(handler func(err error, packet nex.PacketInterface, callID uint32, persistenceSlotID *types.PrimitiveU16, deleteLastObject *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error))
	SetHandlerPrepareGetObjectOrMetaBinary(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStorePrepareGetParam) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetPasswordInfo(handler func(err error, packet nex.PacketInterface, callID uint32, dataID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetPasswordInfos(handler func(err error, packet nex.PacketInterface, callID uint32, dataIDs *types.List[*types.PrimitiveU64]) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetMetasMultipleParam(handler func(err error, packet nex.PacketInterface, callID uint32, params *types.List[*datastore_types.DataStoreGetMetaParam]) (*nex.RMCMessage, *nex.Error))
	SetHandlerCompletePostObjects(handler func(err error, packet nex.PacketInterface, callID uint32, dataIDs *types.List[*types.PrimitiveU64]) (*nex.RMCMessage, *nex.Error))
	SetHandlerChangeMeta(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreChangeMetaParam) (*nex.RMCMessage, *nex.Error))
	SetHandlerChangeMetas(handler func(err error, packet nex.PacketInterface, callID uint32, dataIDs *types.List[*types.PrimitiveU64], params *types.List[*datastore_types.DataStoreChangeMetaParam], transactional *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error))
	SetHandlerRateObjects(handler func(err error, packet nex.PacketInterface, callID uint32, targets *types.List[*datastore_types.DataStoreRatingTarget], params *types.List[*datastore_types.DataStoreRateObjectParam], transactional *types.PrimitiveBool, fetchRatings *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error))
	SetHandlerPostMetaBinaryWithDataID(handler func(err error, packet nex.PacketInterface, callID uint32, dataID *types.PrimitiveU64, param *datastore_types.DataStorePreparePostParam) (*nex.RMCMessage, *nex.Error))
	SetHandlerPostMetaBinariesWithDataID(handler func(err error, packet nex.PacketInterface, callID uint32, dataIDs *types.List[*types.PrimitiveU64], params *types.List[*datastore_types.DataStorePreparePostParam], transactional *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error))
	SetHandlerRateObjectWithPosting(handler func(err error, packet nex.PacketInterface, callID uint32, target *datastore_types.DataStoreRatingTarget, rateParam *datastore_types.DataStoreRateObjectParam, postParam *datastore_types.DataStorePreparePostParam, fetchRatings *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error))
	SetHandlerRateObjectsWithPosting(handler func(err error, packet nex.PacketInterface, callID uint32, targets *types.List[*datastore_types.DataStoreRatingTarget], rateParams *types.List[*datastore_types.DataStoreRateObjectParam], postParams *types.List[*datastore_types.DataStorePreparePostParam], transactional *types.PrimitiveBool, fetchRatings *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetObjectInfos(handler func(err error, packet nex.PacketInterface, callID uint32, dataIDs *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error))
	SetHandlerSearchObjectLight(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreSearchParam) (*nex.RMCMessage, *nex.Error))
}

// Endpoint returns the endpoint implementing the protocol
func (protocol *Protocol) Endpoint() nex.EndpointInterface {
	return protocol.endpoint
}

// SetEndpoint sets the endpoint implementing the protocol
func (protocol *Protocol) SetEndpoint(endpoint nex.EndpointInterface) {
	protocol.endpoint = endpoint
}

// SetHandlerPrepareGetObjectV1 sets the handler for the PrepareGetObjectV1 method
func (protocol *Protocol) SetHandlerPrepareGetObjectV1(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStorePrepareGetParamV1) (*nex.RMCMessage, *nex.Error)) {
	protocol.PrepareGetObjectV1 = handler
}

// SetHandlerPreparePostObjectV1 sets the handler for the PreparePostObjectV1 method
func (protocol *Protocol) SetHandlerPreparePostObjectV1(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStorePreparePostParamV1) (*nex.RMCMessage, *nex.Error)) {
	protocol.PreparePostObjectV1 = handler
}

// SetHandlerCompletePostObjectV1 sets the handler for the CompletePostObjectV1 method
func (protocol *Protocol) SetHandlerCompletePostObjectV1(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreCompletePostParamV1) (*nex.RMCMessage, *nex.Error)) {
	protocol.CompletePostObjectV1 = handler
}

// SetHandlerDeleteObject sets the handler for the DeleteObject method
func (protocol *Protocol) SetHandlerDeleteObject(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreDeleteParam) (*nex.RMCMessage, *nex.Error)) {
	protocol.DeleteObject = handler
}

// SetHandlerDeleteObjects sets the handler for the DeleteObjects method
func (protocol *Protocol) SetHandlerDeleteObjects(handler func(err error, packet nex.PacketInterface, callID uint32, params *types.List[*datastore_types.DataStoreDeleteParam], transactional *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error)) {
	protocol.DeleteObjects = handler
}

// SetHandlerChangeMetaV1 sets the handler for the ChangeMetaV1 method
func (protocol *Protocol) SetHandlerChangeMetaV1(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreChangeMetaParamV1) (*nex.RMCMessage, *nex.Error)) {
	protocol.ChangeMetaV1 = handler
}

// SetHandlerChangeMetasV1 sets the handler for the ChangeMetasV1 method
func (protocol *Protocol) SetHandlerChangeMetasV1(handler func(err error, packet nex.PacketInterface, callID uint32, dataIDs *types.List[*types.PrimitiveU64], params *types.List[*datastore_types.DataStoreChangeMetaParamV1], transactional *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error)) {
	protocol.ChangeMetasV1 = handler
}

// SetHandlerGetMeta sets the handler for the GetMeta method
func (protocol *Protocol) SetHandlerGetMeta(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreGetMetaParam) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetMeta = handler
}

// SetHandlerGetMetas sets the handler for the GetMetas method
func (protocol *Protocol) SetHandlerGetMetas(handler func(err error, packet nex.PacketInterface, callID uint32, dataIDs *types.List[*types.PrimitiveU64], param *datastore_types.DataStoreGetMetaParam) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetMetas = handler
}

// SetHandlerPrepareUpdateObject sets the handler for the PrepareUpdateObject method
func (protocol *Protocol) SetHandlerPrepareUpdateObject(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStorePrepareUpdateParam) (*nex.RMCMessage, *nex.Error)) {
	protocol.PrepareUpdateObject = handler
}

// SetHandlerCompleteUpdateObject sets the handler for the CompleteUpdateObject method
func (protocol *Protocol) SetHandlerCompleteUpdateObject(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreCompleteUpdateParam) (*nex.RMCMessage, *nex.Error)) {
	protocol.CompleteUpdateObject = handler
}

// SetHandlerSearchObject sets the handler for the SearchObject method
func (protocol *Protocol) SetHandlerSearchObject(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreSearchParam) (*nex.RMCMessage, *nex.Error)) {
	protocol.SearchObject = handler
}

// SetHandlerGetNotificationURL sets the handler for the GetNotificationURL method
func (protocol *Protocol) SetHandlerGetNotificationURL(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreGetNotificationURLParam) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetNotificationURL = handler
}

// SetHandlerGetNewArrivedNotificationsV1 sets the handler for the GetNewArrivedNotificationsV1 method
func (protocol *Protocol) SetHandlerGetNewArrivedNotificationsV1(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreGetNewArrivedNotificationsParam) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetNewArrivedNotificationsV1 = handler
}

// SetHandlerRateObject sets the handler for the RateObject method
func (protocol *Protocol) SetHandlerRateObject(handler func(err error, packet nex.PacketInterface, callID uint32, target *datastore_types.DataStoreRatingTarget, param *datastore_types.DataStoreRateObjectParam, fetchRatings *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error)) {
	protocol.RateObject = handler
}

// SetHandlerGetRating sets the handler for the GetRating method
func (protocol *Protocol) SetHandlerGetRating(handler func(err error, packet nex.PacketInterface, callID uint32, target *datastore_types.DataStoreRatingTarget, accessPassword *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetRating = handler
}

// SetHandlerGetRatings sets the handler for the GetRatings method
func (protocol *Protocol) SetHandlerGetRatings(handler func(err error, packet nex.PacketInterface, callID uint32, dataIDs *types.List[*types.PrimitiveU64], accessPassword *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetRatings = handler
}

// SetHandlerResetRating sets the handler for the ResetRating method
func (protocol *Protocol) SetHandlerResetRating(handler func(err error, packet nex.PacketInterface, callID uint32, target *datastore_types.DataStoreRatingTarget, accessPassword *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)) {
	protocol.ResetRating = handler
}

// SetHandlerResetRatings sets the handler for the ResetRatings method
func (protocol *Protocol) SetHandlerResetRatings(handler func(err error, packet nex.PacketInterface, callID uint32, target *datastore_types.DataStoreRatingTarget, transactional *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error)) {
	protocol.ResetRatings = handler
}

// SetHandlerGetSpecificMetaV1 sets the handler for the GetSpecificMetaV1 method
func (protocol *Protocol) SetHandlerGetSpecificMetaV1(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreGetSpecificMetaParamV1) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetSpecificMetaV1 = handler
}

// SetHandlerPostMetaBinary sets the handler for the PostMetaBinary method
func (protocol *Protocol) SetHandlerPostMetaBinary(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStorePreparePostParam) (*nex.RMCMessage, *nex.Error)) {
	protocol.PostMetaBinary = handler
}

// SetHandlerTouchObject sets the handler for the TouchObject method
func (protocol *Protocol) SetHandlerTouchObject(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreTouchObjectParam) (*nex.RMCMessage, *nex.Error)) {
	protocol.TouchObject = handler
}

// SetHandlerGetRatingWithLog sets the handler for the GetRatingWithLog method
func (protocol *Protocol) SetHandlerGetRatingWithLog(handler func(err error, packet nex.PacketInterface, callID uint32, target *datastore_types.DataStoreRatingTarget, accessPassword *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetRatingWithLog = handler
}

// SetHandlerPreparePostObject sets the handler for the PreparePostObject method
func (protocol *Protocol) SetHandlerPreparePostObject(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStorePreparePostParam) (*nex.RMCMessage, *nex.Error)) {
	protocol.PreparePostObject = handler
}

// SetHandlerPrepareGetObject sets the handler for the PrepareGetObject method
func (protocol *Protocol) SetHandlerPrepareGetObject(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStorePrepareGetParam) (*nex.RMCMessage, *nex.Error)) {
	protocol.PrepareGetObject = handler
}

// SetHandlerCompletePostObject sets the handler for the CompletePostObject method
func (protocol *Protocol) SetHandlerCompletePostObject(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreCompletePostParam) (*nex.RMCMessage, *nex.Error)) {
	protocol.CompletePostObject = handler
}

// SetHandlerGetNewArrivedNotifications sets the handler for the GetNewArrivedNotifications method
func (protocol *Protocol) SetHandlerGetNewArrivedNotifications(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreGetNewArrivedNotificationsParam) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetNewArrivedNotifications = handler
}

// SetHandlerGetSpecificMeta sets the handler for the GetSpecificMeta method
func (protocol *Protocol) SetHandlerGetSpecificMeta(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreGetSpecificMetaParam) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetSpecificMeta = handler
}

// SetHandlerGetPersistenceInfo sets the handler for the GetPersistenceInfo method
func (protocol *Protocol) SetHandlerGetPersistenceInfo(handler func(err error, packet nex.PacketInterface, callID uint32, ownerID *types.PID, persistenceSlotID *types.PrimitiveU16) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetPersistenceInfo = handler
}

// SetHandlerGetPersistenceInfos sets the handler for the GetPersistenceInfos method
func (protocol *Protocol) SetHandlerGetPersistenceInfos(handler func(err error, packet nex.PacketInterface, callID uint32, ownerID *types.PID, persistenceSlotIDs *types.List[*types.PrimitiveU16]) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetPersistenceInfos = handler
}

// SetHandlerPerpetuateObject sets the handler for the PerpetuateObject method
func (protocol *Protocol) SetHandlerPerpetuateObject(handler func(err error, packet nex.PacketInterface, callID uint32, persistenceSlotID *types.PrimitiveU16, dataID *types.PrimitiveU64, deleteLastObject *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error)) {
	protocol.PerpetuateObject = handler
}

// SetHandlerUnperpetuateObject sets the handler for the UnperpetuateObject method
func (protocol *Protocol) SetHandlerUnperpetuateObject(handler func(err error, packet nex.PacketInterface, callID uint32, persistenceSlotID *types.PrimitiveU16, deleteLastObject *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error)) {
	protocol.UnperpetuateObject = handler
}

// SetHandlerPrepareGetObjectOrMetaBinary sets the handler for the PrepareGetObjectOrMetaBinary method
func (protocol *Protocol) SetHandlerPrepareGetObjectOrMetaBinary(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStorePrepareGetParam) (*nex.RMCMessage, *nex.Error)) {
	protocol.PrepareGetObjectOrMetaBinary = handler
}

// SetHandlerGetPasswordInfo sets the handler for the GetPasswordInfo method
func (protocol *Protocol) SetHandlerGetPasswordInfo(handler func(err error, packet nex.PacketInterface, callID uint32, dataID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetPasswordInfo = handler
}

// SetHandlerGetPasswordInfos sets the handler for the GetPasswordInfos method
func (protocol *Protocol) SetHandlerGetPasswordInfos(handler func(err error, packet nex.PacketInterface, callID uint32, dataIDs *types.List[*types.PrimitiveU64]) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetPasswordInfos = handler
}

// SetHandlerGetMetasMultipleParam sets the handler for the GetMetasMultipleParam method
func (protocol *Protocol) SetHandlerGetMetasMultipleParam(handler func(err error, packet nex.PacketInterface, callID uint32, params *types.List[*datastore_types.DataStoreGetMetaParam]) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetMetasMultipleParam = handler
}

// SetHandlerCompletePostObjects sets the handler for the CompletePostObjects method
func (protocol *Protocol) SetHandlerCompletePostObjects(handler func(err error, packet nex.PacketInterface, callID uint32, dataIDs *types.List[*types.PrimitiveU64]) (*nex.RMCMessage, *nex.Error)) {
	protocol.CompletePostObjects = handler
}

// SetHandlerChangeMeta sets the handler for the ChangeMeta method
func (protocol *Protocol) SetHandlerChangeMeta(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreChangeMetaParam) (*nex.RMCMessage, *nex.Error)) {
	protocol.ChangeMeta = handler
}

// SetHandlerChangeMetas sets the handler for the ChangeMetas method
func (protocol *Protocol) SetHandlerChangeMetas(handler func(err error, packet nex.PacketInterface, callID uint32, dataIDs *types.List[*types.PrimitiveU64], params *types.List[*datastore_types.DataStoreChangeMetaParam], transactional *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error)) {
	protocol.ChangeMetas = handler
}

// SetHandlerRateObjects sets the handler for the RateObjects method
func (protocol *Protocol) SetHandlerRateObjects(handler func(err error, packet nex.PacketInterface, callID uint32, targets *types.List[*datastore_types.DataStoreRatingTarget], params *types.List[*datastore_types.DataStoreRateObjectParam], transactional *types.PrimitiveBool, fetchRatings *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error)) {
	protocol.RateObjects = handler
}

// SetHandlerPostMetaBinaryWithDataID sets the handler for the PostMetaBinaryWithDataID method
func (protocol *Protocol) SetHandlerPostMetaBinaryWithDataID(handler func(err error, packet nex.PacketInterface, callID uint32, dataID *types.PrimitiveU64, param *datastore_types.DataStorePreparePostParam) (*nex.RMCMessage, *nex.Error)) {
	protocol.PostMetaBinaryWithDataID = handler
}

// SetHandlerPostMetaBinariesWithDataID sets the handler for the PostMetaBinariesWithDataID method
func (protocol *Protocol) SetHandlerPostMetaBinariesWithDataID(handler func(err error, packet nex.PacketInterface, callID uint32, dataIDs *types.List[*types.PrimitiveU64], params *types.List[*datastore_types.DataStorePreparePostParam], transactional *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error)) {
	protocol.PostMetaBinariesWithDataID = handler
}

// SetHandlerRateObjectWithPosting sets the handler for the RateObjectWithPosting method
func (protocol *Protocol) SetHandlerRateObjectWithPosting(handler func(err error, packet nex.PacketInterface, callID uint32, target *datastore_types.DataStoreRatingTarget, rateParam *datastore_types.DataStoreRateObjectParam, postParam *datastore_types.DataStorePreparePostParam, fetchRatings *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error)) {
	protocol.RateObjectWithPosting = handler
}

// SetHandlerRateObjectsWithPosting sets the handler for the RateObjectsWithPosting method
func (protocol *Protocol) SetHandlerRateObjectsWithPosting(handler func(err error, packet nex.PacketInterface, callID uint32, targets *types.List[*datastore_types.DataStoreRatingTarget], rateParams *types.List[*datastore_types.DataStoreRateObjectParam], postParams *types.List[*datastore_types.DataStorePreparePostParam], transactional *types.PrimitiveBool, fetchRatings *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error)) {
	protocol.RateObjectsWithPosting = handler
}

// SetHandlerGetObjectInfos sets the handler for the GetObjectInfos method
func (protocol *Protocol) SetHandlerGetObjectInfos(handler func(err error, packet nex.PacketInterface, callID uint32, dataIDs *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetObjectInfos = handler
}

// SetHandlerSearchObjectLight sets the handler for the SearchObjectLight method
func (protocol *Protocol) SetHandlerSearchObjectLight(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreSearchParam) (*nex.RMCMessage, *nex.Error)) {
	protocol.SearchObjectLight = handler
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	message := packet.RMCMessage()

	if !message.IsRequest || message.ProtocolID != ProtocolID {
		return
	}

	if protocol.Patches != nil && slices.Contains(protocol.PatchedMethods, message.MethodID) {
		protocol.Patches.HandlePacket(packet)
		return
	}

	switch message.MethodID {
	case MethodPrepareGetObjectV1:
		protocol.handlePrepareGetObjectV1(packet)
	case MethodPreparePostObjectV1:
		protocol.handlePreparePostObjectV1(packet)
	case MethodCompletePostObjectV1:
		protocol.handleCompletePostObjectV1(packet)
	case MethodDeleteObject:
		protocol.handleDeleteObject(packet)
	case MethodDeleteObjects:
		protocol.handleDeleteObjects(packet)
	case MethodChangeMetaV1:
		protocol.handleChangeMetaV1(packet)
	case MethodChangeMetasV1:
		protocol.handleChangeMetasV1(packet)
	case MethodGetMeta:
		protocol.handleGetMeta(packet)
	case MethodGetMetas:
		protocol.handleGetMetas(packet)
	case MethodPrepareUpdateObject:
		protocol.handlePrepareUpdateObject(packet)
	case MethodCompleteUpdateObject:
		protocol.handleCompleteUpdateObject(packet)
	case MethodSearchObject:
		protocol.handleSearchObject(packet)
	case MethodGetNotificationURL:
		protocol.handleGetNotificationURL(packet)
	case MethodGetNewArrivedNotificationsV1:
		protocol.handleGetNewArrivedNotificationsV1(packet)
	case MethodRateObject:
		protocol.handleRateObject(packet)
	case MethodGetRating:
		protocol.handleGetRating(packet)
	case MethodGetRatings:
		protocol.handleGetRatings(packet)
	case MethodResetRating:
		protocol.handleResetRating(packet)
	case MethodResetRatings:
		protocol.handleResetRatings(packet)
	case MethodGetSpecificMetaV1:
		protocol.handleGetSpecificMetaV1(packet)
	case MethodPostMetaBinary:
		protocol.handlePostMetaBinary(packet)
	case MethodTouchObject:
		protocol.handleTouchObject(packet)
	case MethodGetRatingWithLog:
		protocol.handleGetRatingWithLog(packet)
	case MethodPreparePostObject:
		protocol.handlePreparePostObject(packet)
	case MethodPrepareGetObject:
		protocol.handlePrepareGetObject(packet)
	case MethodCompletePostObject:
		protocol.handleCompletePostObject(packet)
	case MethodGetNewArrivedNotifications:
		protocol.handleGetNewArrivedNotifications(packet)
	case MethodGetSpecificMeta:
		protocol.handleGetSpecificMeta(packet)
	case MethodGetPersistenceInfo:
		protocol.handleGetPersistenceInfo(packet)
	case MethodGetPersistenceInfos:
		protocol.handleGetPersistenceInfos(packet)
	case MethodPerpetuateObject:
		protocol.handlePerpetuateObject(packet)
	case MethodUnperpetuateObject:
		protocol.handleUnperpetuateObject(packet)
	case MethodPrepareGetObjectOrMetaBinary:
		protocol.handlePrepareGetObjectOrMetaBinary(packet)
	case MethodGetPasswordInfo:
		protocol.handleGetPasswordInfo(packet)
	case MethodGetPasswordInfos:
		protocol.handleGetPasswordInfos(packet)
	case MethodGetMetasMultipleParam:
		protocol.handleGetMetasMultipleParam(packet)
	case MethodCompletePostObjects:
		protocol.handleCompletePostObjects(packet)
	case MethodChangeMeta:
		protocol.handleChangeMeta(packet)
	case MethodChangeMetas:
		protocol.handleChangeMetas(packet)
	case MethodRateObjects:
		protocol.handleRateObjects(packet)
	case MethodPostMetaBinaryWithDataID:
		protocol.handlePostMetaBinaryWithDataID(packet)
	case MethodPostMetaBinariesWithDataID:
		protocol.handlePostMetaBinariesWithDataID(packet)
	case MethodRateObjectWithPosting:
		protocol.handleRateObjectWithPosting(packet)
	case MethodRateObjectsWithPosting:
		protocol.handleRateObjectsWithPosting(packet)
	case MethodGetObjectInfos:
		protocol.handleGetObjectInfos(packet)
	case MethodSearchObjectLight:
		protocol.handleSearchObjectLight(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported DataStore method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new DataStore protocol
func NewProtocol() *Protocol {
	return &Protocol{}
}
