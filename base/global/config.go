package global

import (
	"strings"

	"github.com/rockin0098/flash/base/datasource"
)

type ServerConfig struct {
	NodeID     string
	DataSource []datasource.DataSource
}

var serverConfig = &ServerConfig{}

func PurifyConfig(content string) string {

	lines := strings.Split(content, "\n")

	res := ""

	for _, line := range lines {
		idx := strings.Index(line, "//")
		if idx == -1 { // 无注释,返回trim过后的行
			res = res + strings.TrimSpace(line) + "\n"
			continue
		}

		// 截取注释符号之前的字符串
		sub := string(line[:idx])
		sub = strings.TrimSpace(line)

		res = res + sub + "\n"
	}

	return res
}
