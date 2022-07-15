package state

import (
	"binlog_spread/app/dal"
	"binlog_spread/app/dto"
	"binlog_spread/app/models"
	"binlog_spread/cmd/command/handle_binlog/common"
	"errors"
)

var (
	Classes    = make(map[string]int)
	Directions = make(map[int][]models.StateDirection)
)

func InitState() {
	InitStateClass()
	InitStateDirection()
}

func InitStateClass() {
	form := dto.StateClassForm{}
	form.IsDeleted = 0
	list, _ := dal.StatClassDAO.GetList(&form)

	for _, v := range list {
		key := common.GetKey(v.DbName, v.TableName, v.FieldName)
		value := v.StateClassId
		Classes[key] = value
	}
}

func InitStateDirection() {
	form := dto.StateDirectionForm{}
	form.IsDeleted = 0
	list, _ := dal.StateDirectionDAO.GetList(&form)

	for _, v := range list {
		key := v.StateClassId
		SetDirection(key, *v)
	}
}

func SetDirection(classId int, row models.StateDirection) {
	_, ok := Directions[classId]
	if !ok {
		Directions[classId] = []models.StateDirection{}
	}

	Directions[classId] = append(Directions[classId], row)
}

func CheckStatDirection(classId int, from string, to string) (bool, error) {
	list, ok := Directions[classId]
	if !ok {
		return false, errors.New("state class direction not exist")
	}

	for _, v := range list {
		if v.StateFrom == from && v.StateTo == to {
			return true, nil
		}
	}

	return false, errors.New("direction not defined")
}

func GetStatClassId(dbName string, tableName string, fieldName string) (int, error) {
	key := common.GetKey(dbName, tableName, fieldName)
	classId, ok := Classes[key]
	if !ok {
		return 0, errors.New("state class not defined")
	}

	return classId, nil
}
