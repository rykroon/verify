package telnyx

import (
	"os"

	"github.com/rykroon/verify/internal/utils"
	"github.com/rykroon/verify/pkg/telnyx"
	"github.com/spf13/cobra"
)

func newCreateProfileCmd() *cobra.Command {
	var params telnyx.CreateVerifyProfileParams

	var cmd = &cobra.Command{
		Use:   "create-profile",
		Short: "Create Verification Profile",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			client := telnyx.NewClient(nil, os.Getenv("TELNYX_API_KEY"))
			content, err := client.CreateVerifyProfile(params)
			if err != nil {
				return err
			}

			utils.PrintContent(content)
			return nil
		},
	}

	cmd.Flags().StringVarP(&params.Name, "name", "n", "", "The name of the profile")
	// createProfileCmd.Flags().StringVarP(&cvpp.Sms.AppName, "app-name", "a", "", "The Nname of the application")
	cmd.MarkFlagRequired("name")

	return cmd
}
