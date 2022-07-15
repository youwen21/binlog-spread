package service

import (
	"binlog_spread/app/dal"
	"binlog_spread/app/dto"
	"binlog_spread/app/models"
	"errors"
	"strconv"
	"strings"
)

type erSrv struct{}

var (
	ErSrv = &erSrv{}
)

func (srv *erSrv) Query(search *dto.ErSearch) (map[string]interface{}, error) {
	var list []*models.ApiBinlog
	var err error
	switch search.Type {
	case "stream_ids":
		list, err = dal.ApiBinlogDAO.GetList(map[string]interface{}{"id_list": search.SearchIdSlice()})
		break
	case "event_id":
		list, err = ApiDefineSrv.GetBinlogListByApiId(search.Search)
		break
	//case "transaction_id":
	//	list, err = dao.GetStreamListByTransactionId(search.Search)
	//	break
	default:
		return nil, errors.New("不支持类型" + search.Type)
	}

	data := processStreamListToErData(list, search.Scope, true)

	return data, err
}

type EntityItem struct {
	Name  string `json:"name"`
	IsKey bool   `json:"isKey"`
}

type Entity struct {
	Key   string       `json:"key"`
	Items []EntityItem `json:"items"`
}

type linkData struct {
	key     string
	columns []string
}

type link struct {
	From string `json:"from"`
	To   string `json:"to"`
	Text string `json:"text"`
}

var ignoreField []string = []string{"deleted_at", "created_at", "updated_at", "is_deleted", "status", "sys_update_dc"}

func processStreamListToErData(list []*models.ApiBinlog, scope string, filter bool) map[string]interface{} {
	var entityList []Entity

	var links []linkData

	for _, value := range list {
		columns := extractColumns(value, scope)
		entityInfo := &Entity{}
		entityInfo.Key = value.DbName + "." + value.TableName + "-" + strconv.Itoa(value.ApiBinlogId)

		keyName := value.TableName + "_id"
		for _, fieldName := range columns {
			item := &EntityItem{Name: fieldName, IsKey: fieldName == keyName}
			entityInfo.Items = append(entityInfo.Items, *item)
		}

		entityList = append(entityList, *entityInfo)

		// 组装linkData数据
		linked := &linkData{key: entityInfo.Key}
		if filter {
			linked.columns = filterColumns(columns)
		} else {
			linked.columns = columns
		}
		links = append(links, *linked)
	}

	link := genRealLink(links)

	return map[string]interface{}{"nodeData": entityList, "linkData": link}
}

func arrayIntersect(a []string, b []string) []string {
	var newS []string
	for _, av := range a {
		for _, bv := range b {
			if av == bv {
				newS = append(newS, av)
			}
		}
	}

	return newS
}

func genRealLink(d []linkData) []link {
	var linkEntities []link
	len := len(d)
	if len < 1 {
		return linkEntities
	}

	for i := 0; i < len; i++ {
		for j := i + 1; j < len; j++ {
			remainColumns := arrayIntersect(d[i].columns, d[j].columns)
			linkRow := &link{From: d[i].key, To: d[j].key, Text: strings.Join(remainColumns, ",")}
			linkEntities = append(linkEntities, *linkRow)
		}
	}

	return linkEntities
}

func extractColumns(info *models.ApiBinlog, scope string) []string {
	switch scope {
	case "columns":
		return strings.Split(info.Columns, ",")
	case "update_columns":
		return strings.Split(info.UpdateColumns, ",")
	}

	return nil
}

func filterColumns(columns []string) []string {
	var remain []string
	for _, name := range columns {
		if inIgnoreField(name) == false {
			remain = append(remain, name)
		}
	}

	return remain
}

func inIgnoreField(name string) bool {
	for _, ignore := range ignoreField {
		if name == ignore {
			return true
		}
	}

	return false
}
