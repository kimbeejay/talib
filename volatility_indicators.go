package talib

// ATR - Wilder-smoothed average of the True Range over a period, measuring price volatility regardless of direction.
// Higher ATR means greater volatility; no directional bias.
//
// TR_t = max(high-low, |prevClose-high|, |prevClose-low|)
// ATR seed = simple average of first period TR values
// ATR_t = (ATR_{t-1} * (period-1) + TR_t) / period
func ATR(inHigh, inLow, inClose []float64, optInTimePeriod int) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inClose))
	)

	if retCode := atr(
		startIdx,
		endIdx,
		inHigh,
		inLow,
		inClose,
		int32(optInTimePeriod),
		&outBegIdx,
		&outNBElement,
		outReal,
	); retCode != 0 {
		return nil
	}

	return outReal
}

// NATR - Average True Range expressed as a percentage of the current close, making volatility comparable across price levels and securities.
// Same computation as ATR, then normalized by close. Higher values mean greater relative volatility; unit is percent of price.
//
// NATR = (ATR / Close) * 100
// ATR: first value = SMA of TRANGE over period; then Wilder smoothing ATR_t = (ATR_{t-1}*(period-1) + TR_t) / period
func NATR(inHigh, inLow, inClose []float64, optInTimePeriod int) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inClose))
	)

	if retCode := natr(
		startIdx,
		endIdx,
		inHigh,
		inLow,
		inClose,
		int32(optInTimePeriod),
		&outBegIdx,
		&outNBElement,
		outReal,
	); retCode != 0 {
		return nil
	}

	return outReal
}

// TRANGE - True Range: the greatest of today's high-low span and the two gaps between yesterday's close and today's high/low.
// Base volatility measure used to build ATR/NATR. Larger values mean wider or gappier bars (higher volatility).
//
// TR = max( high - low, |prevClose - high|, |prevClose - low| )
//
// Note: The first bar produces no value because it has no prior close;
// unlike some definitions, it does not fall back to the high-low range for that bar.
func TRANGE(inHigh, inLow, inClose []float64) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inClose))
	)

	if retCode := trange(
		startIdx,
		endIdx,
		inHigh,
		inLow,
		inClose,
		&outBegIdx,
		&outNBElement,
		outReal,
	); retCode != 0 {
		return nil
	}

	return outReal
}
