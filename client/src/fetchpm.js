// import { Message } from 'element-ui'

const HOST = 'https://www.maszfglzx.com:3000'
// const HOST = '//localhost:3000'


const ThirtyMinute = 30 * 60 * 1000
function refreshToken (vm) {
  return new Promise((resolve, reject) => {
    var user = JSON.parse(sessionStorage.getItem('user')) || {}
    var expireTime = user.expire ? new Date(user.expire) : new Date()
    var timeLeft = expireTime.getTime() - new Date ().getTime()
    console.info('timeLeft', timeLeft)
    if ( 0 < timeLeft && timeLeft <= ThirtyMinute) {
      console.info('refreshToken')
      fetch(HOST + '/pm/refresh_token', {
        method: 'POST',
        headers: {
          'Authorization': 'Bearer ' + user.token
        }
      }).then(response => {
        console.info('refreshToken',response)
        return response.json()
      }).then(body => {
        if (body.error === 0) {
          Object.assign(user, body.data)
          sessionStorage.setItem('user', JSON.stringify(user))
        }
        resolve()
      })
    } else {
      resolve()
    }
  })
}
/**
 * 统一处理请求
 * @param  {Object} vm      Vue Example：this
 * @param  {Boolean} auth    是否认证
 * @param  {String} path    目标地址
 * @param  {Object} options 请求的options
 * @return {Promise}         [description]
 */
export default function (vm, auth, path, options) {
  return refreshToken(vm).then(() => {
    var user = JSON.parse(sessionStorage.getItem('user')) || {}
    var token = user.token

    if (options.method !== 'GET' && !options.body) options.body = {}

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
    if(options.method !== 'GET') options.body = JSON.stringify(options.body)
    console.info(path, options)
    return fetch(path, options).then(resp => {
      if (resp.status == 401 && vm.$route.path !== '/login') {//认证失败，重新登录
        vm.$router.push({path:'/'})
        return resp
      } else {
        return resp
      }
    })
  })
}
