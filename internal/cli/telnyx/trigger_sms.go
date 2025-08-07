package telnyx

import (
	"os"

	"github.com/rykroon/verify/internal/utils"
	"github.com/rykroon/verify/pkg/telnyx"
	"github.com/spf13/cobra"
)

var triggerSmsCmd = &cobra.Command{
	Use:   "trigger-sms",
	Short: "Trigger SMS Verification",
	Long:  ``,
	RunE:  runTriggerSms,
}

var tsp telnyx.TriggerSmsParams

func runTriggerSms(cmd *cobra.Command, args []string) error {
	client := telnyx.NewClient(nil, os.Getenv("TELNYX_API_KEY"))

	resp, err := client.TriggerSmsVerification(tsp)
	if err != nil {
		return err
	}

	utils.PrintResponse(resp)
	return nil
}

func init() {
	triggerSmsCmd.Flags().StringVarP(&tsp.PhoneNumber, "phone-number", "p", "", "+E164 formatted phone number.")
	triggerSmsCmd.Flags().StringVarP(&tsp.VerifyProfileId, "verify-profile-id", "V", "", "The identifier of the associated Verify profile.")

	triggerSmsCmd.MarkFlagRequired("phone-number")
	triggerSmsCmd.MarkFlagRequired("verify-profile-id")
}
