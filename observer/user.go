package observer

type User interface {
	update(float64)
	getUserId() int
}