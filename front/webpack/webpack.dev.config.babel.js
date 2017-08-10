import config from "./webpack.base.config.babel"

config.devtool = "inline-source-map"
config.devServer = {
  inline: true,
  contentBase: "../server/src/web/static",
  port: "3000",
  host: "0.0.0.0",
}

export default config
