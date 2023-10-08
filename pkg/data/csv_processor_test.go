package data_test

import (
	"testing"

	"renatoaraujo/uk-visa-sponsors/pkg/data"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type MockStruct struct {
	Field1 string `csv:"field1"`
	Field2 int    `csv:"field2"`
}

func TestProcessRawData(t *testing.T) {
	cp := data.NewCSVProcessor()

	var validRef []MockStruct
	validCSV := []byte("field1,field2\nvalue1,10\nvalue2,20")
	err := cp.ProcessRawData(validCSV, &validRef)
	require.NoError(t, err)
	assert.Equal(t, "value1", validRef[0].Field1)
	assert.Equal(t, 10, validRef[0].Field2)

	var invalidRef MockStruct
	err = cp.ProcessRawData(validCSV, &invalidRef)
	require.Error(t, err)

	var mismatchRef []*MockStruct
	mismatchCSV := []byte("field3,field4\nvalue3,30\nvalue4,40")
	err = cp.ProcessRawData(mismatchCSV, &mismatchRef)
	require.Error(t, err)
}
