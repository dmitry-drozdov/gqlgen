package stubgen

import (
	_ "embed"
	"path/filepath"
	"syscall"

	"github.com/dmitry-drozdov/gqlgen/codegen"
	"github.com/dmitry-drozdov/gqlgen/codegen/config"
	"github.com/dmitry-drozdov/gqlgen/codegen/templates"
	"github.com/dmitry-drozdov/gqlgen/internal/code"
	"github.com/dmitry-drozdov/gqlgen/plugin"
)

//go:embed stubs.gotpl
var stubsTemplate string

func New(filename, typename string) plugin.Plugin {
	return &Plugin{filename: filename, typeName: typename}
}

type Plugin struct {
	filename string
	typeName string
}

var (
	_ plugin.CodeGenerator = &Plugin{}
	_ plugin.ConfigMutator = &Plugin{}
)

func (m *Plugin) Name() string {
	return "stubgen"
}

func (m *Plugin) MutateConfig(cfg *config.Config) error {
	_ = syscall.Unlink(m.filename)
	return nil
}

func (m *Plugin) GenerateCode(data *codegen.Data) error {
	abs, err := filepath.Abs(m.filename)
	if err != nil {
		return err
	}
	pkgName := code.NameForDir(filepath.Dir(abs))

	return templates.Render(templates.Options{
		PackageName: pkgName,
		Filename:    m.filename,
		Data: &ResolverBuild{
			Data:     data,
			TypeName: m.typeName,
		},
		GeneratedHeader: true,
		Packages:        data.Config.Packages,
		Template:        stubsTemplate,
	})
}

type ResolverBuild struct {
	*codegen.Data

	TypeName string
}
