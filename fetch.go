package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/popmonkey/irdata"
)

const cacheDir = ".cache"

func main() {
	args := os.Args[1:]

	if len(args) != 3 {
		fmt.Println("Usage: go run fetch.go <path to keyfile> <path to creds> <path to doc repository>")
		os.Exit(1)
	}

	keyFn, credsFn, docDir := args[0], args[1], args[2]

	api := irdata.Open(context.Background())

	defer api.Close()

	api.EnableCache(cacheDir)

	if _, err := os.Stat(credsFn); err != nil {
		irdata.SaveProvidedCredsToFile(keyFn, credsFn, irdata.CredsFromTerminal{})
	}

	err := api.AuthWithCredsFromFile(keyFn, credsFn)
	if err != nil {
		log.Panic(err)
	}

	data, err := api.GetWithCache("data/doc", time.Duration(8)*time.Hour)
	if err != nil {
		log.Panic(err)
	}

	err = os.WriteFile(filepath.Join(docDir, "doc.json"), data, 0644)
	if err != nil {
		log.Panic(err)
	}
}
