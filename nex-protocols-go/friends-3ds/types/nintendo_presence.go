// Package types implements all the types used by the Friends3DS protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// NintendoPresence is a type within the Friends3DS protocol
type NintendoPresence struct {
	types.Structure
	*types.Data
	ChangedFlags      *types.PrimitiveU32
	GameKey           *GameKey
	Message           *types.String
	JoinAvailableFlag *types.PrimitiveU32
	MatchmakeType     *types.PrimitiveU8
	JoinGameID        *types.PrimitiveU32
	JoinGameMode      *types.PrimitiveU32
	OwnerPID          *types.PID
	JoinGroupID       *types.PrimitiveU32
	ApplicationArg    *types.Buffer
}

// WriteTo writes the NintendoPresence to the given writable
func (np *NintendoPresence) WriteTo(writable types.Writable) {
	np.Data.WriteTo(writable)

	contentWritable := writable.CopyNew()

	np.ChangedFlags.WriteTo(contentWritable)
	np.GameKey.WriteTo(contentWritable)
	np.Message.WriteTo(contentWritable)
	np.JoinAvailableFlag.WriteTo(contentWritable)
	np.MatchmakeType.WriteTo(contentWritable)
	np.JoinGameID.WriteTo(contentWritable)
	np.JoinGameMode.WriteTo(contentWritable)
	np.OwnerPID.WriteTo(contentWritable)
	np.JoinGroupID.WriteTo(contentWritable)
	np.ApplicationArg.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	np.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the NintendoPresence from the given readable
func (np *NintendoPresence) ExtractFrom(readable types.Readable) error {
	var err error

	err = np.Data.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresence.Data. %s", err.Error())
	}

	err = np.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresence header. %s", err.Error())
	}

	err = np.ChangedFlags.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresence.ChangedFlags. %s", err.Error())
	}

	err = np.GameKey.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresence.GameKey. %s", err.Error())
	}

	err = np.Message.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresence.Message. %s", err.Error())
	}

	err = np.JoinAvailableFlag.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresence.JoinAvailableFlag. %s", err.Error())
	}

	err = np.MatchmakeType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresence.MatchmakeType. %s", err.Error())
	}

	err = np.JoinGameID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresence.JoinGameID. %s", err.Error())
	}

	err = np.JoinGameMode.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresence.JoinGameMode. %s", err.Error())
	}

	err = np.OwnerPID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresence.OwnerPID. %s", err.Error())
	}

	err = np.JoinGroupID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresence.JoinGroupID. %s", err.Error())
	}

	err = np.ApplicationArg.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresence.ApplicationArg. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of NintendoPresence
func (np *NintendoPresence) Copy() types.RVType {
	copied := NewNintendoPresence()

	copied.StructureVersion = np.StructureVersion
	copied.Data = np.Data.Copy().(*types.Data)
	copied.ChangedFlags = np.ChangedFlags.Copy().(*types.PrimitiveU32)
	copied.GameKey = np.GameKey.Copy().(*GameKey)
	copied.Message = np.Message.Copy().(*types.String)
	copied.JoinAvailableFlag = np.JoinAvailableFlag.Copy().(*types.PrimitiveU32)
	copied.MatchmakeType = np.MatchmakeType.Copy().(*types.PrimitiveU8)
	copied.JoinGameID = np.JoinGameID.Copy().(*types.PrimitiveU32)
	copied.JoinGameMode = np.JoinGameMode.Copy().(*types.PrimitiveU32)
	copied.OwnerPID = np.OwnerPID.Copy().(*types.PID)
	copied.JoinGroupID = np.JoinGroupID.Copy().(*types.PrimitiveU32)
	copied.ApplicationArg = np.ApplicationArg.Copy().(*types.Buffer)

	return copied
}

// Equals checks if the given NintendoPresence contains the same data as the current NintendoPresence
func (np *NintendoPresence) Equals(o types.RVType) bool {
	if _, ok := o.(*NintendoPresence); !ok {
		return false
	}

	other := o.(*NintendoPresence)

	if np.StructureVersion != other.StructureVersion {
		return false
	}

	if !np.Data.Equals(other.Data) {
		return false
	}

	if !np.ChangedFlags.Equals(other.ChangedFlags) {
		return false
	}

	if !np.GameKey.Equals(other.GameKey) {
		return false
	}

	if !np.Message.Equals(other.Message) {
		return false
	}

	if !np.JoinAvailableFlag.Equals(other.JoinAvailableFlag) {
		return false
	}

	if !np.MatchmakeType.Equals(other.MatchmakeType) {
		return false
	}

	if !np.JoinGameID.Equals(other.JoinGameID) {
		return false
	}

	if !np.JoinGameMode.Equals(other.JoinGameMode) {
		return false
	}

	if !np.OwnerPID.Equals(other.OwnerPID) {
		return false
	}

	if !np.JoinGroupID.Equals(other.JoinGroupID) {
		return false
	}

	return np.ApplicationArg.Equals(other.ApplicationArg)
}

// String returns the string representation of the NintendoPresence
func (np *NintendoPresence) String() string {
	return np.FormatToString(0)
}

// FormatToString pretty-prints the NintendoPresence using the provided indentation level
func (np *NintendoPresence) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("NintendoPresence{\n")
	b.WriteString(fmt.Sprintf("%sData (parent): %s,\n", indentationValues, np.Data.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sChangedFlags: %s,\n", indentationValues, np.ChangedFlags))
	b.WriteString(fmt.Sprintf("%sGameKey: %s,\n", indentationValues, np.GameKey.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sMessage: %s,\n", indentationValues, np.Message))
	b.WriteString(fmt.Sprintf("%sJoinAvailableFlag: %s,\n", indentationValues, np.JoinAvailableFlag))
	b.WriteString(fmt.Sprintf("%sMatchmakeType: %s,\n", indentationValues, np.MatchmakeType))
	b.WriteString(fmt.Sprintf("%sJoinGameID: %s,\n", indentationValues, np.JoinGameID))
	b.WriteString(fmt.Sprintf("%sJoinGameMode: %s,\n", indentationValues, np.JoinGameMode))
	b.WriteString(fmt.Sprintf("%sOwnerPID: %s,\n", indentationValues, np.OwnerPID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sJoinGroupID: %s,\n", indentationValues, np.JoinGroupID))
	b.WriteString(fmt.Sprintf("%sApplicationArg: %s,\n", indentationValues, np.ApplicationArg))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewNintendoPresence returns a new NintendoPresence
func NewNintendoPresence() *NintendoPresence {
	np := &NintendoPresence{
		Data:              types.NewData(),
		ChangedFlags:      types.NewPrimitiveU32(0),
		GameKey:           NewGameKey(),
		Message:           types.NewString(""),
		JoinAvailableFlag: types.NewPrimitiveU32(0),
		MatchmakeType:     types.NewPrimitiveU8(0),
		JoinGameID:        types.NewPrimitiveU32(0),
		JoinGameMode:      types.NewPrimitiveU32(0),
		OwnerPID:          types.NewPID(0),
		JoinGroupID:       types.NewPrimitiveU32(0),
		ApplicationArg:    types.NewBuffer(nil),
	}

	return np
}
