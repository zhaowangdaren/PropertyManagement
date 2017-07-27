export default function (status) {
  switch (status) {
    case 1:
      return 'Submit'//居民提交
    case 2:
      return 'Review'//待审核
    case 3:
      return 'Handle'//待处理
    case 4:
      return 'Finish'//已处理
  }
}