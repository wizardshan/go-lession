
package ent

import (
	"database/sql"
	entSql "entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
)

func (c *Client) DB() *sql.DB {
	return c.driver.(*entSql.Driver).DB()
}
