package talib

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"

	"github.com/ebitengine/purego"
)

func Load() (uintptr, error) {
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
