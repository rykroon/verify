package telnyx

import (
	"os"

	"github.com/rykroon/verify/internal/utils"
	"github.com/rykroon/verify/pkg/telnyx"
	"github.com/spf13/cobra"
)

func newListTempaltesCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "list-templates",
		Short: "List Verification Profile Templates",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			client := telnyx.NewClient(nil, os.Getenv("TELNYX_API_KEY"))
			content, err := client.ListMessageTemplates()
			if err != nil {
				return err
			}

			utils.PrintContent(content)
			return nil
		},
	}

	return cmd
}
