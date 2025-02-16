// Package types implements all the types used by the Matchmaking protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// MatchmakeBlockListParam is a type within the Matchmaking protocol
type MatchmakeBlockListParam struct {
	types.Structure
	OptionFlag *types.PrimitiveU32
}

// WriteTo writes the MatchmakeBlockListParam to the given writable
func (mblp *MatchmakeBlockListParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	mblp.OptionFlag.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	mblp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the MatchmakeBlockListParam from the given readable
func (mblp *MatchmakeBlockListParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = mblp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeBlockListParam header. %s", err.Error())
	}

	err = mblp.OptionFlag.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeBlockListParam.OptionFlag. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of MatchmakeBlockListParam
func (mblp *MatchmakeBlockListParam) Copy() types.RVType {
	copied := NewMatchmakeBlockListParam()

	copied.StructureVersion = mblp.StructureVersion
	copied.OptionFlag = mblp.OptionFlag.Copy().(*types.PrimitiveU32)

	return copied
}

// Equals checks if the given MatchmakeBlockListParam contains the same data as the current MatchmakeBlockListParam
func (mblp *MatchmakeBlockListParam) Equals(o types.RVType) bool {
	if _, ok := o.(*MatchmakeBlockListParam); !ok {
		return false
	}

	other := o.(*MatchmakeBlockListParam)

	if mblp.StructureVersion != other.StructureVersion {
		return false
	}

	return mblp.OptionFlag.Equals(other.OptionFlag)
}

// String returns the string representation of the MatchmakeBlockListParam
func (mblp *MatchmakeBlockListParam) String() string {
	return mblp.FormatToString(0)
}

// FormatToString pretty-prints the MatchmakeBlockListParam using the provided indentation level
func (mblp *MatchmakeBlockListParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("MatchmakeBlockListParam{\n")
	b.WriteString(fmt.Sprintf("%sOptionFlag: %s,\n", indentationValues, mblp.OptionFlag))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewMatchmakeBlockListParam returns a new MatchmakeBlockListParam
func NewMatchmakeBlockListParam() *MatchmakeBlockListParam {
	mblp := &MatchmakeBlockListParam{
		OptionFlag: types.NewPrimitiveU32(0),
	}

	return mblp
}
