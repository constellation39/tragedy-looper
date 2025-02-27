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
	Attributes   Attributes                   // 位置属性值
	Characters   map[CharacterName]*Character // 当前在此位置的角色

	// 相邻位置关系
	left     *Location // 左侧位置
	right    *Location // 右侧位置
	top      *Location // 上方位置
	bottom   *Location // 下方位置
	diagonal *Location // 斜向位置
}

func (l *Location) Intrigue() int {
	return l.GetAttribute(IntrigueAttribute)
}

func (l *Location) SetIntrigue(i int) {
	// 设置阴谋值
	l.SetAttribute(IntrigueAttribute, i)
}

func (l *Location) Paranoia() int {
	return l.GetAttribute(ParanoiaAttribute)
}

func (l *Location) SetParanoia(i int) {
	// 位置不支持不安值，忽略设置
	return
}

func (l *Location) Goodwill() int {
	return l.GetAttribute(GoodwillAttribute)
}

func (l *Location) SetGoodwill(i int) {
	// 位置不支持好感度，忽略设置
	return
}

func (l *Location) GetAttribute(attr AttributeType) int {
	if attr == IntrigueAttribute {
		return l.Attributes.Get(attr)
	}
	return 0 // 保持原有逻辑，location只有intrigue有意义
}

func (l *Location) SetAttribute(attr AttributeType, value int) {
	if attr == IntrigueAttribute {
		if value < 0 {
			value = 0
		}
		l.Attributes.Set(attr, value)
	}
	// 其他attr忽略，保持原SetParanoia/SetGoodwill逻辑
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
