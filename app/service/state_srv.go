package service

import (
	"binlog_spread/app/dal"
	"binlog_spread/app/dto"
)

type stateSrv struct{}

var (
	StateSrv = &stateSrv{}
)

func (srv *stateSrv) Query(f *dto.StateForm) (map[string]interface{}, error) {
	ret, _ := dal.StateDAO.Query(f)
	var stateClassIdList []int
	for _, v := range ret.List {
		stateClassIdList = append(stateClassIdList, v.StateClassId)
	}

	stateClassMap, _ := dal.StatClassDAO.GetMulti(stateClassIdList)

	var retList []dto.StateText
	for _, v := range ret.List {
		state := dto.StateText{State: v}
		if stateClass, ok := stateClassMap[v.StateClassId]; ok {
			state.DbName = stateClass.DbName
			state.TableName = stateClass.TableName
			state.FieldName = stateClass.FieldName
			state.StateClassName = stateClass.StateName
		}

		retList = append(retList, state)
	}

	return map[string]interface{}{"count": ret.Count, "list": retList}, nil
}
