package talib

import "log/slog"

// AccBands - Acceleration Bands: three overlap lines around price.
// The middle band is an SMA of the close; the upper/lower bands are SMAs of the high/low scaled by an intraday-range factor.
//
// factor = 4*(H-L)/(H+L)
// upperRaw = H*(1+factor), lowerRaw = L*(1-factor)
// Upper = SMA(upperRaw, N), Middle = SMA(Close, N), Lower = SMA(lowerRaw, N)
func AccBands(inHigh, inLow, inClose []float64, inTimePeriod int) ([]float64, []float64, []float64) {
	var (
		startIdx          int32
		endIdx            = int32(len(inClose) - 1)
		outBegIdx         int32
		outNBElement      int32
		outRealUpperBand  = make([]float64, len(inClose))
		outRealMiddleBand = make([]float64, len(inClose))
		outRealLowerBand  = make([]float64, len(inClose))
	)

	if retCode := accbands(
		startIdx,
		endIdx,
		inHigh,
		inLow,
		inClose,
		int32(inTimePeriod),
		&outBegIdx,
		&outNBElement,
		outRealUpperBand,
		outRealMiddleBand,
		outRealLowerBand,
	); retCode != 0 {
		slog.Debug("AccBands", "result", retCode)
		return nil, nil, nil
	}

	return outRealUpperBand, outRealMiddleBand, outRealLowerBand
}

// BBands - Bollinger Bands: a moving-average middle band with upper and lower bands offset by a multiple of the standard deviation.
// Used to gauge relative price volatility.
//
// middle = MA(inReal, period); sd = stddev(inReal, period);
// upper = middle + nbDevUpsd;
// lower = middle - nbDevDnsd
//
// Note: The standard deviation uses the population form (dividing by the period), not the sample form.
// The standard deviation is always computed with a simple moving average regardless of the selected MA type.
func BBands(inReal []float64, inTimePeriod int, inNbDevUp, inNbDevDn float64, inMAType MAType) ([]float64, []float64, []float64) {
	var (
		startIdx          int32
		endIdx            = int32(len(inReal) - 1)
		outBegIdx         int32
		outNBElement      int32
		outRealUpperBand  = make([]float64, len(inReal))
		outRealMiddleBand = make([]float64, len(inReal))
		outRealLowerBand  = make([]float64, len(inReal))
	)

	if retCode := bbands(
		startIdx,
		endIdx,
		inReal,
		int32(inTimePeriod),
		inNbDevUp,
		inNbDevDn,
		int32(inMAType),
		&outBegIdx,
		&outNBElement,
		outRealUpperBand,
		outRealMiddleBand,
		outRealLowerBand,
	); retCode != 0 {
		slog.Debug("BBands", "result", retCode)
		return nil, nil, nil
	}

	return outRealUpperBand, outRealMiddleBand, outRealLowerBand
}

// DEMA - Double Exponential Moving Average: an EMA combined with an EMA-of-EMA to reduce lag versus a plain EMA.
//
// EMA1 = EMA(inReal, period); EMA2 = EMA(EMA1, period); DEMA = 2*EMA1 - EMA2
//
// Note: A period of 1 performs no smoothing: the output is a copy of the input. Allowed since 0.6.5 (issues #48/#59).
func DEMA(inReal []float64, inTimePeriod int) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inReal))
	)

	if retCode := dema(
		startIdx,
		endIdx,
		inReal,
		int32(inTimePeriod),
		&outBegIdx,
		&outNBElement,
		outReal,
	); retCode != 0 {
		slog.Debug("DEMA", "result", retCode)
		return nil
	}

	return outReal
}

// EMA - Exponential moving average that weights recent prices more heavily via a recursive smoothing factor.
// A core building block seeding or composing many other indicators. Reacts faster than SMA; price above/below EMA suggests up/down trend.
//
// k = 2 / (period + 1);
// EMA_t = (price_t - EMA_{t-1}) * k + EMA_{t-1}.
// Seed: EMA = SMA of first period bars.
//
// Note: A period of 1 performs no smoothing: the output is a copy of the input. Allowed since 0.6.5 (issues #48/#59).
func EMA(inReal []float64, inTimePeriod int) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inReal))
	)

	if retCode := ema(
		startIdx,
		endIdx,
		inReal,
		int32(inTimePeriod),
		&outBegIdx,
		&outNBElement,
		outReal,
	); retCode != 0 {
		slog.Debug("EMA", "result", retCode)
		return nil
	}

	return outReal
}

// HT_TRENDLINE - Ehlers' Hilbert Transform Instantaneous Trendline:
// a smoothed, low-lag overlay whose averaging window adapts to the dominant cycle period measured via Hilbert-transform quadrature (I/Q) analysis of price.
func HT_TRENDLINE(inReal []float64) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inReal))
	)

	if retCode := ht_trendline(
		startIdx,
		endIdx,
		inReal,
		&outBegIdx,
		&outNBElement,
		outReal,
	); retCode != 0 {
		slog.Debug("HT_TRENDLINE", "result", retCode)
		return nil
	}

	return outReal
}

// KAMA - Kaufman Adaptive Moving Average: an EMA whose smoothing factor adapts each bar to an efficiency ratio (directional move vs. total volatility).
// Reacts fast in trends and smooths in ranging markets. Flat KAMA = non-trending/ranging market.
// KAMA tracking price closely = efficient trend.
//
// ER = |price[t] - price[t-period]| / sum(|price[i]-price[i-1]|, last period bars)
// SC = (ER*(2/3 - 2/31) + 2/31)^2
// KAMA[t] = KAMA[t-1] + SC*(price[t] - KAMA[t-1])
//
// Note: A period of 1 performs no smoothing: the output is a copy of the input, consistent with MA(period=1) for every MAType.
// (The natural KAMA math at period 1 would degenerate to a fixed-alpha EMA because the efficiency ratio is always 1, so the copy is made explicit.)
// Allowed since 0.6.5.
func KAMA(inReal []float64, inTimePeriod int) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inReal))
	)

	if retCode := kama(
		startIdx,
		endIdx,
		inReal,
		int32(inTimePeriod),
		&outBegIdx,
		&outNBElement,
		outReal,
	); retCode != 0 {
		slog.Debug("KAMA", "result", retCode)
		return nil
	}

	return outReal
}
