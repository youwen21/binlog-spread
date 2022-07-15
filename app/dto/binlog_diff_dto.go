package dto

import (
	"binlog_spread/app/models"
	lib2 "binlog_spread/lib"
	"encoding/json"
	"fmt"
	"strings"
)

type ErSearch struct {
	Search string `json:"search" form:"search"`
	Type   string `json:"type,default=event_id" form:"type,default=event_id"`
	Scope  string `json:"scope,default=update_columns" form:"scope,default=update_columns"`
}

// 事件对比搜索条件
type BinlogDiffSearch struct {
	ApiA      string `json:"api_a" form:"api_a"`
	ApiB      string `json:"api_b" form:"api_b"`
	BinlogIds string `json:"binlog_ids" form:"binlog_ids"`
}

func (er *ErSearch) SearchIdSlice() []string {
	return strings.Split(er.Search, ",")
}

type SortStreams []*models.ApiBinlog

func (s SortStreams) Len() int { return len(s) }

func (s SortStreams) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s SortStreams) Less(i, j int) bool {
	if s[i].DbName > s[j].DbName {
		return true
	}

	if s[i].TableName > s[j].TableName {
		return true
	}

	if s[i].EventType > s[j].EventType {
		return true
	}

	return false
}

// @see https://stackoverflow.com/questions/36854408/how-to-append-to-a-slice-pointer-receiver
// @see https://www.pauladamsmith.com/blog/2016/07/go-modify-slice-iteration.html
// extractAndRemove
func (s *SortStreams) ExtractAndRemove(db string, table string) *models.ApiBinlog {
	for i, v := range *s {
		if v.DbName == db && v.TableName == table {
			// 从s中移出匹配的元素
			*s = append((*s)[:i], (*s)[i+1:]...)
			return v
		}
	}
	return s.GetDefault()
}

func (s SortStreams) GetDefault() *models.ApiBinlog {
	stream := models.ApiBinlog{}
	stream.EventType = -100
	return &stream
}

func (s SortStreams) Process() {
	// 处理Columns，UpdateColumns
	// 处理update_value
	for i, v := range s {
		if v.Columns != "" {
			v.Columns = strings.ReplaceAll(v.Columns, ",", "\n")
		}
		if v.UpdateColumns != "" {
			v.UpdateColumns = strings.ReplaceAll(v.UpdateColumns, ",", "\n")
		}
		v.UpdateValue = processValues(v.UpdateValue)

		s[i] = v
	}
}

func processValues(jsonStr string) string {
	jmap := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonStr), &jmap)
	if err != nil {
		fmt.Print(err)
		return ""
	}
	str := ""
	for k, v := range jmap {
		strV, _ := lib2.GetStringValue(v)
		str += k + ":" + strV + "\n"
	}
	return str
}
