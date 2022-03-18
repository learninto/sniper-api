package main

import (
	"github.com/learninto/sniper-api/cmd/http"
	"github.com/learninto/sniper-api/cmd/job"
	_ "github.com/learninto/sniper-api/init"

	"github.com/spf13/cobra"
)

func main() {
	root := cobra.Command{Use: "sniper-api"}

	root.AddCommand(
		http.Cmd,
		job.Cmd,
	)

	_ = root.Execute()
}
