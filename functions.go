package talib

// Cycle Indicators (HT) - Hilbert Transform
var (
	ht_dcperiod func(
		startIdx int32,
		endIdx int32,
		inReal []float64,
		outBegIdx *int32,
		outNBElement *int32,
		outReal []float64,
	) int32

	ht_dcphase func(
		startIdx int32,
		endIdx int32,
		inReal []float64,
		outBegIdx *int32,
		outNBElement *int32,
		outReal []float64,
	) int32

	ht_phasor func(
		startIdx int32,
		endIdx int32,
		inReal []float64,
		outBegIdx *int32,
		outNBElement *int32,
		outInPhase []float64,
		outQuadrature []float64,
	) int32

	ht_sine func(
		startIdx int32,
		endIdx int32,
		inReal []float64,
		outBegIdx *int32,
		outNBElement *int32,
		outSine []float64,
		outLeadSine []float64,
	) int32

	ht_trendmode func(
		startIdx int32,
		endIdx int32,
		inReal []float64,
		outBegIdx *int32,
		outNBElement *int32,
		outInteger []int32,
	) int32
)

// Math Operators
var (
	add func(
		startIdx int32,
		endIdx int32,
		inReal0 []float64,
		inReal1 []float64,
		outBegIdx *int32,
		outNBElement *int32,
		outReal []float64,
	) int32

	div func(
		startIdx int32,
		endIdx int32,
		inReal0 []float64,
		inReal1 []float64,
		outBegIdx *int32,
		outNBElement *int32,
		outReal []float64,
	) int32

	max func(
		startIdx int32,
		endIdx int32,
		inReal []float64,
		optInTimePeriod int32,
		outBegIdx *int32,
		outNBElement *int32,
		outReal []float64,
	) int32

	maxIndex func(
		startIdx int32,
		endIdx int32,
		inReal []float64,
		optInTimePeriod int32,
		outBegIdx *int32,
		outNBElement *int32,
		outInteger []int32,
	) int32

	min func(
		startIdx int32,
		endIdx int32,
		inReal []float64,
		optInTimePeriod int32,
		outBegIdx *int32,
		outNBElement *int32,
		outReal []float64,
	) int32

	minIndex func(
		startIdx int32,
		endIdx int32,
		inReal []float64,
		optInTimePeriod int32,
		outBegIdx *int32,
		outNBElement *int32,
		outInteger []int32,
	) int32

	minMax func(
		startIdx int32,
		endIdx int32,
		inReal []float64,
		optInTimePeriod int32,
		outBegIdx *int32,
		outNBElement *int32,
		outMin []float64,
		outMax []float64,
	) int32

	minMaxIndex func(
		startIdx int32,
		endIdx int32,
		inReal []float64,
		optInTimePeriod int32,
		outBegIdx *int32,
		outNBElement *int32,
		outMinIndex []int32,
		outMaxIndex []int32,
	) int32

	mult func(
		startIdx int32,
		endIdx int32,
		inReal0 []float64,
		inReal1 []float64,
		outBegIdx *int32,
		outNBElement *int32,
		outReal []float64,
	) int32

	sub func(
		startIdx int32,
		endIdx int32,
		inReal0 []float64,
		inReal1 []float64,
		outBegIdx *int32,
		outNBElement *int32,
		outReal []float64,
	) int32

	sum func(
		startIdx int32,
		endIdx int32,
		inReal []float64,
		optInTimePeriod int32,
		outBegIdx *int32,
		outNBElement *int32,
		outReal []float64,
	) int32
)
