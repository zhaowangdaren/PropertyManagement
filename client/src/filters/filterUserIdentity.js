export default function (type) {
  switch (type) {
    case 1:
      return '系统管理员'
    case 2:
      return '政府主管单位'
    case 3:
      return '街道办公人员'
    case 4:
      return '物业公司'
    default:
      return '居民提交'
  }
}