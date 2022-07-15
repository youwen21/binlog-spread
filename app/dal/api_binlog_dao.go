package dal

import (
	"binlog_spread/app/dto"
	"binlog_spread/app/models"
	"binlog_spread/comps"
	"gorm.io/gorm"
	"time"
)

type apiBinlogDAO struct {
}

var (
	ApiBinlogDAO = &apiBinlogDAO{}
)

func (d *apiBinlogDAO) Query(f *dto.ApiBinlogForm) (*dto.ApiBinlogResult, error) {
	session := d.newSession()

	var total int64
	var list []models.ApiBinlog


	if len(f.OrderBy) > 0 {
		for _, v := range f.OrderBy {
			session.Order(v)
		}
	}

	if err := session.Count(&total).Error; err != nil {
		return nil, err
	}
	if err := session.Limit(f.Limit()).Offset(f.Offset()).Find(&list).Error; err != nil {
		return nil, err
	}

	return &dto.ApiBinlogResult{Count: total, List: list}, nil
}

func (d *apiBinlogDAO) GetList(m map[string]interface{}) ([]*models.ApiBinlog, error) {
	var results []*models.ApiBinlog

	session := d.newSession()

	if v, ok := m["is_delete"]; ok {
		session.Where("is_deleted= ?", v)
	}

	if v, ok := m["id_list"]; ok {
		session.Where("api_binlog_id in ?", v)
	}

	if v, ok := m["state_class_id"]; ok {
		session.Where("state_class_id = ?", v)
	}

	limit := 2000
	offset := 0
	if v, ok := m["pageSize"]; ok {
		limit = v.(int)
	}
	if v, ok := m["pageNumber"]; ok {
		size := v.(int)
		if size != 0 {
			offset = (size - 1) * limit
		}
	}

	err := session.Limit(limit).Offset(offset).Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (d *apiBinlogDAO) Get(id int) (*models.ApiBinlog, error) {
	info := &models.ApiBinlog{}
	session := d.newSession()
	if err := session.Where(" api_binlog_id= ?", id).First(info).Error; err != nil {
		return nil, err
	}

	return info, nil
}

func (d *apiBinlogDAO) GetMulti(idList []int) (map[int]models.ApiBinlog, error) {
	var mMap = make(map[int]models.ApiBinlog, 0)
	var results []models.ApiBinlog
	session := d.newSession()
	query := session.Where("api_binlog_id IN ?", idList)
	err := query.Find(&results).Error
	if err != nil {
		return nil, err
	}
	for _, v := range results {
		mMap[v.ApiBinlogId] = v
	}
	return mMap, nil
}

func (d *apiBinlogDAO) Insert(m *models.ApiBinlog) error {
	session := d.newSession()
	err := session.Create(m).Error
	return err
}

func (d *apiBinlogDAO) Update(m *models.ApiBinlog) error {
	session := d.newSession()
	err := session.Where("api_binlog_id = ?", m.ApiBinlogId).UpdateColumns(m).Error
	return err
}

func (d *apiBinlogDAO) Delete(id int) error {
	session := d.newSession()
	err := session.Where("api_binlog_id = ?", id).Updates(map[string]interface{}{"is_deleted": 1, "deleted_at": time.Now()}).Error
	return err
}

func (d *apiBinlogDAO) RealDelete(id int) error {
	session := d.newSession()
	err := session.Where("api_binlog_id = ?", id).Delete(models.ApiBinlog{}).Error
	return err
}

func (d *apiBinlogDAO) newEngine() *gorm.DB {
	return comps.GetDb()
}

func (d *apiBinlogDAO) newSession() *gorm.DB {
	return comps.GetSession().Table("api_binlog")
}
