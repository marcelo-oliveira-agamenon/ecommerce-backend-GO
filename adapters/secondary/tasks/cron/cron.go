package cronjob

import (
	"errors"
	"log"
	"time"

	"github.com/ecommerce/core/domain/order"
	"github.com/ecommerce/core/domain/product"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

var (
	ErrorToken = errors.New("invalid generated token")
)

func NewCronTasks(ps *gorm.DB) {
	br, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		log.Println("Error in timezone setting, cronjob")
		return
	}
	cr := cron.New(cron.WithLocation(br))
	cr.AddFunc("0 0 1 * * *", func() {
		VerifyRateOrderAndModifyProduct(ps)
	})
	go cr.Start()
}

// TODO: improve this with a field to check order is reviewed
func VerifyRateOrderAndModifyProduct(ps *gorm.DB) {
	var or []order.Order
	res := ps.Preload("OrderDetails").Where("status = 'ENTREGUE' AND is_order_rated = false").Find(&or)
	if res.Error != nil {
		log.Println(res.Error)
		return
	}
	if len(or) == 0 {
		log.Println("No order to check rate in cronjob")
		return
	}

	for i := range or {
		odrs := or[i].OrderDetails
		for j := range odrs {
			prd := odrs[j]
			res1 := ps.Model(&product.Product{}).Where("id = ?", prd.ProductID).Update("rate", or[i].Rate/len(odrs))
			if res1.Error != nil {
				log.Println(res1.Error)
				return
			}
		}
	}
}
