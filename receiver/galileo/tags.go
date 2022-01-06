package galileo

import "fmt"

type tags []tag

type tag struct {
	Tag   uint8       `json:"tag"`
	Value interface{} `json:"value"`
}

func (t *tag) SetValue(tagType string, val []byte) error {
	var v tagParser

	switch tagType {
	case "uint":
		v = &UintTag{}
	case "string":
		v = &StringTag{}
	case "time":
		v = &TimeTag{}
	case "coord":
		v = &CoordTag{}
	case "speed":
		v = &SpeedTag{}
	case "int":
		v = &IntTag{}
	case "bitstring":
		v = &BitsTag{}
	default:
		return fmt.Errorf("Unknown data type: %s. Meaning: %x", tagType, val)
	}

	if v == nil {
		return fmt.Errorf("Invalid tag pointer")
	}

	if err := v.Parse(val); err != nil {
		return err
	}

	t.Value = v

	return nil
}
