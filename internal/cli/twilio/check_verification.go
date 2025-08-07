package twilio

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/rykroon/verify/pkg/twilio"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var checkVerificationCmd = &cobra.Command{
	Use:   "check-verification",
	Short: "Check Verification",
	Long:  ``,
	RunE:  runCheckVerificationCmd,
}

type checkVerificationParams struct {
	serviceSid string
	*twilio.CheckVerificationParams
}

var cvp *checkVerificationParams

func runCheckVerificationCmd(cmd *cobra.Command, args []string) error {
	client := twilio.NewClient(nil, os.Getenv("TWILIO_API_KEY_SID"), os.Getenv("TWILIO_API_KEY_SECRET"))

	resp, err := client.CheckVerification(cvp.serviceSid, cvp.CheckVerificationParams)
	if err != nil {
		return err
	}

	fmt.Println(resp.Status)

	var m map[string]any
	if err := json.Unmarshal(resp.Body, &m); err != nil {
		return err
	}

	rawYaml, err := yaml.Marshal(m)
	if err != nil {
		return nil
	}

	fmt.Println(string(rawYaml))

	return nil
}

func init() {
	cvp = &checkVerificationParams{"", &twilio.CheckVerificationParams{}}
	checkVerificationCmd.Flags().StringVarP(&cvp.serviceSid, "service-sid", "s", "", "The SID of the verification Service.")
	checkVerificationCmd.Flags().StringVarP(&cvp.To, "to", "t", "", "The phone number or email to verify.")
	checkVerificationCmd.Flags().StringVarP(&cvp.VerificationSid, "verification-sid", "V", "", "A SID that uniquely identifies the Verification Check.")
	checkVerificationCmd.Flags().StringVarP(&cvp.Code, "code", "c", "", "The 4-10 character string being verified.")
	checkVerificationCmd.MarkFlagRequired("service-sid")
	checkVerificationCmd.MarkFlagsOneRequired("to", "verification-sid")
	checkVerificationCmd.MarkFlagRequired("code")
}
