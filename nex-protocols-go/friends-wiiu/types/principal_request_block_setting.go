// Package types implements all the types used by the FriendsWiiU protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// PrincipalRequestBlockSetting is a type within the FriendsWiiU protocol
type PrincipalRequestBlockSetting struct {
	types.Structure
	*types.Data
	PID       *types.PrimitiveU32
	IsBlocked *types.PrimitiveBool
}

// WriteTo writes the PrincipalRequestBlockSetting to the given writable
func (prbs *PrincipalRequestBlockSetting) WriteTo(writable types.Writable) {
	prbs.Data.WriteTo(writable)

	contentWritable := writable.CopyNew()

	prbs.PID.WriteTo(contentWritable)
	prbs.IsBlocked.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	prbs.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the PrincipalRequestBlockSetting from the given readable
func (prbs *PrincipalRequestBlockSetting) ExtractFrom(readable types.Readable) error {
	var err error

	err = prbs.Data.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract PrincipalRequestBlockSetting.Data. %s", err.Error())
	}

	err = prbs.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract PrincipalRequestBlockSetting header. %s", err.Error())
	}

	err = prbs.PID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract PrincipalRequestBlockSetting.PID. %s", err.Error())
	}

	err = prbs.IsBlocked.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract PrincipalRequestBlockSetting.IsBlocked. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of PrincipalRequestBlockSetting
func (prbs *PrincipalRequestBlockSetting) Copy() types.RVType {
	copied := NewPrincipalRequestBlockSetting()

	copied.StructureVersion = prbs.StructureVersion
	copied.Data = prbs.Data.Copy().(*types.Data)
	copied.PID = prbs.PID.Copy().(*types.PrimitiveU32)
	copied.IsBlocked = prbs.IsBlocked.Copy().(*types.PrimitiveBool)

	return copied
}

// Equals checks if the given PrincipalRequestBlockSetting contains the same data as the current PrincipalRequestBlockSetting
func (prbs *PrincipalRequestBlockSetting) Equals(o types.RVType) bool {
	if _, ok := o.(*PrincipalRequestBlockSetting); !ok {
		return false
	}

	other := o.(*PrincipalRequestBlockSetting)

	if prbs.StructureVersion != other.StructureVersion {
		return false
	}

	if !prbs.Data.Equals(other.Data) {
		return false
	}

	if !prbs.PID.Equals(other.PID) {
		return false
	}

	return prbs.IsBlocked.Equals(other.IsBlocked)
}

// String returns the string representation of the PrincipalRequestBlockSetting
func (prbs *PrincipalRequestBlockSetting) String() string {
	return prbs.FormatToString(0)
}

// FormatToString pretty-prints the PrincipalRequestBlockSetting using the provided indentation level
func (prbs *PrincipalRequestBlockSetting) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("PrincipalRequestBlockSetting{\n")
	b.WriteString(fmt.Sprintf("%sData (parent): %s,\n", indentationValues, prbs.Data.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sPID: %s,\n", indentationValues, prbs.PID))
	b.WriteString(fmt.Sprintf("%sIsBlocked: %s,\n", indentationValues, prbs.IsBlocked))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewPrincipalRequestBlockSetting returns a new PrincipalRequestBlockSetting
func NewPrincipalRequestBlockSetting() *PrincipalRequestBlockSetting {
	prbs := &PrincipalRequestBlockSetting{
		Data:      types.NewData(),
		PID:       types.NewPrimitiveU32(0),
		IsBlocked: types.NewPrimitiveBool(false),
	}

	return prbs
}
