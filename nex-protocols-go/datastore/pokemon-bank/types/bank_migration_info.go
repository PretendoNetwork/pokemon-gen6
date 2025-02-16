// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// BankMigrationInfo is a type within the DataStore protocol
type BankMigrationInfo struct {
	types.Structure
	MigrationStatus *types.PrimitiveU32
	UpdatedTime     *types.DateTime
}

// WriteTo writes the BankMigrationInfo to the given writable
func (bmi *BankMigrationInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	bmi.MigrationStatus.WriteTo(contentWritable)
	bmi.UpdatedTime.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	bmi.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the BankMigrationInfo from the given readable
func (bmi *BankMigrationInfo) ExtractFrom(readable types.Readable) error {
	var err error

	err = bmi.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract BankMigrationInfo header. %s", err.Error())
	}

	err = bmi.MigrationStatus.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract BankMigrationInfo.MigrationStatus. %s", err.Error())
	}

	err = bmi.UpdatedTime.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract BankMigrationInfo.UpdatedTime. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of BankMigrationInfo
func (bmi *BankMigrationInfo) Copy() types.RVType {
	copied := NewBankMigrationInfo()

	copied.StructureVersion = bmi.StructureVersion
	copied.MigrationStatus = bmi.MigrationStatus.Copy().(*types.PrimitiveU32)
	copied.UpdatedTime = bmi.UpdatedTime.Copy().(*types.DateTime)

	return copied
}

// Equals checks if the given BankMigrationInfo contains the same data as the current BankMigrationInfo
func (bmi *BankMigrationInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*BankMigrationInfo); !ok {
		return false
	}

	other := o.(*BankMigrationInfo)

	if bmi.StructureVersion != other.StructureVersion {
		return false
	}

	if !bmi.MigrationStatus.Equals(other.MigrationStatus) {
		return false
	}

	return bmi.UpdatedTime.Equals(other.UpdatedTime)
}

// String returns the string representation of the BankMigrationInfo
func (bmi *BankMigrationInfo) String() string {
	return bmi.FormatToString(0)
}

// FormatToString pretty-prints the BankMigrationInfo using the provided indentation level
func (bmi *BankMigrationInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("BankMigrationInfo{\n")
	b.WriteString(fmt.Sprintf("%sMigrationStatus: %s,\n", indentationValues, bmi.MigrationStatus))
	b.WriteString(fmt.Sprintf("%sUpdatedTime: %s,\n", indentationValues, bmi.UpdatedTime.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewBankMigrationInfo returns a new BankMigrationInfo
func NewBankMigrationInfo() *BankMigrationInfo {
	bmi := &BankMigrationInfo{
		MigrationStatus: types.NewPrimitiveU32(0),
		UpdatedTime:     types.NewDateTime(0),
	}

	return bmi
}
