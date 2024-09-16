use glob::glob;
use std::collections::HashMap;

pub fn run() -> Result<(), Box<dyn std::error::Error>> {
    let files = glob("**/*")?;
    let mut map: HashMap<String, i32> = HashMap::new();
    for f in files.flatten() {
        let ext = f
            .extension()
            .unwrap_or_default()
            .to_string_lossy()
            .to_string();
        if f.is_file() && !ext.is_empty() {
            if let Some(val) = map.get_mut(&ext) {
                *val += 1;
            } else {
                map.insert(ext, 1);
            }
        }
    }
    let mut mapvec: Vec<_> = map.iter().map(|(k, v)| (v, k)).collect();
    mapvec.sort();
    for (v, k) in mapvec.iter().rev() {
        println!("{:3} {}", v, k);
    }
    Ok(())
}
