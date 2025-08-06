package telnyx

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/rykroon/verify/internal/utils"
	"github.com/rykroon/verify/pkg/telnyx"
	"github.com/spf13/cobra"
)

var listProfilesCmd = &cobra.Command{
	Use:   "list-profiles",
	Short: "List verification profiles.",
	Long:  ``,
	RunE:  runListProfiles,
}

type listProfilesParams struct {
	pageNumber int
	pageSize   int
}

var lpp listProfilesParams

func runListProfiles(cmd *cobra.Command, args []string) error {
	client := telnyx.NewClient(os.Getenv("TELNYX_API_KEY"))
	params := telnyx.NewListVerifyProfilesParams()
	if lpp.pageNumber != 0 {
		params.SetPageNumber(lpp.pageNumber)
	}
	if lpp.pageSize != 0 {
		params.SetPageSize(lpp.pageSize)
	}
	req, err := client.NewListVerifyProfilesRequest(params)
	if err != nil {
		return err
	}

	resp, err := utils.DoAndReadAll(http.DefaultClient, req)
	if err != nil {
		return err
	}

	fmt.Println(resp.Status)

	var m map[string]any
	if err := json.Unmarshal(resp.Body, &m); err != nil {
		return err
	}
	yamlBytes, err := yaml.Marshal(m)
	if err != nil {
		return err
	}
	fmt.Println(string(yamlBytes))
	return nil
}

func init() {
	listProfilesCmd.Flags().IntVar(&lpp.pageNumber, "page-number", 0, "The page number")
	listProfilesCmd.Flags().IntVar(&lpp.pageSize, "page-size", 0, "The page size")
}
