<template>
  <div class="content">
    <div class="home"  >
      <div class="home1">
       <router-link to="/content1">
        <Button>添加</Button>
      </router-link>
      <router-link to="/exchange">
        <Button>交流区</Button>
      </router-link>
      <router-link to="/">
        <Button>退出登录</Button>
      </router-link>
       <router-link to="/homecopy">
          <Button style="margin-left: 8px">修改</Button>
        </router-link>
      </div>
      <!-- <Button @click="list">显示</Button>    点击触发-->
      <!-- 不影响自动实现 -->
      
      <div class="biede"  >  
        学校
      <ol>
        <li v-for="item in arr">
         <a   @click="search1(item.school)"> {{item.school}}</a>
             <!-- <button  @click="search1(item.school)" >查看</button>    -->
      </li>
      </ol>
    </div>
    
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
//import HelloWorld from '@/components/HelloWorld.vue'
import min from "../api/min"; //导入包

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
  .biede{
    top: 100px;
   width: 10%;
   height: 25%;
   position: relative;
   background-color: #ccc
  }
</style>


<style>
  .home1{
   left: 1350px;
    top: 10px;
   position: absolute;
  }
</style>