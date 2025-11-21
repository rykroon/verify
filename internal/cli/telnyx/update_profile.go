package telnyx

import (
	"os"

	"github.com/rykroon/verify/internal/utils"
	"github.com/rykroon/verify/pkg/telnyx"
	"github.com/spf13/cobra"
)

func newUpdateProfileCmd() *cobra.Command {
	var params telnyx.UpdateVerifyProfileParams

	var cmd = &cobra.Command{
		Use:   "update-profile",
		Short: "Update Verification Profile",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			client := telnyx.NewClient(nil, os.Getenv("TELNYX_API_KEY"))
			resp, err := client.UpdateVerifyProfile(params.VerifyProfileId, params.UpdateVerifyProfilePayload)
			if err != nil {
				return err
			}

			utils.PrintResponse(resp)
			return nil
		},
	}

	cmd.Flags().StringVar(
		&params.VerifyProfileId, "id", "", "The verify profile id.",
	)
	cmd.Flags().StringVar(&params.Name, "name", "", "Profile name")
	cmd.Flags().StringVar(
		&params.Sms.AppName, "app-name", "", "Application name",
	)
	cmd.Flags().StringVar(
		&params.Sms.MessagingTemplateId, "template-id", "", "Messaging template id",
	)
	cmd.Flags().StringArrayVar(
		&params.Sms.WhitelistedDestinations,
		"whitelisted-destinations",
		nil,
		"List of whitelisted destinations",
	)
	cmd.Flags().IntVar(
		&params.Sms.CodeLength, "code-length", 0, "Code length",
	)
	cmd.MarkFlagRequired("id")

	return cmd
}
