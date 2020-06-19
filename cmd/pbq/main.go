package pbq

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pbq",
	Short: "pbq is a tool to write a query to ProtocolBuffer's .proto file.",
	Long:  "pbq is a tool to write a query to ProtocolBuffer's .proto file.",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
