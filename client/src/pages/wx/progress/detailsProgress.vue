<template>
  <div>
    <action-bar :OPTIONS='headerOptions'></action-bar>
    <div :class='s.content'>
      <div :class='s.up'>
        <div :class='s.warn' v-text='warn'></div>
        <div :class='s.item'>
          <div :class='s.key'>进度：</div>
          <div :class='s.value'>{{event.Status | filterEventStatus }}</div>
          <el-button 
            v-if='event.Status == 0'
            type='danger' 
            :class='s.revoke'
            @click='onRevoke(event.Index)'>撤销事件</el-button>
        </div>
      </div>
      <div :class='s.event' v-for='handle in eventHandles'>
        <div :class='s.left'>
          <div :class='s.identity'>{{handle.AuthorCategory | filterUserIdentity}}</div>
          <div><span :class='s.key'>提交时间:</span><span>{{handle.Time | filterTime}}</span></div>
          <div><span :class='s.key'>提交内容:</span><span v-text='handle.HandleInfo'></span></div>
          <div><span :class='s.key' v-if='handle.Imgs.length > 0'>提交图片:</span><span class='iconfont icon-pic' @click='onShowImgs(handle.Imgs)'>点击查看图片</span></div>
        </div>
      </div>
      
      <div :class='s.event'>
        <div :class='s.left'>
          <div :class='s.identity'>居民</div>
          <div><span :class='s.key'>提交时间:</span><span>{{event.Time | filterTime}}</span></div>
          <div><span :class='s.key'>提交内容:</span><span v-text='event.Content'></span></div>
          <div><span :class='s.key' v-if='event.Imgs.length > 0'>提交图片:</span><span class='iconfont icon-pic' @click='onShowImgs(event.Imgs)'>点击查看图片</span></div>
          <div><span :class='s.key'>联系电话:</span><span v-text='event.Tel'></span></div>
        </div>
      </div>
    </div>
    <el-dialog
      title='图片'
      size='large'
      :visible.sync='showImgs'>
        <imgs :imgs='showingImgs' v-if='showImgs'></imgs>
      </el-dialog>
  </div>
</template>

<script>
import ActionBar from '@/components/mobile/ActionBar'
import BDialog from '@/components/dialog/BasicDialog'
import Imgs from '@/components/dialog/Imgs'
import filterEventStatus from '@/filters/filterEventStatus'
import filterTime from '@/filters/filterTime'
import fetchpm from '@/fetchpm'
import filterUserIdentity from '@/filters/filterUserIdentity'
export default {
  components: { ActionBar, Imgs },
  filters: { filterEventStatus, filterTime, filterUserIdentity },
  data () {
    return {
      headerOptions: {
        title: '查看进度详情'
      },
      eventIndex: '',
      eventStatus: 0,
      event: {
        Status: 0,
        Imgs:''
      },
      eventHandles: [],
      showImgs: false,
      showingImgs: ''
    }
  },
  mounted () {
    this.eventIndex = this.$route.query.index
    this.eventStatus = this.$route.query.status
    this.fetchEvent()
    this.fetchEventHandle()
  },
  methods: {
    onShowImgs (imgs) {
      this.showingImgs = imgs
      this.showImgs = true
      console.info('onShowImgs',this.showingImgs)
    },
    onRevoke () {
      this.event.Status = -1
      fetchpm(this, false,'/open/event/update', {
        method: 'POST',
        body: this.event
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('onRevoke', body)
        if (body.error === 0 ) this.warn = '撤销成功'
        else this.warn = body.data
      })
    },
    fetchEvent () {
      fetchpm(this, false, '/open/event/id/' + this.eventIndex, {
        method: 'POST'
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('fetchEvent', body)
        if (body.error === 0 ) {
          this.event = body.data[0] || {}
        }
      })
    },
    fetchEventHandle () {
      fetchpm(this, false, '/open/eventHandle/kvs', {
        method: 'POST',
        body: {index: this.eventIndex}
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('fetchEventHandle', body)
        if (body.error === 0){
          this.eventHandles = body.data || []
        }
      })
    }
  }
}
</script>

<style lang="less" module='s'>
.content{
  margin-top: 80px;
  font-size: 25px;
  .up{
    background-color: #fff;
    box-shadow: 0 1px 2px 0 rgba(0, 0, 0, 0.3);
    .item{
      padding: 10px;
      display: flex;
      align-items: center;
      .key{
        color: #555;
      }
      .value{
        flex: 1;
      }
      .revoke {

      }
    }
  }
  .event{
    margin: 10px;
    background-color: #fff;
    border-radius: 2px;
    padding: 10px;
    display: flex;
    justify-content: center;
    align-items: center;
    .left{
      flex: 1;
      .key{
        color: #555;
      }
    }
    .btn{
      background-color: #20a0ff;
      border-radius: 5px;
      padding: 10px;
      color: #fff;
      font-size: 12px;
    }
  }
}
</style>