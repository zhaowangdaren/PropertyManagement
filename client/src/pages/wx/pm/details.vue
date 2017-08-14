<!-- XQID: xq.ID, XQName: xq.Name -->
<template>
  <div :class='s.wrap'>
    <action-bar :OPTIONS='headerOptions'></action-bar>
    <div :class='s.warn' v-text='warn'></div>
    <div :class='s.tab'>
      <div :class='s.line'></div>
      PM基本信息
      <div :class='s.line'></div>
    </div>
    <div :class='s.item'>
      <div :class='s.key'>XQ Name</div>
      <div :class='s.value' v-text='$route.query.XQName'></div>
    </div>
    <div :class='s.item'>
      <div :class='s.key'>PM Name</div>
      <div :class='s.value' v-text='pm.Name'></div>
    </div>
    <div :class='s.item'>
      <div :class='s.key'>LegalPerson</div>
      <div :class='s.value'>{{pm.LegalPerson === '' ? '暂无信息': pm.LegalPerson}}</div>
    </div>
    <div :class='s.item'>
      <div :class='s.key'>PM zizhi</div>
      <div :class='s.value'>{{pm.WuYeZiZhi === '' ? '暂无信息': pm.WuYeZiZhi}}</div>
    </div>
    <div :class='s.item'>
      <div :class='s.key'>PM类型</div>
      <div :class='s.value'>{{pm.WuYeXinZhi === '' ? '暂无信息': pm.WuYeXinZhi}}</div>
    </div>
    <div :class='s.item'>
      <div :class='s.key'>联系电话</div>
      <div :class='s.value'>{{pm.Tel === '' ? '暂无信息': pm.Tel}}</div>
    </div>
    <div :class='s.tab'>
      <div :class='s.line'></div>
      XQ PM基本信息
      <div :class='s.line'></div>
    </div>
    <div :class='s.item'>
      <div :class='s.key'>XQ环境</div>
      <div :class='s.value'>{{pm.XQEnv === '' ? '暂无信息': pm.XQEnv}}</div>
    </div>
    <div :class='s.item'>
      <div :class='s.key'>XQBJ</div>
      <div :class='s.value'>{{pm.XQCleaning === '' ? '暂无信息': pm.XQCleaning}}</div>
    </div>
    <div :class='s.item'>
      <div :class='s.key'>GreenVegetatio</div>
      <div :class='s.value'>{{pm.GreenVegetatio === '' ? '暂无信息': pm.GreenVegetatio}}</div>
    </div>
    <div :class='s.item'>
      <div :class='s.key'>GuanYangBaoHu</div>
      <div :class='s.value'>{{pm.GuanYangBaoHu === '' ? '暂无信息': pm.GuanYangBaoHu}}</div>
    </div>
    <div :class='s.item'>
      <div :class='s.key'>XiaoFanJianCha</div>
      <div :class='s.value'>{{pm.XiaoFanJianCha === '' ? '暂无信息': pm.XiaoFanJianCha}}</div>
    </div>
    <div :class='s.item'>
      <div :class='s.key'>DianTiBaoYang</div>
      <div :class='s.value'>{{pm.DianTiBaoYang === '' ? '暂无信息': pm.DianTiBaoYang}}</div>
    </div>
    <div :class='s.item'>
      <div :class='s.key'>CarParkInOrder</div>
      <div :class='s.value'>{{pm.CarParkInOrder === '' ? '暂无信息': pm.CarParkInOrder}}</div>
    </div>
    <div :class='s.item'>
      <div :class='s.key'>YeZhuCommunity</div>
      <div :class='s.value'>{{pm.YeZhuCommunity === '' ? '暂无信息': pm.YeZhuCommunity}}</div>
    </div>
    <div :class='s.item'>
      <div :class='s.key'>YeZhuCommunityTel</div>
      <div :class='s.value'>{{pm.YeZhuCommunityTel === '' ? '暂无信息': pm.YeZhuCommunityTel}}</div>
    </div>
    <div :class='s.item'>
      <div :class='s.key'>PastYearLevel</div>
      <div :class='s.value'>{{pm.PastYearLevel === '' ? '暂无信息': pm.PastYearLevel}}</div>
    </div>
    <div :class='s.item'>
      <div :class='s.key'>JianYiZhengGaiCuoShi</div>
      <div :class='s.value'>{{pm.JianYiZhengGaiCuoShi === '' ? '暂无信息': pm.JianYiZhengGaiCuoShi}}</div>
    </div>
    <div :class='s.tab'>
      <div :class='s.line'></div>
      XQ PM展示
      <div :class='s.line'></div>
    </div>
    <div :class='s.showImgs'>
      <img v-for='img in imgs' :src='host + "/open/image/" + img'>
    </div>
  </div>
</template>

<script>
import ActionBar from '@/components/mobile/ActionBar'
import filterImgs2Array from '@/filters/filterImgs2Array'
import fetchpm from '@/fetchpm'
export default {
  components: { ActionBar },
  filters: { filterImgs2Array },
  data () {
    return {
      headerOptions: {
        leftBtns: [{text:'上一步', event: null}],
        title: 'PM信息',
        rightBtns: []
      },
      pm: {},
      imgs:[]
    }
  },
  mounted () {
    this.headerOptions.leftBtns[0].event = this.onPrev
    this.fetchPM(this.$route.query.XQID)
  },
  methods: {
    onPrev () {
      this.$router.go(-1)
    },
    fetchPM (xqID) {
      fetchpm(this, false, '/open/pm/kvs', {
        method: 'POST',
        body: {XQID: xqID}
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('fetchPM', body)
        if (body.error === 0 && body.data.length > 0) {
          this.pm = body.data[0] || {}
          this.imgs = filterImgs2Array(this.pm.Imgs)
        }
        else this.warn = '查询失败'
      })
    }
  }
}
</script>

<style lang="less" module='s'>
.wrap{
  margin-top: 80px;
  .warn{
    color: red;
    text-align: center;
    margin: 10px;
  }
  .tab{
    text-align: center;
    font-size: 25px;
    font-weight: bold;
    color: #555;
    padding: 10px;
    display: flex;
    align-items: center;
    justify-content: center;
    .line{
      border-top: solid 1px #ccc;
      flex: 1;
      margin: 10px;
    }
  }
  .item{
    color: #999;
    padding: 10px;
    font-size: 20px;
    .value{
      color: #555;
      box-shadow: 0px 1px 1px 1px rgba(0,0,0,0.2) inset;
      padding: 10px;
      background-color: #fff;
    }
  }
  .showImgs {
    text-align: center;
    img{
      width: 90%;
      margin: 5px auto;
    }
  }
}
</style>