<template>
  <div :class='s.wrap'>
    <div :class='s.searchWrap'>
      <div :class='s.searchIcon'><i class="iconfont icon-ic_serach"></i></div>
      <div :class='s.inputWrap'>
        <input type="text" name="" placeholder="输入小区名" 
          v-model='xqName'
          @focus='showSearchResults = true'>
      </div>
      <div :class='s.closeIcon' v-if='xqName.length > 0' @click='onClearXQName'><i class="iconfont icon-close"></i></div>
    </div>
    <div :class='s.searchResults' v-if='showSearchResults && xqs.length > 0'>
      <div v-for='xq in xqs' :class='s.item' @click='onXQ(xq)'>
        <div>{{xq.Name}}</div>
        <div :class='s.address'>{{xq.Address}}</div>
      </div>
    </div>
    <div :class='s.parkIcon'>
      <i class="iconfont icon-tingchewei"></i>
    </div>
    <park-free :class='s.parkFree' 
      v-if='selectedXQ.ID.length > 0' 
      :xqID='selectedXQ.ID'
      :xqName='selectedXQ.Name'></park-free>
    <div :class='s.manager' @click='onManager'>我是管理员</div>
  </div>
</template>

<script>
import fetchpm from '@/fetchpm'
import ParkFree from '@/components/mobile/ParkFree'
export default {
  components: { ParkFree },
  data () {
    return {
      xqName: '',
      showSearchResults: false,
      xqs: [],
      selectedXQ: {
        Name: '小区名',
        ID: ''
      }
    }
  },
  watch: {
    xqName: function (val) {
      if (val.length >= 2) this.fetchXQs(val)
      else this.xqs = []
    }
  },
  mounted () {
    this.fetchOpenID()
  },
  methods: {
    onCheck () {

    },
    onManager () {
      this.$router.push({path:'/wx/park/manager'})
    },
    onClearXQName () {
      this.xqName = ''
    },
    fetchXQs (name) {
      fetchpm(this, false, '/open/xq/name/' + name, {
        method: 'GET'
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('fetchXQs', body)
        if (body.error === 0) this.xqs = body.data || []
        else this.xqs = []
      })
    },
    onXQ (xq) {
      this.selectedXQ = xq
      this.showSearchResults = false
    },
    fetchOpenID () {
      fetchpm(this, false, '/open/wx/openid/' + this.$route.query.code, {
        method: 'POST'
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('fetchOpenID', body)
        if (body.error === 0) {
          sessionStorage.setItem('WXUser', JSON.stringify(body.data))
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
}
.searchWrap{
  position: relative;
  display: flex;
  margin: 20px 20px 0 20px;
  background-color: #fff;
  justify-content: center;
  align-content: center;
  font-size: 23px;
  height: 50px;
  border-radius: 25px;
  .searchIcon, .closeIcon{
    color: #999;
    height: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
    border: solid 1px transparent;
    i{
      font-size: 23px;
      margin: 5px;
    }
  }
  .inputWrap{
    flex: 1;
    display: flex;
    align-items: center;
    input{
      border: none;
      width: 96%;
      box-sizing: border-box;
    }
  }
}
.searchResults{
  background-color: #fff;
  color: #565656;
  margin: 0 auto;
  width: 86%;
  font-size: 22px;
  .item{
    padding: 10px;
    border-bottom: solid 1px #ccc;
    .address{
      font-size: 20px;
      color: #999;
      width: 100%;
      overflow: hidden;
      text-overflow:ellipsis;
      white-space: nowrap;
    }
  }
}
.parkIcon{
  color: #fff;
  position: fixed;
  top: 50%;
  left: 0;
  right: 0;
  text-align: center;
  i{
    font-size: 100px;
  }
}
.parkFree{
  position: absolute;
  top: 30%;
  width: 80%;
  left: 10%;
}
.manager{
  position: fixed;
  bottom: 10px;
  left: 0;
  right: 0;
  text-align: center;
  color: #fff;
  font-size: 22px;
  &:active {
    color: rgba(255,255,255, 0.5);
  }
}
</style>