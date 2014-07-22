package foo

import(
	"net/http"
	"fmt"
)

func HttpGet() error {
	res, err := http.Get("http://www.google.com.hk")
	if err != nil {
		return err
	}
	fmt.Println(res.Status)
	return nil
}
