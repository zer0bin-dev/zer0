use std::{fs::File, path::PathBuf};

use serde::{Serialize, Deserialize};

const DEFAULT_INSTANCE: &str = "https://zer0b.in";

#[derive(Serialize, Deserialize, Clone)]
pub struct Config {
    #[serde(default = "default_instance")]
    pub instance: String
}

fn default_instance() -> String {
    DEFAULT_INSTANCE.to_string()
}

pub fn load(path: PathBuf) -> Config {
    let file = File::open(path).expect("Failed to load config");
    serde_json::from_reader(file).unwrap()
}
