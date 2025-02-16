// Package protocol implements the Super Smash Bros. 4 DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	datastore "github.com/PretendoNetwork/nex-protocols-go/v2/datastore"
	datastore_super_smash_bros_4_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/super-smash-bros-4/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	"golang.org/x/exp/slices"
)

const (
	// ProtocolID is the protocol ID for the DataStore (Super Smash Bros 4) protocol. ID is the same as the DataStore protocol
	ProtocolID = 0x73

	// MethodPostProfile is the method ID for the method PostProfile
	MethodPostProfile = 0x2D

	// MethodGetProfiles is the method ID for the method GetProfiles
	MethodGetProfiles = 0x2E

	// MethodSendPlayReport is the method ID for the method SendPlayReport
	MethodSendPlayReport = 0x2F

	// MethodGetWorldPlayReport is the method ID for the method GetWorldPlayReport
	MethodGetWorldPlayReport = 0x30

	// MethodGetReplayMeta is the method ID for the method GetReplayMeta
	MethodGetReplayMeta = 0x31

	// MethodPrepareGetReplay is the method ID for the method PrepareGetReplay
	MethodPrepareGetReplay = 0x32

	// MethodPreparePostReplay is the method ID for the method PreparePostReplay
	MethodPreparePostReplay = 0x33

	// MethodCompletePostReplay is the method ID for the method CompletePostReplay
	MethodCompletePostReplay = 0x34

	// MethodCheckPostReplay is the method ID for the method CheckPostReplay
	MethodCheckPostReplay = 0x35

	// MethodGetNextReplay is the method ID for the method GetNextReplay
	MethodGetNextReplay = 0x36

	// MethodPreparePostSharedData is the method ID for the method PreparePostSharedData
	MethodPreparePostSharedData = 0x37

	// MethodCompletePostSharedData is the method ID for the method CompletePostSharedData
	MethodCompletePostSharedData = 0x38

	// MethodSearchSharedData is the method ID for the method SearchSharedData
	MethodSearchSharedData = 0x39

	// MethodGetApplicationConfig is the method ID for the method GetApplicationConfig
	MethodGetApplicationConfig = 0x3A

	// MethodSearchReplay is the method ID for the method SearchReplay
	MethodSearchReplay = 0x3B

	// MethodPostFightingPowerScore is the method ID for the method PostFightingPowerScore
	MethodPostFightingPowerScore = 0x3C

	// MethodGetFightingPowerChart is the method ID for the method GetFightingPowerChart
	MethodGetFightingPowerChart = 0x3D

	// MethodGetFightingPowerChartAll is the method ID for the method GetFightingPowerChartAll
	MethodGetFightingPowerChartAll = 0x3E

	// MethodReportSharedData is the method ID for the method ReportSharedData
	MethodReportSharedData = 0x3F

	// MethodGetSharedDataMeta is the method ID for the method GetSharedDataMeta
	MethodGetSharedDataMeta = 0x40
)

var patchedMethods = []uint32{
	MethodPostProfile,
	MethodGetProfiles,
	MethodSendPlayReport,
	MethodGetWorldPlayReport,
	MethodGetReplayMeta,
	MethodPrepareGetReplay,
	MethodPreparePostReplay,
	MethodCompletePostReplay,
	MethodCheckPostReplay,
	MethodGetNextReplay,
	MethodPreparePostSharedData,
	MethodCompletePostSharedData,
	MethodSearchSharedData,
	MethodGetApplicationConfig,
	MethodSearchReplay,
	MethodPostFightingPowerScore,
	MethodGetFightingPowerChart,
	MethodGetFightingPowerChartAll,
	MethodReportSharedData,
}

type dataStoreProtocol = datastore.Protocol

// Protocol stores all the RMC method handlers for the DataStore (Super Smash Bros 4) protocol and listens for requests
// Embeds the DataStore protocol
type Protocol struct {
	endpoint nex.EndpointInterface
	dataStoreProtocol
	PostProfile              func(err error, packet nex.PacketInterface, callID uint32, param *datastore_super_smash_bros_4_types.DataStorePostProfileParam) (*nex.RMCMessage, *nex.Error)
	GetProfiles              func(err error, packet nex.PacketInterface, callID uint32, pidList *types.List[*types.PID]) (*nex.RMCMessage, *nex.Error)
	SendPlayReport           func(err error, packet nex.PacketInterface, callID uint32, playReport *types.List[*types.PrimitiveS32]) (*nex.RMCMessage, *nex.Error)
	GetWorldPlayReport       func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	GetReplayMeta            func(err error, packet nex.PacketInterface, callID uint32, param *datastore_super_smash_bros_4_types.DataStoreGetReplayMetaParam) (*nex.RMCMessage, *nex.Error)
	PrepareGetReplay         func(err error, packet nex.PacketInterface, callID uint32, param *datastore_super_smash_bros_4_types.DataStorePrepareGetReplayParam) (*nex.RMCMessage, *nex.Error)
	PreparePostReplay        func(err error, packet nex.PacketInterface, callID uint32, param *datastore_super_smash_bros_4_types.DataStorePreparePostReplayParam) (*nex.RMCMessage, *nex.Error)
	CompletePostReplay       func(err error, packet nex.PacketInterface, callID uint32, param *datastore_super_smash_bros_4_types.DataStoreCompletePostReplayParam) (*nex.RMCMessage, *nex.Error)
	CheckPostReplay          func(err error, packet nex.PacketInterface, callID uint32, param *datastore_super_smash_bros_4_types.DataStorePreparePostReplayParam) (*nex.RMCMessage, *nex.Error)
	GetNextReplay            func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	PreparePostSharedData    func(err error, packet nex.PacketInterface, callID uint32, param *datastore_super_smash_bros_4_types.DataStorePreparePostSharedDataParam) (*nex.RMCMessage, *nex.Error)
	CompletePostSharedData   func(err error, packet nex.PacketInterface, callID uint32, param *datastore_super_smash_bros_4_types.DataStoreCompletePostSharedDataParam) (*nex.RMCMessage, *nex.Error)
	SearchSharedData         func(err error, packet nex.PacketInterface, callID uint32, param *datastore_super_smash_bros_4_types.DataStoreSearchSharedDataParam) (*nex.RMCMessage, *nex.Error)
	GetApplicationConfig     func(err error, packet nex.PacketInterface, callID uint32, applicationID *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	SearchReplay             func(err error, packet nex.PacketInterface, callID uint32, param *datastore_super_smash_bros_4_types.DataStoreSearchReplayParam) (*nex.RMCMessage, *nex.Error)
	PostFightingPowerScore   func(err error, packet nex.PacketInterface, callID uint32, params *types.List[*datastore_super_smash_bros_4_types.DataStorePostFightingPowerScoreParam]) (*nex.RMCMessage, *nex.Error)
	GetFightingPowerChart    func(err error, packet nex.PacketInterface, callID uint32, mode *types.PrimitiveU8) (*nex.RMCMessage, *nex.Error)
	GetFightingPowerChartAll func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	ReportSharedData         func(err error, packet nex.PacketInterface, callID uint32, dataID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)
	GetSharedDataMeta        func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error)
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
	case MethodPostProfile:
		protocol.handlePostProfile(packet)
	case MethodGetProfiles:
		protocol.handleGetProfiles(packet)
	case MethodSendPlayReport:
		protocol.handleSendPlayReport(packet)
	case MethodGetWorldPlayReport:
		protocol.handleGetWorldPlayReport(packet)
	case MethodGetReplayMeta:
		protocol.handleGetReplayMeta(packet)
	case MethodPrepareGetReplay:
		protocol.handlePrepareGetReplay(packet)
	case MethodPreparePostReplay:
		protocol.handlePreparePostReplay(packet)
	case MethodCompletePostReplay:
		protocol.handleCompletePostReplay(packet)
	case MethodCheckPostReplay:
		protocol.handleCheckPostReplay(packet)
	case MethodGetNextReplay:
		protocol.handleGetNextReplay(packet)
	case MethodPreparePostSharedData:
		protocol.handlePreparePostSharedData(packet)
	case MethodCompletePostSharedData:
		protocol.handleCompletePostSharedData(packet)
	case MethodSearchSharedData:
		protocol.handleSearchSharedData(packet)
	case MethodGetApplicationConfig:
		protocol.handleGetApplicationConfig(packet)
	case MethodSearchReplay:
		protocol.handleSearchReplay(packet)
	case MethodPostFightingPowerScore:
		protocol.handlePostFightingPowerScore(packet)
	case MethodGetFightingPowerChart:
		protocol.handleGetFightingPowerChart(packet)
	case MethodGetFightingPowerChartAll:
		protocol.handleGetFightingPowerChartAll(packet)
	case MethodReportSharedData:
		protocol.handleReportSharedData(packet)
	case MethodGetSharedDataMeta:
		protocol.handleGetSharedDataMeta(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported DataStoreSuperSmashBros4 method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new DataStore (Super Smash Bros 4) protocol
func NewProtocol(endpoint nex.EndpointInterface) *Protocol {
	protocol := &Protocol{endpoint: endpoint}
	protocol.dataStoreProtocol.SetEndpoint(endpoint)

	return protocol
}
