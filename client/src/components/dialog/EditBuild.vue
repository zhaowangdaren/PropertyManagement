<template>
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
      </table>
      <div :class='s.item'>
      <!-- 房屋使用变更 -->
        房屋使用变更
        <el-input :class='s.elInput' v-model='house.UseChange' type='textarea'></el-input>
      </div>
      <div :class='s.item'>
      <!-- 主题结构裂缝情况 -->
        主题结构裂缝情况
        <el-input :class='s.elInput' v-model='house.MainCrack' type='textarea'></el-input>
      </div>
      <div :class='s.item'>
      <!-- 地基沉降 -->
        地基沉降
        <el-input :class='s.elInput' v-model='house.FoundationDown' type='textarea'></el-input>
      </div>
      <div :class='s.item'>
      <!-- 主题结构倾斜状况 -->
        主题结构倾斜状况
        <el-input :class='s.elInput' v-model='house.MainSlant' type='textarea'></el-input>
      </div>
      <div :class='s.item'>
      <!-- 悬梁结构破坏 -->
        悬梁结构破坏
        <el-input :class='s.elInput' v-model='house.CantileverCrack' type='textarea'></el-input>
      </div>
      <div :class='s.item'>
      <!-- 女儿墙脱落情况 -->
        女儿墙脱落情况
        <el-input :class='s.elInput' v-model='house.ParapetOff' type='textarea'></el-input>
      </div>
      <div :class='s.item'>
      <!-- 外墙抹灰层剥落情况 -->
        外墙抹灰层剥落情况
        <el-input :class='s.elInput' v-model='house.OuterLloatedCoatOff' type='textarea'></el-input>
      </div>
      <div :class='s.item'>
      <!-- 房屋变形 -->
        地质灾害
        <el-input :class='s.elInput' v-model='house.HouseDeform' type='textarea'></el-input>
      </div>
      <!-- 地质灾害
      排水系统
      内部装修变更
      违规搭建加层等改变房屋寨河
      建筑物附近地质灾害治理情况
      房屋鉴定等级
      房屋图片 -->
      <div :class='s.item'>
        地址灾害治理情况
        <el-input :class='s.elInput' v-model='house.Disaster' type='textarea'></el-input>
      </div>
      <div :class='s.item'>
        地址灾害治理情况
        <el-input :class='s.elInput' v-model='house.DisasterManage' type='textarea'></el-input>
      </div>
      <div :class='s.item'>
        排水系统
        <el-input :class='s.elInput' v-model='house.DrainageSsystem' type='textarea'></el-input>
      </div>
      <div :class='s.item'>
        房屋内部装修主体结构变更破坏情况
        <el-input :class='s.elInput' v-model='house.InnerChange' type='textarea'></el-input>
      </div>
      <div :class='s.item'>
        违规搭建加层
        <el-input :class='s.elInput' v-model='house.IllegalBuild' type='textarea'></el-input>
      </div>
      <div :class='s.item'>
        房屋鉴定等级
        <el-input :class='s.elInput' v-model='house.RankAppraisal' type='textarea'></el-input>
      </div>
      <div :class='s.item'>
        <div :class='s.title'>图片</div>
        <div :class='s.uploadWrap'>
          <el-upload
            action="//47.94.7.154:3000/open/upload"
            list-type="picture-card"
            :on-preview="handlePictureCardPreview"
            :on-remove="handleRemove">
            <i class="el-icon-plus"></i>
          </el-upload>
        </div>
      </div>
    </div>
    <div slot="footer" :class="s.bts">
      <el-button @click="onCancel">取 消</el-button>
      <el-button type="primary" @click="onSave">提 交</el-button>
    </div>
    <el-dialog v-model="dialogVisible" size='tiny'>
      <img width="100%" :src="dialogImageUrl" alt="">
    </el-dialog>
  </div>
</template>

<script>
import SearchSelect from '@/components/SearchSelect'
import { Message } from 'element-ui'
import fetchpm from '@/fetchpm'
export default {
  components: {},
  props: {
    HOUSE: Object
  },
  data () {
    return {
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
        XQID: '',
        Year:''
      },
      dialogImageUrl: '',
      dialogVisible: false
    }
  },
  watch: {
    inputStreetID: function (val, old) {
      if (old == this.house.StreetID) this.inputCommunityID = ''
      this.fetchAllCommunitiesByStreetID(val)
    },
    inputCommunityID: function (val, old) {
      if (old == this.house.CommunityID) this.inputXQID = ''
      this.fetchAllXQByCommunityID(val)
    }
  },
  mounted () {
    Object.assign(this.house, this.HOUSE)
    this.inputStreetID = this.house.StreetID
    this.inputCommunityID = this.house.CommunityID
    this.inputXQID = this.house.XQID
    this.fetchAllXQs()
    this.fetchAllCommunities()
    this.fetchAllStreets()
    
  },
  methods: {
    handleRemove(file, fileList) {
      console.log(file, fileList);
    },
    handlePictureCardPreview(file) {
      this.dialogImageUrl = file.url;
      this.dialogVisible = true;
    },
    upload (e) {
      console.info('upload', e)
    },
    onFocus () {
      this.warn = ''
    },
    onSave () {
      this.house.StreetID = this.inputStreetID
      this.house.CommunityID = this.inputCommunityID
      this.house.XQID = this.inputXQID
      if (!this.checkHouse()) return
      this.updateHouse()
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
    updateHouse () {
      fetchpm(this, true, '/pm/house/update', {
        method: 'POST',
        body:this.house
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('updateHouse', body)
        if(body.error === 1) this.warn = body.data
        else {
          // this.warn = '修改成功'
          Message({message:'恭喜，修改成功', type:'success'})
          this.onCancel()
          this.$emit('editSucc')
        }
      })
    },
    fetchAllStreets () {//获取所有街道名称
      fetchpm( this, true, '/pm/street', {
        method: 'POST'
      }).then(resp => {
        return resp.json()
      }).then(body => {
        this.streets = body.data
      })
    },
    fetchAllCommunitiesByStreetID (streetID) {
      fetchpm( this, true, '/pm/community/kvs', {
        method: 'POST',
        body: {streetID: streetID}
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('fetchAllCommunities', body)
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
    },
    fetchAllCommunities () {
      fetchpm( this, true, '/pm/community', {
        method: 'POST'
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('fetchAllCommunities', body)
        this.communities = body.data
      })
    },
    fetchAllXQs () {
      fetchpm( this, true, '/pm/xq', {
        method: 'POST'
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('fetchAllXQs', body)
        this.xqs = body.data
      })
    },
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
        font-size: 20px;
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