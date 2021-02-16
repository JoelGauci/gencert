package cert

import (
	"encoding/csv"
	"os"
)

func ParseCSV(filename string) ([]*Cert, error) {
	certs := make([]*Cert, 0)
	f, err := os.Open(filename)
	if err != nil {
		return certs, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		return certs, err
	}

	for _, rec := range records {
		c, err := New(rec[0], rec[1], rec[2])
		if err != nil {
			return certs, err
		}
		certs = append(certs, c)
	}

	return certs, nil
}
