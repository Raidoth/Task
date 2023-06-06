package test

import (
	"test/service/config"
	"test/service/database"
	"test/service/model"
	"testing"
)

func TestDbGetBooks(t *testing.T) {
	service := model.Service{
		Cfg: config.InitService("config.json"),
	}
	database.ConnectionDb(service)
	res, err := database.GetAuthorsByBook("Ken")
	if err != nil {
		t.Error("Error get", err)
	}
	if len(res) != 2 || res[0] != "Mark Tven" || res[1] != "Sem Len" {
		t.Error("Geted incorrect data")
	}
}

func TestDbGetAuthors(t *testing.T) {
	service := model.Service{
		Cfg: config.InitService("config.json"),
	}
	database.ConnectionDb(service)
	res, err := database.GetBooksByAuthor("Mark Tven")
	if err != nil {
		t.Error("Error get", err)
	}
	if len(res) != 3 || res[0] != "Ben" || res[1] != "Sen" || res[2] != "Ken" {
		t.Error("Geted incorrect data")
	}
}
