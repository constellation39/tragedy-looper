package models

import (
	"time"
)

// PlotType 剧情类型枚举
type PlotType string

const (
	MainPlot PlotType = "MAIN" // 主要剧情
	SubPlot  PlotType = "SUB"  // 次要剧情
)

// RuleType 规则类型枚举
type RuleType string

const (
	Mandatory RuleType = "MANDATORY" // 强制规则
	Optional  RuleType = "OPTIONAL"  // 可选规则
	Failure   RuleType = "FAILURE"   // 失败条件
)

// PlotRule 剧情规则接口
type PlotRule interface {
	CheckCondition(gameState *GameState) bool
	GetTiming() DayPhase    // 游戏阶段
	GetRuleType() RuleType  // 返回规则类型
	GetDescription() string // 规则描述
}

// Plot 剧情结构体
type Plot struct {
	ID            string         // 剧情唯一标识符
	Name          string         // 剧情名称
	Type          PlotType       // 剧情类型（主要剧情或次要剧情）
	Description   string         // 剧情描述
	Rules         []PlotRule     // 剧情规则列表
	RequiredRoles []string       // 必需角色列表
	RoleCounts    map[string]int // 每个角色需要的数量
	IsActive      bool           // 是否激活
	CreatedAt     time.Time      // 创建时间
	UpdatedAt     time.Time      // 更新时间
}

// NewPlot 创建新剧情
func NewPlot(id, name string, plotType PlotType, description string) *Plot {
	return &Plot{
		ID:            id,
		Name:          name,
		Type:          plotType,
		Description:   description,
		Rules:         make([]PlotRule, 0),
		RequiredRoles: make([]string, 0),
		RoleCounts:    make(map[string]int),
		IsActive:      false,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
}

// 添加剧情规则
func (p *Plot) AddRule(rule PlotRule) {
	p.Rules = append(p.Rules, rule)
	p.UpdatedAt = time.Now()
}

// 添加必需角色及其数量
func (p *Plot) AddRequiredRole(roleID string, count int) {
	p.RequiredRoles = append(p.RequiredRoles, roleID)
	p.RoleCounts[roleID] = count
	p.UpdatedAt = time.Now()
}
