// Package types implements all the types used by the Friends3DS protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// FriendPersistentInfo is a type within the Friends3DS protocol
type FriendPersistentInfo struct {
	types.Structure
	*types.Data
	PID              *types.PID
	Region           *types.PrimitiveU8
	Country          *types.PrimitiveU8
	Area             *types.PrimitiveU8
	Language         *types.PrimitiveU8
	Platform         *types.PrimitiveU8
	GameKey          *GameKey
	Message          *types.String
	MessageUpdatedAt *types.DateTime
	MiiModifiedAt    *types.DateTime
	LastOnline       *types.DateTime
}

// WriteTo writes the FriendPersistentInfo to the given writable
func (fpi *FriendPersistentInfo) WriteTo(writable types.Writable) {
	fpi.Data.WriteTo(writable)

	contentWritable := writable.CopyNew()

	fpi.PID.WriteTo(contentWritable)
	fpi.Region.WriteTo(contentWritable)
	fpi.Country.WriteTo(contentWritable)
	fpi.Area.WriteTo(contentWritable)
	fpi.Language.WriteTo(contentWritable)
	fpi.Platform.WriteTo(contentWritable)
	fpi.GameKey.WriteTo(contentWritable)
	fpi.Message.WriteTo(contentWritable)
	fpi.MessageUpdatedAt.WriteTo(contentWritable)
	fpi.MiiModifiedAt.WriteTo(contentWritable)
	fpi.LastOnline.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	fpi.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the FriendPersistentInfo from the given readable
func (fpi *FriendPersistentInfo) ExtractFrom(readable types.Readable) error {
	var err error

	err = fpi.Data.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendPersistentInfo.Data. %s", err.Error())
	}

	err = fpi.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendPersistentInfo header. %s", err.Error())
	}

	err = fpi.PID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendPersistentInfo.PID. %s", err.Error())
	}

	err = fpi.Region.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendPersistentInfo.Region. %s", err.Error())
	}

	err = fpi.Country.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendPersistentInfo.Country. %s", err.Error())
	}

	err = fpi.Area.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendPersistentInfo.Area. %s", err.Error())
	}

	err = fpi.Language.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendPersistentInfo.Language. %s", err.Error())
	}

	err = fpi.Platform.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendPersistentInfo.Platform. %s", err.Error())
	}

	err = fpi.GameKey.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendPersistentInfo.GameKey. %s", err.Error())
	}

	err = fpi.Message.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendPersistentInfo.Message. %s", err.Error())
	}

	err = fpi.MessageUpdatedAt.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendPersistentInfo.MessageUpdatedAt. %s", err.Error())
	}

	err = fpi.MiiModifiedAt.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendPersistentInfo.MiiModifiedAt. %s", err.Error())
	}

	err = fpi.LastOnline.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendPersistentInfo.LastOnline. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of FriendPersistentInfo
func (fpi *FriendPersistentInfo) Copy() types.RVType {
	copied := NewFriendPersistentInfo()

	copied.StructureVersion = fpi.StructureVersion
	copied.Data = fpi.Data.Copy().(*types.Data)
	copied.PID = fpi.PID.Copy().(*types.PID)
	copied.Region = fpi.Region.Copy().(*types.PrimitiveU8)
	copied.Country = fpi.Country.Copy().(*types.PrimitiveU8)
	copied.Area = fpi.Area.Copy().(*types.PrimitiveU8)
	copied.Language = fpi.Language.Copy().(*types.PrimitiveU8)
	copied.Platform = fpi.Platform.Copy().(*types.PrimitiveU8)
	copied.GameKey = fpi.GameKey.Copy().(*GameKey)
	copied.Message = fpi.Message.Copy().(*types.String)
	copied.MessageUpdatedAt = fpi.MessageUpdatedAt.Copy().(*types.DateTime)
	copied.MiiModifiedAt = fpi.MiiModifiedAt.Copy().(*types.DateTime)
	copied.LastOnline = fpi.LastOnline.Copy().(*types.DateTime)

	return copied
}

// Equals checks if the given FriendPersistentInfo contains the same data as the current FriendPersistentInfo
func (fpi *FriendPersistentInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*FriendPersistentInfo); !ok {
		return false
	}

	other := o.(*FriendPersistentInfo)

	if fpi.StructureVersion != other.StructureVersion {
		return false
	}

	if !fpi.Data.Equals(other.Data) {
		return false
	}

	if !fpi.PID.Equals(other.PID) {
		return false
	}

	if !fpi.Region.Equals(other.Region) {
		return false
	}

	if !fpi.Country.Equals(other.Country) {
		return false
	}

	if !fpi.Area.Equals(other.Area) {
		return false
	}

	if !fpi.Language.Equals(other.Language) {
		return false
	}

	if !fpi.Platform.Equals(other.Platform) {
		return false
	}

	if !fpi.GameKey.Equals(other.GameKey) {
		return false
	}

	if !fpi.Message.Equals(other.Message) {
		return false
	}

	if !fpi.MessageUpdatedAt.Equals(other.MessageUpdatedAt) {
		return false
	}

	if !fpi.MiiModifiedAt.Equals(other.MiiModifiedAt) {
		return false
	}

	return fpi.LastOnline.Equals(other.LastOnline)
}

// String returns the string representation of the FriendPersistentInfo
func (fpi *FriendPersistentInfo) String() string {
	return fpi.FormatToString(0)
}

// FormatToString pretty-prints the FriendPersistentInfo using the provided indentation level
func (fpi *FriendPersistentInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("FriendPersistentInfo{\n")
	b.WriteString(fmt.Sprintf("%sData (parent): %s,\n", indentationValues, fpi.Data.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sPID: %s,\n", indentationValues, fpi.PID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sRegion: %s,\n", indentationValues, fpi.Region))
	b.WriteString(fmt.Sprintf("%sCountry: %s,\n", indentationValues, fpi.Country))
	b.WriteString(fmt.Sprintf("%sArea: %s,\n", indentationValues, fpi.Area))
	b.WriteString(fmt.Sprintf("%sLanguage: %s,\n", indentationValues, fpi.Language))
	b.WriteString(fmt.Sprintf("%sPlatform: %s,\n", indentationValues, fpi.Platform))
	b.WriteString(fmt.Sprintf("%sGameKey: %s,\n", indentationValues, fpi.GameKey.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sMessage: %s,\n", indentationValues, fpi.Message))
	b.WriteString(fmt.Sprintf("%sMessageUpdatedAt: %s,\n", indentationValues, fpi.MessageUpdatedAt.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sMiiModifiedAt: %s,\n", indentationValues, fpi.MiiModifiedAt.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sLastOnline: %s,\n", indentationValues, fpi.LastOnline.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewFriendPersistentInfo returns a new FriendPersistentInfo
func NewFriendPersistentInfo() *FriendPersistentInfo {
	fpi := &FriendPersistentInfo{
		Data:             types.NewData(),
		PID:              types.NewPID(0),
		Region:           types.NewPrimitiveU8(0),
		Country:          types.NewPrimitiveU8(0),
		Area:             types.NewPrimitiveU8(0),
		Language:         types.NewPrimitiveU8(0),
		Platform:         types.NewPrimitiveU8(0),
		GameKey:          NewGameKey(),
		Message:          types.NewString(""),
		MessageUpdatedAt: types.NewDateTime(0),
		MiiModifiedAt:    types.NewDateTime(0),
		LastOnline:       types.NewDateTime(0),
	}

	return fpi
}
