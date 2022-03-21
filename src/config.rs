use std::{fs::File, path::PathBuf};

use serde::{Serialize, Deserialize};


#[derive(Serialize, Deserialize, Clone)]
pub struct Config {
    pub instance: Option<String>
}

pub fn load(path: PathBuf) -> Config {
    let file = File::open(path).expect("Failed to load config");
    serde_json::from_reader(file).unwrap()
}
