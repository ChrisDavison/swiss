pub fn run(text: String) -> Result<(), Box<dyn std::error::Error>> {
    let dashes = "─".repeat(text.len() + 2);
    println!("┌{}┐", dashes);
    println!("│ {} │", text);
    println!("└{}┘", dashes);
    Ok(())
}
