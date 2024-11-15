package service

import (
	"Lotso_Airdrop_Server/database"
	"Lotso_Airdrop_Server/model/base"
)

func GetAddressesShouldAirdrop() (response *base.Response) {
	airdropItems, err := database.GetAddressesShouldAirdrop()
	if err != nil {
		return base.NewErrorResponse(err, base.GetAirdropItemFailed)
	}

	response = base.NewDataResponse(airdropItems)
	return
}
