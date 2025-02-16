// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreMetaInfo is a type within the DataStore protocol
type DataStoreMetaInfo struct {
	types.Structure
	DataID        *types.PrimitiveU64
	OwnerID       *types.PID
	Size          *types.PrimitiveU32
	DataType      *types.PrimitiveU16
	Name          *types.String
	MetaBinary    *types.QBuffer
	Permission    *DataStorePermission
	DelPermission *DataStorePermission
	CreatedTime   *types.DateTime
	UpdatedTime   *types.DateTime
	Period        *types.PrimitiveU16
	Status        *types.PrimitiveU8
	ReferredCnt   *types.PrimitiveU32
	ReferDataID   *types.PrimitiveU32
	Flag          *types.PrimitiveU32
	ReferredTime  *types.DateTime
	ExpireTime    *types.DateTime
	Tags          *types.List[*types.String]
	Ratings       *types.List[*DataStoreRatingInfoWithSlot]
}

// WriteTo writes the DataStoreMetaInfo to the given writable
func (dsmi *DataStoreMetaInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dsmi.DataID.WriteTo(contentWritable)
	dsmi.OwnerID.WriteTo(contentWritable)
	dsmi.Size.WriteTo(contentWritable)
	dsmi.DataType.WriteTo(contentWritable)
	dsmi.Name.WriteTo(contentWritable)
	dsmi.MetaBinary.WriteTo(contentWritable)
	dsmi.Permission.WriteTo(contentWritable)
	dsmi.DelPermission.WriteTo(contentWritable)
	dsmi.CreatedTime.WriteTo(contentWritable)
	dsmi.UpdatedTime.WriteTo(contentWritable)
	dsmi.Period.WriteTo(contentWritable)
	dsmi.Status.WriteTo(contentWritable)
	dsmi.ReferredCnt.WriteTo(contentWritable)
	dsmi.ReferDataID.WriteTo(contentWritable)
	dsmi.Flag.WriteTo(contentWritable)
	dsmi.ReferredTime.WriteTo(contentWritable)
	dsmi.ExpireTime.WriteTo(contentWritable)
	dsmi.Tags.WriteTo(contentWritable)
	dsmi.Ratings.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dsmi.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreMetaInfo from the given readable
func (dsmi *DataStoreMetaInfo) ExtractFrom(readable types.Readable) error {
	var err error

	err = dsmi.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo header. %s", err.Error())
	}

	err = dsmi.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo.DataID. %s", err.Error())
	}

	err = dsmi.OwnerID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo.OwnerID. %s", err.Error())
	}

	err = dsmi.Size.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo.Size. %s", err.Error())
	}

	err = dsmi.DataType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo.DataType. %s", err.Error())
	}

	err = dsmi.Name.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo.Name. %s", err.Error())
	}

	err = dsmi.MetaBinary.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo.MetaBinary. %s", err.Error())
	}

	err = dsmi.Permission.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo.Permission. %s", err.Error())
	}

	err = dsmi.DelPermission.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo.DelPermission. %s", err.Error())
	}

	err = dsmi.CreatedTime.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo.CreatedTime. %s", err.Error())
	}

	err = dsmi.UpdatedTime.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo.UpdatedTime. %s", err.Error())
	}

	err = dsmi.Period.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo.Period. %s", err.Error())
	}

	err = dsmi.Status.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo.Status. %s", err.Error())
	}

	err = dsmi.ReferredCnt.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo.ReferredCnt. %s", err.Error())
	}

	err = dsmi.ReferDataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo.ReferDataID. %s", err.Error())
	}

	err = dsmi.Flag.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo.Flag. %s", err.Error())
	}

	err = dsmi.ReferredTime.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo.ReferredTime. %s", err.Error())
	}

	err = dsmi.ExpireTime.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo.ExpireTime. %s", err.Error())
	}

	err = dsmi.Tags.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo.Tags. %s", err.Error())
	}

	err = dsmi.Ratings.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo.Ratings. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreMetaInfo
func (dsmi *DataStoreMetaInfo) Copy() types.RVType {
	copied := NewDataStoreMetaInfo()

	copied.StructureVersion = dsmi.StructureVersion
	copied.DataID = dsmi.DataID.Copy().(*types.PrimitiveU64)
	copied.OwnerID = dsmi.OwnerID.Copy().(*types.PID)
	copied.Size = dsmi.Size.Copy().(*types.PrimitiveU32)
	copied.DataType = dsmi.DataType.Copy().(*types.PrimitiveU16)
	copied.Name = dsmi.Name.Copy().(*types.String)
	copied.MetaBinary = dsmi.MetaBinary.Copy().(*types.QBuffer)
	copied.Permission = dsmi.Permission.Copy().(*DataStorePermission)
	copied.DelPermission = dsmi.DelPermission.Copy().(*DataStorePermission)
	copied.CreatedTime = dsmi.CreatedTime.Copy().(*types.DateTime)
	copied.UpdatedTime = dsmi.UpdatedTime.Copy().(*types.DateTime)
	copied.Period = dsmi.Period.Copy().(*types.PrimitiveU16)
	copied.Status = dsmi.Status.Copy().(*types.PrimitiveU8)
	copied.ReferredCnt = dsmi.ReferredCnt.Copy().(*types.PrimitiveU32)
	copied.ReferDataID = dsmi.ReferDataID.Copy().(*types.PrimitiveU32)
	copied.Flag = dsmi.Flag.Copy().(*types.PrimitiveU32)
	copied.ReferredTime = dsmi.ReferredTime.Copy().(*types.DateTime)
	copied.ExpireTime = dsmi.ExpireTime.Copy().(*types.DateTime)
	copied.Tags = dsmi.Tags.Copy().(*types.List[*types.String])
	copied.Ratings = dsmi.Ratings.Copy().(*types.List[*DataStoreRatingInfoWithSlot])

	return copied
}

// Equals checks if the given DataStoreMetaInfo contains the same data as the current DataStoreMetaInfo
func (dsmi *DataStoreMetaInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreMetaInfo); !ok {
		return false
	}

	other := o.(*DataStoreMetaInfo)

	if dsmi.StructureVersion != other.StructureVersion {
		return false
	}

	if !dsmi.DataID.Equals(other.DataID) {
		return false
	}

	if !dsmi.OwnerID.Equals(other.OwnerID) {
		return false
	}

	if !dsmi.Size.Equals(other.Size) {
		return false
	}

	if !dsmi.DataType.Equals(other.DataType) {
		return false
	}

	if !dsmi.Name.Equals(other.Name) {
		return false
	}

	if !dsmi.MetaBinary.Equals(other.MetaBinary) {
		return false
	}

	if !dsmi.Permission.Equals(other.Permission) {
		return false
	}

	if !dsmi.DelPermission.Equals(other.DelPermission) {
		return false
	}

	if !dsmi.CreatedTime.Equals(other.CreatedTime) {
		return false
	}

	if !dsmi.UpdatedTime.Equals(other.UpdatedTime) {
		return false
	}

	if !dsmi.Period.Equals(other.Period) {
		return false
	}

	if !dsmi.Status.Equals(other.Status) {
		return false
	}

	if !dsmi.ReferredCnt.Equals(other.ReferredCnt) {
		return false
	}

	if !dsmi.ReferDataID.Equals(other.ReferDataID) {
		return false
	}

	if !dsmi.Flag.Equals(other.Flag) {
		return false
	}

	if !dsmi.ReferredTime.Equals(other.ReferredTime) {
		return false
	}

	if !dsmi.ExpireTime.Equals(other.ExpireTime) {
		return false
	}

	if !dsmi.Tags.Equals(other.Tags) {
		return false
	}

	return dsmi.Ratings.Equals(other.Ratings)
}

// String returns the string representation of the DataStoreMetaInfo
func (dsmi *DataStoreMetaInfo) String() string {
	return dsmi.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreMetaInfo using the provided indentation level
func (dsmi *DataStoreMetaInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreMetaInfo{\n")
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, dsmi.DataID))
	b.WriteString(fmt.Sprintf("%sOwnerID: %s,\n", indentationValues, dsmi.OwnerID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sSize: %s,\n", indentationValues, dsmi.Size))
	b.WriteString(fmt.Sprintf("%sDataType: %s,\n", indentationValues, dsmi.DataType))
	b.WriteString(fmt.Sprintf("%sName: %s,\n", indentationValues, dsmi.Name))
	b.WriteString(fmt.Sprintf("%sMetaBinary: %s,\n", indentationValues, dsmi.MetaBinary))
	b.WriteString(fmt.Sprintf("%sPermission: %s,\n", indentationValues, dsmi.Permission.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sDelPermission: %s,\n", indentationValues, dsmi.DelPermission.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sCreatedTime: %s,\n", indentationValues, dsmi.CreatedTime.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sUpdatedTime: %s,\n", indentationValues, dsmi.UpdatedTime.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sPeriod: %s,\n", indentationValues, dsmi.Period))
	b.WriteString(fmt.Sprintf("%sStatus: %s,\n", indentationValues, dsmi.Status))
	b.WriteString(fmt.Sprintf("%sReferredCnt: %s,\n", indentationValues, dsmi.ReferredCnt))
	b.WriteString(fmt.Sprintf("%sReferDataID: %s,\n", indentationValues, dsmi.ReferDataID))
	b.WriteString(fmt.Sprintf("%sFlag: %s,\n", indentationValues, dsmi.Flag))
	b.WriteString(fmt.Sprintf("%sReferredTime: %s,\n", indentationValues, dsmi.ReferredTime.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sExpireTime: %s,\n", indentationValues, dsmi.ExpireTime.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sTags: %s,\n", indentationValues, dsmi.Tags))
	b.WriteString(fmt.Sprintf("%sRatings: %s,\n", indentationValues, dsmi.Ratings))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// FilterPropertiesByResultOption zeroes out certain struct properties based on the input flags
func (dsmi *DataStoreMetaInfo) FilterPropertiesByResultOption(resultOption *types.PrimitiveU8) {
	// * This is kind of backwards
	// *
	// * This method assumes all struct data exists
	// * by default. This is done in order to simplify
	// * database calls by just querying for all fields
	// * at once. Therefore, instead of the ResultOption
	// * flags being used to conditionally ADD properties,
	// * it's used to conditionally REMOVE them

	if resultOption.PAND(0x1) == 0 {
		dsmi.Tags = types.NewList[*types.String]()
		dsmi.Tags.Type = types.NewString("")
	}

	if resultOption.PAND(0x2) == 0 {
		dsmi.Ratings = types.NewList[*DataStoreRatingInfoWithSlot]()
		dsmi.Ratings.Type = NewDataStoreRatingInfoWithSlot()
	}

	if resultOption.PAND(0x4) == 0 {
		dsmi.MetaBinary = types.NewQBuffer(nil)
	}
}

// NewDataStoreMetaInfo returns a new DataStoreMetaInfo
func NewDataStoreMetaInfo() *DataStoreMetaInfo {
	dsmi := &DataStoreMetaInfo{
		DataID:        types.NewPrimitiveU64(0),
		OwnerID:       types.NewPID(0),
		Size:          types.NewPrimitiveU32(0),
		DataType:      types.NewPrimitiveU16(0),
		Name:          types.NewString(""),
		MetaBinary:    types.NewQBuffer(nil),
		Permission:    NewDataStorePermission(),
		DelPermission: NewDataStorePermission(),
		CreatedTime:   types.NewDateTime(0),
		UpdatedTime:   types.NewDateTime(0),
		Period:        types.NewPrimitiveU16(0),
		Status:        types.NewPrimitiveU8(0),
		ReferredCnt:   types.NewPrimitiveU32(0),
		ReferDataID:   types.NewPrimitiveU32(0),
		Flag:          types.NewPrimitiveU32(0),
		ReferredTime:  types.NewDateTime(0),
		ExpireTime:    types.NewDateTime(0),
		Tags:          types.NewList[*types.String](),
		Ratings:       types.NewList[*DataStoreRatingInfoWithSlot](),
	}

	dsmi.Tags.Type = types.NewString("")
	dsmi.Ratings.Type = NewDataStoreRatingInfoWithSlot()

	return dsmi
}
