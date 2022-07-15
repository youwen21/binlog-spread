package service

import (
	"binlog_spread/app/dal"
	"binlog_spread/app/dto"
)

type stateDirectionSrv struct{}

var (
	StateDirectionSrv = &stateDirectionSrv{}
)

func (srv *stateDirectionSrv) Query(f *dto.StateDirectionForm) (map[string]interface{}, error) {
	ret, _ := dal.StateDirectionDAO.Query(f)
	var stateValueList []string
	for _, v := range ret.List {
		stateValueList = append(stateValueList, v.StateFrom, v.StateTo)
	}

	stateMap, _ := dal.StateDAO.GetMultiByValue(f.StateClassId, stateValueList)

	var retList []dto.StateDirectionText
	for _, v := range ret.List {
		text := dto.StateDirectionText{StateDirection: v}
		if state, ok := stateMap[v.StateFrom]; ok {
			text.StateFromDesc = state.StateValueDesc
		}
		if state, ok := stateMap[v.StateTo]; ok {
			text.StateToDesc = state.StateValueDesc
		}

		retList = append(retList, text)
	}

	return map[string]interface{}{"count": ret.Count, "list": retList}, nil
}
