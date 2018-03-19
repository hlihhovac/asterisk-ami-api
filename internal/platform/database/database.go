package database

import (
	"fmt"
	"https://github.com/hlihhovac/asterisk-ami-api/tree/master/internal/utils/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

type CDR struct {
	Calldate     time.Time `json:"calldate"`
	Formateddate string    `json:"formateddate"`
	Src          string    `json:"src"`
	Dst          string    `json:"dst"`
	Dcontext     string    `json:"dcontext"`
	Channel      string    `json:"channel"`
	Disposition  string    `json:"disposition"`
	Dstchannel   string    `json:"dstchannel"`
	Lastapp      string    `json:"lastapp"`
	Duration     int       `json:"duration"`
	Billsec      int       `json:"billsec"`
	Uniqueid     string    `json:"uniqueid"`
	Actionid     string    `json:"actionid"`
}

var dbInstance *gorm.DB

func Connect(tomlConfig *config.TomlConfig) (*gorm.DB, error) {
	if dbInstance == nil {
		connString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			tomlConfig.DB.Username, tomlConfig.DB.Password, tomlConfig.DB.Host, tomlConfig.DB.Database)

		database, err := gorm.Open("mysql", connString)
		if err != nil {
			return nil, err
		}

		database.LogMode(tomlConfig.DB.Debug)
		dbInstance = database
	}

	return dbInstance, nil
}

func GetStatByMSISDN(MSISDN, startdate, enddate string) []CDR {
	query, err := Connect(config.GetConfig())
	if err != nil {
		log.Fatal(err)
	}

	cdrs := []CDR{}

	if startdate == "" || enddate == "" {
		query.Table("cdr").Where("src=? or dst=?", MSISDN, MSISDN).Find(&cdrs)
	} else {
		query.Table("cdr").Where("calldate between ? and ? and (src=? or dst=?)",
			startdate+" 00:00:00", enddate+" 23:59:00", MSISDN, MSISDN).Find(&cdrs)
	}

	return cdrs
}

func GetStatByActionID(MSISDN, actionID string) []CDR {
	query, err := Connect(config.GetConfig())
	if err != nil {
		log.Fatal(err)
	}

	cdrs := []CDR{}

	query.Table("cdr").Where("(src=? or dst=?) and actionid=?", MSISDN, MSISDN, actionID).
		Find(&cdrs)

	return cdrs
}
