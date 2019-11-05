package geoip

import (
	"errors"
	"github.com/esrrhs/go-engine/src/common"
	"github.com/oschwald/geoip2-golang"
	"net"
)

var gdb *geoip2.Reader

func Load(file string) {

	if len(file) <= 0 {
		file = common.GetDataDir() + "/geoip/" + "GeoLite2-Country.mmdb"
	}

	db, err := geoip2.Open(file)
	if err != nil {
		panic(err)
	}
	gdb = db
}

func GetCountryIsoCode(ipaddr string) (string, error) {

	ip := net.ParseIP(ipaddr)
	if ip == nil {
		return "", errors.New("ip " + ipaddr + " ParseIP nil")
	}
	record, err := gdb.City(ip)
	if err != nil {
		return "", err
	}

	return record.Country.IsoCode, nil
}

func GetCountryName(ipaddr string) (string, error) {

	ip := net.ParseIP(ipaddr)
	if ip == nil {
		return "", errors.New("ip " + ipaddr + "ParseIP nil")
	}
	record, err := gdb.City(ip)
	if err != nil {
		return "", err
	}

	return record.Country.Names["en"], nil
}
