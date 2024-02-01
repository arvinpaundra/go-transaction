package config

import "clean-arch/pkg/util"

func GetMysqlUser() string {
	return util.GetEnv("MYSQL_USER", "root")
}

func GetMysqlPassword() string {
	return util.GetEnv("MYSQL_PASSWORD", "")
}

func GetMysqlHost() string {
	return util.GetEnv("MYSQL_HOST", "localhost")
}

func GetMysqlPort() string {
	return util.GetEnv("MYSQL_PORT", "3306")
}

func GetEnvironment() string {
	return util.GetEnv("ENVIRONMENT", "development")
}

func IsDevelopmentEnv() bool {
	return util.GetEnv("ENVIRONMENT", "development") == "development"
}
