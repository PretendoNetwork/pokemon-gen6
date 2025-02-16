package datastore

import (
	"context"
	"slices"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	datastore "github.com/PretendoNetwork/nex-protocols-go/v2/datastore"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/types"
	"github.com/minio/minio-go/v7"
)

type CommonProtocol struct {
	endpoint                                     nex.EndpointInterface
	protocol                                     datastore.Interface
	S3Bucket                                     string
	s3DataKeyBase                                string
	s3NotifyKeyBase                              string
	RootCACert                                   []byte
	minIOClient                                  *minio.Client
	S3Presigner                                  S3PresignerInterface
	GetUserFriendPIDs                            func(pid uint32) []uint32
	GetObjectInfoByDataID                        func(dataID *types.PrimitiveU64) (*datastore_types.DataStoreMetaInfo, *nex.Error)
	UpdateObjectPeriodByDataIDWithPassword       func(dataID *types.PrimitiveU64, dataType *types.PrimitiveU16, password *types.PrimitiveU64) *nex.Error
	UpdateObjectMetaBinaryByDataIDWithPassword   func(dataID *types.PrimitiveU64, metaBinary *types.QBuffer, password *types.PrimitiveU64) *nex.Error
	UpdateObjectDataTypeByDataIDWithPassword     func(dataID *types.PrimitiveU64, period *types.PrimitiveU16, password *types.PrimitiveU64) *nex.Error
	GetObjectSizeByDataID                        func(dataID *types.PrimitiveU64) (uint32, *nex.Error)
	UpdateObjectUploadCompletedByDataID          func(dataID *types.PrimitiveU64, uploadCompleted bool) *nex.Error
	GetObjectInfoByPersistenceTargetWithPassword func(persistenceTarget *datastore_types.DataStorePersistenceTarget, password *types.PrimitiveU64) (*datastore_types.DataStoreMetaInfo, *nex.Error)
	GetObjectInfoByDataIDWithPassword            func(dataID *types.PrimitiveU64, password *types.PrimitiveU64) (*datastore_types.DataStoreMetaInfo, *nex.Error)
	S3GetRequestHeaders                          func() ([]*datastore_types.DataStoreKeyValue, *nex.Error)
	S3PostRequestHeaders                         func() ([]*datastore_types.DataStoreKeyValue, *nex.Error)
	InitializeObjectByPreparePostParam           func(ownerPID *types.PID, param *datastore_types.DataStorePreparePostParam) (uint64, *nex.Error)
	InitializeObjectRatingWithSlot               func(dataID uint64, param *datastore_types.DataStoreRatingInitParamWithSlot) *nex.Error
	RateObjectWithPassword                       func(dataID *types.PrimitiveU64, slot *types.PrimitiveU8, ratingValue *types.PrimitiveS32, accessPassword *types.PrimitiveU64) (*datastore_types.DataStoreRatingInfo, *nex.Error)
	DeleteObjectByDataIDWithPassword             func(dataID *types.PrimitiveU64, password *types.PrimitiveU64) *nex.Error
	DeleteObjectByDataID                         func(dataID *types.PrimitiveU64) *nex.Error
	GetObjectInfosByDataStoreSearchParam         func(param *datastore_types.DataStoreSearchParam) ([]*datastore_types.DataStoreMetaInfo, uint32, *nex.Error)
	GetObjectOwnerByDataID                       func(dataID *types.PrimitiveU64) (uint32, *nex.Error)
	OnAfterDeleteObject                          func(packet nex.PacketInterface, param *datastore_types.DataStoreDeleteParam)
	OnAfterGetMeta                               func(packet nex.PacketInterface, param *datastore_types.DataStoreGetMetaParam)
	OnAfterGetMetas                              func(packet nex.PacketInterface, dataIDs *types.List[*types.PrimitiveU64], param *datastore_types.DataStoreGetMetaParam)
	OnAfterSearchObject                          func(packet nex.PacketInterface, param *datastore_types.DataStoreSearchParam)
	OnAfterRateObject                            func(packet nex.PacketInterface, target *datastore_types.DataStoreRatingTarget, param *datastore_types.DataStoreRateObjectParam, fetchRatings *types.PrimitiveBool)
	OnAfterPostMetaBinary                        func(packet nex.PacketInterface, param *datastore_types.DataStorePreparePostParam)
	OnAfterPreparePostObject                     func(packet nex.PacketInterface, param *datastore_types.DataStorePreparePostParam)
	OnAfterPrepareGetObject                      func(packet nex.PacketInterface, param *datastore_types.DataStorePrepareGetParam)
	OnAfterCompletePostObject                    func(packet nex.PacketInterface, param *datastore_types.DataStoreCompletePostParam)
	OnAfterGetMetasMultipleParam                 func(packet nex.PacketInterface, params *types.List[*datastore_types.DataStoreGetMetaParam])
	OnAfterCompletePostObjects                   func(packet nex.PacketInterface, dataIDs *types.List[*types.PrimitiveU64])
	OnAfterChangeMeta                            func(packet nex.PacketInterface, param *datastore_types.DataStoreChangeMetaParam)
	OnAfterRateObjects                           func(packet nex.PacketInterface, targets *types.List[*datastore_types.DataStoreRatingTarget], params *types.List[*datastore_types.DataStoreRateObjectParam], transactional *types.PrimitiveBool, fetchRatings *types.PrimitiveBool)
}

func (c *CommonProtocol) S3StatObject(bucket, key string) (minio.ObjectInfo, error) {
	return c.minIOClient.StatObject(context.TODO(), bucket, key, minio.StatObjectOptions{})
}

func (c *CommonProtocol) S3ObjectSize(bucket, key string) (uint64, error) {
	info, err := c.S3StatObject(bucket, key)
	if err != nil {
		return 0, err
	}

	return uint64(info.Size), nil
}

func (c *CommonProtocol) VerifyObjectPermission(ownerPID, accessorPID *types.PID, permission *datastore_types.DataStorePermission) *nex.Error {
	if permission.Permission.Value > 3 {
		return nex.NewError(nex.ResultCodes.DataStore.InvalidArgument, "change_error")
	}

	// * Owner can always access their own objects
	if ownerPID.Equals(accessorPID) {
		return nil
	}

	// * Allow anyone
	if permission.Permission.Value == 0 {
		return nil
	}

	// * Allow only friends of the owner
	if permission.Permission.Value == 1 {
		// TODO - This assumes a legacy client. Will not work on the Switch
		friendsList := c.GetUserFriendPIDs(ownerPID.LegacyValue())

		if !slices.Contains(friendsList, accessorPID.LegacyValue()) {
			return nex.NewError(nex.ResultCodes.DataStore.PermissionDenied, "change_error")
		}
	}

	// * Allow only users whose PIDs are defined in permission.RecipientIDs
	if permission.Permission.Value == 2 {
		if !permission.RecipientIDs.Contains(accessorPID) {
			return nex.NewError(nex.ResultCodes.DataStore.PermissionDenied, "change_error")
		}
	}

	// * Allow only the owner
	if permission.Permission.Value == 3 {
		if !ownerPID.Equals(accessorPID) {
			return nex.NewError(nex.ResultCodes.DataStore.PermissionDenied, "change_error")
		}
	}

	return nil
}

// SetDataKeyBase sets the base for the key to be used when uploading standard DataStore objects
func (c *CommonProtocol) SetDataKeyBase(base string) {
	// * Just in case someone passes a badly formatted key
	base = strings.TrimPrefix(base, "/")
	base = strings.TrimSuffix(base, "/")
	c.s3DataKeyBase = base
}

// SetNotifyKeyBase sets the base for the key to be used when uploading DataStore notification data
func (c *CommonProtocol) SetNotifyKeyBase(base string) {
	// * Just in case someone passes a badly formatted key
	base = strings.TrimPrefix(base, "/")
	base = strings.TrimSuffix(base, "/")
	c.s3NotifyKeyBase = base
}

// SetMinIOClient sets the MinIO S3 client
func (c *CommonProtocol) SetMinIOClient(client *minio.Client) {
	c.minIOClient = client
	c.S3Presigner = NewS3Presigner(c.minIOClient)
}

// NewCommonProtocol returns a new CommonProtocol
func NewCommonProtocol(protocol datastore.Interface) *CommonProtocol {
	commonProtocol := &CommonProtocol{
		endpoint:   protocol.Endpoint(),
		protocol:   protocol,
		RootCACert: []byte{},
		S3GetRequestHeaders: func() ([]*datastore_types.DataStoreKeyValue, *nex.Error) {
			return []*datastore_types.DataStoreKeyValue{}, nil
		},
		S3PostRequestHeaders: func() ([]*datastore_types.DataStoreKeyValue, *nex.Error) {
			return []*datastore_types.DataStoreKeyValue{}, nil
		},
	}

	protocol.SetHandlerDeleteObject(commonProtocol.deleteObject)
	protocol.SetHandlerGetMeta(commonProtocol.getMeta)
	protocol.SetHandlerGetMetas(commonProtocol.getMetas)
	protocol.SetHandlerSearchObject(commonProtocol.searchObject)
	protocol.SetHandlerRateObject(commonProtocol.rateObject)
	protocol.SetHandlerPostMetaBinary(commonProtocol.postMetaBinary)
	protocol.SetHandlerPreparePostObject(commonProtocol.preparePostObject)
	protocol.SetHandlerPrepareGetObject(commonProtocol.prepareGetObject)
	protocol.SetHandlerCompletePostObject(commonProtocol.completePostObject)
	protocol.SetHandlerGetMetasMultipleParam(commonProtocol.getMetasMultipleParam)
	protocol.SetHandlerCompletePostObjects(commonProtocol.completePostObjects)
	protocol.SetHandlerChangeMeta(commonProtocol.changeMeta)
	protocol.SetHandlerRateObjects(commonProtocol.rateObjects)

	return commonProtocol
}
