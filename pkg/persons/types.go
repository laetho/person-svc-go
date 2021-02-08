package persons



type Persons []*Person

type Person struct {
	Id      int32  `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address,omitempty"`
	Phone   int32  `json:"phone,omitempty"`
}

type CreatePerson struct {
	Name    string `json:"name"`
	Address string `json:"address,omitempty"`
	Phone   int32  `json:"phone,omitempty"`
}
