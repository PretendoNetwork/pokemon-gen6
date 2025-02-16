// Package types implements all the types used by the FriendsWiiU protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// PrincipalBasicInfo is a type within the FriendsWiiU protocol
type PrincipalBasicInfo struct {
	types.Structure
	*types.Data
	PID     *types.PID
	NNID    *types.String
	Mii     *MiiV2
	Unknown *types.PrimitiveU8
}

// WriteTo writes the PrincipalBasicInfo to the given writable
func (pbi *PrincipalBasicInfo) WriteTo(writable types.Writable) {
	pbi.Data.WriteTo(writable)

	contentWritable := writable.CopyNew()

	pbi.PID.WriteTo(contentWritable)
	pbi.NNID.WriteTo(contentWritable)
	pbi.Mii.WriteTo(contentWritable)
	pbi.Unknown.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	pbi.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the PrincipalBasicInfo from the given readable
func (pbi *PrincipalBasicInfo) ExtractFrom(readable types.Readable) error {
	var err error

	err = pbi.Data.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract PrincipalBasicInfo.Data. %s", err.Error())
	}

	err = pbi.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract PrincipalBasicInfo header. %s", err.Error())
	}

	err = pbi.PID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract PrincipalBasicInfo.PID. %s", err.Error())
	}

	err = pbi.NNID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract PrincipalBasicInfo.NNID. %s", err.Error())
	}

	err = pbi.Mii.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract PrincipalBasicInfo.Mii. %s", err.Error())
	}

	err = pbi.Unknown.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract PrincipalBasicInfo.Unknown. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of PrincipalBasicInfo
func (pbi *PrincipalBasicInfo) Copy() types.RVType {
	copied := NewPrincipalBasicInfo()

	copied.StructureVersion = pbi.StructureVersion
	copied.Data = pbi.Data.Copy().(*types.Data)
	copied.PID = pbi.PID.Copy().(*types.PID)
	copied.NNID = pbi.NNID.Copy().(*types.String)
	copied.Mii = pbi.Mii.Copy().(*MiiV2)
	copied.Unknown = pbi.Unknown.Copy().(*types.PrimitiveU8)

	return copied
}

// Equals checks if the given PrincipalBasicInfo contains the same data as the current PrincipalBasicInfo
func (pbi *PrincipalBasicInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*PrincipalBasicInfo); !ok {
		return false
	}

	other := o.(*PrincipalBasicInfo)

	if pbi.StructureVersion != other.StructureVersion {
		return false
	}

	if !pbi.Data.Equals(other.Data) {
		return false
	}

	if !pbi.PID.Equals(other.PID) {
		return false
	}

	if !pbi.NNID.Equals(other.NNID) {
		return false
	}

	if !pbi.Mii.Equals(other.Mii) {
		return false
	}

	return pbi.Unknown.Equals(other.Unknown)
}

// String returns the string representation of the PrincipalBasicInfo
func (pbi *PrincipalBasicInfo) String() string {
	return pbi.FormatToString(0)
}

// FormatToString pretty-prints the PrincipalBasicInfo using the provided indentation level
func (pbi *PrincipalBasicInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("PrincipalBasicInfo{\n")
	b.WriteString(fmt.Sprintf("%sData (parent): %s,\n", indentationValues, pbi.Data.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sPID: %s,\n", indentationValues, pbi.PID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sNNID: %s,\n", indentationValues, pbi.NNID))
	b.WriteString(fmt.Sprintf("%sMii: %s,\n", indentationValues, pbi.Mii.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sUnknown: %s,\n", indentationValues, pbi.Unknown))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewPrincipalBasicInfo returns a new PrincipalBasicInfo
func NewPrincipalBasicInfo() *PrincipalBasicInfo {
	pbi := &PrincipalBasicInfo{
		Data:    types.NewData(),
		PID:     types.NewPID(0),
		NNID:    types.NewString(""),
		Mii:     NewMiiV2(),
		Unknown: types.NewPrimitiveU8(0),
	}

	return pbi
}
