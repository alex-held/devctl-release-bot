package devctl

import (
	"bytes"
	"fmt"
	"io"

	"github.com/alex-held/devctl/pkg/index/scanner"
	"github.com/alex-held/devctl/pkg/index/validate"
	"github.com/spf13/afero"
)

//ValidatePlugin validates the plugin spec
func ValidatePlugin(name, file string) error {
	plugin, err := scanner.ReadPluginFromFile(afero.NewOsFs(), file)
	if err != nil {
		return err
	}

	return validate.ValidatePlugin(name, plugin)
}

type nopCloser struct {
	io.Reader
}

func (n *nopCloser) Close() error {
	return nil
}

func NewReadCloser(spec []byte) io.ReadCloser {
	return &nopCloser{bytes.NewReader(spec)}
}

//GetPluginName gets the plugin name from template .devctl.yaml file
func GetPluginName(spec []byte) (string, error) {

	plugin, err := scanner.ReadPlugin(NewReadCloser(spec))
	if err != nil {
		return "", err
	}

	return plugin.GetName(), nil
}

//PluginFileName returns the plugin file with extension
func PluginFileName(name string) string {
	return fmt.Sprintf("%s%s", name, ".yaml")
}
