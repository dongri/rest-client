require 'test/unit'
require './client.rb'

class ClientTest < Test::Unit::TestCase

  @@endpoint ='https://yourserver.herokuapp.com'

  def test_get
    client = Client.new(@@endpoint, nil)
    res = client.get("/users", nil)
    assert_equal 200, res.status
    body = JSON.parse(res.body)
    assert_equal 2, body.length
  end

  def test_post
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
    assert_equal 201, res.status
    body = JSON.parse(res.body)
    assert_equal "ruby", body["name"]
    assert_equal "ruby@gmail.com", body["email"]
  end

  def test_put
    client = Client.new(@@endpoint, nil)
    body = {
      'name' => 'ruby2',
      'email' => 'ruby2@gmail.com'
    }
    res = client.put("/users/1", body)
    assert_equal 200, res.status
    body = JSON.parse(res.body)
    assert_equal "ruby2", body["name"]
    assert_equal "ruby2@gmail.com", body["email"]
  end

  def test_delete
    client = Client.new(@@endpoint, nil)
    res = client.delete("/users/1", nil)
    assert_equal 200, res.status
  end

end
