package talib

// Beta - Beta: the slope of a least-squares linear regression of one series' percentage returns (y, from inReal1) against another's (x, from inReal0) over a rolling window.
// Measures how much a security moves relative to a market index.
// Beta = 1 moves with the index; < 1 less volatile, > 1 more volatile.
//
// @param inReal0: Series whose returns are the regression x (market/index)
// @param inReal1: Series whose returns are the regression y (security)
// @param optInTimePeriod: Rolling window length (number of returns) for the regression sums; default is 5 (1-100000)
//
// @return: regression slope of inReal1-returns on inReal0-returns
func Beta(inReal0, inReal1 []float64, optInTimePeriod int) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal0) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inReal0))
	)

	if retCode := beta(
		startIdx,
		endIdx,
		inReal0,
		inReal1,
		int32(optInTimePeriod),
		&outBegIdx,
		&outNBElement,
		outReal,
	); retCode != 0 {
		return nil
	}

	return outReal
}

// Correl - Pearson's correlation coefficient (r) between two input series over a rolling window of optInTimePeriod bars.
// Measures how linearly the two series move together. r near +1: strong positive co-movement; near -1: strong inverse; near 0: no linear relationship.
//
// r = (sumXY - sumX*sumY/n) / sqrt((sumX2 - sumX^2/n) * (sumY2 - sumY^2/n)), n = optInTimePeriod, sums over the window
//
// Note: When the correlation is undefined for a window (for example a constant series), the output is 0 rather than an error or NaN.
//
// @param inReal0: First data series (X)
// @param inReal1: Second data series (Y)
// @param optInTimePeriod: Rolling window length (number of bars) for the correlation sums; default is 30 (2-100000)
//
// @return: Correlation coefficient r in [-1, 1]
func Correl(inReal0, inReal1 []float64, optInTimePeriod int) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal0) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inReal0))
	)

	if retCode := correl(
		startIdx,
		endIdx,
		inReal0,
		inReal1,
		int32(optInTimePeriod),
		&outBegIdx,
		&outNBElement,
		outReal,
	); retCode != 0 {
		return nil
	}

	return outReal
}

// LinearReg - Least-squares straight-line fit over the last optInTimePeriod bars, reported as the fitted line value at the window endpoint (b + m*(period-1)).
//
// @param inReal: Input series
// @param optInTimePeriod: Number of bars (period) for the regression; default is 14 (2-100000)
//
// @return: Fitted line value at the window endpoint
func LinearReg(inReal []float64, optInTimePeriod int) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inReal))
	)

	if retCode := linearreg(
		startIdx,
		endIdx,
		inReal,
		int32(optInTimePeriod),
		&outBegIdx,
		&outNBElement,
		outReal,
	); retCode != 0 {
		return nil
	}

	return outReal
}

// LinearRegAngle - The angle, in degrees, of the least-squares best-fit line over the last N points.
// It is the LINEARREG_SLOPE value passed through atan and converted to degrees.
// Positive angle = rising fit line, negative = falling; magnitude reflects steepness.
//
// m = (N·SumXY − SumX·SumY) / (SumX² − N·SumXSqr), with SumX=N(N−1)/2, SumXSqr=N(N−1)(2N−1)/6; angle = atan(m)·(180/π)
func LinearRegAngle(inReal []float64, optInTimePeriod int) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inReal))
	)

	if retCode := linearreg_angle(
		startIdx,
		endIdx,
		inReal,
		int32(optInTimePeriod),
		&outBegIdx,
		&outNBElement,
		outReal,
	); retCode != 0 {
		return nil
	}

	return outReal
}

// LinearRegIntercept - Returns the y-intercept (b) of the least-squares regression line fitted over the last optInTimePeriod values.
// Part of the linear-regression family (LINEARREG, SLOPE, ANGLE, TSF).
//
// Fit y = b + m·x over the window with x = bars-ago (x=0 is the current bar, x=period-1 the oldest).
// With SumX = period(period-1)/2, SumXSqr = period(period-1)(2·period-1)/6, Divisor = SumX² − period·SumXSqr:
// m = (period·SumXY − SumX·SumY) / Divisor
// b = (SumY − m·SumX) / period ← output
func LinearRegIntercept(inReal []float64, optInTimePeriod int) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inReal))
	)

	if retCode := linearreg_intercept(
		startIdx,
		endIdx,
		inReal,
		int32(optInTimePeriod),
		&outBegIdx,
		&outNBElement,
		outReal,
	); retCode != 0 {
		return nil
	}

	return outReal
}

// LinearRegSlope - Slope 'm' of the least-squares best-fit line (y = b + m*x) over the last optInTimePeriod bars.
// Reports the per-bar rate of change of the fitted trend line.
// Positive slope = rising trend, negative = falling; magnitude is price change per bar.
//
// m = (n·SumXY − SumX·SumY) / Divisor
// SumX = n(n−1)/2, SumXSqr = n(n−1)(2n−1)/6, Divisor = SumX² − n·SumXSqr
// SumXY = Σ i·y[today−i], SumY = Σ y[today−i], i=0..n−1, n=period, y=inReal
func LinearRegSlope(inReal []float64, optInTimePeriod int) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inReal))
	)

	if retCode := linearreg_slope(
		startIdx,
		endIdx,
		inReal,
		int32(optInTimePeriod),
		&outBegIdx,
		&outNBElement,
		outReal,
	); retCode != 0 {
		return nil
	}

	return outReal
}

// StdDev - Rolling standard deviation of a series over a window, scaled by a deviations multiplier.
// Delegates to VAR, then takes the square root.
//
// Note: Uses population variance (divides by the period, not period minus one), so results differ slightly from the sample standard deviation used by some tools.
// @param inReal: Input series
// @param optInTimePeriod: Number of bars (period) for the rolling window; default is 5 (2-100000)
// @param optInNbDev: Multiplier for the standard deviation; default is 1.0 (0.0-500.0)
//
// @return: Rolling standard deviation scaled by optInNbDev
func StdDev(inReal []float64, optInTimePeriod int, optInNbDev float64) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inReal))
	)

	if retCode := stddev(
		startIdx,
		endIdx,
		inReal,
		int32(optInTimePeriod),
		optInNbDev,
		&outBegIdx,
		&outNBElement,
		outReal,
	); retCode != 0 {
		return nil
	}

	return outReal
}

// TSF - Time Series Forecast: fits a least-squares linear regression line over the last N bars and projects it one x-step beyond talib.LinearReg.
// Same regression as talib.LinearReg but evaluated at x=period instead of x=period-1.
//
// Fit y=b+mx over window (x=0..N-1): m = (NSumXY - SumXSumY)/(SumX^2 - NSumXSqr), b = (SumY - mSumX)/N; output = b + mN.
// With SumX=N(N-1)/2, SumXSqr=N(N-1)(2N-1)/6.
func TSF(inReal []float64, optInTimePeriod int) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inReal))
	)

	if retCode := tsf(
		startIdx,
		endIdx,
		inReal,
		int32(optInTimePeriod),
		&outBegIdx,
		&outNBElement,
		outReal,
	); retCode != 0 {
		return nil
	}

	return outReal
}

// Var - Rolling population variance of a real series over a given period.
// Measures dispersion of values around their mean. Higher values indicate greater dispersion; 0 means constant input.
//
// Var = (SumX2 - SumX^2/n) / n, n = optInTimePeriod, sums over the window
//
// Note: Computes population variance (divides by the period), not the sample variance (n-1) used by some definitions.
// The deviation-count parameter is accepted but has no effect on the result.
func Var(inReal []float64, optInTimePeriod int, optInNbDev float64) []float64 {
	var (
		startIdx     int32
		endIdx       = int32(len(inReal) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]float64, len(inReal))
	)

	if retCode := variance(
		startIdx,
		endIdx,
		inReal,
		int32(optInTimePeriod),
		optInNbDev,
		&outBegIdx,
		&outNBElement,
		outReal,
	); retCode != 0 {
		return nil
	}

	return outReal
}
