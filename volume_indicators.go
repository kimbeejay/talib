package talib

// AD - Chaikin Accumulation/Distribution Line, a cumulative volume-flow indicator.
// Sums a volume-weighted money-flow multiplier per bar to gauge buying vs. selling pressure.
// Rising line = accumulation (buying pressure); falling = distribution.
//
// MFM = ((close-low) - (high-close)) / (high-low);
// AD_t = AD_{t-1} + MFM_t * volume_t (running sum, seeded at 0)
func AD(inHigh, inLow, inClose, inVolume []float64) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inClose))
	)

	if retCode := ad(
		startIdx,
		endIdx,
		inHigh,
		inLow,
		inClose,
		inVolume,
		&outBegIdx,
		&outNBElement,
		outReal,
	); retCode != 0 {
		return nil
	}

	return outReal
}

// ADOsc - Chaikin A/D Oscillator: the difference between a fast and a slow EMA of the Accumulation/Distribution line.
// Highlights momentum in accumulation/distribution volume flow.
// Positive/rising suggests accumulation; negative/falling suggests distribution.
//
// ad += ((close-low)-(high-close))/(high-low) * volume (only when high>low)
// fastEMA = fastk*ad + (1-fastk)fastEMA, fastk = 2/(optInFastPeriod+1)
// slowEMA = slowkad + (1-slowk)*slowEMA, slowk = 2/(optInSlowPeriod+1)
// ADOSC = fastEMA - slowEMA
//
// @param inHigh []float64
// @param inLow []float64
// @param inClose []float64
// @param inVolume []float64
// @param fastPeriod int Period of the fast A/D EMA; default is 3 (2–100000)
// @param slowPeriod int Period of the slow A/D EMA; default is 10 (2–100000)
// @return []float64 Fast-EMA minus slow-EMA of the A/D line
func ADOsc(inHigh, inLow, inClose, inVolume []float64, fastPeriod, slowPeriod int) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inClose))
	)

	if retCode := adosc(
		startIdx,
		endIdx,
		inHigh,
		inLow,
		inClose,
		inVolume,
		int32(fastPeriod),
		int32(slowPeriod),
		&outBegIdx,
		&outNBElement,
		outReal,
	); retCode != 0 {
		return nil
	}

	return outReal
}

// CMF - Chaikin Money Flow: over a trailing window of optInTimePeriod bars, the sum of each bar's money flow volume divided by the sum of its volume.
// The result is a ratio in [-1, +1].
// A bar's money flow volume is its volume scaled by where the close sat inside the bar's range: a close at the high contributes the full volume,
// a close at the low contributes minus the full volume, and a close at the midpoint contributes nothing.
// Summing that over a window and dividing by the window's volume answers "over these N bars, what share of the traded volume closed near the top of its range?"
//
// Above zero is accumulation, below zero is distribution, and the distance from zero measures conviction.
// Because the divisor is the window's own volume, the output is comparable across instruments and across time in a way a raw accumulation total is not.
//
// Created by Marc Chaikin, who also authored the AD line this shares its per-bar multiplier with.
// CMF is that same multiplier summed over a fixed window and normalised, where AD accumulates it from the start of the series without bound.
//
// t = high[i] - low[i]
// mfv[i] = ((close[i] - low[i]) - (high[i] - close[i])) / t * volume[i], or 0 when t is not positive
// CMF[i] = ( sum_{k=i-N+1..i} mfv[k] ) / ( sum_{k=i-N+1..i} volume[k] ), N = optInTimePeriod
// There is no seeding and no recursion, hence no unstable period. Each output depends only on the N bars in its own window.
//
// Note: The output is the raw ratio in [-1, +1], matching every published definition.
// Some retail platforms display it multiplied by 100; that is a presentation choice, not a different indicator.
//
// Each bar's close is expected to lie within its own [low, high], and its volume to be finite and non-negative.
// A close outside its bar makes the multiplier exceed ±1 and is passed through unclamped, exactly as AD does.
//
// A bar whose high equals its low has no range for the close to sit inside, so it contributes exactly zero money flow volume rather than dividing by zero.
// Its volume still counts toward the divisor.
//
// A window whose volume is entirely zero has no money flow to distribute and reports 0.0.
// Published references are silent here and other implementations divide by zero; TA-Lib does not return NaN from a successful call.
//
// Bars where the low exceeds the high are malformed rather than degenerate, and also contribute zero.
// The default period of 20 follows the original write-up, which describes 20 or 21 bars.
//
// @param inHigh   []float64
// @param inLow    []float64
// @param inClose  []float64
// @param inVolume []float64
// @param period   int       Lookback window; default is 20 (2–100000)
// @return []float64
//
// since TA-Lib 0.8.1
func CMF(inHigh, inLow, inClose, inVolume []float64, period int) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inClose))
	)

	if retCode := cmf(
		startIdx,
		endIdx,
		inHigh,
		inLow,
		inClose,
		inVolume,
		int32(period),
		&outBegIdx,
		&outNBElement,
		outReal,
	); retCode != 0 {
		return nil
	}

	return outReal
}

// NVI - Negative Volume Index: a running cumulative index that changes only on days when
// volume falls versus the prior day, compounding that day's percentage price change.
// The premise is that quiet, low-volume days reflect the actions of well-informed
// "smart money", so NVI is read as a proxy for that cohort's positioning.
//
// NVI[startIdx] = 1000
//
// For each subsequent bar i:
// NVI[i] = NVI[i-1] + ( inVolume[i] < inVolume[i-1] ? ((inClose[i] - inClose[i-1]) / inClose[i-1]) * NVI[i-1] : 0 )
//
// The index carries forward unchanged on bars whose volume did not fall (and on the
// degenerate case of a zero previous close, which would otherwise divide by zero).
//
// since TA-Lib 0.8.1
func NVI(inReal, inVolume []float64) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inReal))
	)

	if retCode := nvi(
		startIdx,
		endIdx,
		inReal,
		inVolume,
		&outBegIdx,
		&outNBElement,
		outReal,
	); retCode != 0 {
		return nil
	}

	return outReal
}

// OBV - On Balance Volume: a running cumulative total of volume, added on up-price bars and subtracted on down-price bars.
// Relates volume flow to price direction.
func OBV(inReal, inVolume []float64) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inReal))
	)

	if retCode := obv(
		startIdx,
		endIdx,
		inReal,
		inVolume,
		&outBegIdx,
		&outNBElement,
		outReal,
	); retCode != 0 {
		return nil
	}

	return outReal
}

// PVI - Positive Volume Index: a running cumulative index that changes only on days when
// volume rises versus the prior day, compounding that day's percentage price change.
// The premise is that active, high-volume days reflect the actions of the
// less-informed "crowd", so PVI is read as a proxy for that cohort's positioning.
//
// PVI[startIdx] = 1000
//
// For each subsequent bar i:
// PVI[i] = PVI[i-1] + ( inVolume[i] > inVolume[i-1] ? ((inClose[i] - inClose[i-1]) / inClose[i-1]) * PVI[i-1] : 0 )
//
// The index carries forward unchanged on bars whose volume did not rise (and on the
// degenerate case of a zero previous close, which would otherwise divide by zero).
//
// since TA-Lib 0.8.1
func PVI(inReal, inVolume []float64) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inReal))
	)

	if retCode := pvi(
		startIdx,
		endIdx,
		inReal,
		inVolume,
		&outBegIdx,
		&outNBElement,
		outReal,
	); retCode != 0 {
		return nil
	}

	return outReal
}

// PVO - Percentage Volume Oscillator: a variation of the Percentage Price Oscillator (PPO, created by Gerald Appel)
// applied to the volume series instead of price. It is the difference between a fast and slow moving average of volume,
// expressed as a percentage of the slow talib.MAType.
// Positive when short-term volume is above its longer-term average (rising participation), negative when below.
// The default periods (12, 26) match MACD and PPO.
//
// PVO = ((fastMA(inVolume) - slowMA(inVolume)) / slowMA(inVolume)) * 100, both MAs of type optInMAType; output = 0 when slowMA == 0
//
// The standard form is exponential with periods 12 and 26 — ((12-day EMA of Volume - 26-day EMA of Volume) / 26-day EMA of Volume) * 100,
// i.e. the talib.PPO/talib.MACD oscillator computed on volume.
// optInMAType therefore defaults to talib.EMA — the moving average Gerald Appel used for the original talib.PPO/talib.MACD;
// pass another type (e.g. talib.MA_SMA) to override.
//
// @param inVolume []float64
// @param fastPeriod int Period of the fast moving average; default is 12 (2–100000)
// @param slowPeriod int Period of the slow moving average; default is 26 (2–100000)
// @param ma MAType Type of moving average to use; default is talib.EMA
// @return []float64 Percentage Volume Oscillator
//
// since TA-Lib 0.8.1
func PVO(inVolume []float64, fastPeriod, slowPeriod int, ma MAType) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inVolume) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inVolume))
	)

	if retCode := pvo(
		startIdx,
		endIdx,
		inVolume,
		int32(fastPeriod),
		int32(slowPeriod),
		int32(ma),
		&outBegIdx,
		&outNBElement,
		outReal,
	); retCode != 0 {
		return nil
	}

	return outReal
}
