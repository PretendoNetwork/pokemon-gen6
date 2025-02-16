// Package types implements all the types used by the FriendsWiiU protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// NintendoPresenceV2 is a type within the FriendsWiiU protocol
type NintendoPresenceV2 struct {
	types.Structure
	*types.Data
	ChangedFlags    *types.PrimitiveU32
	Online          *types.PrimitiveBool
	GameKey         *GameKey
	Unknown1        *types.PrimitiveU8
	Message         *types.String
	Unknown2        *types.PrimitiveU32
	Unknown3        *types.PrimitiveU8
	GameServerID    *types.PrimitiveU32
	Unknown4        *types.PrimitiveU32
	PID             *types.PID
	GatheringID     *types.PrimitiveU32
	ApplicationData *types.Buffer
	Unknown5        *types.PrimitiveU8
	Unknown6        *types.PrimitiveU8
	Unknown7        *types.PrimitiveU8
}

// WriteTo writes the NintendoPresenceV2 to the given writable
func (npv *NintendoPresenceV2) WriteTo(writable types.Writable) {
	npv.Data.WriteTo(writable)

	contentWritable := writable.CopyNew()

	npv.ChangedFlags.WriteTo(contentWritable)
	npv.Online.WriteTo(contentWritable)
	npv.GameKey.WriteTo(contentWritable)
	npv.Unknown1.WriteTo(contentWritable)
	npv.Message.WriteTo(contentWritable)
	npv.Unknown2.WriteTo(contentWritable)
	npv.Unknown3.WriteTo(contentWritable)
	npv.GameServerID.WriteTo(contentWritable)
	npv.Unknown4.WriteTo(contentWritable)
	npv.PID.WriteTo(contentWritable)
	npv.GatheringID.WriteTo(contentWritable)
	npv.ApplicationData.WriteTo(contentWritable)
	npv.Unknown5.WriteTo(contentWritable)
	npv.Unknown6.WriteTo(contentWritable)
	npv.Unknown7.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	npv.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the NintendoPresenceV2 from the given readable
func (npv *NintendoPresenceV2) ExtractFrom(readable types.Readable) error {
	var err error

	err = npv.Data.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.Data. %s", err.Error())
	}

	err = npv.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2 header. %s", err.Error())
	}

	err = npv.ChangedFlags.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.ChangedFlags. %s", err.Error())
	}

	err = npv.Online.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.Online. %s", err.Error())
	}

	err = npv.GameKey.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.GameKey. %s", err.Error())
	}

	err = npv.Unknown1.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.Unknown1. %s", err.Error())
	}

	err = npv.Message.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.Message. %s", err.Error())
	}

	err = npv.Unknown2.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.Unknown2. %s", err.Error())
	}

	err = npv.Unknown3.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.Unknown3. %s", err.Error())
	}

	err = npv.GameServerID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.GameServerID. %s", err.Error())
	}

	err = npv.Unknown4.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.Unknown4. %s", err.Error())
	}

	err = npv.PID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.PID. %s", err.Error())
	}

	err = npv.GatheringID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.GatheringID. %s", err.Error())
	}

	err = npv.ApplicationData.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.ApplicationData. %s", err.Error())
	}

	err = npv.Unknown5.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.Unknown5. %s", err.Error())
	}

	err = npv.Unknown6.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.Unknown6. %s", err.Error())
	}

	err = npv.Unknown7.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.Unknown7. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of NintendoPresenceV2
func (npv *NintendoPresenceV2) Copy() types.RVType {
	copied := NewNintendoPresenceV2()

	copied.StructureVersion = npv.StructureVersion
	copied.Data = npv.Data.Copy().(*types.Data)
	copied.ChangedFlags = npv.ChangedFlags.Copy().(*types.PrimitiveU32)
	copied.Online = npv.Online.Copy().(*types.PrimitiveBool)
	copied.GameKey = npv.GameKey.Copy().(*GameKey)
	copied.Unknown1 = npv.Unknown1.Copy().(*types.PrimitiveU8)
	copied.Message = npv.Message.Copy().(*types.String)
	copied.Unknown2 = npv.Unknown2.Copy().(*types.PrimitiveU32)
	copied.Unknown3 = npv.Unknown3.Copy().(*types.PrimitiveU8)
	copied.GameServerID = npv.GameServerID.Copy().(*types.PrimitiveU32)
	copied.Unknown4 = npv.Unknown4.Copy().(*types.PrimitiveU32)
	copied.PID = npv.PID.Copy().(*types.PID)
	copied.GatheringID = npv.GatheringID.Copy().(*types.PrimitiveU32)
	copied.ApplicationData = npv.ApplicationData.Copy().(*types.Buffer)
	copied.Unknown5 = npv.Unknown5.Copy().(*types.PrimitiveU8)
	copied.Unknown6 = npv.Unknown6.Copy().(*types.PrimitiveU8)
	copied.Unknown7 = npv.Unknown7.Copy().(*types.PrimitiveU8)

	return copied
}

// Equals checks if the given NintendoPresenceV2 contains the same data as the current NintendoPresenceV2
func (npv *NintendoPresenceV2) Equals(o types.RVType) bool {
	if _, ok := o.(*NintendoPresenceV2); !ok {
		return false
	}

	other := o.(*NintendoPresenceV2)

	if npv.StructureVersion != other.StructureVersion {
		return false
	}

	if !npv.Data.Equals(other.Data) {
		return false
	}

	if !npv.ChangedFlags.Equals(other.ChangedFlags) {
		return false
	}

	if !npv.Online.Equals(other.Online) {
		return false
	}

	if !npv.GameKey.Equals(other.GameKey) {
		return false
	}

	if !npv.Unknown1.Equals(other.Unknown1) {
		return false
	}

	if !npv.Message.Equals(other.Message) {
		return false
	}

	if !npv.Unknown2.Equals(other.Unknown2) {
		return false
	}

	if !npv.Unknown3.Equals(other.Unknown3) {
		return false
	}

	if !npv.GameServerID.Equals(other.GameServerID) {
		return false
	}

	if !npv.Unknown4.Equals(other.Unknown4) {
		return false
	}

	if !npv.PID.Equals(other.PID) {
		return false
	}

	if !npv.GatheringID.Equals(other.GatheringID) {
		return false
	}

	if !npv.ApplicationData.Equals(other.ApplicationData) {
		return false
	}

	if !npv.Unknown5.Equals(other.Unknown5) {
		return false
	}

	if !npv.Unknown6.Equals(other.Unknown6) {
		return false
	}

	return npv.Unknown7.Equals(other.Unknown7)
}

// String returns the string representation of the NintendoPresenceV2
func (npv *NintendoPresenceV2) String() string {
	return npv.FormatToString(0)
}

// FormatToString pretty-prints the NintendoPresenceV2 using the provided indentation level
func (npv *NintendoPresenceV2) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("NintendoPresenceV2{\n")
	b.WriteString(fmt.Sprintf("%sData (parent): %s,\n", indentationValues, npv.Data.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sChangedFlags: %s,\n", indentationValues, npv.ChangedFlags))
	b.WriteString(fmt.Sprintf("%sOnline: %s,\n", indentationValues, npv.Online))
	b.WriteString(fmt.Sprintf("%sGameKey: %s,\n", indentationValues, npv.GameKey.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sUnknown1: %s,\n", indentationValues, npv.Unknown1))
	b.WriteString(fmt.Sprintf("%sMessage: %s,\n", indentationValues, npv.Message))
	b.WriteString(fmt.Sprintf("%sUnknown2: %s,\n", indentationValues, npv.Unknown2))
	b.WriteString(fmt.Sprintf("%sUnknown3: %s,\n", indentationValues, npv.Unknown3))
	b.WriteString(fmt.Sprintf("%sGameServerID: %s,\n", indentationValues, npv.GameServerID))
	b.WriteString(fmt.Sprintf("%sUnknown4: %s,\n", indentationValues, npv.Unknown4))
	b.WriteString(fmt.Sprintf("%sPID: %s,\n", indentationValues, npv.PID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sGatheringID: %s,\n", indentationValues, npv.GatheringID))
	b.WriteString(fmt.Sprintf("%sApplicationData: %s,\n", indentationValues, npv.ApplicationData))
	b.WriteString(fmt.Sprintf("%sUnknown5: %s,\n", indentationValues, npv.Unknown5))
	b.WriteString(fmt.Sprintf("%sUnknown6: %s,\n", indentationValues, npv.Unknown6))
	b.WriteString(fmt.Sprintf("%sUnknown7: %s,\n", indentationValues, npv.Unknown7))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewNintendoPresenceV2 returns a new NintendoPresenceV2
func NewNintendoPresenceV2() *NintendoPresenceV2 {
	npv := &NintendoPresenceV2{
		Data:            types.NewData(),
		ChangedFlags:    types.NewPrimitiveU32(0),
		Online:          types.NewPrimitiveBool(false),
		GameKey:         NewGameKey(),
		Unknown1:        types.NewPrimitiveU8(0),
		Message:         types.NewString(""),
		Unknown2:        types.NewPrimitiveU32(0),
		Unknown3:        types.NewPrimitiveU8(0),
		GameServerID:    types.NewPrimitiveU32(0),
		Unknown4:        types.NewPrimitiveU32(0),
		PID:             types.NewPID(0),
		GatheringID:     types.NewPrimitiveU32(0),
		ApplicationData: types.NewBuffer(nil),
		Unknown5:        types.NewPrimitiveU8(0),
		Unknown6:        types.NewPrimitiveU8(0),
		Unknown7:        types.NewPrimitiveU8(0),
	}

	return npv
}
