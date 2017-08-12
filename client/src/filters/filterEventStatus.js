export default function (status) {
  switch (status) {
    case 0:
      return '居民提交'//居民提交
    case 1:
      return '已审核待处理'//已审核待处理
    case 2:
      return '已解决'//已解决
    case -1:
      return '已关闭'//已关闭
  }
}