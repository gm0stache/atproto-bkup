package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/gm0stache/atproto-bkup/pkg/atproto"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "atp-bkup",
	Short: "atp-bkup provides some small utilities for the 'ATproto' protocol.",
}

var dloadCmd = &cobra.Command{
	Use:     "dload",
	Short:   "Download a ATproto repository based on a provided handle.",
	Example: "atputil dload mona.lisa.space",
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			return err
		}
		return verifyATprotoHandle(args[0])
	},
	Run: func(cmd *cobra.Command, args []string) {
		handle := args[0]
		path, err := cmd.LocalFlags().GetString("output")
		if err != nil {
			log.Fatal(err)
		}
		if err := downloadCar(cmd.Context(), handle, path); err != nil {
			log.Fatal(err)
		}
	},
	Version: "v0.1.0",
}

func verifyATprotoHandle(handle string) error {
	u, err := url.Parse(handle)
	if err != nil {
		return err
	}

	if strings.Contains(u.String(), "/") {
		return errors.New("handle must be a valid DNS name")
	}

	return nil
}

func downloadCar(ctx context.Context, handle string, path string) error {
	atID, err := atproto.GetATID(ctx, handle)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("resolved handle %q:\n", handle)
	fmt.Printf("	ATproto ID: %s\n", atID.DID.String())
	fmt.Printf("	PDS URL: %s\n", atID.PDSEndpoint())

	if err := atproto.DownloadRepo(ctx, path, atID); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("export saved at: %s\n", path)

	return nil
}

func main() {
	tstamp := time.Now().Format("2006-01-02T15-04-05")
	defaultOutputPath := fmt.Sprintf("atp-repo-%s.car", tstamp)
	dloadCmd.Flags().StringP("output", "o", defaultOutputPath, "")

	rootCmd.AddCommand(dloadCmd)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
