package talib

type taResult int32

const (
	SUCCESS                   taResult = 0 // No error
	LIB_NOT_INITIALIZE                 = 1 // TA_Initialize was not successfully called
	BAD_PARAM                          = 2 // A parameter is out of range
	ALLOC_ERR                          = 3 // Possibly out-of-memory
	GROUP_NOT_FOUND                    = 4
	FUNC_NOT_FOUND                     = 5
	INVALID_HANDLE                     = 6
	INVALID_PARAM_HOLDER               = 7
	INVALID_PARAM_HOLDER_TYPE          = 8
	INVALID_PARAM_FUNCTION             = 9
	INPUT_NOT_ALL_INITIALIZE           = 10
	OUTPUT_NOT_ALL_INITIALIZE          = 11
	OUT_OF_RANGE_START_INDEX           = 12
	OUT_OF_RANGE_END_INDEX             = 13
	INVALID_LIST_TYPE                  = 14
	BAD_OBJECT                         = 15
	NOT_SUPPORTED                      = 16
	INTERNAL_ERROR                     = 5000
)
