use websocket::sync::Server;
use websocket::OwnedMessage;

fn main() {
    let server = Server::bind("p2:8080").unwrap();

    for request in server.filter_map(Result::ok) {
        // Spawn a new thread for each connection.

        // if !request.protocols().contains(&"abby-test".to_string()) {
        // 	request.reject().unwrap();
        // 	return;
        // }

        let client = request.accept().unwrap();
        //use_protocol("abby-test")

        let ip = client.peer_addr().unwrap();
        println!("Connection from {}", ip);

        let message = OwnedMessage::Text(std::iter::repeat("X").take(10000).collect::<String>());

        let (mut receiver, mut sender) = client.split().unwrap();
        sender.send_message(&message).unwrap();

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
                    for _ in 0..1000 {
                        println!("{:?}", message);
                        sender.send_message(&message).unwrap();
                    }
                }
            }
        }
    }
}
