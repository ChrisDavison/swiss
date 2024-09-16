use std::io::Write;

pub fn run(
    files: &[std::path::PathBuf],
    filename_out: String,
) -> Result<(), Box<dyn std::error::Error>> {
    let mut contents = Vec::new();
    for filename in files {
        let filetext = std::fs::read_to_string(filename)?;
        contents.push(filetext);
    }
    let mut f_out = std::fs::File::create_new(filename_out)?;
    f_out.write_all(contents.join("\n").as_bytes())?;
    Ok(())
}
