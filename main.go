package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"tag-to-sha/models"

	"github.com/spf13/cobra"
)

func main() {
	var logger *slog.Logger
	handler := slog.NewJSONHandler(os.Stdout, nil)
	logger = slog.New(handler)
	logger.Info("logger initialized")
	defer func() {
		logger.Info("Application complete")
	}()

	arguments, err := models.NewArguments()
	if err != nil {
		logger.Error("there is no source filename")
	}

	var rootCmd = &cobra.Command{}

	rootCmd.Flags().StringVarP(&arguments.SourceFileName, "file", "f", "", "path/name to souce file")
	rootCmd.MarkFlagRequired("file")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if _, err := os.Stat(arguments.SourceFileName); os.IsNotExist(err) {
		logger.Error("input file not found", "file", arguments.SourceFileName)
		os.Exit(1)
	} else if err != nil {
		logger.Error("error during accessing the source file", "file", arguments.SourceFileName)
		os.Exit(1)
	}

	imageFile, err := os.Open(arguments.SourceFileName)
	if err != nil {
		logger.Error("Error opening file", "file", arguments.SourceFileName, "err", err)
		return
	}
	defer imageFile.Close()

	var images []*models.Image
	scanner := bufio.NewScanner(imageFile)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		logger.Info("parsing", "line", line)

		if line == "" {
			logger.Info("empty line, skipping")
			continue
		}

		image, _ := models.NewImage()
		err := image.ParseImage(line, logger)

		if err != nil {
			logger.Error("error parsing", "line", line, "error", err)
			continue
		}

		images = append(images, image)
	}

	if err := scanner.Err(); err != nil {
		logger.Error("error reading file", "error", err)
	}

	for _, img := range images {
		logger.Info("got image", "image", img)
	}

	for index, image := range images {
		fmt.Printf("Image[%d]: Registry=%s, Path=%s, Tag=%s, Sha256digest=%s\n", index, image.Registry, image.Path, image.Tag, image.Sha256digest)
	}

}
