<template>
  <div :class='s.wrap'>
    <action-bar :OPTIONS='headerOptions'></action-bar>
    <div :class='s.body'>
      <div :class='s.header'>
        <span :class='s.name' v-text='xq.Name'></span>
        XQ
      </div>
      <div :class='s.item'>
        <div :class='s.title'>Event Type:</div>
        <popup-picker
          :class='s.value'
          :data='eventTypes'
          :value='selectedEventType.value'
          @change='onEventTypePick'
          @onScroll='onEventTypePick'>
          <div 
            :class='s.typeValue'>
            <span
              v-if='selectedEventType.index > 0' 
              v-text='eventTypes[0][selectedEventType.index].name'
              ></span>
            <span v-else :class='s.typeValueEmpty'>--请选择--</span>
            <i class="iconfont icon-unfold"></i>
          </div>
        </popup-picker>
      </div>
      <div :class='s.item'>
        <div :class='s.title'>联系电话:</div>
        <input 
          :class='s.value'
          type="number">
      </div>
      <div :class='s.contentWrap'>
        <div :class='s.title'>Content</div>
        <div :class='s.value'>
          <textarea 
            rows='10'
            placeholder="请输入内容"
            v-model="event.Content" 
            ></textarea>
        </div>
      </div>
      <div :class='s.btn'>下一步</div>
    </div>
  </div>
</template>

<script>
import ActionBar from '@/components/mobile/ActionBar'
import PopupPicker from '@/components/popup-picker'
export default{
  components: { ActionBar, PopupPicker },
  data () {
    return {
      headerOptions: {
        leftBtns: ['上一步'],
        title: 'Fill in Event',
        rightBtns: ['下一步']
      },
      xq: {
        ID: '',
        Name: ''
      },
      event: {
        Content: ''
      },
      eventTypes: [[
        {name:'Type1', value:'0'},
        {name:'Type2', value:'1'},
        {name:'Type3', value:'2'},
        {name:'Type4', value:'3'},
        {name:'Type5', value:'4'},
        {name:'Type6', value:'5'}
      ]],
      selectedEventType:{index: -1, value:['']},
      host:'http://10.176.118.61:3000'
    }
  },
  mounted () {
    this.xq.CommunityID = this.$route.query.XQID
    this.xq.Name = this.$route.query.XQName
  },
  methods: {
    onEventTypePick (value) {
      console.info('onEventTypePick', value)
      this.selectedEventType.value[0] = value
      this.selectedEventType.index = parseInt(value)
    }
  }
}
</script>

<style lang="less" module='s'>
.wrap{
  .body{
    margin-top: 80px;
    .header{
      font-size: 30px;
      margin: 10px;
      font-weight: bold;
      color: #999;
      .name{
        color: #000;
      }
    }
    .item{
      margin: 10px auto;
      font-size: 25px;
      display: flex;
      padding: 0px 10px;
      background-color: #fff;
      .title{
        padding: 15px 0;
        color: #999;
      }
      .value{
        padding: 15px 10px;
        border-bottom: solid 2px #20a0ff;
        border-top: solid 0px;
        border-left: solid 0;
        border-right: solid 0;
        flex: 1;
        font-size: 30px;
        .typeValue{
          display: flex;
          align-items: center;
          span{
            flex: 1;
          }
          .typeValueEmpty{
            color: #999;
          }
        }
      }
    }
    .contentWrap{
      font-size:25px;
      margin-top: 20px;
      .title{
        margin: 10px;
        font-size: 30px;
      }
      .value{
        width: 90%;
        margin: 0 auto;
        textarea{
          width: 100%;
          min-width: 50px;
        }
      }
    }
    .btn{
      background-color: #20a0ff;
      font-size: 30px;
      color: #fff;
      text-align: center;
      padding: 15px 0;
      margin: 20px;
    }
  }
}
</style>