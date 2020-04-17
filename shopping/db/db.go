package db

import (
	"shopping/models",
	"github.com/mattn/go-sqlite3"
)

func LoadItem(id int) *models.Item {
	return &models.Item{
		Price: 9.001,
	}
}
