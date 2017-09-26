export default function (handleType) {
  switch (handleType) {
    case 1 :
      return '街道办推送给物业公司'
    case 2 :
      return '房管中心约谈物业公司'
    case 3 :
      return '询问'
    case 4:
      return '回复'
    case 5:
      return '物业公司处理'
    case 6:
      return '用户评价'
    case 7:
      return '审核事件等级'
    case 8:
      return '街道请求关闭投诉'
    case 9:
      return '政府同意关闭'
    case 10:
      return '街道推送给政府'
    case 11:
      return '政府推送给法官'
    case 12:
      return '法官询问物业公司'
    case 13:
      return '法官询问投诉人'
    case 14:
      return '物业回复法官'
    case 15:
      return '居民回复法官'
    case 16:
      return '法官受理'
    default:
      return '提交'
  }
}
