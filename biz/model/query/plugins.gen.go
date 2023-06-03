// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/lyleshaw/open-plugin/biz/model/orm_gen"
)

func newPlugin(db *gorm.DB, opts ...gen.DOOption) plugin {
	_plugin := plugin{}

	_plugin.pluginDo.UseDB(db, opts...)
	_plugin.pluginDo.UseModel(&orm_gen.Plugin{})

	tableName := _plugin.pluginDo.TableName()
	_plugin.ALL = field.NewAsterisk(tableName)
	_plugin.PluginID = field.NewInt32(tableName, "plugin_id")
	_plugin.PluginName = field.NewString(tableName, "plugin_name")
	_plugin.PluginConfigURL = field.NewString(tableName, "plugin_config_url")
	_plugin.PluginOpenapiURL = field.NewString(tableName, "plugin_openapi_url")
	_plugin.PluginConfig = field.NewString(tableName, "plugin_config")
	_plugin.PluginOpenapi = field.NewString(tableName, "plugin_openapi")
	_plugin.IsDeleted = field.NewBool(tableName, "is_deleted")
	_plugin.CreatedAt = field.NewTime(tableName, "created_at")
	_plugin.UpdatedAt = field.NewTime(tableName, "updated_at")

	_plugin.fillFieldMap()

	return _plugin
}

type plugin struct {
	pluginDo

	ALL              field.Asterisk
	PluginID         field.Int32
	PluginName       field.String
	PluginConfigURL  field.String
	PluginOpenapiURL field.String
	PluginConfig     field.String
	PluginOpenapi    field.String
	IsDeleted        field.Bool
	CreatedAt        field.Time
	UpdatedAt        field.Time

	fieldMap map[string]field.Expr
}

func (p plugin) Table(newTableName string) *plugin {
	p.pluginDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p plugin) As(alias string) *plugin {
	p.pluginDo.DO = *(p.pluginDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *plugin) updateTableName(table string) *plugin {
	p.ALL = field.NewAsterisk(table)
	p.PluginID = field.NewInt32(table, "plugin_id")
	p.PluginName = field.NewString(table, "plugin_name")
	p.PluginConfigURL = field.NewString(table, "plugin_config_url")
	p.PluginOpenapiURL = field.NewString(table, "plugin_openapi_url")
	p.PluginConfig = field.NewString(table, "plugin_config")
	p.PluginOpenapi = field.NewString(table, "plugin_openapi")
	p.IsDeleted = field.NewBool(table, "is_deleted")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")

	p.fillFieldMap()

	return p
}

func (p *plugin) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *plugin) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 9)
	p.fieldMap["plugin_id"] = p.PluginID
	p.fieldMap["plugin_name"] = p.PluginName
	p.fieldMap["plugin_config_url"] = p.PluginConfigURL
	p.fieldMap["plugin_openapi_url"] = p.PluginOpenapiURL
	p.fieldMap["plugin_config"] = p.PluginConfig
	p.fieldMap["plugin_openapi"] = p.PluginOpenapi
	p.fieldMap["is_deleted"] = p.IsDeleted
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
}

func (p plugin) clone(db *gorm.DB) plugin {
	p.pluginDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p plugin) replaceDB(db *gorm.DB) plugin {
	p.pluginDo.ReplaceDB(db)
	return p
}

type pluginDo struct{ gen.DO }

type IPluginDo interface {
	gen.SubQuery
	Debug() IPluginDo
	WithContext(ctx context.Context) IPluginDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPluginDo
	WriteDB() IPluginDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPluginDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPluginDo
	Not(conds ...gen.Condition) IPluginDo
	Or(conds ...gen.Condition) IPluginDo
	Select(conds ...field.Expr) IPluginDo
	Where(conds ...gen.Condition) IPluginDo
	Order(conds ...field.Expr) IPluginDo
	Distinct(cols ...field.Expr) IPluginDo
	Omit(cols ...field.Expr) IPluginDo
	Join(table schema.Tabler, on ...field.Expr) IPluginDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPluginDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPluginDo
	Group(cols ...field.Expr) IPluginDo
	Having(conds ...gen.Condition) IPluginDo
	Limit(limit int) IPluginDo
	Offset(offset int) IPluginDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPluginDo
	Unscoped() IPluginDo
	Create(values ...*orm_gen.Plugin) error
	CreateInBatches(values []*orm_gen.Plugin, batchSize int) error
	Save(values ...*orm_gen.Plugin) error
	First() (*orm_gen.Plugin, error)
	Take() (*orm_gen.Plugin, error)
	Last() (*orm_gen.Plugin, error)
	Find() ([]*orm_gen.Plugin, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*orm_gen.Plugin, err error)
	FindInBatches(result *[]*orm_gen.Plugin, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*orm_gen.Plugin) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPluginDo
	Assign(attrs ...field.AssignExpr) IPluginDo
	Joins(fields ...field.RelationField) IPluginDo
	Preload(fields ...field.RelationField) IPluginDo
	FirstOrInit() (*orm_gen.Plugin, error)
	FirstOrCreate() (*orm_gen.Plugin, error)
	FindByPage(offset int, limit int) (result []*orm_gen.Plugin, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPluginDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p pluginDo) Debug() IPluginDo {
	return p.withDO(p.DO.Debug())
}

func (p pluginDo) WithContext(ctx context.Context) IPluginDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p pluginDo) ReadDB() IPluginDo {
	return p.Clauses(dbresolver.Read)
}

func (p pluginDo) WriteDB() IPluginDo {
	return p.Clauses(dbresolver.Write)
}

func (p pluginDo) Session(config *gorm.Session) IPluginDo {
	return p.withDO(p.DO.Session(config))
}

func (p pluginDo) Clauses(conds ...clause.Expression) IPluginDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p pluginDo) Returning(value interface{}, columns ...string) IPluginDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p pluginDo) Not(conds ...gen.Condition) IPluginDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p pluginDo) Or(conds ...gen.Condition) IPluginDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p pluginDo) Select(conds ...field.Expr) IPluginDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p pluginDo) Where(conds ...gen.Condition) IPluginDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p pluginDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IPluginDo {
	return p.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (p pluginDo) Order(conds ...field.Expr) IPluginDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p pluginDo) Distinct(cols ...field.Expr) IPluginDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p pluginDo) Omit(cols ...field.Expr) IPluginDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p pluginDo) Join(table schema.Tabler, on ...field.Expr) IPluginDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p pluginDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPluginDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p pluginDo) RightJoin(table schema.Tabler, on ...field.Expr) IPluginDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p pluginDo) Group(cols ...field.Expr) IPluginDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p pluginDo) Having(conds ...gen.Condition) IPluginDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p pluginDo) Limit(limit int) IPluginDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p pluginDo) Offset(offset int) IPluginDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p pluginDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPluginDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p pluginDo) Unscoped() IPluginDo {
	return p.withDO(p.DO.Unscoped())
}

func (p pluginDo) Create(values ...*orm_gen.Plugin) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p pluginDo) CreateInBatches(values []*orm_gen.Plugin, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p pluginDo) Save(values ...*orm_gen.Plugin) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p pluginDo) First() (*orm_gen.Plugin, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*orm_gen.Plugin), nil
	}
}

func (p pluginDo) Take() (*orm_gen.Plugin, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*orm_gen.Plugin), nil
	}
}

func (p pluginDo) Last() (*orm_gen.Plugin, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*orm_gen.Plugin), nil
	}
}

func (p pluginDo) Find() ([]*orm_gen.Plugin, error) {
	result, err := p.DO.Find()
	return result.([]*orm_gen.Plugin), err
}

func (p pluginDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*orm_gen.Plugin, err error) {
	buf := make([]*orm_gen.Plugin, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p pluginDo) FindInBatches(result *[]*orm_gen.Plugin, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p pluginDo) Attrs(attrs ...field.AssignExpr) IPluginDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p pluginDo) Assign(attrs ...field.AssignExpr) IPluginDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p pluginDo) Joins(fields ...field.RelationField) IPluginDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p pluginDo) Preload(fields ...field.RelationField) IPluginDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p pluginDo) FirstOrInit() (*orm_gen.Plugin, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*orm_gen.Plugin), nil
	}
}

func (p pluginDo) FirstOrCreate() (*orm_gen.Plugin, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*orm_gen.Plugin), nil
	}
}

func (p pluginDo) FindByPage(offset int, limit int) (result []*orm_gen.Plugin, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p pluginDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p pluginDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p pluginDo) Delete(models ...*orm_gen.Plugin) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *pluginDo) withDO(do gen.Dao) *pluginDo {
	p.DO = *do.(*gen.DO)
	return p
}