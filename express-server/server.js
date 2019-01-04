import express from 'express'
import bodyParser from 'body-parser'
const server = express()
const port = 9000;

server.use(bodyParser.json())
server.use(bodyParser.urlencoded({
    extended: true
}))

server.get('/', (req, res) => {
    res.send('ok')
})

server.listen(port, function (err, result) {
    console.log('running in port http://localhost:' + port)
})

export default server