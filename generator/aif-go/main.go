package main

import (
	"log"

	"github.com/spf13/cobra"

	"allinfinance.com/golang/cmd/aif-go/v2/internal/change"
	"allinfinance.com/golang/cmd/aif-go/v2/internal/project"
	"allinfinance.com/golang/cmd/aif-go/v2/internal/proto"
	"allinfinance.com/golang/cmd/aif-go/v2/internal/run"
	"allinfinance.com/golang/cmd/aif-go/v2/internal/upgrade"
)

var rootCmd = &cobra.Command{
	Use:     "aif-go",
	Short:   "aif-go: An elegant toolkit for Go microservices.",
	Long:    `aif-go: An elegant toolkit for Go microservices.`,
	Version: "1.0.0",
}

func init() {
	rootCmd.AddCommand(project.CmdNew)
	rootCmd.AddCommand(proto.CmdProto)
	rootCmd.AddCommand(upgrade.CmdUpgrade)
	rootCmd.AddCommand(change.CmdChange)
	rootCmd.AddCommand(run.CmdRun)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
