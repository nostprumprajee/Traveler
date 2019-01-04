var express = require('express');
var app = express();
var path = require('path');
var port = 8081;
 
var server = app.listen(port, function() {
    console.log('Listening on port: ' + port);
});
var io = require('socket.io').listen(server);
 

app.set('views', path.join(__dirname, 'views'));
app.set('view engine', 'jade'); 
app.use(express.static('public'));
 
app.get('/', function(req, res) {
    res.render('index');
});
 
io.on('connection', function(socket) {
    socket.on('chat', function(data) {
        io.emit('chat', data);
    });
});