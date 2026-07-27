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

// MA - Generic moving-average dispatcher that forwards the job to a concrete MA implementation selected by optInMAType.
// Single uniform interface over all TA-Lib moving averages.
//
// outReal = MA_of_type(optInMAType)(inReal, optInTimePeriod); default type = SMA
//
// Note: A period of 1 performs no smoothing for every MAType: the output is a copy of the input.
// `TA_MAType_DISABLED` bypasses smoothing explicitly, for any period: the output is a copy of the input with a lookback of 0.
// Every function that takes an MAType parameter accepts it.
func MA(inReal []float64, inTimePeriod int, inMAType MAType) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inReal))
	)

	if retCode := ma(
		startIdx,
		endIdx,
		inReal,
		int32(inTimePeriod),
		int32(inMAType),
		&outBegIdx,
		&outNBElement,
		outReal,
	); retCode != 0 {
		slog.Debug("MA", "result", retCode)
		return nil
	}

	return outReal
}

// MAMA - MESA Adaptive Moving Average: an adaptive EMA whose smoothing factor is driven by the dominant-cycle phase rate measured with a Hilbert transform.
// Emits two lines, MAMA and its slower follower FAMA.
// MAMA crossing above FAMA is bullish; crossing below is bearish.
//
// phase = atan(Q1/I1) in degrees; deltaPhase = max(1, prevPhase - phase)
// alpha = max(fastLimit/deltaPhase, slowLimit) if deltaPhase>1 else fastLimit
// MAMA = alpha*price + (1-alpha)*MAMA_prev
// FAMA = (alpha/2)*MAMA + (1-alpha/2)*FAMA_prev
// @param inReal Input data series
// @param inFastLimit Upper bound on the adaptive smoothing factor; default 0.5 (0.01 to 0.99)
// @param inSlowLimit Lower bound on the adaptive smoothing factor; default 0.05 (0.01 to 0.99)
func MAMA(inReal []float64, inFastLimit, inSlowLimit float64) ([]float64, []float64) {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal) - 1)
		outBegIdx    int32
		outNBElement int32
		outMAMA      = make([]float64, len(inReal))
		outFAMA      = make([]float64, len(inReal))
	)

	if retCode := mama(
		startIdx,
		endIdx,
		inReal,
		inFastLimit,
		inSlowLimit,
		&outBegIdx,
		&outNBElement,
		outMAMA,
		outFAMA,
	); retCode != 0 {
		slog.Debug("MAMA", "result", retCode)
		return nil, nil
	}

	return outMAMA, outFAMA
}

// MAVP - Moving average whose period varies per bar, driven by a companion period series.
// For each bar it computes an MA of the selected type over the (clamped) period given by inPeriods.
//
// p_i = clamp((int)inPeriods[startIdx+i], optInMinPeriod, optInMaxPeriod); outReal[i] = MA(inReal, p_i, optInMAType) at bar startIdx+i
//
// Note: Fractional per-bar periods are truncated to whole numbers before being clamped to the minimum and maximum period.
// Period values of 1 perform no smoothing (the bar's output equals its input); the minimum allowed period is 1 since 0.6.5.
// @param inReal Input data series
// @param inPeriods per-bar desired MA period
// @param inMinPeriod Lower clamp for the per-bar period; default 2 (1-100000)
// @param inMaxPeriod Upper clamp for the per-bar period; default 30 (1-100000)
// @param inMAType Type of moving average to compute; default talib.MA_SMA
func MAVP(inReal, inPeriods []float64, inMinPeriod, inMaxPeriod int, inMAType MAType) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inReal))
	)

	if retCode := mavp(
		startIdx,
		endIdx,
		inReal,
		inPeriods,
		int32(inMinPeriod),
		int32(inMaxPeriod),
		int32(inMAType),
		&outBegIdx,
		&outNBElement,
		outReal,
	); retCode != 0 {
		slog.Debug("MAVP", "result", retCode)
		return nil
	}

	return outReal
}

// MidPoint - Midpoint over a period: the average of the highest and lowest input values within the lookback window.
// A single-series overlap smoother (use MIDPRICE for separate high/low price bars).
//
// MIDPOINT = (Highest(inReal, period) + Lowest(inReal, period)) / 2
// @param inReal Input data series
// @param inTimePeriod Lookback window length; default 14 (2-100000)
func MidPoint(inReal []float64, inTimePeriod int) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inReal))
	)

	if retCode := midpoint(
		startIdx,
		endIdx,
		inReal,
		int32(inTimePeriod),
		&outBegIdx,
		&outNBElement,
		outReal,
	); retCode != 0 {
		slog.Debug("MidPoint", "result", retCode)
		return nil
	}

	return outReal
}

// MidPrice - Midpoint of the price range over a rolling window: the average of the highest high and lowest low across the last optInTimePeriod bars.
// An overlap-study line plotted on price.
//
// MIDPRICE = (Highest(High, N) + Lowest(Low, N)) / 2, over the N=optInTimePeriod bars ending at each index
// @param inHigh Input high price series
// @param inLow Input low price series
// @param inTimePeriod Lookback window length; default 14 (2-100000)
func MidPrice(inHigh, inLow []float64, inTimePeriod int) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inHigh) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inHigh))
	)

	if retCode := midprice(
		startIdx,
		endIdx,
		inHigh,
		inLow,
		int32(inTimePeriod),
		&outBegIdx,
		&outNBElement,
		outReal,
	); retCode != 0 {
		slog.Debug("MidPrice", "result", retCode)
		return nil
	}

	return outReal
}

// SAR - Wilder's Parabolic SAR (Stop And Reverse): a trailing stop/reverse level that accelerates toward price via an acceleration factor.
// Signals trend direction and trailing exit points.
// SAR below price = uptrend (long);
// SAR above price = downtrend (short).
// Price crossing SAR flips direction.
//
// SAR_next = SAR + af * (EP - SAR)
// EP = extreme point (highest high in long / lowest low in short); af starts at Acceleration, += Acceleration each new EP, capped at Maximum.
// On penetration: reverse, SAR := prior EP, reset af = Acceleration. SAR clamped each bar so it does not penetrate the prior/current bar's range.
// @param inHigh Input high price series
// @param inLow Input low price series
// @param inAcceleration Step added to the acceleration factor on each new extreme point; default 0.02 (>=0)
// @param inMaximum Ceiling on the acceleration factor; default 0.2 (>=0)
func SAR(inHigh, inLow []float64, inAcceleration, inMaximum float64) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inHigh) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inHigh))
	)

	if retCode := sar(
		startIdx,
		endIdx,
		inHigh,
		inLow,
		inAcceleration,
		inMaximum,
		&outBegIdx,
		&outNBElement,
		outReal,
	); retCode != 0 {
		slog.Debug("SAR", "result", retCode)
		return nil
	}

	return outReal
}

// SARExt - Extended Parabolic SAR (stop and reverse) giving the caller full control over the initial state and separate acceleration factors for long and short positions.
// Unlike SAR, it returns negative values while short so reversals are distinguishable.
// Sign flip of the output marks a trend reversal (positive=long stop, negative=short stop).
//
// SAR_next = SAR + AF*(EP - SAR), then clamped within the prior and current bar's range.
// On penetration, reverse: set SAR=EP (clamped), reset AF to its Init value, EP=extreme of the new direction.
// Output is +SAR when long, -SAR when short. On reversal an optional offset is applied: long->short SAR*(1+offset), short->long SAR*(1-offset).
// @param inHigh Input high price series
// @param inLow Input low price series
// @param inStartValue Initial SAR/direction: 0 auto, >0 start long at value, <0 start short at |value|; default 0 (any real number)
// @param inOffsetOnReverse Fractional offset applied to the stop on each reversal; default 0 (>=0)
// @param inAccelerationInitLong Initial acceleration factor when long; default 0.02 (>=0)
// @param inAccelerationLong AF increment per new long extreme; default 0.02 (>=0)
// @param inAccelerationMaxLong Cap on the long acceleration factor; default 0.2 (>=0)
// @param inAccelerationInitShort Initial acceleration factor when short; default 0.02 (>=0)
// @param inAccelerationShort AF increment per new short extreme; default 0.02 (>=0)
// @param inAccelerationMaxShort Cap on the short acceleration factor; default 0.2 (>=0)
func SARExt(
	inHigh, inLow []float64,
	inStartValue, inOffsetOnReverse,
	inAccelerationInitLong, inAccelerationLong, inAccelerationMaxLong,
	inAccelerationInitShort, inAccelerationShort, inAccelerationMaxShort float64,
) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inHigh) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inHigh))
	)

	if retCode := sarext(
		startIdx,
		endIdx,
		inHigh,
		inLow,
		inStartValue,
		inOffsetOnReverse,
		inAccelerationInitLong,
		inAccelerationLong,
		inAccelerationMaxLong,
		inAccelerationInitShort,
		inAccelerationShort,
		inAccelerationMaxShort,
		&outBegIdx,
		&outNBElement,
		outReal,
	); retCode != 0 {
		slog.Debug("SARExt", "result", retCode)
		return nil
	}

	return outReal
}

// SMA - Simple Moving Average: the unweighted arithmetic mean of the last N input values. Used to smooth a series.
//
// SMA_t = (1/N) * sum_{i=t-N+1}^{t} inReal_i
//
// Note: A period of 1 performs no smoothing: the output is a copy of the input. Allowed since 0.6.5 (issues #48/#59).
func SMA(inReal []float64, inTimePeriod int) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inReal))
	)

	if retCode := sma(
		startIdx,
		endIdx,
		inReal,
		int32(inTimePeriod),
		&outBegIdx,
		&outNBElement,
		outReal,
	); retCode != 0 {
		slog.Debug("SMA", "result", retCode)
		return nil
	}

	return outReal
}

// T3 - Tillson's T3: a low-lag moving average built from six chained EMAs, combined via volume-factor-weighted coefficients.
// Not the same as EMA3, despite both being called "triple EMA".
//
// k = 2/(period+1); e1=EMA(x), e2=EMA(e1), ... e6=EMA(e5) (six chained EMAs).
// v = vFactor: c1 = -v^3; c2 = 3(v^2 - c1); c3 = -6v^2 - 3(v - c1); c4 = 1 + 3v - c1 + 3v^2.
// T3 = c1e6 + c2e5 + c3e4 + c4e3
//
// Note: A period of 1 performs no smoothing: the output is a copy of the input. Allowed since 0.6.5 (issues #48/#59).
// @param inReal Input data series
// @param inTimePeriod EMA period for each of the six stages; default 5 (2-100000)
// @param inVFactor Volume factor weighting the coefficients (0 = plain triple EMA, higher = more DEMA-like sharpening); default 0.7 (0-1)
func T3(inReal []float64, inTimePeriod int, inVFactor float64) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inReal))
	)

	if retCode := t3(
		startIdx,
		endIdx,
		inReal,
		int32(inTimePeriod),
		inVFactor,
		&outBegIdx,
		&outNBElement,
		outReal,
	); retCode != 0 {
		slog.Debug("T3", "result", retCode)
		return nil
	}

	return outReal
}

// TEMA - Triple Exponential Moving Average: a smoothed price overlay built from three successively-applied EMAs to reduce lag versus a plain EMA.
// Distinct from EMA3, also called "triple EMA" in the literature.
//
// EMA1=EMA(t,period); EMA2=EMA(EMA1,period); EMA3=EMA(EMA2,period); TEMA = 3EMA1 - 3EMA2 + EMA3
//
// Note: A period of 1 performs no smoothing: the output is a copy of the input. Allowed since 0.6.5 (issues #48/#59).
func TEMA(inReal []float64, inTimePeriod int) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inReal))
	)

	if retCode := tema(
		startIdx,
		endIdx,
		inReal,
		int32(inTimePeriod),
		&outBegIdx,
		&outNBElement,
		outReal,
	); retCode != 0 {
		slog.Debug("TEMA", "result", retCode)
		return nil
	}

	return outReal
}

// TRIMA - Triangular Moving Average: a double-smoothed moving average that weights prices toward the middle of the window most heavily.
// Equivalent to an SMA of an SMA, computed here via an incremental triangular-weighted running numerator.
//
// Weights rise then fall (4-period: (1a+2b+2c+1d)/6; 5-period: (1a+2b+3c+2d+1e)/9).
// With n = period>>1: odd divides by (n+1)^2, even by n(n+1).
// Equivalent to odd: SMA(SMA(x,(period+1)/2),(period+1)/2); even: SMA(SMA(x,period/2),period/2+1).
//
// Note: Follows the generally accepted (Metastock) definition rather than the TradeStation variant.
// A period of 1 performs no smoothing: the output is a copy of the input. Allowed since 0.6.5 (issues #48/#59).
func TRIMA(inReal []float64, inTimePeriod int) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inReal))
	)

	if retCode := trima(
		startIdx,
		endIdx,
		inReal,
		int32(inTimePeriod),
		&outBegIdx,
		&outNBElement,
		outReal,
	); retCode != 0 {
		slog.Debug("TRIMA", "result", retCode)
		return nil
	}

	return outReal
}

// WMA - Linearly weighted moving average: each of the last N prices is weighted by its position,
// oldest getting weight 1 and newest weight N. Smooths price while emphasizing recent bars.
//
// WMA = ( sum_{k=1..N} k * P_k ) / (N(N+1)/2), where P_N is the most recent bar
//
// Note: A period of 1 performs no smoothing: the output is a copy of the input. Allowed since 0.6.5 (issues #48/#59).
func WMA(inReal []float64, inTimePeriod int) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inReal))
	)

	if retCode := wma(
		startIdx,
		endIdx,
		inReal,
		int32(inTimePeriod),
		&outBegIdx,
		&outNBElement,
		outReal,
	); retCode != 0 {
		slog.Debug("WMA", "result", retCode)
		return nil
	}

	return outReal
}
