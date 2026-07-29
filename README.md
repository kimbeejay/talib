# talib

Go wrapper for all [TA-Lib](https://ta-lib.org/) functions using [`purego`](https://github.com/ebitengine/purego) to dynamically load the native TA-Lib shared library at runtime.

This package exposes the complete TA-Lib function set across these categories:

- **Cycle Indicators**: Hilbert Transform (`HT_*`)
- **Math Operators**: Vector operations (Add, Sub, Mult, Div, Max, Min, Sum, etc.)
- **Math Transforms**: Trigonometric, hyperbolic, exponential, and rounding functions
- **Momentum Indicators**: ADX, MACD, RSI, Stoch, Aroon, oscillators, ROC, directional movement
- **Overlap Studies**: Moving averages and bands (EMA, SMA, BBands, KAMA, MAMA, SAR, etc.)
- **Price Transforms**: Per-bar price composites (AvgPrice, TypPrice, WCLPrice, etc.)
- **Statistic Functions**: Regression, correlation, and dispersion (Beta, Correl, LinearReg*, StdDev, Var)
- **Volatility Indicators**: ATR, NATR, TRANGE
- **Volume Indicators**: AD, ADOsc, CMF, OBV, PVI, PVO
- **Pattern Recognition**: 61 candlestick patterns (`Cdl*`)

For a complete function reference with descriptions, see [TA-Lib's official function list](https://ta-lib.org/functions/).

## Requirements

- Go `1.26`
- A native TA-Lib shared library installed on the host system

This package does **not** bundle TA-Lib itself. You must install the native library separately.

## How library loading works

Call `talib.Load()` once before using any indicator function. After loading, you can call `talib.Version()` to read the native TA-Lib version string. The loader looks for the platform-specific library name:

- macOS: `libta-lib.dylib`
- Linux: `libta-lib.so`
- Windows: `libta-lib.dll`

Lookup order:

1. `$TA_LIB_PATH/<library file>`
2. `./<library file>`
3. `/usr/local/lib/<library file>`
4. `/usr/lib/<library file>`

## Installing TA-Lib

### macOS

If you use Homebrew:

```bash
brew install ta-lib
```

If the library is installed outside the default lookup paths, set:

```bash
export TA_LIB_PATH=/path/to/lib
```

### Linux

Install TA-Lib with your package manager if available, or build it from source and place the shared library in a standard library directory such as `/usr/local/lib`.

### Windows

Install the TA-Lib DLL and make sure `libta-lib.dll` is reachable through `TA_LIB_PATH` or from the current working directory.

## Usage

```go
package main

import (
	"fmt"
	"log"
	"math"

	"beejay.kim/lib/talib"
)

func main() {
	if _, err := talib.Load(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("TA-Lib version:", talib.Version())

	var series []float64
	for i := 0; i < 100; i++ {
        series = append(series, math.Sin(float64(i)/7)*4.5+math.Cos(float64(i)/19)*1.75)
    }

	period := talib.HT_DCPERIOD(series)
	fmt.Println(period)

	sum := talib.Add(series, series)
	fmt.Println(sum)

	rollingMax := talib.Max(series, 14)
	fmt.Println(rollingMax)
}
```

## API notes

- Indicator functions return `nil` when the underlying TA-Lib call fails.
- `Version()` returns the TA-Lib version string exposed by `TA_GetVersionString` (call after `Load()`).
- Native TA-Lib functions are exposed with Go-style exported names such as `Add`, `Div`, `Max`, and `Sum`.
- Functions accepting a moving-average type use the `MAType` constant (`MA_SMA`, `MA_EMA`, `MA_WMA`, `MA_DEMA`, `MA_TEMA`, `MA_TRIMA`, `MA_KAMA`, `MA_MAMA`, `MA_T3`).
- Functions returning multiple outputs (e.g. `Aroon`, `MACD`, `Stoch`) return multiple slices; all are `nil` on failure.
- Pattern-recognition functions (`Cdl*`) return `[]int32` signals; typical values are `-100` (bearish), `0` (no pattern), and `100` (bullish).

## Project status

This wrapper provides **complete coverage** of the TA-Lib function library. All ~100 indicators across 10 function categories are exposed with Go-idiomatic APIs. New versions of TA-Lib may introduce additional functions; updating this wrapper requires adding function pointer declarations in `functions.go`, registering them in `loader.go`, and exposing Go wrappers in the appropriate `*_indicators.go` file.
