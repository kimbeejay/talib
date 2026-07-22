package talib

import "log/slog"

// ADD - Vector arithmetic addition. Outputs the element-wise sum of two input series.
//
// outReal[i] = inReal0[i] + inReal1[i]
func ADD(inReal0, inReal1 []float64) []float64 {
	var (
		startIdx     = 0
		endIdx       = len(inReal0) - 1
		outBegIdx    int
		outNBElement int
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
		slog.Debug("ADD", "result", retCode)
		return nil
	}

	return outReal[:outNBElement]
}
