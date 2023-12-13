/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"github.com/dasalgadoc/e-Invoicing-Processor/internal/extracting"
	"github.com/dasalgadoc/e-Invoicing-Processor/internal/platform/gmail"
	"github.com/dasalgadoc/e-Invoicing-Processor/internal/scraping"
	"log"
)

func main() {
	fmt.Println("Hello world!")

	source, err := gmail.NewGmailScrapSource()
	if err != nil {
		fmt.Println(err)
	}

	srv := scraping.NewScrapSource(source)
	resp, err := srv.Invoke()
	if err != nil {
		log.Fatalf("Error invoking scrap service: %s", err)
	}

	extSrv := extracting.NewExtractService()
	_, err = extSrv.Invoke(*resp)
	if err != nil {
		log.Fatalf("Error invoking extract service: %s", err)
	}

	//cmd.Execute()
}
