package dal

import (
	"binlog_spread/app/dto"
	"binlog_spread/app/models"
	"binlog_spread/comps"
	"gorm.io/gorm"
)

type adminDAO struct{}

var (
	AdminDAO = &adminDAO{}

	//queryErr = errors.New("查询失败")
)

func (d *adminDAO) Query(menu *models.Administrator) (*models.Administrator, error) {
	return nil, nil
}

func (d *adminDAO) Get(id int) (*models.Administrator, error) {
	session := d.newSession()
	info := models.Administrator{}

	err := session.Where("administrator_id = ?", id).First(&info).Error
	if err != nil {
		return nil, err
	}
	return &info, nil
}

func (d *adminDAO) GetByUsername(username string) (*models.Administrator, error) {
	session := d.newSession()
	info := models.Administrator{}

	err := session.Where("username = ?", username).First(&info).Error
	if err != nil {
		return nil, err
	}
	return &info, nil
}

func (d *adminDAO) GetList(f *dto.MenuForm) ([]*models.Administrator, error) {
	var results []*models.Administrator

	session := d.newSession()
	if f.IsDeleted != 0 {
		session.Where("is_deleted= ?", f.IsDeleted)
	}
	err := session.Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (d *adminDAO) Insert(menu *models.Administrator) (*models.Administrator, error) {
	return menu, nil
}

func (d *adminDAO) Update(menu *models.Administrator) (*models.Administrator, error) {
	return menu, nil
}

func (d *adminDAO) Delete(id int) (int64, error) {
	return 0, nil
}

func (d *adminDAO) newEngine() *gorm.DB {
	return comps.GetDb()
}

func (d *adminDAO) newSession() *gorm.DB {
	return comps.GetSession().Table("administrator")
}
