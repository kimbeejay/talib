package talib

import "C"
import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/ebitengine/purego"
)

var (
	internalVersion func() string
	once            sync.Once
)

func Version() string {
	return internalVersion()
}

func Load() (uintptr, error) {
	var (
		ptr uintptr
		err error
	)

	once.Do(func() {
		ptr, err = internalLoad()
	})

	return ptr, err
}

func internalLoad() (uintptr, error) {
	var (
		path string
		ptr  uintptr
		err  error
	)

	if path, err = getExpectedLibraryPath(); err != nil {
		return 0, err
	}

	slog.Debug("Loading TA-Lib library", "path", path)
	if ptr, err = purego.Dlopen(path, purego.RTLD_NOW|purego.RTLD_GLOBAL); err != nil {
		return 0, err
	}

	purego.RegisterLibFunc(&internalVersion, ptr, "TA_GetVersionString")

	// Register the functions we need from the TA-Lib library
	purego.RegisterLibFunc(&ht_dcperiod, ptr, "TA_HT_DCPERIOD")
	purego.RegisterLibFunc(&ht_dcphase, ptr, "TA_HT_DCPHASE")
	purego.RegisterLibFunc(&ht_phasor, ptr, "TA_HT_PHASOR")
	purego.RegisterLibFunc(&ht_sine, ptr, "TA_HT_SINE")
	purego.RegisterLibFunc(&ht_trendmode, ptr, "TA_HT_TRENDMODE")

	purego.RegisterLibFunc(&add, ptr, "TA_ADD")
	purego.RegisterLibFunc(&div, ptr, "TA_DIV")
	purego.RegisterLibFunc(&max, ptr, "TA_MAX")
	purego.RegisterLibFunc(&maxIndex, ptr, "TA_MAXINDEX")
	purego.RegisterLibFunc(&min, ptr, "TA_MIN")
	purego.RegisterLibFunc(&minIndex, ptr, "TA_MININDEX")
	purego.RegisterLibFunc(&minMax, ptr, "TA_MINMAX")
	purego.RegisterLibFunc(&minMaxIndex, ptr, "TA_MINMAXINDEX")
	purego.RegisterLibFunc(&mult, ptr, "TA_MULT")
	purego.RegisterLibFunc(&sub, ptr, "TA_SUB")
	purego.RegisterLibFunc(&sum, ptr, "TA_SUM")

	purego.RegisterLibFunc(&cos, ptr, "TA_COS")
	purego.RegisterLibFunc(&acos, ptr, "TA_ACOS")
	purego.RegisterLibFunc(&cosh, ptr, "TA_COSH")
	purego.RegisterLibFunc(&sin, ptr, "TA_SIN")
	purego.RegisterLibFunc(&asin, ptr, "TA_ASIN")
	purego.RegisterLibFunc(&sinh, ptr, "TA_SINH")
	purego.RegisterLibFunc(&tan, ptr, "TA_TAN")
	purego.RegisterLibFunc(&atan, ptr, "TA_ATAN")
	purego.RegisterLibFunc(&tanh, ptr, "TA_TANH")

	purego.RegisterLibFunc(&ceil, ptr, "TA_CEIL")
	purego.RegisterLibFunc(&floor, ptr, "TA_FLOOR")
	purego.RegisterLibFunc(&sqrt, ptr, "TA_SQRT")

	purego.RegisterLibFunc(&exp, ptr, "TA_EXP")
	purego.RegisterLibFunc(&ln, ptr, "TA_LN")
	purego.RegisterLibFunc(&log10, ptr, "TA_LOG10")

	purego.RegisterLibFunc(&adx, ptr, "TA_ADX")
	purego.RegisterLibFunc(&adxr, ptr, "TA_ADXR")
	purego.RegisterLibFunc(&apo, ptr, "TA_APO")
	purego.RegisterLibFunc(&aroon, ptr, "TA_AROON")
	purego.RegisterLibFunc(&aroonosc, ptr, "TA_AROONOSC")
	purego.RegisterLibFunc(&bop, ptr, "TA_BOP")
	purego.RegisterLibFunc(&cci, ptr, "TA_CCI")
	purego.RegisterLibFunc(&cmo, ptr, "TA_CMO")
	purego.RegisterLibFunc(&cmou, ptr, "TA_CMOU")
	purego.RegisterLibFunc(&dx, ptr, "TA_DX")
	purego.RegisterLibFunc(&imi, ptr, "TA_IMI")
	purego.RegisterLibFunc(&macd, ptr, "TA_MACD")
	purego.RegisterLibFunc(&macdext, ptr, "TA_MACDEXT")
	purego.RegisterLibFunc(&macdfix, ptr, "TA_MACDFIX")
	purego.RegisterLibFunc(&mfi, ptr, "TA_MFI")
	purego.RegisterLibFunc(&mom, ptr, "TA_MOM")
	purego.RegisterLibFunc(&minus_di, ptr, "TA_MINUS_DI")
	purego.RegisterLibFunc(&minus_dm, ptr, "TA_MINUS_DM")
	purego.RegisterLibFunc(&plus_di, ptr, "TA_PLUS_DI")
	purego.RegisterLibFunc(&plus_dm, ptr, "TA_PLUS_DM")
	purego.RegisterLibFunc(&ppo, ptr, "TA_PPO")
	purego.RegisterLibFunc(&roc, ptr, "TA_ROC")
	purego.RegisterLibFunc(&rocp, ptr, "TA_ROCP")
	purego.RegisterLibFunc(&rocr, ptr, "TA_ROCR")
	purego.RegisterLibFunc(&rocr100, ptr, "TA_ROCR100")
	purego.RegisterLibFunc(&rsi, ptr, "TA_RSI")
	purego.RegisterLibFunc(&stoch, ptr, "TA_STOCH")
	purego.RegisterLibFunc(&stochf, ptr, "TA_STOCHF")
	purego.RegisterLibFunc(&stochrsi, ptr, "TA_STOCHRSI")
	purego.RegisterLibFunc(&trix, ptr, "TA_TRIX")
	purego.RegisterLibFunc(&ultosc, ptr, "TA_ULTOSC")
	purego.RegisterLibFunc(&willr, ptr, "TA_WILLR")

	purego.RegisterLibFunc(&accbands, ptr, "TA_ACCBANDS")
	purego.RegisterLibFunc(&bbands, ptr, "TA_BBANDS")
	purego.RegisterLibFunc(&dema, ptr, "TA_DEMA")
	purego.RegisterLibFunc(&ema, ptr, "TA_EMA")
	purego.RegisterLibFunc(&ht_trendline, ptr, "TA_HT_TRENDLINE")
	purego.RegisterLibFunc(&kama, ptr, "TA_KAMA")
	purego.RegisterLibFunc(&ma, ptr, "TA_MA")
	purego.RegisterLibFunc(&mama, ptr, "TA_MAMA")
	purego.RegisterLibFunc(&mavp, ptr, "TA_MAVP")
	purego.RegisterLibFunc(&midpoint, ptr, "TA_MIDPOINT")
	purego.RegisterLibFunc(&midprice, ptr, "TA_MIDPRICE")
	purego.RegisterLibFunc(&sar, ptr, "TA_SAR")
	purego.RegisterLibFunc(&sarext, ptr, "TA_SAREXT")
	purego.RegisterLibFunc(&sma, ptr, "TA_SMA")
	purego.RegisterLibFunc(&t3, ptr, "TA_T3")
	purego.RegisterLibFunc(&tema, ptr, "TA_TEMA")
	purego.RegisterLibFunc(&trima, ptr, "TA_TRIMA")
	purego.RegisterLibFunc(&wma, ptr, "TA_WMA")

	purego.RegisterLibFunc(&avgdev, ptr, "TA_AVGDEV")
	purego.RegisterLibFunc(&avgprice, ptr, "TA_AVGPRICE")
	purego.RegisterLibFunc(&medprice, ptr, "TA_MEDPRICE")
	purego.RegisterLibFunc(&typprice, ptr, "TA_TYPPRICE")
	purego.RegisterLibFunc(&wclprice, ptr, "TA_WCLPRICE")

	purego.RegisterLibFunc(&beta, ptr, "TA_BETA")
	purego.RegisterLibFunc(&correl, ptr, "TA_CORREL")
	purego.RegisterLibFunc(&linearreg, ptr, "TA_LINEARREG")
	purego.RegisterLibFunc(&linearreg_angle, ptr, "TA_LINEARREG_ANGLE")
	purego.RegisterLibFunc(&linearreg_intercept, ptr, "TA_LINEARREG_INTERCEPT")
	purego.RegisterLibFunc(&linearreg_slope, ptr, "TA_LINEARREG_SLOPE")
	purego.RegisterLibFunc(&stddev, ptr, "TA_STDDEV")
	purego.RegisterLibFunc(&tsf, ptr, "TA_TSF")
	purego.RegisterLibFunc(&variance, ptr, "TA_VAR")

	purego.RegisterLibFunc(&atr, ptr, "TA_ATR")
	purego.RegisterLibFunc(&natr, ptr, "TA_NATR")
	purego.RegisterLibFunc(&trange, ptr, "TA_TRANGE")

	purego.RegisterLibFunc(&ad, ptr, "TA_AD")
	purego.RegisterLibFunc(&adosc, ptr, "TA_ADOSC")
	purego.RegisterLibFunc(&cmf, ptr, "TA_CMF")
	purego.RegisterLibFunc(&nvi, ptr, "TA_NVI")
	purego.RegisterLibFunc(&obv, ptr, "TA_OBV")
	purego.RegisterLibFunc(&pvi, ptr, "TA_PVI")
	purego.RegisterLibFunc(&pvo, ptr, "TA_PVO")

	purego.RegisterLibFunc(&cdl2crows, ptr, "TA_CDL2CROWS")
	purego.RegisterLibFunc(&cdl3blackcrows, ptr, "TA_CDL3BLACKCROWS")
	purego.RegisterLibFunc(&cdl3inside, ptr, "TA_CDL3INSIDE")
	purego.RegisterLibFunc(&cdl3linestrike, ptr, "TA_CDL3LINESTRIKE")
	purego.RegisterLibFunc(&cdl3outside, ptr, "TA_CDL3OUTSIDE")
	purego.RegisterLibFunc(&cdl3starsinsouth, ptr, "TA_CDL3STARSINSOUTH")
	purego.RegisterLibFunc(&cdl3whitesoldiers, ptr, "TA_CDL3WHITESOLDIERS")
	purego.RegisterLibFunc(&cdlabandonedbaby, ptr, "TA_CDLABANDONEDBABY")
	purego.RegisterLibFunc(&cdladvanceblock, ptr, "TA_CDLADVANCEBLOCK")
	purego.RegisterLibFunc(&cdlbelthold, ptr, "TA_CDLBELTHOLD")
	purego.RegisterLibFunc(&cdlbreakaway, ptr, "TA_CDLBREAKAWAY")
	purego.RegisterLibFunc(&cdlclosingmarubozu, ptr, "TA_CDLCLOSINGMARUBOZU")
	purego.RegisterLibFunc(&cdlconcealbabyswall, ptr, "TA_CDLCONCEALBABYSWALL")
	purego.RegisterLibFunc(&cdlcounterattack, ptr, "TA_CDLCOUNTERATTACK")
	purego.RegisterLibFunc(&cdldarkcloudcover, ptr, "TA_CDLDARKCLOUDCOVER")
	purego.RegisterLibFunc(&cdldoji, ptr, "TA_CDLDOJI")
	purego.RegisterLibFunc(&cdldojistar, ptr, "TA_CDLDOJISTAR")
	purego.RegisterLibFunc(&cdldragonflydoji, ptr, "TA_CDLDRAGONFLYDOJI")
	purego.RegisterLibFunc(&cdlengulfing, ptr, "TA_CDLENGULFING")
	purego.RegisterLibFunc(&cdleveningdojistar, ptr, "TA_CDLEVENINGDOJISTAR")
	purego.RegisterLibFunc(&cdleveningstar, ptr, "TA_CDLEVENINGSTAR")
	purego.RegisterLibFunc(&cdlgapsidesidewhite, ptr, "TA_CDLGAPSIDESIDEWHITE")
	purego.RegisterLibFunc(&cdlgravestonedoji, ptr, "TA_CDLGRAVESTONEDOJI")
	purego.RegisterLibFunc(&cdlhammer, ptr, "TA_CDLHAMMER")
	purego.RegisterLibFunc(&cdlhangingman, ptr, "TA_CDLHANGINGMAN")
	purego.RegisterLibFunc(&cdlharami, ptr, "TA_CDLHARAMI")
	purego.RegisterLibFunc(&cdlharamicross, ptr, "TA_CDLHARAMICROSS")
	purego.RegisterLibFunc(&cdlhighwave, ptr, "TA_CDLHIGHWAVE")
	purego.RegisterLibFunc(&cdlhikkake, ptr, "TA_CDLHIKKAKE")
	purego.RegisterLibFunc(&cdlhikkakemod, ptr, "TA_CDLHIKKAKEMOD")
	purego.RegisterLibFunc(&cdlhomingpigeon, ptr, "TA_CDLHOMINGPIGEON")
	purego.RegisterLibFunc(&cdlidentical3crows, ptr, "TA_CDLIDENTICAL3CROWS")
	purego.RegisterLibFunc(&cdlinneck, ptr, "TA_CDLINNECK")
	purego.RegisterLibFunc(&cdlinvertedhammer, ptr, "TA_CDLINVERTEDHAMMER")
	purego.RegisterLibFunc(&cdlkicking, ptr, "TA_CDLKICKING")
	purego.RegisterLibFunc(&cdlkickingbylength, ptr, "TA_CDLKICKINGBYLENGTH")
	purego.RegisterLibFunc(&cdlladderbottom, ptr, "TA_CDLLADDERBOTTOM")
	purego.RegisterLibFunc(&cdllongleggeddoji, ptr, "TA_CDLLONGLEGGEDDOJI")
	purego.RegisterLibFunc(&cdllongline, ptr, "TA_CDLLONGLINE")
	purego.RegisterLibFunc(&cdlmarubozu, ptr, "TA_CDLMARUBOZU")
	purego.RegisterLibFunc(&cdlmatchinglow, ptr, "TA_CDLMATCHINGLOW")
	purego.RegisterLibFunc(&cdlmathold, ptr, "TA_CDLMATHOLD")
	purego.RegisterLibFunc(&cdlmorningdojistar, ptr, "TA_CDLMORNINGDOJISTAR")
	purego.RegisterLibFunc(&cdlmorningstar, ptr, "TA_CDLMORNINGSTAR")
	purego.RegisterLibFunc(&cdlonneck, ptr, "TA_CDLONNECK")
	purego.RegisterLibFunc(&cdlpiercing, ptr, "TA_CDLPIERCING")
	purego.RegisterLibFunc(&cdlrickshawman, ptr, "TA_CDLRICKSHAWMAN")
	purego.RegisterLibFunc(&cdlrisefall3methods, ptr, "TA_CDLRISEFALL3METHODS")
	purego.RegisterLibFunc(&cdlseparatinglines, ptr, "TA_CDLSEPARATINGLINES")
	purego.RegisterLibFunc(&cdlshootingstar, ptr, "TA_CDLSHOOTINGSTAR")
	purego.RegisterLibFunc(&cdlshortline, ptr, "TA_CDLSHORTLINE")
	purego.RegisterLibFunc(&cdlspinningtop, ptr, "TA_CDLSPINNINGTOP")
	purego.RegisterLibFunc(&cdlstalledpattern, ptr, "TA_CDLSTALLEDPATTERN")
	purego.RegisterLibFunc(&cdlsticksandwich, ptr, "TA_CDLSTICKSANDWICH")
	purego.RegisterLibFunc(&cdltakuri, ptr, "TA_CDLTAKURI")
	purego.RegisterLibFunc(&cdltasukigap, ptr, "TA_CDLTASUKIGAP")
	purego.RegisterLibFunc(&cdlthrusting, ptr, "TA_CDLTHRUSTING")
	purego.RegisterLibFunc(&cdltristar, ptr, "TA_CDLTRISTAR")
	purego.RegisterLibFunc(&cdlunique3river, ptr, "TA_CDLUNIQUE3RIVER")
	purego.RegisterLibFunc(&cdlupsidegap2crows, ptr, "TA_CDLUPSIDEGAP2CROWS")
	purego.RegisterLibFunc(&cdlxsidegap3methods, ptr, "TA_CDLXSIDEGAP3METHODS")

	return ptr, nil
}

func getExpectedLibraryName() (string, error) {
	switch runtime.GOOS {
	case "windows":
		return "libta-lib.dll", nil
	case "darwin":
		return "libta-lib.dylib", nil
	case "linux":
		return "libta-lib.so", nil
	default:
		return "", fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}
}

func getExpectedLibraryPath() (string, error) {
	libName, err := getExpectedLibraryName()
	if err != nil {
		return "", err
	}

	isExist := func(path string) bool {
		_, err := os.Stat(path)
		return err == nil
	}

	if envpath, ok := os.LookupEnv("TA_LIB_PATH"); ok {
		p := filepath.Join(envpath, libName)
		if isExist(p) {
			return p, nil
		}

		slog.Debug("TA_LIB_PATH environment variable is set but library not found", "path", p)
	}

	for _, dir := range []string{
		".",
		"/usr/local/lib",
		"/usr/lib",
	} {
		p := filepath.Join(dir, libName)
		if isExist(p) {
			return p, nil
		}
	}

	return "", fmt.Errorf("library not found: %s", libName)
}
