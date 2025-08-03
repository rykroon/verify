package telnyx

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"

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
	result, err := client.ListVerifyProfiles(params)
	if err != nil {
		return err
	}
	bytes_, err := yaml.Marshal(result)
	if err != nil {
		return err
	}
	fmt.Println(string(bytes_))
	return nil
}

func init() {
	listProfilesCmd.Flags().IntVar(&lpp.pageNumber, "page-number", 0, "The page number")
	listProfilesCmd.Flags().IntVar(&lpp.pageSize, "page-size", 0, "The page size")
}
