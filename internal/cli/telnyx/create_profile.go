package telnyx

import (
	"os"

	"github.com/rykroon/verify/internal/utils"
	"github.com/rykroon/verify/pkg/telnyx"
	"github.com/spf13/cobra"
)

var createProfileCmd = &cobra.Command{
	Use:   "create-profile",
	Short: "Create Verification Profile",
	Long:  ``,
	RunE:  runCreateProfiles,
}

var cvpp telnyx.CreateVerifyProfileParams

func runCreateProfiles(cmd *cobra.Command, args []string) error {
	client := telnyx.NewClient(nil, os.Getenv("TELNYX_API_KEY"))
	resp, err := client.CreateVerifyProfile(cvpp)
	if err != nil {
		return err
	}

	utils.PrintResponse(resp)
	return nil
}

func init() {
	createProfileCmd.Flags().StringVarP(&cvpp.Name, "name", "n", "", "The name of the profile")
	// createProfileCmd.Flags().StringVarP(&cvpp.Sms.AppName, "app-name", "a", "", "The Nname of the application")
	createProfileCmd.MarkFlagRequired("name")
}
