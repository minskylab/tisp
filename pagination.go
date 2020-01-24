package tisp


type Pagination struct {
	BySelectors *[]string
	ByIDs *[]string

	First int32
	Skip int32
}


var DefaultPagination = Pagination{
	BySelectors: nil,
	ByIDs:       nil,
	First:       20,
	Skip:        0,
}
