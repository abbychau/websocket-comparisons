const WebSocket = require('ws');

const wss = new WebSocket.Server({ port: 8080 });

wss.on('connection', function connection(ws) {
  ws.on('message', function incoming(message) {
    for(var i =0 ;i<10;i++){
      ws.send('0'.repeat(10000));
    }
  });
});