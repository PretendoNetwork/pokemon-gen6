// Package types implements all the types used by the DataStoreSuperSmashBros.4 protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStorePreparePostSharedDataParam is a type within the DataStoreSuperSmashBros.4 protocol
type DataStorePreparePostSharedDataParam struct {
	types.Structure
	DataType   *types.PrimitiveU8
	Region     *types.PrimitiveU8
	Attribute1 *types.PrimitiveU8
	Attribute2 *types.PrimitiveU8
	Fighter    *types.Buffer
	Size       *types.PrimitiveU32
	Comment    *types.String
	MetaBinary *types.QBuffer
	ExtraData  *types.List[*types.String]
}

// WriteTo writes the DataStorePreparePostSharedDataParam to the given writable
func (dsppsdp *DataStorePreparePostSharedDataParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dsppsdp.DataType.WriteTo(contentWritable)
	dsppsdp.Region.WriteTo(contentWritable)
	dsppsdp.Attribute1.WriteTo(contentWritable)
	dsppsdp.Attribute2.WriteTo(contentWritable)
	dsppsdp.Fighter.WriteTo(contentWritable)
	dsppsdp.Size.WriteTo(contentWritable)
	dsppsdp.Comment.WriteTo(contentWritable)
	dsppsdp.MetaBinary.WriteTo(contentWritable)
	dsppsdp.ExtraData.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dsppsdp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStorePreparePostSharedDataParam from the given readable
func (dsppsdp *DataStorePreparePostSharedDataParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = dsppsdp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostSharedDataParam header. %s", err.Error())
	}

	err = dsppsdp.DataType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostSharedDataParam.DataType. %s", err.Error())
	}

	err = dsppsdp.Region.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostSharedDataParam.Region. %s", err.Error())
	}

	err = dsppsdp.Attribute1.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostSharedDataParam.Attribute1. %s", err.Error())
	}

	err = dsppsdp.Attribute2.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostSharedDataParam.Attribute2. %s", err.Error())
	}

	err = dsppsdp.Fighter.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostSharedDataParam.Fighter. %s", err.Error())
	}

	err = dsppsdp.Size.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostSharedDataParam.Size. %s", err.Error())
	}

	err = dsppsdp.Comment.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostSharedDataParam.Comment. %s", err.Error())
	}

	err = dsppsdp.MetaBinary.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostSharedDataParam.MetaBinary. %s", err.Error())
	}

	err = dsppsdp.ExtraData.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostSharedDataParam.ExtraData. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStorePreparePostSharedDataParam
func (dsppsdp *DataStorePreparePostSharedDataParam) Copy() types.RVType {
	copied := NewDataStorePreparePostSharedDataParam()

	copied.StructureVersion = dsppsdp.StructureVersion
	copied.DataType = dsppsdp.DataType.Copy().(*types.PrimitiveU8)
	copied.Region = dsppsdp.Region.Copy().(*types.PrimitiveU8)
	copied.Attribute1 = dsppsdp.Attribute1.Copy().(*types.PrimitiveU8)
	copied.Attribute2 = dsppsdp.Attribute2.Copy().(*types.PrimitiveU8)
	copied.Fighter = dsppsdp.Fighter.Copy().(*types.Buffer)
	copied.Size = dsppsdp.Size.Copy().(*types.PrimitiveU32)
	copied.Comment = dsppsdp.Comment.Copy().(*types.String)
	copied.MetaBinary = dsppsdp.MetaBinary.Copy().(*types.QBuffer)
	copied.ExtraData = dsppsdp.ExtraData.Copy().(*types.List[*types.String])

	return copied
}

// Equals checks if the given DataStorePreparePostSharedDataParam contains the same data as the current DataStorePreparePostSharedDataParam
func (dsppsdp *DataStorePreparePostSharedDataParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStorePreparePostSharedDataParam); !ok {
		return false
	}

	other := o.(*DataStorePreparePostSharedDataParam)

	if dsppsdp.StructureVersion != other.StructureVersion {
		return false
	}

	if !dsppsdp.DataType.Equals(other.DataType) {
		return false
	}

	if !dsppsdp.Region.Equals(other.Region) {
		return false
	}

	if !dsppsdp.Attribute1.Equals(other.Attribute1) {
		return false
	}

	if !dsppsdp.Attribute2.Equals(other.Attribute2) {
		return false
	}

	if !dsppsdp.Fighter.Equals(other.Fighter) {
		return false
	}

	if !dsppsdp.Size.Equals(other.Size) {
		return false
	}

	if !dsppsdp.Comment.Equals(other.Comment) {
		return false
	}

	if !dsppsdp.MetaBinary.Equals(other.MetaBinary) {
		return false
	}

	return dsppsdp.ExtraData.Equals(other.ExtraData)
}

// String returns the string representation of the DataStorePreparePostSharedDataParam
func (dsppsdp *DataStorePreparePostSharedDataParam) String() string {
	return dsppsdp.FormatToString(0)
}

// FormatToString pretty-prints the DataStorePreparePostSharedDataParam using the provided indentation level
func (dsppsdp *DataStorePreparePostSharedDataParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStorePreparePostSharedDataParam{\n")
	b.WriteString(fmt.Sprintf("%sDataType: %s,\n", indentationValues, dsppsdp.DataType))
	b.WriteString(fmt.Sprintf("%sRegion: %s,\n", indentationValues, dsppsdp.Region))
	b.WriteString(fmt.Sprintf("%sAttribute1: %s,\n", indentationValues, dsppsdp.Attribute1))
	b.WriteString(fmt.Sprintf("%sAttribute2: %s,\n", indentationValues, dsppsdp.Attribute2))
	b.WriteString(fmt.Sprintf("%sFighter: %s,\n", indentationValues, dsppsdp.Fighter))
	b.WriteString(fmt.Sprintf("%sSize: %s,\n", indentationValues, dsppsdp.Size))
	b.WriteString(fmt.Sprintf("%sComment: %s,\n", indentationValues, dsppsdp.Comment))
	b.WriteString(fmt.Sprintf("%sMetaBinary: %s,\n", indentationValues, dsppsdp.MetaBinary))
	b.WriteString(fmt.Sprintf("%sExtraData: %s,\n", indentationValues, dsppsdp.ExtraData))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStorePreparePostSharedDataParam returns a new DataStorePreparePostSharedDataParam
func NewDataStorePreparePostSharedDataParam() *DataStorePreparePostSharedDataParam {
	dsppsdp := &DataStorePreparePostSharedDataParam{
		DataType:   types.NewPrimitiveU8(0),
		Region:     types.NewPrimitiveU8(0),
		Attribute1: types.NewPrimitiveU8(0),
		Attribute2: types.NewPrimitiveU8(0),
		Fighter:    types.NewBuffer(nil),
		Size:       types.NewPrimitiveU32(0),
		Comment:    types.NewString(""),
		MetaBinary: types.NewQBuffer(nil),
		ExtraData:  types.NewList[*types.String](),
	}

	dsppsdp.ExtraData.Type = types.NewString("")

	return dsppsdp
}
