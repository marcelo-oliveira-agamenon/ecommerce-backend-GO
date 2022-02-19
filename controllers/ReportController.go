package controller

import (
	"bytes"
	"encoding/csv"

	"github.com/ecommerce/db"
	u "github.com/ecommerce/models"
	"github.com/gofiber/fiber"
)

// Return a csv file with users
func ExportUsersReport(w *fiber.Ctx)  {
	createdAtStart := w.Query("user_created_at_start")
	createdAtEnd := w.Query("user_created_at_end")
	var users []u.User
	baseQuery := db.DBConn

	if createdAtStart != "" && createdAtEnd != "" {
		baseQuery = baseQuery.Where("created_at BETWEEN ?::date AND ?::date", createdAtStart, createdAtEnd)
	}

	result := baseQuery.Find(&users)
	if result.Error != nil {
		w.Status(500).JSON("Error getting users with filters")
		return
	}

	if len(users) > 0 {
		var keys []string
		bytesFile := new(bytes.Buffer)
		file := csv.NewWriter(bytesFile)
		
		resultNames, namesErr := db.DBConn.Debug().Migrator().ColumnTypes(&u.User{})
		if namesErr != nil {
			w.Status(500).JSON("Error getting users column names")
			return
		}

		for _, v := range resultNames {
			keys = append(keys, v.Name())
		}

		if keyErr := file.Write(keys); keyErr != nil {
			w.Status(500).JSON("Error on write csv file")
			return
		}

		for _, v := range users {
			row := []string{v.ID.String(), v.CreatedAt.String(), v.UpdatedAt.String(), v.DeletedAt.Time.String(), v.Name, v.Email, v.Address, v.ImageKey, v.ImageURL, v.Phone, "", "", v.Birthday, v.Gender, ""}
			if valueErr := file.Write(row); valueErr != nil {
				w.Status(500).JSON("Error on write csv file")
				return
			}
		}

		file.Flush()
		if errFile := file.Error(); errFile != nil {
			w.Status(500).JSON("Error on write csv file")
			return
		}

		w.Status(200).JSON(bytesFile.Bytes())
		return
	}

	w.Status(404).JSON("No users found with this filter")
}