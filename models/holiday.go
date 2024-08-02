package models

type Holiday struct {
	Date        string `json:"date" xml:"date"`
	Title       string `json:"title" xml:"title"`
	Type        string `json:"type" xml:"type"`
	Inalienable bool   `json:"inalienable" xml:"inalienable"`
	Extra       string `json:"extra" xml:"extra"`
}

type HolidayResponse struct {
	Data []Holiday `json:"data"`
}
