package main

import (
	"fmt"
	"github.com/renevanderark/goharvest/oai"
)

func main() {
	(&oai.Request{
		BaseUrl: "http://services.kb.nl/mdo/oai", Set: "DTS", MetadataPrefix: "dcx",
		From: "2012-09-06T014:00:00.000Z",
	}).HarvestRecords(func(record *oai.Record) {
		fmt.Printf("%s\n\n", record.Metadata.Body[0:500])
	})
}
