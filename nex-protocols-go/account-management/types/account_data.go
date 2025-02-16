// Package types implements all the types used by the AccountManagement protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// AccountData is a type within the AccountManagement protocol
type AccountData struct {
	types.Structure
	PID                *types.PID
	StrName            *types.String
	UIGroups           *types.PrimitiveU32
	StrEmail           *types.String
	DTCreationDate     *types.DateTime
	DTEffectiveDate    *types.DateTime
	StrNotEffectiveMsg *types.String
	DTExpiryDate       *types.DateTime
	StrExpiredMsg      *types.String
}

// WriteTo writes the AccountData to the given writable
func (ad *AccountData) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	ad.PID.WriteTo(contentWritable)
	ad.StrName.WriteTo(contentWritable)
	ad.UIGroups.WriteTo(contentWritable)
	ad.StrEmail.WriteTo(contentWritable)
	ad.DTCreationDate.WriteTo(contentWritable)
	ad.DTEffectiveDate.WriteTo(contentWritable)
	ad.StrNotEffectiveMsg.WriteTo(contentWritable)
	ad.DTExpiryDate.WriteTo(contentWritable)
	ad.StrExpiredMsg.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	ad.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the AccountData from the given readable
func (ad *AccountData) ExtractFrom(readable types.Readable) error {
	var err error

	err = ad.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract AccountData header. %s", err.Error())
	}

	err = ad.PID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract AccountData.PID. %s", err.Error())
	}

	err = ad.StrName.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract AccountData.StrName. %s", err.Error())
	}

	err = ad.UIGroups.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract AccountData.UIGroups. %s", err.Error())
	}

	err = ad.StrEmail.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract AccountData.StrEmail. %s", err.Error())
	}

	err = ad.DTCreationDate.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract AccountData.DTCreationDate. %s", err.Error())
	}

	err = ad.DTEffectiveDate.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract AccountData.DTEffectiveDate. %s", err.Error())
	}

	err = ad.StrNotEffectiveMsg.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract AccountData.StrNotEffectiveMsg. %s", err.Error())
	}

	err = ad.DTExpiryDate.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract AccountData.DTExpiryDate. %s", err.Error())
	}

	err = ad.StrExpiredMsg.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract AccountData.StrExpiredMsg. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of AccountData
func (ad *AccountData) Copy() types.RVType {
	copied := NewAccountData()

	copied.StructureVersion = ad.StructureVersion
	copied.PID = ad.PID.Copy().(*types.PID)
	copied.StrName = ad.StrName.Copy().(*types.String)
	copied.UIGroups = ad.UIGroups.Copy().(*types.PrimitiveU32)
	copied.StrEmail = ad.StrEmail.Copy().(*types.String)
	copied.DTCreationDate = ad.DTCreationDate.Copy().(*types.DateTime)
	copied.DTEffectiveDate = ad.DTEffectiveDate.Copy().(*types.DateTime)
	copied.StrNotEffectiveMsg = ad.StrNotEffectiveMsg.Copy().(*types.String)
	copied.DTExpiryDate = ad.DTExpiryDate.Copy().(*types.DateTime)
	copied.StrExpiredMsg = ad.StrExpiredMsg.Copy().(*types.String)

	return copied
}

// Equals checks if the given AccountData contains the same data as the current AccountData
func (ad *AccountData) Equals(o types.RVType) bool {
	if _, ok := o.(*AccountData); !ok {
		return false
	}

	other := o.(*AccountData)

	if ad.StructureVersion != other.StructureVersion {
		return false
	}

	if !ad.PID.Equals(other.PID) {
		return false
	}

	if !ad.StrName.Equals(other.StrName) {
		return false
	}

	if !ad.UIGroups.Equals(other.UIGroups) {
		return false
	}

	if !ad.StrEmail.Equals(other.StrEmail) {
		return false
	}

	if !ad.DTCreationDate.Equals(other.DTCreationDate) {
		return false
	}

	if !ad.DTEffectiveDate.Equals(other.DTEffectiveDate) {
		return false
	}

	if !ad.StrNotEffectiveMsg.Equals(other.StrNotEffectiveMsg) {
		return false
	}

	if !ad.DTExpiryDate.Equals(other.DTExpiryDate) {
		return false
	}

	return ad.StrExpiredMsg.Equals(other.StrExpiredMsg)
}

// String returns the string representation of the AccountData
func (ad *AccountData) String() string {
	return ad.FormatToString(0)
}

// FormatToString pretty-prints the AccountData using the provided indentation level
func (ad *AccountData) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("AccountData{\n")
	b.WriteString(fmt.Sprintf("%sPID: %s,\n", indentationValues, ad.PID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sStrName: %s,\n", indentationValues, ad.StrName))
	b.WriteString(fmt.Sprintf("%sUIGroups: %s,\n", indentationValues, ad.UIGroups))
	b.WriteString(fmt.Sprintf("%sStrEmail: %s,\n", indentationValues, ad.StrEmail))
	b.WriteString(fmt.Sprintf("%sDTCreationDate: %s,\n", indentationValues, ad.DTCreationDate.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sDTEffectiveDate: %s,\n", indentationValues, ad.DTEffectiveDate.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sStrNotEffectiveMsg: %s,\n", indentationValues, ad.StrNotEffectiveMsg))
	b.WriteString(fmt.Sprintf("%sDTExpiryDate: %s,\n", indentationValues, ad.DTExpiryDate.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sStrExpiredMsg: %s,\n", indentationValues, ad.StrExpiredMsg))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewAccountData returns a new AccountData
func NewAccountData() *AccountData {
	ad := &AccountData{
		PID:                types.NewPID(0),
		StrName:            types.NewString(""),
		UIGroups:           types.NewPrimitiveU32(0),
		StrEmail:           types.NewString(""),
		DTCreationDate:     types.NewDateTime(0),
		DTEffectiveDate:    types.NewDateTime(0),
		StrNotEffectiveMsg: types.NewString(""),
		DTExpiryDate:       types.NewDateTime(0),
		StrExpiredMsg:      types.NewString(""),
	}

	return ad
}
