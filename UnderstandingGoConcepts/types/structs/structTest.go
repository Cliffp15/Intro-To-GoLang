package Structs

type StructTest struct {
	Test           string
	JSONString     string `json:"abc"`
	privateInteger int
}
