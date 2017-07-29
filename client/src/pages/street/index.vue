<template>
  <div :class='s.wrap'>
    <div :class='s.street'>Street</div>
    <to-be-done></to-be-done>
  </div>
</template>

<script>
  import filterEventStatus from '@/filters/filterEventStatus'
  import filterEventLevel from '@/filters/filterEventLevel'
  import ToBeDone from '@/components/table/ToBeDone'
  export default {
    components: {ToBeDone},
    filters: {filterEventStatus, filterEventLevel},
    data () {
      return {
        host: '//localhost:3000',
        events: []
      }
    },
    mounted () {
      this.fetchEvents('')
    },
    methods: {
      fetchEvents (index) {
        if (!index ) {
          console.info('index', index)
        }
        fetch(this.host + '/event', {
          method: 'POST',
          body: JSON.stringify({index: index})
        }).then(resp => {
          return resp.json()
        }).then(body => {
          console.info('fetchEvents', body)
          this.events = body.data
        })
      },
      toDetails (event) {
        console.info('toDetails', event.Index)
        // this.$router.push({path:'/street/detail/', params:{index: event.Index}})
        this.$router.push({path:'/street/detail/' + event.Index})
      }
    }
  }
</script>
<style lang="less" module='s'>
.wrap{
  margin: 10px;
  width: 100%;
  .street{
    background-color: #ddd;
    width: 100%;
    font-size: 25px;
    padding: 5px;
  }
  
  
}
</style>