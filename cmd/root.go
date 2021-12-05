package cmd

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
	firstOnly bool
	url       string
	selector  string
	rootCmd   = &cobra.Command{
		Use:   "jh",
		Short: "HTML content parser with selector handling",
	}

	parseSelectorCmd = &cobra.Command{
		Use:     "parse",
		Aliases: []string{"p"},
		Short:   "parses html content",
		Run: func(cmd *cobra.Command, args []string) {
			c := http.Client{}
			resp, err := c.Get(url)
			if err != nil {
				cmd.PrintErrf("failed to fetch html: %v\n", err)
				return
			}
			defer resp.Body.Close()
			if resp.StatusCode != 200 {
				cmd.PrintErrf("status code error: %d %s\n", resp.StatusCode, resp.Status)
				return
			}
			doc, err := goquery.NewDocumentFromReader(resp.Body)
			if err != nil {
				cmd.PrintErrf("parse error: %v\n", err)
				return
			}
			if firstOnly {
				cmd.Println(doc.Find(selector).First().Text())
			} else {
				cmd.Println(doc.Find(selector).Text())
			}
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().StringVar(&url, "url", "", "url for html source")
	rootCmd.PersistentFlags().StringVar(&selector, "selector", "", "selector for parse")
	rootCmd.PersistentFlags().BoolVar(&firstOnly, "fo", false, "returns first occurance only")
	rootCmd.AddCommand(parseSelectorCmd)
}
