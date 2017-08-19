<template>
  <div :class='s.wrap'>
    <div v-for='(bg, index) in bgs' :class='s.bgImage'>
      <transition name='fade'>
        <div :style='{backgroundImage: "url(" + bg + ")"}' 
            v-if='bgIndex == index' :class='s.item'>
        </div>
      </transition>
    </div>
    
    <div :class='s.content'>
      <div :class='s.titleWrap'>
        <img src="~@/res/images/logo.png">
        <img src="~@/res/images/title.png">
      </div>
      <div :class='s.body'>
        <div :class='s.left'>
          <div :class='s.adminWrap' @click='toAdmin'>
            <div>
              管理员
            </div>
            <div :class='s.hide'>
              <span :class='s.yellow'>管理员</span>后台统一管理
            </div>
          </div>
          <div :class='s.streetWrap' @click='toStreet'>
            街道工作人员
            <div :class='s.hide'>
              <span :class='s.yellow'>街道工作人员</span>居民纠纷处理
            </div>
          </div>
        </div>
        <div :class='s.right' @click='toGov'>
          <div :class='s.gov'>
            花山区住房管理中心工作人员
            <div :class='s.hide'>
              <span :class='s.yellow'>花山区住房管理<br/>中心工作人员</span>
              <br/>纠纷、法院调解、推送政府公告
            </div>
          </div>
          
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  export default {
    data () {
      return {
        bgs:[require('@/res/images/ysh1.jpg'),require('@/res/images/ysh2.jpg'),require('@/res/images/ysh3.jpg'),require('@/res/images/ysh4.jpg'),require('@/res/images/ysh5.jpg'),require('@/res/images/ysh6.jpg'),],
        bgIndex: 0,
        bgTimer: null
      }
    },
    mounted () {
      this.slideBg()
    },
    methods: {
      slideBg () {
        this.bgTimer = setInterval(() => {
          this.bgIndex = ( this.bgIndex + 1 ) % this.bgs.length
        }, 4000)
      },
      toAdmin () {
        this.$router.push({path:'/login', query: {title: '管理员', target: '/admin'}})
      },
      toStreet () {
        this.$router.push({path:'/login', query: {title:'街道工作人员', target:'/street'}})
      },
      toGov () {
        this.$router.push({path:'/login', query: {title:'区住房管理中心工作人员', target: '/gov'}})
      }
    }
  }
</script>

<style lang="less" module='s'>
  .wrap{
    background-color: #f0f0f0;
    background-repeat: no-repeat;
    background-size: cover;
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    .bgImage {
      position: absolute;
      height: 100%;
      width: 100%;
      
      .item{
        width: 100%;
        height: 100%;
      }
    }
    .content{
      z-index: 1;
      // background-color: #fff;
      width: 80%;
      .titleWrap{
        font-size: 50px;
        display: flex;
        justify-content: center;
        align-items: center;
        margin-bottom: 20px;
      }
      .body{
        display: flex;
        justify-content: center;
        align-items: center;
        width: 60%;
        height: 310px;
        font-size: 20px;
        margin: auto;
        color: #fff;
        .left{
          display: flex;
          flex: 1;
          height: 310px;
          margin: 5px;
          flex-direction: column;
          .adminWrap > .hide, .streetWrap > .hide {
            display: none;
            font-size: 15px;
            .yellow{
              color: yellow;
              font-size: 25px;
              font-style:italic;
            }
          }
          .adminWrap{
            // margin-top: 5px;
            flex: 1;
            // height: 50%;
            background-color: rgb(0,191,254);
            // height: 145px;
            padding: 5px;
            &:hover .hide{
              display: block;
            }
          }
          .streetWrap{
            margin-top: 5px;
            flex: 1;
            padding: 5px;
            background-color: rgb(158,152,74);
            // height: 145px;
            &:hover .hide{
              display: block;
            }
          }
        }
        .right{
          display: flex;
          flex: 1;
          height: 310px;
          margin: 5px;
          flex-direction: column;
          .gov{
            padding: 5px;
            flex: 1;
            background-color: rgb(27,185,127);
            position: relative;
            .hide{
              display: none;
              font-size: 15px;
              position: absolute;
              bottom: 10px;
              right: 10px;
              .yellow{
                color: yellow;
                font-size: 25px;
                font-style:italic;
              }
            }
            &:hover .hide{
              display: block;
            }
          }
        }
      }
    }
  }
</style>