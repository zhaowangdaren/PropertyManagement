var path = require('path');
var webpack = require('webpack');
var HtmlWebpackPlugin = require('html-webpack-plugin');
var DashboardPlugin = require('webpack-dashboard/plugin');
var px2rem = require('postcss-px2rem');

const vuxLoader = require('vux-loader')

const webpackConfig = {
  entry: ['whatwg-fetch','./src/app.js'],
  output: {
    path: path.resolve(__dirname, 'dist'),
    publicPath: '/',
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
      { test: /\.css$/, loader: 'style-loader!css-loader!px2rem-loader?remUnit=150&remPrecision=8' },
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
          },
          postcss: function() {
            return [px2rem({remUnit: 150})];
          }
        }
      },
      {
        test: /\.(png|jpe?g|gif|svg)(\?.*)?$/, 
        loader: 'url-loader',
        query: {
          limit: 8192,
          name: 'images/[name].[ext]'
        }
      },
      {
        test: /\.(eot|svg|ttf|woff|woff2)(\?\S*)?$/,
        loader: 'file-loader'
      }
    ]
  },
  externals: {
    "vue": "Vue"
  },
  plugins:[
    new DashboardPlugin(),
    new webpack.HotModuleReplacementPlugin(),
    new webpack.NoEmitOnErrorsPlugin(),
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
    port: 9000,
    host: '0.0.0.0',
    disableHostCheck: true
  }
}

// module.exports = vuxLoader.merge(webpackConfig, {
//   plugins: ['vux-ui']
// })
module.exports = webpackConfig
