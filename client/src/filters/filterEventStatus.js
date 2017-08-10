export default function (status) {
  switch (status) {
    case 0:
      return 'Submit'//居民提交
    case 1:
      return 'Review'//待审核
    case 2:
      return 'Handle'//待处理
    case 3:
      return 'Finish'//已处理
    case 4:
      return 'Close'//已关闭
  }
}