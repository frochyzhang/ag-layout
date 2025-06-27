package proto

import (
	"allinfinance.com/golang/cmd/aif-go/v2/internal/proto/service"
	"github.com/spf13/cobra"

	"allinfinance.com/golang/cmd/aif-go/v2/internal/proto/add"
	"allinfinance.com/golang/cmd/aif-go/v2/internal/proto/client"
	"allinfinance.com/golang/cmd/aif-go/v2/internal/proto/server"
)

// CmdProto represents the proto command.
var CmdProto = &cobra.Command{
	Use:   "proto",
	Short: "Generate the proto files",
	Long:  "Generate the proto files.",
}

func init() {
	CmdProto.AddCommand(add.CmdAdd)
	CmdProto.AddCommand(client.CmdClient)
	CmdProto.AddCommand(server.CmdServer)
	CmdProto.AddCommand(service.CmdService)
}
