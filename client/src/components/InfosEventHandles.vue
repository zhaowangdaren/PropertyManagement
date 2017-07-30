<template>
  <div :class='s.wrap'>
    <div :class='s.street'>Street</div>
    <event-handles></event-handles>
  </div>
</template>

<script>
  import filterEventStatus from '@/filters/filterEventStatus'
  import filterEventLevel from '@/filters/filterEventLevel'
  import EventHandles from '@/components/table/EventHandles'
  import fetchpm from '@/fetchpm'
  export default {
    components: { EventHandles },
    filters: {filterEventStatus, filterEventLevel},
    data () {
      return {
        host: '//localhost:3000',
        eventHandle: {
          eventLevel:'',
          eventType: '',
          startTime:'',
          endTime:''
        },
        searchResult:[{}],
        eventHandles: [],
        eventLevels:["全部", "特急", "加急", "急"],
        eventTypes: ["全部", "type1", "type2", 'type3'],
        eventStatus: ["全部", 'status1', 'status2', 'status3', 'status4'],
        communities: [],
        xqs:[]
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
        fetchpm(this, true, '/pm/event', {
          method: 'POST',
          body: {index: index}
        }).then(resp => {
          return resp.json()
        }).then(body => {
          console.info('fetchEvents', body)
          this.searchResult = body.data
        })
      },
      toDetails (event) {
        console.info('toDetails', event.Index)
        this.$router.push({name: 'eventDetail', params:{index: event.Index}})
        // this.$router.push({path:'/street/detail/' + event.Index})
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