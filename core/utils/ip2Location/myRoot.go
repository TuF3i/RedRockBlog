package ip2Location

type Location struct {
	ip string
}

type RootEntity struct {
	Status       string       `json:"status"`
	T            string       `json:"t"`
	SetCacheTime string       `json:"set_cache_time"`
	Data         []DataEntity `json:"data"`
}

type DataEntity struct {
	ExtendedLocation string             `json:"ExtendedLocation"`
	OriginQuery      string             `json:"OriginQuery"`
	SchemaVer        string             `json:"SchemaVer"`
	Appinfo          string             `json:"appinfo"`
	DispType         int64              `json:"disp_type"`
	Fetchkey         string             `json:"fetchkey"`
	Location         string             `json:"location"`
	Origip           string             `json:"origip"`
	Origipquery      string             `json:"origipquery"`
	Resourceid       string             `json:"resourceid"`
	RoleId           int64              `json:"role_id"`
	SchemaID         string             `json:"schemaID"`
	ShareImage       int64              `json:"shareImage"`
	ShowLikeShare    int64              `json:"showLikeShare"`
	Showlamp         string             `json:"showlamp"`
	StrategyData     StrategyDataEntity `json:"strategyData"`
	Titlecont        string             `json:"titlecont"`
	Tplt             string             `json:"tplt"`
}

type StrategyDataEntity struct {
}
