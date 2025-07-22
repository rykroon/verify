package telnyx

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/rykroon/verify/internal/telnyx"
	"github.com/spf13/cobra"
)

var triggerSmsCmd = &cobra.Command{
	Use:   "trigger-sms",
	Short: "Trigger SMS Verification",
	Long:  ``,
	RunE:  runTriggerSms,
}

var phoneNumber string
var verifyProfileId string

func runTriggerSms(cmd *cobra.Command, args []string) error {
	client := telnyx.NewClient(os.Getenv("TELNYX_API_KEY"))

	params := telnyx.NewTriggerSmsVerificationParams(phoneNumber, verifyProfileId)
	result, err := client.TriggerSmsVerification(params)
	if err != nil {
		return err
	}

	bites, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return nil
	}

	fmt.Println(string(bites))

	return nil
}

func init() {
	triggerSmsCmd.Flags().StringVarP(&phoneNumber, "phone-number", "p", "", "+E164 formatted phone number.")
	triggerSmsCmd.Flags().StringVarP(&verifyProfileId, "verify-profile-id", "v", "", "The identifier of the associated Verify profile.")

	triggerSmsCmd.MarkFlagRequired("phone-number")
	triggerSmsCmd.MarkFlagRequired("verify-profile-id")
}
