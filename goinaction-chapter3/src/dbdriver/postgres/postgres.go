package postgres

import (
    "database/sql"
    "database/sql/driver"
    "errors"
    "fmt"
)

// 自定义数据驱动。
type PostgresDriver struct{}

// 为驱动提供获取数据连接的方法，实现sql的Open接口。
func (dr PostgresDriver) Open(string) (driver.Conn, error) {
    return nil, errors.New("Unimplemented")
}

var d *PostgresDriver

// 初始化方法在main方法之前被调用。
func init() {
    fmt.Println("postgres init")

    d = new(PostgresDriver)

    // 注册自定义驱动
    sql.Register("postgres", d)
}
