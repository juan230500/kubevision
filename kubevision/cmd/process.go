/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"io"
	"regexp"
	"strings"
	"time"

	"github.com/spf13/cobra"
	corev1 "k8s.io/api/core/v1"
)

type Log struct {
	Class   string
	Time    time.Time
	Message string
}

func stringToLog(txt string) *Log {
	re := regexp.MustCompile(`\[(\w+)\]\[(.*?)\] (.*)`)
	matches := re.FindStringSubmatch(txt)

	if len(matches) != 4 {
		fmt.Println("Formato de log no reconocido")
		return nil
	}

	class := matches[1]
	timeRaw := matches[2]
	message := matches[3]

	time, err := time.Parse(time.RFC3339, timeRaw)
	if err != nil {
		fmt.Println("Fecha y hora del log no reconocidas")
		return nil
	}

	newLog := &Log{
		Class:   class,
		Time:    time,
		Message: message,
	}

	return newLog
}

// processCmd represents the process command
var processCmd = &cobra.Command{
	Use:   "process",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		for i, pod := range pods.Items {
			fmt.Printf("%d) Name: %s, Status: %s\n", i+1, pod.Name, pod.Status.Phase)
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
			logStrings = logStrings[:len(logStrings)-1]
			now := time.Now().UTC()

			counts := map[string]int{"INFO": 0, "WARNING": 0, "ERROR": 0}

			for _, log := range logStrings {
				currentLog := stringToLog(log)
				if now.Sub(currentLog.Time).Minutes() <= 10 {
					// fmt.Println(currentLog, now.Sub(currentLog.Time).Minutes())
					counts[currentLog.Class]++
				}
			}

			total := 0
			for k, v := range counts {
				fmt.Printf("%s = %d\n", k, v)
				total += v
			}
			fmt.Printf("(SUM = %d)\n", total)
			fmt.Println("==============================")
		}
	},
}

func init() {
	rootCmd.AddCommand(processCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// processCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// processCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
