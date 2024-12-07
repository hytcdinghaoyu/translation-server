package config

import (
	"encoding/csv"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/pelletier/go-toml"
	"github.com/tealeg/xlsx/v3"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

var (
	Conf                = New()
	ItemConfMap         = make(map[int]string)
)

type JwtClaims struct {
	*jwt.StandardClaims
	UserInfo
}

type UserInfo struct {
	Uid      uint
	Username string
	RoleId   uint
	AuthStr  string
}


// 前端渲染option
type Option struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func Reload(){
}

/**
 * 返回单例实例
 * @method New
 */
func New() *toml.Tree {
	config, err := toml.LoadFile("./config.toml")

	if err != nil {
		fmt.Println("TomlError ", err.Error())
	}

	return config
}

func NewItemConf() []Option {
	itemsList := ReadXlsxToOption("Item", "Item", "Name", "ID")
	equipList := ReadXlsxToOption("Equip", "Equip", "Name", "ID")
	sliItem := make([]Option, len(itemsList)+len(equipList))
	copy(sliItem, itemsList)
	for i := 0; i < len(equipList); i++ {
		sliItem[len(itemsList)+i] = equipList[i]
	}

	for _, v := range sliItem {
		ItemConfMap[v.Value] = v.Name
	}
	return sliItem
}

func NewItemConfFairy() []Option {
	itemsList := NewSliConf("Items", 5)
	return itemsList
}

func NewSliConf(name string, offset int) []Option {
	path := fmt.Sprintf("./upload/%s.csv", name)
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	r := csv.NewReader(strings.NewReader(string(dat[:])))

	var sliItem []Option
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		id, err := strconv.Atoi(record[0])
		if id == 0 {
			continue
		}

		if err != nil {
			panic(err)
		}

		i := Option{Value: id, Name: record[offset]}
		sliItem = append(sliItem, i)

		ItemConfMap[id] = record[offset]
	}
	return sliItem
}

func NewConfMap(name string, offset int) map[int]string {

	path := fmt.Sprintf("./upload/%s.csv", name)
	dat, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}
	r := csv.NewReader(strings.NewReader(string(dat[:])))

	m := map[int]string{}
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		id, err := strconv.Atoi(record[0])
		if id == 0 {
			continue
		}

		if err != nil {
			panic(err)
		}

		m[id] = record[offset]
	}

	return m
}

func ReadXlsxToMap(name string,sheetName string, key string, val string) map[int]string {
	path := fmt.Sprintf("./upload/%s.xlsx", name)
	fmt.Println(path)
	wb, err := xlsx.OpenFile(path)
	if err != nil {
		panic(err)

	}
	sheet, ok := wb.Sheet[sheetName]
	if !ok {
		panic("sheet name not exist")
	}

	fieldRow, _ := sheet.Row(2)
	var keyIndex, valIndex int
	err = fieldRow.ForEachCell(func (c *xlsx.Cell) error {
		value, err := c.FormattedValue()
		if err != nil {
			fmt.Println(err.Error())
		} else if value == key {
			keyIndex, _ = c.GetCoordinates()
		} else if value == val {
			valIndex, _ = c.GetCoordinates()
		}
		return err
	})

	ret := make(map[int]string, sheet.MaxRow)
	err = sheet.ForEachRow(func (r *xlsx.Row) error {
		i := r.GetCoordinate()
		if i < 4 {
			return nil
		}

		var err error
		keyVal, err := r.GetCell(keyIndex).Int()
		valVal, err := r.GetCell(valIndex).FormattedValue()
		if err != nil || keyVal <= 0 {
			return err
		}
		ret[keyVal] = valVal
		return err
	})

	return ret
}

func ReadXlsxToOption(name string,sheetName string, key string, val string) []Option {
	path := fmt.Sprintf("./upload/%s.xlsx", name)
	wb, err := xlsx.OpenFile(path)
	if err != nil {
		panic(err)
	}
	sheet, ok := wb.Sheet[sheetName]
	if !ok {
		panic("sheet name not exist")
	}

	fieldRow, _ := sheet.Row(2)
	var keyIndex, valIndex int
	err = fieldRow.ForEachCell(func (c *xlsx.Cell) error {
		value, err := c.FormattedValue()
		if err != nil {
			fmt.Println(err.Error())
		} else if value == key {
			keyIndex, _ = c.GetCoordinates()
		} else if value == val {
			valIndex, _ = c.GetCoordinates()
		}
		return err
	})

	ret := make([]Option, 0, sheet.MaxRow)
	err = sheet.ForEachRow(func (r *xlsx.Row) error {
		i := r.GetCoordinate()
		if i < 4 {
			return nil
		}

		var err error
		keyVal, err := r.GetCell(keyIndex).FormattedValue()
		valVal, err := r.GetCell(valIndex).Int()
		if err != nil {
			return err
		}
		ret = append(ret, Option{Name: keyVal, Value: valVal})
		return err
	})

	return ret
}

