// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreSearchParam is a type within the DataStore protocol
type DataStoreSearchParam struct {
	types.Structure
	SearchTarget           *types.PrimitiveU8
	OwnerIDs               *types.List[*types.PID]
	OwnerType              *types.PrimitiveU8
	DestinationIDs         *types.List[*types.PID]
	DataType               *types.PrimitiveU16
	CreatedAfter           *types.DateTime
	CreatedBefore          *types.DateTime
	UpdatedAfter           *types.DateTime
	UpdatedBefore          *types.DateTime
	ReferDataID            *types.PrimitiveU32
	Tags                   *types.List[*types.String]
	ResultOrderColumn      *types.PrimitiveU8
	ResultOrder            *types.PrimitiveU8
	ResultRange            *types.ResultRange
	ResultOption           *types.PrimitiveU8
	MinimalRatingFrequency *types.PrimitiveU32
	UseCache               *types.PrimitiveBool
	TotalCountEnabled      *types.PrimitiveBool
	DataTypes              *types.List[*types.PrimitiveU16]
}

// WriteTo writes the DataStoreSearchParam to the given writable
func (dssp *DataStoreSearchParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dssp.SearchTarget.WriteTo(contentWritable)
	dssp.OwnerIDs.WriteTo(contentWritable)
	dssp.OwnerType.WriteTo(contentWritable)
	dssp.DestinationIDs.WriteTo(contentWritable)
	dssp.DataType.WriteTo(contentWritable)
	dssp.CreatedAfter.WriteTo(contentWritable)
	dssp.CreatedBefore.WriteTo(contentWritable)
	dssp.UpdatedAfter.WriteTo(contentWritable)
	dssp.UpdatedBefore.WriteTo(contentWritable)
	dssp.ReferDataID.WriteTo(contentWritable)
	dssp.Tags.WriteTo(contentWritable)
	dssp.ResultOrderColumn.WriteTo(contentWritable)
	dssp.ResultOrder.WriteTo(contentWritable)
	dssp.ResultRange.WriteTo(contentWritable)
	dssp.ResultOption.WriteTo(contentWritable)
	dssp.MinimalRatingFrequency.WriteTo(contentWritable)
	dssp.UseCache.WriteTo(contentWritable)
	dssp.TotalCountEnabled.WriteTo(contentWritable)
	dssp.DataTypes.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dssp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreSearchParam from the given readable
func (dssp *DataStoreSearchParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = dssp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchParam header. %s", err.Error())
	}

	err = dssp.SearchTarget.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchParam.SearchTarget. %s", err.Error())
	}

	err = dssp.OwnerIDs.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchParam.OwnerIDs. %s", err.Error())
	}

	err = dssp.OwnerType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchParam.OwnerType. %s", err.Error())
	}

	err = dssp.DestinationIDs.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchParam.DestinationIDs. %s", err.Error())
	}

	err = dssp.DataType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchParam.DataType. %s", err.Error())
	}

	err = dssp.CreatedAfter.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchParam.CreatedAfter. %s", err.Error())
	}

	err = dssp.CreatedBefore.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchParam.CreatedBefore. %s", err.Error())
	}

	err = dssp.UpdatedAfter.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchParam.UpdatedAfter. %s", err.Error())
	}

	err = dssp.UpdatedBefore.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchParam.UpdatedBefore. %s", err.Error())
	}

	err = dssp.ReferDataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchParam.ReferDataID. %s", err.Error())
	}

	err = dssp.Tags.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchParam.Tags. %s", err.Error())
	}

	err = dssp.ResultOrderColumn.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchParam.ResultOrderColumn. %s", err.Error())
	}

	err = dssp.ResultOrder.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchParam.ResultOrder. %s", err.Error())
	}

	err = dssp.ResultRange.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchParam.ResultRange. %s", err.Error())
	}

	err = dssp.ResultOption.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchParam.ResultOption. %s", err.Error())
	}

	err = dssp.MinimalRatingFrequency.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchParam.MinimalRatingFrequency. %s", err.Error())
	}

	err = dssp.UseCache.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchParam.UseCache. %s", err.Error())
	}

	err = dssp.TotalCountEnabled.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchParam.TotalCountEnabled. %s", err.Error())
	}

	err = dssp.DataTypes.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchParam.DataTypes. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreSearchParam
func (dssp *DataStoreSearchParam) Copy() types.RVType {
	copied := NewDataStoreSearchParam()

	copied.StructureVersion = dssp.StructureVersion
	copied.SearchTarget = dssp.SearchTarget.Copy().(*types.PrimitiveU8)
	copied.OwnerIDs = dssp.OwnerIDs.Copy().(*types.List[*types.PID])
	copied.OwnerType = dssp.OwnerType.Copy().(*types.PrimitiveU8)
	copied.DestinationIDs = dssp.DestinationIDs.Copy().(*types.List[*types.PID])
	copied.DataType = dssp.DataType.Copy().(*types.PrimitiveU16)
	copied.CreatedAfter = dssp.CreatedAfter.Copy().(*types.DateTime)
	copied.CreatedBefore = dssp.CreatedBefore.Copy().(*types.DateTime)
	copied.UpdatedAfter = dssp.UpdatedAfter.Copy().(*types.DateTime)
	copied.UpdatedBefore = dssp.UpdatedBefore.Copy().(*types.DateTime)
	copied.ReferDataID = dssp.ReferDataID.Copy().(*types.PrimitiveU32)
	copied.Tags = dssp.Tags.Copy().(*types.List[*types.String])
	copied.ResultOrderColumn = dssp.ResultOrderColumn.Copy().(*types.PrimitiveU8)
	copied.ResultOrder = dssp.ResultOrder.Copy().(*types.PrimitiveU8)
	copied.ResultRange = dssp.ResultRange.Copy().(*types.ResultRange)
	copied.ResultOption = dssp.ResultOption.Copy().(*types.PrimitiveU8)
	copied.MinimalRatingFrequency = dssp.MinimalRatingFrequency.Copy().(*types.PrimitiveU32)
	copied.UseCache = dssp.UseCache.Copy().(*types.PrimitiveBool)
	copied.TotalCountEnabled = dssp.TotalCountEnabled.Copy().(*types.PrimitiveBool)
	copied.DataTypes = dssp.DataTypes.Copy().(*types.List[*types.PrimitiveU16])

	return copied
}

// Equals checks if the given DataStoreSearchParam contains the same data as the current DataStoreSearchParam
func (dssp *DataStoreSearchParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreSearchParam); !ok {
		return false
	}

	other := o.(*DataStoreSearchParam)

	if dssp.StructureVersion != other.StructureVersion {
		return false
	}

	if !dssp.SearchTarget.Equals(other.SearchTarget) {
		return false
	}

	if !dssp.OwnerIDs.Equals(other.OwnerIDs) {
		return false
	}

	if !dssp.OwnerType.Equals(other.OwnerType) {
		return false
	}

	if !dssp.DestinationIDs.Equals(other.DestinationIDs) {
		return false
	}

	if !dssp.DataType.Equals(other.DataType) {
		return false
	}

	if !dssp.CreatedAfter.Equals(other.CreatedAfter) {
		return false
	}

	if !dssp.CreatedBefore.Equals(other.CreatedBefore) {
		return false
	}

	if !dssp.UpdatedAfter.Equals(other.UpdatedAfter) {
		return false
	}

	if !dssp.UpdatedBefore.Equals(other.UpdatedBefore) {
		return false
	}

	if !dssp.ReferDataID.Equals(other.ReferDataID) {
		return false
	}

	if !dssp.Tags.Equals(other.Tags) {
		return false
	}

	if !dssp.ResultOrderColumn.Equals(other.ResultOrderColumn) {
		return false
	}

	if !dssp.ResultOrder.Equals(other.ResultOrder) {
		return false
	}

	if !dssp.ResultRange.Equals(other.ResultRange) {
		return false
	}

	if !dssp.ResultOption.Equals(other.ResultOption) {
		return false
	}

	if !dssp.MinimalRatingFrequency.Equals(other.MinimalRatingFrequency) {
		return false
	}

	if !dssp.UseCache.Equals(other.UseCache) {
		return false
	}

	if !dssp.TotalCountEnabled.Equals(other.TotalCountEnabled) {
		return false
	}

	return dssp.DataTypes.Equals(other.DataTypes)
}

// String returns the string representation of the DataStoreSearchParam
func (dssp *DataStoreSearchParam) String() string {
	return dssp.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreSearchParam using the provided indentation level
func (dssp *DataStoreSearchParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreSearchParam{\n")
	b.WriteString(fmt.Sprintf("%sSearchTarget: %s,\n", indentationValues, dssp.SearchTarget))
	b.WriteString(fmt.Sprintf("%sOwnerIDs: %s,\n", indentationValues, dssp.OwnerIDs))
	b.WriteString(fmt.Sprintf("%sOwnerType: %s,\n", indentationValues, dssp.OwnerType))
	b.WriteString(fmt.Sprintf("%sDestinationIDs: %s,\n", indentationValues, dssp.DestinationIDs))
	b.WriteString(fmt.Sprintf("%sDataType: %s,\n", indentationValues, dssp.DataType))
	b.WriteString(fmt.Sprintf("%sCreatedAfter: %s,\n", indentationValues, dssp.CreatedAfter.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sCreatedBefore: %s,\n", indentationValues, dssp.CreatedBefore.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sUpdatedAfter: %s,\n", indentationValues, dssp.UpdatedAfter.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sUpdatedBefore: %s,\n", indentationValues, dssp.UpdatedBefore.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sReferDataID: %s,\n", indentationValues, dssp.ReferDataID))
	b.WriteString(fmt.Sprintf("%sTags: %s,\n", indentationValues, dssp.Tags))
	b.WriteString(fmt.Sprintf("%sResultOrderColumn: %s,\n", indentationValues, dssp.ResultOrderColumn))
	b.WriteString(fmt.Sprintf("%sResultOrder: %s,\n", indentationValues, dssp.ResultOrder))
	b.WriteString(fmt.Sprintf("%sResultRange: %s,\n", indentationValues, dssp.ResultRange.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sResultOption: %s,\n", indentationValues, dssp.ResultOption))
	b.WriteString(fmt.Sprintf("%sMinimalRatingFrequency: %s,\n", indentationValues, dssp.MinimalRatingFrequency))
	b.WriteString(fmt.Sprintf("%sUseCache: %s,\n", indentationValues, dssp.UseCache))
	b.WriteString(fmt.Sprintf("%sTotalCountEnabled: %s,\n", indentationValues, dssp.TotalCountEnabled))
	b.WriteString(fmt.Sprintf("%sDataTypes: %s,\n", indentationValues, dssp.DataTypes))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreSearchParam returns a new DataStoreSearchParam
func NewDataStoreSearchParam() *DataStoreSearchParam {
	dssp := &DataStoreSearchParam{
		SearchTarget:           types.NewPrimitiveU8(0),
		OwnerIDs:               types.NewList[*types.PID](),
		OwnerType:              types.NewPrimitiveU8(0),
		DestinationIDs:         types.NewList[*types.PID](),
		DataType:               types.NewPrimitiveU16(0),
		CreatedAfter:           types.NewDateTime(0),
		CreatedBefore:          types.NewDateTime(0),
		UpdatedAfter:           types.NewDateTime(0),
		UpdatedBefore:          types.NewDateTime(0),
		ReferDataID:            types.NewPrimitiveU32(0),
		Tags:                   types.NewList[*types.String](),
		ResultOrderColumn:      types.NewPrimitiveU8(0),
		ResultOrder:            types.NewPrimitiveU8(0),
		ResultRange:            types.NewResultRange(),
		ResultOption:           types.NewPrimitiveU8(0),
		MinimalRatingFrequency: types.NewPrimitiveU32(0),
		UseCache:               types.NewPrimitiveBool(false),
		TotalCountEnabled:      types.NewPrimitiveBool(false),
		DataTypes:              types.NewList[*types.PrimitiveU16](),
	}

	dssp.OwnerIDs.Type = types.NewPID(0)
	dssp.DestinationIDs.Type = types.NewPID(0)
	dssp.Tags.Type = types.NewString("")
	dssp.DataTypes.Type = types.NewPrimitiveU16(0)

	return dssp
}
