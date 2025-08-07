package twilio

import (
	"os"

	"github.com/rykroon/verify/internal/utils"
	"github.com/rykroon/verify/pkg/twilio"
	"github.com/spf13/cobra"
)

var sendVerificationCmd = &cobra.Command{
	Use:   "send-verification",
	Short: "Send Verification",
	Long:  ``,
	RunE:  runSendVerificationCmd,
}

type sendVerificationParams struct {
	serviceSid string
	twilio.SendVerificationParams
}

var svp sendVerificationParams

func runSendVerificationCmd(cmd *cobra.Command, args []string) error {
	client := twilio.NewClient(nil, os.Getenv("TWILIO_API_KEY_SID"), os.Getenv("TWILIO_API_KEY_SECRET"))

	resp, err := client.SendVerification(svp.serviceSid, svp.SendVerificationParams)
	if err != nil {
		return err
	}

	utils.PrintResponse(resp)
	return nil
}

func init() {
	sendVerificationCmd.Flags().StringVar(&svp.serviceSid, "service-sid", "", "The SID of the verification Service.")
	sendVerificationCmd.Flags().StringVar(&svp.To, "to", "", "The phone number or email to verify.")
	sendVerificationCmd.Flags().StringVar(&svp.Channel, "channel", "", "The verification method to use.")
	sendVerificationCmd.MarkFlagRequired("service-sid")
	sendVerificationCmd.MarkFlagRequired("to")
	sendVerificationCmd.MarkFlagRequired("channel")
}
