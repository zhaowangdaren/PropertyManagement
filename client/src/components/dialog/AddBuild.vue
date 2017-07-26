<template>
  <div>
    <!-- 新增建筑信息 -->
      <table :class='s.searchWrap'>
        <tr>
          <td :class='s.title'>BuildNo</td>
          <td>
            <div :class='s.inputWrap'>
              <input type="text" v-model='house.BuildNo'>
            </div>
          </td>
          <td :class='s.title'>Owner</td>
          <td>
            <div :class='s.inputWrap'>
              <input type="text" v-model='house.Owner'>
            </div>
          </td>
          <td :class='s.title'>Type</td>
          <td>
            <div :class='s.inputWrap'>
              <input type="text" v-model='house.HouseType'>
            </div>
          </td>
        </tr>
        <tr>
        <!-- 街道 社区 小区 -->
          <td :class='s.title'>Street</td>
          <td>
            <search-select v-model='house.Street' :values='streets'/>
          </td>
          <td :class='s.title'>Community</td>
          <td>
            <search-select v-model='house.Community' :values='communities'/>
          </td>
          <td :class='s.title'>XQ</td>
          <td>
            <search-select v-model='house.XQ' :values='xqNames' />
          </td>
        </tr>
        <tr>
        <!-- 楼栋号 门牌号 建筑年代 -->
          <td :class='s.title'>HouseBuildNo</td>
          <td>
            <div :class='s.inputWrap'>
              <input type="text" v-model='house.HouseBuildNo'>
            </div>
          </td>
          <td :class='s.title'>HouseNo</td>
          <td>
            <div :class='s.inputWrap'>
              <input type="text" v-model='house.HouseNo'>
            </div>
          </td>
          <td :class='s.title'>Year</td>
          <td>
            <div :class='s.inputWrap'>
              <input type="text" v-model='house.Year'>
            </div>
          </td>
        </tr>
      </table>
      <div :class='s.item'>
      <!-- 房屋使用变更 -->
        UseChange
        <textarea v-model='house.UseChange'></textarea>
      </div>
      <div :class='s.item'>
      <!-- 主题结构裂缝情况 -->
        MainCrack
        <textarea v-model='house.MainCrack'></textarea>
      </div>
      <div :class='s.item'>
      <!-- 地基沉降 -->
        FoundationDown
        <textarea v-model='house.FoundationDown'></textarea>
      </div>
      <div :class='s.item'>
      <!-- 主题结构倾斜状况 -->
        MainSlant
        <textarea v-model='house.MainSlant'></textarea>
      </div>
      <div :class='s.item'>
      <!-- 悬梁结构破坏 -->
        CantileverCrack
        <textarea v-model='house.CantileverCrack'></textarea>
      </div>
      <div :class='s.item'>
      <!-- 女儿墙脱落情况 -->
        ParapetFallOff
        <textarea v-model='house.ParapetOff'></textarea>
      </div>
      <div :class='s.item'>
      <!-- 外墙抹灰层剥落情况 -->
        OuterLloatedCoatOff
        <textarea v-model='house.OuterLloatedCoatOff'></textarea>
      </div>
      <div :class='s.item'>
      <!-- 房屋变形 -->
        HouseDeform
        <textarea v-model='house.HouseDeform'></textarea>
      </div>
      <!-- 地质灾害
      排水系统
      内部装修变更
      违规搭建加层等改变房屋寨河
      建筑物附近地质灾害治理情况
      房屋鉴定等级
      房屋图片 -->
      <div :class='s.item'>
        Disaster
        <textarea v-model='house.Disaster'></textarea>
      </div>
      <div :class='s.item'>
        DisasterManage
        <textarea v-model='house.DisasterManage'></textarea>
      </div>
      <div :class='s.item'>
        DrainageSsystem
        <textarea v-model='house.DrainageSsystem'></textarea>
      </div>
      <div :class='s.item'>
        InnerChange
        <textarea v-model='house.InnerChange'></textarea>
      </div>
      <div :class='s.item'>
        IllegalBuild
        <textarea v-model='house.IllegalBuild'></textarea>
      </div>
      <div :class='s.item'>
        RankAppraisal
        <textarea v-model='house.RankAppraisal'></textarea>
      </div>
      <div :class='s.item'>
        <div :class='s.title'>Images</div>
        <div :class='s.uploadWrap'>
          <el-upload
            action="https://jsonplaceholder.typicode.com/posts/"
            list-type="picture-card"
            :on-preview="handlePictureCardPreview"
            :on-remove="handleRemove">
            <i class="el-icon-plus"></i>
          </el-upload>
        </div>
      </div>
      <div slot="footer" class="dialog-footer">
        <el-button @click="onCancel">取 消</el-button>
        <el-button type="primary" @click="onSave">提 交</el-button>
      </div>
    <el-dialog v-model="dialogVisible" size='tiny'>
      <img width="100%" :src="dialogImageUrl" alt="">
    </el-dialog>
  </div>
</template>

<style lang="less" module='s'>
  .searchWrap{
    border: solid 1px #ddd;
    .title{
      background: #f0f0f0;
      padding: 10px;
      font-size: 20px;
    }
    .inputWrap {
        position: relative;
        font-size: 18px;
        flex: 1;
        margin-left: 10px;
      input{
        width: 100%;
        border: solid 1px #ddd;
        height: 30px;
      }
    } 
    .searchBt{
      position: absolute;
      right: 10px;
    }
  }
  .item{
    display: flex;
    justify-content: flex-end;
    align-items: center;
    padding: 10px 0px;
    font-size: 18px;
    .red{
      color: red;
    }
    textarea {
      flex: 1;
      margin-left: 10px;
      height: 30px;
    }
    .uploadWrap{
      flex: 1;
      padding: 24px;
    }
  }
  
</style>

<script>
import BasicDialog from '@/components/dialog/index'
import SearchSelect from '@/components/SearchSelect'
export default {
  components: { BasicDialog, SearchSelect },
  data () {
    return {
      host: 'http://10.176.118.61:3000',
      showSelf: false,
      warn: '',
      streets: [],
      communities: [],
      xqNames: [],
      house: {
        name: '',
        address: '',
        street: '',
        community: '',
        contactName: '',//联系人
        tel: '',
        intro:''
      },
      dialogImageUrl: '',
      dialogVisible: false
    }
  },
  mounted () {
    this.showSelf = true
    this.fetchAllStreetName()
    this.fetchAllCommunityName()
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
      if (!this.checkCountry()) return
      this.addCountry()
      console.info('onSave')
      this.$emit('save')
    },
    onCancel () {
      this.$emit('close')
    },
    checkCountry () {
      if (this.country.name !== ''
        && this.country.address !== ''
        && this.country.street !== ''
        && this.country.community !== ''
        && this.country.contactName !== ''
        && this.country.tel !== ''
        && this.country.intro !== '') return true
      return false
    },
    addCountry () {
      fetch(this.host + '/xq/add', {
        method: 'POST',
        body:JSON.stringify(this.country)
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('addCountry', body)
        if(body.error === 1) this.warn = body.data
        else {
          this.warn = 'Add Succ'
          this.country.name = ''
          this.country.address = ''
          this.country.street = ''
          this.country.community = ''
          this.country.contactName = ''
          this.country.tel = ''
          this.country.intro = ''
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
    fetchAllCommunityName () {
      fetch( this.host + '/community/key/name', {
        method: 'POST',
        body: '{}'
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('fetchAllCommunityName', body)
        this.communities = body.data
      })
    }
  },
  beforeDestroy() {
    this.showSelf = false
  }
}
</script>