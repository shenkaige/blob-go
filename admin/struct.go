package admin

import (
	"../core"
	"../db"
)

type OverviewStruct struct {
	Core    core.CoreStruct
	SubData db.OverviewDb
}
