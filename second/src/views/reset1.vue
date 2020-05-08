<template>
  
      <!-- <Button @click="list">显示</Button>    点击触发-->
      <!-- 不影响自动实现 -->
      <div class="useredit"  >
     <dl>
       <dt>用户名:</dt> 
       <dd>{{userData.username}}</dd>
       <dt>更改密码</dt>
       <dd>
         <!-- {{userData.pass}} -->
        <Input type="text" v-model="newPass" />
       </dd>
     </dl>
      <button @click="edit">更改</button>
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
      userData: [], // 用户信息
      newPass: '' // 新密码
    };
  },

  //自定义函数在加载时自动执行
    mounted() {
        this.list()
    },

  //列表方法课改
  methods: {
    list() { 
      // console.log(location.href.match('[^/]+(?!.*/)')[0])
      var urlStr = location.href.match('[^/]+(?!.*/)')[0];
      console.log(urlStr);
    min.get("/api/user/search/"+urlStr).then(resp => {
        console.log(resp);
        if (resp.code == 200) {
          this.userData = resp.data;
          //  this.userData.pass = ''
          // console.log(this.userData)
        }
      });
    },
    edit(){
      this.userData.password = this.newPass
      min.post('/api/user/edit', this.userData).then(res => {
        console.log(res)
        if(res.code==200){
        this.$Message.success(res.msg);
      }else {
              this.$Message.error(res.msg);
            }
      })
      
    }
  }
 
};
</script>

<style>
  .useredit{
    height: 100vh;
    width: 25vw;
    /* background-image: url('../assets/2.png'); */
    position: relative;
    left: 400px;
    top: 20px;
  }
</style>