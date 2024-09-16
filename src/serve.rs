const template: &str = include_str!("template.html");

#[get("/hello/<name>/<age>")]
fn hello(name: &str, age: u8) -> String {
    format!("Hello, {} year old named {}!", age, name)
}

#[get("/")]
fn index() -> String {
    template.to_string()
}

// #[launch]
// fn rocket() -> _ {
// }

pub fn run() -> Result<(), Box<dyn std::error::Error>> {
    rocket::build().mount("/", routes![hello, index]);
    Ok(())
}
