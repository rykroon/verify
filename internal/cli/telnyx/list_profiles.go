package telnyx

import (
	"os"

	"github.com/rykroon/verify/internal/utils"
	"github.com/rykroon/verify/pkg/telnyx"
	"github.com/spf13/cobra"
)

var listProfilesCmd = &cobra.Command{
	Use:   "list-profiles",
	Short: "List verification profiles.",
	Long:  ``,
	RunE:  runListProfiles,
}

var lpp telnyx.ListVerifyProfilesParams

func runListProfiles(cmd *cobra.Command, args []string) error {
	client := telnyx.NewClient(nil, os.Getenv("TELNYX_API_KEY"))

	resp, err := client.ListVerifyProfiles(lpp)
	if err != nil {
		return err
	}

	utils.PrintResponse(resp)
	return nil
}

func init() {
	listProfilesCmd.Flags().IntVar(&lpp.PageNumber, "page-number", 0, "The page number")
	listProfilesCmd.Flags().IntVar(&lpp.PageSize, "page-size", 0, "The page size")
}
