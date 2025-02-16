// Package types implements all the types used by the Friends3DS protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// Mii is a type within the Friends3DS protocol
type Mii struct {
	types.Structure
	*types.Data
	Name     *types.String
	Unknown2 *types.PrimitiveBool
	Unknown3 *types.PrimitiveU8
	MiiData  *types.Buffer
}

// WriteTo writes the Mii to the given writable
func (m *Mii) WriteTo(writable types.Writable) {
	m.Data.WriteTo(writable)

	contentWritable := writable.CopyNew()

	m.Name.WriteTo(contentWritable)
	m.Unknown2.WriteTo(contentWritable)
	m.Unknown3.WriteTo(contentWritable)
	m.MiiData.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	m.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the Mii from the given readable
func (m *Mii) ExtractFrom(readable types.Readable) error {
	var err error

	err = m.Data.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Mii.Data. %s", err.Error())
	}

	err = m.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Mii header. %s", err.Error())
	}

	err = m.Name.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Mii.Name. %s", err.Error())
	}

	err = m.Unknown2.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Mii.Unknown2. %s", err.Error())
	}

	err = m.Unknown3.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Mii.Unknown3. %s", err.Error())
	}

	err = m.MiiData.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Mii.MiiData. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of Mii
func (m *Mii) Copy() types.RVType {
	copied := NewMii()

	copied.StructureVersion = m.StructureVersion
	copied.Data = m.Data.Copy().(*types.Data)
	copied.Name = m.Name.Copy().(*types.String)
	copied.Unknown2 = m.Unknown2.Copy().(*types.PrimitiveBool)
	copied.Unknown3 = m.Unknown3.Copy().(*types.PrimitiveU8)
	copied.MiiData = m.MiiData.Copy().(*types.Buffer)

	return copied
}

// Equals checks if the given Mii contains the same data as the current Mii
func (m *Mii) Equals(o types.RVType) bool {
	if _, ok := o.(*Mii); !ok {
		return false
	}

	other := o.(*Mii)

	if m.StructureVersion != other.StructureVersion {
		return false
	}

	if !m.Data.Equals(other.Data) {
		return false
	}

	if !m.Name.Equals(other.Name) {
		return false
	}

	if !m.Unknown2.Equals(other.Unknown2) {
		return false
	}

	if !m.Unknown3.Equals(other.Unknown3) {
		return false
	}

	return m.MiiData.Equals(other.MiiData)
}

// String returns the string representation of the Mii
func (m *Mii) String() string {
	return m.FormatToString(0)
}

// FormatToString pretty-prints the Mii using the provided indentation level
func (m *Mii) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("Mii{\n")
	b.WriteString(fmt.Sprintf("%sData (parent): %s,\n", indentationValues, m.Data.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sName: %s,\n", indentationValues, m.Name))
	b.WriteString(fmt.Sprintf("%sUnknown2: %s,\n", indentationValues, m.Unknown2))
	b.WriteString(fmt.Sprintf("%sUnknown3: %s,\n", indentationValues, m.Unknown3))
	b.WriteString(fmt.Sprintf("%sMiiData: %s,\n", indentationValues, m.MiiData))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewMii returns a new Mii
func NewMii() *Mii {
	m := &Mii{
		Data:     types.NewData(),
		Name:     types.NewString(""),
		Unknown2: types.NewPrimitiveBool(false),
		Unknown3: types.NewPrimitiveU8(0),
		MiiData:  types.NewBuffer(nil),
	}

	return m
}
