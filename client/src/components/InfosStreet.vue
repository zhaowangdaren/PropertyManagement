<template>
  <div :class='s.wrap'>
    <div :class='s.addDel'>
      <image-button :class='s.bt' :clickMethod='onAdd'
        text='新增'
        :img='require("@/res/images/add.png")'
        bgColor='#3598dc'
        fontSize='20px'
        />
      <image-button :class='s.bt'
        text='删除'
        :img='require("@/res/images/delete.png")'
        bgColor='#cb5a5e'
        fontSize='20px'
        />
    </div>
    <table>
      <tr >
        <th>选择</th>
        <th>name</th>
        <!-- 街道名 -->
        <th>charger</th>
        <!-- 负责人 -->
        <th>Tel</th>
        <!-- 电话号码 -->
        <th>code</th>
        <!-- 授权码 -->
        <th :class='s.descr'>intro</th>
        <!-- 描述 -->
        <th>操作</th>
      </tr>
      <tr v-for='street in streets' :class='s.street'>
        <td>
          <input type="checkbox" v-model='street.checked'>
        </td>
        <td v-text='street.Name'></td>
        <td v-text='street.PersonInCharge'></td>
        <td v-text='street.Tel'></td>
        <td v-text='street.AuthCode'></td>
        <td v-text='street.Intro'></td>
        <td>
          <el-button type="primary" icon="edit" @click="onEdit(street)">编辑</el-button>
        </td>
      </tr>
    </table>
    <el-dialog 
      title='Add Street'
      :visible.sync='showAddDialog'
      size='small'>
      <add-street></add-street>
    </el-dialog>
    <el-dialog 
      title='Edit Street'
      :visible.sync='showEditDialog'
      size='small'>
      <edit-street></edit-street>
    </el-dialog>
  </div>
</template>

<script type="text/javascript">
  import Ajax from '@/Ajax'
  import ImageButton from '@/components/ImageButton'
  import AddStreet from '@/components/dialog/AddStreet'
  import EditStreet from '@/components/dialog/EditStreet'
  import fetchpm from '@/fetchpm'
  export default {
    components: {ImageButton, AddStreet, EditStreet},
    data () {
      return {
        streets:[],
        showAddDialog: false,
        showEditDialog: false
      }
    },
    mounted () {
      this.fetechStreets()
    },
    methods: {
      fetechStreets () {
        // var request = {
        //   method: 'POST',
        //   url:'http://10.176.118.61:3000/streetInfo'
        // }
        // Ajax(request, data => {
        //   if (data === null) return
        //   this.streets = JSON.parse(data).Streets
        // })
        fetchpm(this, true, '/pm/street',{
          method: 'POST'
        }).then(resp => {
          console.info(resp)
          return resp.json()
        }).then( data => {
          // console.info('fetechStreets', data)
          if (data.error === 0) {
            console.info (data)
            this.streets = data.data
          }
        })
      },
      onAdd () {
        console.info('onAdd')
        this.showAddDialog = true
      },
      onEdit (street) {
        this.showEditDialog = true
      }
    }
  }
</script>

<style lang="less" module='s'>
.wrap{
  // margin: 10px;
  .addDel{
    display: flex;
    align-items: center;
    margin-top: 10px;
    .bt {
      margin: 5px;
    }
  }
  table{
    width: 99%;
    font-size: 15px;
    color: #555;
    margin: 10px auto;
    text-align: center;
    th{
      text-align: center;
      padding: 5px;
      border: solid 1px #ddd;
    }
    td{
      padding: 5px;
      border: solid 1px #ddd;
    }
    .street{
      &:hover {
        background-color: #ddd;
      }
    }
  }
}
</style>