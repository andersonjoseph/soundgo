package shared

type IDEncoder interface {
	Encode(id int) (string, error)
	Decode(id string) int
}
