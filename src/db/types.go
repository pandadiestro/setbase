package db

type Register struct {
    Name string     `json:"name"`
    Path string     `json:"path"`
    Sets []string   `json:"sets"`
}

type Query struct {
    Sets []string   `json:"sets"`
    Expr string     `json:"expr"`
}

type DB struct {
    ListedData []Register
}


