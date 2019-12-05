
use std::thread;
use websocket::sync::Server;
use websocket::OwnedMessage;

fn main() {
	let server = Server::bind("127.0.0.1:8080").unwrap();

	for request in server.filter_map(Result::ok) {
		// Spawn a new thread for each connection.
		thread::spawn(|| {
			// if !request.protocols().contains(&"abby-test".to_string()) {
			// 	request.reject().unwrap();
			// 	return;
			// }

			let mut client = request.accept().unwrap();//use_protocol("abby-test")

			let ip = client.peer_addr().unwrap();

			//println!("Connection from {}", ip);

			let message = OwnedMessage::Text("Hello".to_string());
			client.send_message(&message).unwrap();

			let (mut receiver, mut sender) = client.split().unwrap();

			for message in receiver.incoming_messages() {
				let message = message.unwrap();
                

                
				match message {
					OwnedMessage::Close(_) => {
						let message = OwnedMessage::Close(None);
						sender.send_message(&message).unwrap();
						//println!("Client {} disconnected", ip);
						return;
					}
					OwnedMessage::Ping(ping) => {
                        
						let message = OwnedMessage::Pong(ping);
						sender.send_message(&message).unwrap();
					}
					_ => {
                        println!("{:?}",message);
                        sender.send_message(&message).unwrap()
                    }
				}
			}
		});
	}
}