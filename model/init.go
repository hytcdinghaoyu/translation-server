package models

import (
	"bytes"
	"encoding/json"
	"fg-admin/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/patrickmn/go-cache"
	"github.com/pelletier/go-toml"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	DB       = New()
	MemCache = cache.New(10*time.Minute, 1*time.Hour)
)

func HttpPost(url string, body []byte) map[string]interface{} {
	rsp, err := http.Post(url, "application/json", bytes.NewReader(body))
	ret := make(map[string]interface{})
	if err != nil {
		fmt.Println("http request error", err)
		return ret
	}

	defer rsp.Body.Close()
	rspBody, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		fmt.Println("read error", err)
	}

	err = json.Unmarshal(rspBody, &ret)
	if err != nil {
		fmt.Println("json decode error", err)
	}

	return ret

}


// New setup DB connection
func New() *gorm.DB {

	driver := config.Conf.Get("database.driver").(string)
	configTree := config.Conf.Get(driver).(*toml.Tree)
	userName := configTree.Get("username").(string)
	password := configTree.Get("password").(string)
	databaseName := configTree.Get("db").(string)
	connect := userName + ":" + password + "@/" + databaseName + "?charset=utf8&parseTime=True&loc=Local"

	DB, err := gorm.Open(driver, connect)

	if err != nil {
		panic(fmt.Sprintf("failed to connect to database, got err=%+v", err))
	}

	return DB
}

func GetAll(string, orderBy string, page, limit int) *gorm.DB {
	TDB := DB
	if len(string) > 0 {
		TDB = TDB.Where("name LIKE  ?", "%"+string+"%")
	}

	if len(orderBy) > 0 {
		TDB = TDB.Order(orderBy)
	} else {
		TDB = TDB.Order("created_at desc")
	}

	if page > 0 {
		TDB = TDB.Offset((page - 1) * limit)
	}

	if limit > 0 {
		TDB = TDB.Limit(limit)
	}

	return TDB.Debug()
}
