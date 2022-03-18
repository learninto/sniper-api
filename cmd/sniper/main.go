package main

import (
	"flag"
	"fmt"

	"github.com/learninto/sniper-api/cmd/sniper/new"
	"github.com/learninto/sniper-api/cmd/sniper/rpc"
	"github.com/learninto/sniper-api/cmd/sniper/twirp"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/compiler/protogen"
)

var version bool
var protocHelp bool

// Cmd 脚手架命令
var Cmd = &cobra.Command{
	Use:   "sniper",
	Short: "sniper 脚手架",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if version {
			fmt.Println(twirp.Version)
			return
		}

		g := twirp.NewGenerator()

		var flags flag.FlagSet

		flags.StringVar(&g.OptionPrefix, "option_prefix", "sniper", "legacy option prefix")
		flags.StringVar(&g.RootPackage, "root_package", "github.com/learninto/sniper-api", "root package of pkg")
		flags.BoolVar(&g.ValidateEnable, "validate_enable", false, "generate *.validate.go")

		if protocHelp {
			fmt.Println("protoc-gen-twirp " + twirp.Version)
			flags.PrintDefaults()
			return
		}

		protogen.Options{
			ParamFunc: flags.Set,
		}.Run(g.Generate)
	},
}

func init() {
	Cmd.Flags().BoolVar(&version, "version", false, "工具版本")
	Cmd.Flags().BoolVar(&protocHelp, "protoc-help", false, "查看 protoc-gen-twirp 帮助")
}

func main() {
	Cmd.AddCommand(rpc.Cmd)
	Cmd.AddCommand(new.Cmd)
	_ = Cmd.Execute()
}
