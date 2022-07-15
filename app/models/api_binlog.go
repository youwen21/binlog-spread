package models

type ApiBinlog struct {
	ApiBinlogId       int      `gorm:"primary_key;" json:"api_binlog_id"`
	DbName            string   `json:"db_name"`
	TableName         string   `json:"table_name"`
	TransactionTag    string   `json:"transaction_tag"`
	EventType         int      `json:"event_type"`
	Columns           string   `json:"columns"`
	UpdateColumns     string   `json:"update_columns"`
	UpdateValue       string   `json:"update_value"`
	IgnoreColumnValue string   `json:"ignore_column_value"`
	Comment           string   `json:"comment"`
	CreatedAt         DateTime `gorm:"autoUpdateTime;column:created_at;default:null" json:"created_at"`
	UpdatedAt         DateTime `gorm:"autoUpdateTime;column:updated_at;default:null" json:"updated_at"`
}

type ApiBinlog2 struct {
	ApiBinlogId       int      `gorm:"primary_key;" xorm:"not null pk autoincr INT(11)" json:"api_binlog_id"`
	DbName            string   `xorm:"VARCHAR(255)" json:"db_name"`
	TableName         string   `xorm:"VARCHAR(255)" json:"table_name"`
	TransactionTag    string   `xorm:"VARCHAR(64)" json:"transaction_tag"`
	EventType         int      `xorm:"default -100 INT(11)" json:"event_type"`
	Columns           string   `xorm:"TEXT" json:"columns"`
	UpdateColumns     string   `xorm:"comment('更新的字段') TEXT" json:"update_columns"`
	UpdateValue       string   `xorm:"comment('更新字段的值') TEXT" json:"update_value"`
	IgnoreColumnValue string   `xorm:"comment('忽略的字段值') TEXT" json:"ignore_column_value"`
	Comment           string   `xorm:"TEXT" json:"comment"`
	CreatedAt         DateTime `gorm:"autoUpdateTime;column:created_at;default:null" xorm:"default 'CURRENT_TIMESTAMP' DATETIME" json:"created_at"`
	UpdatedAt         DateTime `gorm:"autoUpdateTime;column:updated_at;default:null" xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME" json:"updated_at"`
}

type ApiBinlog3 struct {
	ApiBinlogId       int    `gorm:"primary_key;" json:"api_binlog_id"`
	DbName            string `json:"db_name"`
	TableName         string `json:"table_name"`
	TransactionTag    string `json:"transaction_tag"`
	EventType         int    `json:"event_type"`
	Columns           string `json:"columns"`
	UpdateColumns     string `json:"update_columns"`
	UpdateValue       string `json:"update_value"`
	IgnoreColumnValue string `json:"ignore_column_value"`
	Comment           string `json:"comment"`
	CreatedAt         string `gorm:"autoUpdateTime;column:created_at;default:null" xorm:"default 'CURRENT_TIMESTAMP' DATETIME" json:"created_at"`
	UpdatedAt         string `gorm:"autoUpdateTime;column:updated_at;default:null" xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME" json:"updated_at"`
}
