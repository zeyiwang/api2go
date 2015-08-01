package api2go

// The CRUD interface MUST be implemented in order to use the api2go api.
type CRUD interface {
	// FindOne returns an object by its ID
	FindOne(ID string, req Request) (interface{}, error)

	// Create a new object and return its ID
	Create(obj interface{}, req Request) (string, error)

	// Delete an object
	Delete(id string, req Request) error

	// Update an object
	Update(obj interface{}, req Request) error
}

// ContentMarshaler controls how requests from clients are unmarshaled
// and responses from the server are marshaled. The content marshaler
// is in charge of encoding and decoding data to and from a particular
// format (e.g. JSON). The encoding and decoding processes follow the
// rules of the standard encoding/json package.
type ContentMarshaler interface {
	Marshal(i interface{}) ([]byte, error)
	Unmarshal(data []byte, i interface{}) error
}

// The PaginatedFindAll interface can be optionally implemented to fetch a subset of all records.
// Pagination query parameters must be used to limit the result. Pagination URLs will automatically
// be generated by the api. You can use a combination of the following 2 query parameters:
// page[number] AND page[size]
// OR page[offset] AND page[limit]
type PaginatedFindAll interface {
	PaginatedFindAll(req Request) (obj interface{}, totalCount uint, err error)
}

// The FindAll interface can be optionally implemented to fetch all records at once.
type FindAll interface {
	// FindAll returns all objects
	FindAll(req Request) (interface{}, error)
}
