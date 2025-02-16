// Package types implements all the types used by the Shop protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// SubscriberContent is a type within the Shop protocol
type SubscriberContent struct {
	types.Structure
	Unknown1 *types.PrimitiveU64
	Unknown2 *types.String
	Unknown3 *types.Buffer
	Unknown4 *types.PrimitiveU64
	Unknown5 *types.List[*types.String]
	Unknown6 *types.DateTime
}

// WriteTo writes the SubscriberContent to the given writable
func (sc *SubscriberContent) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	sc.Unknown1.WriteTo(contentWritable)
	sc.Unknown2.WriteTo(contentWritable)
	sc.Unknown3.WriteTo(contentWritable)
	sc.Unknown4.WriteTo(contentWritable)
	sc.Unknown5.WriteTo(contentWritable)
	sc.Unknown6.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	sc.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the SubscriberContent from the given readable
func (sc *SubscriberContent) ExtractFrom(readable types.Readable) error {
	var err error

	err = sc.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SubscriberContent header. %s", err.Error())
	}

	err = sc.Unknown1.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SubscriberContent.Unknown1. %s", err.Error())
	}

	err = sc.Unknown2.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SubscriberContent.Unknown2. %s", err.Error())
	}

	err = sc.Unknown3.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SubscriberContent.Unknown3. %s", err.Error())
	}

	err = sc.Unknown4.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SubscriberContent.Unknown4. %s", err.Error())
	}

	err = sc.Unknown5.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SubscriberContent.Unknown5. %s", err.Error())
	}

	err = sc.Unknown6.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SubscriberContent.Unknown6. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of SubscriberContent
func (sc *SubscriberContent) Copy() types.RVType {
	copied := NewSubscriberContent()

	copied.StructureVersion = sc.StructureVersion
	copied.Unknown1 = sc.Unknown1.Copy().(*types.PrimitiveU64)
	copied.Unknown2 = sc.Unknown2.Copy().(*types.String)
	copied.Unknown3 = sc.Unknown3.Copy().(*types.Buffer)
	copied.Unknown4 = sc.Unknown4.Copy().(*types.PrimitiveU64)
	copied.Unknown5 = sc.Unknown5.Copy().(*types.List[*types.String])
	copied.Unknown6 = sc.Unknown6.Copy().(*types.DateTime)

	return copied
}

// Equals checks if the given SubscriberContent contains the same data as the current SubscriberContent
func (sc *SubscriberContent) Equals(o types.RVType) bool {
	if _, ok := o.(*SubscriberContent); !ok {
		return false
	}

	other := o.(*SubscriberContent)

	if sc.StructureVersion != other.StructureVersion {
		return false
	}

	if !sc.Unknown1.Equals(other.Unknown1) {
		return false
	}

	if !sc.Unknown2.Equals(other.Unknown2) {
		return false
	}

	if !sc.Unknown3.Equals(other.Unknown3) {
		return false
	}

	if !sc.Unknown4.Equals(other.Unknown4) {
		return false
	}

	if !sc.Unknown5.Equals(other.Unknown5) {
		return false
	}

	return sc.Unknown6.Equals(other.Unknown6)
}

// String returns the string representation of the SubscriberContent
func (sc *SubscriberContent) String() string {
	return sc.FormatToString(0)
}

// FormatToString pretty-prints the SubscriberContent using the provided indentation level
func (sc *SubscriberContent) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("SubscriberContent{\n")
	b.WriteString(fmt.Sprintf("%sUnknown1: %s,\n", indentationValues, sc.Unknown1))
	b.WriteString(fmt.Sprintf("%sUnknown2: %s,\n", indentationValues, sc.Unknown2))
	b.WriteString(fmt.Sprintf("%sUnknown3: %s,\n", indentationValues, sc.Unknown3))
	b.WriteString(fmt.Sprintf("%sUnknown4: %s,\n", indentationValues, sc.Unknown4))
	b.WriteString(fmt.Sprintf("%sUnknown5: %s,\n", indentationValues, sc.Unknown5))
	b.WriteString(fmt.Sprintf("%sUnknown6: %s,\n", indentationValues, sc.Unknown6.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewSubscriberContent returns a new SubscriberContent
func NewSubscriberContent() *SubscriberContent {
	sc := &SubscriberContent{
		Unknown1: types.NewPrimitiveU64(0),
		Unknown2: types.NewString(""),
		Unknown3: types.NewBuffer(nil),
		Unknown4: types.NewPrimitiveU64(0),
		Unknown5: types.NewList[*types.String](),
		Unknown6: types.NewDateTime(0),
	}

	sc.Unknown5.Type = types.NewString("")

	return sc
}
