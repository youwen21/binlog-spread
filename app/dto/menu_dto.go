package dto

import (
	"binlog_spread/app/models"
	"strings"
)

type MenuForm struct {
	models.Menu
	PageParam
}

type MenuResult struct {
	Count int64         `json:"count"`
	List  []models.Menu `json:"list"`
}

type MenuText struct {
	models.Menu
}

// 左侧菜单树
type MenuList []*models.Menu

type MenuTreeItem struct {
	models.Menu
	Child []MenuTreeItem `json:"_child"`
}

// 左侧菜单树 -end

// 关于子结构体转父结构体的说明
//https://stackoverflow.com/questions/46989247/how-to-convert-parent-type-to-child-type
func (m MenuTreeItem) ConvertToMenu() models.Menu {
	menu := models.Menu{}

	menu.MenuId = m.MenuId
	menu.Title = m.Title
	menu.Pid = m.Pid
	menu.Sort = m.Sort
	menu.Hide = m.Hide
	menu.Pathname = m.Pathname
	menu.Iconfont = m.Iconfont
	menu.CreatedAt = m.CreatedAt
	menu.UpdatedAt = m.UpdatedAt
	menu.IsDeleted = m.IsDeleted
	menu.DeletedAt = m.DeletedAt

	return menu
}

// 菜单选择列表
type MenuSelectItem struct {
	models.Menu
	Level     int    `json:"level"`
	TitleShow string `json:"title_show"`
}

func (m *MenuList) ProcessToTree(pid int, level int) []MenuTreeItem {
	var mTree []MenuTreeItem
	if level == 10 {
		return mTree
	}

	list := m.findChildren(pid)
	if len(list) == 0 {
		return mTree
	}

	for _, v := range list {
		child := m.ProcessToTree(v.MenuId, level+1)
		mTree = append(mTree, MenuTreeItem{*v, child})
	}

	return mTree
}

func (m *MenuList) findChildren(pid int) []*models.Menu {
	var child []*models.Menu

	for _, v := range *m {
		if v.Pid == pid {
			child = append(child, v)
		}
	}
	return child
}

// -- 菜单选择列表 --
func TreeToSelect(msP *[]MenuSelectItem, mt []MenuTreeItem, level int) {
	for _, v := range mt {
		prefixStr := strings.Repeat("&nbsp;", level*2)
		prefixStr += "└"
		showTitle := prefixStr + v.Menu.Title
		item := MenuSelectItem{v.ConvertToMenu(), level, showTitle}
		*msP = append(*msP, item)
		if len(v.Child) > 0 {
			TreeToSelect(msP, v.Child, level+1)
		}
	}
}
