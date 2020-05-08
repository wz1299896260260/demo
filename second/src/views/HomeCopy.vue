<template>
  <div class="content">
    <div class="home1">
      <router-link to="/content1">
        <Button>添加</Button>
      </router-link>
      <router-link to="/exchange">
        <Button>交流区</Button>
      </router-link>
        <Button @click="edit" >修改</Button>
    </div>
      <div class="biede1" >
        <!-- 学校
        <ol>
        <li v-for="item in arr">
         <a   @click="search1(item.school)"> {{item.school}}</a>
        </li>
        </ol> -->
          <strong>学校</strong>
         <!-- <List item-layout="vertical">
            <ListItem v-for="item in arr" :key="item.id">
              <a   @click="search1(item.school)"> {{item.school}}</a></ListItem>
        </List> -->
        <List class="list-container"  >
            <ListItem  item-layout="vertical" v-for="(item, index) in arr"  v-if='index<5' :key="index"> <a   @click="search1(item.school)"> {{item.school}}</a></ListItem>
        </List>
     <router-link to="/home">更多</router-link>
    </div>
      </div>
</template>

<script>
// @ is an alias to /src
//import HelloWorld from '@/components/HelloWorld.vue'
import min from "../api/min"; //导入包
import jwtDecode from 'jwt-decode'

export default {
 
  //  name: "Home",
  data() {
    return {
      arr: [],

    };
  },

  //自定义函数在加载时自动执行
    mounted() {
        this.list()
    },

  //列表方法课改
  methods: {
    list() {
    min.get("/api/particular/all").then(resp => {
        // console.log(resp);
        if (resp.code == 200) {
          this.arr = resp.data;
        }
      });
    },
    search1(val){
      //  console.log(val)
      //  this.resp.data.school=this.val
      //  console.log(resp.data.school)
      min.get("/api/particular/all", val).then(res=>{
        //  console.log(res)
              this.$router.push({ path: `/schoolsearch/${val}` })
          // if (res.code == 200) {       
          //     //this.$router.push({ path: "/password/"+num });
              
          //     this.$router.push({ path: `/schoolsearch/${val}` })
              
          //   } else {
          //     this.$Message.error(res.msg);
          //   }
      })
     
    },
    edit(){
       let token = localStorage.getItem('token')
      let newtoken = jwtDecode(token)
      var uid=newtoken.id
      console.log(uid)
          min.get("/adm/particular/get/"+uid).then(res=>{
        console.log(res)
              this.$router.push({ path: "/editstudent/"+uid })
          
      })
    }
  },
 
 
};
</script>

<style>
  .content{
    /* 100为百分百填充网页 */
    height: 100vh;
    width: 100vw;
    background-image: url('../assets/1.jpg');
  }
</style>

<style>
  .biede1{
    top: 100px;
    display: flex;
    flex-direction: column;
    position: absolute;
    background-color: rgba(168, 135, 135, 0.7)
  }
  .biede1 > .list-container {
    padding: 5em;
    display: flex;
    flex-direction: column;
  }
  .ivu-list-item{
    border: none !important;
  }
  .ivu-list-split{
    border: none
  }
  
</style>




