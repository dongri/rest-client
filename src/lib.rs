extern crate reqwest;
use reqwest::StatusCode;
use std::collections::HashMap;
use std::io::Read;

#[derive(Debug)]
pub struct Client {
    pub base_url: String,
    pub headers: String,
}

impl Client {

  fn get(&self, path: String, params: HashMap<&str, &str>) -> Result<(), String> {
      self.send(path, params)
  }

  fn post(&self, path: String, params: HashMap<&str, &str>) -> Result<(), String> {
      self.send(path, params)
  }

  fn put() {

  }

  fn delete() {

  }

  fn patch() {

  }

  fn send(&self, path: String, params: HashMap<&str, &str>) -> Result<(), String> {
      let client = reqwest::Client::new();
      // json
      // let mut res = client
      //     .post("http://example.com/sample/v1/SampleService")
      //     .header(ContentType::json())
      //     .body(json_request)
      //     .send()
      //     .unwrap();
      // let mut buf = String::new();
      // res.read_to_string(&mut buf)
      //     .expect("Failed to read response");
      // println!("{}", buf);

      // from
    let mut res = client
        .post(&format!("{}{}", self.base_url, path))
        .form(&params)
        .query(&params)
        .send()
        .unwrap();

    // let mut buf = String::new();
    // res.read_to_string(&mut buf)
    //     .expect("Failed to read response");
    // println!("{}", buf);
    // return Ok(());


    if res.status() != StatusCode::OK {
        match res.error_for_status() {
            Err(text) => {
                return Err(text.to_string());
            }
            Ok(_) => {
                return Ok(());
            }
        }
    }

    let mut buf = String::new();
    res.read_to_string(&mut buf)
        .expect("Failed to read response");
    println!("{}", buf);

    return Ok(());
  }

}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let client = Client{
            base_url: "http://localhost:8080".to_string(),
            headers: "test".to_string(),
        };
        let mut params = HashMap::new();
        params.insert("name", "dongri");
        client.post("/".to_string(), params).unwrap();

        // assert_eq!(e.encode_word("abc def".to_string()), "=?UTF-8?Q?abc_def?=");
    }
}
