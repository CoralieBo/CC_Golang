package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func createCsv(response []RepoInfos) {
	fileName := fmt.Sprintf("%s.csv", time.Now().UTC().Format("2006-01-02"))

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()

	head := []string{
		"Id",
		"Name",
		"Full name",
		"Owner login",
		"Owner id",
		"Description",
		"Url",
		"Language",
		"Updated_at",
	}
	if err := w.Write(head); err != nil {
		log.Fatal("error writing header to file", err)
	}

	for i := 0; i < len(response); i++ {
		writeCsv(response[i], w)
	}
}

func writeCsv(data RepoInfos, w *csv.Writer) {
	row := []string{
		strconv.Itoa(data.Id),
		data.Name,
		data.Full_name,
		data.Owner.Login,
		strconv.Itoa(data.Owner.Id),
		data.Description,
		data.Url,
		data.Language,
		data.Updated_at.String(),
	}
	if err := w.Write(row); err != nil {
		log.Fatal("error writing data to file", err)
	}
}
