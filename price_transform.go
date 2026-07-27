package talib

import "log/slog"

// AvgDev - Rolling average absolute deviation of a series from its own simple moving average over the last N periods.
// Measures dispersion around the window mean. Higher values indicate greater spread; zero when all values in the window are equal.
func AvgDev(inReal []float64, inTimePeriod int) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inReal))
	)

	if retCode := avgdev(
		startIdx,
		endIdx,
		inReal,
		int32(inTimePeriod),
		&outBegIdx,
		&outNBElement,
		outReal,
	); retCode != 0 {
		slog.Debug("AVGDEV", "result", retCode)
		return nil
	}

	return outReal
}

// AvgPrice - Average Price: the arithmetic mean of each bar's open, high, low, and close.
// A price-transform overlap condensing OHLC into a single representative price.
func AvgPrice(inOpen, inHigh, inLow, inClose []float64) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inOpen) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inOpen))
	)

	if retCode := avgprice(
		startIdx,
		endIdx,
		inOpen,
		inHigh,
		inLow,
		inClose,
		&outBegIdx,
		&outNBElement,
		outReal,
	); retCode != 0 {
		slog.Debug("AVGPRICE", "result", retCode)
		return nil
	}

	return outReal
}

// MedPrice - Median Price: the midpoint of each bar's high and low. A price-transform overlay.
//
// Median Price = (High + Low) / 2
func MedPrice(inHigh, inLow []float64) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inHigh) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inHigh))
	)

	if retCode := medprice(
		startIdx,
		endIdx,
		inHigh,
		inLow,
		&outBegIdx,
		&outNBElement,
		outReal,
	); retCode != 0 {
		slog.Debug("MEDPRICE", "result", retCode)
		return nil
	}

	return outReal
}

// TypPrice - Typical Price: the average of the high, low, and close of each bar. A single representative price per period.
//
// Typical Price = (High + Low + Close) / 3
func TypPrice(inHigh, inLow, inClose []float64) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inHigh) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inHigh))
	)

	if retCode := typprice(
		startIdx,
		endIdx,
		inHigh,
		inLow,
		inClose,
		&outBegIdx,
		&outNBElement,
		outReal,
	); retCode != 0 {
		slog.Debug("TYPPRICE", "result", retCode)
		return nil
	}

	return outReal
}

// WCLPrice - Weighted Close Price: a per-bar price average giving the close double weight relative to high and low.
//
// Weighted Close Price = (High + Low + 2 * Close) / 4
func WCLPrice(inHigh, inLow, inClose []float64) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inHigh) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inHigh))
	)

	if retCode := wclprice(
		startIdx,
		endIdx,
		inHigh,
		inLow,
		inClose,
		&outBegIdx,
		&outNBElement,
		outReal,
	); retCode != 0 {
		slog.Debug("WCLPRICE", "result", retCode)
		return nil
	}

	return outReal
}
