package persistence

// DataSource defines the API of the persistence layer package
type DataSource interface {
	AddItem(name string, price string) (id int, err error)
}
