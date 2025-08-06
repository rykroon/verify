package telnyx

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/rykroon/verify/pkg/telnyx"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var updateProfileCmd = &cobra.Command{
	Use:   "update-profile",
	Short: "Update Verification Profile",
	Long:  ``,
	RunE:  runUpdateProfiles,
}

type UpdateVerifyProfileParams struct {
	VerifyProfileId string
	*telnyx.UpdateVerifyProfileParams
}

var upp *UpdateVerifyProfileParams

func runUpdateProfiles(cmd *cobra.Command, args []string) error {
	client := telnyx.NewClient(nil, os.Getenv("TELNYX_API_KEY"))
	resp, err := client.UpdateVerifyProfile(upp.VerifyProfileId, upp.UpdateVerifyProfileParams)
	if err != nil {
		return err
	}

	fmt.Println(resp.Status)

	var m map[string]any
	if err := json.Unmarshal(resp.Body, &m); err != nil {
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
	upp = &UpdateVerifyProfileParams{"", &telnyx.UpdateVerifyProfileParams{}}
	updateProfileCmd.Flags().StringVar(&upp.VerifyProfileId, "id", "", "The id of the verification profile.")
	updateProfileCmd.Flags().StringVar(&upp.Name, "name", "", "The name of the profile")
	//updateProfileCmd.Flags().StringVar(&upp.Sms.AppName, "app-name", "", "The Name of the application")
	updateProfileCmd.MarkFlagRequired("id")
}
