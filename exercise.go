package golang

import "fmt"
import "strings"

// func myPrintf(err error, s fmt.Stringer) string {
// 	return err.Error() + " " + s.String()
// }

func myPrintf(v ...interface{}) string {
	s := []string{}
	for _, v := range v {
		switch v.(type) {
		case string:
			s = append(s, v.(string))
		case error:
			s = append(s, v.(error).Error())
		case fmt.Stringer:
			s = append(s, v.(fmt.Stringer).String())
		case int:
			s = append(s, number(v.(int)).String())
		}
	}
	return strings.Join(s, " ")
}
