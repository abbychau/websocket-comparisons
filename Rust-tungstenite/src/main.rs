#[macro_use]
extern crate may;

use may::net::TcpListener;
use tungstenite::server::accept;
// use tokio::net::TcpListener;
use tungstenite::Message;

fn main() {
    let server = TcpListener::bind("ps2:8080").unwrap();
    for stream in server.incoming() {
        go!(move || {
            let mut websocket = accept(stream.unwrap()).unwrap();
            loop {
                match websocket.read_message() {
                    Ok(msg) => {
                        // We do not want to send back ping/pong messages.
                        if msg.is_binary() || msg.is_text() {
                            for _ in 0..10 {
                                match websocket.write_message(Message::Text("0".repeat(10000))) {
                                    Ok(_) => {}
                                    Err(e) => println!("End point disconnected? {}", e),
                                }
                            }
                        }
                    }
                    Err(e) => 
                    {println!("End point disconnected? {}", e);
                    break;
                    },
                }
            }
        });
    }
}

//tokio listener with thread
// use std::net::TcpListener;
// use std::thread::spawn;
// use tungstenite::server::accept;
// // use tokio::net::TcpListener;
// use tungstenite::Message;
// use tokio::prelude::*;

// /// A WebSocket echo server
// fn main(){
// let server = TcpListener::bind("127.0.0.1:9001").unwrap();
// for stream in server.incoming() {
//     spawn (move || {
//         let mut websocket = accept(stream.unwrap()).unwrap();
//         loop {
//             let msg = websocket.read_message().unwrap();

//             // We do not want to send back ping/pong messages.
//             if msg.is_binary() || msg.is_text() {
//                 websocket.write_message(
//                     Message::Text("0".repeat(1000))
//                 ).unwrap();
//             }
//         }
//     });
// }
// }
