
package main

import (
    "database/sql"

    // 如果只是利用包的初始化而不真正使用它，则在导入时使用下划线作为别名
    _ "dbdriver/postgres"
)

//  程序入口
func main() {
    sql.Open("postgres", "mydb")
}
