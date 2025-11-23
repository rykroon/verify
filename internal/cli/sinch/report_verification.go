package sinch

import (
	"os"

	"github.com/rykroon/verify/internal/utils"
	"github.com/rykroon/verify/pkg/sinch"
	"github.com/spf13/cobra"
)

func newReportVerificationCmd() *cobra.Command {
	var params reportVerificationParams

	cmd := &cobra.Command{
		Use:   "report-verification",
		Short: "Report Verification",
		Long:  `See API documentation: https://developers.sinch.com/docs/verification/api-reference/verification/tag/Verifications-report/`,
		RunE: func(cmd *cobra.Command, args []string) error {
			client := sinch.NewClient(nil, os.Getenv("SINCH_APP_KEY"), os.Getenv("SINCH_APP_SECRET"))

			content, err := client.ReportVerificationById(params.id, params.ReportVerificationParams)
			if err != nil {
				return err
			}

			utils.PrintContent(content)
			return nil
		},
	}

	cmd.Flags().StringVarP(&params.id, "id", "i", "", "Verification ID of the verification request")
	cmd.Flags().StringVarP(&params.Method, "method", "m", "", "The type of the verification")
	cmd.Flags().StringVarP(&params.Code, "code", "c", "", "The code which was received by the user submitting the SMS verification")
	cmd.MarkFlagRequired("id")
	cmd.MarkFlagRequired("method")
	cmd.MarkFlagRequired("code")

	return cmd
}

type reportVerificationParams struct {
	id string
	sinch.ReportVerificationParams
}
