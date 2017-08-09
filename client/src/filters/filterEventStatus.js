export default function (status) {
  switch (status) {
    case 0:
      return 'Submit'//居民提交
    case 1:
      return 'Aduit'//已审核待处理
    case 2:
      return 'Handle'//已解决
    case 3:
      return 'Closed'//已关闭
  }
}