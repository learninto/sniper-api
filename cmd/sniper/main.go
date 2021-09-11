package main

import (
	"github.com/learninto/sniper-api/cmd/sniper/rename"
	"github.com/learninto/sniper-api/cmd/sniper/rpc"

	"github.com/spf13/cobra"
)

func init() {
	Cmd.AddCommand(rpc.Cmd)
	Cmd.AddCommand(rename.Cmd)
}

// Cmd 脚手架命令
var Cmd = &cobra.Command{
	Use:   "sniper",
	Short: "sniper 脚手架",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func main() {
	Cmd.Execute()
}
