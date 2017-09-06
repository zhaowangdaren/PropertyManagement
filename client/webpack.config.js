var path = require('path');
var webpack = require('webpack');
var HtmlWebpackPlugin = require('html-webpack-plugin');
var DashboardPlugin = require('webpack-dashboard/plugin');
var px2rem = require('postcss-px2rem');
var OptimizeCSSPlugin = require('optimize-css-assets-webpack-plugin')

function resolve (dir) {
  return path.join(__dirname, '..', dir)
}

const webpackConfig = {
  entry: ['whatwg-fetch','./src/app.js'],
  // entry: ['./src/app.js'],
  output: {
    path: path.resolve(__dirname, 'dist'),
    publicPath: './',
    filename:'[name].[hash].js'
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
        test: /\.(less|css)$/,
        use: [
          {
              loader: "style-loader" // creates style nodes from JS strings
          },
          {
              loader: "css-loader" // translates CSS into CommonJS
          },
          {
              loader: "less-loader" // compiles Less to CSS
          },
          {
            loader: "px2rem-loader",
            options: {
              remUnit: 150
            }
          }
        ]
      },
      {
        test: /\.js$/,
        loader:'babel-loader',
        exclude: /node_modules/
        // include: [resolve('src'), resolve('test'), resolve('node_modules/vux')]
      },
      {
        test: /\.vue$/,
        loader: 'vue-loader',
        options: {
          cssModules: {
            localIdentName: '[local]-[hash:base64:5]',
            camelCase: true
          },
          loaders: {
            css: [
              "vue-style-loader",
              {
                loader: "css-loader",
                options: {
                  minimize: false,
                  sourceMap: false
                }
              },
              {
                loader: "px2rem-loader",
                options: {
                  remUnit: 150
                }
              }
            ],
            less: [
              "vue-style-loader",
              {
                loader: "css-loader",
                options: {
                  minimize: false,
                  sourceMap: false
                }
              },
              {
                loader: "px2rem-loader",
                options: {
                  remUnit: 150
                }
              },
              {
                loader: "less-loader",
                options: {
                  sourceMap: false
                }
              }
            ]
          }
          // postcss: function() {
          //   return [px2rem({remUnit: 150})];
          // }
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
    // 当接收到热更新信号时，在浏览器console控制台打印更多可读性高的模块名称等信息
    // new webpack.NamedModulesPlugin(),
    new webpack.NoEmitOnErrorsPlugin(),
    new HtmlWebpackPlugin({
      title: 'OP.V2',
      filename: 'index.html',
      template: 'index.html'
    }),
    // Compress extracted CSS. We are using this plugin so that possible
    // duplicated CSS from different components can be deduped.
    new OptimizeCSSPlugin(),
    // new webpack.optimize.UglifyJsPlugin({
    //   compress: {
    //     warnings: false,
    //     drop_console: false,
    //   },
    //   comments: false
    // })
  ],
  devtool: 'source-map',
  // devtool: false,
  devServer: {
    contentBase: path.join(__dirname, 'dist'),
    compress: false,
    port: 9000,
    host: '0.0.0.0',
    disableHostCheck: true,
    hot: true
    // https: true
  }
}

module.exports = webpackConfig
