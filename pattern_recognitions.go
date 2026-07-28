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
