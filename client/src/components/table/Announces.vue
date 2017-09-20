<template>
  <div :class='s.wrap'>
    <div :class='s.content'>
      <div :class='s.title'>
        <img src="~@/res/images/earth.png">
        政府公告<!-- 推送法院调解请求 -->
      </div>
      <div :class='s.body'>
        <!-- search result -->
        <div :class='s.uploadWrap' v-if='from === "gov"'>
          <el-button type='primary' class='add' icon='plus' @click='showUploadDialog = true'>新 增</el-button>
          <el-dialog
            title='上传文件'
            :visible.sync='showUploadDialog'>
            <upload
              :uploading='uploading'
              @cancel='showUploadDialog = false'
              @sure='onSure'></upload>
          </el-dialog>
        </div>
        <div>已发布公告列表</div>
        <div :class='s.fileItem' v-for='announce in announces'>
          <div 
            :class='s.file'
            ><a :href='host + "/open/file/" + announce.md5'>{{announce.filename}}</a>
          </div>
          <i class="iconfont icon-close" @click='onDel(announce)'  v-if='from === "gov"'></i>
        </div>
        <el-pagination
          layout="total, prev, pager, next"
          @current-change='onChangePage'
          :page-size='query.pageSize'
          :total="sum">
        </el-pagination>
      </div>
    </div>
    <el-dialog
      title='确认删除'
      :visible.sync='showDelDialog'>
      <a :href='host + "/open/file/" + delAnnounce.md5'>{{delAnnounce.filename}}</a>
      <div :class='s.bottomBtns'>
        <el-button type='primary' @click='sureDelAnnounce(delAnnounce)' :loading='delingAnnounce'>确 定</el-button>
        <el-button @click='showDelDialog = false'>取 消</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
  import fetchpm from '@/fetchpm'
  import Upload from '@/components/dialog/Upload'
  import { Message } from 'element-ui'
  export default {
    components: { Upload },
    filters: {},
    props: {
      from: {
        type: String,
        default: ''
      }
    },
    data () {
      return {
        host: "https://www.maszfglzx.com:3000",
        user: {},
        announces: [],
        newAnnounces: [],
        query: {
          fileName: '',
          pageNo: 0,
          pageSize: 10
        },
        sum: 0,
        showUploadDialog: false,
        uploading: false,
        showDelDialog: false,
        delAnnounce: {
          filename: '',
          md5: ''
        },
        delingAnnounce: false
      }
    },
    mounted () {
      this.user = JSON.parse(sessionStorage.getItem('user') || '{}') || {}
      this.fetchAnnounces(this.query)
    },
    methods: {
      onDel (announce) {
        this.showDelDialog = true
        this.delAnnounce = announce
      },
      sureDelAnnounce (announce) {
        this.delingAnnounce = true
        fetchpm(this, true, '/pm/announce/del/id/'+ announce._id, {
          method: 'GET'
        }).then(resp => {
          return resp.json()
        }).then(body => {
          console.info('sureDelAnnounce', body)
          if (body.error === 0) {
            Message({type:'success', message: '删除成功'})
            this.fetchAnnounces(this.query)
            this.showDelDialog = false
          } else {
            Message({type:'error', message: body.data})
          }
          this.delingAnnounce = false
        })
      },
      onChangePage (curPage) {
        this.query.pageNo = curPage - 1
        this.fetchAnnounces(this.query)
      },
      fetchAnnounces(query) {
        fetchpm(this, true, '/pm/announce/search/name', {
          method: 'POST',
          body: query
        }).then(resp => {
          return resp.json()
        }).then(body => {
          console.info('fetchAnnounces', body)
          if (body.error === 0 ) {
            this.announces = body.data.list || []
            this.sum = body.data.sum || 0
          }
        })
      },
      onSure (files) {
        if (!files || files.length === 0) {
          this.showUploadDialog = false
          return
        }
        // files[0].name
        // files[0].response.data.md5
        this.uploading = true
        for (var i = 0; i < files.length; i++) {
          let announce = {}
          announce.UserID = this.user.ID
          announce.MD5 = files[i].response.data.md5
          announce.FileName = files[i].name
          this.newAnnounces.push(announce)
        }
        this.postAnnounces(this.newAnnounces)
      },
      postAnnounces (announces) {
        fetchpm(this, true, '/pm/announce/add', {
          method: 'POST',
          body: {list: announces}
        }).then(resp => {
          return resp.json()
        }).then(body => {
          console.info('postAnnounces', body)
          if (body.error === 0) {
            this.fetchAnnounces(this.query)
            Message({type:'success', message: '恭喜，发布成功'})
            this.showUploadDialog = false
          } else {
            Message({type:'error', message: body.data})
          }
          this.uploading = false
        })
      }
    }
  }
</script>

<style lang="less" module='s'>
.wrap{
  // width: 100%;
  margin: 10px;
  .content{
    border: solid 1px #4c87b9;
    background-color: #fff;
    .title{
      color: #fff;
      font-size: 20px;
      background: #4c87b9;
      padding: 10px;
      display: flex;
      align-items: center;

      color: #fff;
      font-size: 30px;
      font-family: "华文行楷";
      background: #4c87b9;
      background: -webkit-gradient(linear,center top, center bottom,from(#ff0000), to(#000000));
      
      img{
        width: 20px;
        margin-right: 10px;
      }
    }
    .body{
      margin: 10px;
    }    
  }
}
.uploadWrap{
  
}
.fileItem{
  display: flex;
  align-items: center;
  font-size: 22px;
  .file{
    flex: 1;
  }
  i{
    color: #999;
  }
}
</style>