// Package types implements all the types used by the MatchmakeExtension protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// FriendUserInfo is a type within the MatchmakeExtension protocol
type FriendUserInfo struct {
	types.Structure
	PID      *types.PID
	Name     *types.String
	Presence *types.PrimitiveU32
}

// WriteTo writes the FriendUserInfo to the given writable
func (fui *FriendUserInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	fui.PID.WriteTo(contentWritable)
	fui.Name.WriteTo(contentWritable)
	fui.Presence.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	fui.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the FriendUserInfo from the given readable
func (fui *FriendUserInfo) ExtractFrom(readable types.Readable) error {
	var err error

	err = fui.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendUserInfo header. %s", err.Error())
	}

	err = fui.PID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendUserInfo.PID. %s", err.Error())
	}

	err = fui.Name.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendUserInfo.Name. %s", err.Error())
	}

	err = fui.Presence.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendUserInfo.Presence. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of FriendUserInfo
func (fui *FriendUserInfo) Copy() types.RVType {
	copied := NewFriendUserInfo()

	copied.StructureVersion = fui.StructureVersion
	copied.PID = fui.PID.Copy().(*types.PID)
	copied.Name = fui.Name.Copy().(*types.String)
	copied.Presence = fui.Presence.Copy().(*types.PrimitiveU32)

	return copied
}

// Equals checks if the given FriendUserInfo contains the same data as the current FriendUserInfo
func (fui *FriendUserInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*FriendUserInfo); !ok {
		return false
	}

	other := o.(*FriendUserInfo)

	if fui.StructureVersion != other.StructureVersion {
		return false
	}

	if !fui.PID.Equals(other.PID) {
		return false
	}

	if !fui.Name.Equals(other.Name) {
		return false
	}

	return fui.Presence.Equals(other.Presence)
}

// String returns the string representation of the FriendUserInfo
func (fui *FriendUserInfo) String() string {
	return fui.FormatToString(0)
}

// FormatToString pretty-prints the FriendUserInfo using the provided indentation level
func (fui *FriendUserInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("FriendUserInfo{\n")
	b.WriteString(fmt.Sprintf("%sPID: %s,\n", indentationValues, fui.PID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sName: %s,\n", indentationValues, fui.Name))
	b.WriteString(fmt.Sprintf("%sPresence: %s,\n", indentationValues, fui.Presence))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewFriendUserInfo returns a new FriendUserInfo
func NewFriendUserInfo() *FriendUserInfo {
	fui := &FriendUserInfo{
		PID:      types.NewPID(0),
		Name:     types.NewString(""),
		Presence: types.NewPrimitiveU32(0),
	}

	return fui
}
