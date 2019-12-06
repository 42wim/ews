package ewsutil

import (
	"github.com/mhewedy/ews"
	"math"
)

// FindPeople find persona slice by query string
func FindPeople(c ews.Client, q string) ([]ews.Persona, error) {

	req := &ews.FindPeopleRequest{IndexedPageItemView: ews.IndexedPageItemView{
		MaxEntriesReturned: math.MaxInt32,
		Offset:             0,
		BasePoint:          ews.BasePointBeginning,
	}, ParentFolderId: ews.ParentFolderId{
		DistinguishedFolderId: ews.DistinguishedFolderId{Id: "directory"}},
		PersonaShape: &ews.PersonaShape{BaseShape: ews.BaseShapeIdOnly,
			AdditionalProperties: ews.AdditionalProperties{
				FieldURI: []ews.FieldURI{
					{FieldURI: "persona:DisplayName"},
					{FieldURI: "persona:Title"},
					{FieldURI: "persona:EmailAddress"},
				},
			}},
		QueryString: q,
	}

	resp, err := ews.FindPeople(c, req)

	if err != nil {
		return nil, err
	}

	return resp.People.Persona, nil
}
