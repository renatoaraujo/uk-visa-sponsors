package data

import (
	"errors"
	"reflect"

	"github.com/jszwec/csvutil"
)

type CSVProcessor struct{}

func NewCSVProcessor() *CSVProcessor {
	return &CSVProcessor{}
}

func isSliceOfStructsPointer(v interface{}) bool {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Ptr {
		return false
	}

	elem := val.Elem()
	if elem.Kind() != reflect.Slice {
		return false
	}
	if elem.Type().Elem().Kind() != reflect.Struct {
		return false
	}
	return true
}

func (cp *CSVProcessor) ProcessRawData(data []byte, ref interface{}) error {
	if !isSliceOfStructsPointer(ref) {
		return errors.New("expected a pointer to a slice of structs")
	}
	return csvutil.Unmarshal(data, ref)
}
