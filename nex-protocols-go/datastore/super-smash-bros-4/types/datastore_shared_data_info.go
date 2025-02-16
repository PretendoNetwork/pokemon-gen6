// Package types implements all the types used by the DataStoreSuperSmashBros.4 protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreSharedDataInfo is a type within the DataStoreSuperSmashBros.4 protocol
type DataStoreSharedDataInfo struct {
	types.Structure
	DataID      *types.PrimitiveU64
	OwnerID     *types.PrimitiveU32
	DataType    *types.PrimitiveU8
	Comment     *types.String
	MetaBinary  *types.QBuffer
	Profile     *types.QBuffer
	Rating      *types.PrimitiveS64
	CreatedTime *types.DateTime
	Info        *DataStoreFileServerObjectInfo
}

// WriteTo writes the DataStoreSharedDataInfo to the given writable
func (dssdi *DataStoreSharedDataInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dssdi.DataID.WriteTo(contentWritable)
	dssdi.OwnerID.WriteTo(contentWritable)
	dssdi.DataType.WriteTo(contentWritable)
	dssdi.Comment.WriteTo(contentWritable)
	dssdi.MetaBinary.WriteTo(contentWritable)
	dssdi.Profile.WriteTo(contentWritable)
	dssdi.Rating.WriteTo(contentWritable)
	dssdi.CreatedTime.WriteTo(contentWritable)
	dssdi.Info.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dssdi.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreSharedDataInfo from the given readable
func (dssdi *DataStoreSharedDataInfo) ExtractFrom(readable types.Readable) error {
	var err error

	err = dssdi.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSharedDataInfo header. %s", err.Error())
	}

	err = dssdi.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSharedDataInfo.DataID. %s", err.Error())
	}

	err = dssdi.OwnerID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSharedDataInfo.OwnerID. %s", err.Error())
	}

	err = dssdi.DataType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSharedDataInfo.DataType. %s", err.Error())
	}

	err = dssdi.Comment.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSharedDataInfo.Comment. %s", err.Error())
	}

	err = dssdi.MetaBinary.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSharedDataInfo.MetaBinary. %s", err.Error())
	}

	err = dssdi.Profile.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSharedDataInfo.Profile. %s", err.Error())
	}

	err = dssdi.Rating.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSharedDataInfo.Rating. %s", err.Error())
	}

	err = dssdi.CreatedTime.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSharedDataInfo.CreatedTime. %s", err.Error())
	}

	err = dssdi.Info.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSharedDataInfo.Info. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreSharedDataInfo
func (dssdi *DataStoreSharedDataInfo) Copy() types.RVType {
	copied := NewDataStoreSharedDataInfo()

	copied.StructureVersion = dssdi.StructureVersion
	copied.DataID = dssdi.DataID.Copy().(*types.PrimitiveU64)
	copied.OwnerID = dssdi.OwnerID.Copy().(*types.PrimitiveU32)
	copied.DataType = dssdi.DataType.Copy().(*types.PrimitiveU8)
	copied.Comment = dssdi.Comment.Copy().(*types.String)
	copied.MetaBinary = dssdi.MetaBinary.Copy().(*types.QBuffer)
	copied.Profile = dssdi.Profile.Copy().(*types.QBuffer)
	copied.Rating = dssdi.Rating.Copy().(*types.PrimitiveS64)
	copied.CreatedTime = dssdi.CreatedTime.Copy().(*types.DateTime)
	copied.Info = dssdi.Info.Copy().(*DataStoreFileServerObjectInfo)

	return copied
}

// Equals checks if the given DataStoreSharedDataInfo contains the same data as the current DataStoreSharedDataInfo
func (dssdi *DataStoreSharedDataInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreSharedDataInfo); !ok {
		return false
	}

	other := o.(*DataStoreSharedDataInfo)

	if dssdi.StructureVersion != other.StructureVersion {
		return false
	}

	if !dssdi.DataID.Equals(other.DataID) {
		return false
	}

	if !dssdi.OwnerID.Equals(other.OwnerID) {
		return false
	}

	if !dssdi.DataType.Equals(other.DataType) {
		return false
	}

	if !dssdi.Comment.Equals(other.Comment) {
		return false
	}

	if !dssdi.MetaBinary.Equals(other.MetaBinary) {
		return false
	}

	if !dssdi.Profile.Equals(other.Profile) {
		return false
	}

	if !dssdi.Rating.Equals(other.Rating) {
		return false
	}

	if !dssdi.CreatedTime.Equals(other.CreatedTime) {
		return false
	}

	return dssdi.Info.Equals(other.Info)
}

// String returns the string representation of the DataStoreSharedDataInfo
func (dssdi *DataStoreSharedDataInfo) String() string {
	return dssdi.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreSharedDataInfo using the provided indentation level
func (dssdi *DataStoreSharedDataInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreSharedDataInfo{\n")
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, dssdi.DataID))
	b.WriteString(fmt.Sprintf("%sOwnerID: %s,\n", indentationValues, dssdi.OwnerID))
	b.WriteString(fmt.Sprintf("%sDataType: %s,\n", indentationValues, dssdi.DataType))
	b.WriteString(fmt.Sprintf("%sComment: %s,\n", indentationValues, dssdi.Comment))
	b.WriteString(fmt.Sprintf("%sMetaBinary: %s,\n", indentationValues, dssdi.MetaBinary))
	b.WriteString(fmt.Sprintf("%sProfile: %s,\n", indentationValues, dssdi.Profile))
	b.WriteString(fmt.Sprintf("%sRating: %s,\n", indentationValues, dssdi.Rating))
	b.WriteString(fmt.Sprintf("%sCreatedTime: %s,\n", indentationValues, dssdi.CreatedTime.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sInfo: %s,\n", indentationValues, dssdi.Info.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreSharedDataInfo returns a new DataStoreSharedDataInfo
func NewDataStoreSharedDataInfo() *DataStoreSharedDataInfo {
	dssdi := &DataStoreSharedDataInfo{
		DataID:      types.NewPrimitiveU64(0),
		OwnerID:     types.NewPrimitiveU32(0),
		DataType:    types.NewPrimitiveU8(0),
		Comment:     types.NewString(""),
		MetaBinary:  types.NewQBuffer(nil),
		Profile:     types.NewQBuffer(nil),
		Rating:      types.NewPrimitiveS64(0),
		CreatedTime: types.NewDateTime(0),
		Info:        NewDataStoreFileServerObjectInfo(),
	}

	return dssdi
}
