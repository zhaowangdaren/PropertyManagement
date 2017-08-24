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
      orient: 'vertical',
      x: 'left',
      data: data.data
    },
    series: [
      {
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