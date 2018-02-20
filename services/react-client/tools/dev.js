import 'colors'
import _ from 'lodash'
import opn from 'opn'
import express from 'express'
import webpack from 'webpack'
import webpackDevMiddleware from 'webpack-dev-middleware'
import webpackHotMiddleware from 'webpack-hot-middleware'
import jsonServer from 'json-server'

import webpackConfig from '../webpack.dev'
import mockData from '../api/data'

const port = process.env.PORT || 4000
const hmr = 'webpack-hot-middleware/client?path=/webpack_hmr&timeout=2000'

const app = express()
const transpiler = webpack(webpackConfig)

// Use json-server for mocking api backend
const apiRouter = jsonServer.router(mockData)
const apiMiddleware = jsonServer.defaults()

// Update config for hot reloading
if (_.isArray(webpackConfig.entry.app)) {
  webpackConfig.entry.app.unshift(hmr)
} else if (_.isString(webpackConfig.entry.app)) {
  webpackConfig.entry.app = [ hmr, webpackConfig.entry.app ]
}

app.use(webpackDevMiddleware(transpiler, {
  noInfo: true,
  publicPath: webpackConfig.output.publicPath
}))

app.use(webpackHotMiddleware(transpiler, {
  path: '/webpack_hmr',
  heartbeat: 1000
}))

// Api mock server
app.use('/api/v1', apiMiddleware, apiRouter)

app.listen(port, err => {
  if (err) {
    console.log(err.red)
    return 1
  }
  console.log(`Dev server listening on http://localhost:${port}`.bold.green)
  opn(`http://localhost:${port}`)
})
