// Package types implements all the types used by the MatchmakeExtension protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// SimpleSearchObject is a type within the MatchmakeExtension protocol
type SimpleSearchObject struct {
	types.Structure
	ObjectID            *types.PrimitiveU32
	OwnerPID            *types.PID
	Attributes          *types.List[*types.PrimitiveU32]
	Metadata            *types.QBuffer
	CommunityIDMiiverse *types.PrimitiveU32
	CommunityCode       *types.String
	DatetimeAttribute   *SimpleSearchDateTimeAttribute
}

// WriteTo writes the SimpleSearchObject to the given writable
func (sso *SimpleSearchObject) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	sso.ObjectID.WriteTo(contentWritable)
	sso.OwnerPID.WriteTo(contentWritable)
	sso.Attributes.WriteTo(contentWritable)
	sso.Metadata.WriteTo(contentWritable)
	sso.CommunityIDMiiverse.WriteTo(contentWritable)
	sso.CommunityCode.WriteTo(contentWritable)
	sso.DatetimeAttribute.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	sso.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the SimpleSearchObject from the given readable
func (sso *SimpleSearchObject) ExtractFrom(readable types.Readable) error {
	var err error

	err = sso.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchObject header. %s", err.Error())
	}

	err = sso.ObjectID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchObject.ObjectID. %s", err.Error())
	}

	err = sso.OwnerPID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchObject.OwnerPID. %s", err.Error())
	}

	err = sso.Attributes.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchObject.Attributes. %s", err.Error())
	}

	err = sso.Metadata.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchObject.Metadata. %s", err.Error())
	}

	err = sso.CommunityIDMiiverse.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchObject.CommunityIDMiiverse. %s", err.Error())
	}

	err = sso.CommunityCode.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchObject.CommunityCode. %s", err.Error())
	}

	err = sso.DatetimeAttribute.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchObject.DatetimeAttribute. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of SimpleSearchObject
func (sso *SimpleSearchObject) Copy() types.RVType {
	copied := NewSimpleSearchObject()

	copied.StructureVersion = sso.StructureVersion
	copied.ObjectID = sso.ObjectID.Copy().(*types.PrimitiveU32)
	copied.OwnerPID = sso.OwnerPID.Copy().(*types.PID)
	copied.Attributes = sso.Attributes.Copy().(*types.List[*types.PrimitiveU32])
	copied.Metadata = sso.Metadata.Copy().(*types.QBuffer)
	copied.CommunityIDMiiverse = sso.CommunityIDMiiverse.Copy().(*types.PrimitiveU32)
	copied.CommunityCode = sso.CommunityCode.Copy().(*types.String)
	copied.DatetimeAttribute = sso.DatetimeAttribute.Copy().(*SimpleSearchDateTimeAttribute)

	return copied
}

// Equals checks if the given SimpleSearchObject contains the same data as the current SimpleSearchObject
func (sso *SimpleSearchObject) Equals(o types.RVType) bool {
	if _, ok := o.(*SimpleSearchObject); !ok {
		return false
	}

	other := o.(*SimpleSearchObject)

	if sso.StructureVersion != other.StructureVersion {
		return false
	}

	if !sso.ObjectID.Equals(other.ObjectID) {
		return false
	}

	if !sso.OwnerPID.Equals(other.OwnerPID) {
		return false
	}

	if !sso.Attributes.Equals(other.Attributes) {
		return false
	}

	if !sso.Metadata.Equals(other.Metadata) {
		return false
	}

	if !sso.CommunityIDMiiverse.Equals(other.CommunityIDMiiverse) {
		return false
	}

	if !sso.CommunityCode.Equals(other.CommunityCode) {
		return false
	}

	return sso.DatetimeAttribute.Equals(other.DatetimeAttribute)
}

// String returns the string representation of the SimpleSearchObject
func (sso *SimpleSearchObject) String() string {
	return sso.FormatToString(0)
}

// FormatToString pretty-prints the SimpleSearchObject using the provided indentation level
func (sso *SimpleSearchObject) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("SimpleSearchObject{\n")
	b.WriteString(fmt.Sprintf("%sObjectID: %s,\n", indentationValues, sso.ObjectID))
	b.WriteString(fmt.Sprintf("%sOwnerPID: %s,\n", indentationValues, sso.OwnerPID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sAttributes: %s,\n", indentationValues, sso.Attributes))
	b.WriteString(fmt.Sprintf("%sMetadata: %s,\n", indentationValues, sso.Metadata))
	b.WriteString(fmt.Sprintf("%sCommunityIDMiiverse: %s,\n", indentationValues, sso.CommunityIDMiiverse))
	b.WriteString(fmt.Sprintf("%sCommunityCode: %s,\n", indentationValues, sso.CommunityCode))
	b.WriteString(fmt.Sprintf("%sDatetimeAttribute: %s,\n", indentationValues, sso.DatetimeAttribute.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewSimpleSearchObject returns a new SimpleSearchObject
func NewSimpleSearchObject() *SimpleSearchObject {
	sso := &SimpleSearchObject{
		ObjectID:            types.NewPrimitiveU32(0),
		OwnerPID:            types.NewPID(0),
		Attributes:          types.NewList[*types.PrimitiveU32](),
		Metadata:            types.NewQBuffer(nil),
		CommunityIDMiiverse: types.NewPrimitiveU32(0),
		CommunityCode:       types.NewString(""),
		DatetimeAttribute:   NewSimpleSearchDateTimeAttribute(),
	}

	sso.Attributes.Type = types.NewPrimitiveU32(0)

	return sso
}
