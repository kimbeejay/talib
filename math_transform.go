package talib

import "log/slog"

// ACos - Vector trigonometric arc cosine: applies acos() to each input value.
// A Math Transform passthrough with zero lookback.
//
// outReal[i] = acos(inReal[i])
func ACos(inReal []float64) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inReal))
	)

	if retCode := acos(
		startIdx,
		endIdx,
		inReal,
		&outBegIdx,
		&outNBElement,
		outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("ACos", "result", retCode)
		return nil
	}

	return outReal
}

// Cos - Element-wise trigonometric cosine of the input series.
// Applies the C library cos() to each sample.
func Cos(inReal []float64) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inReal))
	)

	if retCode := cos(
		startIdx,
		endIdx,
		inReal,
		&outBegIdx,
		&outNBElement,
		outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("Cos", "result", retCode)
		return nil
	}

	return outReal
}

// CosH - Vector hyperbolic cosine: applies cosh element-wise to each input value.
// A Math Transform primitive with no lookback.
//
// outReal[i] = cosh(inReal[i]) = (e^{inReal[i]} + e^{-inReal[i]}) / 2
func CosH(inReal []float64) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inReal))
	)

	if retCode := cosh(
		startIdx,
		endIdx,
		inReal,
		&outBegIdx,
		&outNBElement,
		outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("CosH", "result", retCode)
		return nil
	}

	return outReal
}

// Sin - Vector trigonometric sine: applies sin() element-wise to each input value.
// Part of the Math Transform group.
func Sin(inReal []float64) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inReal))
	)

	if retCode := sin(
		startIdx,
		endIdx,
		inReal,
		&outBegIdx,
		&outNBElement,
		outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("Sin", "result", retCode)
		return nil
	}

	return outReal
}

// ASin - Element-wise arcsine (inverse sine) of each input value.
// A vector math transform, not a market indicator.
func ASin(inReal []float64) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inReal))
	)

	if retCode := asin(
		startIdx,
		endIdx,
		inReal,
		&outBegIdx,
		&outNBElement,
		outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("ASin", "result", retCode)
		return nil
	}

	return outReal
}

// SinH - Element-wise hyperbolic sine of the input series.
// A vector math transform applying sinh() to each value.
//
// outReal[i] = sinh(inReal[i])
func SinH(inReal []float64) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inReal))
	)

	if retCode := sinh(
		startIdx,
		endIdx,
		inReal,
		&outBegIdx,
		&outNBElement,
		outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("SinH", "result", retCode)
		return nil
	}

	return outReal
}

// Tan - Vector trigonometric tangent: applies tan() element-wise to each input value.
func Tan(inReal []float64) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inReal))
	)

	if retCode := tan(
		startIdx,
		endIdx,
		inReal,
		&outBegIdx,
		&outNBElement,
		outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("Tan", "result", retCode)
		return nil
	}

	return outReal
}

// ATan - Vector trigonometric arc tangent: applies atan element-wise to each input.
// Pure math transform with no lookback.
func ATan(inReal []float64) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inReal))
	)

	if retCode := atan(
		startIdx,
		endIdx,
		inReal,
		&outBegIdx,
		&outNBElement,
		outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("ATan", "result", retCode)
		return nil
	}

	return outReal
}

// TanH - Vector hyperbolic tangent: applies tanh element-wise to the input series.
//
// outReal[i] = tanh(inReal[i])
func TanH(inReal []float64) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inReal))
	)

	if retCode := tanh(
		startIdx,
		endIdx,
		inReal,
		&outBegIdx,
		&outNBElement,
		outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("TanH", "result", retCode)
		return nil
	}

	return outReal
}

// Ceil - Vector ceiling: element-wise ceiling of each input value (smallest integer >= input).
func Ceil(inReal []float64) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inReal))
	)

	if retCode := ceil(
		startIdx,
		endIdx,
		inReal,
		&outBegIdx,
		&outNBElement,
		outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("Ceil", "result", retCode)
		return nil
	}

	return outReal
}

// Exp - Vector arithmetic exponential: applies the base-e exponential to each input value. Element-wise math transform.
//
// outReal[i] = exp(inReal[i]) = e^
func Exp(inReal []float64) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inReal))
	)

	if retCode := exp(
		startIdx,
		endIdx,
		inReal,
		&outBegIdx,
		&outNBElement,
		outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("Exp", "result", retCode)
		return nil
	}

	return outReal
}

// Floor - Vector floor: rounds each input value down to the nearest integer. Element-wise math transform.
func Floor(inReal []float64) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inReal))
	)

	if retCode := floor(
		startIdx,
		endIdx,
		inReal,
		&outBegIdx,
		&outNBElement,
		outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("Floor", "result", retCode)
		return nil
	}

	return outReal
}

// Sqrt - Vector square root: applies the square-root function element-wise to each input value.
func Sqrt(inReal []float64) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inReal))
	)

	if retCode := sqrt(
		startIdx,
		endIdx,
		inReal,
		&outBegIdx,
		&outNBElement,
		outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("Sqrt", "result", retCode)
		return nil
	}

	return outReal
}

// Ln - Vector natural logarithm: applies the natural log (base e) elementwise to the input series.
func Ln(inReal []float64) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inReal))
	)

	if retCode := ln(
		startIdx,
		endIdx,
		inReal,
		&outBegIdx,
		&outNBElement,
		outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("Ln", "result", retCode)
		return nil
	}

	return outReal
}

// Log10 - Vector base-10 logarithm. Applies log10 element-wise over each input value.
func Log10(inReal []float64) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inReal))
	)

	if retCode := log10(
		startIdx,
		endIdx,
		inReal,
		&outBegIdx,
		&outNBElement,
		outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("Log10", "result", retCode)
		return nil
	}

	return outReal
}
