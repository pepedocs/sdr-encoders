package htmencoders

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	inputValue  float64
	inputSpec   NumberInputSpec
	expectedSdr NumberSdr
}

func TestNumberEncoder(t *testing.T) {
	var tests = []TestCase{}
	spec, _ := NewNumberInputSpec(0.0, 10.0, 10.0, 3.0)
	expectSdr, _ := NewNumberSdr(14.0, 10.0, 3.0)
	tcase := TestCase{10.0, spec, expectSdr}
	tests = append(tests, tcase)

	for _, test := range tests {
		testName := fmt.Sprintf(
			"input=%f, maxVal=%f, minVal=%f, numBuckets=%f, numActiveBits=%f",
			test.inputValue,
			test.inputSpec.MaxVal,
			test.inputSpec.MinVal,
			test.inputSpec.NumBuckets,
			test.inputSpec.NumActiveBits)
		t.Run(testName, func(t *testing.T) {
			encoder := NewNumberEncoder(test.inputSpec)
			result, err := encoder.EncodeNumberSdr(test.inputValue, test.inputSpec)
			if err != nil {
				t.Errorf("encoding failed")
			}
			if !reflect.DeepEqual(result, test.expectedSdr) {
				t.Errorf("got %v, expect %v", result, test.expectedSdr)
			}
		})
	}
}
