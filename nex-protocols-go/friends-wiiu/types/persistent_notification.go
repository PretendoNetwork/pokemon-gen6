// Package types implements all the types used by the FriendsWiiU protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// PersistentNotification is a type within the FriendsWiiU protocol
type PersistentNotification struct {
	types.Structure
	*types.Data
	Unknown1 *types.PrimitiveU64
	Unknown2 *types.PrimitiveU32
	Unknown3 *types.PrimitiveU32
	Unknown4 *types.PrimitiveU32
	Unknown5 *types.String
}

// WriteTo writes the PersistentNotification to the given writable
func (pn *PersistentNotification) WriteTo(writable types.Writable) {
	pn.Data.WriteTo(writable)

	contentWritable := writable.CopyNew()

	pn.Unknown1.WriteTo(contentWritable)
	pn.Unknown2.WriteTo(contentWritable)
	pn.Unknown3.WriteTo(contentWritable)
	pn.Unknown4.WriteTo(contentWritable)
	pn.Unknown5.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	pn.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the PersistentNotification from the given readable
func (pn *PersistentNotification) ExtractFrom(readable types.Readable) error {
	var err error

	err = pn.Data.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract PersistentNotification.Data. %s", err.Error())
	}

	err = pn.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract PersistentNotification header. %s", err.Error())
	}

	err = pn.Unknown1.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract PersistentNotification.Unknown1. %s", err.Error())
	}

	err = pn.Unknown2.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract PersistentNotification.Unknown2. %s", err.Error())
	}

	err = pn.Unknown3.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract PersistentNotification.Unknown3. %s", err.Error())
	}

	err = pn.Unknown4.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract PersistentNotification.Unknown4. %s", err.Error())
	}

	err = pn.Unknown5.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract PersistentNotification.Unknown5. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of PersistentNotification
func (pn *PersistentNotification) Copy() types.RVType {
	copied := NewPersistentNotification()

	copied.StructureVersion = pn.StructureVersion
	copied.Data = pn.Data.Copy().(*types.Data)
	copied.Unknown1 = pn.Unknown1.Copy().(*types.PrimitiveU64)
	copied.Unknown2 = pn.Unknown2.Copy().(*types.PrimitiveU32)
	copied.Unknown3 = pn.Unknown3.Copy().(*types.PrimitiveU32)
	copied.Unknown4 = pn.Unknown4.Copy().(*types.PrimitiveU32)
	copied.Unknown5 = pn.Unknown5.Copy().(*types.String)

	return copied
}

// Equals checks if the given PersistentNotification contains the same data as the current PersistentNotification
func (pn *PersistentNotification) Equals(o types.RVType) bool {
	if _, ok := o.(*PersistentNotification); !ok {
		return false
	}

	other := o.(*PersistentNotification)

	if pn.StructureVersion != other.StructureVersion {
		return false
	}

	if !pn.Data.Equals(other.Data) {
		return false
	}

	if !pn.Unknown1.Equals(other.Unknown1) {
		return false
	}

	if !pn.Unknown2.Equals(other.Unknown2) {
		return false
	}

	if !pn.Unknown3.Equals(other.Unknown3) {
		return false
	}

	if !pn.Unknown4.Equals(other.Unknown4) {
		return false
	}

	return pn.Unknown5.Equals(other.Unknown5)
}

// String returns the string representation of the PersistentNotification
func (pn *PersistentNotification) String() string {
	return pn.FormatToString(0)
}

// FormatToString pretty-prints the PersistentNotification using the provided indentation level
func (pn *PersistentNotification) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("PersistentNotification{\n")
	b.WriteString(fmt.Sprintf("%sData (parent): %s,\n", indentationValues, pn.Data.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sUnknown1: %s,\n", indentationValues, pn.Unknown1))
	b.WriteString(fmt.Sprintf("%sUnknown2: %s,\n", indentationValues, pn.Unknown2))
	b.WriteString(fmt.Sprintf("%sUnknown3: %s,\n", indentationValues, pn.Unknown3))
	b.WriteString(fmt.Sprintf("%sUnknown4: %s,\n", indentationValues, pn.Unknown4))
	b.WriteString(fmt.Sprintf("%sUnknown5: %s,\n", indentationValues, pn.Unknown5))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewPersistentNotification returns a new PersistentNotification
func NewPersistentNotification() *PersistentNotification {
	pn := &PersistentNotification{
		Data:     types.NewData(),
		Unknown1: types.NewPrimitiveU64(0),
		Unknown2: types.NewPrimitiveU32(0),
		Unknown3: types.NewPrimitiveU32(0),
		Unknown4: types.NewPrimitiveU32(0),
		Unknown5: types.NewString(""),
	}

	return pn
}
