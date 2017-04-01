package achieve

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/hcl/ast"
)

// LoadConfig reads the file at the path given and parses it and returns a new
// Config
func LoadConfig(root string) (*Config, error) {
	data, err := readFile(root)
	if err != nil {
		return nil, err
	}

	f, err := hcl.Parse(data)
	if err != nil {
		return nil, err
	}

	conf, err := parse(f)

	return conf, err
}

func parse(f *ast.File) (*Config, error) {
	// Top-level item should be the object list
	list, ok := f.Node.(*ast.ObjectList)
	if !ok {
		return nil, fmt.Errorf("error parsing: file does not contain root node object")
	}

	conf := new(Config)

	if variables := list.Filter("variable"); len(variables.Items) > 0 {
		var err error
		conf.Variables, err = loadVariables(variables)
		if err != nil {
			return nil, err
		}
	}

	if tasks := list.Filter("task"); len(tasks.Items) > 0 {
		var err error
		conf.Tasks, err = loadTasks(tasks)
		if err != nil {
			return nil, err
		}
	}

	if providers := list.Filter("provider"); len(providers.Items) > 0 {
		var err error
		conf.ProviderConfigs, err = loadProviders(providers)
		if err != nil {
			return nil, err
		}
	}

	return conf, nil
}

func loadTasks(list *ast.ObjectList) (map[string]*Task, error) {
	list = list.Children()
	if len(list.Items) == 0 {
		return nil, nil
	}

	result := make(map[string]*Task)

	for _, item := range list.Items {
		name := item.Keys[0].Token.Value().(string)

		var listVal *ast.ObjectList
		if ot, ok := item.Val.(*ast.ObjectType); ok {
			listVal = ot.List
		} else {
			return nil, fmt.Errorf("task '%s': should be an object", name)
		}

		var actions []*Action
		if o := listVal.Filter("action"); len(o.Items) > 0 {
			var err error
			actions, err = loadActions(o)
			if err != nil {
				return nil, fmt.Errorf(
					"Error parsing response for %s: %s",
					name,
					err)
			}
		}

		result[name] = &Task{
			Actions: actions,
		}
	}

	return result, nil
}

func loadActions(list *ast.ObjectList) ([]*Action, error) {
	list = list.Children()
	if len(list.Items) == 0 {
		return nil, nil
	}

	var result []*Action

	for _, item := range list.Items {
		action := item.Keys[0].Token.Value().(string)
		provider := strings.Split(action, "_")[0]

		var config map[string]interface{}
		if err := hcl.DecodeObject(&config, item.Val); err != nil {
			return nil, fmt.Errorf(
				"Error reading config for task %s: %s",
				action,
				err)
		}

		result = append(result, &Action{
			Provider:   provider,
			Name:       action,
			RawOptions: config,
		})
	}

	return result, nil
}

func loadProviders(list *ast.ObjectList) ([]*ProviderConfig, error) {
	list = list.Children()
	if len(list.Items) == 0 {
		return nil, nil
	}

	var result []*ProviderConfig

	for _, item := range list.Items {
		name := item.Keys[0].Token.Value().(string)

		var config map[string]interface{}
		if err := hcl.DecodeObject(&config, item.Val); err != nil {
			return nil, fmt.Errorf(
				"Error reading provider %s: %s",
				name,
				err)
		}

		result = append(result, &ProviderConfig{
			Name:      name,
			RawConfig: config,
		})
	}

	return result, nil
}

func loadVariables(list *ast.ObjectList) ([]*Variable, error) {
	list = list.Children()
	if len(list.Items) == 0 {
		return nil, nil
	}

	var result []*Variable

	for _, item := range list.Items {
		name := item.Keys[0].Token.Value().(string)

		if _, ok := item.Val.(*ast.ObjectType); !ok {
			return nil, fmt.Errorf("variable '%s': should be an object", name)
		}

		result = append(result, &Variable{
			Name: name,
		})
	}

	return result, nil
}

func readFile(file string) (string, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
