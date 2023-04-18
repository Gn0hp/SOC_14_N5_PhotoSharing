package mysql

import "fmt"

type MySqlConfig struct {
	Host     string `mapstructure:"host"`
	Port     int64  `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`

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
