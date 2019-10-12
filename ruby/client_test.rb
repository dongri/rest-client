require 'test/unit'
require './client.rb'

class ClientTest < Test::Unit::TestCase

  @@endpoint ='https://yourserver.herokuapp.com'
  # @@endpoint ='http://localhost:8080'

  def test_get
    headers = {
      "Content-Type" => "application/x-www-form-urlencoded"
    }
    client = Client.new(@@endpoint, headers)
    params = {
      "name" => "dongri"
    }
    res = client.get("/", params)
    assert_equal "200", res.code
    body = JSON.parse(res.body)
    puts body
  end

  def test_post
    headers = {
      "Content-Type" => "application/x-www-form-urlencoded"
    }
    client = Client.new(@@endpoint, headers)
    params = {
      "name" => "dongri"
    }
    res = client.post("/", params)
    assert_equal "200", res.code
    body = JSON.parse(res.body)
    puts body
  end

  def test_put
    headers = {
      "Content-Type" => "application/x-www-form-urlencoded"
    }
    client = Client.new(@@endpoint, headers)
    params = {
      "name" => "dongri",
    }
    res = client.put("/", params)
    assert_equal "200", res.code
    body = JSON.parse(res.body)
    puts body
  end

  def test_delete
    headers = {
      "Content-Type" => "application/x-www-form-urlencoded"
    }
    client = Client.new(@@endpoint, headers)
    params = {
      "name" => "dongri",
    }
    res = client.delete("/", params)
    assert_equal "200", res.code
    body = JSON.parse(res.body)
    puts body
  end

  def test_post_json
    headers = {
      "Content-Type" => "application/json"
    }
    client = Client.new(@@endpoint, headers)
    params = {
      "name" => "dongri",
    }
    res = client.post("/", params)
    assert_equal "200", res.code
    body = JSON.parse(res.body)
    puts body
  end

end
