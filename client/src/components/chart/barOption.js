export default function (data) {
  return {
    title: {
      text: data.title,
      left: 'center'
    },
    tooltip: {
      trigger: 'axis',
      axisPointer: {
          type: 'shadow'
      }
    },
    grid: {
      left: '3%',
      right: '4%',
      bottom: '3%',
      containLabel: true
    },
    xAxis: {
      type: 'value',
      boundaryGap: [0, 0.01]
    },
    yAxis: {
      type: 'category',
      data: data.data.y
    },
    series: [
      {
        type: 'bar',
        barWidth: 10,
        data: data.data.x
      }
    ]
  }
}