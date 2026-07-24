package talib

import "log/slog"

// ADX - Wilder's Average Directional Movement Index, a smoothed measure of trend strength derived from the directional indicators (+DI/-DI).
// Quantifies how strongly a market is trending, regardless of direction.
// Higher values indicate a stronger trend (a common convention treats >25 as trending); says nothing about direction.
//
// +DI = 100*(+DM_p/TR_p), -DI = 100*(-DM_p/TR_p); DX = 100*|(-DI)-(+DI)| / ((-DI)+(+DI));
// first ADX = mean of the first period DX; then ADX = (prevADX*(period-1) + DX)/period.
// +DM_p/-DM_p/TR_p use Wilder smoothing: X = X - X/period + today's one-bar value.
//
// Note: Wilder's original integer rounding is not applied.
func ADX(inHigh, inLow, inClose []float64, optInTimePeriod int) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inClose))
	)

	if retCode := adx(
		startIdx,
		endIdx,
		inHigh,
		inLow,
		inClose,
		int32(optInTimePeriod),
		&outBegIdx,
		&outNBElement,
		outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("ADX", "result", retCode)
		return nil
	}

	return outReal
}

// ADXR - Smoothed variant of ADX: the average of the current ADX value and the ADX value from (period-1) bars earlier.
// Further damps ADX to gauge trend strength. Higher values mean a stronger trend; smoother and more lagging than ADX.
//
// ADXR[i] = (ADX[i] + ADX[i-(period-1)]) / 2
// Note: Wilder's original integer rounding is not applied (unreliable when values are near 1).
func ADXR(inHigh, inLow, inClose []float64, optInTimePeriod int) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inClose))
	)

	if retCode := adxr(
		startIdx,
		endIdx,
		inHigh,
		inLow,
		inClose,
		int32(optInTimePeriod),
		&outBegIdx,
		&outNBElement,
		outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("ADXR", "result", retCode)
		return nil
	}

	return outReal
}

// APO - Absolute Price Oscillator: the difference between a fast and a slow moving average of the input, in price units.
// Measures short- vs long-term momentum. Positive when fast MA > slow MA (upward momentum); negative otherwise.
//
// APO = fastMA - slowMA
// fastMA = MA(inReal, optInFastPeriod, optInMAType)
// slowMA = MA(inReal, optInSlowPeriod, optInMAType)
//
// The standard form is exponential — APO with EMA and periods 12/26 is the fast-minus-slow EMA construction underlying the MACD (in price units).
// optInMAType therefore defaults to EMA — the moving average Gerald Appel used for the original MACD;
// pass another type (e.g. talib.MATypeSMA) to override.
func APO(inReal []float64, optInFastPeriod, optInSlowPeriod int, optInMAType MAType) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inReal))
	)

	if retCode := apo(
		startIdx,
		endIdx,
		inReal,
		int32(optInFastPeriod),
		int32(optInSlowPeriod),
		int32(optInMAType),
		&outBegIdx,
		&outNBElement,
		outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("APO", "result", retCode)
		return nil
	}

	return outReal
}

// AROON - Aroon reports how recently the highest high and lowest low occurred within a rolling window of length optInTimePeriod, as two 0-100 oscillators.
// Indicates trend strength and direction. Up near 100 = a very recent new high (strong uptrend);
// Down near 100 = a very recent new low. Up/Down crossovers signal trend shifts.
//
// Up = 100*(period-(today-highestIdx))/period
// Down = 100*(period-(today-lowestIdx))/period
// where highestIdx/lowestIdx index the highest high / lowest low over the window [today-period .. today].
func AROON(inHigh, inLow []float64, optInTimePeriod int) ([]float64, []float64) {
	var (
		startIdx     int32
		endIdx       = int32(len(inHigh) - 1)
		outBegIdx    int32
		outNBElement int32
		outAroonUp   = make([]float64, len(inHigh))
		outAroonDown = make([]float64, len(inLow))
	)

	if retCode := aroon(
		startIdx,
		endIdx,
		inHigh,
		inLow,
		int32(optInTimePeriod),
		&outBegIdx,
		&outNBElement,
		outAroonUp,
		outAroonDown,
	); SUCCESS != taResult(retCode) {
		slog.Debug("AROON", "result", retCode)
		return nil, nil
	}

	return outAroonUp, outAroonDown
}

// AROONOSC - Aroon Oscillator: AroonUp minus AroonDown over a lookback window.
// Measures trend direction and strength on a -100..+100 scale.
// Positive when the high is more recent than the low (up-trend); negative when the low is more recent (down-trend).
//
// factor = 100 / optInTimePeriod
// AroonUp = factor * (period - (today - highestIdx))
// AroonDown = factor * (period - (today - lowestIdx))
// AroonOsc = AroonUp - AroonDown = factor * (highestIdx - lowestIdx)
// highestIdx/lowestIdx = bar index of the highest high / lowest low in the last (period+1) bars.
func AROONOSC(inHigh, inLow []float64, optInTimePeriod int) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inHigh) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inHigh))
	)

	if retCode := aroonosc(
		startIdx,
		endIdx,
		inHigh,
		inLow,
		int32(optInTimePeriod),
		&outBegIdx,
		&outNBElement,
		outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("AROONOSC", "result", retCode)
		return nil
	}

	return outReal
}

// BOP - Balance Of Power compares where the close sits relative to the open, normalized by the bar's high-low range.
// A per-bar oscillator with no smoothing. Positive: close above open (buyers dominated); negative: sellers dominated.
//
// BOP = (Close - Open) / (High - Low)
func BOP(inOpen, inHigh, inLow, inClose []float64) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inClose))
	)

	if retCode := bop(
		startIdx,
		endIdx,
		inOpen,
		inHigh,
		inLow,
		inClose,
		&outBegIdx,
		&outNBElement,
		outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("BOP", "result", retCode)
		return nil
	}

	return outReal
}

// CCI - Commodity Channel Index: measures the current typical price relative to its simple moving average, scaled by mean absolute deviation.
// Momentum oscillator flagging overbought/oversold extremes. CCI > +100 overbought; CCI < -100 oversold.
//
// TP_i = (High_i + Low_i + Close_i)/3
// SMA = (1/N) * sum(TP over N bars)
// meanDev = (1/N) * sum(|TP - SMA| over N bars)
// CCI = (TP_last - SMA) / (0.015 * meanDev)
func CCI(inHigh, inLow, inClose []float64, optInTimePeriod int) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inClose))
	)

	if retCode := cci(
		startIdx,
		endIdx,
		inHigh,
		inLow,
		inClose,
		int32(optInTimePeriod),
		&outBegIdx,
		&outNBElement,
		outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("CCI", "result", retCode)
		return nil
	}

	return outReal
}

// CMO - Chande Momentum Oscillator: bounded momentum measure from Wilder-smoothed average up-moves and down-moves.
// Identical to RSI except the numerator uses (gain-loss) instead of gain.
// Bounded in [-100,+100]; positive = net upward momentum, negative = net downward.
//
// d = P[t]-P[t-1]; over the initial period accumulate gain = sum of positive d, loss = sum of -d for negative d.
// Wilder-smooth each: prevGain = (prevGain*(period-1) + gain_today)/period (same for loss).
// CMO = 100 * (prevGain-prevLoss)/(prevGain+prevLoss); 0 when prevGain+prevLoss == 0.
//
// Note: Gains and losses are smoothed with Wilder's method (as in RSI) rather than the simple period sums of Chande's original definition.
func CMO(inReal []float64, optInTimePeriod int) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inReal))
	)

	if retCode := cmo(
		startIdx,
		endIdx,
		inReal,
		int32(optInTimePeriod),
		&outBegIdx,
		&outNBElement,
		outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("CMO", "result", retCode)
		return nil
	}

	return outReal
}

// DX - Wilder's Directional Movement Index: the normalized spread between +DI and -DI.
// Measures the strength of directional (trending) movement, irrespective of direction.
// Higher DX = stronger trend (either direction); low DX = ranging market.
//
// Seed +DM14, -DM14, TR14 as sums of the first (period-1) one-period values, then Wilder-smooth each: X = X - X/period + today.
// +DI = 100*(+DM14/TR14), -DI = 100*(-DM14/TR14). DX = 100 * |(-DI) - (+DI)| / ((-DI) + (+DI)).
//
// Note: Wilder's original integer rounding is not applied (it can be unreliable when values are near 1).
// When +DI and -DI sum to zero the value is undefined; the previous bar's DX is carried forward instead (the first such bar outputs zero).
func DX(inHigh, inLow, inClose []float64, optInTimePeriod int) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inClose))
	)

	if retCode := dx(
		startIdx,
		endIdx,
		inHigh,
		inLow,
		inClose,
		int32(optInTimePeriod),
		&outBegIdx,
		&outNBElement,
		outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("DX", "result", retCode)
		return nil
	}

	return outReal
}

// IMI - Intraday Momentum Index: an RSI-like 0-100 oscillator built from the open-to-close body of each bar.
// Over a rolling window it ratios cumulative up-body moves against total up+down body moves.
//
// upsum = Σ(close-open) for bars with close>open;
// downsum = Σ(open-close) for bars with close<=open, over window [i-lookback, i];
// IMI = 100 * upsum/(upsum+downsum)
func IMI(inOpen, inClose []float64, optInTimePeriod int) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inClose))
	)

	if retCode := imi(
		startIdx,
		endIdx,
		inOpen,
		inClose,
		int32(optInTimePeriod),
		&outBegIdx,
		&outNBElement,
		outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("IMI", "result", retCode)
		return nil
	}

	return outReal
}

// MACD - Moving Average Convergence/Divergence: the difference between a fast and a slow EMA of the input, plus an EMA-smoothed signal line and their histogram.
// MACD crossing its signal line and histogram sign changes flag momentum shifts.
//
// MACD = EMA_fast - EMA_slow; Signal = EMA(MACD, signalPeriod); Hist = MACD - Signal
//
// Note: If the slow period is set smaller than the fast period, the two are swapped so the slow EMA is always the longer one.
// A signal period of 1 disables signal-line smoothing: the signal equals the MACD line and the histogram is zero.
// Before 0.6.5 this parameter value produced misaligned output (issues #48/#59).
func MACD(inReal []float64, optInFastPeriod, optInSlowPeriod, optInSignalPeriod int) ([]float64, []float64, []float64) {
	var (
		startIdx      int32
		endIdx        = int32(len(inReal) - 1)
		outBegIdx     int32
		outNBElement  int32
		outMACD       = make([]float64, len(inReal))
		outMACDSignal = make([]float64, len(inReal))
		outMACDHist   = make([]float64, len(inReal))
	)

	if retCode := macd(
		startIdx,
		endIdx,
		inReal,
		int32(optInFastPeriod),
		int32(optInSlowPeriod),
		int32(optInSignalPeriod),
		&outBegIdx,
		&outNBElement,
		outMACD,
		outMACDSignal,
		outMACDHist,
	); SUCCESS != taResult(retCode) {
		slog.Debug("MACD", "result", retCode)
		return nil, nil, nil
	}

	return outMACD, outMACDSignal, outMACDHist
}

// MACDEXT - MACD variant where the fast, slow, and signal moving averages each use a user-selectable MA type.
// Outputs the MACD line, its signal line, and their difference (histogram).
// Hist sign change (MACD crossing its signal line) flags momentum shifts.
//
// MACD = MA_fast(inReal) - MA_slow(inReal)
// Signal = MA_signal(MACD)
// Hist = MACD - Signal
// (each MA_* uses its own MA type and period)
//
// Note: If the slow period is set smaller than the fast period, the fast and slow periods and their MA types are swapped so the slow moving average is always the longer one.
// A signal period of 1 disables signal-line smoothing for every signal MAType: the signal equals the MACD line and the histogram is zero.
func MACDEXT(
	inReal []float64,
	optInFastPeriod int, optInFastMAType MAType,
	optInSlowPeriod int, optInSlowMAType MAType,
	optInSignalPeriod int, optInSignalMAType MAType,
) ([]float64, []float64, []float64) {
	var (
		startIdx      int32
		endIdx        = int32(len(inReal) - 1)
		outBegIdx     int32
		outNBElement  int32
		outMACD       = make([]float64, len(inReal))
		outMACDSignal = make([]float64, len(inReal))
		outMACDHist   = make([]float64, len(inReal))
	)

	if retCode := macdext(
		startIdx,
		endIdx,
		inReal,
		int32(optInFastPeriod),
		int32(optInFastMAType),
		int32(optInSlowPeriod),
		int32(optInSlowMAType),
		int32(optInSignalPeriod),
		int32(optInSignalMAType),
		&outBegIdx,
		&outNBElement,
		outMACD,
		outMACDSignal,
		outMACDHist,
	); SUCCESS != taResult(retCode) {
		slog.Debug("MACDEXT", "result", retCode)
		return nil, nil, nil
	}

	return outMACD, outMACDSignal, outMACDHist
}

// MACDFIX - MACD with the fast/slow EMAs fixed to the classic 12/26 periods (with the classic fixed smoothing factors 0.15 and 0.075), exposing only the signal period.
// Signal-line crossovers and histogram sign flag momentum shifts.
//
// MACD = EMA_12 - EMA_26 (fixed k: 0.15 for 12, 0.075 for 26)
// Signal = EMA(MACD, signalPeriod), k = 2/(signalPeriod+1)
// Hist = MACD - Signal
//
// Note: A signal period of 1 disables signal-line smoothing: the signal equals the MACD line and the histogram is zero.
// Before 0.6.5 this parameter value produced misaligned output (issues #48/#59).
func MACDFIX(inReal []float64, optInSignalPeriod int) ([]float64, []float64, []float64) {
	var (
		startIdx      int32
		endIdx        = int32(len(inReal) - 1)
		outBegIdx     int32
		outNBElement  int32
		outMACD       = make([]float64, len(inReal))
		outMACDSignal = make([]float64, len(inReal))
		outMACDHist   = make([]float64, len(inReal))
	)

	if retCode := macdfix(
		startIdx,
		endIdx,
		inReal,
		int32(optInSignalPeriod),
		&outBegIdx,
		&outNBElement,
		outMACD,
		outMACDSignal,
		outMACDHist,
	); SUCCESS != taResult(retCode) {
		slog.Debug("MACDFIX", "result", retCode)
		return nil, nil, nil
	}

	return outMACD, outMACDSignal, outMACDHist
}

// MFI - Money Flow Index: a volume-weighted momentum oscillator (0-100) comparing positive vs negative money flow over a period.
// A volume-based analog of RSI. >80 overbought, <20 oversold.
//
// TP = (High+Low+Close)/3; MF = TP*Volume, classed positive if TP>prevTP, negative if TP<prevTP, neither if equal.
// MFI = 100 * posSumMF/(posSumMF+negSumMF).
//
// Note: When the typical price is unchanged from the prior bar, that bar's money flow is counted as neither positive nor negative.
func MFI(inHigh, inLow, inClose, inVolume []float64, optInTimePeriod int) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inClose))
	)

	if retCode := mfi(
		startIdx,
		endIdx,
		inHigh,
		inLow,
		inClose,
		inVolume,
		int32(optInTimePeriod),
		&outBegIdx,
		&outNBElement,
		outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("MFI", "result", retCode)
		return nil
	}

	return outReal
}

// MOM - Momentum: current price minus the price optInTimePeriod bars ago.
// The absolute (unnormalized) rate of change. Positive = price rose over the period, negative = fell; centered at zero.
//
// MOM[i] = inReal[i] - inReal[i - optInTimePeriod]
func MOM(inReal []float64, optInTimePeriod int) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inReal))
	)

	if retCode := mom(
		startIdx,
		endIdx,
		inReal,
		int32(optInTimePeriod),
		&outBegIdx,
		&outNBElement,
		outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("MOM", "result", retCode)
		return nil
	}

	return outReal
}

// MINUS_DI - Wilder's Minus Directional Indicator: the Wilder-smoothed downward directional movement (-DM) normalized by smoothed True Range.
// Measures the strength of downward price movement. Higher -DI indicates a stronger downtrend; compared against +DI to gauge directional dominance.
//
// -DM1 = (prevLow - low) if (prevLow-low)>0 and (high-prevHigh)<(prevLow-low), else 0.
// Seed -DM/TR = sum of first (period-1) -DM1/TR1, then Wilder-smooth each: X = X - X/period + today.
// -DI = 100 * (-DM / TR); TR from ta_true_range.
// If period<=1: -DI1 = -DM1/TR1 (no ×100).
//
// Note: Wilder's original integer rounding is not applied (it was removed as unreliable when values are near 1).
func MINUS_DI(inHigh, inLow, inClose []float64, optInTimePeriod int) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inClose))
	)

	if retCode := minus_di(
		startIdx,
		endIdx,
		inHigh,
		inLow,
		inClose,
		int32(optInTimePeriod),
		&outBegIdx,
		&outNBElement,
		outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("MINUS_DI", "result", retCode)
		return nil
	}

	return outReal
}

// MINUS_DM - Minus Directional Movement, the downward component of Wilder's directional movement system.
// Measures Wilder-smoothed downward price motion over the period. Higher -DM indicates stronger downward directional movement.
//
// diffP = high - prevHigh; diffM = prevLow - low
// -DM1 = diffM if (diffM > 0 and diffP < diffM) else 0
// period<=1: output raw -DM1 per bar.
// period>1: seed = sum of first (period-1) -DM1; then Wilder smooth each bar:
// -DM = prevMinusDM - prevMinusDM/period (+ -DM1 when the bar qualifies)
func MINUS_DM(inHigh, inLow []float64, optInTimePeriod int) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inLow) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inLow))
	)

	if retCode := minus_dm(
		startIdx,
		endIdx,
		inHigh,
		inLow,
		int32(optInTimePeriod),
		&outBegIdx,
		&outNBElement,
		outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("MINUS_DM", "result", retCode)
		return nil
	}

	return outReal
}

// PLUS_DI - Plus Directional Indicator: the Wilder-smoothed positive directional movement expressed as a percentage of the true range.
// Measures the strength of upward price movement. Rising +DI signals strengthening upward direction; compared against MINUS_DI to judge trend direction.
//
// +DM1 = (H-Hprev) if (H-Hprev) > 0 and (H-Hprev) > (Lprev-L), else 0.
// TR1 = true range = max(H-L, |H-Cprev|, |L-Cprev|).
// Seed +DM/TR = sum of first (period-1) one-period values; then Wilder smooth: X = X - X/period + X1.
// +DI = 100 * (+DM / TR); if TR = 0, +DI = 0.
// When period <= 1: +DI = +DM1 / TR1 (no *100).
//
// Note: Wilder's original integer rounding of intermediate values is not applied (it was unreliable when values are near 1).
func PLUS_DI(inHigh, inLow, inClose []float64, optInTimePeriod int) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inClose))
	)

	if retCode := plus_di(
		startIdx,
		endIdx,
		inHigh,
		inLow,
		inClose,
		int32(optInTimePeriod),
		&outBegIdx,
		&outNBElement,
		outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("PLUS_DI", "result", retCode)
		return nil
	}

	return outReal
}

// PLUS_DM - Plus Directional Movement: the Wilder-smoothed accumulation of upward directional movement (+DM1).
// A component of the Directional Movement System used to build +DI/DX/ADX.
//
// +DM1 = (high - prevHigh) if (high-prevHigh) > 0 and > (prevLow-low), else 0.
// period<=1: output = +DM1 per bar.
// period>1: seed = sum of first (period-1) +DM1; then Wilder smoothing:
// +DM = prevPlusDM - prevPlusDM/period + +DM1(today)
func PLUS_DM(inHigh, inLow []float64, optInTimePeriod int) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inLow) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inLow))
	)

	if retCode := plus_dm(
		startIdx,
		endIdx,
		inHigh,
		inLow,
		int32(optInTimePeriod),
		&outBegIdx,
		&outNBElement,
		outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("PLUS_DM", "result", retCode)
		return nil
	}

	return outReal
}

// PPO - Percentage Price Oscillator: the difference between a fast and slow moving average expressed as a percentage of the slow MA.
// A normalized (scale-invariant) variant of APO. Positive when the fast MA is above the slow MA (upward momentum), negative otherwise; magnitude is the % deviation.
//
// PPO = ((fastMA(inReal) - slowMA(inReal)) / slowMA(inReal)) * 100, both MAs of type optInMAType; output = 0 when slowMA == 0
// The standard form is exponential with periods 12 and 26 — ((12-day EMA - 26-day EMA) / 26-day EMA) * 100,
// i.e. the MACD oscillator expressed as a percentage. optInMAType therefore defaults to EMA — the moving average Gerald Appel used for the original PPO/MACD;
// pass another type (e.g. talib.SMA) to override.
func PPO(inReal []float64, optInFastPeriod, optInSlowPeriod int, optInMAType MAType) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inReal))
	)

	if retCode := ppo(
		startIdx,
		endIdx,
		inReal,
		int32(optInFastPeriod),
		int32(optInSlowPeriod),
		int32(optInMAType),
		&outBegIdx,
		&outNBElement,
		outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("PPO", "result", retCode)
		return nil
	}

	return outReal
}

// ROC - Rate-of-change momentum oscillator: the percent change of price versus the price optInTimePeriod bars earlier.
// Centered at zero with positive and negative values. Positive when price rose over the period, negative when it fell; magnitude scales the move.
//
// ROC = ((price / prevPrice) - 1) * 100, where prevPrice = inReal[i - optInTimePeriod]
func ROC(inReal []float64, optInTimePeriod int) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inReal))
	)

	if retCode := roc(
		startIdx,
		endIdx,
		inReal,
		int32(optInTimePeriod),
		&outBegIdx,
		&outNBElement,
		outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("ROC", "result", retCode)
		return nil
	}

	return outReal
}

// ROCP - Rate of change expressed as a fraction of the price optInTimePeriod bars ago.
// Normalized and centered at zero (positive or negative). >0 rising vs N bars ago, <0 falling; equals ROC/100.
//
// ROCP = (price - prevPrice) / prevPrice, prevPrice = inReal[i - optInTimePeriod]
func ROCP(inReal []float64, optInTimePeriod int) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inReal))
	)

	if retCode := rocp(
		startIdx,
		endIdx,
		inReal,
		int32(optInTimePeriod),
		&outBegIdx,
		&outNBElement,
		outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("ROCP", "result", retCode)
		return nil
	}

	return outReal
}

// ROCR - Rate of Change Ratio: the ratio of the current price to the price optInTimePeriod bars ago.
// A momentum measure centered at 1. Always positive, centered at 1: >1 rising, <1 falling.
//
// ROCR = price / price[t - optInTimePeriod]
func ROCR(inReal []float64, optInTimePeriod int) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inReal))
	)

	if retCode := rocr(
		startIdx,
		endIdx,
		inReal,
		int32(optInTimePeriod),
		&outBegIdx,
		&outNBElement,
		outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("ROCR", "result", retCode)
		return nil
	}

	return outReal
}

// ROCR100 - Rate-of-change ratio scaled by 100: current price as a percentage of the price optInTimePeriod bars ago.
// Momentum measure centered at 100 and always positive. Above 100 = price rose vs n bars ago; below 100 = price fell.
//
// ROCR100 = (price / price[t - optInTimePeriod]) * 100
func ROCR100(inReal []float64, optInTimePeriod int) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inReal))
	)

	if retCode := rocr100(
		startIdx,
		endIdx,
		inReal,
		int32(optInTimePeriod),
		&outBegIdx,
		&outNBElement,
		outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("ROCR100", "result", retCode)
		return nil
	}

	return outReal
}

// RSI - Wilder's Relative Strength Index, a momentum oscillator bounded 0-100 from the ratio of average gains to average losses over the period.
// Used to gauge overbought/oversold conditions. >70 overbought, <30 oversold.
//
// Seed avgGain/avgLoss as the sum of the first (period-1) gains/losses, then Wilder-smooth each: X = X - X/period + today.
// RS = avgGain / avgLoss; RSI = 100 - (100 / (1 + RS)); when avgLoss == 0, RSI = 100; when avgGain == 0, RSI = 0.
//
// Note: Wilder's original integer rounding is not applied (it was unreliable when values are near 1).
func RSI(inReal []float64, optInTimePeriod int) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inReal))
	)

	if retCode := rsi(
		startIdx,
		endIdx,
		inReal,
		int32(optInTimePeriod),
		&outBegIdx,
		&outNBElement,
		outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("RSI", "result", retCode)
		return nil
	}

	return outReal
}

// STOCH - Slow Stochastic oscillator: locates the close within the high-low range over a lookback period, then double-smooths it.
// Returns the Slow-%K and Slow-%D lines. SlowK/SlowD > 80 overbought, < 20 oversold; %K crossing %D signals momentum shifts.
//
// FastK = 100*(Close - LL_n)/(HH_n - LL_n), n = FastK_Period (LL/HH = lowest low / highest high over n)
// SlowK = MA(FastK, SlowK_Period, SlowK_MAType)
// SlowD = MA(SlowK, SlowD_Period, SlowD_MAType)
//
// Note: When the high-low range over the window is zero, the raw stochastic is set to 0 instead of being undefined.
func STOCH(
	inHigh, inLow, inClose []float64,
	optInFastKPeriod int,
	optInSlowKPeriod int,
	optInSlowKMAType MAType,
	optInSlowDPeriod int,
	optInSlowDMAType MAType,
) ([]float64, []float64) {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outSlowK     = make([]float64, len(inClose))
		outSlowD     = make([]float64, len(inClose))
	)

	if retCode := stoch(
		startIdx,
		endIdx,
		inHigh,
		inLow,
		inClose,
		int32(optInFastKPeriod),
		int32(optInSlowKPeriod),
		int32(optInSlowKMAType),
		int32(optInSlowDPeriod),
		int32(optInSlowDMAType),
		&outBegIdx,
		&outNBElement,
		outSlowK,
		outSlowD,
	); SUCCESS != taResult(retCode) {
		slog.Debug("STOCH", "result", retCode)
		return nil, nil
	}

	return outSlowK, outSlowD
}

// STOCHF - Fast Stochastic Oscillator: the raw %K line and its moving-average-smoothed %D line.
// Unlike STOCH (which slows both lines), STOCHF returns the unsmoothed FastK and FastD. Oscillates 0-100; >80 overbought, <20 oversold.
//
// FastK = 100 * (Close - LowestLow) / (HighestHigh - LowestLow), over the last FastK_Period bars (incl. today)
// FastD = MA(FastK, FastD_Period, FastD_MAType)
//
// Note: When the high-low range over the window is zero, %K is set to 0 instead of being undefined.
func STOCHF(
	inHigh, inLow, inClose []float64,
	optInFastKPeriod int,
	optInFastDPeriod int,
	optInFastDMAType MAType,
) ([]float64, []float64) {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outFastK     = make([]float64, len(inClose))
		outFastD     = make([]float64, len(inClose))
	)

	if retCode := stochf(
		startIdx,
		endIdx,
		inHigh,
		inLow,
		inClose,
		int32(optInFastKPeriod),
		int32(optInFastDPeriod),
		int32(optInFastDMAType),
		&outBegIdx,
		&outNBElement,
		outFastK,
		outFastD,
	); SUCCESS != taResult(retCode) {
		slog.Debug("STOCHF", "result", retCode)
		return nil, nil
	}

	return outFastK, outFastD
}

// STOCHRSI - Applies the Fast Stochastic (STOCHF) oscillator to an RSI series instead of price,
// measuring where RSI sits within its recent min/max range.
// Oscillates 0-100; high = RSI near its recent top, low = near its recent bottom.
//
// rsi = RSI(inReal, optInTimePeriod)
// FastK = 100 * (rsi_t - min(rsi, FastK_Period)) / (max(rsi, FastK_Period) - min(rsi, FastK_Period))
// FastD = MA(FastK, FastD_Period, FastD_MAType)
//
// Note: To reproduce the original article's unsmoothed Stochastic RSI, set the RSI period equal to the %K period and read the raw %K output.
// When the RSI's recent range is zero, %K is set to 0 instead of being undefined.
func STOCHRSI(
	inReal []float64,
	optInTimePeriod int,
	optInFastKPeriod int,
	optInFastDPeriod int,
	optInFastDMAType MAType,
) ([]float64, []float64) {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal) - 1)
		outBegIdx    int32
		outNBElement int32
		outFastK     = make([]float64, len(inReal))
		outFastD     = make([]float64, len(inReal))
	)

	if retCode := stochrsi(
		startIdx,
		endIdx,
		inReal,
		int32(optInTimePeriod),
		int32(optInFastKPeriod),
		int32(optInFastDPeriod),
		int32(optInFastDMAType),
		&outBegIdx,
		&outNBElement,
		outFastK,
		outFastD,
	); SUCCESS != taResult(retCode) {
		slog.Debug("STOCHRSI", "result", retCode)
		return nil, nil
	}

	return outFastK, outFastD
}

// TRIX - 1-day Rate-Of-Change of a triple-smoothed EMA of the input.
// Momentum oscillator that filters out price moves shorter than the chosen period.
// Oscillates around zero; sign, zero-crossings and slope signal momentum direction.
//
// E1 = EMA(inReal, n); E2 = EMA(E1, n); E3 = EMA(E2, n); TRIX = ROC_1(E3) = 100 * (E3_today/E3_yesterday - 1)
//
// Note: The final rate-of-change step yields 0 when the previous smoothed value is exactly zero, rather than being undefined.
func TRIX(inReal []float64, optInTimePeriod int) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inReal))
	)

	if retCode := trix(
		startIdx,
		endIdx,
		inReal,
		int32(optInTimePeriod),
		&outBegIdx,
		&outNBElement,
		outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("TRIX", "result", retCode)
		return nil
	}

	return outReal
}

// ULTOSC - Ultimate Oscillator: momentum indicator combining buying-pressure/true-range ratios over three time periods into one 0-100 weighted average.
// Blends short-, medium-, and long-term momentum to damp single-period noise. Ranges 0-100; conventionally >70 overbought, <30 oversold.
//
// trueLow = min(low, prevClose); BP = close - trueLow
// TR = max(high-low, |prevClose-high|, |prevClose-low|)
// avg_n = (sum BP over n bars) / (sum TR over n bars)
// ULTOSC = 100 * (4avg_short + 2avg_mid + avg_long) / 7
//
// Note: The three periods are sorted internally, so the 4/2/1 weighting always applies to the shortest, middle, and longest period regardless of the order in which you pass them.
func ULTOSC(
	inHigh, inLow, inClose []float64,
	optInTimePeriod1, optInTimePeriod2, optInTimePeriod3 int,
) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inClose))
	)

	if retCode := ultosc(
		startIdx,
		endIdx,
		inHigh,
		inLow,
		inClose,
		int32(optInTimePeriod1),
		int32(optInTimePeriod2),
		int32(optInTimePeriod3),
		&outBegIdx,
		&outNBElement,
		outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("ULTOSC", "result", retCode)
		return nil
	}

	return outReal
}

// WILLR - Williams' %R momentum oscillator over a rolling period, bounded in [-100, 0].
// Measures where the current close sits relative to the high-low range of the last N bars.
// Near 0 = close at period high (overbought); near -100 = close at period low (oversold).
//
// %R = -100 * (highestHigh - close) / (highestHigh - lowestLow) over the trailing optInTimePeriod bars; if highestHigh == lowestLow, output 0.
func WILLR(inHigh, inLow, inClose []float64, optInTimePeriod int) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inClose))
	)

	if retCode := willr(
		startIdx,
		endIdx,
		inHigh,
		inLow,
		inClose,
		int32(optInTimePeriod),
		&outBegIdx,
		&outNBElement,
		outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("WILLR", "result", retCode)
		return nil
	}

	return outReal
}
