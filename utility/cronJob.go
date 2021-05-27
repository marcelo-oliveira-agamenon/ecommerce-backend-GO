package utility

import (
	"github.com/ecommerce/db"
	"github.com/ecommerce/models"
	"github.com/robfig/cron"
)

//Centralize all jobs
func CallerAllJobs() {
	cronJob := cron.New()
	cronJob.AddFunc("0 0 2 * * *", verifyRateOrderAndModifyProduct)
	cronJob.Start()
}

//Check if a order receive a rate, and mirror this change on order
func verifyRateOrderAndModifyProduct() {
	var orders []models.Order

	result := db.DBConn.Where("status = 'ENTREGUE'").Find(&orders)
	if result.Error != nil {
		return
	}

	dict := make(map[string]int)
	rate := 0
	for _, num := range orders {
		for i := 0; i < len(num.ProductID); i++ {
			dict[num.ProductID[i]] = dict[num.ProductID[i]]+1
		}
		
		rate = rate + num.Rate
	}

	for k := range dict {
		res := db.DBConn.Model(&models.Product{}).Where("id = ?", k).Update("rate", rate / len(orders))
		if res.Error != nil {
			return
		}
	}
}