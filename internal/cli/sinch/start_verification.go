package sinch

import (
	"fmt"
	"os"

	"github.com/rykroon/verify/pkg/sinch"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func newStartVerificationCmd() *cobra.Command {
	var payload sinch.StartVerificationPayload

	cmd := &cobra.Command{
		Use:   "start-verification",
		Short: "Start Verification",
		Long:  `See API documentation: https://developers.sinch.com/docs/verification/api-reference/verification/tag/Verifications-start/#tag/Verifications-start/operation/StartVerification`,
		RunE: func(cmd *cobra.Command, args []string) error {
			client := sinch.NewClient(nil, os.Getenv("SINCH_APP_KEY"), os.Getenv("SINCH_APP_SECRET"))

			result, err := client.StartVerification(payload)
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

	cmd.Flags().StringVarP(&payload.Identity.Type, "identity-type", "t", "number", "")
	cmd.Flags().StringVarP(&payload.Identity.Endpoint, "identity-endpoint", "e", "", "E.164 formatted phone number")
	cmd.Flags().StringVarP(&payload.Method, "method", "m", "sms", "The type of the verification request")
	cmd.MarkFlagRequired("identity-endpoint")

	return cmd
}
