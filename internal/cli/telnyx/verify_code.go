package telnyx

import (
	"os"

	"github.com/rykroon/verify/internal/utils"
	"github.com/rykroon/verify/pkg/telnyx"
	"github.com/spf13/cobra"
)

var verifyCodeCmd = &cobra.Command{
	Use:   "verify-code",
	Short: "Verify SMS Verification",
	Long:  ``,
	RunE:  runVerifyCode,
}

type VerifyCodeParams struct {
	VerificationId string
	telnyx.VerifyCodeParams
}

var vcp VerifyCodeParams

func runVerifyCode(cmd *cobra.Command, args []string) error {
	client := telnyx.NewClient(nil, os.Getenv("TELNYX_API_KEY"))

	resp, err := client.VerifyCode(vcp.VerificationId, vcp.VerifyCodeParams)
	if err != nil {
		return err
	}

	utils.PrintResponse(resp)
	return nil
}

func init() {
	verifyCodeCmd.Flags().StringVar(&vcp.VerificationId, "id", "", "The verification id")
	verifyCodeCmd.Flags().StringVarP(&vcp.Code, "code", "c", "", "The code")

	verifyCodeCmd.MarkFlagRequired("id")
	verifyCodeCmd.MarkFlagRequired("code")
}
