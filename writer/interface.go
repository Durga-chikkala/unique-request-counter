package writer

type CountWriter interface {
	WriteCount(count int64)
	WriteStatus(endpoint string, statusCode int)
}
