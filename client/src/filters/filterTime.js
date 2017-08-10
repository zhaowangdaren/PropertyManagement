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
      + (result.getHours() + 1) + ':'
      + result.getMinutes() + ':'
      + result.getSeconds()
}