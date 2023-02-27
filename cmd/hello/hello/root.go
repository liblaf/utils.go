package hello

import (
	"fmt"

	"github.com/dsnet/try"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:  "hello",
	Args: cobra.NoArgs,

	RunE: func(cmd *cobra.Command, args []string) (err error) {
		defer try.Handle(&err)
		name := try.E1(cmd.Flags().GetString("name"))
		fmt.Printf("Hello, %s!\n", name)
		return nil
	},
}

func init() {
	RootCmd.Flags().StringP("name", "n", "world", "")
}
