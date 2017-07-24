/**
 * 请求
 * @param  {[type]}   request  {
 *                               method: 'GET',
 *                               url:'http://XXX.com',
 *                               headers: {
 *                                 KEY: VALUE
 *                               },
 *                               query: {
 *                                 KEY: VALUE
 *                               }
 *                             }
 * @param  {Function} callback [description]
 * @return {[type]}            [description]
 */
export default function (request, callback) {
  // body...
  console.info('request', request)
  var xmlhttp = null;
  if(window.XMLHttpRequest){
    xmlhttp = new XMLHttpRequest();
  }else if(window.ActiveXObject){
    xmlhttp = new ActiveXObject('Microsoft.XMLHTTP');
  }

  if(xmlhttp != null){
    xmlhttp.open(request.method ? request.method : 'GET', request.url);
    xmlhttp.onload = handler;
    xmlhttp.setRequestHeader("Content-type","application/x-www-form-urlencoded")
    // xmlhttp.setRequestHeader('Content-type','application/json; charset=utf-8');
 
    if (request.query) {
      xmlhttp.send(JSON.stringify(request.query))
    } else {
      xmlhttp.send(null)
    }
  }

  function handler(e){
    console.info('handler', this.responseText)
    callback(this.responseText);
  }
}

function queries(query) {
  if (!query) return ''
  var result = ''
  for ( let key in query) {
    result += key + "=" + query[key] + "&"
  }
  console.info(result)
  return result
}