use std::path::PathBuf;

use rocket::fs::NamedFile;
use rocket::response::content::RawHtml;
use rocket::response::status::NotFound;
use rocket::tokio::runtime;

#[get("/<file..>")]
async fn files(file: PathBuf) -> Result<NamedFile, NotFound<String>> {
    let path = std::path::Path::new(&file);
    NamedFile::open(&path)
        .await
        .map_err(|e| NotFound(e.to_string()))
}

#[get("/")]
pub fn index() -> RawHtml<String> {
    RawHtml(std::fs::read_to_string("index.html").unwrap().to_string())
}

pub fn run() -> Result<(), Box<dyn std::error::Error>> {
    let _ = runtime::Runtime::new()?.block_on(async {
        rocket::build()
            .mount("/", routes![index, files])
            .launch()
            .await
    });
    Ok(())
}
