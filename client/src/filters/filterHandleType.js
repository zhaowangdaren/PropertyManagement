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
    default:
      return '提交'
  }
}