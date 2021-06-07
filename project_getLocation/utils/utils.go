package utils

//定义结构体接收查询百度地图返回结果
type MapRes struct {
	Status     int
	Message    string
	ResultType string
	Results    []DetailedLocation
}

type DetailedLocation struct {
	Name     string
	Location LatAndLng
	Address  string
	Province string
	City     string
	Area     string
	Detail   int
	Uid      string
}

type LatAndLng struct {
	Lat float64
	Lng float64
}

//定义前端接收结构体
type PCALocation struct {
	Province string //省
	City     string //市
	Area     string //区（县）
}
