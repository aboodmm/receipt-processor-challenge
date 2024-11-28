package cache

import (
	"github.com/aboodmm/receipt-processor/models"
)

var cacheO map[string]models.Receipt

func GetCache() map[string]models.Receipt {

	if cacheO != nil {
		return cacheO
	}
	cacheO = make(map[string]models.Receipt)
	return cacheO
}
