package trading

import (
	"github.com/starter-go/application"
	"github.com/starter-go/trading"
)

func Module() application.Module {
	mb := trading.NewMainModule()

	// mb.Components( nil  )

	return mb.Create()
}

func ModuleForTest() application.Module {
	mb := trading.NewTestModule()

	// mb.Components( nil  )

	mb.Depend(Module())
	return mb.Create()
}
