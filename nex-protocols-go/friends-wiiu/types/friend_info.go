// Package types implements all the types used by the FriendsWiiU protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// FriendInfo is a type within the FriendsWiiU protocol
type FriendInfo struct {
	types.Structure
	*types.Data
	NNAInfo      *NNAInfo
	Presence     *NintendoPresenceV2
	Status       *Comment
	BecameFriend *types.DateTime
	LastOnline   *types.DateTime
	Unknown      *types.PrimitiveU64
}

// WriteTo writes the FriendInfo to the given writable
func (fi *FriendInfo) WriteTo(writable types.Writable) {
	fi.Data.WriteTo(writable)

	contentWritable := writable.CopyNew()

	fi.NNAInfo.WriteTo(contentWritable)
	fi.Presence.WriteTo(contentWritable)
	fi.Status.WriteTo(contentWritable)
	fi.BecameFriend.WriteTo(contentWritable)
	fi.LastOnline.WriteTo(contentWritable)
	fi.Unknown.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	fi.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the FriendInfo from the given readable
func (fi *FriendInfo) ExtractFrom(readable types.Readable) error {
	var err error

	err = fi.Data.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendInfo.Data. %s", err.Error())
	}

	err = fi.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendInfo header. %s", err.Error())
	}

	err = fi.NNAInfo.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendInfo.NNAInfo. %s", err.Error())
	}

	err = fi.Presence.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendInfo.Presence. %s", err.Error())
	}

	err = fi.Status.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendInfo.Status. %s", err.Error())
	}

	err = fi.BecameFriend.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendInfo.BecameFriend. %s", err.Error())
	}

	err = fi.LastOnline.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendInfo.LastOnline. %s", err.Error())
	}

	err = fi.Unknown.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendInfo.Unknown. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of FriendInfo
func (fi *FriendInfo) Copy() types.RVType {
	copied := NewFriendInfo()

	copied.StructureVersion = fi.StructureVersion
	copied.Data = fi.Data.Copy().(*types.Data)
	copied.NNAInfo = fi.NNAInfo.Copy().(*NNAInfo)
	copied.Presence = fi.Presence.Copy().(*NintendoPresenceV2)
	copied.Status = fi.Status.Copy().(*Comment)
	copied.BecameFriend = fi.BecameFriend.Copy().(*types.DateTime)
	copied.LastOnline = fi.LastOnline.Copy().(*types.DateTime)
	copied.Unknown = fi.Unknown.Copy().(*types.PrimitiveU64)

	return copied
}

// Equals checks if the given FriendInfo contains the same data as the current FriendInfo
func (fi *FriendInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*FriendInfo); !ok {
		return false
	}

	other := o.(*FriendInfo)

	if fi.StructureVersion != other.StructureVersion {
		return false
	}

	if !fi.Data.Equals(other.Data) {
		return false
	}

	if !fi.NNAInfo.Equals(other.NNAInfo) {
		return false
	}

	if !fi.Presence.Equals(other.Presence) {
		return false
	}

	if !fi.Status.Equals(other.Status) {
		return false
	}

	if !fi.BecameFriend.Equals(other.BecameFriend) {
		return false
	}

	if !fi.LastOnline.Equals(other.LastOnline) {
		return false
	}

	return fi.Unknown.Equals(other.Unknown)
}

// String returns the string representation of the FriendInfo
func (fi *FriendInfo) String() string {
	return fi.FormatToString(0)
}

// FormatToString pretty-prints the FriendInfo using the provided indentation level
func (fi *FriendInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("FriendInfo{\n")
	b.WriteString(fmt.Sprintf("%sData (parent): %s,\n", indentationValues, fi.Data.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sNNAInfo: %s,\n", indentationValues, fi.NNAInfo.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sPresence: %s,\n", indentationValues, fi.Presence.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sStatus: %s,\n", indentationValues, fi.Status.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sBecameFriend: %s,\n", indentationValues, fi.BecameFriend.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sLastOnline: %s,\n", indentationValues, fi.LastOnline.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sUnknown: %s,\n", indentationValues, fi.Unknown))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewFriendInfo returns a new FriendInfo
func NewFriendInfo() *FriendInfo {
	fi := &FriendInfo{
		Data:         types.NewData(),
		NNAInfo:      NewNNAInfo(),
		Presence:     NewNintendoPresenceV2(),
		Status:       NewComment(),
		BecameFriend: types.NewDateTime(0),
		LastOnline:   types.NewDateTime(0),
		Unknown:      types.NewPrimitiveU64(0),
	}

	return fi
}
