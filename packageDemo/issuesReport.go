package packageDemo

import "time"

const temp1 = `{{.TotalCount}} issues:
{{range .Items}}-------------
Number:{{.Number}}
User:{{.User.login}}
Title:{{.Title|printf "%.64s"}}  //|表示将当前表达式的结果作为后一个函数输入
Age:{{.CreateAt|daysAgo}} days
{{end}}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}
