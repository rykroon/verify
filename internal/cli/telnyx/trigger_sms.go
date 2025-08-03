package telnyx

import (
	"fmt"
	"os"

	"github.com/rykroon/verify/pkg/telnyx"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var triggerSmsCmd = &cobra.Command{
	Use:   "trigger-sms",
	Short: "Trigger SMS Verification",
	Long:  ``,
	RunE:  runTriggerSms,
}

type triggerSmsParams struct {
	phoneNumber     string
	verifyProfileId string
}

var tsp triggerSmsParams

func runTriggerSms(cmd *cobra.Command, args []string) error {
	client := telnyx.NewClient(os.Getenv("TELNYX_API_KEY"))
	params := telnyx.NewTriggerSmsVerificationParams()
	params.SetPhoneNumber(tsp.phoneNumber)
	params.SetVerifyProfileId(tsp.verifyProfileId)
	result, err := client.TriggerSmsVerification(params)
	if err != nil {
		return err
	}

	bites, err := yaml.Marshal(result)
	if err != nil {
		return nil
	}

	fmt.Println(string(bites))

	return nil
}

func init() {
	triggerSmsCmd.Flags().StringVarP(&tsp.phoneNumber, "phone-number", "p", "", "+E164 formatted phone number.")
	triggerSmsCmd.Flags().StringVarP(&tsp.verifyProfileId, "verify-profile-id", "V", "", "The identifier of the associated Verify profile.")

	triggerSmsCmd.MarkFlagRequired("phone-number")
	triggerSmsCmd.MarkFlagRequired("verify-profile-id")
}
