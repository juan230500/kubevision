/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
)

func watchPodLogs(clientset *kubernetes.Clientset, pod corev1.Pod) {
	podLogOpts := corev1.PodLogOptions{
		Follow: true, // Seguir los logs para obtener actualizaciones en tiempo real
	}

	req := clientset.CoreV1().Pods(pod.Namespace).GetLogs(pod.Name, &podLogOpts)
	podLogs, err := req.Stream(context.Background())
	if err != nil {
		fmt.Printf("Error getting logs for pod %s: %s\n", pod.Name, err)
		return
	}
	defer podLogs.Close()

	buf := make([]byte, 2000) // Ajusta el tamaño del buffer según tus necesidades.
	start := time.Now()
	for {
		n, err := podLogs.Read(buf)
		if err != nil {
			fmt.Printf("Error reading logs for pod %s: %s\n", pod.Name, err)
			return
		}

		if time.Since(start).Seconds() < 0.5 {
			continue
		}

		fmt.Printf("[%s] %q\n", pod.Name, string(buf[:n]))
	}
}

// streamCmd represents the stream command
var streamCmd = &cobra.Command{
	Use:   "stream",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, pod := range pods.Items {
			go watchPodLogs(clientset, pod)
		}

		select {}
	},
}

func init() {
	rootCmd.AddCommand(streamCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// streamCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// streamCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
