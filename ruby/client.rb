require 'httpclient'
require 'json'

uri ='http://localhost:8080/users'

client = HTTPClient.new()

body = { 'name' => 'ruby', 'email' => 'ruby@gmail.com' }


res = client.post(uri, body)

p res.status
p res.contenttype
puts JSON.parse(res.body)


require 'digest/sha2'
hashed = Digest::SHA256.hexdigest "4123234567879090"
puts hashed
