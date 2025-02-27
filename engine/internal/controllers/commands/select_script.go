// engine/internal/controllers/commands/select_script.go
package commands

type SelectScriptCommand struct {
    ScriptName string
}

func NewSelectScriptCommand(name string) *SelectScriptCommand {
    return &SelectScriptCommand{ScriptName: name}
}

func (c *SelectScriptCommand) Type() CommandType { return CmdSelectScript }
func (c *SelectScriptCommand) Validate() error  { return nil }
func (c *SelectScriptCommand) RequiredInputs() []string { return []string{"剧本名称"} }

func (c *SelectScriptCommand) Execute(ctx CommandContext) error {
    // 剧本验证和加载逻辑将放在控制器层
    return nil
}
