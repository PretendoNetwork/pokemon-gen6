// Package types implements all the types used by the ServiceItem protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ServiceItemTransaction is a type within the ServiceItem protocol
type ServiceItemTransaction struct {
	types.Structure
	TransactionID          *types.String
	ExtTransactionID       *types.String
	Time                   *types.DateTime
	TransactionType        *types.PrimitiveU32
	TransactionDescription *types.String
	TransactionAmount      *ServiceItemAmount
	ItemCode               *types.String
	ReferenceID            *types.String
	Limitation             *ServiceItemLimitation
}

// WriteTo writes the ServiceItemTransaction to the given writable
func (sit *ServiceItemTransaction) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	sit.TransactionID.WriteTo(contentWritable)
	sit.ExtTransactionID.WriteTo(contentWritable)
	sit.Time.WriteTo(contentWritable)
	sit.TransactionType.WriteTo(contentWritable)
	sit.TransactionDescription.WriteTo(contentWritable)
	sit.TransactionAmount.WriteTo(contentWritable)
	sit.ItemCode.WriteTo(contentWritable)
	sit.ReferenceID.WriteTo(contentWritable)
	sit.Limitation.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	sit.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemTransaction from the given readable
func (sit *ServiceItemTransaction) ExtractFrom(readable types.Readable) error {
	var err error

	err = sit.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemTransaction header. %s", err.Error())
	}

	err = sit.TransactionID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemTransaction.TransactionID. %s", err.Error())
	}

	err = sit.ExtTransactionID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemTransaction.ExtTransactionID. %s", err.Error())
	}

	err = sit.Time.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemTransaction.Time. %s", err.Error())
	}

	err = sit.TransactionType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemTransaction.TransactionType. %s", err.Error())
	}

	err = sit.TransactionDescription.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemTransaction.TransactionDescription. %s", err.Error())
	}

	err = sit.TransactionAmount.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemTransaction.TransactionAmount. %s", err.Error())
	}

	err = sit.ItemCode.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemTransaction.ItemCode. %s", err.Error())
	}

	err = sit.ReferenceID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemTransaction.ReferenceID. %s", err.Error())
	}

	err = sit.Limitation.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemTransaction.Limitation. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemTransaction
func (sit *ServiceItemTransaction) Copy() types.RVType {
	copied := NewServiceItemTransaction()

	copied.StructureVersion = sit.StructureVersion
	copied.TransactionID = sit.TransactionID.Copy().(*types.String)
	copied.ExtTransactionID = sit.ExtTransactionID.Copy().(*types.String)
	copied.Time = sit.Time.Copy().(*types.DateTime)
	copied.TransactionType = sit.TransactionType.Copy().(*types.PrimitiveU32)
	copied.TransactionDescription = sit.TransactionDescription.Copy().(*types.String)
	copied.TransactionAmount = sit.TransactionAmount.Copy().(*ServiceItemAmount)
	copied.ItemCode = sit.ItemCode.Copy().(*types.String)
	copied.ReferenceID = sit.ReferenceID.Copy().(*types.String)
	copied.Limitation = sit.Limitation.Copy().(*ServiceItemLimitation)

	return copied
}

// Equals checks if the given ServiceItemTransaction contains the same data as the current ServiceItemTransaction
func (sit *ServiceItemTransaction) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemTransaction); !ok {
		return false
	}

	other := o.(*ServiceItemTransaction)

	if sit.StructureVersion != other.StructureVersion {
		return false
	}

	if !sit.TransactionID.Equals(other.TransactionID) {
		return false
	}

	if !sit.ExtTransactionID.Equals(other.ExtTransactionID) {
		return false
	}

	if !sit.Time.Equals(other.Time) {
		return false
	}

	if !sit.TransactionType.Equals(other.TransactionType) {
		return false
	}

	if !sit.TransactionDescription.Equals(other.TransactionDescription) {
		return false
	}

	if !sit.TransactionAmount.Equals(other.TransactionAmount) {
		return false
	}

	if !sit.ItemCode.Equals(other.ItemCode) {
		return false
	}

	if !sit.ReferenceID.Equals(other.ReferenceID) {
		return false
	}

	return sit.Limitation.Equals(other.Limitation)
}

// String returns the string representation of the ServiceItemTransaction
func (sit *ServiceItemTransaction) String() string {
	return sit.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemTransaction using the provided indentation level
func (sit *ServiceItemTransaction) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemTransaction{\n")
	b.WriteString(fmt.Sprintf("%sTransactionID: %s,\n", indentationValues, sit.TransactionID))
	b.WriteString(fmt.Sprintf("%sExtTransactionID: %s,\n", indentationValues, sit.ExtTransactionID))
	b.WriteString(fmt.Sprintf("%sTime: %s,\n", indentationValues, sit.Time.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sTransactionType: %s,\n", indentationValues, sit.TransactionType))
	b.WriteString(fmt.Sprintf("%sTransactionDescription: %s,\n", indentationValues, sit.TransactionDescription))
	b.WriteString(fmt.Sprintf("%sTransactionAmount: %s,\n", indentationValues, sit.TransactionAmount.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sItemCode: %s,\n", indentationValues, sit.ItemCode))
	b.WriteString(fmt.Sprintf("%sReferenceID: %s,\n", indentationValues, sit.ReferenceID))
	b.WriteString(fmt.Sprintf("%sLimitation: %s,\n", indentationValues, sit.Limitation.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemTransaction returns a new ServiceItemTransaction
func NewServiceItemTransaction() *ServiceItemTransaction {
	sit := &ServiceItemTransaction{
		TransactionID:          types.NewString(""),
		ExtTransactionID:       types.NewString(""),
		Time:                   types.NewDateTime(0),
		TransactionType:        types.NewPrimitiveU32(0),
		TransactionDescription: types.NewString(""),
		TransactionAmount:      NewServiceItemAmount(),
		ItemCode:               types.NewString(""),
		ReferenceID:            types.NewString(""),
		Limitation:             NewServiceItemLimitation(),
	}

	return sit
}
