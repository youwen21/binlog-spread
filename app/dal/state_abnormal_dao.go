package dal

import (
	"binlog_spread/app/dto"
	"binlog_spread/app/models"
	"binlog_spread/comps"
	"gorm.io/gorm"
	"time"
)

type stateAbnormalDAO struct {
}

var (
	StateAbnormalDAO = &stateAbnormalDAO{}
)

func (d *stateAbnormalDAO) Query(f *dto.StateAbnormalForm) (*dto.StateAbnormalQueryResult, error) {
	session := d.newSession()

	var total int64
	var list []models.StateAbnormal

	if f.IsDeleted != -1 {
		session.Where("is_deleted = ?", f.IsDeleted)
	}

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

	return &dto.StateAbnormalQueryResult{Count: total, List: list}, nil
}

func (d *stateAbnormalDAO) GetList(f *dto.StateAbnormalForm) ([]*models.StateAbnormal, error) {
	var results []*models.StateAbnormal

	session := d.newSession()
	if f.IsDeleted != -1 {
		session.Where("is_deleted= ?", f.IsDeleted)
	}
	err := session.Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (d *stateAbnormalDAO) Get(id int) (*models.StateAbnormal, error) {
	info := &models.StateAbnormal{}
	session := d.newSession()
	if err := session.Where(" state_abnormal_id= ?", id).First(info).Error; err != nil {
		return nil, err
	}

	return info, nil
}

func (d *stateAbnormalDAO) Insert(m *models.StateAbnormal) error {
	session := d.newSession()
	err := session.Create(m).Error
	return err
}

func (d *stateAbnormalDAO) Update(m *models.StateAbnormal) error {
	session := d.newSession()
	err := session.Where("state_abnormal_id = ?", m.StateAbnormalId).UpdateColumns(m).Error
	return err
}

func (d *stateAbnormalDAO) Delete(id int) error {
	session := d.newSession()
	err := session.Where("state_abnormal_id = ?", id).Updates(map[string]interface{}{"is_deleted": 1, "deleted_at": time.Now()}).Error
	return err
}

func (d *stateAbnormalDAO) RealDelete(id int) error {
	session := d.newSession()
	err := session.Where("state_abnormal_id = ?", id).Delete(models.StateAbnormal{}).Error
	return err
}

func (d *stateAbnormalDAO) newEngine() *gorm.DB {
	return comps.GetDb()
}

func (d *stateAbnormalDAO) newSession() *gorm.DB {
	return comps.GetSession().Table("state_abnormal")
}
