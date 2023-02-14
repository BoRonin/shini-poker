module.exports = {
    apps: [
      {
        name: 'shini-poker',
        exec_mode: 'cluster',
        instances: 'max',
        script: './.output/server/index.mjs',
        port: 3005,
      }
    ]
  }