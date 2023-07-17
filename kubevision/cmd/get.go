/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/spf13/cobra"
	corev1 "k8s.io/api/core/v1"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		for i, pod := range pods.Items {
			fmt.Printf("%d\tName: %s, Status: %s\n", i+1, pod.Name, pod.Status.Phase)
			podLogOpts := corev1.PodLogOptions{}
			req := clientset.CoreV1().Pods(pod.Namespace).GetLogs(pod.Name, &podLogOpts)
			podLogs, err := req.Stream(context.Background())
			if err != nil {
				panic(err)
			}
			defer podLogs.Close()

			logs, err := io.ReadAll(podLogs)
			if err != nil {
				panic(err)
			}

			logStrings := strings.Split(string(logs), "\n")
			logStringsLast := logStrings[len(logStrings)-11 : len(logStrings)-1]

			for j, log := range logStringsLast {
				fmt.Printf("%d.%d\t[%s] %q\n", i+1, j+1, pod.Name, log)
			}
			fmt.Println("==============================")
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
