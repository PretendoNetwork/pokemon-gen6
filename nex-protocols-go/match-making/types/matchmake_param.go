// Package types implements all the types used by the Matchmaking protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// MatchmakeParam is a type within the Matchmaking protocol
type MatchmakeParam struct {
	types.Structure
	Params *types.Map[*types.String, *types.Variant]
}

// WriteTo writes the MatchmakeParam to the given writable
func (mp *MatchmakeParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	mp.Params.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	mp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the MatchmakeParam from the given readable
func (mp *MatchmakeParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = mp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeParam header. %s", err.Error())
	}

	err = mp.Params.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeParam.Params. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of MatchmakeParam
func (mp *MatchmakeParam) Copy() types.RVType {
	copied := NewMatchmakeParam()

	copied.StructureVersion = mp.StructureVersion
	copied.Params = mp.Params.Copy().(*types.Map[*types.String, *types.Variant])

	return copied
}

// Equals checks if the given MatchmakeParam contains the same data as the current MatchmakeParam
func (mp *MatchmakeParam) Equals(o types.RVType) bool {
	if _, ok := o.(*MatchmakeParam); !ok {
		return false
	}

	other := o.(*MatchmakeParam)

	if mp.StructureVersion != other.StructureVersion {
		return false
	}

	return mp.Params.Equals(other.Params)
}

// String returns the string representation of the MatchmakeParam
func (mp *MatchmakeParam) String() string {
	return mp.FormatToString(0)
}

// FormatToString pretty-prints the MatchmakeParam using the provided indentation level
func (mp *MatchmakeParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("MatchmakeParam{\n")
	b.WriteString(fmt.Sprintf("%sParams: %s,\n", indentationValues, mp.Params.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewMatchmakeParam returns a new MatchmakeParam
func NewMatchmakeParam() *MatchmakeParam {
	mp := &MatchmakeParam{
		Params: types.NewMap[*types.String, *types.Variant](),
	}

	mp.Params.KeyType = types.NewString("")
	mp.Params.ValueType = types.NewVariant()

	return mp
}
