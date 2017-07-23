<template>
  <div>
  <!-- 新增街道信息 -->
    <basic-dialog 
      title='Title' 
      :onSave='onSave'
      :onCancel='onCancel'
      :warn='warn'
      >
      <div :class='s.item'>
        <div :class='s.red'>*</div>
        name
        <!-- 街道名 -->
        <input type="text" v-model='street.name' @focus='onFocus'>
      </div>
      <div :class='s.item'>
        <div :class='s.red'>*</div>
        charger
        <!-- 负责人 -->
        <input type="text" v-model='street.personInCharge' @focus='onFocus'>
      </div>
      <div :class='s.item'>
        <div :class='s.red'>*</div>
        tel
        <!-- 电话号码 -->
        <input type="text" v-model='street.tel' @focus='onFocus'>
      </div>
      <div :class='s.item'>
        <div :class='s.red'>*</div>
        authCode
        <!-- 授权码 -->
        <input type="text" v-model='street.authCode' @focus='onFocus'>
      </div>
      <div :class='s.item'>
        <div :class='s.red'>*</div>
        intro
        <!-- 描述 -->
        <input type="text" v-model='street.intro' @focus='onFocus'>
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

  export default {
    components: { BasicDialog },
    data () {
      return {
        street: {
          name: '',
          psersonInCharge:'',
          tel:'',
          authCode:'',
          intro: ''
        },
        warn:''
      }
    },
    methods:{
      onSave () {
        if (!this.checkStreet()) return
        this.addStreet()
        console.info('onSave')

      },
      onCancel () {
        console.info('onCancel')
        this.$emit('close')
      },
      checkStreet () {
        if( this.street.name !== ''
          && this.street.personInCharge !== ''
          && this.street.tel !== ''
          && this.street.authCode !== ''
          && this.street.intro !== '') return true
        return false
      },
      addStreet () {
        // var request = {
        //   method: 'POST',
        //   url:'http://10.176.118.61:3000/street',
        //   query: {
        //     // put: JSON.stringify(this.street)
        //     put:''
        //   }
        // }
        // Ajax(request, data => {
        //   console.info('addStreet',data)
        // })
        fetch('http://10.176.118.61:3000/street/add', {
          method: 'POST',
          body: JSON.stringify(this.street)
        }).then(resp => {
          console.info(resp)
          return resp.json()
        }).then(body => {
          if(body.error ==1) this.warn = body.data
          else {
            this.warn = '街道新增成功'
            this.street.name = '',
            this.street.personInCharge = ''
            this.street.tel = ''
            this.street.authCode = ''
            this.street.intro = ''
          }
        })
      },
      onFocus () {
        this.warn = ''
      }
    }
  }
</script>