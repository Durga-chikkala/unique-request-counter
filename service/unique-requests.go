package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/Durga-chikkala/unique-request-counter/model"
	"github.com/Durga-chikkala/unique-request-counter/store"
	"github.com/Durga-chikkala/unique-request-counter/writer"
)

type Service struct {
	store.Store
	writer writer.CountWriter
}

func New(s store.Store, w writer.CountWriter) Service {
	svc := Service{Store: s, writer: w}

	return svc
}

func (s *Service) Get(ctx context.Context, f model.Filter) error {
	if s.Store.LockId(ctx, f.Id) {
		s.Store.IncrementCount(ctx)
	} else {
		return errors.New("id already processed")
	}

	if f.Endpoint != "" {
		go s.postUniqueCount(ctx, f.Endpoint)
	}

	return nil
}

// postUniqueCount sends the count to the endpoint
func (s *Service) postUniqueCount(ctx context.Context, endpoint string) {
	u, err := url.Parse(endpoint)
	if err != nil {
		log.Println("Invalid endpoint:", err)
		return
	}

	count := s.RequestCount(ctx)

	query := u.Query()
	query.Set("unique_count", fmt.Sprintf("%d", count))
	u.RawQuery = query.Encode()

	// Make HTTP POST request
	resp, err := http.Post(u.String(), "application/json", nil)
	if err != nil {
		log.Println("Error sending POST request:", err)
		return
	}

	defer resp.Body.Close()

	// Log the HTTP status code
	s.writer.WriteStatus(endpoint, resp.StatusCode)
}

// LogUniqueRequestCount logs the unique request count every minute using a time.Ticker
func (s *Service) LogUniqueRequestCount() {
	// Create a ticker that ticks every 1 minute
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop() // Stop the ticker when the function exits

	for {
		select {
		case <-ticker.C:
			uniqueCount := s.Store.RequestCount(context.Background())
			s.Store.Flush(context.Background()) // Reset unique requests for the new minute

			// Log the count of unique requests in the past minute
			s.writer.WriteCount(uniqueCount)
		}
	}
}
