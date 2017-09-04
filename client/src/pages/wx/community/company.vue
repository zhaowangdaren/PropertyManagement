<!-- 工程服务列表 -->
<template>
  <div>
    <action-bar :OPTIONS='headerOptions'></action-bar>
    <div :class="s.content">
      <div :class="s.up">
        <div :class="s.container">
          <div :class="s.logo">
            LOGO
          </div>
          <div :class="s.name">{{company.Name}}</div>
          <div :class='s.charger'><span>企业负责人：</span>{{company.Charger}}</div>
        </div>
      </div>
      <div :class="s.down">
        <div><span>联系电话</span>{{company.Tel}}</div>
        <div><span>加盟时间</span>{{company.JoinTime}}</div>
        <div><span>服务内容</span>{{company.MainContent}}</div>
        <div><span>服务时间</span>{{company.WorkSchedule}}</div>
      </div>
      

    </div>
  </div>
</template>

<script>
import ActionBar from '@/components/mobile/ActionBar'
import ListMenu from '@/components/mobile/ListMenu'
export default {
  components: { ListMenu, ActionBar },
  data () {
    return {
      headerOptions: {
        leftBtns: [],
        title: '公司详情',
        rightBtns: []
      },
      company: {
        Logo: '',
        Name: '马鞍山市雨山区洁美家政服务部',
        Charger: '左军',
        Tel: '96345',
        JoinTime: '2015-05-21',
        MainContent: '家政/保洁',
        WorkSchedule: '8:00-18:00'
      }
    }
  },
  methods: {
    
  }
}
</script>

<style lang="less" module="s">
.content{
  position: fixed;
  top: 80px;
  bottom: 0;
  left: 0;
  right: 0;
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
}
.up{
  width: 100%;
  height: 250px;
  background-image: url('../../../res/images/yw.jpg');
  background-repeat: no-repeat;
  background-size: 100%;
  position: relative;
  color: #fff;

}
.container{
  display: block;
  background-color: rgba(0, 0, 0, 0.5);
  height: 100%;
  width: 100%;
  position: absolute;
  top: 0;
  display: flex;
  justify-content: center;
  align-items: center;
}
.down{
  width: 100%;
  flex: 1;
  background-color: #fff;
}
.logo{
  display: none;
  background-color: rgb(143, 206, 105);
  color: #fff;
  font-weight: bold;
  width: 100px;
  height: 100px;
  border-radius: 50%;
  line-height: 100px;
  text-align: center;
  z-index: 2;
}
.name{
  font-size: 30px;
  margin: 10px; 
  z-index: 2;
}
.charger{
  font-size: 20px;
  span{
    color: #ccc;
  }
}
</style>