// Package types implements all the types used by the Utility protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// UniqueIDInfo is a type within the Utility protocol
type UniqueIDInfo struct {
	types.Structure
	NEXUniqueID         *types.PrimitiveU64
	NEXUniqueIDPassword *types.PrimitiveU64
}

// WriteTo writes the UniqueIDInfo to the given writable
func (uidi *UniqueIDInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	uidi.NEXUniqueID.WriteTo(contentWritable)
	uidi.NEXUniqueIDPassword.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	uidi.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the UniqueIDInfo from the given readable
func (uidi *UniqueIDInfo) ExtractFrom(readable types.Readable) error {
	var err error

	err = uidi.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UniqueIDInfo header. %s", err.Error())
	}

	err = uidi.NEXUniqueID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UniqueIDInfo.NEXUniqueID. %s", err.Error())
	}

	err = uidi.NEXUniqueIDPassword.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UniqueIDInfo.NEXUniqueIDPassword. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of UniqueIDInfo
func (uidi *UniqueIDInfo) Copy() types.RVType {
	copied := NewUniqueIDInfo()

	copied.StructureVersion = uidi.StructureVersion
	copied.NEXUniqueID = uidi.NEXUniqueID.Copy().(*types.PrimitiveU64)
	copied.NEXUniqueIDPassword = uidi.NEXUniqueIDPassword.Copy().(*types.PrimitiveU64)

	return copied
}

// Equals checks if the given UniqueIDInfo contains the same data as the current UniqueIDInfo
func (uidi *UniqueIDInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*UniqueIDInfo); !ok {
		return false
	}

	other := o.(*UniqueIDInfo)

	if uidi.StructureVersion != other.StructureVersion {
		return false
	}

	if !uidi.NEXUniqueID.Equals(other.NEXUniqueID) {
		return false
	}

	return uidi.NEXUniqueIDPassword.Equals(other.NEXUniqueIDPassword)
}

// String returns the string representation of the UniqueIDInfo
func (uidi *UniqueIDInfo) String() string {
	return uidi.FormatToString(0)
}

// FormatToString pretty-prints the UniqueIDInfo using the provided indentation level
func (uidi *UniqueIDInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("UniqueIDInfo{\n")
	b.WriteString(fmt.Sprintf("%sNEXUniqueID: %s,\n", indentationValues, uidi.NEXUniqueID))
	b.WriteString(fmt.Sprintf("%sNEXUniqueIDPassword: %s,\n", indentationValues, uidi.NEXUniqueIDPassword))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewUniqueIDInfo returns a new UniqueIDInfo
func NewUniqueIDInfo() *UniqueIDInfo {
	uidi := &UniqueIDInfo{
		NEXUniqueID:         types.NewPrimitiveU64(0),
		NEXUniqueIDPassword: types.NewPrimitiveU64(0),
	}

	return uidi
}
