package sinch

import (
	"os"

	"github.com/rykroon/verify/internal/utils"
	"github.com/rykroon/verify/pkg/sinch"
	"github.com/spf13/cobra"
)

func newGetVerificationCmd() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "get-verification",
		Short: "Get Verification",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			client := sinch.NewClient(nil, os.Getenv("SINCH_APP_KEY"), os.Getenv("SINCH_APP_SECRET"))

			resp, err := client.GetVerificationById(id)
			if err != nil {
				return err
			}

			utils.PrintResponse(resp)
			return nil
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "ID of the verification request")
	return cmd
}
