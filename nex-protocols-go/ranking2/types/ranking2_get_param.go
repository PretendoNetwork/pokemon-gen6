// Package types implements all the types used by the Ranking2 protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// Ranking2GetParam is a type within the Ranking2 protocol
type Ranking2GetParam struct {
	types.Structure
	NexUniqueID        *types.PrimitiveU64
	PrincipalID        *types.PID
	Category           *types.PrimitiveU32
	Offset             *types.PrimitiveU32
	Length             *types.PrimitiveU32
	SortFlags          *types.PrimitiveU32
	OptionFlags        *types.PrimitiveU32
	Mode               *types.PrimitiveU8
	NumSeasonsToGoBack *types.PrimitiveU8
}

// WriteTo writes the Ranking2GetParam to the given writable
func (rgp *Ranking2GetParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	rgp.NexUniqueID.WriteTo(contentWritable)
	rgp.PrincipalID.WriteTo(contentWritable)
	rgp.Category.WriteTo(contentWritable)
	rgp.Offset.WriteTo(contentWritable)
	rgp.Length.WriteTo(contentWritable)
	rgp.SortFlags.WriteTo(contentWritable)
	rgp.OptionFlags.WriteTo(contentWritable)
	rgp.Mode.WriteTo(contentWritable)
	rgp.NumSeasonsToGoBack.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	rgp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the Ranking2GetParam from the given readable
func (rgp *Ranking2GetParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = rgp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2GetParam header. %s", err.Error())
	}

	err = rgp.NexUniqueID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2GetParam.NexUniqueID. %s", err.Error())
	}

	err = rgp.PrincipalID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2GetParam.PrincipalID. %s", err.Error())
	}

	err = rgp.Category.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2GetParam.Category. %s", err.Error())
	}

	err = rgp.Offset.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2GetParam.Offset. %s", err.Error())
	}

	err = rgp.Length.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2GetParam.Length. %s", err.Error())
	}

	err = rgp.SortFlags.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2GetParam.SortFlags. %s", err.Error())
	}

	err = rgp.OptionFlags.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2GetParam.OptionFlags. %s", err.Error())
	}

	err = rgp.Mode.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2GetParam.Mode. %s", err.Error())
	}

	err = rgp.NumSeasonsToGoBack.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2GetParam.NumSeasonsToGoBack. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of Ranking2GetParam
func (rgp *Ranking2GetParam) Copy() types.RVType {
	copied := NewRanking2GetParam()

	copied.StructureVersion = rgp.StructureVersion
	copied.NexUniqueID = rgp.NexUniqueID.Copy().(*types.PrimitiveU64)
	copied.PrincipalID = rgp.PrincipalID.Copy().(*types.PID)
	copied.Category = rgp.Category.Copy().(*types.PrimitiveU32)
	copied.Offset = rgp.Offset.Copy().(*types.PrimitiveU32)
	copied.Length = rgp.Length.Copy().(*types.PrimitiveU32)
	copied.SortFlags = rgp.SortFlags.Copy().(*types.PrimitiveU32)
	copied.OptionFlags = rgp.OptionFlags.Copy().(*types.PrimitiveU32)
	copied.Mode = rgp.Mode.Copy().(*types.PrimitiveU8)
	copied.NumSeasonsToGoBack = rgp.NumSeasonsToGoBack.Copy().(*types.PrimitiveU8)

	return copied
}

// Equals checks if the given Ranking2GetParam contains the same data as the current Ranking2GetParam
func (rgp *Ranking2GetParam) Equals(o types.RVType) bool {
	if _, ok := o.(*Ranking2GetParam); !ok {
		return false
	}

	other := o.(*Ranking2GetParam)

	if rgp.StructureVersion != other.StructureVersion {
		return false
	}

	if !rgp.NexUniqueID.Equals(other.NexUniqueID) {
		return false
	}

	if !rgp.PrincipalID.Equals(other.PrincipalID) {
		return false
	}

	if !rgp.Category.Equals(other.Category) {
		return false
	}

	if !rgp.Offset.Equals(other.Offset) {
		return false
	}

	if !rgp.Length.Equals(other.Length) {
		return false
	}

	if !rgp.SortFlags.Equals(other.SortFlags) {
		return false
	}

	if !rgp.OptionFlags.Equals(other.OptionFlags) {
		return false
	}

	if !rgp.Mode.Equals(other.Mode) {
		return false
	}

	return rgp.NumSeasonsToGoBack.Equals(other.NumSeasonsToGoBack)
}

// String returns the string representation of the Ranking2GetParam
func (rgp *Ranking2GetParam) String() string {
	return rgp.FormatToString(0)
}

// FormatToString pretty-prints the Ranking2GetParam using the provided indentation level
func (rgp *Ranking2GetParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("Ranking2GetParam{\n")
	b.WriteString(fmt.Sprintf("%sNexUniqueID: %s,\n", indentationValues, rgp.NexUniqueID))
	b.WriteString(fmt.Sprintf("%sPrincipalID: %s,\n", indentationValues, rgp.PrincipalID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sCategory: %s,\n", indentationValues, rgp.Category))
	b.WriteString(fmt.Sprintf("%sOffset: %s,\n", indentationValues, rgp.Offset))
	b.WriteString(fmt.Sprintf("%sLength: %s,\n", indentationValues, rgp.Length))
	b.WriteString(fmt.Sprintf("%sSortFlags: %s,\n", indentationValues, rgp.SortFlags))
	b.WriteString(fmt.Sprintf("%sOptionFlags: %s,\n", indentationValues, rgp.OptionFlags))
	b.WriteString(fmt.Sprintf("%sMode: %s,\n", indentationValues, rgp.Mode))
	b.WriteString(fmt.Sprintf("%sNumSeasonsToGoBack: %s,\n", indentationValues, rgp.NumSeasonsToGoBack))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewRanking2GetParam returns a new Ranking2GetParam
func NewRanking2GetParam() *Ranking2GetParam {
	rgp := &Ranking2GetParam{
		NexUniqueID:        types.NewPrimitiveU64(0),
		PrincipalID:        types.NewPID(0),
		Category:           types.NewPrimitiveU32(0),
		Offset:             types.NewPrimitiveU32(0),
		Length:             types.NewPrimitiveU32(0),
		SortFlags:          types.NewPrimitiveU32(0),
		OptionFlags:        types.NewPrimitiveU32(0),
		Mode:               types.NewPrimitiveU8(0),
		NumSeasonsToGoBack: types.NewPrimitiveU8(0),
	}

	return rgp
}
