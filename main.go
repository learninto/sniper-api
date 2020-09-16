package main

import (
	"github.com/learninto/sniper-api/cmd/job"
	"github.com/learninto/sniper-api/cmd/server"
	_ "github.com/learninto/sniper-api/init"

	"github.com/spf13/cobra"
)

func main() {
	root := cobra.Command{Use: "sniper-api"}

	root.AddCommand(
		server.Cmd,
		job.Cmd,
	)

	_ = root.Execute()
}
