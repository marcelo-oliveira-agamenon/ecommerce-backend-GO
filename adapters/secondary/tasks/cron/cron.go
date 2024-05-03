package cronjob

import (
	"errors"

	"github.com/ecommerce/core/domain/order"
	"github.com/ecommerce/core/domain/product"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

var (
	ErrorToken = errors.New("invalid generated token")
)

func NewCronTasks(ps *gorm.DB) {
	cr := cron.New(cron.WithChain((cron.SkipIfStillRunning(cron.DefaultLogger))))
	cr.AddFunc("0 0 1 * * *", func() {
		VerifyRateOrderAndModifyProduct(ps)
	})
	go cr.Start()
}

func VerifyRateOrderAndModifyProduct(ps *gorm.DB) {
	var or []order.Order
	res := ps.Where("status = 'ENTREGUE'").Find(&or)
	if res.Error != nil {
		return
	}
	if len(or) == 0 {
		return
	}

	lis := make(map[string]int)
	rt := 0
	for _, num := range or {
		for i := 0; i < len(num.ProductID); i++ {
			lis[num.ProductID[i]] = lis[num.ProductID[i]] + 1
		}

		rt = rt + num.Rate
	}

	for k := range lis {
		res := ps.Model(&product.Product{}).Where("id = ?", k).Update("rate", rt/len(or))
		if res.Error != nil {
			return
		}
	}
}
