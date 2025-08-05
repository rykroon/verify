package telnyx

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/rykroon/verify/pkg/telnyx"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var createProfileCmd = &cobra.Command{
	Use:   "create-profile",
	Short: "Create Verification Profile",
	Long:  ``,
	RunE:  runCreateProfiles,
}

type createProfileParams struct {
	name    string
	appName string
}

var cpp createProfileParams

func runCreateProfiles(cmd *cobra.Command, args []string) error {
	params := telnyx.NewCreateVerifyProfileParams()
	params.SetName(cpp.name)
	if cpp.appName != "" {
		params.SetSmsAppName(cpp.appName)
	}

	client := telnyx.NewClient(os.Getenv("TELNYX_API_KEY"))
	rawJson, err := client.CreateVerifyProfile(params)
	if err != nil {
		return err
	}

	var m map[string]any
	if err := json.Unmarshal(rawJson, &m); err != nil {
		return err
	}
	rawYaml, err := yaml.Marshal(m)
	if err != nil {
		return err
	}
	fmt.Println(string(rawYaml))
	return nil
}

func init() {
	createProfileCmd.Flags().StringVarP(&cpp.name, "name", "n", "", "The name of the profile")
	createProfileCmd.Flags().StringVarP(&cpp.appName, "app-name", "a", "", "The Nname of the application")

	createProfileCmd.MarkFlagRequired("name")
}
