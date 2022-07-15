package lib

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func GetStringValue(a interface{}) (string, error) {

	switch a.(type) {
	case nil:
		return "", nil
	case int8:
		return strconv.FormatInt(int64(a.(int8)), 10), nil
	case []uint8:
		return string(a.([]uint8)), nil
	case uint8: //byte
		return strconv.FormatUint(uint64(a.(uint8)), 10), nil
	case int16:
		return strconv.FormatInt(int64(a.(int16)), 10), nil
	case uint16:
		return strconv.FormatUint(uint64(a.(uint16)), 10), nil
	case int32: // rune
		return strconv.FormatInt(int64(a.(int32)), 10), nil
	case uint32:
		return strconv.FormatUint(uint64(a.(uint32)), 10), nil
	case int64:
		return strconv.FormatInt(a.(int64), 10), nil
	case uint64:
		return strconv.FormatUint(a.(uint64), 10), nil
	case int:
		return strconv.Itoa(a.(int)), nil
	case uint:
		return strconv.FormatUint(uint64(a.(uint)), 10), nil
	case float32:
		return strconv.FormatFloat(float64(a.(float32)), 'f', 0, 64), nil
	case float64:
		return strconv.FormatFloat(a.(float64), 'f', 0, 64), nil
	case complex64:
		return fmt.Sprint(a), nil
	case complex128:
		return fmt.Sprint(a), nil
	case uintptr:
		return fmt.Sprint(a), nil
	case string:
		return a.(string), nil
	default: // 其他类型有pointer， struct， array, slice ,map, interface, function, channel
		val := reflect.ValueOf(a)
		typ := reflect.Indirect(val).Type()
		fmt.Println(typ)

		return "", errors.New("type not support:" + fmt.Sprint(typ))
	}
}
