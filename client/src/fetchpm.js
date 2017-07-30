const HOST = '//localhost:3000'
/**
 * 统一处理请求
 * @param  {Object} vm      Vue Example：this
 * @param  {Boolean} auth    是否认证
 * @param  {String} path    目标地址
 * @param  {Object} options 请求的options
 * @return {Promise}         [description]
 */
export default function (vm, auth, path, options) {
  var token = sessionStorage.getItem('token')

  if (options.method == 'POST' && !options.body) options.body = {}

  if (auth) {
    if (token == null || token === '') {
      vm.$router.push({path:'/'})
      return
    } else {
      if(!options.headers) options.headers = {}
      Object.assign(options.headers, {
        'Authorization': 'Bearer ' + token
      })
    }
  }
  path = HOST + path
  if(options.method == 'POST')
    options.body = JSON.stringify(options.body)
  console.info(path, options)
  return fetch(path, options)
}