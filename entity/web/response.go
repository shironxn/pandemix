package web

type ResponseSuccess struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseError struct {
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
}

type ResponseCovid struct {
	TotalCase int    `json:"total_case"`
	Positive  int    `json:"positive"`
	Recovered int    `json:"recovered"`
	Dead      int    `json:"dead"`
	Date      string `json:"date"`
}

type ResponsePatient struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Gender   string `json:"gender"`
	Status   string `json:"status"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
}
