package mysqldb

type User struct {
	ID     int64 `column:"id" json:"id"`
	Number int64 `column:"number" json:"number"`
}
