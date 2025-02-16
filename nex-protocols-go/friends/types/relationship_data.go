// Package types implements all the types used by the Friends protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// RelationshipData is a type within the Friends protocol
type RelationshipData struct {
	types.Structure
	PID            *types.PrimitiveU32
	StrName        *types.String
	ByRelationship *types.PrimitiveU8
	UIDetails      *types.PrimitiveU32
	ByStatus       *types.PrimitiveU8
}

// WriteTo writes the RelationshipData to the given writable
func (rd *RelationshipData) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	rd.PID.WriteTo(contentWritable)
	rd.StrName.WriteTo(contentWritable)
	rd.ByRelationship.WriteTo(contentWritable)
	rd.UIDetails.WriteTo(contentWritable)
	rd.ByStatus.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	rd.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the RelationshipData from the given readable
func (rd *RelationshipData) ExtractFrom(readable types.Readable) error {
	var err error

	err = rd.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RelationshipData header. %s", err.Error())
	}

	err = rd.PID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RelationshipData.PID. %s", err.Error())
	}

	err = rd.StrName.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RelationshipData.StrName. %s", err.Error())
	}

	err = rd.ByRelationship.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RelationshipData.ByRelationship. %s", err.Error())
	}

	err = rd.UIDetails.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RelationshipData.UIDetails. %s", err.Error())
	}

	err = rd.ByStatus.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RelationshipData.ByStatus. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of RelationshipData
func (rd *RelationshipData) Copy() types.RVType {
	copied := NewRelationshipData()

	copied.StructureVersion = rd.StructureVersion
	copied.PID = rd.PID.Copy().(*types.PrimitiveU32)
	copied.StrName = rd.StrName.Copy().(*types.String)
	copied.ByRelationship = rd.ByRelationship.Copy().(*types.PrimitiveU8)
	copied.UIDetails = rd.UIDetails.Copy().(*types.PrimitiveU32)
	copied.ByStatus = rd.ByStatus.Copy().(*types.PrimitiveU8)

	return copied
}

// Equals checks if the given RelationshipData contains the same data as the current RelationshipData
func (rd *RelationshipData) Equals(o types.RVType) bool {
	if _, ok := o.(*RelationshipData); !ok {
		return false
	}

	other := o.(*RelationshipData)

	if rd.StructureVersion != other.StructureVersion {
		return false
	}

	if !rd.PID.Equals(other.PID) {
		return false
	}

	if !rd.StrName.Equals(other.StrName) {
		return false
	}

	if !rd.ByRelationship.Equals(other.ByRelationship) {
		return false
	}

	if !rd.UIDetails.Equals(other.UIDetails) {
		return false
	}

	return rd.ByStatus.Equals(other.ByStatus)
}

// String returns the string representation of the RelationshipData
func (rd *RelationshipData) String() string {
	return rd.FormatToString(0)
}

// FormatToString pretty-prints the RelationshipData using the provided indentation level
func (rd *RelationshipData) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("RelationshipData{\n")
	b.WriteString(fmt.Sprintf("%sPID: %s,\n", indentationValues, rd.PID))
	b.WriteString(fmt.Sprintf("%sStrName: %s,\n", indentationValues, rd.StrName))
	b.WriteString(fmt.Sprintf("%sByRelationship: %s,\n", indentationValues, rd.ByRelationship))
	b.WriteString(fmt.Sprintf("%sUIDetails: %s,\n", indentationValues, rd.UIDetails))
	b.WriteString(fmt.Sprintf("%sByStatus: %s,\n", indentationValues, rd.ByStatus))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewRelationshipData returns a new RelationshipData
func NewRelationshipData() *RelationshipData {
	rd := &RelationshipData{
		PID:            types.NewPrimitiveU32(0),
		StrName:        types.NewString(""),
		ByRelationship: types.NewPrimitiveU8(0),
		UIDetails:      types.NewPrimitiveU32(0),
		ByStatus:       types.NewPrimitiveU8(0),
	}

	return rd
}
