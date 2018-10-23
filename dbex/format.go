package dbex

var (
	// DataSourceFormat 可帶入參數置換產生完整的 DataSource
	DataSourceFormat = "%s:%s@tcp(%s:%s)/%s?charset=%s"
	// DataSourceWithoutDatabaseFormat 可帶入參數置換產生沒有 database 的 DataSource
	DataSourceWithoutDatabaseFormat = "%s:%s@tcp(%s:%s)/?charset=%s"
	// DatabaseURLFormat 可帶入參數置換產生 database url
	DatabaseURLFormat = "%s://%s:%s@tcp(%s:%s)/%s?charset=%s"
)
