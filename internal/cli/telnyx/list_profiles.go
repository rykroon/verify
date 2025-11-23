package telnyx

import (
	"os"

	"github.com/rykroon/verify/internal/utils"
	"github.com/rykroon/verify/pkg/telnyx"
	"github.com/spf13/cobra"
)

func newListProfilesCmd() *cobra.Command {
	var params telnyx.ListVerifyProfilesParams

	var cmd = &cobra.Command{
		Use:   "list-profiles",
		Short: "List verification profiles.",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			client := telnyx.NewClient(nil, os.Getenv("TELNYX_API_KEY"))

			content, err := client.ListVerifyProfiles(params)
			if err != nil {
				return err
			}

			utils.PrintContent(content)
			return nil
		},
	}

	cmd.Flags().IntVar(&params.PageNumber, "page-number", 0, "The page number")
	cmd.Flags().IntVar(&params.PageSize, "page-size", 0, "The page size")
	return cmd
}
