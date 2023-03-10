package pkg

import (
	"fmt"
)

func DumpTable(result SearchResult) {
	for _, row := range result.Data {
		fmt.Printf("%s %.25s %.25s %.200s\n", row.Id[7:18], row.Attributes.Timestamp, row.Attributes.Service, row.Attributes.Message)
	}
}
