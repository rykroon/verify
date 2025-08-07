package telnyx

import (
	"encoding/json"
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

var tsp telnyx.TriggerSmsParams

func runTriggerSms(cmd *cobra.Command, args []string) error {
	client := telnyx.NewClient(nil, os.Getenv("TELNYX_API_KEY"))

	resp, err := client.TriggerSmsVerification(tsp)
	if err != nil {
		return err
	}

	fmt.Println(resp.Status)

	var m map[string]any
	if err := json.Unmarshal(resp.Body, &m); err != nil {
		return err
	}

	yamlBytes, err := yaml.Marshal(m)
	if err != nil {
		return err
	}

	fmt.Println(string(yamlBytes))
	return nil
}

func init() {
	triggerSmsCmd.Flags().StringVarP(&tsp.PhoneNumber, "phone-number", "p", "", "+E164 formatted phone number.")
	triggerSmsCmd.Flags().StringVarP(&tsp.VerifyProfileId, "verify-profile-id", "V", "", "The identifier of the associated Verify profile.")

	triggerSmsCmd.MarkFlagRequired("phone-number")
	triggerSmsCmd.MarkFlagRequired("verify-profile-id")
}
