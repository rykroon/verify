package telnyx

import (
	"os"

	"github.com/rykroon/verify/internal/utils"
	"github.com/rykroon/verify/pkg/telnyx"
	"github.com/spf13/cobra"
)

var updateProfileCmd = &cobra.Command{
	Use:   "update-profile",
	Short: "Update Verification Profile",
	Long:  ``,
	RunE:  runUpdateProfiles,
}

var upp telnyx.UpdateVerifyProfileParams

func runUpdateProfiles(cmd *cobra.Command, args []string) error {
	client := telnyx.NewClient(nil, os.Getenv("TELNYX_API_KEY"))
	resp, err := client.UpdateVerifyProfile(upp.VerifyProfileId, upp.UpdateVerifyProfilePayload)
	if err != nil {
		return err
	}

	utils.PrintResponse(resp)
	return nil
}

func init() {
	updateProfileCmd.Flags().StringVar(
		&upp.VerifyProfileId, "id", "", "The verify profile id.",
	)
	updateProfileCmd.Flags().StringVar(&upp.Name, "name", "", "Profile name")
	updateProfileCmd.Flags().StringVar(
		&upp.Sms.AppName, "app-name", "", "Application name",
	)
	updateProfileCmd.Flags().StringVar(
		&upp.Sms.MessagingTemplateId, "template-id", "", "Messaging template id",
	)
	updateProfileCmd.Flags().StringArrayVar(
		&upp.Sms.WhitelistedDestinations,
		"whitelisted-destinations",
		nil,
		"List of whitelisted destinations",
	)
	updateProfileCmd.Flags().IntVar(
		&upp.Sms.CodeLength, "code-length", 0, "Code length",
	)
	updateProfileCmd.MarkFlagRequired("id")
}
