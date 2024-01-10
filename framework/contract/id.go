package contract

const IDKey = "gob:id"

type ID interface {
	NewID() string
}
