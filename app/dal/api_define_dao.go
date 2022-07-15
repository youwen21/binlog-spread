package dal

import (
	"binlog_spread/app/dto"
	"binlog_spread/app/models"
	"binlog_spread/comps"
	"gorm.io/gorm"
	"time"
)

type apiDefineDAO struct {
}

var (
	ApiDefineDAO = &apiDefineDAO{}
)

func (d *apiDefineDAO) Query(f *dto.PageParam) (*dto.ApiDefineResult, error) {
	session := d.newSession()

	var total int64
	var list []models.ApiDefine

	if err := session.Count(&total).Error; err != nil {
		return nil, err
	}

	if len(f.OrderBy) > 0 {
		for _, v := range f.OrderBy {
			session.Order(v)
		}
	}

	if err := session.Limit(f.Limit()).Offset(f.Offset()).Find(&list).Error; err != nil {
		return nil, err
	}

	return &dto.ApiDefineResult{Count: total, List: list}, nil
}

func (d *apiDefineDAO) GetList(m map[string]interface{}) ([]*models.ApiDefine, error) {
	var results []*models.ApiDefine

	session := d.newSession()

	if v, ok := m["id_list"]; ok {
		session.Where("api_define_id in ?", v)
	}

	err := session.Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (d *apiDefineDAO) GetMulti(idList []int) (map[int]models.ApiDefine, error) {
	var cMap = make(map[int]models.ApiDefine, 0)
	var results []models.ApiDefine
	session := d.newSession()
	query := session.Where("api_define_id IN ?", idList)
	err := query.Find(&results).Error
	if err != nil {
		return nil, err
	}
	for _, v := range results {
		cMap[v.ApiDefineId] = v
	}
	return cMap, nil
}

func (d *apiDefineDAO) Get(id int) (*models.ApiDefine, error) {
	info := &models.ApiDefine{}
	session := d.newSession()
	if err := session.Where(" api_define_id= ?", id).First(info).Error; err != nil {
		return nil, err
	}

	return info, nil
}

func (d *apiDefineDAO) Insert(m *models.ApiDefine) error {
	session := d.newSession()
	err := session.Create(m).Error
	return err
}

func (d *apiDefineDAO) Update(m *models.ApiDefine) error {
	session := d.newSession()
	err := session.Where("api_define_id = ?", m.ApiDefineId).UpdateColumns(m).Error
	return err
}

func (d *apiDefineDAO) Delete(id int) error {
	session := d.newSession()
	err := session.Where("api_define_id = ?", id).Updates(map[string]interface{}{"is_deleted": 1, "deleted_at": time.Now()}).Error
	return err
}

func (d *apiDefineDAO) RealDelete(id int) error {
	session := d.newSession()
	err := session.Where("api_define_id = ?", id).Delete(models.ApiDefine{}).Error
	return err
}

func (d *apiDefineDAO) newEngine() *gorm.DB {
	return comps.GetDb().Table("api_define")
}

func (d *apiDefineDAO) newSession() *gorm.DB {
	return comps.GetSession().Table("api_define")
}
