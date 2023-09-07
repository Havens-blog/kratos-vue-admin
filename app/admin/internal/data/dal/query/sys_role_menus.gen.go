// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/data/dal/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newSysRoleMenu(db *gorm.DB, opts ...gen.DOOption) sysRoleMenu {
	_sysRoleMenu := sysRoleMenu{}

	_sysRoleMenu.sysRoleMenuDo.UseDB(db, opts...)
	_sysRoleMenu.sysRoleMenuDo.UseModel(&model.SysRoleMenu{})

	tableName := _sysRoleMenu.sysRoleMenuDo.TableName()
	_sysRoleMenu.ALL = field.NewAsterisk(tableName)
	_sysRoleMenu.ID = field.NewInt64(tableName, "id")
	_sysRoleMenu.RoleID = field.NewInt64(tableName, "role_id")
	_sysRoleMenu.MenuID = field.NewInt64(tableName, "menu_id")
	_sysRoleMenu.RoleName = field.NewString(tableName, "role_name")

	_sysRoleMenu.fillFieldMap()

	return _sysRoleMenu
}

type sysRoleMenu struct {
	sysRoleMenuDo sysRoleMenuDo

	ALL      field.Asterisk
	ID       field.Int64  // 主键id
	RoleID   field.Int64  // 角色id
	MenuID   field.Int64  // 菜单id
	RoleName field.String // 角色名称

	fieldMap map[string]field.Expr
}

func (s sysRoleMenu) Table(newTableName string) *sysRoleMenu {
	s.sysRoleMenuDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysRoleMenu) As(alias string) *sysRoleMenu {
	s.sysRoleMenuDo.DO = *(s.sysRoleMenuDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysRoleMenu) updateTableName(table string) *sysRoleMenu {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.RoleID = field.NewInt64(table, "role_id")
	s.MenuID = field.NewInt64(table, "menu_id")
	s.RoleName = field.NewString(table, "role_name")

	s.fillFieldMap()

	return s
}

func (s *sysRoleMenu) WithContext(ctx context.Context) *sysRoleMenuDo {
	return s.sysRoleMenuDo.WithContext(ctx)
}

func (s sysRoleMenu) TableName() string { return s.sysRoleMenuDo.TableName() }

func (s sysRoleMenu) Alias() string { return s.sysRoleMenuDo.Alias() }

func (s *sysRoleMenu) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysRoleMenu) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 4)
	s.fieldMap["id"] = s.ID
	s.fieldMap["role_id"] = s.RoleID
	s.fieldMap["menu_id"] = s.MenuID
	s.fieldMap["role_name"] = s.RoleName
}

func (s sysRoleMenu) clone(db *gorm.DB) sysRoleMenu {
	s.sysRoleMenuDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysRoleMenu) replaceDB(db *gorm.DB) sysRoleMenu {
	s.sysRoleMenuDo.ReplaceDB(db)
	return s
}

type sysRoleMenuDo struct{ gen.DO }

func (s sysRoleMenuDo) Debug() *sysRoleMenuDo {
	return s.withDO(s.DO.Debug())
}

func (s sysRoleMenuDo) WithContext(ctx context.Context) *sysRoleMenuDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysRoleMenuDo) ReadDB() *sysRoleMenuDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysRoleMenuDo) WriteDB() *sysRoleMenuDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysRoleMenuDo) Session(config *gorm.Session) *sysRoleMenuDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysRoleMenuDo) Clauses(conds ...clause.Expression) *sysRoleMenuDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysRoleMenuDo) Returning(value interface{}, columns ...string) *sysRoleMenuDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysRoleMenuDo) Not(conds ...gen.Condition) *sysRoleMenuDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysRoleMenuDo) Or(conds ...gen.Condition) *sysRoleMenuDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysRoleMenuDo) Select(conds ...field.Expr) *sysRoleMenuDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysRoleMenuDo) Where(conds ...gen.Condition) *sysRoleMenuDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysRoleMenuDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *sysRoleMenuDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s sysRoleMenuDo) Order(conds ...field.Expr) *sysRoleMenuDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysRoleMenuDo) Distinct(cols ...field.Expr) *sysRoleMenuDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysRoleMenuDo) Omit(cols ...field.Expr) *sysRoleMenuDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysRoleMenuDo) Join(table schema.Tabler, on ...field.Expr) *sysRoleMenuDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysRoleMenuDo) LeftJoin(table schema.Tabler, on ...field.Expr) *sysRoleMenuDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysRoleMenuDo) RightJoin(table schema.Tabler, on ...field.Expr) *sysRoleMenuDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysRoleMenuDo) Group(cols ...field.Expr) *sysRoleMenuDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysRoleMenuDo) Having(conds ...gen.Condition) *sysRoleMenuDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysRoleMenuDo) Limit(limit int) *sysRoleMenuDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysRoleMenuDo) Offset(offset int) *sysRoleMenuDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysRoleMenuDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *sysRoleMenuDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysRoleMenuDo) Unscoped() *sysRoleMenuDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysRoleMenuDo) Create(values ...*model.SysRoleMenu) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysRoleMenuDo) CreateInBatches(values []*model.SysRoleMenu, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysRoleMenuDo) Save(values ...*model.SysRoleMenu) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysRoleMenuDo) First() (*model.SysRoleMenu, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysRoleMenu), nil
	}
}

func (s sysRoleMenuDo) Take() (*model.SysRoleMenu, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysRoleMenu), nil
	}
}

func (s sysRoleMenuDo) Last() (*model.SysRoleMenu, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysRoleMenu), nil
	}
}

func (s sysRoleMenuDo) Find() ([]*model.SysRoleMenu, error) {
	result, err := s.DO.Find()
	return result.([]*model.SysRoleMenu), err
}

func (s sysRoleMenuDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysRoleMenu, err error) {
	buf := make([]*model.SysRoleMenu, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysRoleMenuDo) FindInBatches(result *[]*model.SysRoleMenu, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysRoleMenuDo) Attrs(attrs ...field.AssignExpr) *sysRoleMenuDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysRoleMenuDo) Assign(attrs ...field.AssignExpr) *sysRoleMenuDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysRoleMenuDo) Joins(fields ...field.RelationField) *sysRoleMenuDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysRoleMenuDo) Preload(fields ...field.RelationField) *sysRoleMenuDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysRoleMenuDo) FirstOrInit() (*model.SysRoleMenu, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysRoleMenu), nil
	}
}

func (s sysRoleMenuDo) FirstOrCreate() (*model.SysRoleMenu, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysRoleMenu), nil
	}
}

func (s sysRoleMenuDo) FindByPage(offset int, limit int) (result []*model.SysRoleMenu, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s sysRoleMenuDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysRoleMenuDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysRoleMenuDo) Delete(models ...*model.SysRoleMenu) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysRoleMenuDo) withDO(do gen.Dao) *sysRoleMenuDo {
	s.DO = *do.(*gen.DO)
	return s
}
