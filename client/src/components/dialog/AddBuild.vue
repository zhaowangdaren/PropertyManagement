<template>
  <el-dialog
    title='新增房屋信息'
      :visible.sync='showDialog'
      @close='onColse'
      size='large'>
    <div :class='s.wrap'>
      <div :class='s.warn' v-if='warn !== ""' v-text='warn'></div>
      <div :class='s.content'>
        <!-- 新增建筑信息 -->
        <table :class='s.searchWrap'>
          <tr>
            <td :class='s.title'>编号</td>
            <td>
              <div :class='s.inputWrap'>
                <el-input v-model='house.BuildNo'></el-input>
              </div>
            </td>
            <td :class='s.title'>房产登记人</td>
            <td>
              <div :class='s.inputWrap'>
                <el-input v-model='house.Owner'></el-input>
              </div>
            </td>
            <td :class='s.title'>房屋类型</td>
            <td>
              <div :class='s.inputWrap'>
                <el-input v-model='house.HouseType'></el-input>
              </div>
            </td>
          </tr>
          <tr>
          <!-- 街道 社区 小区 -->
            <td :class='s.title'>街道</td>
            <td>
              <el-select :class='s.elSelect' v-model="inputStreetID" placeholder="请选择">
                <el-option
                  v-for="item in streets"
                  :key="item.ID"
                  :label="item.Name"
                  :value="item.ID">
                </el-option>
              </el-select>
            </td>
            <td :class='s.title'>社区</td>
            <td>
              <el-select :class='s.elSelect' v-model="inputCommunityID" placeholder="请选择">
                <el-option
                  v-for="item in communities"
                  :key="item.ID"
                  :label="item.Name"
                  :value="item.ID">
                </el-option>
              </el-select>
            </td>
            <td :class='s.title'>小区</td>
            <td>
              <el-select :class='s.elSelect' v-model="inputXQID" placeholder="请选择">
                <el-option
                  v-for="item in xqs"
                  :key="item.ID"
                  :label="item.Name"
                  :value="item.ID">
                </el-option>
              </el-select>
            </td>
          </tr>
          <tr>
          <!-- 楼栋号 门牌号 建筑年代 -->
            <td :class='s.title'>房屋楼栋号</td>
            <td>
              <div :class='s.inputWrap'>
                <el-input v-model='house.HouseBuildNo'></el-input>
              </div>
            </td>
            <td :class='s.title'>房屋门牌号</td>
            <td>
              <div :class='s.inputWrap'>
                <el-input v-model='house.HouseNo'></el-input>
              </div>
            </td>
            <td :class='s.title'>建筑年代</td>
            <td>
              <div :class='s.inputWrap'>
                <el-input v-model='house.Year'></el-input>
              </div>
            </td>
          </tr>
          <tr>
            <td :class='s.title'>房屋使用变更</td>
            <td colspan="5">
              <el-input :class='s.elInput' v-model='house.UseChange' type='textarea'></el-input>
            </td>
          </tr>
          <tr>
            <td :class='s.title'>主体结构裂缝情况</td>
            <td colspan="5">
              <el-input :class='s.elInput' v-model='house.MainCrack' type='textarea'></el-input>
            </td>
          </tr>
          <tr>
            <td :class='s.title'>地基沉降</td>
            <td colspan="5">
              <el-input :class='s.elInput' v-model='house.FoundationDown' type='textarea'></el-input>
            </td>
          </tr>
          <tr>
            <td :class='s.title'>主体结构倾斜状况</td>
            <td colspan="5">
              <el-input :class='s.elInput' v-model='house.MainSlant' type='textarea'></el-input>
            </td>
          </tr>
          <tr>
            <td :class='s.title'>悬梁结构破坏</td>
            <td colspan="5">
              <el-input :class='s.elInput' v-model='house.CantileverCrack' type='textarea'></el-input>
            </td>
          </tr>
          <tr>
            <td :class='s.title'>女儿墙脱落情况</td>
            <td colspan="5">
              <el-input :class='s.elInput' v-model='house.ParapetOff' type='textarea'></el-input>
            </td>
          </tr>
          <tr>
            <td :class='s.title'>外墙抹灰层剥落情况</td>
            <td colspan="5">
              <el-input :class='s.elInput' v-model='house.OuterLloatedCoatOff' type='textarea'></el-input>
            </td>
          </tr>
          <tr>
            <td :class='s.title'>房屋变形</td>
            <td colspan="5">
              <el-input :class='s.elInput' v-model='house.HouseDeform' type='textarea'></el-input>
            </td>
          </tr>
          <tr>
            <td :class='s.title'>地质灾害</td>
            <td colspan="5">
              <el-input :class='s.elInput' v-model='house.Disaster' type='textarea'></el-input>
            </td>
          </tr>
          <tr>
            <td :class='s.title'>地址灾害治理情况</td>
            <td colspan="5">
              <el-input :class='s.elInput' v-model='house.DisasterManage' type='textarea'></el-input>
            </td>
          </tr>
          <tr>
            <td :class='s.title'>排水系统</td>
            <td colspan="5">
              <el-input :class='s.elInput' v-model='house.DrainageSsystem' type='textarea'></el-input>
            </td>
          </tr>
          <tr>
            <td :class='s.title'>房屋内部装修主体结构变更破坏情况</td>
            <td colspan="5">
              <el-input :class='s.elInput' v-model='house.InnerChange' type='textarea'></el-input>
            </td>
          </tr>
          <tr>
            <td :class='s.title'>违规搭建加层</td>
            <td colspan="5">
              <el-input :class='s.elInput' v-model='house.IllegalBuild' type='textarea'></el-input>
            </td>
          </tr>
          <tr>
            <td :class='s.title'>房屋鉴定等级</td>
            <td colspan="5">
              <el-input :class='s.elInput' v-model='house.RankAppraisal' type='textarea'></el-input>
            </td>
          </tr>
          <!-- <tr>
            <td :class='s.title'>上传图片</td>
            <td colspan="5">
              <div :class='s.uploadWrap'>
                <el-upload
                  action="https://www.maszfglzx.com:3000/open/upload"
                  list-type="picture-card"
                  :on-preview="handlePictureCardPreview"
                  :on-remove="handleRemove">
                  <i class="el-icon-plus"></i>
                </el-upload>
              </div>
            </td>
          </tr> -->
        </table>
      </div>
      <div slot="footer" :class="s.bts">
        <el-button @click="onCancel">取 消</el-button>
        <el-button type="primary" @click="onAdd">提 交</el-button>
      </div>
    </div>
    <basic-dialog :visible.sync="showImg" title='图片'>
      <img width="100%" :src="dialogImageUrl" alt="">
    </basic-dialog>
  </el-dialog>
</template>

<script>
import BasicDialog from '@/components/dialog/BasicDialog'
import SearchSelect from '@/components/SearchSelect'
import { Message } from 'element-ui'
import fetchpm from '@/fetchpm'
export default {
  components: { BasicDialog, SearchSelect },
  data () {
    return {
      showDialog: true,
      warn:'',
      warn: '',
      inputStreetID: '',
      inputCommunityID: '',
      inputXQID: '',
      streets: [],
      communities: [],
      xqs: [],
      house: {
        BuildNo:'',
        Owner: '',
        StreetID: '',
        CommunityID: '',
        XQID: ''
      },
      dialogImageUrl: '',
      showImg: false,

    }
  },
  watch: {
    inputStreetID: function (val) {
      this.fetchAllCommunitiesByStreetID(val)
    },
    inputCommunityID: function (val) {
      this.fetchAllXQByCommunityID(val)
    }
  },
  mounted () {
    this.fetchAllStreets()
  },
  methods: {
    onColse () {
      this.$emit('close')
    },
    resetData () {
      this.house = {
        BuildNo:'',
        Owner: '',
        StreetID: '',
        CommunityID: '',
        XQID: ''
      }
    },
    handleRemove(file, fileList) {
      console.log(file, fileList);
    },
    handlePictureCardPreview(file) {
      this.dialogImageUrl = file.url;
      this.showImg = true;
    },
    upload (e) {
      console.info('upload', e)
    },
    onFocus () {
      this.warn = ''
    },
    onAdd () {
      this.house.StreetID = this.inputStreetID
      this.house.CommunityID = this.inputCommunityID
      this.house.XQID = this.inputXQID
      if (!this.checkHouse()) return
      this.addHouse()
      console.info('onSave')
      this.$emit('add')
    },
    onCancel () {
      this.$emit('close')
    },
    checkHouse () {
      if (this.house.BuildNo == ''){
        this.warn = '请填写BuildNo'
        return false
      }
      if (this.house.Owner == ''){
        this.warn = '请填写Owner'
        return false
      }
      if (this.house.StreetID == ''){
        this.warn = '请选择StreetID'
        return false
      }
      if (this.house.CommunityID == ''){
        this.warn = '请填写CommunityID'
        return false
      }
      return true
    },
    addHouse () {
      fetchpm(this, true, '/pm/house/add', {
        method: 'POST',
        body:this.house
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('addCountry', body)
        if(body.error === 1) this.warn = body.data
        else {
          // this.warn = 'Add Succ'
          Message({message:'恭喜，成功新增建筑信息', type:'success'})
          this.$emit('addSucc')
          this.onCancel()
        }
      })
    },
    fetchAllStreets () {//获取所有街道名称
      fetchpm( this, true, '/pm/street', {
        method: 'POST'
      }).then(resp => {
        return resp.json()
      }).then(body => {
        this.streets = body.data.streets
      })
    },
    
    fetchAllCommunitiesByStreetID (streetID) {
      fetchpm( this, true, '/pm/community/kvs', {
        method: 'POST',
        body: {streetID: streetID}
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('fetchAllCommunitiesByStreetID', body)
        this.communities = body.data
      })
    },
    fetchAllXQByCommunityID (communityID) {
      if ( !communityID || communityID == '') return null
      fetchpm( this, true, '/pm/xq/kvs', {
        method: 'POST',
        body: {communityID: communityID}
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('fetchAllXQByCommunityID', body)
        this.xqs = body.data
      })
    }
  },
  beforeDestroy() {
  }
}
</script>

<style lang="less" module='s'>
.wrap{
  width: 100%;
  .warn{
    text-align: center;
    color: red;
    font-size: 30px;
  }
  .content{
    // width: 100%;
    overflow-y: scroll;
    height: 500px;
    .searchWrap{
      border: solid 1px #ddd;
      .title{
        background: #f0f0f0;
        padding: 10px;
        font-size: 16px;
        max-width: 50px;
      }
      .elSelect{
        width: 100%;
      }
    }
    .item{
      display: flex;
      justify-content: flex-end;
      align-items: center;
      // padding: 10px 0px;
      font-size: 18px;
      margin: 20px;
      .red{
        color: red;
      }
      .elInput{
        width: 75%;
        margin-left: 10px
      }
    }
  }
  .bts{
    padding-top: 20px;
    display: flex;
    justify-content: flex-end;
  }
}
</style>