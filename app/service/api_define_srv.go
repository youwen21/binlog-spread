package service

import (
	"binlog_spread/app/dal"
	"binlog_spread/app/models"
	"strconv"
	"strings"
)

type apiDefineSrv struct{}

var (
	ApiDefineSrv = &apiDefineSrv{}
)

func (srv *apiDefineSrv) GetBinlogListByApiId(apiDefineId string) ([]*models.ApiBinlog, error) {
	id, _ := strconv.Atoi(apiDefineId)
	defineInfo, _ := dal.ApiDefineDAO.Get(id)
	ids := strings.Split(defineInfo.StreamIds, ",")

	var idsInt []int
	for _, v := range ids {
		intV, _ := strconv.Atoi(v)
		idsInt = append(idsInt, intV)
	}
	binlogList, err := dal.ApiBinlogDAO.GetList(map[string]interface{}{"id_list": idsInt})
	return binlogList, err
}
