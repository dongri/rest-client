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
    
    pub fn new(base_url: &'static str, headers: HashMap<&'static str, &'static str>) -> Client {
        Client {
            client: reqwest::Client::new(),
            base_url: base_url,
            headers: headers,
        }
    }

    pub fn get(&self, path: &'static str, params: HashMap<&str, &str>) -> Result<String, String> {
        self.send("GET", path, params)
    }

    pub fn post(&self, path: &'static str, params: HashMap<&str, &str>) -> Result<String, String> {
        self.send("POST", path, params)
    }

    pub fn put(&self, path: &'static str, params: HashMap<&str, &str>) -> Result<String, String> {
        self.send("PUT", path, params)
    }

    pub fn delete(&self, path: &'static str, params: HashMap<&str, &str>) -> Result<String, String> {
        self.send("DELETE", path, params)
    }

    pub fn patch(&self, path: &'static str, params: HashMap<&str, &str>) -> Result<String, String> {
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

    fn send(&self, method: &str, path: &'static str, params: HashMap<&str, &str>) -> Result<String, String> {
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
        match headers.get(CONTENT_TYPE) {
            Some(content_type) => {
                if content_type == HeaderValue::from_static("application/json") {
                    builder = builder.json(&params);
                }
            },
            None => {}
        };
        builder = builder.headers(headers);
        let mut res = builder.send().unwrap();
        if res.status() != StatusCode::OK {
            match res.error_for_status() {
                Err(err) => {
                    return Err(err.to_string());
                }
                Ok(res) => {
                    return Err(res.status().to_string());
                }
            }
        }
        let mut buf = String::new();
        res.read_to_string(&mut buf).expect("Failed to read response");
        return Ok(buf);
    }

    fn string_to_static_str(s: String) -> &'static str {
        Box::leak(s.into_boxed_str())
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    const BASE_URL: &'static str = "http://localhost:8080";

    #[test]
    fn get() {
        let mut headers = HashMap::new();
        headers.insert("Content-Type", "application/json");
        let client = Client::new(BASE_URL, headers);
        let mut params = HashMap::new();
        params.insert("name", "dongri");
        match client.get("/", params) {
            Ok(body) => {
                println!("Body: {:?}", body);
            },
            Err(err) => {
                println!("Error: {:?}", err);
            }
        }
    }

    #[test]
    fn post() {
        let mut headers = HashMap::new();
        headers.insert("Content-Type", "application/x-www-form-urlencoded");
        let client = Client::new(BASE_URL, headers);
        let mut params = HashMap::new();
        params.insert("name", "dongri");
        match client.post("/", params) {
            Ok(body) => {
                println!("Body: {:?}", body);
            },
            Err(err) => {
                println!("Error: {:?}", err);
            }
        }
    }

    #[test]
    fn put() {
        let mut headers = HashMap::new();
        headers.insert("Content-Type", "application/x-www-form-urlencoded");
        let client = Client::new(BASE_URL, headers);
        let mut params = HashMap::new();
        params.insert("name", "dongri");
        match client.put("/", params) {
            Ok(body) => {
                println!("Body: {:?}", body);
            },
            Err(err) => {
                println!("Error: {:?}", err);
            }
        }
    }

    #[test]
    fn delete() {
        let mut headers = HashMap::new();
        headers.insert("Content-Type", "application/x-www-form-urlencoded");
        let client = Client::new(BASE_URL, headers);
        let mut params = HashMap::new();
        params.insert("name", "dongri");
        match client.delete("/", params) {
            Ok(body) => {
                println!("Body: {:?}", body);
            },
            Err(err) => {
                println!("Error: {:?}", err);
            }
        }
    }

    #[test]
    fn patch() {
        let mut headers = HashMap::new();
        headers.insert("Content-Type", "application/x-www-form-urlencoded");
        let client = Client::new(BASE_URL, headers);
        let mut params = HashMap::new();
        params.insert("name", "dongri");
        match client.patch("/", params) {
            Ok(body) => {
                println!("Body: {:?}", body);
            },
            Err(err) => {
                println!("Error: {:?}", err);
            }
        }
    }

    #[test]
    fn json() {
        let mut headers = HashMap::new();
        headers.insert("Content-Type", "application/json");
        let client = Client::new(BASE_URL, headers);
        let mut params = HashMap::new();
        params.insert("name", "dongri");
        match client.post("/", params) {
            Ok(body) => {
                println!("Body: {:?}", body);
            },
            Err(err) => {
                println!("Error: {:?}", err);
            }
        }
    }

}
