package flush

type FlushEvent interface {
	Flush(events []map[string]interface{}) error
}
