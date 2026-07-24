package talib

type MAType int32

const (
	SMA   MAType = 0 // Simple Moving Average
	EMA   MAType = 1 // Exponential Moving Average
	WMA   MAType = 2 // Weighted Moving Average
	DEMA  MAType = 3 // Double Exponential Moving Average
	TEMA  MAType = 4 // Triple Exponential Moving Average
	TRIMA MAType = 5 // Triangular Moving Average
	KAMA  MAType = 6 // Kaufman Adaptive Moving Average
	MAMA  MAType = 7 // MESA Adaptive Moving Average
	T3    MAType = 8 // Triple Exponential Moving Average (T3)
)
