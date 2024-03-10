package structs

import "database/sql"


type Image struct {
    Id int
    Name sql.NullString
    Path string
    Description sql.NullString
}

type Video struct {
    Id int
    Name sql.NullString
    Path string
    Description sql.NullString
}

type Item struct {
    Id int
    Name sql.NullString
    Path string
    Description sql.NullString
}

