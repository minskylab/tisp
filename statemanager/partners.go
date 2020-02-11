package statemanager

import (
	"context"

	"github.com/minskylab/tisp"
	p "github.com/minskylab/tisp/repository"
)

func (db *StateManager) RegisterNewPartner(newClient tisp.NewPartnerInformation) (*p.Partner, error) {
	var contacts *p.ContactCreateManyInput

	if newClient.Email != nil {
		contacts = &p.ContactCreateManyInput{
			Create:  []p.ContactCreateInput{
				{
					Type:  p.ContactTypeEmail,
					Value: *newClient.Email,
				},
			},
		}
	}

	if newClient.Phone != nil {
		if contacts == nil {
			contacts = &p.ContactCreateManyInput{Create: []p.ContactCreateInput{}}
		}

		contacts.Create = append(contacts.Create, p.ContactCreateInput{
			Type: p.ContactTypePhone,
			Value: *newClient.Phone,
		})
	}


	return db.client.CreatePartner(p.PartnerCreateInput{
		CompanyName: newClient.CompanyName,
		Selector:    newClient.Selector,
		Contacts:    contacts,
		Projects:    nil,
	}).Exec(context.Background())
}
