// Package types implements all the types used by the DataStoreMiitopia protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// MiiTubeSearchResult is a type within the DataStoreMiitopia protocol
type MiiTubeSearchResult struct {
	types.Structure

	Result  *types.List[*MiiTubeMiiInfo]
	Count   *types.PrimitiveU32
	Page    *types.PrimitiveU32
	HasNext *types.PrimitiveBool
}

// WriteTo writes the MiiTubeSearchResult to the given variable
func (mtsr *MiiTubeSearchResult) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	mtsr.Result.WriteTo(contentWritable)
	mtsr.Count.WriteTo(contentWritable)
	mtsr.Page.WriteTo(contentWritable)
	mtsr.HasNext.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	mtsr.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the MiiTubeSearchResult from the given readable
func (mtsr *MiiTubeSearchResult) ExtractFrom(readable types.Readable) error {
	var err error

	err = mtsr.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MiiTubeSearchResult header. %s", err.Error())
	}

	err = mtsr.Result.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MiiTubeSearchResult.Result. %s", err.Error())
	}

	err = mtsr.Count.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MiiTubeSearchResult.Count. %s", err.Error())
	}

	err = mtsr.Page.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MiiTubeSearchResult.Page. %s", err.Error())
	}

	err = mtsr.HasNext.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MiiTubeSearchResult.ExtractNext. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of MiiTubeSearchResult
func (mtsr *MiiTubeSearchResult) Copy() types.RVType {
	copied := NewMiiTubeSearchResult()

	copied.Result = mtsr.Result
	copied.Count = mtsr.Count
	copied.Page = mtsr.Page
	copied.HasNext = mtsr.HasNext

	return copied
}

// Equals checks if the given MiiTubeSearchResult contains the same data as the current MiiTubeSearchResult
func (mtsr *MiiTubeSearchResult) Equals(o types.RVType) bool {
	if _, ok := o.(*MiiTubeSearchResult); !ok {
		return false
	}

	other := o.(*MiiTubeSearchResult)

	if !mtsr.Result.Equals(other.Result) {
		return false
	}

	if !mtsr.Count.Equals(other.Count) {
		return false
	}

	if !mtsr.Page.Equals(other.Page) {
		return false
	}

	return mtsr.HasNext.Equals(other.HasNext)
}

// String returns the string representation of the MiiTubeSearchResult
func (mtsr *MiiTubeSearchResult) String() string {
	return mtsr.FormatToString(0)
}

// FormatToString pretty-prints the MiiTubeSearchResult using the provided indentation level
func (mtsr *MiiTubeSearchResult) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("MiiTubeSearchResult{\n")
	b.WriteString(fmt.Sprintf("%sResult: %s,\n", indentationValues, mtsr.Result))
	b.WriteString(fmt.Sprintf("%sCount: %s,\n", indentationValues, mtsr.Count))
	b.WriteString(fmt.Sprintf("%sPage: %s,\n", indentationValues, mtsr.Page))
	b.WriteString(fmt.Sprintf("%sHasNext: %s,\n", indentationValues, mtsr.HasNext))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewMiiTubeSearchResult returns a new MiiTubeSearchResult
func NewMiiTubeSearchResult() *MiiTubeSearchResult {
	mtsr := &MiiTubeSearchResult{
		Result:  types.NewList[*MiiTubeMiiInfo](),
		Count:   types.NewPrimitiveU32(0),
		Page:    types.NewPrimitiveU32(0),
		HasNext: types.NewPrimitiveBool(false),
	}

	mtsr.Result.Type = NewMiiTubeMiiInfo()

	return mtsr
}
