package dal

import (
	"binlog_spread/app/dto"
	"binlog_spread/app/models"
	"binlog_spread/comps"
	"gorm.io/gorm"
	"time"
)

type stateDAO struct {
}

var (
	StateDAO = &stateDAO{}
)

func (d *stateDAO) Query(f *dto.StateForm) (*dto.StateResult, error) {
	session := d.newSession()

	var total int64
	var list []models.State

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

	return &dto.StateResult{Count: total, List: list}, nil
}

func (d *stateDAO) GetList(f *dto.StateForm) ([]*models.State, error) {
	var results []*models.State

	session := d.newSession()
	if f.IsDeleted != -1 {
		session.Where("is_deleted= ?", f.IsDeleted)
	}

	if f.StateClassId != 0 {
		session.Where("state_class_id = ?", f.StateClassId)
	}

	err := session.Limit(f.Limit()).Offset(f.Offset()).Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (d *stateDAO) Get(id int) (*models.State, error) {
	info := &models.State{}
	session := d.newSession()
	if err := session.Where(" state_id= ?", id).First(info).Error; err != nil {
		return nil, err
	}

	return info, nil
}

func (d *stateDAO) GetMulti(idList []int) (map[int]models.State, error) {
	var mMap = make(map[int]models.State, 0)
	var results []models.State
	session := d.newSession()
	query := session.Where("state_id IN ?", idList)
	err := query.Find(&results).Error
	if err != nil {
		return nil, err
	}
	for _, state := range results {
		mMap[state.StateId] = state
	}
	return mMap, nil
}

func (d *stateDAO) GetMultiByValue(classId int, stateValueList []string) (map[string]models.State, error) {
	var mMap = make(map[string]models.State, 0)
	var results []models.State
	session := d.newSession()
	query := session.Where("state_class_id = ?", classId).Where("state_value IN ?", stateValueList)
	err := query.Find(&results).Error
	if err != nil {
		return nil, err
	}
	for _, state := range results {
		mMap[state.StateValue] = state
	}
	return mMap, nil
}

func (d *stateDAO) Insert(m *models.State) error {
	session := d.newSession()
	err := session.Create(m).Error
	return err
}

func (d *stateDAO) Update(m *models.State) error {
	session := d.newSession()
	err := session.Where("state_id = ?", m.StateId).UpdateColumns(m).Error
	return err
}

func (d *stateDAO) Delete(id int) error {
	session := d.newSession()
	err := session.Where("state_id = ?", id).Updates(map[string]interface{}{"is_deleted": 1, "deleted_at": time.Now()}).Error
	return err
}

func (d *stateDAO) RealDelete(id int) error {
	session := d.newSession()
	err := session.Where("state_id = ?", id).Delete(models.State{}).Error
	return err
}

func (d *stateDAO) newEngine() *gorm.DB {
	return comps.GetDb()
}

func (d *stateDAO) newSession() *gorm.DB {
	return comps.GetSession().Table("state")
}
