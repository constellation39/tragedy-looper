package models

// Script 剧本结构
type Script struct {
	// Title 剧本标题
	Title string
	// 主要剧情
	MainPlot *Plot
	// 子剧情(Basic Tragedy Set 会有两个子剧情)
	SubPlots []*Plot
	// 角色列表
	Characters []*Character
	// 事件列表
	Incidents []Incident
	// 循环次数限制
	MaxLoops int
	// 每个循环的天数
	DaysPerLoop int
}
