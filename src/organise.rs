use shellexpand::tilde;
use std::path::PathBuf;

fn organise_folder(direc: String) -> Result<Vec<PathBuf>, Box<dyn std::error::Error>> {
    let resolved = tilde(&direc);
    let pat = format!("{}/*", resolved);
    dbg!(&pat);
    // let files = glob::glob(pat.to_string().as_str())?;
    // dbg!(files);
    Ok(vec![])
}

pub fn run() -> Result<(), Box<dyn std::error::Error>> {
    let mut remaining = Vec::new();
    for direc in [
        "~/Downloads",
        "~/syncthing",
        "/mnt/c/Users/davison/Downloads",
    ] {
        remaining.extend(organise_folder(direc.into()))
    }

    if !remaining.is_empty() {
        println!("Unorganised:");
        for thing in remaining {
            println!("    {:?}", thing);
        }
    }

    Ok(())
}
