package db

type Register struct {
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


