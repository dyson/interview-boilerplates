const express = require('express');
const app = express();

const port = process.env.port || 3000;

app.get('/hello', (req, res) => {
  res.json({'hello': 'world'});
});

const server = app.listen(port, () => {
  console.log('Listening on http://localhost:%s', port);
});
module.exports = server;
