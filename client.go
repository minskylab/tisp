package tisp

type ContactType int

const Email ContactType = 0
const Phone ContactType = 1

type Contact struct {
	Type  ContactType
	Value string
}

type Client struct {
	ID          ID
	CompanyName string
	Contacts    []Contact
}
