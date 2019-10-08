extern crate reqwest;
use reqwest::{RequestBuilder, StatusCode};
use reqwest::header::{HeaderMap, HeaderName, HeaderValue, CONTENT_TYPE};
use std::collections::HashMap;
use std::io::Read;

#[derive(Debug)]
pub struct Client {
    pub client: reqwest::Client,
    pub base_url: &'static str,
    pub headers: HashMap<&'static str, &'static str>,
}

impl Client {
    
    fn new(base_url: &'static str, headers: HashMap<&'static str, &'static str>) -> Client {
        Client {
            client: reqwest::Client::new(),
            base_url: base_url,
            headers: headers,
        }
    }

    fn get(&self, path: &'static str, params: HashMap<&str, &str>) -> Result<(), String> {
        self.send("GET", path, params)
    }

    fn post(&self, path: &'static str, params: HashMap<&str, &str>) -> Result<(), String> {
        self.send("POST", path, params)
    }

    fn put(&self, path: &'static str, params: HashMap<&str, &str>) -> Result<(), String> {
        self.send("PUT", path, params)
    }

    fn delete(&self, path: &'static str, params: HashMap<&str, &str>) -> Result<(), String> {
        self.send("DELETE", path, params)
    }

    fn patch(&self, path: &'static str, params: HashMap<&str, &str>) -> Result<(), String> {
        self.send("PATCH", path, params)
    }

    fn construct_headers(&self) -> HeaderMap {
        let mut header_map = HeaderMap::new();
        for (k, v) in &self.headers {
            let k = String::from(*k);
            let v = String::from(*v);
            header_map.insert(
                HeaderName::from_static(Client::string_to_static_str(k.to_lowercase())), 
                HeaderValue::from_static(Client::string_to_static_str(v.to_lowercase()))
            );
        }
        header_map
    }

    fn send(&self, method: &str, path: &'static str, params: HashMap<&str, &str>) -> Result<(), String> {
        let client = reqwest::Client::new();
        let headers = self.construct_headers();
        let mut builder: RequestBuilder;
        let url = &format!("{}{}", self.base_url, path);
        match method {
            "GET" => {
                builder = client.get(url);
                builder = builder.query(&params);
            }
            "POST" => {
                builder = client.post(url);
                builder = builder.form(&params);
            }
            "PUT" => {
                builder = client.put(url);
                builder = builder.form(&params);
            }
            "DELETE" => {
                builder = client.delete(url);
                builder = builder.query(&params);
            }
            "PATCH" => {
                builder = client.patch(url);
                builder = builder.form(&params);
            }
            _ => {
                return Err("method error".to_string());
            }
        }

        let content_type = headers.get(CONTENT_TYPE).unwrap();
        if content_type == HeaderValue::from_static("application/json") {
            builder = builder.json(&params);
        }

        builder = builder.headers(headers);

        let mut res = builder.send().unwrap();

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

    fn string_to_static_str(s: String) -> &'static str {
        Box::leak(s.into_boxed_str())
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn get() {
        let mut headers = HashMap::new();
        headers.insert("Content-Type", "application/json");
        let client = Client::new("http://localhost:8080", headers);
        let mut params = HashMap::new();
        params.insert("name", "dongri");
        client.get("/", params).unwrap();
        // assert_eq!(e.encode_word("abc def".to_string()), "=?UTF-8?Q?abc_def?=");
    }

    #[test]
    fn post() {
        let mut headers = HashMap::new();
        headers.insert("Content-Type", "application/x-www-form-urlencoded");
        let client = Client::new("http://localhost:8080", headers);
        let mut params = HashMap::new();
        params.insert("name", "dongri");
        client.post("/", params).unwrap();
        // assert_eq!(e.encode_word("abc def".to_string()), "=?UTF-8?Q?abc_def?=");
    }

    #[test]
    fn post_json() {
        let mut headers = HashMap::new();
        headers.insert("Content-Type", "application/json");
        let client = Client::new("http://localhost:8080", headers);
        let mut params = HashMap::new();
        params.insert("name", "dongri");
        client.put("/", params).unwrap();
        // assert_eq!(e.encode_word("abc def".to_string()), "=?UTF-8?Q?abc_def?=");
    }

    #[test]
    fn put() {
        let mut headers = HashMap::new();
        headers.insert("Content-Type", "application/x-www-form-urlencoded");
        let client = Client::new("http://localhost:8080", headers);
        let mut params = HashMap::new();
        params.insert("name", "dongri");
        client.put("/", params).unwrap();
        // assert_eq!(e.encode_word("abc def".to_string()), "=?UTF-8?Q?abc_def?=");
    }

    #[test]
    fn delete() {
        let mut headers = HashMap::new();
        headers.insert("Content-Type", "application/x-www-form-urlencoded");
        let client = Client::new("http://localhost:8080", headers);
        let mut params = HashMap::new();
        params.insert("name", "dongri");
        client.delete("/", params).unwrap();
        // assert_eq!(e.encode_word("abc def".to_string()), "=?UTF-8?Q?abc_def?=");
    }

    #[test]
    fn patch() {
        let mut headers = HashMap::new();
        headers.insert("Content-Type", "application/x-www-form-urlencoded");
        let client = Client::new("http://localhost:8080", headers);
        let mut params = HashMap::new();
        params.insert("name", "dongri");
        client.patch("/", params).unwrap();
        // assert_eq!(e.encode_word("abc def".to_string()), "=?UTF-8?Q?abc_def?=");
    }

}
