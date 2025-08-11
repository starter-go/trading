package trading

import (
	"embed"

	"github.com/starter-go/application"
	"github.com/starter-go/starter"
)

const (
	theModuleName     = "github.com/starter-go/trading"
	theModuleVersion  = "0.0.1"
	theModuleRevision = 1
)

////////////////////////////////////////////////////////////////////////////////

const (
	theMainModuleResPath = "src/main/resources"
	theTestModuleResPath = "src/test/resources"
)

//go:embed "src/main/resources"
var theMainModuleResFS embed.FS

//go:embed "src/test/resources"
var theTestModuleResFS embed.FS

////////////////////////////////////////////////////////////////////////////////

func NewMainModule() *application.ModuleBuilder {
	mb := new(application.ModuleBuilder)

	mb.Name(theModuleName).Version(theModuleVersion).Revision(theModuleRevision)
	mb.EmbedResources(theMainModuleResFS, theMainModuleResPath)

	mb.Depend(starter.Module())

	return mb
}

func NewTestModule() *application.ModuleBuilder {
	mb := new(application.ModuleBuilder)

	mb.Name(theModuleName).Version(theModuleVersion).Revision(theModuleRevision)
	mb.EmbedResources(theTestModuleResFS, theTestModuleResPath)

	mb.Depend(starter.Module())

	return mb
}
