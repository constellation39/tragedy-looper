package commands

import (
	"fmt"
	"tragedy-looper/engine/internal/models"
)

// SelectCharCommand 选择角色作为目标命令
type SelectCharCommand struct {
	CharacterName string
}

// NewSelectCharCommand 创建新的选择角色命令
func NewSelectCharCommand(characterName string) *SelectCharCommand {
	return &SelectCharCommand{
		CharacterName: characterName,
	}
}

func (c *SelectCharCommand) Type() CommandType {
	return CmdSelectChar
}

// Validate 验证命令参数
func (c *SelectCharCommand) Validate() error {
	if c.CharacterName == "" {
		return fmt.Errorf("角色名称不能为空")
	}
	return nil
}

// RequiredInputs 返回命令需要的输入描述
func (c *SelectCharCommand) RequiredInputs() []string {
	return []string{
		"角色名称: 要选择的角色名",
	}
}

func (c *SelectCharCommand) Execute(ctx CommandContext) error {
	if err := c.Validate(); err != nil {
		return err
	}

	// 检查角色是否存在
	character := ctx.GameState.Character(models.CharacterName(c.CharacterName))
	if character == nil {
		return fmt.Errorf("找不到角色: %s", c.CharacterName)
	}

	// 单独执行该命令时，只是显示角色信息
	fmt.Printf("已选择角色: %s\n", c.CharacterName)
	fmt.Printf("当前位置: %s\n", character.Location())

	return nil
}

// SelectLocationCommand 选择位置作为目标命令
type SelectLocationCommand struct {
	LocationType string
}

// NewSelectLocationCommand 创建新的选择位置命令
func NewSelectLocationCommand(locationType string) *SelectLocationCommand {
	return &SelectLocationCommand{
		LocationType: locationType,
	}
}

func (c *SelectLocationCommand) Type() CommandType {
	return CmdSelectLocation
}

// Validate 验证命令参数
func (c *SelectLocationCommand) Validate() error {
	if c.LocationType == "" {
		return fmt.Errorf("位置类型不能为空")
	}
	return nil
}

// RequiredInputs 返回命令需要的输入描述
func (c *SelectLocationCommand) RequiredInputs() []string {
	return []string{
		"位置类型: 要选择的位置类型",
	}
}

func (c *SelectLocationCommand) Execute(ctx CommandContext) error {
	if err := c.Validate(); err != nil {
		return err
	}

	// 检查位置是否存在
	location := ctx.GameState.Location(models.LocationType(c.LocationType))
	if location == nil {
		return fmt.Errorf("找不到位置: %s", c.LocationType)
	}

	// 单独执行该命令时，只是显示位置信息
	fmt.Printf("已选择位置: %s\n", c.LocationType)
	fmt.Printf("阴谋标记: %d\n", location.CurIntrigue)

	// 显示该位置上的角色
	if len(location.Characters) > 0 {
		fmt.Print("角色: ")
		for name := range location.Characters {
			fmt.Printf("%s ", name)
		}
		fmt.Println()
	}

	return nil
}
