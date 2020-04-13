package main

import (
	"sniper-api/cmd/sniper/rename"
	"sniper-api/cmd/sniper/rpc"

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
		_ = cmd.Help()
	},
}

func main() {
	_ = Cmd.Execute()
}
