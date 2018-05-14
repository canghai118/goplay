package api

import "fmt"

func ApiURL(name string) string {
	return fmt.Sprintf("%s/%s", JIKE_SERVER, name)
}
