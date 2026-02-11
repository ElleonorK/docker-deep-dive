const express = require('express');
const os = require('os');

const app = express();
const PORT = process.env.PORT || 8080;
const VERSION = process.env.APP_VERSION || 'unknown';

app.get('/', (req, res) => {
  res.json({
    message: 'Hello from Docker!',
    version: VERSION,
    hostname: os.hostname(),
    platform: os.platform(),
    timestamp: new Date().toISOString()
  });
});

app.get('/health', (req, res) => {
  res.json({ status: 'healthy' });
});

app.listen(PORT, '0.0.0.0', () => {
  console.log(`Simple web app v${VERSION} listening on port ${PORT}`);
});
