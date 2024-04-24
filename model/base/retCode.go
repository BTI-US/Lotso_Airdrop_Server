package base

var (
	SUCCESS = NewRetCode(0, "Success")

	UnknownError                  = NewRetCode(10000, "Unknown error")
	WrongParams                   = NewRetCode(10001, "Wrong params")
	DialApiServerFailed           = NewRetCode(10002, "Dial API server failed")
	InvalidAddress                = NewRetCode(10003, "Invalid address")
	EthApiCallFailed              = NewRetCode(10004, "ETH API call failed")
	ConvertFailed                 = NewRetCode(10005, "Convert failed")
	TransactionCountAlreadyExists = NewRetCode(10006, "Transaction count already exists")
	GetTransactionCountFailed     = NewRetCode(10007, "Get transaction count failed")
	SaveTransactionCountFailed    = NewRetCode(10008, "Save transaction count failed")
	DistributeAirdropsFailed      = NewRetCode(10009, "Distribute airdrops failed")
	ClaimAirdropsFailed           = NewRetCode(10010, "Claim airdrops failed")
	RecipientsCountFailed         = NewRetCode(100011, "Get recipients count failed")
)

type RetCode struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewRetCode(code int, message string) *RetCode {
	return &RetCode{Code: code, Message: message}
}
