package twilio

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/rykroon/verify/pkg/twilio"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var sendVerificationCmd = &cobra.Command{
	Use:   "send-verification",
	Short: "Send Verification",
	Long:  ``,
	RunE:  runSendVerificationCmd,
}

type sendVerificationParams struct {
	serviceSid string
	to         string
	channel    string
}

var svp sendVerificationParams

func runSendVerificationCmd(cmd *cobra.Command, args []string) error {
	client := twilio.NewClient(os.Getenv("TWILIO_API_KEY_SID"), os.Getenv("TWILIO_API_KEY_SECRET"))
	params := twilio.NewSendVerificationParams(svp.to, svp.channel)
	rawJson, err := client.SendVerification(svp.serviceSid, params)
	if err != nil {
		return err
	}

	var m map[string]any
	if err := json.Unmarshal(rawJson, &m); err != nil {
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
	sendVerificationCmd.Flags().StringVarP(&svp.serviceSid, "service-sid", "s", "", "The SID of the verification Service.")
	sendVerificationCmd.Flags().StringVarP(&svp.to, "to", "t", "", "The phone number or email to verify.")
	sendVerificationCmd.Flags().StringVarP(&svp.channel, "channel", "c", "", "The verification method to use.")
	sendVerificationCmd.MarkFlagRequired("service-sid")
	sendVerificationCmd.MarkFlagRequired("to")
	sendVerificationCmd.MarkFlagRequired("channel")
}
