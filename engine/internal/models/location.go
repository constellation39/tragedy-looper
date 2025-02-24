package models

import "fmt"

// LocationType 代表游戏板上的位置类型
type LocationType string

const (
	LocationHospital LocationType = "Hospital" // 医院
	LocationCity     LocationType = "City"     // 城市
	LocationSchool   LocationType = "School"   // 学校
	LocationShrine   LocationType = "Shrine"   // 神社
)

// Location 代表游戏板上的一个具体位置
type Location struct {
	LocationType LocationType                 // 位置类型
	CurIntrigue  int                          // 阴谋标记数量
	Characters   map[CharacterName]*Character // 当前在此位置的角色

	// 相邻位置关系
	left     *Location // 左侧位置
	right    *Location // 右侧位置
	top      *Location // 上方位置
	bottom   *Location // 下方位置
	diagonal *Location // 斜向位置
}

func (l *Location) Intrigue() int {
	return l.CurIntrigue
}

func (l *Location) SetIntrigue(i int) {
	// 确保阴谋值不为负数
	if i < 0 {
		l.CurIntrigue = 0
		return
	}
	l.CurIntrigue = i
}

func (l *Location) Paranoia() int {
	return 0
}

func (l *Location) SetParanoia(i int) {
	return
}

func (l *Location) Goodwill() int {
	return 0
}

func (l *Location) SetGoodwill(i int) {
	// 位置不支持好感度，忽略设置
	return
}

func (l *Location) Location() LocationType {
	return l.LocationType
}

func (l *Location) ToLocation(board *Board, movementDirection MovementDirection) {
}

// getNextLocation determines the next location based on movement direction
func (l *Location) getNextLocation(movementDirection MovementDirection) (*Location, error) {
	switch movementDirection {
	case HorizontalMovement:
		return l.getHorizontalLocation()
	case VerticalMovement:
		return l.getVerticalLocation()
	case DiagonalMovement:
		return l.getDiagonalLocation()
	default:
		return nil, fmt.Errorf("invalid movement movementDirection")
	}
}

// getHorizontalLocation handles horizontal movement logic
func (l *Location) getHorizontalLocation() (*Location, error) {
	if l.left == nil && l.right == nil {
		return nil, fmt.Errorf("invalid map definition: no horizontal movement possible")
	}
	if l.left != nil && l.right != nil {
		return nil, fmt.Errorf("invalid map definition: multiple horizontal paths")
	}

	if l.left != nil {
		return l.left, nil
	}
	return l.right, nil
}

// getVerticalLocation handles vertical movement logic
func (l *Location) getVerticalLocation() (*Location, error) {
	if l.top == nil && l.bottom == nil {
		return nil, fmt.Errorf("invalid map definition: no vertical movement possible")
	}
	if l.top != nil && l.bottom != nil {
		return nil, fmt.Errorf("invalid map definition: multiple vertical paths")
	}

	if l.top != nil {
		return l.top, nil
	}
	return l.bottom, nil
}

// getDiagonalLocation handles diagonal movement logic
func (l *Location) getDiagonalLocation() (*Location, error) {
	if l.diagonal == nil {
		return nil, fmt.Errorf("invalid map definition: no diagonal movement possible")
	}
	return l.diagonal, nil
}
