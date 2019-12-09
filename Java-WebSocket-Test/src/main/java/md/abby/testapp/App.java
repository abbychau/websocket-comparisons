package md.abby.testapp;

import java.net.InetSocketAddress;

import org.java_websocket.server.WebSocketServer;


/**
 * Hello world!
 *
 */
public class App 
{

	public static void main(String[] args) {
		String host = "p2";
		int port = 8080;

		WebSocketServer server = new SimpleServer(new InetSocketAddress(host, port));
		server.run();
	}
}
