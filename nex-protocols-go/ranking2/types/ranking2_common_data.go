// Package types implements all the types used by the Ranking2 protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// Ranking2CommonData is a type within the Ranking2 protocol
type Ranking2CommonData struct {
	types.Structure
	UserName   *types.String
	Mii        *types.QBuffer
	BinaryData *types.QBuffer
}

// WriteTo writes the Ranking2CommonData to the given writable
func (rcd *Ranking2CommonData) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	rcd.UserName.WriteTo(contentWritable)
	rcd.Mii.WriteTo(contentWritable)
	rcd.BinaryData.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	rcd.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the Ranking2CommonData from the given readable
func (rcd *Ranking2CommonData) ExtractFrom(readable types.Readable) error {
	var err error

	err = rcd.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2CommonData header. %s", err.Error())
	}

	err = rcd.UserName.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2CommonData.UserName. %s", err.Error())
	}

	err = rcd.Mii.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2CommonData.Mii. %s", err.Error())
	}

	err = rcd.BinaryData.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2CommonData.BinaryData. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of Ranking2CommonData
func (rcd *Ranking2CommonData) Copy() types.RVType {
	copied := NewRanking2CommonData()

	copied.StructureVersion = rcd.StructureVersion
	copied.UserName = rcd.UserName.Copy().(*types.String)
	copied.Mii = rcd.Mii.Copy().(*types.QBuffer)
	copied.BinaryData = rcd.BinaryData.Copy().(*types.QBuffer)

	return copied
}

// Equals checks if the given Ranking2CommonData contains the same data as the current Ranking2CommonData
func (rcd *Ranking2CommonData) Equals(o types.RVType) bool {
	if _, ok := o.(*Ranking2CommonData); !ok {
		return false
	}

	other := o.(*Ranking2CommonData)

	if rcd.StructureVersion != other.StructureVersion {
		return false
	}

	if !rcd.UserName.Equals(other.UserName) {
		return false
	}

	if !rcd.Mii.Equals(other.Mii) {
		return false
	}

	return rcd.BinaryData.Equals(other.BinaryData)
}

// String returns the string representation of the Ranking2CommonData
func (rcd *Ranking2CommonData) String() string {
	return rcd.FormatToString(0)
}

// FormatToString pretty-prints the Ranking2CommonData using the provided indentation level
func (rcd *Ranking2CommonData) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("Ranking2CommonData{\n")
	b.WriteString(fmt.Sprintf("%sUserName: %s,\n", indentationValues, rcd.UserName))
	b.WriteString(fmt.Sprintf("%sMii: %s,\n", indentationValues, rcd.Mii))
	b.WriteString(fmt.Sprintf("%sBinaryData: %s,\n", indentationValues, rcd.BinaryData))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewRanking2CommonData returns a new Ranking2CommonData
func NewRanking2CommonData() *Ranking2CommonData {
	rcd := &Ranking2CommonData{
		UserName:   types.NewString(""),
		Mii:        types.NewQBuffer(nil),
		BinaryData: types.NewQBuffer(nil),
	}

	return rcd
}
