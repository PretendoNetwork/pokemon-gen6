package nex_subscription_types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ActivePlayerSubscriptionData is a type within the Subscription protocol
type ActivePlayerSubscriptionData struct {
	types.Structure
	SubscriptionData
	ByteUnk types.UInt8
}

// WriteTo writes the ActivePlayerSubscriptionData to the given writable
func (apsd ActivePlayerSubscriptionData) WriteTo(writable types.Writable) {
	apsd.SubscriptionData.WriteTo(writable)

	contentWritable := writable.CopyNew()

	apsd.ByteUnk.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	apsd.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ActivePlayerSubscriptionData from the given readable
func (apsd *ActivePlayerSubscriptionData) ExtractFrom(readable types.Readable) error {
	var err error

	err = apsd.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ActivePlayerSubscriptionData header. %s", err.Error())
	}

	err = apsd.SubscriptionData.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ActivePlayerSubscriptionData.SubscriptionData. %s", err.Error())
	}

	err = apsd.ByteUnk.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ActivePlayerSubscriptionData.ByteUnk. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ActivePlayerSubscriptionData
func (apsd ActivePlayerSubscriptionData) Copy() types.RVType {
	copied := NewActivePlayerSubscriptionData()

	copied.StructureVersion = apsd.StructureVersion
	copied.SubscriptionData = apsd.SubscriptionData.Copy().(SubscriptionData)
	copied.ByteUnk = apsd.ByteUnk.Copy().(types.UInt8)

	return copied
}

// Equals checks if the given ActivePlayerSubscriptionData contains the same data as the current ActivePlayerSubscriptionData
func (apsd ActivePlayerSubscriptionData) Equals(o types.RVType) bool {
	if _, ok := o.(ActivePlayerSubscriptionData); !ok {
		return false
	}

	other := o.(ActivePlayerSubscriptionData)

	if apsd.StructureVersion != other.StructureVersion {
		return false
	}

	if !apsd.SubscriptionData.Equals(other.SubscriptionData) {
		return false
	}

	return apsd.ByteUnk.Equals(other.ByteUnk)
}

// CopyRef copies the current value of the ActivePlayerSubscriptionData
// and returns a pointer to the new copy
func (apsd ActivePlayerSubscriptionData) CopyRef() types.RVTypePtr {
	copied := apsd.Copy().(ActivePlayerSubscriptionData)
	return &copied
}

// Deref takes a pointer to the ActivePlayerSubscriptionData
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (apsd *ActivePlayerSubscriptionData) Deref() types.RVType {
	return *apsd
}

// String returns the string representation of the ActivePlayerSubscriptionData
func (apsd ActivePlayerSubscriptionData) String() string {
	return apsd.FormatToString(0)
}

// FormatToString pretty-prints the ActivePlayerSubscriptionData using the provided indentation level
func (apsd ActivePlayerSubscriptionData) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ActivePlayerSubscriptionData{\n")
	b.WriteString(fmt.Sprintf("%sSubscriptionData (parent): %s,\n", indentationValues, apsd.SubscriptionData.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sByteUnk: %s,\n", indentationValues, apsd.ByteUnk.String()))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewActivePlayerSubscriptionData returns a new ActivePlayerSubscriptionData
func NewActivePlayerSubscriptionData() ActivePlayerSubscriptionData {
	return ActivePlayerSubscriptionData{
		SubscriptionData: NewSubscriptionData(),
		ByteUnk:          types.NewUInt8(0),
	}

}
