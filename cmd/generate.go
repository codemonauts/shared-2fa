package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/xlzd/gotp"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "A brief description of your command",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		totp := gotp.NewDefaultTOTP("D5TDXSRUQASXCAXSM5Q5PV76KEUIV7HR")
		key, expiration := totp.NowWithExpiration()
		exp := expiration - time.Now().Unix()
		fmt.Printf("%s (Expires in %ds)\n", key, exp)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
