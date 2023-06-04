/*
Copyright © 2023 Pierre PELOILLE <pierre@peloille.com>
*/
package cmd

import (
	"fmt"

	"github.com/pale-whale/share.me/internal/sharing"
	"github.com/pale-whale/share.me/internal/ui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/exp/constraints"
)

var g_sharePort string

// shareCmd represents the share command
var shareCmd = &cobra.Command{
	Use:   "share [flags] [file]",
	Short: "Share a single file or start the ui",
	Long: `Start a server to share a single file (or directory)
or start an UI that let you select which files to share`,
	Args: cobra.MatchAll(cobra.MaximumNArgs(1), cobra.OnlyValidArgs),
	Run:  share,
}

func init() {
	rootCmd.AddCommand(shareCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// shareCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// shareCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func displayInfos(address string, file string, id string) {
	fmt.Printf(`addr: %s/%s
file: %s
`, address, id, file)
}

func share(cmd *cobra.Command, args []string) {
	port, ok := viper.Get("port").(int)
	stringPort := ""
	if ok != true {
		stringPort = "0"
	}
	stringPort = fmt.Sprintf("%d", port)
	server := sharing.CreateServer(stringPort)
	if len(args) == 0 {
		mainWindow := ui.CreateUI(server)
		mainWindow.Run()
		return
	}

	id := server.ServeRootFile(args[0])

	displayInfos(server.Addr(), args[0], id)
	server.Serve()
}
