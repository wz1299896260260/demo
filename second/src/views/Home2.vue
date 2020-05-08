<template>
  <div class="content">
    <div class="home1">
      <router-link to="/content1">
        <Button>添加</Button>
      </router-link>
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
        <List  class="list-container">
            <ListItem item-layout="vertical" v-for="(item, index) in arr" :key="index"> <a   @click="search1(item.school)"> {{item.school}}</a></ListItem>
        </List>
    </div>
      </div>
</template>

<script>
// @ is an alias to /src
//import HelloWorld from '@/components/HelloWorld.vue'
import min from "../api/min"; //导入包
import jwtDecode from 'jwt-decode'

export default {
  name: "Home",
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
       let token = localStorage.getItem('token')
        let newtoken = jwtDecode(token)
        console.log(newtoken.id)
        if (newtoken.id=='5'){
           min.get("/api/particular/all", val).then(res=>{
           this.$router.push({ path: `/adminschoolsearch/${val}` })
      })
        }else{
          this.$Message.error("你不是管理员");
        }
    },
  }
 
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
</style>


