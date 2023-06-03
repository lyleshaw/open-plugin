// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package orm_gen

import (
	"time"
)

const TableNamePlugin = "Plugins"

// Plugin mapped from table <Plugins>
type Plugin struct {
	PluginID         int32     `gorm:"column:plugin_id;primaryKey;autoIncrement:true;comment:Unique identifier for the plugin" json:"plugin_id"`
	PluginName       string    `gorm:"column:plugin_name;not null;comment:Name of the plugin" json:"plugin_name"`
	PluginConfigURL  string    `gorm:"column:plugin_config_url;comment:URL for the plugin configuration" json:"plugin_config_url"`
	PluginOpenapiURL string    `gorm:"column:plugin_openapi_url;comment:URL for the plugin OpenAPI documentation" json:"plugin_openapi_url"`
	PluginConfig     string    `gorm:"column:plugin_config;comment:Configuration for the plugin" json:"plugin_config"`
	PluginOpenapi    string    `gorm:"column:plugin_openapi;comment:OpenAPI documentation for the plugin" json:"plugin_openapi"`
	IsDeleted        bool      `gorm:"column:is_deleted;not null;comment:Flag indicating if the plugin has been deleted" json:"is_deleted"`
	CreatedAt        time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP;comment:Timestamp for when the plugin was created" json:"created_at"`
	UpdatedAt        time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP;comment:Timestamp for when the plugin was last updated" json:"updated_at"`
}

// TableName Plugin's table name
func (*Plugin) TableName() string {
	return TableNamePlugin
}
