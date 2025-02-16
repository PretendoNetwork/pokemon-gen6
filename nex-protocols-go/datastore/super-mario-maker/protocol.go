// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	datastore "github.com/PretendoNetwork/nex-protocols-go/v2/datastore"
	datastore_super_mario_maker_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/super-mario-maker/types"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	"golang.org/x/exp/slices"
)

const (
	// ProtocolID is the Protocol ID for the DataStore (Super Mario Maker) protocol
	ProtocolID = 0x73

	// MethodGetObjectInfos is the method ID for the GetObjectInfos method
	MethodGetObjectInfos = 0x2D

	// MethodGetMetaByOwnerID is the method ID for the GetMetaByOwnerID method
	MethodGetMetaByOwnerID = 0x2E

	// MethodCustomSearchObject is the method ID for the CustomSearchObject method
	MethodCustomSearchObject = 0x2F

	// MethodRateCustomRanking is the method ID for the RateCustomRanking method
	MethodRateCustomRanking = 0x30

	// MethodGetCustomRanking is the method ID for the GetCustomRanking method
	MethodGetCustomRanking = 0x31

	// MethodGetCustomRankingByDataID is the method ID for the GetCustomRankingByDataID method
	MethodGetCustomRankingByDataID = 0x32

	// MethodDeleteCustomRanking is the method ID for the DeleteCustomRanking method
	MethodDeleteCustomRanking = 0x33

	// MethodAddToBufferQueue is the method ID for the AddToBufferQueue method
	MethodAddToBufferQueue = 0x34

	// MethodAddToBufferQueues is the method ID for the AddToBufferQueues method
	MethodAddToBufferQueues = 0x35

	// MethodGetBufferQueue is the method ID for the GetBufferQueue method
	MethodGetBufferQueue = 0x36

	// MethodGetBufferQueues is the method ID for the GetBufferQueues method
	MethodGetBufferQueues = 0x37

	// MethodClearBufferQueues is the method ID for the ClearBufferQueues method
	MethodClearBufferQueues = 0x38

	// MethodCompleteAttachFile is the method ID for the CompleteAttachFile method
	MethodCompleteAttachFile = 0x39

	// MethodCompleteAttachFileV1 is the method ID for the CompleteAttachFileV1 method
	MethodCompleteAttachFileV1 = 0x3A

	// MethodPrepareAttachFile is the method ID for the PrepareAttachFile method
	MethodPrepareAttachFile = 0x3B

	// MethodConditionalSearchObject is the method ID for the ConditionalSearchObject method
	MethodConditionalSearchObject = 0x3C

	// MethodGetApplicationConfig is the method ID for the GetApplicationConfig method
	MethodGetApplicationConfig = 0x3D

	// MethodSetApplicationConfig is the method ID for the SetApplicationConfig method
	MethodSetApplicationConfig = 0x3E

	// MethodDeleteApplicationConfig is the method ID for the DeleteApplicationConfig method
	MethodDeleteApplicationConfig = 0x3F

	// MethodLatestCourseSearchObject is the method ID for the LatestCourseSearchObject method
	MethodLatestCourseSearchObject = 0x40

	// MethodFollowingsLatestCourseSearchObject is the method ID for the FollowingsLatestCourseSearchObject method
	MethodFollowingsLatestCourseSearchObject = 0x41

	// MethodRecommendedCourseSearchObject is the method ID for the RecommendedCourseSearchObject method
	MethodRecommendedCourseSearchObject = 0x42

	// MethodScoreRangeCascadedSearchObject is the method ID for the ScoreRangeCascadedSearchObject method
	MethodScoreRangeCascadedSearchObject = 0x43

	// MethodSuggestedCourseSearchObject is the method ID for the SuggestedCourseSearchObject method
	MethodSuggestedCourseSearchObject = 0x44

	// MethodPreparePostObjectWithOwnerIDAndDataID is the method ID for the PreparePostObjectWithOwnerIDAndDataID method
	MethodPreparePostObjectWithOwnerIDAndDataID = 0x45

	// MethodCompletePostObjectWithOwnerID is the method ID for the CompletePostObjectWithOwnerID method
	MethodCompletePostObjectWithOwnerID = 0x46

	// MethodUploadCourseRecord is the method ID for the UploadCourseRecord method
	MethodUploadCourseRecord = 0x47

	// MethodGetCourseRecord is the method ID for the GetCourseRecord method
	MethodGetCourseRecord = 0x48

	// MethodDeleteCourseRecord is the method ID for the DeleteCourseRecord method
	MethodDeleteCourseRecord = 0x49

	// MethodGetApplicationConfigString is the method ID for the GetApplicationConfigString method
	MethodGetApplicationConfigString = 0x4A

	// MethodSetApplicationConfigString is the method ID for the SetApplicationConfigString method
	MethodSetApplicationConfigString = 0x4B

	// MethodGetDeletionReason is the method ID for the GetDeletionReason method
	MethodGetDeletionReason = 0x4C

	// MethodSetDeletionReason is the method ID for the SetDeletionReason method
	MethodSetDeletionReason = 0x4D

	// MethodGetMetasWithCourseRecord is the method ID for the GetMetasWithCourseRecord method
	MethodGetMetasWithCourseRecord = 0x4E

	// MethodCheckRateCustomRankingCounter is the method ID for the CheckRateCustomRankingCounter method
	MethodCheckRateCustomRankingCounter = 0x4F

	// MethodResetRateCustomRankingCounter is the method ID for the ResetRateCustomRankingCounter method
	MethodResetRateCustomRankingCounter = 0x50

	// MethodBestScoreRateCourseSearchObject is the method ID for the BestScoreRateCourseSearchObject method
	MethodBestScoreRateCourseSearchObject = 0x51

	// MethodCTRPickUpCourseSearchObject is the method ID for the CTRPickUpCourseSearchObject method
	MethodCTRPickUpCourseSearchObject = 0x52

	// MethodSetCachedRanking is the method ID for the SetCachedRanking method
	MethodSetCachedRanking = 0x53

	// MethodDeleteCachedRanking is the method ID for the DeleteCachedRanking method
	MethodDeleteCachedRanking = 0x54

	// MethodChangePlayablePlatform is the method ID for the ChangePlayablePlatform method
	MethodChangePlayablePlatform = 0x55

	// MethodSearchUnknownPlatformObjects is the method ID for the SearchUnknownPlatformObjects method
	MethodSearchUnknownPlatformObjects = 0x56

	// MethodReportCourse is the method ID for the ReportCourse method
	MethodReportCourse = 0x57
)

var patchedMethods = []uint32{
	MethodGetObjectInfos,
	MethodGetMetaByOwnerID,
	MethodCustomSearchObject,
	MethodRateCustomRanking,
	MethodGetCustomRanking,
	MethodGetCustomRankingByDataID,
	MethodDeleteCustomRanking,
	MethodAddToBufferQueue,
	MethodAddToBufferQueues,
	MethodGetBufferQueue,
	MethodGetBufferQueues,
	MethodClearBufferQueues,
	MethodCompleteAttachFile,
	MethodCompleteAttachFileV1,
	MethodPrepareAttachFile,
	MethodConditionalSearchObject,
	MethodGetApplicationConfig,
	MethodSetApplicationConfig,
	MethodDeleteApplicationConfig,
	MethodLatestCourseSearchObject,
	MethodFollowingsLatestCourseSearchObject,
	MethodRecommendedCourseSearchObject,
	MethodScoreRangeCascadedSearchObject,
	MethodSuggestedCourseSearchObject,
	MethodPreparePostObjectWithOwnerIDAndDataID,
	MethodCompletePostObjectWithOwnerID,
	MethodUploadCourseRecord,
	MethodGetCourseRecord,
	MethodDeleteCourseRecord,
	MethodGetApplicationConfigString,
	MethodSetApplicationConfigString,
	MethodGetDeletionReason,
	MethodSetDeletionReason,
	MethodGetMetasWithCourseRecord,
	MethodCheckRateCustomRankingCounter,
	MethodResetRateCustomRankingCounter,
	MethodBestScoreRateCourseSearchObject,
	MethodCTRPickUpCourseSearchObject,
	MethodSetCachedRanking,
	MethodDeleteCachedRanking,
	MethodChangePlayablePlatform,
	MethodSearchUnknownPlatformObjects,
	MethodReportCourse,
}

type dataStoreProtocol = datastore.Protocol

// Protocol stores all the RMC method handlers for the DataStore (Super Mario Maker) protocol and listens for requests
// Embeds the DataStore protocol
type Protocol struct {
	endpoint nex.EndpointInterface
	dataStoreProtocol
	GetObjectInfos                        func(err error, packet nex.PacketInterface, callID uint32, dataIDs *types.List[*types.PrimitiveU64]) (*nex.RMCMessage, *nex.Error)
	GetMetaByOwnerID                      func(err error, packet nex.PacketInterface, callID uint32, param *datastore_super_mario_maker_types.DataStoreGetMetaByOwnerIDParam) (*nex.RMCMessage, *nex.Error)
	CustomSearchObject                    func(err error, packet nex.PacketInterface, callID uint32, condition *types.PrimitiveU32, param *datastore_types.DataStoreSearchParam) (*nex.RMCMessage, *nex.Error)
	RateCustomRanking                     func(err error, packet nex.PacketInterface, callID uint32, params *types.List[*datastore_super_mario_maker_types.DataStoreRateCustomRankingParam]) (*nex.RMCMessage, *nex.Error)
	GetCustomRanking                      func(err error, packet nex.PacketInterface, callID uint32, param *datastore_super_mario_maker_types.DataStoreGetCustomRankingParam) (*nex.RMCMessage, *nex.Error)
	GetCustomRankingByDataID              func(err error, packet nex.PacketInterface, callID uint32, param *datastore_super_mario_maker_types.DataStoreGetCustomRankingByDataIDParam) (*nex.RMCMessage, *nex.Error)
	DeleteCustomRanking                   func(err error, packet nex.PacketInterface, callID uint32, dataIDList *types.List[*types.PrimitiveU64]) (*nex.RMCMessage, *nex.Error)
	AddToBufferQueue                      func(err error, packet nex.PacketInterface, callID uint32, param *datastore_super_mario_maker_types.BufferQueueParam, buffer *types.QBuffer) (*nex.RMCMessage, *nex.Error)
	AddToBufferQueues                     func(err error, packet nex.PacketInterface, callID uint32, params *types.List[*datastore_super_mario_maker_types.BufferQueueParam], buffers *types.List[*types.QBuffer]) (*nex.RMCMessage, *nex.Error)
	GetBufferQueue                        func(err error, packet nex.PacketInterface, callID uint32, param *datastore_super_mario_maker_types.BufferQueueParam) (*nex.RMCMessage, *nex.Error)
	GetBufferQueues                       func(err error, packet nex.PacketInterface, callID uint32, params *types.List[*datastore_super_mario_maker_types.BufferQueueParam]) (*nex.RMCMessage, *nex.Error)
	ClearBufferQueues                     func(err error, packet nex.PacketInterface, callID uint32, params *types.List[*datastore_super_mario_maker_types.BufferQueueParam]) (*nex.RMCMessage, *nex.Error)
	CompleteAttachFile                    func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreCompletePostParam) (*nex.RMCMessage, *nex.Error)
	CompleteAttachFileV1                  func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreCompletePostParamV1) (*nex.RMCMessage, *nex.Error)
	PrepareAttachFile                     func(err error, packet nex.PacketInterface, callID uint32, param *datastore_super_mario_maker_types.DataStoreAttachFileParam) (*nex.RMCMessage, *nex.Error)
	ConditionalSearchObject               func(err error, packet nex.PacketInterface, callID uint32, condition *types.PrimitiveU32, param *datastore_types.DataStoreSearchParam, extraData *types.List[*types.String]) (*nex.RMCMessage, *nex.Error)
	GetApplicationConfig                  func(err error, packet nex.PacketInterface, callID uint32, applicationID *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	SetApplicationConfig                  func(err error, packet nex.PacketInterface, callID uint32, applicationID *types.PrimitiveU32, key *types.PrimitiveU32, value *types.PrimitiveS32) (*nex.RMCMessage, *nex.Error)
	DeleteApplicationConfig               func(err error, packet nex.PacketInterface, callID uint32, applicationID *types.PrimitiveU32, key *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	LatestCourseSearchObject              func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreSearchParam, extraData *types.List[*types.String]) (*nex.RMCMessage, *nex.Error)
	FollowingsLatestCourseSearchObject    func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreSearchParam, extraData *types.List[*types.String]) (*nex.RMCMessage, *nex.Error)
	RecommendedCourseSearchObject         func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreSearchParam, extraData *types.List[*types.String]) (*nex.RMCMessage, *nex.Error)
	ScoreRangeCascadedSearchObject        func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreSearchParam, extraData *types.List[*types.String]) (*nex.RMCMessage, *nex.Error)
	SuggestedCourseSearchObject           func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreSearchParam, extraData *types.List[*types.String]) (*nex.RMCMessage, *nex.Error)
	PreparePostObjectWithOwnerIDAndDataID func(err error, packet nex.PacketInterface, callID uint32, ownerID *types.PrimitiveU32, dataID *types.PrimitiveU64, param *datastore_types.DataStorePreparePostParam) (*nex.RMCMessage, *nex.Error)
	CompletePostObjectWithOwnerID         func(err error, packet nex.PacketInterface, callID uint32, ownerID *types.PrimitiveU32, param *datastore_types.DataStoreCompletePostParam) (*nex.RMCMessage, *nex.Error)
	UploadCourseRecord                    func(err error, packet nex.PacketInterface, callID uint32, param *datastore_super_mario_maker_types.DataStoreUploadCourseRecordParam) (*nex.RMCMessage, *nex.Error)
	GetCourseRecord                       func(err error, packet nex.PacketInterface, callID uint32, param *datastore_super_mario_maker_types.DataStoreGetCourseRecordParam) (*nex.RMCMessage, *nex.Error)
	DeleteCourseRecord                    func(err error, packet nex.PacketInterface, callID uint32, param *datastore_super_mario_maker_types.DataStoreGetCourseRecordParam) (*nex.RMCMessage, *nex.Error)
	GetApplicationConfigString            func(err error, packet nex.PacketInterface, callID uint32, applicationID *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	SetApplicationConfigString            func(err error, packet nex.PacketInterface, callID uint32, applicationID *types.PrimitiveU32, key *types.PrimitiveU32, value *types.String) (*nex.RMCMessage, *nex.Error)
	GetDeletionReason                     func(err error, packet nex.PacketInterface, callID uint32, dataIDLst *types.List[*types.PrimitiveU64]) (*nex.RMCMessage, *nex.Error)
	SetDeletionReason                     func(err error, packet nex.PacketInterface, callID uint32, dataIDLst *types.List[*types.PrimitiveU64], deletionReason *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	GetMetasWithCourseRecord              func(err error, packet nex.PacketInterface, callID uint32, params *types.List[*datastore_super_mario_maker_types.DataStoreGetCourseRecordParam], metaParam *datastore_types.DataStoreGetMetaParam) (*nex.RMCMessage, *nex.Error)
	CheckRateCustomRankingCounter         func(err error, packet nex.PacketInterface, callID uint32, applicationID *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	ResetRateCustomRankingCounter         func(err error, packet nex.PacketInterface, callID uint32, applicationID *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	BestScoreRateCourseSearchObject       func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreSearchParam, extraData *types.List[*types.String]) (*nex.RMCMessage, *nex.Error)
	CTRPickUpCourseSearchObject           func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreSearchParam, extraData *types.List[*types.String]) (*nex.RMCMessage, *nex.Error)
	SetCachedRanking                      func(err error, packet nex.PacketInterface, callID uint32, rankingType *types.String, rankingArgs *types.List[*types.String], dataIDLst *types.List[*types.PrimitiveU64]) (*nex.RMCMessage, *nex.Error)
	DeleteCachedRanking                   func(err error, packet nex.PacketInterface, callID uint32, rankingType *types.String, rankingArgs *types.List[*types.String]) (*nex.RMCMessage, *nex.Error)
	ChangePlayablePlatform                func(err error, packet nex.PacketInterface, callID uint32, params *types.List[*datastore_super_mario_maker_types.DataStoreChangePlayablePlatformParam]) (*nex.RMCMessage, *nex.Error)
	SearchUnknownPlatformObjects          func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error)
	ReportCourse                          func(err error, packet nex.PacketInterface, callID uint32, param *datastore_super_mario_maker_types.DataStoreReportCourseParam) (*nex.RMCMessage, *nex.Error)
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	message := packet.RMCMessage()

	if !message.IsRequest || message.ProtocolID != ProtocolID {
		return
	}

	if !slices.Contains(patchedMethods, message.MethodID) {
		protocol.dataStoreProtocol.HandlePacket(packet)
		return
	}

	switch message.MethodID {
	case MethodGetObjectInfos:
		protocol.handleGetObjectInfos(packet)
	case MethodGetMetaByOwnerID:
		protocol.handleGetMetaByOwnerID(packet)
	case MethodCustomSearchObject:
		protocol.handleCustomSearchObject(packet)
	case MethodRateCustomRanking:
		protocol.handleRateCustomRanking(packet)
	case MethodGetCustomRanking:
		protocol.handleGetCustomRanking(packet)
	case MethodGetCustomRankingByDataID:
		protocol.handleGetCustomRankingByDataID(packet)
	case MethodDeleteCustomRanking:
		protocol.handleDeleteCustomRanking(packet)
	case MethodAddToBufferQueue:
		protocol.handleAddToBufferQueue(packet)
	case MethodAddToBufferQueues:
		protocol.handleAddToBufferQueues(packet)
	case MethodGetBufferQueue:
		protocol.handleGetBufferQueue(packet)
	case MethodGetBufferQueues:
		protocol.handleGetBufferQueues(packet)
	case MethodClearBufferQueues:
		protocol.handleClearBufferQueues(packet)
	case MethodCompleteAttachFile:
		protocol.handleCompleteAttachFile(packet)
	case MethodCompleteAttachFileV1:
		protocol.handleCompleteAttachFileV1(packet)
	case MethodPrepareAttachFile:
		protocol.handlePrepareAttachFile(packet)
	case MethodConditionalSearchObject:
		protocol.handleConditionalSearchObject(packet)
	case MethodGetApplicationConfig:
		protocol.handleGetApplicationConfig(packet)
	case MethodSetApplicationConfig:
		protocol.handleSetApplicationConfig(packet)
	case MethodDeleteApplicationConfig:
		protocol.handleDeleteApplicationConfig(packet)
	case MethodLatestCourseSearchObject:
		protocol.handleLatestCourseSearchObject(packet)
	case MethodFollowingsLatestCourseSearchObject:
		protocol.handleFollowingsLatestCourseSearchObject(packet)
	case MethodRecommendedCourseSearchObject:
		protocol.handleRecommendedCourseSearchObject(packet)
	case MethodScoreRangeCascadedSearchObject:
		protocol.handleScoreRangeCascadedSearchObject(packet)
	case MethodSuggestedCourseSearchObject:
		protocol.handleSuggestedCourseSearchObject(packet)
	case MethodPreparePostObjectWithOwnerIDAndDataID:
		protocol.handlePreparePostObjectWithOwnerIDAndDataID(packet)
	case MethodCompletePostObjectWithOwnerID:
		protocol.handleCompletePostObjectWithOwnerID(packet)
	case MethodUploadCourseRecord:
		protocol.handleUploadCourseRecord(packet)
	case MethodGetCourseRecord:
		protocol.handleGetCourseRecord(packet)
	case MethodDeleteCourseRecord:
		protocol.handleDeleteCourseRecord(packet)
	case MethodGetApplicationConfigString:
		protocol.handleGetApplicationConfigString(packet)
	case MethodSetApplicationConfigString:
		protocol.handleSetApplicationConfigString(packet)
	case MethodGetDeletionReason:
		protocol.handleGetDeletionReason(packet)
	case MethodSetDeletionReason:
		protocol.handleSetDeletionReason(packet)
	case MethodGetMetasWithCourseRecord:
		protocol.handleGetMetasWithCourseRecord(packet)
	case MethodCheckRateCustomRankingCounter:
		protocol.handleCheckRateCustomRankingCounter(packet)
	case MethodResetRateCustomRankingCounter:
		protocol.handleResetRateCustomRankingCounter(packet)
	case MethodBestScoreRateCourseSearchObject:
		protocol.handleBestScoreRateCourseSearchObject(packet)
	case MethodCTRPickUpCourseSearchObject:
		protocol.handleCTRPickUpCourseSearchObject(packet)
	case MethodSetCachedRanking:
		protocol.handleSetCachedRanking(packet)
	case MethodDeleteCachedRanking:
		protocol.handleDeleteCachedRanking(packet)
	case MethodChangePlayablePlatform:
		protocol.handleChangePlayablePlatform(packet)
	case MethodSearchUnknownPlatformObjects:
		protocol.handleSearchUnknownPlatformObjects(packet)
	case MethodReportCourse:
		protocol.handleReportCourse(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported DataStore (Super Mario Maker) method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new DataStoreSuperMarioMaker protocol
func NewProtocol(endpoint nex.EndpointInterface) *Protocol {
	protocol := &Protocol{endpoint: endpoint}
	protocol.dataStoreProtocol.SetEndpoint(endpoint)

	return protocol
}
