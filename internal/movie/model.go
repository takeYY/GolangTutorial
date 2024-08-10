package movie

type (
	Genre struct {
		Code string `json:"code"`
		Name string `json:"name"`
	}
	MovieResponse struct {
		ID       int32   `json:"id"`
		Title    string  `json:"title"`
		Overview string  `json:"overview"`
		Genres   []Genre `json:"genres"`
	}
)
