package devctl

import (
	"bytes"
	"fmt"

	"github.com/alex-held/devctl/pkg/index/scanner"
	"github.com/alex-held/devctl/pkg/index/validate"
)

//ValidatePlugin validates the plugin spec
func ValidatePlugin(name, file string) error {
	plugin, err := scanner.ReadPluginFile(file)
	if err != nil {
		return err
	}

	return validate.ValidatePlugin(name, plugin)
}

//GetPluginName gets the plugin name from template .devctl.yaml file
func GetPluginName(spec []byte) (string, error) {
	plugin, err := scanner.DecodePluginFile(bytes.NewReader(spec))
	if err != nil {
		return "", err
	}

	return plugin.GetName(), nil
}

//PluginFileName returns the plugin file with extension
func PluginFileName(name string) string {
	return fmt.Sprintf("%s%s", name, ".yaml")
}
