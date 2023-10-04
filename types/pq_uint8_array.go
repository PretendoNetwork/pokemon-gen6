package types

import (
	"fmt"
	"strconv"
	"strings"
)

type PQUInt8Array struct {
	Value []uint8
}

func (array *PQUInt8Array) Scan(src interface{}) error {

	switch src := src.(type) {
	case []uint8:
		// * Postgres stores uint8 slices as
		// * a string in the format:
		// * {1,2,3,4}
		str := string(src)
		str = strings.TrimSuffix(str, "}")
		str = strings.TrimPrefix(str, "{")
		strSlice := strings.Split(str, ",")

		for i := 0; i < len(strSlice); i++ {
			number, err := strconv.Atoi(strSlice[i])

			if err != nil {
				return err
			}

			array.Value = append(array.Value, uint8(number))
		}

		return nil
	}

	return fmt.Errorf("Type %T does not have a path to convert to UInt8Array", src)
}
