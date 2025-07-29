package twilio

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/rykroon/verify/internal/twilio"
	"github.com/spf13/cobra"
)

var sendVerificationCmd = &cobra.Command{
	Use:   "send-verification",
	Short: "Send Verification",
	Long:  ``,
	RunE:  runSendVerificationCmd,
}

var serviceSid string
var to string
var channel string

func runSendVerificationCmd(cmd *cobra.Command, args []string) error {
	client := twilio.NewClient(os.Getenv("TWILIO_API_KEY_SID"), os.Getenv("TWILIO_API_KEY_SECRET"))
	params := twilio.NewSendVerificationParams(serviceSid, to, channel)
	result, err := client.SendVerification(params)
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
	sendVerificationCmd.Flags().StringVarP(&serviceSid, "service-sid", "s", "", "The SID of the verification Service.")
	sendVerificationCmd.Flags().StringVarP(&to, "to", "t", "", "The phone number or email to verify.")
	sendVerificationCmd.Flags().StringVarP(&channel, "channel", "c", "", "The verification method to use.")
	sendVerificationCmd.MarkFlagRequired("service-sid")
	sendVerificationCmd.MarkFlagRequired("to")
	sendVerificationCmd.MarkFlagRequired("channel")
}
