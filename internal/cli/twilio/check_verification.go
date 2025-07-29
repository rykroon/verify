package twilio

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/rykroon/verify/internal/twilio"
	"github.com/spf13/cobra"
)

var checkVerificationCmd = &cobra.Command{
	Use:   "check-verification",
	Short: "Check Verification",
	Long:  ``,
	RunE:  runCheckVerificationCmd,
}

// var serviceSid string
// var to string
var verificationSid string
var code string

func runCheckVerificationCmd(cmd *cobra.Command, args []string) error {
	client := twilio.NewClient(os.Getenv("TWILIO_API_KEY_SID"), os.Getenv("TWILIO_API_KEY_SECRET"))
	params := twilio.NewCheckVerificationParams(serviceSid)
	if to != "" {
		params.SetTo(to)
	} else if verificationSid != "" {
		params.SetVerificationSid(verificationSid)
	}

	result, err := client.CheckVerification(params)
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
	checkVerificationCmd.Flags().StringVarP(&serviceSid, "service-sid", "s", "", "The SID of the verification Service.")
	checkVerificationCmd.Flags().StringVarP(&to, "to", "t", "", "The phone number or email to verify.")
	checkVerificationCmd.Flags().StringVarP(&verificationSid, "verification-sid", "V", "", "A SID that uniquely identifies the Verification Check.")
	checkVerificationCmd.Flags().StringVarP(&code, "code", "c", "", "The 4-10 character string being verified.")
	sendVerificationCmd.MarkFlagRequired("service-sid")
	sendVerificationCmd.MarkFlagsOneRequired("to", "verification-sid")
}
