package twilio

import (
	"os"

	"github.com/rykroon/verify/internal/utils"
	"github.com/rykroon/verify/pkg/twilio"
	"github.com/spf13/cobra"
)

var checkVerificationCmd = &cobra.Command{
	Use:   "check-verification",
	Short: "Check Verification",
	Long:  ``,
	RunE:  runCheckVerificationCmd,
}

type checkVerificationParams struct {
	serviceSid string
	twilio.CheckVerificationParams
}

var cvp checkVerificationParams

func runCheckVerificationCmd(cmd *cobra.Command, args []string) error {
	client := twilio.NewClient(nil, os.Getenv("TWILIO_API_KEY_SID"), os.Getenv("TWILIO_API_KEY_SECRET"))

	resp, err := client.CheckVerification(cvp.serviceSid, cvp.CheckVerificationParams)
	if err != nil {
		return err
	}

	utils.PrintResponse(resp)

	return nil
}

func init() {
	checkVerificationCmd.Flags().StringVarP(&cvp.serviceSid, "service-sid", "s", "", "The SID of the verification Service.")
	checkVerificationCmd.Flags().StringVarP(&cvp.To, "to", "t", "", "The phone number or email to verify.")
	checkVerificationCmd.Flags().StringVarP(&cvp.VerificationSid, "verification-sid", "V", "", "A SID that uniquely identifies the Verification Check.")
	checkVerificationCmd.Flags().StringVarP(&cvp.Code, "code", "c", "", "The 4-10 character string being verified.")
	checkVerificationCmd.MarkFlagRequired("service-sid")
	checkVerificationCmd.MarkFlagsOneRequired("to", "verification-sid")
	checkVerificationCmd.MarkFlagRequired("code")
}
