export default function (imgsStr) {
  if (!imgsStr || imgsStr === '') return []
  var result = imgsStr.split(',')
  if (result[result.length - 1] === '') result = result.splice(result.length - 1, 1)
  return result
}