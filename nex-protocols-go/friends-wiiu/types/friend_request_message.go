// Package types implements all the types used by the FriendsWiiU protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// FriendRequestMessage is a type within the FriendsWiiU protocol
type FriendRequestMessage struct {
	types.Structure
	*types.Data
	FriendRequestID *types.PrimitiveU64
	Received        *types.PrimitiveBool
	Unknown2        *types.PrimitiveU8
	Message         *types.String
	Unknown3        *types.PrimitiveU8
	Unknown4        *types.String
	GameKey         *GameKey
	Unknown5        *types.DateTime
	ExpiresOn       *types.DateTime
}

// WriteTo writes the FriendRequestMessage to the given writable
func (frm *FriendRequestMessage) WriteTo(writable types.Writable) {
	frm.Data.WriteTo(writable)

	contentWritable := writable.CopyNew()

	frm.FriendRequestID.WriteTo(contentWritable)
	frm.Received.WriteTo(contentWritable)
	frm.Unknown2.WriteTo(contentWritable)
	frm.Message.WriteTo(contentWritable)
	frm.Unknown3.WriteTo(contentWritable)
	frm.Unknown4.WriteTo(contentWritable)
	frm.GameKey.WriteTo(contentWritable)
	frm.Unknown5.WriteTo(contentWritable)
	frm.ExpiresOn.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	frm.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the FriendRequestMessage from the given readable
func (frm *FriendRequestMessage) ExtractFrom(readable types.Readable) error {
	var err error

	err = frm.Data.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendRequestMessage.Data. %s", err.Error())
	}

	err = frm.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendRequestMessage header. %s", err.Error())
	}

	err = frm.FriendRequestID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendRequestMessage.FriendRequestID. %s", err.Error())
	}

	err = frm.Received.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendRequestMessage.Received. %s", err.Error())
	}

	err = frm.Unknown2.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendRequestMessage.Unknown2. %s", err.Error())
	}

	err = frm.Message.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendRequestMessage.Message. %s", err.Error())
	}

	err = frm.Unknown3.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendRequestMessage.Unknown3. %s", err.Error())
	}

	err = frm.Unknown4.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendRequestMessage.Unknown4. %s", err.Error())
	}

	err = frm.GameKey.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendRequestMessage.GameKey. %s", err.Error())
	}

	err = frm.Unknown5.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendRequestMessage.Unknown5. %s", err.Error())
	}

	err = frm.ExpiresOn.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendRequestMessage.ExpiresOn. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of FriendRequestMessage
func (frm *FriendRequestMessage) Copy() types.RVType {
	copied := NewFriendRequestMessage()

	copied.StructureVersion = frm.StructureVersion
	copied.Data = frm.Data.Copy().(*types.Data)
	copied.FriendRequestID = frm.FriendRequestID.Copy().(*types.PrimitiveU64)
	copied.Received = frm.Received.Copy().(*types.PrimitiveBool)
	copied.Unknown2 = frm.Unknown2.Copy().(*types.PrimitiveU8)
	copied.Message = frm.Message.Copy().(*types.String)
	copied.Unknown3 = frm.Unknown3.Copy().(*types.PrimitiveU8)
	copied.Unknown4 = frm.Unknown4.Copy().(*types.String)
	copied.GameKey = frm.GameKey.Copy().(*GameKey)
	copied.Unknown5 = frm.Unknown5.Copy().(*types.DateTime)
	copied.ExpiresOn = frm.ExpiresOn.Copy().(*types.DateTime)

	return copied
}

// Equals checks if the given FriendRequestMessage contains the same data as the current FriendRequestMessage
func (frm *FriendRequestMessage) Equals(o types.RVType) bool {
	if _, ok := o.(*FriendRequestMessage); !ok {
		return false
	}

	other := o.(*FriendRequestMessage)

	if frm.StructureVersion != other.StructureVersion {
		return false
	}

	if !frm.Data.Equals(other.Data) {
		return false
	}

	if !frm.FriendRequestID.Equals(other.FriendRequestID) {
		return false
	}

	if !frm.Received.Equals(other.Received) {
		return false
	}

	if !frm.Unknown2.Equals(other.Unknown2) {
		return false
	}

	if !frm.Message.Equals(other.Message) {
		return false
	}

	if !frm.Unknown3.Equals(other.Unknown3) {
		return false
	}

	if !frm.Unknown4.Equals(other.Unknown4) {
		return false
	}

	if !frm.GameKey.Equals(other.GameKey) {
		return false
	}

	if !frm.Unknown5.Equals(other.Unknown5) {
		return false
	}

	return frm.ExpiresOn.Equals(other.ExpiresOn)
}

// String returns the string representation of the FriendRequestMessage
func (frm *FriendRequestMessage) String() string {
	return frm.FormatToString(0)
}

// FormatToString pretty-prints the FriendRequestMessage using the provided indentation level
func (frm *FriendRequestMessage) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("FriendRequestMessage{\n")
	b.WriteString(fmt.Sprintf("%sData (parent): %s,\n", indentationValues, frm.Data.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sFriendRequestID: %s,\n", indentationValues, frm.FriendRequestID))
	b.WriteString(fmt.Sprintf("%sReceived: %s,\n", indentationValues, frm.Received))
	b.WriteString(fmt.Sprintf("%sUnknown2: %s,\n", indentationValues, frm.Unknown2))
	b.WriteString(fmt.Sprintf("%sMessage: %s,\n", indentationValues, frm.Message))
	b.WriteString(fmt.Sprintf("%sUnknown3: %s,\n", indentationValues, frm.Unknown3))
	b.WriteString(fmt.Sprintf("%sUnknown4: %s,\n", indentationValues, frm.Unknown4))
	b.WriteString(fmt.Sprintf("%sGameKey: %s,\n", indentationValues, frm.GameKey.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sUnknown5: %s,\n", indentationValues, frm.Unknown5.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sExpiresOn: %s,\n", indentationValues, frm.ExpiresOn.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewFriendRequestMessage returns a new FriendRequestMessage
func NewFriendRequestMessage() *FriendRequestMessage {
	frm := &FriendRequestMessage{
		Data:            types.NewData(),
		FriendRequestID: types.NewPrimitiveU64(0),
		Received:        types.NewPrimitiveBool(false),
		Unknown2:        types.NewPrimitiveU8(0),
		Message:         types.NewString(""),
		Unknown3:        types.NewPrimitiveU8(0),
		Unknown4:        types.NewString(""),
		GameKey:         NewGameKey(),
		Unknown5:        types.NewDateTime(0),
		ExpiresOn:       types.NewDateTime(0),
	}

	return frm
}
