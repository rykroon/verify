package sinch

import (
	"os"

	"github.com/rykroon/verify/internal/utils"
	"github.com/rykroon/verify/pkg/sinch"
	"github.com/spf13/cobra"
)

var getVerificationCmd = &cobra.Command{
	Use:   "get-verification",
	Short: "Get Verification",
	Long:  ``,
	RunE:  runGetVerification,
}

func runGetVerification(cmd *cobra.Command, args []string) error {
	client := sinch.NewClient(nil, os.Getenv("SINCH_APP_KEY"), os.Getenv("SINCH_APP_SECRET"))

	resp, err := client.GetVerificationById(rvp.id)
	if err != nil {
		return err
	}

	utils.PrintResponse(resp)
	return nil
}

func init() {
	getVerificationCmd.Flags().StringVar(&rvp.id, "id", "", "ID of the verification request")
}
