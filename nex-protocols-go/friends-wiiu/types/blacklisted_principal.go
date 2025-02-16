// Package types implements all the types used by the FriendsWiiU protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// BlacklistedPrincipal is a type within the FriendsWiiU protocol
type BlacklistedPrincipal struct {
	types.Structure
	*types.Data
	PrincipalBasicInfo *PrincipalBasicInfo
	GameKey            *GameKey
	BlackListedSince   *types.DateTime
}

// WriteTo writes the BlacklistedPrincipal to the given writable
func (bp *BlacklistedPrincipal) WriteTo(writable types.Writable) {
	bp.Data.WriteTo(writable)

	contentWritable := writable.CopyNew()

	bp.PrincipalBasicInfo.WriteTo(contentWritable)
	bp.GameKey.WriteTo(contentWritable)
	bp.BlackListedSince.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	bp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the BlacklistedPrincipal from the given readable
func (bp *BlacklistedPrincipal) ExtractFrom(readable types.Readable) error {
	var err error

	err = bp.Data.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract BlacklistedPrincipal.Data. %s", err.Error())
	}

	err = bp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract BlacklistedPrincipal header. %s", err.Error())
	}

	err = bp.PrincipalBasicInfo.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract BlacklistedPrincipal.PrincipalBasicInfo. %s", err.Error())
	}

	err = bp.GameKey.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract BlacklistedPrincipal.GameKey. %s", err.Error())
	}

	err = bp.BlackListedSince.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract BlacklistedPrincipal.BlackListedSince. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of BlacklistedPrincipal
func (bp *BlacklistedPrincipal) Copy() types.RVType {
	copied := NewBlacklistedPrincipal()

	copied.StructureVersion = bp.StructureVersion
	copied.Data = bp.Data.Copy().(*types.Data)
	copied.PrincipalBasicInfo = bp.PrincipalBasicInfo.Copy().(*PrincipalBasicInfo)
	copied.GameKey = bp.GameKey.Copy().(*GameKey)
	copied.BlackListedSince = bp.BlackListedSince.Copy().(*types.DateTime)

	return copied
}

// Equals checks if the given BlacklistedPrincipal contains the same data as the current BlacklistedPrincipal
func (bp *BlacklistedPrincipal) Equals(o types.RVType) bool {
	if _, ok := o.(*BlacklistedPrincipal); !ok {
		return false
	}

	other := o.(*BlacklistedPrincipal)

	if bp.StructureVersion != other.StructureVersion {
		return false
	}

	if !bp.Data.Equals(other.Data) {
		return false
	}

	if !bp.PrincipalBasicInfo.Equals(other.PrincipalBasicInfo) {
		return false
	}

	if !bp.GameKey.Equals(other.GameKey) {
		return false
	}

	return bp.BlackListedSince.Equals(other.BlackListedSince)
}

// String returns the string representation of the BlacklistedPrincipal
func (bp *BlacklistedPrincipal) String() string {
	return bp.FormatToString(0)
}

// FormatToString pretty-prints the BlacklistedPrincipal using the provided indentation level
func (bp *BlacklistedPrincipal) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("BlacklistedPrincipal{\n")
	b.WriteString(fmt.Sprintf("%sData (parent): %s,\n", indentationValues, bp.Data.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sPrincipalBasicInfo: %s,\n", indentationValues, bp.PrincipalBasicInfo.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sGameKey: %s,\n", indentationValues, bp.GameKey.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sBlackListedSince: %s,\n", indentationValues, bp.BlackListedSince.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewBlacklistedPrincipal returns a new BlacklistedPrincipal
func NewBlacklistedPrincipal() *BlacklistedPrincipal {
	bp := &BlacklistedPrincipal{
		Data:               types.NewData(),
		PrincipalBasicInfo: NewPrincipalBasicInfo(),
		GameKey:            NewGameKey(),
		BlackListedSince:   types.NewDateTime(0),
	}

	return bp
}
