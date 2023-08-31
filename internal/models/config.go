package models

type Config struct {
	HostRedis   string `yml:"host_redis"`
	PortRedis   string `yml:"port_redis"`
	HostSQL     string `yml:"host"`
	PortSQL     string `yml:"port"`
	UsernameSQL string `yml:"username"`
	PasswordSQL string `yml:"password"`
	DBName      string `yml:"dbname"`
	SSLmode     string `yml:"sslmode"`
}
