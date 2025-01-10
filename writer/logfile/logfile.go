package logfile

import "log"

type Logfile struct{}

func (l Logfile) WriteCount(count int64) {
	log.Printf("Total Requests receive in last minute: %d", count)
}

func (l Logfile) WriteStatus(endpoint string, statusCode int) {
	log.Printf("Request sent successfully to %v, StatusCode: %d", endpoint, statusCode)
}
