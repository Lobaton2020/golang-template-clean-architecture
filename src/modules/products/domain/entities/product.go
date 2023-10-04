package products

type Product struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Brand     string  `json:"brand"`
	CreatedAt string  `json:"created_at"`
}
