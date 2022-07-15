package models

type Menu struct {
	MenuId    int      `gorm:"primary_key;" form:"menu_id" json:"menu_id" xorm:"not null pk autoincr comment('{"name":"文档ID","desc":"哈哈哈哈哈哈","type":"password"}') INT(10)"`
	Title     string   `form:"title" json:"title" xorm:"not null default '' comment('标题') VARCHAR(50)"`
	Pid       int      `form:"pid" json:"pid" xorm:"not null default 0 comment('{"name":"上级ID","desc":"","type":"select", "options":{"callback":"getMenuTree"}}') index INT(10)"`
	Sort      int      `form:"sort" json:"sort" xorm:"not null default 0 comment('排序（同级有效）') INT(10)"`
	Hide      int      `form:"hide" json:"hide" xorm:"not null default 0 comment('{"name":"是否隐藏","options":{"1":"否","2": "是"}}') TINYINT(1)"`
	Pathname  string   `form:"pathname" json:"pathname" xorm:"comment('路由') VARCHAR(255)"`
	Iconfont  string   `form:"iconfont" json:"iconfont" xorm:"default '' comment('{"name":"图标"}') VARCHAR(255)"`
	CreatedAt DateTime `form:"created_at" json:"created_at" xorm:"default 'CURRENT_TIMESTAMP' DATETIME"`
	UpdatedAt DateTime `form:"updated_at" json:"updated_at" xorm:"default 'CURRENT_TIMESTAMP' DATETIME"`
	IsDeleted int      `form:"id_deleted" json:"id_deleted" xorm:"default 0 TINYINT(4)"`
	DeletedAt DateTime `form:"deleted_at" json:"deleted_at" xorm:"DATETIME"`
}
