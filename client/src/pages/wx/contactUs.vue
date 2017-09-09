<!-- 工程服务列表 -->
<template>
  <div :class="s.wrap">
    <action-bar :OPTIONS='headerOptions'></action-bar>
    <div :class="s.content">
      <div :class="s.up">
        <div :class="s.container">
          <div :class="s.logo">
            <img src="~@/res/images/logo.png">
          </div>
          <div :class="s.name">{{company.name}}</div>
        </div>
      </div>
      <div :class="s.down">
        <div :class='s.item'><span>单位地址：</span>{{company.address}}</div>
        <div :class='s.item'><span>联 系 人：</span>{{company.man}}</div>
        <div :class='s.item'>
          <span>联系方式：</span>
          <div>
            <a :href="'tel:' + company.Tel">{{company.mobile}}(手机)</a>
            &nbsp;&nbsp;&nbsp;&nbsp;
            <a :href="'tel:' + company.Tel">{{company.tel}}(固话)</a>
          </div>
        </div>
        <div :class='s.item'><span>电子邮件：</span>{{company.email}}</div>
      </div>
    </div>
    <div :class="s.bottom">
      <div :class="s.callBtn"><a :href="'tel:' + company.tel">呼 叫</a></div>
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
        title: '联系我们',
      },
      company: {
        Logo: '',
        name: '马鞍山市花山区住房管理中心',
        address: '花山区花山路152号科技大厦',
        man: '司先生',
        tel: '0555-2373620',
        mobile: '15255528528',
        email: 'maszfglzx@163.com'
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
  background-image: url('../../res/images/upbg5.jpg');
  background-repeat: no-repeat;
  background-size: cover;
  position: relative;
  color: #fff;
}
.container{
  display: block;
  padding-top: 20px;
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
  padding: 20px 10px;
}
.down{
  width: 100%;
  flex: 1;
}
.logo{
  // display: none;
  // background-color: rgb(143, 206, 105);
  color: #fff;
  font-weight: bold;
  width: 100px;
  height: 100px;
  border-radius: 50%;
  line-height: 100px;
  text-align: center;
  z-index: 2;
  img{
    width: 100%;
  }
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