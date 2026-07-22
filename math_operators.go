package talib

import "log/slog"

// Add - Vector arithmetic addition. Outputs the element-wise sum of two input series.
//
// outReal[i] = inReal0[i] + inReal1[i]
func Add(inReal0, inReal1 []float64) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal0) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inReal0))
	)

	if retCode := add(
		startIdx,
		endIdx,
		inReal0,
		inReal1,
		&outBegIdx,
		&outNBElement,
		outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("Add", "result", retCode)
		return nil
	}

	return outReal
}

// Div - Element-wise division of two input series. Computes the quotient of corresponding values from two real inputs.
//
// outReal[i] = inReal0[i] / inReal1[i]
func Div(inReal0, inReal1 []float64) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal0) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inReal0))
	)

	if retCode := div(
		startIdx,
		endIdx,
		inReal0,
		inReal1,
		&outBegIdx,
		&outNBElement,
		outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("Div", "result", retCode)
		return nil
	}

	return outReal
}

// Max - Highest input value over a rolling window of the last optInTimePeriod bars. A moving-window maximum.
func Max(inReal []float64, optInTimePeriod int) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inReal))
	)

	if retCode := max(
		startIdx,
		endIdx,
		inReal,
		int32(optInTimePeriod),
		&outBegIdx,
		&outNBElement,
		outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("Max", "result", retCode)
		return nil
	}

	return outReal
}

// MaxIndex - Returns the index of the highest input value within a rolling window of optInTimePeriod bars.
// Same as Max but outputs the location instead of the value.
//
// outInteger[i] = argmax_{j in [i-optInTimePeriod+1, i]} inReal[j]
func MaxIndex(inReal []float64, optInTimePeriod int) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal) - 1)
		outBegIdx    int32
		outNBElement int32
		outInteger   = make([]int32, len(inReal))
	)

	if retCode := maxIndex(
		startIdx,
		endIdx,
		inReal,
		int32(optInTimePeriod),
		&outBegIdx,
		&outNBElement,
		outInteger,
	); SUCCESS != taResult(retCode) {
		slog.Debug("MaxIndex", "result", retCode)
		return nil
	}

	return outInteger
}

// Min - Rolling minimum: the lowest input value over the trailing period.
//
// outReal[i] = min(inReal[i-optInTimePeriod+1 .. i])
func Min(inReal []float64, optInTimePeriod int) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inReal))
	)

	if retCode := min(
		startIdx,
		endIdx,
		inReal,
		int32(optInTimePeriod),
		&outBegIdx,
		&outNBElement,
		outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("Min", "result", retCode)
		return nil
	}

	return outReal
}

// MinIndex - Returns the absolute index of the lowest value within a rolling window of the given period.
// Same scan as Min but outputs the position of the minimum rather than its value.
//
// outInteger[t] = argmin_{t-period+1 <= i <= t} inReal[i] (absolute index into inReal)
func MinIndex(inReal []float64, optInTimePeriod int) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal) - 1)
		outBegIdx    int32
		outNBElement int32
		outInteger   = make([]int32, len(inReal))
	)

	if retCode := minIndex(
		startIdx,
		endIdx,
		inReal,
		int32(optInTimePeriod),
		&outBegIdx,
		&outNBElement,
		outInteger,
	); SUCCESS != taResult(retCode) {
		slog.Debug("MinIndex", "result", retCode)
		return nil
	}

	return outInteger
}

// MinMax - Returns both the lowest and highest values of the input over a rolling window of the last optInTimePeriod bars.
// An overlap-study companion to Min and Max that computes both extrema in one pass.
// @param inReal The input data series.
// @param optInTimePeriod The number of bars (or periods) to look back for the lowest and highest values.
// @return Two slices: the first contains the rolling minimum values, and the second contains the rolling maximum values.
func MinMax(inReal []float64, optInTimePeriod int) ([]float64, []float64) {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal) - 1)
		outBegIdx    int32
		outNBElement int32
		outMin       = make([]float64, len(inReal))
		outMax       = make([]float64, len(inReal))
	)

	if retCode := minMax(
		startIdx,
		endIdx,
		inReal,
		int32(optInTimePeriod),
		&outBegIdx,
		&outNBElement,
		outMin,
		outMax,
	); SUCCESS != taResult(retCode) {
		slog.Debug("MinMax", "result", retCode)
		return nil, nil
	}

	return outMin, outMax
}

// MinMaxIndex - Returns the absolute input indices of the lowest and highest values within each rolling window of optInTimePeriod bars. Index variant of MinMax.
//
// For each t: outMaxIdx[t] = argmax_{i in [t-N+1, t]} inReal[i]; outMinIdx[t] = argmin over the same window (N = optInTimePeriod).
func MinMaxIndex(inReal []float64, optInTimePeriod int) ([]int32, []int32) {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal) - 1)
		outBegIdx    int32
		outNBElement int32
		outMinIndex  = make([]int32, len(inReal))
		outMaxIndex  = make([]int32, len(inReal))
	)

	if retCode := minMaxIndex(
		startIdx,
		endIdx,
		inReal,
		int32(optInTimePeriod),
		&outBegIdx,
		&outNBElement,
		outMinIndex,
		outMaxIndex,
	); SUCCESS != taResult(retCode) {
		slog.Debug("MinMaxIndex", "result", retCode)
		return nil, nil
	}

	return outMinIndex, outMaxIndex
}

// Mult - Element-wise multiplication of two input series.
// Produces outReal[i] = inReal0[i] * inReal1[i].
func Mult(inReal0, inReal1 []float64) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal0) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inReal0))
	)

	if retCode := mult(
		startIdx,
		endIdx,
		inReal0,
		inReal1,
		&outBegIdx,
		&outNBElement,
		outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("Mult", "result", retCode)
		return nil
	}

	return outReal
}

// Sub - Element-wise vector subtraction of two input series. Outputs inReal0 minus inReal1 at each index.
//
// outReal[i] = inReal0[i] - inReal1[i]
func Sub(inReal0, inReal1 []float64) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal0) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inReal0))
	)

	if retCode := sub(
		startIdx,
		endIdx,
		inReal0,
		inReal1,
		&outBegIdx,
		&outNBElement,
		outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("Sub", "result", retCode)
		return nil
	}

	return outReal
}

// Sum - Rolling sum of the input over a fixed period.
// Each output is the sum of the most recent optInTimePeriod input values.
func Sum(inReal []float64, optInTimePeriod int) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inReal))
	)

	if retCode := sum(
		startIdx,
		endIdx,
		inReal,
		int32(optInTimePeriod),
		&outBegIdx,
		&outNBElement,
		outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("Sum", "result", retCode)
		return nil
	}

	return outReal
}
