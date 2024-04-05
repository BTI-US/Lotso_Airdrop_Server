package base

var (
	SUCCESS = NewRetCode(0, "Success")

	UnknownError              = NewRetCode(10000, "Unknown error")
	WrongParams               = NewRetCode(10001, "Wrong params")
	DialApiServerFailed       = NewRetCode(10002, "Dial API server failed")
	InvalidAddress            = NewRetCode(10003, "Invalid address")
	EthApiCallFailed          = NewRetCode(10004, "ETH API call failed")
	ConvertFailed             = NewRetCode(10005, "Convert failed")
	GetTransactionCountFailed = NewRetCode(10006, "Get transaction count failed")
	DistributeAirdropsFailed  = NewRetCode(10007, "Distribute airdrops failed")
)

type RetCode struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewRetCode(code int, message string) *RetCode {
	return &RetCode{Code: code, Message: message}
}
