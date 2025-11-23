package telnyx

import (
	"fmt"
	"os"

	"github.com/rykroon/verify/pkg/telnyx"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func newTriggerSmsCmd() *cobra.Command {
	var payload telnyx.TriggerSmsPayload

	var cmd = &cobra.Command{
		Use:   "trigger-sms",
		Short: "Trigger SMS Verification",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			client := telnyx.NewClient(nil, os.Getenv("TELNYX_API_KEY"))

			result, err := client.TriggerSmsVerification(payload)
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

	cmd.Flags().StringVarP(&payload.PhoneNumber, "phone-number", "p", "", "+E164 formatted phone number.")
	cmd.Flags().StringVarP(&payload.VerifyProfileId, "verify-profile-id", "V", "", "The identifier of the associated Verify profile.")

	cmd.MarkFlagRequired("phone-number")
	cmd.MarkFlagRequired("verify-profile-id")

	return cmd

}
