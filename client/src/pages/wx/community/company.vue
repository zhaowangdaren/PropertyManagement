<!-- 工程服务列表 -->
<template>
  <div :class="s.wrap">
    <action-bar :OPTIONS='headerOptions'></action-bar>
    <div :class="s.content">
      <div :class="s.up">
        <div :class="s.container">
          <div :class="s.logo">
            LOGO
          </div>
          <div :class="s.name">{{company.Name}}</div>
        </div>
      </div>
      <div :class="s.down">
        <div :class='s.item'><span>企业负责人：</span>{{company.Charger}}</div>
        <div :class='s.item'><span>联系电话：</span><a :href="'tel:' + company.Tel">{{company.Tel}}</a></div>
        <div :class='s.item'><span>加盟时间：</span>{{company.JoinTime}}</div>
        <div :class='s.item'><span>服务内容：</span>{{company.MainContent}}</div>
        <div :class='s.item'><span>服务时间：</span>{{company.WorkSchedule}}</div>
        <div :class="s.assess" @click='showAssess = true'>服务评价</div>
      </div>
    </div>
    <div :class="s.bottom">
      <div :class="s.callBtn"><a :href="'tel:' + company.Tel">呼 叫</a></div>
    </div>
    <y-dialog :visible.sync='showAssess'>
      <assess @close='showAssess = false'></assess>
    </y-dialog>
  </div>
</template>

<script>
import ActionBar from '@/components/mobile/ActionBar'
import ListMenu from '@/components/mobile/ListMenu'
import YDialog from '@/components/mobile/YDialog'
import Assess from '@/components/mobile/Assess'
export default {
  components: { ListMenu, ActionBar, YDialog, Assess },
  data () {
    return {
      headerOptions: {
        leftBtns: [{text: '上一步', event: this.onBack}],
        title: '详情',
        rightBtns: []
      },
      company: {
        Logo: '',
        Name: '名称名称名称名称名称名称名称吗',
        Charger: '名字名字',
        Tel: '96345',
        JoinTime: '2015-05-21',
        MainContent: '内容内容内容',
        WorkSchedule: '8:00-18:00'
      },
      showAssess: false
    }
  },
  methods: {
    onBack () {
      this.$router.go(-1)
    }
  }
}
</script>

<style lang="less" module="s">
.wrap{
  height: 100%;
  background-color: #fff;
}
.content{
  margin-top: 80px;
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
}
.up{
  width: 100%;
  // height: 250px;
  background-image: url('../../../res/images/upbg5.jpg');
  background-repeat: no-repeat;
  background-size: cover;
  position: relative;
  color: #fff;

}
.container{
  display: block;
  background-color: rgba(72, 87, 106, 0.3);
  height: 100%;
  width: 100%;
  // position: absolute;
  top: 0;
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
}
.name{
  font-size: 30px;
  margin: 10px; 
  z-index: 2;
  padding: 100px 10px;
}
.down{
  width: 100%;
  flex: 1;
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
.item{
  font-size: 22px;
  margin: 20px;
  padding: 10px 10px;
  background-color: #f1f3f6;
  span{
    color: #999;
  }
}
.callBtn{
  position: fixed;
  bottom: 10px;
  left: 50%;
  transform: translateX(-50%);
  padding: 10px 20px;
  font-size: 20px;
  font-weight: 500;
  background-color: rgb(25, 163, 24);
  border-radius: 100px;
  box-shadow: 0 2px 5px rgba(0,0,0,.25);
  a{
    color: #fff;
  }
}
.assess{
  background-color: #1f2d3d;
  padding: 10px;
  color: #f1f3f6;
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 22px;
  margin: 50px;
}
</style>