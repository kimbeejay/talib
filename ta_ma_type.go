package talib

type MAType int32

const (
	MA_SMA   MAType = 0 // Simple Moving Average
	MA_EMA   MAType = 1 // Exponential Moving Average
	MA_WMA   MAType = 2 // Weighted Moving Average
	MA_DEMA  MAType = 3 // Double Exponential Moving Average
	MA_TEMA  MAType = 4 // Triple Exponential Moving Average
	MA_TRIMA MAType = 5 // Triangular Moving Average
	MA_KAMA  MAType = 6 // Kaufman Adaptive Moving Average
	MA_MAMA  MAType = 7 // MESA Adaptive Moving Average
	MA_T3    MAType = 8 // Triple Exponential Moving Average (T3)
)
