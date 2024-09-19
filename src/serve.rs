use axum::Router;
use tower_http::services::ServeFile;

#[tokio::main]
pub async fn run() -> Result<(), Box<dyn std::error::Error>> {
    // build our application with a route
    let app = Router::new()
        // `GET /` goes to `root`
        .nest_service("/", ServeFile::new("/"));

    // run our app with hyper, listening globally on port 3000
    let listener = tokio::net::TcpListener::bind("0.0.0.0:8000").await.unwrap();
    axum::serve(listener, app).await.unwrap();
    Ok(())
}
