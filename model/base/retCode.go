package base

var (
	SUCCESS = NewRetCode(0, "Success")

	UnknownError              = NewRetCode(10000, "Unknown error")
	WrongParams               = NewRetCode(10001, "Wrong params")
	DialApiServerFailed       = NewRetCode(10002, "Dial API server failed")
	InvalidAddress            = NewRetCode(10003, "Invalid address")
	EthApiCallFailed          = NewRetCode(10004, "ETH API call failed")
	ConvertFailed             = NewRetCode(10005, "Convert failed")
	AirdropItemAlreadyExists  = NewRetCode(10006, "You have applied the airdrop") // Airdrop item already exists
	AirdropItemNotExists      = NewRetCode(10007, "Airdrop item not exists")
	GetAirdropItemFailed      = NewRetCode(10008, "Get airdrop item failed")
	GetAddressFailed          = NewRetCode(10008, "Get address item failed")
	SaveAirdropItemFailed     = NewRetCode(10009, "Save airdrop item failed")
	DistributeAirdropsFailed  = NewRetCode(10010, "Distribute airdrops failed")
	ClaimAirdropsFailed       = NewRetCode(10011, "Claim airdrops failed")
	RecipientsCountFailed     = NewRetCode(10012, "Get recipients count failed")
	GetTransactionTopicFailed = NewRetCode(10013, "Get transaction topic failed")
	RewardLimitReached        = NewRetCode(10014, "Reward limit reached")
)

type RetCode struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewRetCode(code int, message string) *RetCode {
	return &RetCode{Code: code, Message: message}
}
