package typical

/* Autogenerated by Typical-Go. DO NOT EDIT.

TagName:
	@ctor

Help:
	https://pkg.go.dev/github.com/typical-go/typical-go/pkg/typapp?tab=doc#CtorAnnotation
*/

import (
	"github.com/typical-go/typical-go/pkg/typapp"
	a "github.com/typical-go/typical-rest-server/internal/app/data_access/mysqldb"
	b "github.com/typical-go/typical-rest-server/internal/app/data_access/postgresdb"
	c "github.com/typical-go/typical-rest-server/internal/app/domain/mylibrary/service"
	d "github.com/typical-go/typical-rest-server/internal/app/domain/mymusic/service"
	e "github.com/typical-go/typical-rest-server/internal/app/infra"
)

func init() {
	typapp.AppendCtor(
		&typapp.Constructor{Name: "", Fn: a.NewSongRepo},
		&typapp.Constructor{Name: "", Fn: b.NewBookRepo},
		&typapp.Constructor{Name: "", Fn: c.NewBookSvc},
		&typapp.Constructor{Name: "", Fn: d.NewSongSvc},
		&typapp.Constructor{Name: "", Fn: e.Setup},
	)
}
