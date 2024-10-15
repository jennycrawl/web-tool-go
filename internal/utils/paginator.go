package utils

type Paginator struct {
    CurrentPage int
    PerPage     int
    Total       int64
    TotalPages  int
    Data        interface{}
}
