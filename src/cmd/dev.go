package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"runtimer/utils"
	"strings"

	"github.com/spf13/cobra"
)

var devCmd = &cobra.Command{
	Use:   "dev",
	Short: "Starts a local dev server",
	Run:   dev,
}

func init() {
	devCmd.Flags().StringP("host", "H", "localhost", "host to run the dev server on")
	devCmd.Flags().IntP("port", "p", 4786, "port to run the dev server on")

	rootCmd.AddCommand(devCmd)
}

func dev(cmd *cobra.Command, args []string) {
	host, _ := cmd.Flags().GetString("host")
	port, _ := cmd.Flags().GetInt("port")

	fmt.Printf("Starting dev server on http://%v:%v\n", host, port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		dir, _ := os.ReadDir("./runtimes" + r.URL.Path)

		var files []utils.GithubFile

		stat, _ := os.Stat("./runtimes" + r.URL.Path)
		if stat.IsDir() {
			w.Header().Set("Content-Type", "application/json")
		} else {
			content, _ := os.ReadFile("./runtimes" + r.URL.Path)
			fmt.Fprint(w, string(content))
			return
		}

		for _, f := range dir {
			info, _ := f.Info()
			path := strings.Join(strings.Split(r.URL.JoinPath(info.Name()).String(), "/")[1:], "/")
			var downloadUrl string

			if info.IsDir() {
				downloadUrl = ""
			} else {
				downloadUrl = fmt.Sprintf("http://%v:%v/%v", host, port, path)
			}

			file := utils.GithubFile{
				Name:        info.Name(),
				Path:        path,
				DownloadUrl: downloadUrl,
			}

			files = append(files, file)
		}

		json.NewEncoder(w).Encode(files)
	})

	err := http.ListenAndServe(fmt.Sprintf("%v:%v", host, port), nil)
	if err != nil {
		fmt.Println(err)
	}
}
