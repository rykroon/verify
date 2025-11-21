package telnyx

import (
	"os"

	"github.com/rykroon/verify/internal/utils"
	"github.com/rykroon/verify/pkg/telnyx"
	"github.com/spf13/cobra"
)

func newTriggerSmsCmd() *cobra.Command {
	var params telnyx.TriggerSmsParams

	var cmd = &cobra.Command{
		Use:   "trigger-sms",
		Short: "Trigger SMS Verification",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			client := telnyx.NewClient(nil, os.Getenv("TELNYX_API_KEY"))

			resp, err := client.TriggerSmsVerification(params)
			if err != nil {
				return err
			}

			utils.PrintResponse(resp)
			return nil
		},
	}

	cmd.Flags().StringVarP(&params.PhoneNumber, "phone-number", "p", "", "+E164 formatted phone number.")
	cmd.Flags().StringVarP(&params.VerifyProfileId, "verify-profile-id", "V", "", "The identifier of the associated Verify profile.")

	cmd.MarkFlagRequired("phone-number")
	cmd.MarkFlagRequired("verify-profile-id")

	return cmd

}
