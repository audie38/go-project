package database

var connection string

func init() {
	connection = "MySQL"
}

func GetDataBase() string {
	return connection
}