const path = require('path')

/** @type {import('next').NextConfig} */
module.exports = {
  reactStrictMode: true,
  async rewrites() {
    return [
      {
        source: '/stream',
        destination: 'http://127.0.0.1/stream',
      },
      {
        source: '/meta',
        destination: 'http://127.0.0.1/meta',
      },
    ]
  },
  sassOptions: {
    includePaths: [path.join(__dirname, 'styles')],
  },
}
