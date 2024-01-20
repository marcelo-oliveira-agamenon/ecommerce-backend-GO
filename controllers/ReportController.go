package controller

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"strings"

	"github.com/ecommerce/db"
	u "github.com/ecommerce/models"
	"github.com/gofiber/fiber"
	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Return a csv file with users
func ExportUsersReport(w *fiber.Ctx)  {
	createdAtStart := w.Query("user_created_at_start")
	createdAtEnd := w.Query("user_created_at_end")
	gender := w.Query("gender")
	var users []u.User
	baseQuery := db.DBConn

	if gender != "masc" && gender != "fem" && gender != "other" {
		w.Status(500).JSON("Wrong attribute to gender field")
		return
	}

	if createdAtStart != "" && createdAtEnd != "" {
		baseQuery = baseQuery.Where("created_at BETWEEN ?::date AND ?::date", createdAtStart, createdAtEnd)
	}

	if gender != "" {
		baseQuery = baseQuery.Where("gender = ?", gender)
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
			var roles string
			for _, v := range v.Roles {
				roles = roles + " " + v
			}

			row := []string{v.ID.String(), v.CreatedAt.String(), v.UpdatedAt.String(), v.DeletedAt.Time.String(), v.Name, v.Email, v.Address, v.ImageKey, v.ImageURL, v.Phone, "", "", v.Birthday, v.Gender, roles}
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

// Insert new users on db from csv
func ImportUsersFromCsvFile(w *fiber.Ctx)  {
	file, err := w.FormFile("csv_file")
	if err != nil {
		w.Status(500).JSON("Error reading file")
		return
	}

	fileOpen, errOpenFile := file.Open()
	if errOpenFile != nil {
		w.Status(500).JSON("Error open file")
		return
	}

	var users []u.User
	scanner := bufio.NewScanner(fileOpen)
	for scanner.Scan() {
		var user u.User
		elements := strings.Split(scanner.Text(), ",")

		id, err := uuid.NewV4()
		if err != nil {
			w.Status(500).JSON("Error in uuid generate")
			return
		}

		hashPassword, err := bcrypt.GenerateFromPassword([]byte(elements[5]), bcrypt.DefaultCost)
		if err != nil {
			w.Status(500).JSON("Error in bcrypt")
			return
		}

		user.ID = id
		user.Name = elements[0]
		user.Email = elements[1]
		user.Address = elements[2]
		user.ImageKey = elements[3]
		user.ImageURL = elements[4]
		user.Phone = elements[5]
		user.Password = string(hashPassword)
		user.FacebookID = elements[7]
		user.Birthday = elements[8]
		user.Gender = elements[9]
		user.Roles = strings.Split(elements[10], " ")

		users = append(users, user)
	}	

	result := db.DBConn.Create(&users)
	if result.Error != nil {
		w.Status(500).JSON("Error importing users")
		return
	}

	if err := scanner.Err(); err != nil {
		w.Status(500).JSON("Error scanning lines of file")
        return
    }

	w.Status(200).JSON("Success importing users")
}