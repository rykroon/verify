package twilio

import (
	"fmt"
	"os"

	"github.com/rykroon/verify/pkg/twilio"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func newCheckVerificationCmd() *cobra.Command {
	var serviceSid string
	var form twilio.CheckVerificationForm

	cmd := &cobra.Command{
		Use:   "check-verification",
		Short: "Check Verification",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			client := twilio.NewClient(
				os.Getenv("TWILIO_API_KEY_SID"), os.Getenv("TWILIO_API_KEY_SECRET"), nil,
			)

			result, err := client.CheckVerification(serviceSid, form)
			if err != nil {
				return err
			}

			yamlBytes, err := yaml.Marshal(result)
			if err != nil {
				return fmt.Errorf("failed to encode as yaml: %w", err)
			}
			fmt.Println(string(yamlBytes))

			return nil
		},
	}

	cmd.Flags().StringVarP(&serviceSid, "service-sid", "s", "", "The SID of the verification Service.")
	cmd.Flags().StringVarP(&form.To, "to", "t", "", "The phone number or email to verify.")
	cmd.Flags().StringVarP(&form.VerificationSid, "verification-sid", "V", "", "A SID that uniquely identifies the Verification Check.")
	cmd.Flags().StringVarP(&form.Code, "code", "c", "", "The 4-10 character string being verified.")
	cmd.MarkFlagRequired("service-sid")
	cmd.MarkFlagsOneRequired("to", "verification-sid")
	cmd.MarkFlagRequired("code")

	return cmd
}
