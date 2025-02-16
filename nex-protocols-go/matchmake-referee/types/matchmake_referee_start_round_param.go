// Package types implements all the types used by the MatchmakeReferee protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// MatchmakeRefereeStartRoundParam is a type within the MatchmakeReferee protocol
type MatchmakeRefereeStartRoundParam struct {
	types.Structure
	*types.Data
	PersonalDataCategory *types.PrimitiveU32
	GID                  *types.PrimitiveU32
	PIDs                 *types.List[*types.PID]
}

// WriteTo writes the MatchmakeRefereeStartRoundParam to the given writable
func (mrsrp *MatchmakeRefereeStartRoundParam) WriteTo(writable types.Writable) {
	mrsrp.Data.WriteTo(writable)

	contentWritable := writable.CopyNew()

	mrsrp.PersonalDataCategory.WriteTo(contentWritable)
	mrsrp.GID.WriteTo(contentWritable)
	mrsrp.PIDs.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	mrsrp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the MatchmakeRefereeStartRoundParam from the given readable
func (mrsrp *MatchmakeRefereeStartRoundParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = mrsrp.Data.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStartRoundParam.Data. %s", err.Error())
	}

	err = mrsrp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStartRoundParam header. %s", err.Error())
	}

	err = mrsrp.PersonalDataCategory.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStartRoundParam.PersonalDataCategory. %s", err.Error())
	}

	err = mrsrp.GID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStartRoundParam.GID. %s", err.Error())
	}

	err = mrsrp.PIDs.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStartRoundParam.PIDs. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of MatchmakeRefereeStartRoundParam
func (mrsrp *MatchmakeRefereeStartRoundParam) Copy() types.RVType {
	copied := NewMatchmakeRefereeStartRoundParam()

	copied.StructureVersion = mrsrp.StructureVersion
	copied.Data = mrsrp.Data.Copy().(*types.Data)
	copied.PersonalDataCategory = mrsrp.PersonalDataCategory.Copy().(*types.PrimitiveU32)
	copied.GID = mrsrp.GID.Copy().(*types.PrimitiveU32)
	copied.PIDs = mrsrp.PIDs.Copy().(*types.List[*types.PID])

	return copied
}

// Equals checks if the given MatchmakeRefereeStartRoundParam contains the same data as the current MatchmakeRefereeStartRoundParam
func (mrsrp *MatchmakeRefereeStartRoundParam) Equals(o types.RVType) bool {
	if _, ok := o.(*MatchmakeRefereeStartRoundParam); !ok {
		return false
	}

	other := o.(*MatchmakeRefereeStartRoundParam)

	if mrsrp.StructureVersion != other.StructureVersion {
		return false
	}

	if !mrsrp.Data.Equals(other.Data) {
		return false
	}

	if !mrsrp.PersonalDataCategory.Equals(other.PersonalDataCategory) {
		return false
	}

	if !mrsrp.GID.Equals(other.GID) {
		return false
	}

	return mrsrp.PIDs.Equals(other.PIDs)
}

// String returns the string representation of the MatchmakeRefereeStartRoundParam
func (mrsrp *MatchmakeRefereeStartRoundParam) String() string {
	return mrsrp.FormatToString(0)
}

// FormatToString pretty-prints the MatchmakeRefereeStartRoundParam using the provided indentation level
func (mrsrp *MatchmakeRefereeStartRoundParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("MatchmakeRefereeStartRoundParam{\n")
	b.WriteString(fmt.Sprintf("%sData (parent): %s,\n", indentationValues, mrsrp.Data.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sPersonalDataCategory: %s,\n", indentationValues, mrsrp.PersonalDataCategory))
	b.WriteString(fmt.Sprintf("%sGID: %s,\n", indentationValues, mrsrp.GID))
	b.WriteString(fmt.Sprintf("%sPIDs: %s,\n", indentationValues, mrsrp.PIDs))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewMatchmakeRefereeStartRoundParam returns a new MatchmakeRefereeStartRoundParam
func NewMatchmakeRefereeStartRoundParam() *MatchmakeRefereeStartRoundParam {
	mrsrp := &MatchmakeRefereeStartRoundParam{
		Data:                 types.NewData(),
		PersonalDataCategory: types.NewPrimitiveU32(0),
		GID:                  types.NewPrimitiveU32(0),
		PIDs:                 types.NewList[*types.PID](),
	}

	mrsrp.PIDs.Type = types.NewPID(0)

	return mrsrp
}
