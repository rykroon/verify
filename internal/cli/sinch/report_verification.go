package sinch

import (
	"fmt"
	"os"

	"github.com/rykroon/verify/pkg/sinch"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func newReportVerificationCmd() *cobra.Command {
	var id string
	var payload sinch.ReportVerificationPayload

	cmd := &cobra.Command{
		Use:   "report-verification",
		Short: "Report Verification",
		Long:  `See API documentation: https://developers.sinch.com/docs/verification/api-reference/verification/tag/Verifications-report/`,
		RunE: func(cmd *cobra.Command, args []string) error {
			client := sinch.NewClient(nil, os.Getenv("SINCH_APP_KEY"), os.Getenv("SINCH_APP_SECRET"))

			result, err := client.ReportVerificationById(id, payload)
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

	cmd.Flags().StringVarP(&id, "id", "i", "", "Verification ID of the verification request")
	cmd.Flags().StringVarP(&payload.Method, "method", "m", "", "The type of the verification")
	cmd.Flags().StringVarP(&payload.Code, "code", "c", "", "The code which was received by the user submitting the SMS verification")
	cmd.MarkFlagRequired("id")
	cmd.MarkFlagRequired("method")
	cmd.MarkFlagRequired("code")

	return cmd
}
