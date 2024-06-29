package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/arnaldomf/dns/domain/dns"
	"github.com/spf13/cobra"
)

// decodeCmd represents the decode command
var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "Decodes a DNS message and outputs its contents",
	Long: `Decodes a DNS message from its binary format to
a UTF-8 text, allowing the user to understand the contents
of such a message.`,
	Run: func(cmd *cobra.Command, args []string) {
		fromStdin, _ := cmd.Flags().GetBool("stdin")
		filePath, _ := cmd.Flags().GetString("path")
		file := os.Stdin

		if fromStdin && len(filePath) > 0 {
			fmt.Print("you cannot used --path and --stdin at the same time")
			return
		}

		var err error
		if len(filePath) > 0 {
			file, err = os.Open(filePath)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			defer file.Close()
		}

		content, err := io.ReadAll(file)
		if err != nil && err != io.EOF {
			fmt.Println(err.Error())
			return
		}

		d, err := dns.New(content)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println(d)
	},
}

func init() {
	rootCmd.AddCommand(decodeCmd)
	decodeCmd.Flags().StringP("path", "p", "", "Path of the file to be decoded")
	decodeCmd.Flags().Bool("stdin", false, "Read message from stdin")
}
