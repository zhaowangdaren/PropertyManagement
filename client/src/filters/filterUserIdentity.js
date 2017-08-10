export default function (type) {
  switch (type) {
    case 1:
      return '管理员'
    case 2:
      return 'Gov'
    case 3:
      return 'Street'
    default:
      return '--'
  }
}