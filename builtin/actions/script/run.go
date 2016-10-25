package script

import (
	"fmt"
	"os/exec"

	"github.com/ChrisMcKenzie/achieve/pkg/schema"
)

func runScript() *schema.Action {
	return &schema.Action{
		Exec: execScript,
	}
}

func execScript(c map[string]interface{}, meta interface{}) error {
	fmt.Println("-", c["content"].(string))
	cmd := exec.Command("bash", "-c", c["content"].(string))
	byt, err := cmd.Output()
	if err != nil {
		fmt.Println(string(byt))
		return err
	}
	return nil
}
