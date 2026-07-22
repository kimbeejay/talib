package talib

// Cycle Indicators (HT) - Hilbert Transform
var (
	ht_dcperiod func(
		startIdx int,
		endIdx int,
		inReal []float64,
		outBegIdx *int,
		outNBElement *int,
		outReal []float64,
	) int

	ht_dcphase func(
		startIdx int,
		endIdx int,
		inReal []float64,
		outBegIdx *int,
		outNBElement *int,
		outReal []float64,
	) int

	ht_phasor func(
		startIdx int,
		endIdx int,
		inReal []float64,
		outBegIdx *int,
		outNBElement *int,
		outInPhase []float64,
		outQuadrature []float64,
	) int

	ht_sine func(
		startIdx int,
		endIdx int,
		inReal []float64,
		outBegIdx *int,
		outNBElement *int,
		outSine []float64,
		outLeadSine []float64,
	) int

	ht_trendmode func(
		startIdx int,
		endIdx int,
		inReal []float64,
		outBegIdx *int,
		outNBElement *int,
		outInteger []int,
	) int
)

// Math Operators
var (
	add func(
		startIdx int,
		endIdx int,
		inReal0 []float64,
		inReal1 []float64,
		outBegIdx *int,
		outNBElement *int,
		outReal []float64,
	) int
)
