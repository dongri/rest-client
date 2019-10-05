# Usage

```go
package main

const (
	endpoint = "https://yourserver.herokuapp.com"
	timeout  = 90 //Second
)

func main() {
	header := map[string]string{
		"X-AccessToken": "hoge",
	}
	client := NewClient(endpoint, ContentTypeJSON, header, timeout)
	params := map[string]string{
		"name": "dongri",
	}
	res, err := client.Post("/", params)
	if err != nil {
		fmt.Errorf("got error %v", err)
	}
	defer res.Body.Close()
	var resBody interface{}
	if err := json.NewDecoder(res.Body).Decode(&resBody); err != nil {
		fmt.Errorf("got error %v", err)
	}
	fmt.Println(resBody)
}
```
