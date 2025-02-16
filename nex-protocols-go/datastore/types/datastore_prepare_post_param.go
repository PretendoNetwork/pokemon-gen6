// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStorePreparePostParam is a type within the DataStore protocol
type DataStorePreparePostParam struct {
	types.Structure
	Size                 *types.PrimitiveU32
	Name                 *types.String
	DataType             *types.PrimitiveU16
	MetaBinary           *types.QBuffer
	Permission           *DataStorePermission
	DelPermission        *DataStorePermission
	Flag                 *types.PrimitiveU32
	Period               *types.PrimitiveU16
	ReferDataID          *types.PrimitiveU32
	Tags                 *types.List[*types.String]
	RatingInitParams     *types.List[*DataStoreRatingInitParamWithSlot]
	PersistenceInitParam *DataStorePersistenceInitParam
	ExtraData            *types.List[*types.String] // * NEX v3.5.0
}

// WriteTo writes the DataStorePreparePostParam to the given writable
func (dsppp *DataStorePreparePostParam) WriteTo(writable types.Writable) {
	stream := writable.(*nex.ByteStreamOut)
	libraryVersion := stream.LibraryVersions.DataStore

	contentWritable := writable.CopyNew()

	dsppp.Size.WriteTo(contentWritable)
	dsppp.Name.WriteTo(contentWritable)
	dsppp.DataType.WriteTo(contentWritable)
	dsppp.MetaBinary.WriteTo(contentWritable)
	dsppp.Permission.WriteTo(contentWritable)
	dsppp.DelPermission.WriteTo(contentWritable)
	dsppp.Flag.WriteTo(contentWritable)
	dsppp.Period.WriteTo(contentWritable)
	dsppp.ReferDataID.WriteTo(contentWritable)
	dsppp.Tags.WriteTo(contentWritable)
	dsppp.RatingInitParams.WriteTo(contentWritable)
	dsppp.PersistenceInitParam.WriteTo(contentWritable)

	if libraryVersion.GreaterOrEqual("3.5.0") {
		dsppp.ExtraData.WriteTo(contentWritable)
	}

	content := contentWritable.Bytes()

	dsppp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStorePreparePostParam from the given readable
func (dsppp *DataStorePreparePostParam) ExtractFrom(readable types.Readable) error {
	stream := readable.(*nex.ByteStreamIn)
	libraryVersion := stream.LibraryVersions.DataStore

	var err error

	err = dsppp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParam header. %s", err.Error())
	}

	err = dsppp.Size.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParam.Size. %s", err.Error())
	}

	err = dsppp.Name.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParam.Name. %s", err.Error())
	}

	err = dsppp.DataType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParam.DataType. %s", err.Error())
	}

	err = dsppp.MetaBinary.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParam.MetaBinary. %s", err.Error())
	}

	err = dsppp.Permission.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParam.Permission. %s", err.Error())
	}

	err = dsppp.DelPermission.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParam.DelPermission. %s", err.Error())
	}

	err = dsppp.Flag.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParam.Flag. %s", err.Error())
	}

	err = dsppp.Period.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParam.Period. %s", err.Error())
	}

	err = dsppp.ReferDataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParam.ReferDataID. %s", err.Error())
	}

	err = dsppp.Tags.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParam.Tags. %s", err.Error())
	}

	err = dsppp.RatingInitParams.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParam.RatingInitParams. %s", err.Error())
	}

	err = dsppp.PersistenceInitParam.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParam.PersistenceInitParam. %s", err.Error())
	}

	if libraryVersion.GreaterOrEqual("3.5.0") {
		err = dsppp.ExtraData.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract DataStorePreparePostParam.ExtraData. %s", err.Error())
		}
	}

	return nil
}

// Copy returns a new copied instance of DataStorePreparePostParam
func (dsppp *DataStorePreparePostParam) Copy() types.RVType {
	copied := NewDataStorePreparePostParam()

	copied.StructureVersion = dsppp.StructureVersion
	copied.Size = dsppp.Size.Copy().(*types.PrimitiveU32)
	copied.Name = dsppp.Name.Copy().(*types.String)
	copied.DataType = dsppp.DataType.Copy().(*types.PrimitiveU16)
	copied.MetaBinary = dsppp.MetaBinary.Copy().(*types.QBuffer)
	copied.Permission = dsppp.Permission.Copy().(*DataStorePermission)
	copied.DelPermission = dsppp.DelPermission.Copy().(*DataStorePermission)
	copied.Flag = dsppp.Flag.Copy().(*types.PrimitiveU32)
	copied.Period = dsppp.Period.Copy().(*types.PrimitiveU16)
	copied.ReferDataID = dsppp.ReferDataID.Copy().(*types.PrimitiveU32)
	copied.Tags = dsppp.Tags.Copy().(*types.List[*types.String])
	copied.RatingInitParams = dsppp.RatingInitParams.Copy().(*types.List[*DataStoreRatingInitParamWithSlot])
	copied.PersistenceInitParam = dsppp.PersistenceInitParam.Copy().(*DataStorePersistenceInitParam)
	copied.ExtraData = dsppp.ExtraData.Copy().(*types.List[*types.String])

	return copied
}

// Equals checks if the given DataStorePreparePostParam contains the same data as the current DataStorePreparePostParam
func (dsppp *DataStorePreparePostParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStorePreparePostParam); !ok {
		return false
	}

	other := o.(*DataStorePreparePostParam)

	if dsppp.StructureVersion != other.StructureVersion {
		return false
	}

	if !dsppp.Size.Equals(other.Size) {
		return false
	}

	if !dsppp.Name.Equals(other.Name) {
		return false
	}

	if !dsppp.DataType.Equals(other.DataType) {
		return false
	}

	if !dsppp.MetaBinary.Equals(other.MetaBinary) {
		return false
	}

	if !dsppp.Permission.Equals(other.Permission) {
		return false
	}

	if !dsppp.DelPermission.Equals(other.DelPermission) {
		return false
	}

	if !dsppp.Flag.Equals(other.Flag) {
		return false
	}

	if !dsppp.Period.Equals(other.Period) {
		return false
	}

	if !dsppp.ReferDataID.Equals(other.ReferDataID) {
		return false
	}

	if !dsppp.Tags.Equals(other.Tags) {
		return false
	}

	if !dsppp.RatingInitParams.Equals(other.RatingInitParams) {
		return false
	}

	if !dsppp.PersistenceInitParam.Equals(other.PersistenceInitParam) {
		return false
	}

	return dsppp.ExtraData.Equals(other.ExtraData)
}

// String returns the string representation of the DataStorePreparePostParam
func (dsppp *DataStorePreparePostParam) String() string {
	return dsppp.FormatToString(0)
}

// FormatToString pretty-prints the DataStorePreparePostParam using the provided indentation level
func (dsppp *DataStorePreparePostParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStorePreparePostParam{\n")
	b.WriteString(fmt.Sprintf("%sSize: %s,\n", indentationValues, dsppp.Size))
	b.WriteString(fmt.Sprintf("%sName: %s,\n", indentationValues, dsppp.Name))
	b.WriteString(fmt.Sprintf("%sDataType: %s,\n", indentationValues, dsppp.DataType))
	b.WriteString(fmt.Sprintf("%sMetaBinary: %s,\n", indentationValues, dsppp.MetaBinary))
	b.WriteString(fmt.Sprintf("%sPermission: %s,\n", indentationValues, dsppp.Permission.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sDelPermission: %s,\n", indentationValues, dsppp.DelPermission.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sFlag: %s,\n", indentationValues, dsppp.Flag))
	b.WriteString(fmt.Sprintf("%sPeriod: %s,\n", indentationValues, dsppp.Period))
	b.WriteString(fmt.Sprintf("%sReferDataID: %s,\n", indentationValues, dsppp.ReferDataID))
	b.WriteString(fmt.Sprintf("%sTags: %s,\n", indentationValues, dsppp.Tags))
	b.WriteString(fmt.Sprintf("%sRatingInitParams: %s,\n", indentationValues, dsppp.RatingInitParams))
	b.WriteString(fmt.Sprintf("%sPersistenceInitParam: %s,\n", indentationValues, dsppp.PersistenceInitParam.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sExtraData: %s,\n", indentationValues, dsppp.ExtraData))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStorePreparePostParam returns a new DataStorePreparePostParam
func NewDataStorePreparePostParam() *DataStorePreparePostParam {
	dsppp := &DataStorePreparePostParam{
		Size:                 types.NewPrimitiveU32(0),
		Name:                 types.NewString(""),
		DataType:             types.NewPrimitiveU16(0),
		MetaBinary:           types.NewQBuffer(nil),
		Permission:           NewDataStorePermission(),
		DelPermission:        NewDataStorePermission(),
		Flag:                 types.NewPrimitiveU32(0),
		Period:               types.NewPrimitiveU16(0),
		ReferDataID:          types.NewPrimitiveU32(0),
		Tags:                 types.NewList[*types.String](),
		RatingInitParams:     types.NewList[*DataStoreRatingInitParamWithSlot](),
		PersistenceInitParam: NewDataStorePersistenceInitParam(),
		ExtraData:            types.NewList[*types.String](),
	}

	dsppp.Tags.Type = types.NewString("")
	dsppp.RatingInitParams.Type = NewDataStoreRatingInitParamWithSlot()
	dsppp.ExtraData.Type = types.NewString("")

	return dsppp
}
