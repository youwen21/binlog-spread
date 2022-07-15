package service

import (
	"binlog_spread/app/dal"
	"binlog_spread/app/dto"
	"binlog_spread/app/models"
	"errors"
	"github.com/gin-gonic/gin"
	"sort"
	"strings"
)

type apiBinlogSrv struct{}

var (
	ApiBinlogSrv = &apiBinlogSrv{}
)

func (srv *apiBinlogSrv) Query(f *dto.ApiBinlogForm) (map[string]interface{}, error) {
	ret, _ := dal.ApiBinlogDAO.Query(f)

	var retList []dto.ApiBinlogText
	for _, v := range ret.List {
		d := dto.ApiBinlogText{ApiBinlog: v}
		d.EventName = dto.EventType(v.EventType).String()

		retList = append(retList, d)
	}

	return map[string]interface{}{"count": ret.Count, "list": retList}, nil
}

func (srv *apiBinlogSrv) Diff(c *gin.Context, searchParams dto.BinlogDiffSearch) ([]dto.DiffRow, error) {
	streamListA, streamListB, err := srv.getStreamAB(searchParams)
	if err != nil {
		return nil, errors.New("查询错误")
	}

	streamA := dto.SortStreams(streamListA)
	streamB := dto.SortStreams(streamListB)
	sort.Sort(streamA)
	sort.Sort(streamB)

	streamA.Process()
	streamB.Process()

	result := combineList(streamA, streamB)

	return result, nil
}

func (srv *apiBinlogSrv) getStreamAB(searchParams dto.BinlogDiffSearch) ([]*models.ApiBinlog, []*models.ApiBinlog, error) {
	if searchParams.ApiA == "" && searchParams.ApiB == "" {
		binlogList, _ := dal.ApiBinlogDAO.GetList(map[string]interface{}{"id_list": strings.Split(searchParams.BinlogIds, ",")})
		return binlogList, []*models.ApiBinlog{}, nil
	}

	streamListA, _ := ApiDefineSrv.GetBinlogListByApiId(searchParams.ApiA)

	var streamListB []*models.ApiBinlog
	if searchParams.ApiB == "" {
		streamListB, _ = dal.ApiBinlogDAO.GetList(map[string]interface{}{"id_list": strings.Split(searchParams.BinlogIds, ",")})
	} else {
		streamListB, _ = ApiDefineSrv.GetBinlogListByApiId(searchParams.ApiB)
	}

	return streamListA, streamListB, nil
}

func combineList(a dto.SortStreams, b dto.SortStreams) []dto.DiffRow {
	var diffList []dto.DiffRow

	for _, v := range a {
		// diffEntityRow := &diffEntity{A: v}
		diffEntityRow := dto.DiffRow{}
		diffEntityRow.A = dto.GenApiBinlogText(v)
		diffEntityRow.B = dto.GenApiBinlogText(b.ExtractAndRemove(v.DbName, v.TableName))
		diffList = append(diffList, diffEntityRow)
	}

	for _, v := range b {
		diffEntityRow := dto.DiffRow{A: dto.GenApiBinlogText(b.GetDefault()), B: dto.GenApiBinlogText(v)}
		diffList = append(diffList, diffEntityRow)
	}

	return diffList
}
