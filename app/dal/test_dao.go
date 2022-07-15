package dal

import (
	"binlog_spread/app/models"
	"binlog_spread/comps"
	"gorm.io/gorm"
)

type testDAO struct {
}

var (
	TestDAO = &testDAO{}

	//queryErr = errors.New("查询失败")
)

func (d *testDAO) Query() ([]map[string]interface{}, error) {
	session := d.newSession()

	var total int64
	var list []map[string]interface{}

	if err := session.Count(&total).Error; err != nil {
		return nil, err
	}

	if err := session.Limit(10).Offset(0).Find(&list).Error; err != nil {
		return nil, err
	}

	return list, nil
}

func (d *testDAO) Get(id int) (map[string]interface{}, error) {
	var info map[string]interface{}
	session := d.newSession()
	if err := session.Where(" menu_id= ?", id).First(info).Error; err != nil {
		return nil, err
	}

	return info, nil
}

func (d *testDAO) Insert(menu *models.Menu) error {
	session := d.newSession()
	err := session.Create(&menu).Error
	return err
}

func (d *testDAO) Update(menu *models.Menu) error {
	session := d.newSession()
	err := session.Where("menu_id = ?", menu.MenuId).UpdateColumns(menu).Error
	return err
}

func (d *testDAO) newEngine() *gorm.DB {
	return comps.GetDb()
}

func (d *testDAO) newSession() *gorm.DB {
	return comps.GetSession().Table("menu").Clauses()
}
