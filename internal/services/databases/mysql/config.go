package mysql

import "fmt"

type MySqlConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string

	Params map[string]string
}

func (c MySqlConfig) DSN() string {
	var params string
	if len(c.Params) > 0 {
		var query string
		for key, val := range c.Params {
			if query != "" {
				query += "&"
			}
			query += key + "=" + val
		}
		params = "?" + query
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s%s",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Name,
		params)
}
