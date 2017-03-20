# Usage

```
package main

const (
	endpoint = "https://yourserver.herokuapp.com"
	timeout  = 90 //Second
)

func main() {
  header := map[string][]string{
    "Content-Type": {string(ContentTypeUrlencoded)},
  }
  client := NewClient(endpoint, header, timeout)
  params := map[string][]string{
    "name":  {"dongri"},
    "email": {"dongri@domain.com"},
  }
  res, err := client.Post("/users", params)
  if err != nil {
    fmp.Errorf("got error %v", err)
  }
  defer res.Body.Close()
  var resBody interface{}
  if err := json.NewDecoder(res.Body).Decode(&resBody); err != nil {
    fmp.Errorf("got error %v", err)
  }
  fmp.Print(resBody)
}
```
