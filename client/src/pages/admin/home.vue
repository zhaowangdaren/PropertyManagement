<template>
  <div>
    <h1 :class='s.title'>数据总览</h1>
    <div :class='s.datas'>
      <div :class='s.item'>
        <div :class='s.spec'><span>{{todayEvents.length}}</span>条</div>
        <div :class='s.text'>今日新增投诉</div>
      </div>
      <div :class='s.item'>
        <div :class='s.spec'><span>{{complaintXQNum}}</span>条</div>
        <div :class='s.text'>今日被投诉小区数</div>
      </div>
      <div :class='s.item'>
        <div :class='s.spec'><span>{{todayEventHandles.length}}</span>条</div>
        <div :class='s.text'>今日新增事件处理</div>
      </div>
    </div>
    <div :class='s.datas'>
      <div :class='s.item'>
        <div :class='s.value'><span>{{streetUserNum}}</span>条</div>
        <div :class='s.text'>街道用户数</div>
      </div>
      <div :class='s.item'>
        <div :class='s.value'><span>{{govNum}}</span>条</div>
        <div :class='s.text'>区政府用户数</div>
      </div>
      <div :class='s.item'>
        <div :class='s.value'><span>{{pmsNum}}</span>条</div>
        <div :class='s.text'>物业信息</div>
      </div>
      <div :class='s.item'>
        <div :class='s.value'><span>{{housesNum}}</span>条</div>
        <div :class='s.text'>建筑信息</div>
      </div>
    </div>
    <div :class='s.datas'>
      <div :class='s.item'>
        <div :class='s.value'><span>{{eventsNum}}</span>条</div>
        <div :class='s.text'>投诉记录数</div>
      </div>
      <div :class='s.item'>
        <div :class='s.value'><span>{{streetsNum}}</span>条</div>
        <div :class='s.text'>街道总数</div>
      </div>
      <div :class='s.item'>
        <div :class='s.value'><span>{{communitiesNum}}</span>条</div>
        <div :class='s.text'>社区总数</div>
      </div>
      <div :class='s.item'>
        <div :class='s.value'><span>{{xqsNum}}</span>条</div>
        <div :class='s.text'>小区总数</div>
      </div>
    </div>
  </div>
</template>

<script>
import fetchpm from '@/fetchpm'
export default {
  data () {
    return {
      todayEvents: [],
      complaintXQNum: 0,
      todayEventHandles: [],
      govNum: 0,
      streetUserNum: 0,
      pmsNum: 0,
      housesNum: 0,
      eventsNum: 0,
      communitiesNum: 0,
      streetsNum: 0,
      xqsNum: 0
    }
  },
  mounted () {
    this.fetchTodayEvents()
    this.fetchUserNum(2).then(num => {
      this.govNum = num
    })
    this.fetchUserNum(3).then(num => {
      this.streetUserNum = num
    })
    this.fetchPMsNum()
    this.fetchHousesNum()
    this.fetchEventsNum()
    this.fetchCommunitiesNum()
    this.fetchStreetsNum()
    this.fetchXQsNum()
  },
  methods: {
    fetchPMsNum () {
      fetchpm(this, true, '/pm/pms/num', {
        method: 'GET'
      }).then(resp => {
        return resp.json()
      }).then(body => {
        if (body.error === 0 ) this.pmsNum = body.data || 0
      })
    },
    fetchHousesNum () {
      fetchpm(this, true, '/pm/houses/num', {
        method: 'GET'
      }).then(resp => {
        return resp.json()
      }).then(body => {
        if (body.error === 0 ) this.housesNum = body.data || 0
      })
    },
    fetchEventsNum () {
      fetchpm(this, true, '/pm/events/num', {
        method: 'GET'
      }).then(resp => {
        return resp.json()
      }).then(body => {
        if (body.error === 0 ) this.eventsNum = body.data || 0
      })
    },
    fetchCommunitiesNum () {
      fetchpm(this, true, '/pm/communities/num', {
        method: 'GET'
      }).then(resp => {
        return resp.json()
      }).then(body => {
        if (body.error === 0 ) this.communitiesNum = body.data || 0
      })
    },
    fetchStreetsNum () {
      fetchpm(this, true, '/pm/streets/num', {
        method: 'GET'
      }).then(resp => {
        return resp.json()
      }).then(body => {
        if (body.error === 0 ) this.streetsNum = body.data || 0
      })
    },
    fetchXQsNum () {
      fetchpm(this, true, '/pm/xqs/num', {
        method: 'GET'
      }).then(resp => {
        return resp.json()
      }).then(body => {
        if (body.error === 0 ) this.xqsNum = body.data || 0
      })
    },
    fetchTodayEvents () {
      fetchpm(this, true, '/pm/event/today/events', {
        method: 'GET'
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('fetchTodayEvents', body)
        if (body.error === 0) this.todayEvents = body.data || []
        this.complaintXQNum = this.countTodayCompXQNum(this.todayEvents)
      })
    },
    countTodayCompXQNum (todayEvents) {
      var num = 0
      todayEvents.sort((a, b) => {
        return a.XQID > b.XQID ? 1 : -1
      })
      var tempXQID = ''
      for (var i = 0; i < todayEvents.length; i++) {
        if (todayEvents[i].XQID !== tempXQID) {
          num = num + 1
          tempXQID = todayEvents[i].XQID
        }
      }
      return num
    },
    fetchTodayEventHandles () {
      fetchpm(this, true, '/pm/eventHandle/today/handles', {
        method: 'GET'
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('fetchTodayEventHandles', body)
        if(body.error === 0) this.todayEventHandles = body.data || []
      })
    },
    fetchUserNum (type) {
      return fetchpm(this, true, '/pm/users/num/' + type, {
        method: 'GET'
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('fetchUserNum',body)
        if(body.error === 0) {
          return body.data || 0
        }
      })
    }
  }
}
</script>
<style lang="less" module='s'>
.title{
  font-size: 15px;
  color: #555;
}
.datas{
  display: flex;
  background-color: #fff;
  padding: 5px;
}
.item{
  flex: 1;
  margin: 5px;
  border-radius: 5px;
  padding: 20px 20px;
  font-size: 18px;
  text-align: center;
  background-color: #eee;
  color: #666;
  .value{
    span{
      color: #212138;
      font-size: 25px;
    }
    font-weight: bold;
  }
  .spec{
    span{
      color: #e82d2d;
      font-size: 25px;
    }
    font-weight: bold;
  }
}


</style>