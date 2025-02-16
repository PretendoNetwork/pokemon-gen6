// Package types implements all the types used by the Friends3DS protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// FriendRelationship is a type within the Friends3DS protocol
type FriendRelationship struct {
	types.Structure
	*types.Data
	PID              *types.PID
	LFC              *types.PrimitiveU64
	RelationshipType *types.PrimitiveU8
}

// WriteTo writes the FriendRelationship to the given writable
func (fr *FriendRelationship) WriteTo(writable types.Writable) {
	fr.Data.WriteTo(writable)

	contentWritable := writable.CopyNew()

	fr.PID.WriteTo(contentWritable)
	fr.LFC.WriteTo(contentWritable)
	fr.RelationshipType.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	fr.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the FriendRelationship from the given readable
func (fr *FriendRelationship) ExtractFrom(readable types.Readable) error {
	var err error

	err = fr.Data.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendRelationship.Data. %s", err.Error())
	}

	err = fr.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendRelationship header. %s", err.Error())
	}

	err = fr.PID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendRelationship.PID. %s", err.Error())
	}

	err = fr.LFC.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendRelationship.LFC. %s", err.Error())
	}

	err = fr.RelationshipType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendRelationship.RelationshipType. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of FriendRelationship
func (fr *FriendRelationship) Copy() types.RVType {
	copied := NewFriendRelationship()

	copied.StructureVersion = fr.StructureVersion
	copied.Data = fr.Data.Copy().(*types.Data)
	copied.PID = fr.PID.Copy().(*types.PID)
	copied.LFC = fr.LFC.Copy().(*types.PrimitiveU64)
	copied.RelationshipType = fr.RelationshipType.Copy().(*types.PrimitiveU8)

	return copied
}

// Equals checks if the given FriendRelationship contains the same data as the current FriendRelationship
func (fr *FriendRelationship) Equals(o types.RVType) bool {
	if _, ok := o.(*FriendRelationship); !ok {
		return false
	}

	other := o.(*FriendRelationship)

	if fr.StructureVersion != other.StructureVersion {
		return false
	}

	if !fr.Data.Equals(other.Data) {
		return false
	}

	if !fr.PID.Equals(other.PID) {
		return false
	}

	if !fr.LFC.Equals(other.LFC) {
		return false
	}

	return fr.RelationshipType.Equals(other.RelationshipType)
}

// String returns the string representation of the FriendRelationship
func (fr *FriendRelationship) String() string {
	return fr.FormatToString(0)
}

// FormatToString pretty-prints the FriendRelationship using the provided indentation level
func (fr *FriendRelationship) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("FriendRelationship{\n")
	b.WriteString(fmt.Sprintf("%sData (parent): %s,\n", indentationValues, fr.Data.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sPID: %s,\n", indentationValues, fr.PID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sLFC: %s,\n", indentationValues, fr.LFC))
	b.WriteString(fmt.Sprintf("%sRelationshipType: %s,\n", indentationValues, fr.RelationshipType))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewFriendRelationship returns a new FriendRelationship
func NewFriendRelationship() *FriendRelationship {
	fr := &FriendRelationship{
		Data:             types.NewData(),
		PID:              types.NewPID(0),
		LFC:              types.NewPrimitiveU64(0),
		RelationshipType: types.NewPrimitiveU8(0),
	}

	return fr
}
