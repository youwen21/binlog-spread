package dal

import (
	"binlog_spread/app/dto"
	"binlog_spread/app/models"
	"binlog_spread/comps"
	"gorm.io/gorm"
	"time"
)

type stateDirectionDAO struct {
}

var (
	StateDirectionDAO = &stateDirectionDAO{}
)

func (d *stateDirectionDAO) Query(f *dto.StateDirectionForm) (*dto.StateDirectionQueryResult, error) {
	session := d.newSession()

	var total int64
	var list []models.StateDirection

	if f.IsDeleted != -1 {
		session.Where("is_deleted = ?", f.IsDeleted)
	}

	if f.StateClassId != 0 {
		session.Where("state_class_id = ?", f.StateClassId)
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

	return &dto.StateDirectionQueryResult{Count: total, List: list}, nil
}

func (d *stateDirectionDAO) GetList(f *dto.StateDirectionForm) ([]*models.StateDirection, error) {
	var results []*models.StateDirection

	session := d.newSession()
	if f.IsDeleted != -1 {
		session.Where("is_deleted= ?", f.IsDeleted)
	}

	if f.StateClassId != 0 {
		session.Where("state_class_id = ?", f.StateClassId)
	}

	err := session.Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (d *stateDirectionDAO) GetMulti(idList []int) (map[int]models.StateDirection, error) {
	var cMap = make(map[int]models.StateDirection, 0)
	var results []models.StateDirection
	session := d.newSession()
	query := session.Where("state_direction_id IN ?", idList)
	err := query.Find(&results).Error
	if err != nil {
		return nil, err
	}
	for _, stateClass := range results {
		cMap[stateClass.StateClassId] = stateClass
	}
	return cMap, nil
}

func (d *stateDirectionDAO) Get(id int) (*models.StateDirection, error) {
	info := &models.StateDirection{}
	session := d.newSession()
	if err := session.Where(" state_direction_id= ?", id).First(info).Error; err != nil {
		return nil, err
	}

	return info, nil
}

func (d *stateDirectionDAO) Insert(m *models.StateDirection) error {
	session := d.newSession()
	err := session.Create(m).Error
	return err
}

func (d *stateDirectionDAO) Update(m *models.StateDirection) error {
	session := d.newSession()
	err := session.Where("state_direction_id = ?", m.StateClassId).UpdateColumns(m).Error
	return err
}

func (d *stateDirectionDAO) Delete(id int) error {
	session := d.newSession()
	err := session.Where("state_direction_id = ?", id).Updates(map[string]interface{}{"is_deleted": 1, "deleted_at": time.Now()}).Error
	return err
}

func (d *stateDirectionDAO) RealDelete(id int) error {
	session := d.newSession()
	err := session.Where("state_direction_id = ?", id).Delete(models.StateDirection{}).Error
	return err
}

func (d *stateDirectionDAO) newEngine() *gorm.DB {
	return comps.GetDb()
}

func (d *stateDirectionDAO) newSession() *gorm.DB {
	return comps.GetSession().Table("state_direction")
}
