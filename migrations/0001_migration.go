package main

import (
	"fmt"
	"gin_demo/models"
	"gorm.io/gorm"
)

func main()  {
	models.ConnectDB()
	err := models.DB.Transaction(func(tx *gorm.DB) error {
		err := migrate(tx)
		return err
	})
	if err != nil {
		fmt.Print(err.Error())
	} else {
		fmt.Printf("Successed migrating")
	}
}

func migrate(tx *gorm.DB) error {
	err := tx.Exec("UPDATE book SET count = count + 1").Error
	return err // return nil will commit transaction, otherwise rollback the transaction
}