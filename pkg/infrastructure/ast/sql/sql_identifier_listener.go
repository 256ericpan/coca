package sql

import (
	"fmt"
	"github.com/phodal/coca/languages/sql"
)

type SqlIdentifierListener struct {
	parser.BaseSqlListener
}

type SqlNode struct {
}

var sqlNode SqlNode

func init() {
	sqlNode = *&SqlNode{}
}

func NewSqlIdentifierListener() *SqlIdentifierListener {
	return &SqlIdentifierListener{}
}

func (s *SqlIdentifierListener) EnterSelect_core(ctx *parser.Select_coreContext) {
	columns := ctx.AllResult_column()
	for _, col := range columns {
		column := col.(*parser.Result_columnContext)
		column.GetText()
	}
	if ctx.K_FROM() != nil {
		subqueries := ctx.AllTable_or_subquery()
		for _, subquery := range subqueries {
			subqueryCtx := subquery.(*parser.Table_or_subqueryContext)
			if subqueryCtx.Table_name() != nil {
				fmt.Println(subqueryCtx.Table_name().GetText())
			}
			if subqueryCtx.Database_name() != nil {
				fmt.Println(subqueryCtx.Database_name().GetText())
			}
		}
	}
}

func (s *SqlIdentifierListener) GetNodeInfo() SqlNode {
	return sqlNode
}
