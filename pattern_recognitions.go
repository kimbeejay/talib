package talib

import "log/slog"

// Cdl2Crows - Three-candle bearish reversal pattern: a long white candle, then a black candle gapping up,
// then a black candle that opens inside the second body and closes down inside the first white body.
// A hit (-100) signals a bearish reversal; significant in an uptrend, which this function does not verify.
//
// Note: Does not verify the prior uptrend the pattern classically assumes for significance.
func Cdl2Crows(inOpen, inHigh, inLow, inClose []float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdl2crows(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("Cdl2Crows", "result", retCode)
		return nil
	}

	return outReal
}

// Cdl3BlackCrows - A four-bar pattern: a white candle followed by three consecutive black (down) candles with successively lower closes,
// each opening inside the prior black's real body. It is a bearish reversal signal.
// A hit (-100) signals a bearish reversal.
//
// Note: Does not verify the prior mature uptrend the pattern classically assumes for significance.
func Cdl3BlackCrows(inOpen, inHigh, inLow, inClose []float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outInteger   = make([]int32, len(inClose))
	)

	if retCode := cdl3blackcrows(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose,
		&outBegIdx, &outNBElement, outInteger,
	); SUCCESS != taResult(retCode) {
		slog.Debug("Cdl3BlackCrows", "result", retCode)
		return nil
	}

	return outInteger
}

// Cdl3Inside - A three-candle reversal pattern: a long real body, then a short real body totally engulfed by it (a harami),
// then a third candle of opposite color to the first that closes past the first candle's open.
// Signals a bullish (three inside up) or bearish (three inside down) reversal.
// A hit is a reversal signal: +100 = three inside up (bullish, significant in a downtrend);
// -100 = three inside down (bearish, significant in an uptrend).
//
// Note: Does not verify the prior trend the pattern classically assumes
// (three inside up is meaningful in a downtrend, three inside down in an uptrend).
func Cdl3Inside(inOpen, inHigh, inLow, inClose []float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdl3inside(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("Cdl3Inside", "result", retCode)
		return nil
	}

	return outReal
}

// Cdl3LineStrike - A four-candle pattern: three same-color candles with consecutively higher (or lower) closes,
// each opening within or near the prior real body,
// then a fourth opposite-color candle that opens beyond the third close and closes past the first candle's open.
// TA-Lib emits a signed continuation-style signal keyed to the color of the first three candles.
// +100 = three-white (bullish) strike, -100 = three-black (bearish) strike;
// traditionally read as significant only inside a trend matching the first three candles.
//
// Note: Does not verify the surrounding trend the pattern classically assumes for significance.
func Cdl3LineStrike(inOpen, inHigh, inLow, inClose []float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdl3linestrike(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("Cdl3LineStrike", "result", retCode)
		return nil
	}

	return outReal
}

// Cdl3Outside - A three-candle pattern: an engulfing pair (candle 2's body fully engulfs candle 1's body)
// followed by a third candle that confirms in the engulfing direction.
// Signals a bullish reversal (Three Outside Up) or bearish reversal (Three Outside Down).
// +100 = bullish reversal (Three Outside Up); -100 = bearish reversal (Three Outside Down).
//
// Note: Does not verify the prior trend the pattern classically assumes
// (three outside up is meaningful in a downtrend, three outside down in an uptrend).
func Cdl3Outside(inOpen, inHigh, inLow, inClose []float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdl3outside(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("Cdl3Outside", "result", retCode)
		return nil
	}

	return outReal
}

// Cdl3StarsInSouth - A three-candle bullish reversal pattern of three consecutive black candles that progressively shrink and stabilize:
// a long black candle with a long lower shadow, a smaller black candle probing lower,
// then a small black marubozu contained within the second candle's range.
// A hit (+100) signals a bullish reversal; per the code comment it is meaningful in a downtrend,
// but the function does not verify prior trend.
//
// Note: Does not verify the prior downtrend the pattern classically assumes for significance.
func Cdl3StarsInSouth(inOpen, inHigh, inLow, inClose []float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdl3starsinsouth(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("Cdl3StarsInSouth", "result", retCode)
		return nil
	}

	return outReal
}

// Cdl3WhiteSoldiers - A three-candle pattern of consecutive white candles with progressively higher closes,
// each opening within/near the prior body and each with a very short upper shadow.
// It is a bullish reversal signal. A hit (+100) is bullish, signaling a reversal
// (most meaningful in a downtrend, which the code does not verify).
//
// Note: Does not verify the prior downtrend the pattern classically assumes for significance.
func Cdl3WhiteSoldiers(inOpen, inHigh, inLow, inClose []float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdl3whitesoldiers(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("Cdl3WhiteSoldiers", "result", retCode)
		return nil
	}

	return outReal
}

// CdlAbandonedBaby - A three-candle reversal pattern: a long body, then a gapped-away doji,
// then a body of opposite color that gaps back the other way and closes deep into the first body.
// Bullish (bottom) or bearish (top) reversal signal.
// Nonzero hit signals a reversal: +100 abandoned baby bottom (bullish), -100 abandoned baby top (bearish).
//
// Note: Does not verify the prior trend the pattern classically assumes for significance.
//
// @param penetration - Fraction of the 1st candle's real body the 3rd close must penetrate; default is 0.3 (>=0).
func CdlAbandonedBaby(inOpen, inHigh, inLow, inClose []float64, penetration float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdlabandonedbaby(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose, penetration,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("CdlAbandonedBaby", "result", retCode)
		return nil
	}

	return outReal
}

// CdlAdvanceBlock - Three-candle bearish reversal pattern: three white candles with consecutively higher closes
// whose advance weakens (progressively smaller bodies and/or lengthening upper shadows).
// Signals that an uptrend's advance is being blocked.
// A hit (-100) is bearish: the advance is stalling/blocked; meaningful mainly within an existing uptrend.
//
// Note: Does not verify the prior uptrend the pattern classically assumes for significance.
func CdlAdvanceBlock(inOpen, inHigh, inLow, inClose []float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdladvanceblock(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("CdlAdvanceBlock", "result", retCode)
		return nil
	}

	return outReal
}

// CdlBeltHold - Single-candle pattern with a long real body that opens at (or near) its extreme.
// A bullish belt-hold is a long white candle with no/very short lower shadow;
// a bearish belt-hold is a long black candle with no/very short upper shadow.
// A white hit is bullish (opens at the low, closes strong);
// a black hit is bearish (opens at the high, closes weak).
//
// One candle.
// Requires real body > BodyLong average (long body),
// then either: white body (close>=open) AND lower shadow < ShadowVeryShort average -> bullish;
// OR black body (close<open) AND upper shadow < ShadowVeryShort average -> bearish.
// No prior-trend or gap conditions are checked.
//
// Note: Does not verify the prior trend that the pattern's bullish/bearish reading classically assumes.
func CdlBeltHold(inOpen, inHigh, inLow, inClose []float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdlbelthold(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("CdlBeltHold", "result", retCode)
		return nil
	}

	return outReal
}

// CdlBreakaway - A five-candle reversal pattern: a long first candle, a same-colored second candle that gaps away from it by its real body,
// two more candles extending the move, and an opposite-colored fifth candle that closes back inside the gap.
// Emits a bullish signal (bottom reversal) or bearish signal (top reversal).
// A hit signals a reversal: +100 bullish (bottom), -100 bearish (top).
//
// Note: Does not verify the prior trend the pattern classically assumes (a breakaway matters most against a preceding move).
func CdlBreakaway(inOpen, inHigh, inLow, inClose []float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdlbreakaway(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("CdlBreakaway", "result", retCode)
		return nil
	}

	return outReal
}

// CdlClosingMarubozu - Single-candle pattern: a long real body whose closing end has no or very short shadow,
// so the close sits at the candle's extreme. Non-directional strong bar that emits +100 for a white body and -100 for a black body.
// White (+100) is bullish, black (-100) is bearish; a strong directional bar, not a defined reversal/continuation signal.
//
// One candle.
// Requires: (1) long real body: real body > the BodyLong average; AND (2) very short shadow at the closing end:
// if white (close>=open) upper shadow < the ShadowVeryShort average [close at/near high];
// if black (close<open) lower shadow < the ShadowVeryShort average [close at/near low].
func CdlClosingMarubozu(inOpen, inHigh, inLow, inClose []float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdlclosingmarubozu(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("CdlClosingMarubozu", "result", retCode)
		return nil
	}

	return outReal
}

// CdlConcealBabysWall - A four-candle pattern: two black marubozus,
// then a black candle that gaps down but pokes its upper shadow into the prior body,
// then a larger black candle fully engulfing the third. Bullish reversal signal. A hit signals a bullish reversal.
//
// Note: Does not verify the preceding downtrend the pattern classically assumes.
func CdlConcealBabysWall(inOpen, inHigh, inLow, inClose []float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdlconcealbabyswall(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("CdlConcealBabysWall", "result", retCode)
		return nil
	}

	return outReal
}

// CdlCounterAttack - A two-candle pattern of two long, opposite-colored real bodies whose closing prices are (nearly) equal.
// Emits a bullish signal when the second candle is white and a bearish signal when it is black
// (a reversal signal, though its trend context is not checked).
//
// A hit signals a reversal:
// +100 (white 2nd candle) bullish,
// -100 (black 2nd candle) bearish; significance depends on a prior trend the code does not check.
//
// Note: Does not verify the prior trend the reversal signal classically assumes.
func CdlCounterAttack(inOpen, inHigh, inLow, inClose []float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdlcounterattack(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("CdlCounterAttack", "result", retCode)
		return nil
	}

	return outReal
}

// CdlDarkCloudCover - A two-candle bearish reversal pattern:
// a long white candle followed by a black candle that opens above the prior high and closes deep into the prior white body past a penetration threshold.
// Signals a potential top. A hit (-100) is a bearish reversal signal, most meaningful after an uptrend.
//
// Note: Does not verify the preceding uptrend the bearish reversal classically assumes.
//
// @param penetration - Fraction of candle 1's real body that candle 2's close must penetrate below close[i-1];
// larger values require deeper penetration; default is 0.5 (>=0).
func CdlDarkCloudCover(inOpen, inHigh, inLow, inClose []float64, penetration float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdldarkcloudcover(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose, penetration,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("CdlDarkCloudCover", "result", retCode)
		return nil
	}

	return outReal
}

// CdlDoji - Single-candle Doji recognizer: fires when the real body (|close-open|) is at or below the BodyDoji threshold.
// Returns 100 on a match, 0 otherwise. Market indecision; neither bullish nor bearish on its own.
//
// match if |close-open| <= BodyDoji average (small real body).
func CdlDoji(inOpen, inHigh, inLow, inClose []float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdldoji(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("CdlDoji", "result", retCode)
		return nil
	}

	return outReal
}

// CdlDojiStar - A two-candle reversal pattern: a long real body followed by a doji whose real body gaps away from it
// (up after a white body, down after a black body). Signals a potential reversal of the prevailing trend.
// A hit flags a likely trend reversal; true direction depends on the prevailing trend (bullish in a downtrend, bearish in an uptrend),
// which the code does not itself verify.
//
// Two candles.
// Candle 1: long real body (realbody > BodyLong average).
// Candle 2: doji (realbody <= BodyDoji average).
// Gap: either candle 1 white (color1) AND candle 2 real body gaps up above it (the real bodies gap up),
// or candle 1 black (color-1) AND candle 2 real body gaps down below it (the real bodies gap down).
//
// Note: Does not verify the prior trend the reversal signal classically assumes.
func CdlDojiStar(inOpen, inHigh, inLow, inClose []float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdldojistar(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("CdlDojiStar", "result", retCode)
		return nil
	}

	return outReal
}

// CdlDragonflyDoji - Single-candle pattern: a doji (open and close nearly equal) sitting at the top of the range,
// with no meaningful upper shadow and a long lower shadow.
// A reversal signal, but its bullish/bearish meaning depends on the prior trend (the code does not judge direction).
// A hit marks a dragonfly doji; treated as a potential reversal, but direction (bullish/bearish) must be read from the trend it appears in.
//
// Single candle.
// realbody <= BodyDoji average (doji body) AND upper shadow < ShadowVeryShort average (no/very short upper shadow) AND lower shadow > ShadowVeryShort average (lower shadow present, not very short).
// No color, gap, or trend test.
//
// Note: Does not verify the prior trend that determines the pattern's bullish/bearish meaning.
func CdlDragonflyDoji(inOpen, inHigh, inLow, inClose []float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdldragonflydoji(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("CdlDragonflyDoji", "result", retCode)
		return nil
	}

	return outReal
}

// CdlEngulfing - A two-candle reversal pattern where the second candle's real body engulfs the first candle's opposite-colored real body.
// Bullish (white engulfs black) or bearish (black engulfs white) reversal signal.
// Bullish reversal at +100/+80, bearish at -100/-80;
// ideally after a downtrend (bullish) or uptrend (bearish), which the code does not verify.
//
// Note: Does not verify the prior trend (down for bullish, up for bearish) the reversal classically assumes.
func CdlEngulfing(inOpen, inHigh, inLow, inClose []float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdlengulfing(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("CdlEngulfing", "result", retCode)
		return nil
	}

	return outReal
}

// CdlEveningDojiStar - A three-candle bearish reversal pattern: a long white candle, a doji that gaps up (the star),
// then a black candle closing well down into the first candle's body.
// A stricter Evening Star whose middle candle must be a doji.
// Hit (-100) signals a bearish top reversal.
//
// Note: Does not verify the preceding uptrend the bearish reversal classically assumes.
//
// @param penetration - Fraction of the 1st real body the 3rd candle's close must penetrate; larger demands a deeper close into the first body; default is 0.3 (>=0).
func CdlEveningDojiStar(inOpen, inHigh, inLow, inClose []float64, penetration float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdleveningdojistar(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose, penetration,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("CdlEveningDojiStar", "result", retCode)
		return nil
	}

	return outReal
}

// CdlEveningStar - A three-candle bearish reversal pattern: a long white candle, a short-bodied star gapping up,
// then a black candle closing well down into the first candle's body.
// A hit signals a bearish reversal (most significant in an uptrend).
//
// Note: Does not verify the preceding uptrend the bearish reversal classically assumes.
// The third candle only needs a body longer than short, not the full long body some definitions require.
//
// @param penetration - Fraction of the 1st candle's real body the 3rd close must penetrate below the 1st close;
// larger requires deeper penetration; default is 0.3 (>=0).
func CdlEveningStar(inOpen, inHigh, inLow, inClose []float64, penetration float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdleveningstar(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose, penetration,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("CdlEveningStar", "result", retCode)
		return nil
	}

	return outReal
}

// CdlGapSideBySideWhite - A three-candle pattern:
// a first candle followed by two white candles of similar body size that both gap the same direction (up or down) from the first candle's real body and open at about the same level.
// It is a continuation signal whose sign reports the gap direction; the code does not verify a prior trend.
// A hit signals continuation in the gap's direction: +100 with an upside gap is bullish, -100 with a downside gap is bearish.
//
// Note: Does not verify the prior trend the continuation signal classically assumes.
func CdlGapSideBySideWhite(inOpen, inHigh, inLow, inClose []float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdlgapsidesidewhite(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("CdlGapSideBySideWhite", "result", retCode)
		return nil
	}

	return outReal
}

// CdlGravestoneDoji - Single-candle doji whose open and close sit at the low of the day, leaving a long upper shadow and no lower shadow.
// A doji variant whose bullish/bearish meaning depends on the surrounding trend, which the code does not judge.
// A hit marks a gravestone doji;
// its bullish vs bearish reversal meaning must be read against the prevailing trend, which this function does not check.
//
// One candle.
// Detected when all hold:
// (1) doji body: realbody |close-open| <= BodyDoji average;
// (2) very short/absent lower shadow: lowerShadow < ShadowVeryShort average;
// (3) non-short upper shadow: upperShadow > ShadowVeryShort average (open/close at the low with an upper shadow).
//
// Note: Does not verify the prior trend that determines the pattern's bullish/bearish meaning.
func CdlGravestoneDoji(inOpen, inHigh, inLow, inClose []float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdlgravestonedoji(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("CdlGravestoneDoji", "result", retCode)
		return nil
	}

	return outReal
}

// CdlHammer - Single-candle pattern:
// a small real body at the top of the range with a long lower shadow and little or no upper shadow, sitting at or near the prior candle's low.
// Bullish reversal signal. A hit (+100) flags a potential bullish reversal.
//
// Note: Does not verify the preceding downtrend that the pattern classically assumes; confirm the trend context yourself.
func CdlHammer(inOpen, inHigh, inLow, inClose []float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdlhammer(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("CdlHammer", "result", retCode)
		return nil
	}

	return outReal
}

// CdlHangingMan - Single candle with a small real body,
// a long lower shadow, and little/no upper shadow, sitting at or near the highs of the prior candle.
// Bearish reversal signal. A hit is a bearish reversal signal (meaningful at the top of an uptrend).
//
// Note: Does not verify the preceding uptrend that the pattern classically assumes; confirm the trend context yourself.
func CdlHangingMan(inOpen, inHigh, inLow, inClose []float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdlhangingman(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("CdlHangingMan", "result", retCode)
		return nil
	}

	return outReal
}

// CdlHarami - Two-candle pattern: a long real body followed by a short real body contained within the first candle's real body.
// A reversal signal whose direction is the opposite of the first candle's color.
// A hit signals a potential reversal opposite the 1st candle: positive (black 1st) is bullish, negative (white 1st) is bearish.
//
// Note: Does not verify the prior trend (downtrend for bullish, uptrend for bearish) that the reversal signal assumes.
//
// @returns
//
// +100/+80 when the long 1st candle is black (bullish),
// -100/-80 when it is white (bearish), 0 otherwise;
//
// 80 when the two real bodies share an end,
// 100 when the 1st body strictly overhangs both ends of the 2nd
func CdlHarami(inOpen, inHigh, inLow, inClose []float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdlharami(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("CdlHarami", "result", retCode)
		return nil
	}

	return outReal
}

// CdlHaramiCross - A two-candle reversal pattern:
// a long real body followed by a doji whose real body is contained within the first candle's real body (the doji variant of the Harami).
// Bullish after a black first candle, bearish after a white first candle.
// A hit signals a potential reversal: +100/+80 bullish (black first candle), -100/-80 bearish (white first candle).
//
// Note: Does not verify the prior trend (downtrend for bullish, uptrend for bearish) that the reversal signal assumes.
//
// @returns
//
// +100/+80 when the first candle is black (bullish), -100/-80 when the first candle is white (bearish), 0 otherwise.
// Magnitude 100 for strict containment inside the first body, 80 when one real-body end matches
func CdlHaramiCross(inOpen, inHigh, inLow, inClose []float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdlharamicross(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("CdlHaramiCross", "result", retCode)
		return nil
	}

	return outReal
}

// CdlHighWave - Single-candle pattern: a short real body with both a very long upper and a very long lower shadow.
// Signals market indecision; the output sign reports only candle color, not a bullish/bearish direction.
// A hit marks indecision (long-legged candle); not directional - sign encodes only the candle's color.
//
// One candle at index i. Hit when all hold:
// (1) short real body: real body < the BodyShort average;
// (2) very long upper shadow: upper shadow > the ShadowVeryLong average;
// (3) very long lower shadow: lower shadow > the ShadowVeryLong average.
//
// No color, gap, or trend condition.
//
// @returns
//
// On a hit, +100 when the candle is white (close >= open) or -100 when black (close < open); 0 otherwise.
// Sign denotes color, NOT bull/bear
func CdlHighWave(inOpen, inHigh, inLow, inClose []float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdlhighwave(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("CdlHighWave", "result", retCode)
		return nil
	}

	return outReal
}

// CdlHikkake - A 3-bar pattern:
// an inside bar followed by a false breakout, optionally later confirmed by a follow-through bar.
// Signals a bullish or bearish reversal/continuation depending on the breakout direction.
// A false-breakout setup: positive = bullish, negative = bearish; magnitude 200 flags the confirming bar.
//
// @returns
//
// +100/-100 at the hikkake (breakout) bar for bull/bear;
// +200/-200 at a later confirmation bar; 0 otherwise
func CdlHikkake(inOpen, inHigh, inLow, inClose []float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdlhikkake(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("CdlHikkake", "result", retCode)
		return nil
	}

	return outReal
}

// CdlHikkakeMod - A four-candle pattern:
// two successively narrower inside bars, then a breakout bar, with the second candle closing near one extreme of its range.
// Bullish or bearish reversal signal.
// Bullish (+) or bearish (-) reversal; per the code's note it is significant in a downtrend (bull) or uptrend (bear), context the code does not verify.
//
// Note: Does not verify the prior trend (downtrend for bullish, uptrend for bearish) that this reversal pattern assumes.
//
// @returns
//
// +100 bullish hikkake bar, -100 bearish;
// +200 confirmed bullish, -200 confirmed bearish (confirmation adds another +/-100); 0 otherwise
func CdlHikkakeMod(inOpen, inHigh, inLow, inClose []float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdlhikkakemod(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("CdlHikkakeMod", "result", retCode)
		return nil
	}

	return outReal
}

// CdlHomingPigeon - Two-candle pattern:
// a long black candle followed by a small black candle whose real body sits inside the prior body.
// Bullish reversal signal. A hit signals a bullish reversal (meaningful in a downtrend, which the code does not verify).
//
// Two candles at i-1 and i.
// Both black: close[i-1] < open[i-1] and close[i] < open[i].
// First body long: realbody[i-1] > BodyLong average.
// Second body short: realbody[i] <= BodyShort average.
// Second body contained by first: open[i] < open[i-1] and close[i] > close[i-1].
//
// Note: Does not verify the preceding downtrend that the bullish reversal classically assumes.
//
// @returns
//
// +100 when the pattern is detected, 0 otherwise.
// Never emits -100 (always bullish)
func CdlHomingPigeon(inOpen, inHigh, inLow, inClose []float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdlhomingpigeon(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("CdlHomingPigeon", "result", retCode)
		return nil
	}

	return outReal
}

// CdlIdentical3Crows - A three-candle bearish reversal pattern:
// three consecutive declining black candles, each with a very short (or no) lower shadow,
// where each candle after the first opens at or very near the prior candle's close.
// A hit signals a bearish reversal (pattern is always bearish).
//
// Note: Does not verify the preceding uptrend that the bearish reversal classically assumes.
// Does not require the three bodies to be equal in size; 'identical' refers only to each candle opening at or near the previous candle's close.
//
// @returns
//
// -100 when the pattern is detected (always bearish), 0 otherwise. Never emits +100
func CdlIdentical3Crows(inOpen, inHigh, inLow, inClose []float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdlidentical3crows(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("CdlIdentical3Crows", "result", retCode)
		return nil
	}

	return outReal
}

// CdlInNeck - A two-candle in-neck pattern:
// a long black candle followed by a white candle that opens below the prior low and closes just barely into the prior body (near the prior close).
// It is a bearish continuation signal. A hit signals bearish continuation (the down move is expected to resume).
//
// Two candles.
// First: black (close1 < open1) with a long real body (realbody > candleaverage(BodyLong)).
// Second: white (close2 >= open2), opens below the first candle's low (open2 < low1), and closes slightly into the first body: close2 >= close1 AND close2 <= close1 + candleaverage(Equal).
// No prior-trend check is performed.
//
// Note: Does not verify the preceding downtrend that this bearish continuation pattern assumes.
//
// @returns
//
// -100 when the in-neck pattern is detected, 0 otherwise.
// This pattern only ever emits the negative (bearish) signal; it never emits +100
func CdlInNeck(inOpen, inHigh, inLow, inClose []float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdlinneck(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("CdlInNeck", "result", retCode)
		return nil
	}

	return outReal
}

// CdlInvertedHammer - Single-candle pattern:
// a small real body with a long upper shadow and little-to-no lower shadow that gaps down from the prior candle.
// Bullish reversal signal. A hit (+100) flags a potential bullish reversal.
//
// Note: Does not verify the preceding downtrend that the pattern classically assumes;
// it only checks the gap down from the immediately preceding candle.
//
// @returns
//
// +100 when the inverted hammer is detected, 0 otherwise.
// Never emits -100; the pattern is always bullish
func CdlInvertedHammer(inOpen, inHigh, inLow, inClose []float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdlinvertedhammer(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("CdlInvertedHammer", "result", retCode)
		return nil
	}

	return outReal
}

// CdlKicking - Two-candle pattern of two opposite-color marubozu (long bodies with very short shadows) separated by a price gap.
// A reversal signal whose direction is set by the second candle's color.
// Hit signals a reversal in the direction of the second candle: +100 bullish, -100 bearish.
//
// @returns
//
// +100 or -100 on a hit, 0 otherwise.
// +100 when the second candle is white (bullish), -100 when it is black (bearish), 0 otherwise
func CdlKicking(inOpen, inHigh, inLow, inClose []float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdlkicking(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("CdlKicking", "result", retCode)
		return nil
	}

	return outReal
}

// CdlKickingByLength - A two-candle pattern of two opposite-color marubozu (long body, very short shadows on both ends) separated by a gap.
// A strong directional/reversal signal whose bull/bear bias is set by the longer of the two marubozu.
// A hit signals a strong directional move; +100 bullish / -100 bearish per the color of the longer marubozu.
//
// @returns
//
// +100 or -100 on a hit, 0 otherwise.
// Sign = candlecolor of the candle with the larger realbody (i if realbody(i) > realbody(i-1), else i-1; tie goes to i-1): +100 if that marubozu is white, -100 if black
func CdlKickingByLength(inOpen, inHigh, inLow, inClose []float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdlkickingbylength(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("CdlKickingByLength", "result", retCode)
		return nil
	}

	return outReal
}

// CdlLadderBottom - Five-candle bullish reversal pattern: three consecutively lower black candles,
// a fourth black candle with a non-very-short upper shadow,
// then a white candle that opens above the prior open and closes above the prior high.
// Signals a potential bottom reversal. A hit (+100) is a bullish reversal signal, most meaningful after a downtrend.
//
// Note: Does not verify the preceding downtrend that this bullish reversal classically assumes.
//
// @returns
//
// +100 on a detected ladder bottom, 0 otherwise.
// Only ever emits +100 (never -100); inherently bullish
func CdlLadderBottom(inOpen, inHigh, inLow, inClose []float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdlladderbottom(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("CdlLadderBottom", "result", retCode)
		return nil
	}

	return outReal
}

// CdlLongLeggedDoji - Single-candle doji (open ~ close) with at least one long shadow.
// Signals market indecision, not a directional bias. Marks indecision/uncertainty;
// not inherently bullish or bearish despite the positive sign.
//
// One candle.
// Hit when: real body <= BodyDoji average (doji body) AND (lower shadow > ShadowLong average OR upper shadow > ShadowLong average),
// i.e. at least one long shadow.
//
// Note: Only one long shadow (upper or lower) is required, whereas the classic pattern shows both long upper and lower shadows.
//
// @returns
//
// +100 when the pattern is present, 0 otherwise.
// Only +100 is emitted; the code never emits -100, and the positive sign does NOT mean bullish
func CdlLongLeggedDoji(inOpen, inHigh, inLow, inClose []float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdllongleggeddoji(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("CdlLongLeggedDoji", "result", retCode)
		return nil
	}

	return outReal
}

// CdlLongLine - A single-candle pattern: a long real body with short upper and short lower shadow.
// The signal direction follows the candle color (bullish if white, bearish if black).
// Signals strong directional conviction on the bar: +100 white/bullish, -100 black/bearish.
// Not intrinsically a reversal or continuation signal.
//
// @returns
//
// +100 on a white (close>=open) long line, -100 on a black long line, 0 when no pattern
func CdlLongLine(inOpen, inHigh, inLow, inClose []float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdllongline(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("CdlLongLine", "result", retCode)
		return nil
	}

	return outReal
}

// CdlMarubozu - Single candle with a long real body and no/very-short upper and lower shadows, so open and close sit at the range extremes.
// Bullish (white) or bearish (black) reversal/strength signal per the body color.
// +100 = white marubozu (strong buying pressure); -100 = black marubozu (strong selling pressure).
//
// One candle at i.
// Match when: realbody(i) > BodyLong average AND upperShadow(i) < ShadowVeryShort average AND lowerShadow(i) < ShadowVeryShort average.
// If matched emit candlecolor(i)*100 (+100 white when close>=open, -100 black when close<open); else 0.
//
// @returns
//
// +100 on a white (bullish) marubozu, -100 on a black (bearish) marubozu, 0 when no pattern.
// Sign follows the candle color
func CdlMarubozu(inOpen, inHigh, inLow, inClose []float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdlmarubozu(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("CdlMarubozu", "result", retCode)
		return nil
	}

	return outReal
}

// CdlMatchingLow - A two-candle pattern of two consecutive black (bearish) candles with equal closes (within a tolerance).
// Treated as a bullish reversal signal.
// A hit signals a potential bullish reversal (shared support close after two down candles).
//
// Two candles i-1, i.
// Candle i-1: black (close<open).
// Candle i: black (close<open). E
// qual closes: close[i-1]-E <= close[i] <= close[i-1]+E, where E = the Equal average.
// No shadow, body-size, or gap conditions are checked.
//
// Note: The bullish-reversal reading assumes a prior downtrend, which is not verified.
//
// @returns
//
// +100 when the pattern is present, 0 otherwise.
// Only +100 is ever emitted (matching low is always bullish); never -100
func CdlMatchingLow(inOpen, inHigh, inLow, inClose []float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdlmatchinglow(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("CdlMatchingLow", "result", retCode)
		return nil
	}

	return outReal
}

// CdlMatHold - A five-candle bullish continuation pattern: a long white candle, an upside real-body-gapped small black candle,
// two more small falling candles that hold within the first body, and a final white candle closing above the reaction days' highs.
// Signals continuation of the prior uptrend. Hit = bullish continuation of the existing uptrend.
//
// Note: The colors of the third and fourth (reaction) candles are not checked, although they are classically black.
// The continuation reading assumes a prior uptrend, which is not verified.
//
// @param penetration - Max fraction of the 1st white body the reaction days (3rd, 4th) may penetrate; default is 0.5 (>=0).
// @returns
//
// +100 when the bullish Mat Hold is detected, 0 otherwise.
// Never emits -100
func CdlMatHold(inOpen, inHigh, inLow, inClose []float64, penetration float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdlmathold(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose,
		penetration,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("CdlMatHold", "result", retCode)
		return nil
	}

	return outReal
}

// CdlMorningDojiStar - A three-candle bullish reversal pattern:
// a long black candle, then a doji that gaps down, then a white candle closing well up into the first candle's body.
// It is the doji-star variant of the morning star.
// A hit (+100) signals a bullish reversal; most meaningful after a downtrend, which this function does not verify.
//
// Note: The gap-down is measured between the candles' real bodies, not between their high/low ranges.
// A prior downtrend is not verified.
//
// @param penetration - Fraction of the 1st candle's real body the 3rd close must exceed above close[i-2];
// larger values demand deeper penetration into the black body; default is 0.3 (>=0).
//
// @returns
//
// +100 when the pattern is detected, 0 otherwise.
// Always bullish; never emits -100
func CdlMorningDojiStar(inOpen, inHigh, inLow, inClose []float64, penetration float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdlmorningdojistar(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose, penetration,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("CdlMorningDojiStar", "result", retCode)
		return nil
	}

	return outReal
}

// CdlMorningStar - A three-candle bottom-reversal pattern:
// a long black candle, a small-bodied star gapping down, then a white candle closing well up into the first candle's body.
// Bullish reversal signal.
// A hit signals a bullish reversal (most meaningful after a downtrend, which the code does not check).
//
// Note: The gap-down is measured between the candles' real bodies, not between their high/low ranges.
// A prior downtrend is not verified.
//
// @param penetration - Fraction of the 1st candle's body the 3rd close must exceed above the 1st close; larger = deeper penetration required; default is 0.3 (>=0).
//
// @returns
//
// +100 when the morning star is detected, 0 otherwise. Never negative (pattern is exclusively bullish)
func CdlMorningStar(inOpen, inHigh, inLow, inClose []float64, penetration float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdlmorningstar(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose, penetration,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("CdlMorningStar", "result", retCode)
		return nil
	}

	return outReal
}

// CdlOnNeck - A two-candle on-neck pattern:
// a long black candle followed by a white candle that opens below the prior candle's low and closes right at that low.
// Bearish continuation signal.
// A hit is bearish (bearish continuation); the code does not verify the assumed prior downtrend.
//
// Two candles.
// 1st: black (close<open) with long real body (realbody > BodyLong average).
// 2nd: white (close>=open); open < prior low; close within the Equal band of the prior low,
// i.e. (prior_low - EqualAvg) <= close2 <= (prior_low + EqualAvg).
//
// Note: The bearish-continuation reading assumes a prior downtrend, which is not verified.
//
// @returns
//
// -100 on a match, 0 otherwise.
// Only -100 is ever emitted (never +100); on-neck is always bearish
func CdlOnNeck(inOpen, inHigh, inLow, inClose []float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdlonneck(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("CdlOnNeck", "result", retCode)
		return nil
	}

	return outReal
}

// CdlPiercing - Two-candle pattern:
// a long black candle followed by a long white candle that opens below the prior low and closes back above the midpoint of the prior black body.
// Bullish reversal signal.
// A hit (+100) is a bullish reversal signal.
//
// Note: A prior downtrend is not verified.
//
// @returns
//
// +100 when the piercing pattern is detected; 0 otherwise.
// Always bullish, never emits -100
func CdlPiercing(inOpen, inHigh, inLow, inClose []float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdlpiercing(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("CdlPiercing", "result", retCode)
		return nil
	}

	return outReal
}

// CdlRickshawMan - Single-candle doji with two long shadows whose body sits near the midpoint of the high-low range.
// It is a neutral indecision signal, not a directional (bullish/bearish) reversal.
// A hit marks market indecision/uncertainty; neutral, neither bullish nor bearish.
//
// @returns
//
// +100 when the pattern is present, 0 otherwise.
// Never -100; the code notes the positive value does NOT imply bullish, it signals uncertainty
func CdlRickshawMan(inOpen, inHigh, inLow, inClose []float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdlrickshawman(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("CdlRickshawMan", "result", retCode)
		return nil
	}

	return outReal
}

// CdlRiseFall3Methods - A five-candle continuation pattern:
// a long candle, three small counter-color candles that stay partly within the first candle's high-low range, then a long same-color candle that resumes the trend.
// Bullish (rising) or bearish (falling) continuation signal.
// A hit signals trend continuation: +100 = bullish (rising three methods), -100 = bearish (falling three methods).
//
// Note: Only the three-small-candle variant is detected; the classic pattern allowing two or more small candles is not supported.
// The middle candles need only partially overlap the first candle's range, not be fully contained within it.
// The prior trend the continuation reading assumes is not verified.
//
// @returns
//
// +100 when candle 1 is white (rising/bullish continuation), -100 when candle 1 is black (falling/bearish continuation), 0 otherwise.
// Sign = 100 * color of candle 1
func CdlRiseFall3Methods(inOpen, inHigh, inLow, inClose []float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdlrisefall3methods(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("CdlRiseFall3Methods", "result", retCode)
		return nil
	}

	return outReal
}

// CdlSeparatingLines - A two-candle continuation pattern:
// the second candle opposes the first in color, opens at the same price as the first, and is a long-bodied belt hold.
// Bullish (white second candle) or bearish (black second candle) continuation signal.
// Trend continuation: +100 = bullish (white belt hold), -100 = bearish (black belt hold).
//
// Two consecutive candles i-1, i:
// (1) opposite colors: color(i-1) == -color(i);
// (2) same open: open[i-1]-Equal_avg <= open[i] <= open[i-1]+Equal_avg;
// (3) long body: realbody(i) > BodyLong_avg;
// (4) belt hold: if i is white, lowershadow(i) < ShadowVeryShort_avg; if i is black, uppershadow(i) < ShadowVeryShort_avg.
//
// Note: A prior trend is not verified, nor that the pattern aligns with it.
//
// @returns
//
// +100 for a bullish (white second candle) hit, -100 for a bearish (black second candle) hit, 0 otherwise
func CdlSeparatingLines(inOpen, inHigh, inLow, inClose []float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdlseparatinglines(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("CdlSeparatingLines", "result", retCode)
		return nil
	}

	return outReal
}

// CdlShootingStar - Single-candle pattern:
// a small real body with a long upper shadow and little-to-no lower shadow that gaps up from the prior candle's real body.
// Bearish reversal signal. A hit (-100) flags a bearish reversal at the top of an uptrend.
//
// Note: A preceding uptrend is not verified.
//
// @returns
//
// -100 when the shooting star is detected, 0 otherwise.
// Only ever emits negative (bearish); never +100
func CdlShootingStar(inOpen, inHigh, inLow, inClose []float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdlshootingstar(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("CdlShootingStar", "result", retCode)
		return nil
	}

	return outReal
}

// CdlShortLine - Single-candle pattern: a short real body with short upper and lower shadows (a small-range candle).
// Not a directional signal — the output sign encodes candle color, not bullish/bearish sentiment.
// A hit only flags a small-range candle; the +/- sign is the candle's color (white/black), not a reversal or continuation call.
//
// One candle at i, all three:
//
// short real body: real body < the BodyShort average
// short upper shadow: upper shadow < the ShadowShort average
// short lower shadow: lower shadow < the ShadowShort average
// If matched: output = candle color * 100 (+100 white, -100 black); else 0.
//
// @returns
//
// +100 for a matching white candle (close>=open), -100 for a matching black candle (close<open), 0 when no pattern.
// Sign is candle color, NOT bullish/bearish
func CdlShortLine(inOpen, inHigh, inLow, inClose []float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdlshortline(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("CdlShortLine", "result", retCode)
		return nil
	}

	return outReal
}

// CdlSpinningTop - Single-candle pattern:
// a small real body with both an upper and a lower shadow longer than the body.
// Signals indecision; the code does not classify it as bullish or bearish.
// A hit marks indecision (small body, both shadows long); the sign only reports candle color, not direction.
//
// One candle where: upper shadow > real body AND lower shadow > real body AND real body < the BodyShort average.
// The BodyShort average is the factor-scaled mean body over the prior avgPeriod candles.
//
// @returns
//
// +100 when the candle is white (close>=open), -100 when black (close<open), 0 when no pattern.
// Sign is candle color, NOT bullish/bearish
func CdlSpinningTop(inOpen, inHigh, inLow, inClose []float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdlspinningtop(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("CdlSpinningTop", "result", retCode)
		return nil
	}

	return outReal
}

// CdlStalledPattern - A three-candle pattern of three white candles with consecutively higher closes where the third loses momentum
// (a small body riding on the shoulder of the second's long body).
// It is a bearish reversal signal of a stalling advance.
// A hit (-100) is bearish: the uptrend is stalling and may reverse.
//
// Note: The pattern classically appears in an uptrend, but this function does not verify a prior uptrend; the caller must confirm it.
//
// @returns
//
// -100 when the pattern is detected (always bearish), 0 otherwise.
// Never emits +100
func CdlStalledPattern(inOpen, inHigh, inLow, inClose []float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdlstalledpattern(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("CdlStalledPattern", "result", retCode)
		return nil
	}

	return outReal
}

// CdlStickSandwich - A three-candle bullish reversal pattern:
// two black candles (1st and 3rd) sandwiching a white candle, where the 3rd black candle closes at the same level as the 1st (the "bread").
// A hit signals a bullish reversal (code comment notes it is significant in a downtrend, which the function does not verify).
//
// @returns
//
// +100 when the pattern is present, 0 otherwise.
// Never -100 — Stick Sandwich is always bullish
func CdlStickSandwich(inOpen, inHigh, inLow, inClose []float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdlsticksandwich(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("CdlStickSandwich", "result", retCode)
		return nil
	}

	return outReal
}

// CdlTakuri - Single-candle pattern:
// a doji whose open and close sit at the high (no/very short upper shadow) with a very long lower shadow,
// i.e. a dragonfly doji with an exceptionally long lower shadow.
// Emitted as a positive signal, but its directional meaning depends on the prevailing trend, which the code does not check.
// A hit marks a takuri (dragonfly-doji) line; a potential reversal only when read against the trend
// (typically a bottom/bullish reversal after a downtrend), which the code itself does not verify.
//
// @returns
//
// +100 when the takuri pattern is detected, 0 otherwise.
// Never negative; the positive sign is a convention and does not by itself imply bullishness
func CdlTakuri(inOpen, inHigh, inLow, inClose []float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdltakuri(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("CdlTakuri", "result", retCode)
		return nil
	}

	return outReal
}

// CdlTasukiGap - A three-candle pattern:
// a real-body-gapping candle followed by an opposite-color candle that opens inside its body and closes back into the gap without filling it.
// An upside gap is a bullish continuation signal; a downside gap is a bearish continuation signal.
// Hit signals trend continuation: +100 bullish (in an uptrend), -100 bearish (in a downtrend).
//
// Note: This continuation pattern does not verify the prior trend it classically assumes; the caller must confirm the trend.
//
// @returns
//
// +100 on a bullish (upside-gap) tasuki gap, -100 on a bearish (downside-gap) tasuki gap, 0 otherwise.
// Sign equals the color of the gap candle i-1 (candlecolor(i-1)*100)
func CdlTasukiGap(inOpen, inHigh, inLow, inClose []float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdltasukigap(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("CdlTasukiGap", "result", retCode)
		return nil
	}

	return outReal
}

// CdlThrusting - A two-candle pattern:
// a long black candle followed by a white candle that opens below the prior low and closes back into the prior body but below its midpoint.
// It is a bearish continuation signal.
// A hit is bearish: the failed white push back into the black body signals continuation of the down move.
//
// Note: The pattern is classically meaningful only in a downtrend, but this function does not verify any prior trend.
// Although the pattern can be read as bullish in an uptrend or when it recurs,
// this function ignores trend and always reports it as bearish.
//
// @returns
//
// -100 when the pattern is detected, 0 otherwise.
// Always bearish; never emits +100
func CdlThrusting(inOpen, inHigh, inLow, inClose []float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdlthrusting(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("CdlThrusting", "result", retCode)
		return nil
	}

	return outReal
}

// CdlTristar - A three-candle pattern of three consecutive doji where the middle doji is a star (its body gaps away from the first).
// Bullish or bearish reversal signal. +100 = bullish reversal (middle doji gapped down), -100 = bearish reversal (middle doji gapped up).
//
// Note: This reversal pattern does not verify the prior trend it classically assumes.
//
// @returns
//
// +100 (bullish, star gapped down), -100 (bearish, star gapped up), or 0 when no pattern.
// Both signs are emitted
func CdlTristar(inOpen, inHigh, inLow, inClose []float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdltristar(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("CdlTristar", "result", retCode)
		return nil
	}

	return outReal
}

// CdlUnique3River - A three-candle bullish reversal pattern:
// a long black candle, then a black harami candle that makes a lower low, then a small white candle.
// Signals a potential bullish reversal, ideally in a downtrend (trend not checked by the code).
// A hit (+100) marks a bullish reversal; significant in a downtrend, which the function does not verify.
//
// @returns
//
// +100 when the pattern is present, 0 otherwise.
// Bullish-only: never emits -100
func CdlUnique3River(inOpen, inHigh, inLow, inClose []float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdlunique3river(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("CdlUnique3River", "result", retCode)
		return nil
	}

	return outReal
}

// CdlUpsideGap2Crows - A three-candle bearish reversal pattern:
// a long white candle, then a small black candle gapping up (a gap between the real bodies),
// then a black candle that engulfs the second candle's real body but still closes above the first candle's close.
// Signals a bearish reversal. A hit (-100) is a bearish reversal signal, most meaningful in an uptrend.
//
// Note: The pattern classically assumes a prior uptrend, but this function does not verify any trend.
//
// @returns
//
// -100 on a pattern bar, 0 otherwise.
// Bearish-only: this pattern never emits +100
func CdlUpsideGap2Crows(inOpen, inHigh, inLow, inClose []float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdlupsidegap2crows(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("CdlUpsideGap2Crows", "result", retCode)
		return nil
	}

	return outReal
}

// CdlXSideGap3Methods - A three-candle continuation pattern: two same-color candles separated by a real-body gap,
// followed by an opposite-color candle that fills into the gap.
// Bullish (upside) when the first two candles are white, bearish (downside) when they are black.
// A hit signals trend continuation: +100 bullish (uptrend resumes), -100 bearish (downtrend resumes).
//
// Note: This continuation pattern does not verify the prior trend it classically assumes; the caller must confirm the trend.
//
// @returns
//
// +100 when the two same-color candles are white (bullish/upside continuation),
// -100 when black (bearish/downside continuation), 0 otherwise.
// Equals candlecolor(1st candle) * 100
func CdlXSideGap3Methods(inOpen, inHigh, inLow, inClose []float64) []int32 {
	var (
		startIdx     int32
		endIdx       = int32(len(inClose) - 1)
		outBegIdx    int32
		outNBElement int32
		outReal      = make([]int32, len(inClose))
	)

	if retCode := cdlxsidegap3methods(
		startIdx, endIdx,
		inOpen, inHigh, inLow, inClose,
		&outBegIdx, &outNBElement, outReal,
	); SUCCESS != taResult(retCode) {
		slog.Debug("CdlXSideGap3Methods", "result", retCode)
		return nil
	}

	return outReal
}
