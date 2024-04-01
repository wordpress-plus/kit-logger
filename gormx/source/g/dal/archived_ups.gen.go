// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/wordpress-plus/kit-common/gormx/source/g/model"
)

func newArchivedUp(db *gorm.DB, opts ...gen.DOOption) archivedUp {
	_archivedUp := archivedUp{}

	_archivedUp.archivedUpDo.UseDB(db, opts...)
	_archivedUp.archivedUpDo.UseModel(&model.ArchivedUp{})

	tableName := _archivedUp.archivedUpDo.TableName()
	_archivedUp.ALL = field.NewAsterisk(tableName)
	_archivedUp.CreateTime = field.NewTime(tableName, "create_time")
	_archivedUp.UpdateTime = field.NewTime(tableName, "update_time")
	_archivedUp.DeleteTime = field.NewTime(tableName, "delete_time")
	_archivedUp.TagID = field.NewInt64(tableName, "tag_id")
	_archivedUp.Sign = field.NewString(tableName, "sign")
	_archivedUp.Uname = field.NewString(tableName, "uname")
	_archivedUp.Mid = field.NewInt64(tableName, "mid")
	_archivedUp.Level = field.NewString(tableName, "level")
	_archivedUp.Rank = field.NewString(tableName, "rank")
	_archivedUp.Following = field.NewString(tableName, "following")
	_archivedUp.Follower = field.NewString(tableName, "follower")
	_archivedUp.View = field.NewString(tableName, "view")
	_archivedUp.Likes = field.NewString(tableName, "likes")
	_archivedUp.Resp = field.NewString(tableName, "resp")
	_archivedUp.Video = field.NewString(tableName, "video")
	_archivedUp.ArchivedUpsTag = archivedUpHasOneArchivedUpsTag{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("ArchivedUpsTag", "model.ArchivedUpsTag"),
	}

	_archivedUp.fillFieldMap()

	return _archivedUp
}

// archivedUp 关注的UP
type archivedUp struct {
	archivedUpDo

	ALL            field.Asterisk
	CreateTime     field.Time
	UpdateTime     field.Time
	DeleteTime     field.Time
	TagID          field.Int64  // my group
	Sign           field.String // up desc
	Uname          field.String // up name
	Mid            field.Int64  // up uid
	Level          field.String // up level
	Rank           field.String // up rank
	Following      field.String // up following
	Follower       field.String // up follower
	View           field.String // up view
	Likes          field.String // up likes
	Resp           field.String
	Video          field.String // up video count
	ArchivedUpsTag archivedUpHasOneArchivedUpsTag

	fieldMap map[string]field.Expr
}

func (a archivedUp) Table(newTableName string) *archivedUp {
	a.archivedUpDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a archivedUp) As(alias string) *archivedUp {
	a.archivedUpDo.DO = *(a.archivedUpDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *archivedUp) updateTableName(table string) *archivedUp {
	a.ALL = field.NewAsterisk(table)
	a.CreateTime = field.NewTime(table, "create_time")
	a.UpdateTime = field.NewTime(table, "update_time")
	a.DeleteTime = field.NewTime(table, "delete_time")
	a.TagID = field.NewInt64(table, "tag_id")
	a.Sign = field.NewString(table, "sign")
	a.Uname = field.NewString(table, "uname")
	a.Mid = field.NewInt64(table, "mid")
	a.Level = field.NewString(table, "level")
	a.Rank = field.NewString(table, "rank")
	a.Following = field.NewString(table, "following")
	a.Follower = field.NewString(table, "follower")
	a.View = field.NewString(table, "view")
	a.Likes = field.NewString(table, "likes")
	a.Resp = field.NewString(table, "resp")
	a.Video = field.NewString(table, "video")

	a.fillFieldMap()

	return a
}

func (a *archivedUp) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *archivedUp) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 16)
	a.fieldMap["create_time"] = a.CreateTime
	a.fieldMap["update_time"] = a.UpdateTime
	a.fieldMap["delete_time"] = a.DeleteTime
	a.fieldMap["tag_id"] = a.TagID
	a.fieldMap["sign"] = a.Sign
	a.fieldMap["uname"] = a.Uname
	a.fieldMap["mid"] = a.Mid
	a.fieldMap["level"] = a.Level
	a.fieldMap["rank"] = a.Rank
	a.fieldMap["following"] = a.Following
	a.fieldMap["follower"] = a.Follower
	a.fieldMap["view"] = a.View
	a.fieldMap["likes"] = a.Likes
	a.fieldMap["resp"] = a.Resp
	a.fieldMap["video"] = a.Video

}

func (a archivedUp) clone(db *gorm.DB) archivedUp {
	a.archivedUpDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a archivedUp) replaceDB(db *gorm.DB) archivedUp {
	a.archivedUpDo.ReplaceDB(db)
	return a
}

type archivedUpHasOneArchivedUpsTag struct {
	db *gorm.DB

	field.RelationField
}

func (a archivedUpHasOneArchivedUpsTag) Where(conds ...field.Expr) *archivedUpHasOneArchivedUpsTag {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a archivedUpHasOneArchivedUpsTag) WithContext(ctx context.Context) *archivedUpHasOneArchivedUpsTag {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a archivedUpHasOneArchivedUpsTag) Session(session *gorm.Session) *archivedUpHasOneArchivedUpsTag {
	a.db = a.db.Session(session)
	return &a
}

func (a archivedUpHasOneArchivedUpsTag) Model(m *model.ArchivedUp) *archivedUpHasOneArchivedUpsTagTx {
	return &archivedUpHasOneArchivedUpsTagTx{a.db.Model(m).Association(a.Name())}
}

type archivedUpHasOneArchivedUpsTagTx struct{ tx *gorm.Association }

func (a archivedUpHasOneArchivedUpsTagTx) Find() (result *model.ArchivedUpsTag, err error) {
	return result, a.tx.Find(&result)
}

func (a archivedUpHasOneArchivedUpsTagTx) Append(values ...*model.ArchivedUpsTag) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a archivedUpHasOneArchivedUpsTagTx) Replace(values ...*model.ArchivedUpsTag) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a archivedUpHasOneArchivedUpsTagTx) Delete(values ...*model.ArchivedUpsTag) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a archivedUpHasOneArchivedUpsTagTx) Clear() error {
	return a.tx.Clear()
}

func (a archivedUpHasOneArchivedUpsTagTx) Count() int64 {
	return a.tx.Count()
}

type archivedUpDo struct{ gen.DO }

type IArchivedUpDo interface {
	gen.SubQuery
	Debug() IArchivedUpDo
	WithContext(ctx context.Context) IArchivedUpDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IArchivedUpDo
	WriteDB() IArchivedUpDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IArchivedUpDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IArchivedUpDo
	Not(conds ...gen.Condition) IArchivedUpDo
	Or(conds ...gen.Condition) IArchivedUpDo
	Select(conds ...field.Expr) IArchivedUpDo
	Where(conds ...gen.Condition) IArchivedUpDo
	Order(conds ...field.Expr) IArchivedUpDo
	Distinct(cols ...field.Expr) IArchivedUpDo
	Omit(cols ...field.Expr) IArchivedUpDo
	Join(table schema.Tabler, on ...field.Expr) IArchivedUpDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IArchivedUpDo
	RightJoin(table schema.Tabler, on ...field.Expr) IArchivedUpDo
	Group(cols ...field.Expr) IArchivedUpDo
	Having(conds ...gen.Condition) IArchivedUpDo
	Limit(limit int) IArchivedUpDo
	Offset(offset int) IArchivedUpDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IArchivedUpDo
	Unscoped() IArchivedUpDo
	Create(values ...*model.ArchivedUp) error
	CreateInBatches(values []*model.ArchivedUp, batchSize int) error
	Save(values ...*model.ArchivedUp) error
	First() (*model.ArchivedUp, error)
	Take() (*model.ArchivedUp, error)
	Last() (*model.ArchivedUp, error)
	Find() ([]*model.ArchivedUp, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ArchivedUp, err error)
	FindInBatches(result *[]*model.ArchivedUp, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.ArchivedUp) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IArchivedUpDo
	Assign(attrs ...field.AssignExpr) IArchivedUpDo
	Joins(fields ...field.RelationField) IArchivedUpDo
	Preload(fields ...field.RelationField) IArchivedUpDo
	FirstOrInit() (*model.ArchivedUp, error)
	FirstOrCreate() (*model.ArchivedUp, error)
	FindByPage(offset int, limit int) (result []*model.ArchivedUp, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IArchivedUpDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a archivedUpDo) Debug() IArchivedUpDo {
	return a.withDO(a.DO.Debug())
}

func (a archivedUpDo) WithContext(ctx context.Context) IArchivedUpDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a archivedUpDo) ReadDB() IArchivedUpDo {
	return a.Clauses(dbresolver.Read)
}

func (a archivedUpDo) WriteDB() IArchivedUpDo {
	return a.Clauses(dbresolver.Write)
}

func (a archivedUpDo) Session(config *gorm.Session) IArchivedUpDo {
	return a.withDO(a.DO.Session(config))
}

func (a archivedUpDo) Clauses(conds ...clause.Expression) IArchivedUpDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a archivedUpDo) Returning(value interface{}, columns ...string) IArchivedUpDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a archivedUpDo) Not(conds ...gen.Condition) IArchivedUpDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a archivedUpDo) Or(conds ...gen.Condition) IArchivedUpDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a archivedUpDo) Select(conds ...field.Expr) IArchivedUpDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a archivedUpDo) Where(conds ...gen.Condition) IArchivedUpDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a archivedUpDo) Order(conds ...field.Expr) IArchivedUpDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a archivedUpDo) Distinct(cols ...field.Expr) IArchivedUpDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a archivedUpDo) Omit(cols ...field.Expr) IArchivedUpDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a archivedUpDo) Join(table schema.Tabler, on ...field.Expr) IArchivedUpDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a archivedUpDo) LeftJoin(table schema.Tabler, on ...field.Expr) IArchivedUpDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a archivedUpDo) RightJoin(table schema.Tabler, on ...field.Expr) IArchivedUpDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a archivedUpDo) Group(cols ...field.Expr) IArchivedUpDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a archivedUpDo) Having(conds ...gen.Condition) IArchivedUpDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a archivedUpDo) Limit(limit int) IArchivedUpDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a archivedUpDo) Offset(offset int) IArchivedUpDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a archivedUpDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IArchivedUpDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a archivedUpDo) Unscoped() IArchivedUpDo {
	return a.withDO(a.DO.Unscoped())
}

func (a archivedUpDo) Create(values ...*model.ArchivedUp) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a archivedUpDo) CreateInBatches(values []*model.ArchivedUp, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a archivedUpDo) Save(values ...*model.ArchivedUp) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a archivedUpDo) First() (*model.ArchivedUp, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ArchivedUp), nil
	}
}

func (a archivedUpDo) Take() (*model.ArchivedUp, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ArchivedUp), nil
	}
}

func (a archivedUpDo) Last() (*model.ArchivedUp, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ArchivedUp), nil
	}
}

func (a archivedUpDo) Find() ([]*model.ArchivedUp, error) {
	result, err := a.DO.Find()
	return result.([]*model.ArchivedUp), err
}

func (a archivedUpDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ArchivedUp, err error) {
	buf := make([]*model.ArchivedUp, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a archivedUpDo) FindInBatches(result *[]*model.ArchivedUp, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a archivedUpDo) Attrs(attrs ...field.AssignExpr) IArchivedUpDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a archivedUpDo) Assign(attrs ...field.AssignExpr) IArchivedUpDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a archivedUpDo) Joins(fields ...field.RelationField) IArchivedUpDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a archivedUpDo) Preload(fields ...field.RelationField) IArchivedUpDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a archivedUpDo) FirstOrInit() (*model.ArchivedUp, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ArchivedUp), nil
	}
}

func (a archivedUpDo) FirstOrCreate() (*model.ArchivedUp, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ArchivedUp), nil
	}
}

func (a archivedUpDo) FindByPage(offset int, limit int) (result []*model.ArchivedUp, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a archivedUpDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a archivedUpDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a archivedUpDo) Delete(models ...*model.ArchivedUp) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *archivedUpDo) withDO(do gen.Dao) *archivedUpDo {
	a.DO = *do.(*gen.DO)
	return a
}
