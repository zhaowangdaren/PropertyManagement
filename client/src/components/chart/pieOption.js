export default function function_name(data) {
  return {
    title: {
      text: data.title,
      left: 'center'
    },
    tooltip: {
      trigger: 'item',
      formatter: "{b}: {c} ({d}%)"
    },
    legend: {
      type: 'scroll',
      orient: 'vertical',
      left: 10,
      top: 20,
      bottom: 20,
      data: data.data,
      width: '10'
    },
    series: [
      {
        zlevel: 5,
        type:'pie',
        radius: ['50%', '70%'],
        avoidLabelOverlap: false,
        label: {
            normal: {
                show: false,
                position: 'center'
            },
            emphasis: {
                show: true,
                textStyle: {
                    fontSize: '30',
                    fontWeight: 'bold'
                }
            }
        },
        labelLine: {
            normal: {
                show: false
            }
        },
        data: data.data
      }
    ]
  }
}