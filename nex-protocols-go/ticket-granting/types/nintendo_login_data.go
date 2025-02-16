// Package types implements all the types used by the TicketGranting protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// NintendoLoginData is a type within the TicketGranting protocol
type NintendoLoginData struct {
	types.Structure
	Token *types.String
}

// WriteTo writes the NintendoLoginData to the given writable
func (nld *NintendoLoginData) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	nld.Token.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	nld.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the NintendoLoginData from the given readable
func (nld *NintendoLoginData) ExtractFrom(readable types.Readable) error {
	var err error

	err = nld.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoLoginData header. %s", err.Error())
	}

	err = nld.Token.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoLoginData.Token. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of NintendoLoginData
func (nld *NintendoLoginData) Copy() types.RVType {
	copied := NewNintendoLoginData()

	copied.StructureVersion = nld.StructureVersion
	copied.Token = nld.Token.Copy().(*types.String)

	return copied
}

// Equals checks if the given NintendoLoginData contains the same data as the current NintendoLoginData
func (nld *NintendoLoginData) Equals(o types.RVType) bool {
	if _, ok := o.(*NintendoLoginData); !ok {
		return false
	}

	other := o.(*NintendoLoginData)

	if nld.StructureVersion != other.StructureVersion {
		return false
	}

	return nld.Token.Equals(other.Token)
}

// String returns the string representation of the NintendoLoginData
func (nld *NintendoLoginData) String() string {
	return nld.FormatToString(0)
}

// FormatToString pretty-prints the NintendoLoginData using the provided indentation level
func (nld *NintendoLoginData) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("NintendoLoginData{\n")
	b.WriteString(fmt.Sprintf("%sToken: %s,\n", indentationValues, nld.Token))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewNintendoLoginData returns a new NintendoLoginData
func NewNintendoLoginData() *NintendoLoginData {
	nld := &NintendoLoginData{
		Token: types.NewString(""),
	}

	return nld
}
