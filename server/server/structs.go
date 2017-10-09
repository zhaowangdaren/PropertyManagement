package server

import (
	"../db/table"
)

type QueryEventOverview struct {
	StreetID   string
	StreetName string
	Year       int
	Quarter    int
	Month      int
	UserType   int
	PageSize   int
	PageNo     int
}

type QueryUser struct {
	UserName string
	Type     int
	PageSize int
	PageNo   int
}
type ChangePassword struct {
	Type        int
	UserName    string
	Password    string
	NewPassword string
}
type QueryBasic struct {
	Name     string
	PageNo   int
	PageSize int
}

type QueryKV struct {
	Key      string
	Value    string
	PageNo   int
	PageSize int
}

type QueryKVs struct {
	KVs      map[string]interface{}
	PageNo   int
	PageSize int
}
type QueryHouse struct {
	BuildNo  string
	PageNo   int
	PageSize int
}

type QueryGov struct {
	UserName string
	PageNo   int
	PageSize int
}

type QueryEvent struct {
	Index    string //事件编号
	PageNo   int
	PageSize int
}
type QueryParkManager struct {
	ActualName string
	PageNo     int
	PageSize   int
}

// 查询时间段内的事件
type QueryEventByTime struct {
	StartTime int64 // 查询的开始时间
	EndTime   int64 // 查询的结束时间
}

// QueryEventCheckProgress 查询OpenID所属时间段内的事件
type QueryEventCheckProgress struct {
	OpenID    string //openID
	StartTime int64  // 查询的开始时间
	EndTime   int64  // 查询的结束时间
}

// QueryPMKPI 查询物业KPI 按街道查询
type QueryPMKPI struct {
	StreetID string //街道ID
	XQID     string
	Year     int
	Quarter  int
	PageSize int
	PageNo   int
}

//ResultsPage 分页查询结果
type ResultsPage struct {
	Lists []interface{}
	Sum   int
}

type AddAnnounces struct {
	List []table.Announce `json:"list"`
}

type AddEventImg struct {
	Index string   //事件编号
	Imgs  []string //图片名称
}
type Values struct {
	Values []string
}

type LoginForm struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
