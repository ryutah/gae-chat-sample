import path from "path"

export default {
  entry: path.join(__dirname, "../src/front/app.jsx"),
  output: {
    publicPath: "/",
    path: path.join(__dirname, "../app/web"),
    filename: "bundle.js",
  },
  // when want to export multiple entry
  // entry: {
  //   app: "./app.js",
  //   search: "./search.js"
  // },
  //
  // output: {
  //   filename: "[name].bundle.js",
  //   path: path.join(__dirname, "public"),
  // },

  target: "web",
  resolve: {
    extensions: ["*", ".js", ".jsx"],
  },
  module: {
    loaders: [
      {
        test: /\.jsx?$/,
        // include: [path.resolve(__dirname, "src")],
        exclude: /node_modules/,
        loader: "babel-loader",
      }
    ]
  },
}
