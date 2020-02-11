package tisp


type Selector struct {
	BySelectors *[]string
	ByIDs *[]string

	First int32
	Skip int32
}


var DefaultSelector = Selector{
	BySelectors: nil,
	ByIDs:       nil,
	First:       20,
	Skip:        0,
}
