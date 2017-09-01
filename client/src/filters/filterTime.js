/**
 * 传入time为以秒为单位的时间，所以在new Date()时需要*1000
 * @param  {Int} time 以秒为单位的整数
 * @return {[type]}      [description]
 */
export default function (time) {
  var result = new Date(time * 1000)
  return result.getFullYear() + '年' 
      + (result.getMonth() + 1 ) + '月' 
      + result.getDate() +'日 '
      + (result.getHours() > 9 ? result.getHours() : '0' + result.getHours() ) + ':'
      + (result.getMinutes() > 9 ? result.getMinutes() : '0' + result.getMinutes()) + ':'
      + (result.getSeconds() > 9 ? result.getSeconds() : '0' + result.getSeconds())
}