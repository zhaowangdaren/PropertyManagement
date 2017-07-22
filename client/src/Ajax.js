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
  var xmlhttp = null;
  if(window.XMLHttpRequest){
    xmlhttp = new XMLHttpRequest();
  }else if(window.ActiveXObject){
    xmlhttp - new ActiveXObject('Microsoft.XMLHTTP');
  }

  if(xmlhttp != null){
    
    xmlhttp.open(request.method ? request.method : 'GET', request.url + queries(request.query));
    xmlhttp.onload = handler;
    xmlhttp.setRequestHeader("Content-type","application/x-www-form-urlencoded")
    
    // xmlhttp.send(queries(request.query));
    xmlhttp.send(null)
  }

  function handler(e){
    callback(this.responseText);
  }
}

function queries(query) {
  if (!query) return '?'
  var result = '?'
  for ( let key in query) {
    result += key + "=" + query[key] + "&"
  }
  console.info(result)
  return result
}