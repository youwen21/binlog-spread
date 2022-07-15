package dal

import (
	"binlog_spread/app/dto"
	"binlog_spread/app/models"
	"binlog_spread/comps"
	"gorm.io/gorm"
	"time"
)

type statClassDAO struct {
}

var (
	StatClassDAO = &statClassDAO{}
)

func (d *statClassDAO) Query(f *dto.StateClassForm) (*dto.StateClassQueryResult, error) {
	session := d.newSession()

	var total int64
	var list []models.StateClass

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

	return &dto.StateClassQueryResult{Count: total, List: list}, nil
}

func (d *statClassDAO) GetList(f *dto.StateClassForm) ([]*models.StateClass, error) {
	var results []*models.StateClass

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

func (d *statClassDAO) GetMulti(idList []int) (map[int]models.StateClass, error) {
	var cMap = make(map[int]models.StateClass, 0)
	var results []models.StateClass
	session := d.newSession()
	query := session.Where("state_class_id IN ?", idList)
	err := query.Find(&results).Error
	if err != nil {
		return nil, err
	}
	for _, stateClass := range results {
		cMap[stateClass.StateClassId] = stateClass
	}
	return cMap, nil
}

func (d *statClassDAO) Get(id int) (*models.StateClass, error) {
	info := &models.StateClass{}
	session := d.newSession()
	if err := session.Where(" state_class_id= ?", id).First(info).Error; err != nil {
		return nil, err
	}

	return info, nil
}

func (d *statClassDAO) Insert(m *models.StateClass) error {
	session := d.newSession()
	err := session.Create(m).Error
	return err
}

func (d *statClassDAO) Update(m *models.StateClass) error {
	session := d.newSession()
	err := session.Where("state_class_id = ?", m.StateClassId).UpdateColumns(m).Error
	return err
}

func (d *statClassDAO) Delete(id int) error {
	session := d.newSession()
	err := session.Where("state_class_id = ?", id).Updates(map[string]interface{}{"is_deleted": 1, "deleted_at": time.Now()}).Error
	return err
}

func (d *statClassDAO) RealDelete(id int) error {
	session := d.newSession()
	err := session.Where("state_class_id = ?", id).Delete(models.StateClass{}).Error
	return err
}

func (d *statClassDAO) newEngine() *gorm.DB {
	return comps.GetDb()
}

func (d *statClassDAO) newSession() *gorm.DB {
	return comps.GetSession().Table("state_class")
}
