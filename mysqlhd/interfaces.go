package mysqlhd

type DBHandler interface {
	Query(sql string, param ...any) (any, error)
	Exec(sql string, param ...any) error
}
