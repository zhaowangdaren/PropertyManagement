<template>
  <div :class='s.wrap'>
    <div v-if='identityLevel !== 0 && curStreet !== ""' :class='s.street' v-text='curStreet'></div>
    <pms></pms>
    <component :is='showDialog' @close='showDialog = ""' />
  </div>
</template>

<script>
  import AddPM from '@/components/dialog/AddPM'
  import PMs from '@/components/table/PMs'
  export default {
    components: {AddPM, 'pms': PMs},
    props: {
      identityLevel: {//当前用户的等级 0为admin,
        type: Number, 
        default: 1
      },
      curStreet: String //传入的当前街道
    },
    data () {
      return {
        host:'http://10.176.118.61:3000',
        showDialog: '',
        streetNames: [],
        communityNames: [],
        xqNames: [],
        inputStreetName:'',
        inputCommunityName: '',
        inputXiaoQu:'',
        isLoadingInput: false,
        pms:[]
      }
    },
    computed: {
    },
    mounted () {
      this.fetchPMs()
      this.fetechAllStreetName()
    },
    watch: {
      inputStreetName: function (val) {
        if (this.isLoadingInput ) return
        this.fetchCommunitiesByStreetName(val)
      },
      inputCommunityName: function (val) {
        console.info('inputCommunityName', val)
        if (this.isLoadingInput ) return
        this.fetchXQByCommunityName(val)
      }
    },
    methods: {
      onEdit (item) {

      },
      onAdd () {
        this.showDialog = AddCountry
      },
      onSearch () {
        //小区》社区》街道
        if (this.inputXiaoQu !== '') {
          this.fetchPMKVs({xiaoQu: this.inputXiaoQu})
          return
        }
        if (this.inputCommunityName !== '') {
          this.fetchPMKVs({community: this.inputCommunityName})
          return
        }
        if (this.inputStreetName !== '') {
          this.fetchPMKVs({street: this.inputStreetName})
          return
        }
      },
      fetchPMs () {
        fetch(this.host + '/pm', {
          method: 'POST',
          body: '{}'
        }).then(resp =>{
          return resp.json()
        }).then(body => {
          console.info('fetchPMs', body)
          if (body.error !== 1) this.pms = body.data
        })
      },
      fetchPMKVs (kvs) {
        fetch(this.host + '/pm/kvs', {
          method: 'POST',
          body: JSON.stringify(kvs)
        }).then(resp => {
          return resp.json()
        }).then(body => {
          console.info('fetchPMKVs', body)
          if (body.error !== 1) this.pms = body.data
        })
      },
      fetechAllStreetName () {
        fetch(this.host + '/street/key/name', {
          method: 'POST',
          body: '{}'
        }).then(resp => {
          return resp.json()
        }).then(body => {
          console.info('fetechAllStreetName',body)
          if (body.error !== 1) this.streetNames = body.data
        })
      },
      fetchCommunitiesByStreetName (streetName) {
        if (!streetName) return
        this.isLoadingInput = true
        this.communityNames = []
        this.inputCommunityName = ''
        fetch(this.host + '/community/streetName/'+streetName, {
          method: 'POST'
        }).then(resp => {
          return resp.json()
        }).then( body => {
          console.info('fetchCommunitiesByStreetName', body)
          if(body.error !== 0) {
            console.error("Error: search CommunitiesByStreetName. Reason:" + body.data)
          } else if (body.data !== null ) {
            this.communityNames = body.data.map(item => {
              return item.Name
            })
          }
          this.isLoadingInput = false
        })
      },
      fetchXQByCommunityName (communityName) {
        if (!communityName) return
        this.isLoadingInput = true
        fetch(this.host + '/xq/kvs', {
          method: 'POST',
          body: JSON.stringify({community: communityName})
        }).then(resp => {
          return resp.json()
        }).then( body => {
          console.info('fetchXQByCommunityName', body)
          if(body.error !== 0) console.error("Error: search fetchXQByCommunityName. Reason:" + body.data)
          else if (body.data !== null ) {
            this.xqNames = body.data.map(item => {
              return item.Name
            })
          }
          this.isLoadingInput = false
        })
      }
    }
  }
</script>

<style lang="less" module='s'>
.wrap{
  margin: 10px;
  font-size: 15px;
  .street{
    background-color: #ddd;
    width: 100%;
    font-size: 25px;
    padding: 5px;
  }
  
}
</style>