package db

import (
	"database/sql"
	"fmt"
	"ops/center/cmdb/global"
	"strings"

	"github.com/go-xorm/xorm"
)

type Session struct {
	*xorm.Session
}

// 创建session
func NewSession(args ...*Session) *Session {
	if len(args) > 0 {
		return args[0]
	}
	return &Session{Session: global.DBEngine.NewSession()}
}

// 关闭session
func (p *Session) Close(args ...*Session) {
	if len(args) == 0 {
		p.Session.Close()
	}
}

// 更新表字段
func (p *Session) UpdateTable(table string, condition map[string]interface{}, set map[string]interface{}) (int64, error) {
	var (
		columns []string
		conds   []string
		args    []interface{}
	)

	for k, v := range set {
		columns = append(columns, fmt.Sprintf("%v=?", k))
		args = append(args, v)
	}

	for k, v := range condition {
		conds = append(conds, k+"=?")
		args = append(args, v)
	}
	sql := fmt.Sprintf(`UPDATE %v SET %v WHERE %v`, table, strings.Join(columns, ","), strings.Join(conds, " AND "))

	result, err := p.Exec(append([]interface{}{
		sql,
	}, args...)...)

	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (p *Session) InsertTable(table string, value map[string]interface{}) (int64, error) {
	result, err := p.insertTable(table, value)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

// 更新表字段
func (p *Session) insertTable(table string, value map[string]interface{}) (sql.Result, error) {
	var column []string
	var values []interface{}
	var param []string
	for k, v := range value {
		column = append(column, k)
		param = append(param, "?")
		values = append(values, v)
	}
	sql := fmt.Sprintf(`insert into %s(%s) values(%s)`, table, strings.Join(column, ","), strings.Join(param, ","))
	return p.Exec(sql, values)
}

// ColumnSortUnit 排序
type ColumnSortUnit struct {
	Column string `json:"column" query:"column"` // 列名
	Sort   string `json:"sort" query:"sort"`     // desc  asc
}

// SimpleSQLCondition 单表SQL查询
// map中的key转义未数据库中的key
type SimpleSQLCondition struct {
	COL  []string                 // 查询指定列名
	IN   map[string][]interface{} // in
	EQ   map[string]interface{}   // =
	LIKE map[string]interface{}   // like
	GT   map[string]interface{}   // 大于
	LT   map[string]interface{}   // 小于
	GE   map[string]interface{}   // 大于等于
	LE   map[string]interface{}   // 小于等于
	SORT []ColumnSortUnit         // 排序
}

// AddIN 添加条件
func (p *SimpleSQLCondition) AddIN(column string, value []interface{}) {

	if p.IN == nil {
		p.IN = map[string][]interface{}{
			column: value,
		}
	}
	p.IN[column] = value
}

// AddEQ 添加条件
func (p *SimpleSQLCondition) AddEQ(column string, value interface{}) {

	if p.EQ == nil {
		p.EQ = map[string]interface{}{
			column: value,
		}
	}
	p.EQ[column] = value
}

// AddLIKE 添加条件
func (p *SimpleSQLCondition) AddLIKE(column string, value interface{}) {

	if p.LIKE == nil {
		p.LIKE = map[string]interface{}{
			column: value,
		}
	}
	p.LIKE[column] = value
}

// AddGT 添加条件
func (p *SimpleSQLCondition) AddGT(column string, value interface{}) {

	if p.GT == nil {
		p.GT = map[string]interface{}{
			column: value,
		}
	}
	p.GT[column] = value
}

// AddLT 添加条件
func (p *SimpleSQLCondition) AddLT(column string, value interface{}) {

	if p.LT == nil {
		p.LT = map[string]interface{}{
			column: value,
		}
	}
	p.LT[column] = value
}

// AddGE 添加条件
func (p *SimpleSQLCondition) AddGE(column string, value interface{}) {

	if p.GE == nil {
		p.GE = map[string]interface{}{
			column: value,
		}
	}
	p.GE[column] = value
}

// AddLE 添加条件
func (p *SimpleSQLCondition) AddLE(column string, value interface{}) {

	if p.LE == nil {
		p.LE = map[string]interface{}{
			column: value,
		}
		return
	}
	p.LE[column] = value
}

// AddSORT 添加排序条件
func (p *SimpleSQLCondition) AddSORT(sorts ...ColumnSortUnit) {

	if len(sorts) == 0 {
		return
	}

	if p.SORT == nil {
		p.SORT = []ColumnSortUnit{}
	}
	p.SORT = append(p.SORT, sorts...)
}

// AddCOL 添加指定返回字段
func (p *SimpleSQLCondition) AddCOL(cols ...string) {

	if len(cols) == 0 {
		return
	}

	if p.COL == nil {
		p.COL = []string{}
		return
	}
	p.COL = append(p.COL, cols...)
}

// SearchAndCount 根绝构造方式查找和统计
func SearchAndCount(tableName string, bean interface{}, ssc *SimpleSQLCondition, args ...int64) (int64, error) {
	session := NewSession().Session
	defer session.Close()

	if tableName != "" {
		session = session.Table(tableName)
	}

	if len(ssc.COL) > 0 {
		session = session.Cols(ssc.COL...)
	}

	for k, v := range ssc.IN {
		session = session.In(k, v...)
	}

	for k, v := range ssc.EQ {
		session = session.And(fmt.Sprintf("%v = ?", k), v)
	}

	for k, v := range ssc.LIKE {
		session = session.And(fmt.Sprintf("%v like ?", k), v)
	}

	for k, v := range ssc.GT {
		session = session.And(fmt.Sprintf("%v >  ?", k), v)
	}

	for k, v := range ssc.LT {
		session = session.And(fmt.Sprintf("%v < ?", k), v)
	}

	for k, v := range ssc.GE {
		session = session.And(fmt.Sprintf("%v >=  ?", k), v)
	}

	for k, v := range ssc.LE {
		session = session.And(fmt.Sprintf("%v <= ?", k), v)
	}

	for _, s := range ssc.SORT {
		if s.Sort == "desc" {
			session = session.Desc(s.Column)
			continue
		}
		session = session.Asc(s.Column)
	}

	{
		length := len(args)
		if length > 1 {
			session = session.Limit(int(args[0]), int(args[1]))
		} else {
			session = session.Limit(int(args[0]))
		}
	}

	return session.FindAndCount(bean)
}
