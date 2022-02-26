package starterrestful

import (
	"embed"

	"github.com/bitwormhole/starter"
	"github.com/bitwormhole/starter-restful/gen"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
)

const (
	myModuleName = "github.com/bitwormhole/starter-restful"
	myModuleVer  = "v0.1.5"
	myModuleRev  = 6
)

//go:embed "src/main/resources"
var myModuleMainRes embed.FS

// Module 导出模块【"github.com/bitwormhole/starter-restful"】
func Module() application.Module {

	mb := application.ModuleBuilder{}
	mb.Name(myModuleName).Version(myModuleVer).Revision(myModuleRev)

	mb.Resources(collection.LoadEmbedResources(&myModuleMainRes, "src/main/resources"))
	mb.OnMount(gen.ExportConfigForStarterRESTFul)
	mb.Dependency(starter.Module())

	return mb.Create()
}
