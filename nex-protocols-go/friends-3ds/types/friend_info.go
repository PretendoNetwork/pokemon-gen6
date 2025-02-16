// Package types implements all the types used by the Friends3DS protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// FriendInfo is a type within the Friends3DS protocol
type FriendInfo struct {
	types.Structure
	PID     *types.PID
	Unknown *types.DateTime
}

// WriteTo writes the FriendInfo to the given writable
func (fi *FriendInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	fi.PID.WriteTo(contentWritable)
	fi.Unknown.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	fi.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the FriendInfo from the given readable
func (fi *FriendInfo) ExtractFrom(readable types.Readable) error {
	var err error

	err = fi.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendInfo header. %s", err.Error())
	}

	err = fi.PID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendInfo.PID. %s", err.Error())
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
	copied.PID = fi.PID.Copy().(*types.PID)
	copied.Unknown = fi.Unknown.Copy().(*types.DateTime)

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

	if !fi.PID.Equals(other.PID) {
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
	b.WriteString(fmt.Sprintf("%sPID: %s,\n", indentationValues, fi.PID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sUnknown: %s,\n", indentationValues, fi.Unknown.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewFriendInfo returns a new FriendInfo
func NewFriendInfo() *FriendInfo {
	fi := &FriendInfo{
		PID:     types.NewPID(0),
		Unknown: types.NewDateTime(0),
	}

	return fi
}
