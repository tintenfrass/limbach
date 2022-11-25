package data

var Data FullData
var AdditonalKeys = []string{"A", "B", "C", "D"}

type FullData struct {
	Individual map[int]IndividualRecord `json:"individuals"`
	Events     map[int]EventRecord      `json:"events"`
	Eventtypes []string                 `json:"eventtypes"`
	Places     []string                 `json:"places"`
	CounterI   int                      `json:"counterI"`
	CounterE   int                      `json:"counterE"`
}

type IndividualRecord struct {
	Xref  int
	GName string
	FName string
	Sex   string
}

type EventRecord struct {
	Date         string
	Details      Details //Bei Heirat
	Grandparent1 PersonRecord
	Grandparent2 PersonRecord
	Grandparent3 PersonRecord
	Grandparent4 PersonRecord
	Parent1      PersonRecord
	Parent2      PersonRecord
	Child        PersonRecord //Bei Mehrlingen müssen 2 Events angelegt werden
	Additionals  map[string]AdditionalRecord
	Married      bool
}

type AdditionalRecord struct {
	Parent PersonRecord
	Spouse PersonRecord
	Child  PersonRecord
}

type PersonRecord struct {
	Xref    int `json:",omitempty"` //Wer
	Details Details
}

type Details struct {
	Type  string `json:",omitempty"` //Was
	Place string `json:",omitempty"` //Wo
	Msg   string `json:",omitempty"` //Wie
}
