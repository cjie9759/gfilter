package gfilter

import (
	"log"
	"log/slog"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Filter struct {
}

var s2f = map[string]func(c any, v ...any) clause.Expression{
	// "test":   func(c any, v ...any) clause.Expression { return clause.Lt{} },
	// "gt":     func(c any, v ...any) clause.Expression { return clause.Gt{Column: c, Value: v[0]} },
	"gte":    func(c any, v ...any) clause.Expression { return clause.Gte{Column: c, Value: v[0]} },
	"in":     func(c any, v ...any) clause.Expression { return clause.IN{Column: c, Values: v} },
	"lte":    func(c any, v ...any) clause.Expression { return clause.Lte{Column: c, Value: v[0]} },
	"neq":    func(c any, v ...any) clause.Expression { return clause.Neq{Column: c, Value: v[0]} },
	"eq":     func(c any, v ...any) clause.Expression { return clause.Eq{Column: c, Value: v[0]} },
	"like":   func(c any, v ...any) clause.Expression { return clause.Like{Column: c, Value: v[0]} },
	"isnull": func(c any, v ...any) clause.Expression { return isnull{Column: c} },
}

func New(c *gin.Context, db *gorm.DB) any {
	args := c.Request.URL.Query()
	// args := map[string][]string{"eq|c": {"1"}}
	var res []any
	for k, v := range args {
		ks := strings.SplitN(k, `|`, 2)
		if len(ks) != 2 {
			log.Println(ks)
			continue
		}
		if s2f[ks[0]] != nil {
			db.Where(s2f[ks[0]](ks[1], v[0]))
		}
		// fmt.Println(k, v)
		slog.Info("", k, v)
	}
	return db.Find(&res)
}

// Gte greater than or equal to for where
type isnull clause.Eq

func (gte isnull) Build(builder clause.Builder) {
	builder.WriteQuoted(gte.Column)
	builder.WriteString(" is null ")
}
