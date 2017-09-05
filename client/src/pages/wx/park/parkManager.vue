<template>
  <div :class='s.wrap'>
    <div 
      :class='s.add'
      @click='onAdd'>
      <i class="iconfont icon-add"></i>
      <div>绑定小区</div>
    </div>
    <div :class='s.row'>
      <div v-for='(manager, index) in parkManagers'
        @click='onXQ(manager)'
        :class='s.item' 
        :style='{backgroundColor: manager.Bind < 1 ? "#ddd" : "#fff", color: manager.Bind < 1 ? "#999" : "#565656"}'>
        <div v-if='xqs[index]' :class='s.name' >{{xqs[index].Name}}</div>
        <div :class='s.tip' v-if='manager.Bind === 1 ' :style='{color: "rgb(25, 163, 24)"}'>通过审核</div>
        <div :class='s.tip' v-if='manager.Bind === 0 ' :style='{color: "#999"}'>等待审核</div>
        <div :class='s.tip' v-if='manager.Bind === -1 ' :style='{color: "#999"}'>未通过审核</div>
      </div>
    </div>
    <y-dialog :visible.sync='isUpdatingPark'>
      <update-free-park :xqId='updatingXQID'></update-free-park>
    </y-dialog>
  </div>
</template>

<script>
import fetchpm from '@/fetchpm'
import YDialog from '@/components/mobile/YDialog'
import UpdateFreePark from '@/components/mobile/UpdateFreePark'
export default {
  components: { YDialog, UpdateFreePark },
  data () {
    return {
      openid: '',
      parkManagers: [],
      xqs: [],
      isUpdatingPark: false,
      updatingXQID: ''
    }
  },
  mounted () {
    var WXUser = JSON.parse(sessionStorage.getItem('WXUser') || '{"openid": ""}')
    this.openid = WXUser.openid
    this.fetchParkManagers(this.openid)
  },
  methods: {
    onXQ (manager) {
      console.info('isUpdatingPark', manager)
      if (manager.Bind !== 1) return
      this.updatingXQID = manager.XQID
      this.isUpdatingPark = true
    },
    onAdd () {
      this.$router.push({path: '/wx/park/manager/bind/xq'})
    },
    fetchParkManagers (openid) {
      fetchpm(this, false, '/open/park/manager/query/openid', {
        method: 'POST',
        body: {openid: openid}
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('fetchParkManagers', body)
        if (body.error === 0) this.parkManagers = body.data || []
        this.parkManagers.sort((a, b) => {
          return b.Bind - a.Bind
        })
        let ids = this.parkManagers.map(item => {
          return item.XQID
        })
      this.fetchXQs(ids)
      })
    },
    fetchXQs (ids){
      fetchpm(this, false, '/open/xq/ids', {
        method: 'POST',
        body:{values: ids}
      }).then(resp => {
        return resp.json()
      }).then(body => {
        if (body.error === 0){
          this.xqs = body.data || []
        }
      })
    }
  }
}
</script>

<style lang="less" module='s'>
.wrap{
  background-color: #20a0ff;
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  color: #555;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}
.add {
  border: solid 2px #f1f3f6;
  padding: 20px;
  border-radius: 5px;
  color: #fff;
  text-align: center;
  i{
    font-size: 40px;
  }
}
.row{
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  .item{
    flex: 1;
    padding: 20px 10px;
    margin: 15px;
    border-radius: 5px;
    text-align: center;
    font-size: 25px;
    .name{
      overflow: hidden;
      text-overflow:ellipsis;
      white-space: nowrap;
    }
    .tip{
      font-size: 15px;
    }
  }
}
</style>