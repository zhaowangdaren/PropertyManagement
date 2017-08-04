export default function (data) {
  return {
      tooltip : {
          formatter: "{a} <br/>{b} : {c}%"
      },
      toolbox: {
          feature: {
              restore: {},
              saveAsImage: {}
          }
      },
      series: [
          {
              name: '业务指标',
              type: 'gauge',
              detail: {formatter:'{value}%'},
              data: [data]
          }
      ]
  }
}
