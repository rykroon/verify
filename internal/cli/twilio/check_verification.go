package twilio

import (
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
	serviceSid      string
	to              string
	verificationSid string
	code            string
}

var cvp checkVerificationParams

func runCheckVerificationCmd(cmd *cobra.Command, args []string) error {
	client := twilio.NewClient(os.Getenv("TWILIO_API_KEY_SID"), os.Getenv("TWILIO_API_KEY_SECRET"))
	params := twilio.NewCheckVerificationParams()
	if cvp.to != "" {
		params.SetTo(cvp.to)
	} else if cvp.verificationSid != "" {
		params.SetVerificationSid(cvp.verificationSid)
	}
	params.SetCode(cvp.code)
	result, err := client.CheckVerification(cvp.serviceSid, params)
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
	checkVerificationCmd.Flags().StringVarP(&cvp.serviceSid, "service-sid", "s", "", "The SID of the verification Service.")
	checkVerificationCmd.Flags().StringVarP(&cvp.to, "to", "t", "", "The phone number or email to verify.")
	checkVerificationCmd.Flags().StringVarP(&cvp.verificationSid, "verification-sid", "V", "", "A SID that uniquely identifies the Verification Check.")
	checkVerificationCmd.Flags().StringVarP(&cvp.code, "code", "c", "", "The 4-10 character string being verified.")
	checkVerificationCmd.MarkFlagRequired("service-sid")
	checkVerificationCmd.MarkFlagsOneRequired("to", "verification-sid")
	checkVerificationCmd.MarkFlagRequired("code")
}
