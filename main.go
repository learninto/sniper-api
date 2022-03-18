package main

import (
	"github.com/learninto/goutil/conf"
	"github.com/learninto/sniper-api/cmd/cron"
	"github.com/learninto/sniper-api/cmd/http"
	_ "github.com/learninto/sniper-api/init"

	"github.com/spf13/cobra"
)

func main() {
	root := cobra.Command{Use: conf.AppID}

	root.AddCommand(
		http.Cmd,
		cron.Cmd,
	)

	_ = root.Execute()
}
