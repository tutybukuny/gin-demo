package models

import (
	"fmt"
	"github.com/fatih/structs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"os/exec"
	"path/filepath"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=dev password=secret dbname=gorm port=5432"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	DB = database
}

func Init() {
	ConnectDB()
	runMigrations()
	autoMigrate()
}

func runMigrations() {
	migration := &Migration{}
	migrate(migration, migration.TableName())

	var files []string

	root := "migrations"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if err = DB.Where("name=?", file).First(migration).Error; err != nil {
			fmt.Printf("Run migration %s", file)
			cmd := exec.Command("go", "run", file)
			out, err := cmd.Output()
			fmt.Println(string(out))
			if err != nil {
				panic(fmt.Sprintf("Error when run %s: %s", file, err))
			}
			migration = &Migration{Name: file}
			DB.Create(migration)
		}
	}
}

func autoMigrate() {
	book := &Book{}
	migrate(book, book.TableName())
}

func migrate(tbl interface{}, tableName string) {
	err := DB.AutoMigrate(tbl)
	if err != nil {
		panic(fmt.Sprintf("Failed to migrate db for table %s", tableName))
	}
	for _, field := range structs.Names(tbl) {
		if field == "ID" {
			continue
		}
		err = DB.Migrator().AlterColumn(tbl, field)
		if err != nil {
			panic(fmt.Sprintf("Failed to migrate column %s of table %s", field, tableName))
		}
	}
}
