package telnyx

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/rykroon/verify/pkg/telnyx"
	"github.com/spf13/cobra"
)

var createProfileCmd = &cobra.Command{
	Use:   "create-profile",
	Short: "Create Verification Profile",
	Long:  ``,
	RunE:  runCreateProfiles,
}

var name string
var appName string

func init() {
	createProfileCmd.Flags().StringVarP(&name, "name", "n", "", "Name of profile")
	createProfileCmd.Flags().StringVarP(&appName, "app-name", "a", "", "Name of Application")

	createProfileCmd.MarkFlagRequired("name")
}

func runCreateProfiles(cmd *cobra.Command, args []string) error {
	params := telnyx.NewCreateVerifyProfileParams(name)
	if appName != "" {
		params.SetSmsAppName(appName)
	}

	client := telnyx.NewClient(os.Getenv("TELNYX_API_KEY"))
	result, err := client.CreateVerifyProfile(params)
	if err != nil {
		return err
	}
	bytes_, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(bytes_))
	return nil
}
