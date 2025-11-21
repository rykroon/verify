package telnyx

import (
	"os"

	"github.com/rykroon/verify/internal/utils"
	"github.com/rykroon/verify/pkg/telnyx"
	"github.com/spf13/cobra"
)

func newRetrieveProfileCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "retrieve-profile",
		Short: "Retrieve Verification Profiles",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			client := telnyx.NewClient(nil, os.Getenv("TELNYX_API_KEY"))

			resp, err := client.RetrieveVerifyProfile(verifyProfileId)
			if err != nil {
				return err
			}

			utils.PrintResponse(resp)
			return nil
		},
	}

	cmd.Flags().StringVar(&verifyProfileId, "id", "", "The verification profile id")
	cmd.MarkFlagRequired("id")
	return cmd
}
