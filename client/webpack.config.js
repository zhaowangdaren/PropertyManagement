var path = require('path');
var webpack = require('webpack');
var HtmlWebpackPlugin = require('html-webpack-plugin');
module.exports = {
  entry: ['whatwg-fetch','./src/app.js'],
  output: {
    path: path.resolve(__dirname, 'dist'),
    filename:'[name].js'
  },
  resolve:{
    alias: {
      '@': path.resolve(__dirname, './src')
    },
    extensions: ['.vue', '.js'] //使用的扩展名，解析vue、js等后缀名的文件，在引入时不需填写相应的后缀名了
  },
  module: {
    rules:[
      {
        test: /\.css$/,
        use: [
          'style-loader',
          'css-loader'
        ]
      },
      {
        test: /\.js$/,
        loader:'babel-loader',
        exclude: /node_modules/
      },
      {
        test: /\.vue$/,
        loader: 'vue-loader',
        options: {
          cssModules: {
            localIdentName: '[local]-[hash:base64:5]',
            camelCase: true
          }
        }
      },
      {test: /\.(png|jpe?g|gif|svg)(\?.*)?$/, use:'url-loader?limit=8192'},
      {
        test: /\.(eot|svg|ttf|woff|woff2)(\?\S*)?$/,
        loader: 'file-loader'
      }
    ]
  },
  plugins:[
    new webpack.HotModuleReplacementPlugin(),
    new HtmlWebpackPlugin({
      title: 'OP.V2',
      filename: 'index.html',
      template: 'index.html'
    })
  ],
  devtool: 'source-map',
  devServer: {
    contentBase: path.join(__dirname, 'dist'),
    compress: true,
    port: 9000
  }
}

