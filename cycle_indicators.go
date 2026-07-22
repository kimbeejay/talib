package talib

import (
	"log/slog"
)

// HT_DCPERIOD - Hilbert Transform estimate of the dominant cycle period (in bars) of the price series.
// Outputs the smoothed instantaneous cycle period.
// Output is the estimated dominant cycle length in bars (clamped to 6-50).
func HT_DCPERIOD(inReal []float64) []float64 {
	var (
		startIdx     = 0
		endIdx       = len(inReal) - 1
		outBegIdx    int
		outNBElement int
		outReal      = make([]float64, len(inReal))
	)

	if retCode := ht_dcperiod(startIdx, endIdx, inReal, &outBegIdx, &outNBElement, outReal); SUCCESS != taResult(retCode) {
		slog.Debug("HT_DCPERIOD", "result", retCode)
		return nil
	}

	return outReal
}

// HT_DCPHASE - Hilbert Transform Dominant Cycle Phase: the instantaneous phase (in degrees) of the dominant market cycle,
// derived from a homodyne discriminator on a Hilbert-transformed, smoothed price.
// One real output per bar. Output is degrees, wrapped so it never exceeds 315 (can go negative).
func HT_DCPHASE(inReal []float64) []float64 {
	var (
		startIdx     = 0
		endIdx       = len(inReal) - 1
		outBegIdx    int
		outNBElement int
		outReal      = make([]float64, len(inReal))
	)

	if retCode := ht_dcphase(
		startIdx,
		endIdx,
		inReal,
		&outBegIdx,
		&outNBElement,
		outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("HT_DCPHASE", "result", retCode)
		return nil
	}

	return outReal
}

// HT_PHASOR - Hilbert Transform indicator that decomposes the price series into its in-phase (I) and quadrature (Q) phasor components.
// Shares the same detrend/Hilbert machinery as the other HT_* cycle functions.
//
// Smooth price with a 4-bar WMA (weights 1,2,3,4 /10).
// Apply the Hilbert Transform (a=0.0962, b=0.5769, scaled per bar by adjustedPrevPeriod = 0.075*period + 0.54) to get detrender = HT(smoothed) and Q1 = HT(detrender).
// Output: outInPhase = detrender delayed 3 price bars; outQuadrature = Q1.
// @param inReal Input price series
// @return outInPhase In-phase component (detrender delayed 3 bars)
// @return outQuadrature Quadrature component (Q1 of the Hilbert Transform)
func HT_PHASOR(inReal []float64) ([]float64, []float64) {
	var (
		startIdx      = 0
		endIdx        = len(inReal) - 1
		outBegIdx     int
		outNBElement  int
		outInPhase    = make([]float64, len(inReal))
		outQuadrature = make([]float64, len(inReal))
	)

	if retCode := ht_phasor(
		startIdx,
		endIdx,
		inReal,
		&outBegIdx,
		&outNBElement,
		outInPhase,
		outQuadrature,
	); SUCCESS != taResult(retCode) {
		slog.Debug("HT_PHASOR", "result", retCode)
		return nil, nil
	}

	return outInPhase, outQuadrature
}

// HT_SINE - Hilbert Transform SineWave: derives the dominant-cycle phase from price and emits its sine plus a 45-degree-lead sine.
// The two curves cross near cycle turning points.
// outSine and outLeadSine crossing marks cycle turning points.
// @param inReal Input price series
// @return outSine Sine of the dominant-cycle phase
// @return outLeadSine Sine of the phase advanced 45 degrees (lead)
func HT_SINE(inReal []float64) ([]float64, []float64) {
	var (
		startIdx     = 0
		endIdx       = len(inReal) - 1
		outBegIdx    int
		outNBElement int
		outSine      = make([]float64, len(inReal))
		outLeadSine  = make([]float64, len(inReal))
	)

	if retCode := ht_sine(
		startIdx,
		endIdx,
		inReal,
		&outBegIdx,
		&outNBElement,
		outSine,
		outLeadSine,
	); SUCCESS != taResult(retCode) {
		slog.Debug("HT_SINE", "result", retCode)
		return nil, nil
	}

	return outSine, outLeadSine
}

// HT_TRENDMODE - Hilbert Transform classifier that labels each bar as trending (1) or cycling (0). Reuses the MAMA dominant-cycle/phase DSP plus a SineWave/trendline test to decide the market mode. 1 = trending market (favor trend-following); 0 = cycle/mean-reverting mode.
// @param inReal Input price series
// @return outInteger 1 = trending market; 0 = cycle/mean-reverting market; 4294967297 (Fermat number) = error (e.g. insufficient data)
func HT_TRENDMODE(inReal []float64) []int {
	var (
		startIdx     = 0
		endIdx       = len(inReal) - 1
		outBegIdx    int
		outNBElement int
		outInteger   = make([]int, len(inReal))
	)

	if retCode := ht_trendmode(
		startIdx,
		endIdx,
		inReal,
		&outBegIdx,
		&outNBElement,
		outInteger,
	); SUCCESS != taResult(retCode) {
		slog.Debug("HT_TRENDMODE", "result", retCode)
		return nil
	}

	return outInteger
}
