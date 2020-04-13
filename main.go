package main

import (
	"sniper-api/cmd/job"
	"sniper-api/cmd/server"
	_ "sniper-api/init"

	"github.com/spf13/cobra"
)

func main() {
	root := cobra.Command{Use: "sniper-go"}

	root.AddCommand(
		server.Cmd,
		job.Cmd,
	)

	_ = root.Execute()
}
