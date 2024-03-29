// Command pinger is used to start the app-server, central-server or the
// agent to execute checks. It is used as the entry point for all components
// of the application.
package main

import (
	"os"

	"github.com/sdslabs/pinger/cmd"
	"github.com/sdslabs/pinger/pkg/util/appcontext"

	// Initialize all the plugins here.
	// Standard plugins included in the `pkg/plugins` package.
	_ "github.com/sdslabs/pinger/pkg/plugins"
)

func main() {
	// Parent context for the application.
	ctx, cancel := appcontext.WithSignals(
		cmd.NewAppContext(),
		os.Interrupt, os.Kill, // Exit on interrupt or kill
	)
	defer cancel()

	command, err := cmd.NewRootCmd(ctx)
	if err != nil {
		ctx.Logger().
			WithError(err).
			Fatalln("cannot create pinger command")
		return
	}

	err = command.Execute()
	if err != nil {
		ctx.Logger().
			WithError(err).
			Fatalln("cannot execute pinger command")
		return
	}
}
