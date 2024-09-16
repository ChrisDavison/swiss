pub fn run(level: i8, text: &[String]) -> Result<(), Box<dyn std::error::Error>> {
    let text = text.join(" ");
    match level {
        1 => println!("{}\n{}", text, "=".repeat(text.len())),
        2 => println!("{}\n{}", text, "-".repeat(text.len())),
        n => println!("{} {}", "#".repeat(n as usize), text),
    }
    Ok(())
}
