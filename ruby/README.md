# Usage

```
require './client.rb'

endpoint ='https://yourserver.herokuapp.com'

extheader = {
  "Accept-Charset" => "UTF-8",
  "Content-Type" => "application/x-www-form-urlencoded; charset=UTF-8"
}
client = Client.new(@@endpoint, extheader)
body = {
  'name' => 'ruby',
  'email' => 'ruby@gmail.com'
}
res = client.post("/users", body)
body = JSON.parse(res.body)
puts res.status
puts body
```
