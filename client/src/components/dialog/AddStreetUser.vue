<template>
  <div>
  <!-- 新增街道街道 -->
    <basic-dialog 
      title='Title' 
      :onSave='onSave'
      :onCancel='onCancel'
      :warn='warn'
      >
      <div :class='s.item'>
        <div :class='s.red'>*</div>
        UserName
        <!-- 用户名 -->
        <input type="text" v-model='user.UserName' @focus='onFocus'>
      </div>
      <div :class='s.item'>
        <div :class='s.red'>*</div>
        RealName
        <!-- 真实姓名 -->
        <input type="text" v-model='user.RealName' @focus='onFocus'>
      </div>
      <div :class='s.item'>
        <div :class='s.red'>*</div>
        Street
        <!-- 授权码 -->
        <search-select v-model='user.street' :values='streets'/>
      </div>
      <div :class='s.item'>
        <div :class='s.red'>*</div>
        Password
        <!-- 描述 -->
        <input type="text" v-model='user.Password' @focus='onFocus'>
      </div>
      <div :class='s.item'>
        <div :class='s.red'>*</div>
        Tel
        <!-- 电话号码 -->
        <input type="text" v-model='user.Tel' @focus='onFocus'>
      </div>
    </basic-dialog>
  </div>
</template>

<style lang="less" module='s'>
  .item{
    display: flex;
    justify-content: flex-end;
    align-items: center;
    padding: 10px 40px;
    font-size: 18px;
    .red{
      color: red;
    }
    input {
      flex: 1;
      margin-left: 10px;
    }
  }
  
</style>

<script>
  import BasicDialog from '@/components/dialog/index'
  import { Message } from 'element-ui'
  import SearchSelect from '@/components/SearchSelect'
  export default {
    components: { BasicDialog, SearchSelect},
    data () {
      return {
        host: 'http://10.176.118.61:3000',
        streets: [],
        user: {
          UserName: '',
          RealName:'',
          Street:'',
          Password:'',
          Tel: ''
        },
        warn:''
      }
    },
    mounted () {
      this.fetchAllStreetName()
    },
    methods:{
      onSave () {
        if (!this.checkuser()) return
        this.adduser()
        console.info('onSave')
      },
      onCancel () {
        console.info('onCancel')
        this.$emit('close')
      },
      checkuser () {
        if( this.user.UserName !== ''
          && this.user.RealName !== ''
          && this.user.Tel !== ''
          && this.user.Street !== ''
          && this.user.Password !== '') return true
        return false
      },
      adduser () {
        // var request = {
        //   method: 'POST',
        //   url:'http://10.176.118.61:3000/user',
        //   query: {
        //     // put: JSON.stringify(this.user)
        //     put:''
        //   }
        // }
        // Ajax(request, data => {
        //   console.info('adduser',data)
        // })
        fetch('http://10.176.118.61:3000/streetUser/add', {
          method: 'POST',
          body: JSON.stringify(this.user)
        }).then(resp => {
          return resp.json()
        }).then(body => {
          if(body.error ==1) this.warn = body.data
          else {
            this.warn = '街道用户新增成功'
            Message({message:'恭喜，提交成功', type:'success'})
            this.onCancel()
            this.user.UserName = '',
            this.user.RealName = ''
            this.user.Tel = ''
            this.user.Street = ''
            this.user.Password = ''
          }
        })
      },
      fetchAllStreetName () {//获取所有街道名称
        fetch( this.host + '/street/key/name', {
          method: 'POST',
          body: '{}'
        }).then(resp => {
          return resp.json()
        }).then(body => {
          this.streets = body.data
        })
      },
      onFocus () {
        this.warn = ''
      }
    }
  }
</script>