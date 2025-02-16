// Package types implements all the types used by the FriendsWiiU protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// FriendRequest is a type within the FriendsWiiU protocol
type FriendRequest struct {
	types.Structure
	*types.Data
	PrincipalInfo *PrincipalBasicInfo
	Message       *FriendRequestMessage
	SentOn        *types.DateTime
}

// WriteTo writes the FriendRequest to the given writable
func (fr *FriendRequest) WriteTo(writable types.Writable) {
	fr.Data.WriteTo(writable)

	contentWritable := writable.CopyNew()

	fr.PrincipalInfo.WriteTo(contentWritable)
	fr.Message.WriteTo(contentWritable)
	fr.SentOn.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	fr.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the FriendRequest from the given readable
func (fr *FriendRequest) ExtractFrom(readable types.Readable) error {
	var err error

	err = fr.Data.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendRequest.Data. %s", err.Error())
	}

	err = fr.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendRequest header. %s", err.Error())
	}

	err = fr.PrincipalInfo.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendRequest.PrincipalInfo. %s", err.Error())
	}

	err = fr.Message.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendRequest.Message. %s", err.Error())
	}

	err = fr.SentOn.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendRequest.SentOn. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of FriendRequest
func (fr *FriendRequest) Copy() types.RVType {
	copied := NewFriendRequest()

	copied.StructureVersion = fr.StructureVersion
	copied.Data = fr.Data.Copy().(*types.Data)
	copied.PrincipalInfo = fr.PrincipalInfo.Copy().(*PrincipalBasicInfo)
	copied.Message = fr.Message.Copy().(*FriendRequestMessage)
	copied.SentOn = fr.SentOn.Copy().(*types.DateTime)

	return copied
}

// Equals checks if the given FriendRequest contains the same data as the current FriendRequest
func (fr *FriendRequest) Equals(o types.RVType) bool {
	if _, ok := o.(*FriendRequest); !ok {
		return false
	}

	other := o.(*FriendRequest)

	if fr.StructureVersion != other.StructureVersion {
		return false
	}

	if !fr.Data.Equals(other.Data) {
		return false
	}

	if !fr.PrincipalInfo.Equals(other.PrincipalInfo) {
		return false
	}

	if !fr.Message.Equals(other.Message) {
		return false
	}

	return fr.SentOn.Equals(other.SentOn)
}

// String returns the string representation of the FriendRequest
func (fr *FriendRequest) String() string {
	return fr.FormatToString(0)
}

// FormatToString pretty-prints the FriendRequest using the provided indentation level
func (fr *FriendRequest) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("FriendRequest{\n")
	b.WriteString(fmt.Sprintf("%sData (parent): %s,\n", indentationValues, fr.Data.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sPrincipalInfo: %s,\n", indentationValues, fr.PrincipalInfo.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sMessage: %s,\n", indentationValues, fr.Message.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sSentOn: %s,\n", indentationValues, fr.SentOn.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewFriendRequest returns a new FriendRequest
func NewFriendRequest() *FriendRequest {
	fr := &FriendRequest{
		Data:          types.NewData(),
		PrincipalInfo: NewPrincipalBasicInfo(),
		Message:       NewFriendRequestMessage(),
		SentOn:        types.NewDateTime(0),
	}

	return fr
}
