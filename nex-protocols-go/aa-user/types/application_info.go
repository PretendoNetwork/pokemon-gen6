// Package types implements all the types used by the AAUser protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ApplicationInfo is a type within the AAUser protocol
type ApplicationInfo struct {
	types.Structure
	*types.Data
	TitleID      *types.PrimitiveU64
	TitleVersion *types.PrimitiveU16
}

// WriteTo writes the ApplicationInfo to the given writable
func (ai *ApplicationInfo) WriteTo(writable types.Writable) {
	ai.Data.WriteTo(writable)

	contentWritable := writable.CopyNew()

	ai.TitleID.WriteTo(contentWritable)
	ai.TitleVersion.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	ai.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ApplicationInfo from the given readable
func (ai *ApplicationInfo) ExtractFrom(readable types.Readable) error {
	var err error

	err = ai.Data.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ApplicationInfo.Data. %s", err.Error())
	}

	err = ai.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ApplicationInfo header. %s", err.Error())
	}

	err = ai.TitleID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ApplicationInfo.TitleID. %s", err.Error())
	}

	err = ai.TitleVersion.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ApplicationInfo.TitleVersion. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ApplicationInfo
func (ai *ApplicationInfo) Copy() types.RVType {
	copied := NewApplicationInfo()

	copied.StructureVersion = ai.StructureVersion
	copied.Data = ai.Data.Copy().(*types.Data)
	copied.TitleID = ai.TitleID.Copy().(*types.PrimitiveU64)
	copied.TitleVersion = ai.TitleVersion.Copy().(*types.PrimitiveU16)

	return copied
}

// Equals checks if the given ApplicationInfo contains the same data as the current ApplicationInfo
func (ai *ApplicationInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*ApplicationInfo); !ok {
		return false
	}

	other := o.(*ApplicationInfo)

	if ai.StructureVersion != other.StructureVersion {
		return false
	}

	if !ai.Data.Equals(other.Data) {
		return false
	}

	if !ai.TitleID.Equals(other.TitleID) {
		return false
	}

	return ai.TitleVersion.Equals(other.TitleVersion)
}

// String returns the string representation of the ApplicationInfo
func (ai *ApplicationInfo) String() string {
	return ai.FormatToString(0)
}

// FormatToString pretty-prints the ApplicationInfo using the provided indentation level
func (ai *ApplicationInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ApplicationInfo{\n")
	b.WriteString(fmt.Sprintf("%sData (parent): %s,\n", indentationValues, ai.Data.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sTitleID: %s,\n", indentationValues, ai.TitleID))
	b.WriteString(fmt.Sprintf("%sTitleVersion: %s,\n", indentationValues, ai.TitleVersion))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewApplicationInfo returns a new ApplicationInfo
func NewApplicationInfo() *ApplicationInfo {
	ai := &ApplicationInfo{
		Data:         types.NewData(),
		TitleID:      types.NewPrimitiveU64(0),
		TitleVersion: types.NewPrimitiveU16(0),
	}

	return ai
}
