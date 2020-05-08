<template>
  <!-- <Button @click="list">显示</Button>    点击触发-->
  <!-- 不影响自动实现 -->
  <!-- search2(item.class) 传参数 -->
  
  <div>
  <div class="buttonclass">
    <ButtonGroup vertical>
  <Button @click="value2 = true" type="primary" ghost>Open</Button>
  <Button class="buttonclass1" to="/HomeCopy" >返回</Button>
    </ButtonGroup>
    <Drawer title="学校" placement="left" :closable="false" v-model="value2">
      <ol>
        <li v-for="item in arr">
          <a @click="search2(item.class)">{{item.class}}</a>
        </li>
      </ol>
    </Drawer>
    </div>
    <!-- <ol>
          <li v-for="item in arr">
            <a   @click="search2(item.class)"> {{item.class}} </a> 
        
        </li>
    </ol>-->
    <!-- <div>{{school}}</div>  -->
    <!-- <div v-for="item in student">
          {{item.name}}
        </div> 
    </div>-->
    <!-- <Table width="550" border :columns="columns2" :data="data3"></Table> -->
     <div class="classroom">
    <Table border :columns="columns12" :data="userData">
      <template slot-scope="{ row }" slot="name">
        <strong>{{ row.name }}</strong>
      </template>
      <template slot-scope="{ row, index }" slot="action">
        <!-- <Button type="primary" size="small" style="margin-right: 5px" @click="show(index)">编辑</Button> -->
        <Button type="primary" size="small" @click="remove(row)">编辑</Button>
        <!-- row是值，index是索引 -->
      </template>
    </Table>
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
      school: "",
      student: [],
      value2: false,
      columns12: [
        {
          title: "Name",
          slot: "name"
        },
        {
          title: "Age",
          key: "age"
        },
        {
          title: "Phone",
          key: "phone"
        },
        {
          title: "City",
          key: "city"
        },
        {
          title: "Sex",
          key: "sex"
        },
        {
          title: "QQ",
          key: "qq"
        },
        // {
        //   title: "Action",
        //   slot: "action",
        //   width: 150,
        //   align: "center"
        // }
      ],
      userData: []
    };
  },
  //自定义函数在加载时自动执行
  mounted() {
    this.list();
  },
 
  //列表方法课改
  methods: {
    list() {
      //路由路径
      //  var url=this.$route.path
      //  console.log(url)

      this.school = decodeURI(location.href.match("[^/]+(?!.*/)")[0]);
      // console.log(urlStr);
      min.get("/api/particular/get/" + this.school).then(resp => {
        // console.log(resp);
        if (resp.code == 200) {
          this.arr = resp.data;
        }
      });
    },
    search2(val) {
      //接收值
      min.get("/api/particular/get2/" + val).then(res => {
        // this.student = this.normalizeData(res.data)
        console.log(res);
        this.userData = this.normalizeData(res.data);
        // console.log(this.userData)

        // this.$router.push({ path: `/particular-school/${val}` })
      });
    },
     remove (row) {
       var name =row.name
       min.get("/api/particular/get3/"+name).then(res=>{
         console.log(res)
         if (res.code == 200) {  
               
              //this.$router.push({ path: "/password/"+num });
              
              this.$router.push({ path: "/editstudent/"+name });
              
            } else {
              this.$Message.error(resp.msg);
            }
       })
    },
    // 格式化数据
    normalizeData(list) {
      let result = [];
      list.map(item => {
        if (item.school === this.school) {
          result.push(item);
        }
      });
      return result;
    }
  }
};
</script>
<style>
  .classroom{
    height: 100vh;
    width: 75vw;
    /* background-image: url('../assets/2.png'); */
    position: relative;
    left: 300px;
    top: 50px;
  }
</style>

<style>
  .buttonclass{
    /* height: 100vh;
    width: 50vw; */
    /* background-image: url('../assets/2.png'); */
    position: absolute;
    left: 10px;
    top: 20px;
  }
</style>

<style>
  .buttonclass1{
    /* height: 100vh;
    width: 50vw; */
    /* background-image: url('../assets/2.png'); */
    position: absolute;
    left: 1px;
    top: 10px;
  }
</style>