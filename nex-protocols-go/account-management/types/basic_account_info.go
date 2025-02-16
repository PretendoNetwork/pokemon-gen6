// Package types implements all the types used by the AccountManagement protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// BasicAccountInfo is a type within the AccountManagement protocol
type BasicAccountInfo struct {
	types.Structure
	PIDOwner *types.PID
	StrName  *types.String
}

// WriteTo writes the BasicAccountInfo to the given writable
func (bai *BasicAccountInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	bai.PIDOwner.WriteTo(contentWritable)
	bai.StrName.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	bai.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the BasicAccountInfo from the given readable
func (bai *BasicAccountInfo) ExtractFrom(readable types.Readable) error {
	var err error

	err = bai.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract BasicAccountInfo header. %s", err.Error())
	}

	err = bai.PIDOwner.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract BasicAccountInfo.PIDOwner. %s", err.Error())
	}

	err = bai.StrName.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract BasicAccountInfo.StrName. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of BasicAccountInfo
func (bai *BasicAccountInfo) Copy() types.RVType {
	copied := NewBasicAccountInfo()

	copied.StructureVersion = bai.StructureVersion
	copied.PIDOwner = bai.PIDOwner.Copy().(*types.PID)
	copied.StrName = bai.StrName.Copy().(*types.String)

	return copied
}

// Equals checks if the given BasicAccountInfo contains the same data as the current BasicAccountInfo
func (bai *BasicAccountInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*BasicAccountInfo); !ok {
		return false
	}

	other := o.(*BasicAccountInfo)

	if bai.StructureVersion != other.StructureVersion {
		return false
	}

	if !bai.PIDOwner.Equals(other.PIDOwner) {
		return false
	}

	return bai.StrName.Equals(other.StrName)
}

// String returns the string representation of the BasicAccountInfo
func (bai *BasicAccountInfo) String() string {
	return bai.FormatToString(0)
}

// FormatToString pretty-prints the BasicAccountInfo using the provided indentation level
func (bai *BasicAccountInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("BasicAccountInfo{\n")
	b.WriteString(fmt.Sprintf("%sPIDOwner: %s,\n", indentationValues, bai.PIDOwner.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sStrName: %s,\n", indentationValues, bai.StrName))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewBasicAccountInfo returns a new BasicAccountInfo
func NewBasicAccountInfo() *BasicAccountInfo {
	bai := &BasicAccountInfo{
		PIDOwner: types.NewPID(0),
		StrName:  types.NewString(""),
	}

	return bai
}
