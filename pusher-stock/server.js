const stockData = require('./stock.json');
const Pusher = require('pusher');

var pusher = new Pusher({
  appId: 'APP_ID',
  key: 'APP_KEY',
  secret: 'APP_SECRET',
  cluster: 'APP_CLUSTER',
  encrypted: true
});

let i = 0;
setInterval(() => {
  // const NDAQ = stockData[0]['Trades'][i];
  const GOOG = stockData[1]['Trades'][i];
  pusher.trigger('trade', 'stock', GOOG);
  i++
}, 2000);
