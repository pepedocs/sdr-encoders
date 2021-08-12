package sdrencoders

import (
	"errors"
	"math"
)

type NumberSdr struct {
	Size          float64
	BucketIndex   float64
	NumActiveBits float64
}

func NewNumberSdr(size float64, bucketIndex float64, numActiveBits float64) (sdr NumberSdr, err error) {
	sdr = NumberSdr{
		Size:          size,
		BucketIndex:   bucketIndex,
		NumActiveBits: numActiveBits,
	}
	return
}

type NumberInputSpec struct {
	MinVal        float64
	MaxVal        float64
	NumBuckets    float64
	NumActiveBits float64
}

func NewNumberInputSpec(minVal float64, maxVal float64, numBuckets float64, numActiveBits float64) (spec NumberInputSpec, err error) {
	if maxVal <= minVal {
		return spec, errors.New("maxVal must be > minVal")
	}
	spec = NumberInputSpec{
		MinVal:        minVal,
		MaxVal:        maxVal,
		NumBuckets:    numBuckets,
		NumActiveBits: numActiveBits,
	}
	return
}

func (spec *NumberInputSpec) ValidateInput(value float64) (err error) {
	if value < spec.MinVal || value > spec.MaxVal {
		err = errors.New("value must be within the minVal and maxVal range")
	}
	return
}

type NumberEncoder struct {
	spec NumberInputSpec
}

func NewNumberEncoder(spec NumberInputSpec) (encoder NumberEncoder) {
	return NumberEncoder{spec}
}

// https://numenta.com/assets/pdf/biological-and-machine-intelligence/BaMI-Encoders.pdf
func (e *NumberEncoder) EncodeNumberSdr(value float64, spec NumberInputSpec) (sdr NumberSdr, err error) {
	err = spec.ValidateInput(value)
	if err != nil {
		return sdr, err
	}

	valRange := spec.MaxVal - spec.MinVal
	size := spec.NumBuckets + spec.NumActiveBits + 1
	bucketIndex := math.Floor(spec.NumBuckets * (value - spec.MinVal) / valRange)
	sdr, err = NewNumberSdr(size, bucketIndex, spec.NumActiveBits)
	return
}
