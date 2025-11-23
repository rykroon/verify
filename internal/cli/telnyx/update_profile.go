package telnyx

import (
	"fmt"
	"os"

	"github.com/rykroon/verify/pkg/telnyx"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func newUpdateProfileCmd() *cobra.Command {
	var verifyProfileId string
	var payload telnyx.UpdateVerifyProfilePayload

	var cmd = &cobra.Command{
		Use:   "update-profile",
		Short: "Update Verification Profile",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			client := telnyx.NewClient(nil, os.Getenv("TELNYX_API_KEY"))
			result, err := client.UpdateVerifyProfile(verifyProfileId, payload)
			if err != nil {
				return err
			}

			yamlBytes, err := yaml.Marshal(result)
			if err != nil {
				return fmt.Errorf("failed to encode as yaml: %w", err)
			}
			fmt.Println(string(yamlBytes))
			return nil
		},
	}

	cmd.Flags().StringVar(
		&verifyProfileId, "id", "", "The verify profile id.",
	)
	cmd.Flags().StringVar(&payload.Name, "name", "", "Profile name")
	cmd.Flags().StringVar(
		&payload.Sms.AppName, "app-name", "", "Application name",
	)
	cmd.Flags().StringVar(
		&payload.Sms.MessagingTemplateId, "template-id", "", "Messaging template id",
	)
	cmd.Flags().StringArrayVar(
		&payload.Sms.WhitelistedDestinations,
		"whitelisted-destinations",
		nil,
		"List of whitelisted destinations",
	)
	cmd.Flags().IntVar(
		&payload.Sms.CodeLength, "code-length", 0, "Code length",
	)
	cmd.MarkFlagRequired("id")

	return cmd
}
