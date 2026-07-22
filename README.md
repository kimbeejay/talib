# talib

Go wrapper for selected [TA-Lib](https://ta-lib.org/) functions using [`purego`](https://github.com/ebitengine/purego) to dynamically load the native TA-Lib shared library at runtime.

This project currently exposes a small subset of TA-Lib functions:

- `HT_DCPERIOD`
- `HT_DCPHASE`
- `HT_PHASOR`
- `HT_SINE`
- `HT_TRENDMODE`
- `Add`
- `Div`
- `Max`
- `MaxIndex`
- `Min`
- `MinIndex`
- `MinMax`
- `MinMaxIndex`
- `Mult`
- `Sub`
- `Sum`

## Requirements

- Go `1.26`
- A native TA-Lib shared library installed on the host system

This package does **not** bundle TA-Lib itself. You must install the native library separately.

## How library loading works

Call `talib.Load()` once before using any indicator function. The loader looks for the platform-specific library name:

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
- Native TA-Lib functions are exposed with Go-style exported names such as `Add`, `Div`, `Max`, and `Sum`.

## Project status

This is currently a focused wrapper around a limited subset of TA-Lib. Additional indicators can be added by registering more native functions in `Load()` and exposing Go wrappers for them.
