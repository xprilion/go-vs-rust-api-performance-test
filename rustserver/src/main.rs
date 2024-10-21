use warp::Filter;

#[tokio::main]
async fn main() {
    // Define a route
    let route = warp::path::param()
        .map(|path: String| format!("Hello, visitor! You've requested: /{}", path));

    // Start the server on port 8081
    warp::serve(route).run(([127, 0, 0, 1], 8081)).await;
}
