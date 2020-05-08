
<template>
<div class="exchange-container">
    <div class="text1">
    <List>
        <ListItem item-layout="vertical"  v-for="(item, index) in historyData"  :key="index">      
           <ListItemMeta :title="item.name"  :description="item.content"  />
        </ListItem>
    </List>
    <Page :total="dataCount" :page-size="pageSize" show-total  @on-change="changepage"></Page>
    </div>


  <div class="textview">
    <Form ref="formValidate" :model="formValidate" :rules="ruleValidate" :label-width="80">
      <FormItem label="Text" prop="content">
        <Input v-model="formValidate.content" type="textarea" :autosize="{minRows: 2,maxRows: 5}" />
      </FormItem>
      <FormItem>
        <Button type="primary" @click="handleSubmit('formValidate')">确定</Button>
        <router-link to="/homecopy">
          <Button style="margin-left: 8px">返回</Button>
        </router-link>
      </FormItem>
    </Form>
  </div>
  </div>
</template>
<script>
import jwtDecode from "jwt-decode";
import min from "../api/min";
export default {
  data() {
    return {
     ajaxHistoryData:[],
    // 初始化信息总条数
      dataCount:0,
      // 每页显示多少条
     pageSize:10,
     historyData: [],
      formValidate: {
        content: ""
      },
      ruleValidate: {
        text: [{ required: true, message: "不能为空", trigger: "blur" }]
      }
    };
  },
   mounted() {
        this.list()
    },

  methods: {
      list() {
    min.get("/api/article/page").then(resp => {
        console.log(resp);
        if (resp.code == 200) {
          // this.arr = resp.data.items;
           this.ajaxHistoryData = resp.data.items;
                this.dataCount = resp.data.count;
                // 初始化显示，小于每页显示条数，全显，大于每页显示条数，取前每页条数显示
                if(resp.data.count < this.pageSize){
                    this.historyData = this.ajaxHistoryData;
                }else{
                    this.historyData = this.ajaxHistoryData.slice(0,this.pageSize);
                }
        }
      });
    },
    changepage(index){
      console.log(index)
                var _start = ( index - 1 ) * this.pageSize;
                var _end = index * this.pageSize;
                this.historyData = this.ajaxHistoryData.slice(_start,_end);
            },
             created(){
             this.list();
        },
          handleSubmit(name) {
      let token = localStorage.getItem("token");
      let newtoken = jwtDecode(token);
      this.formValidate.uid = newtoken.id;
      console.log(this.formValidate);
      this.$refs[name].validate(valid => {
        if (valid) {
          min.post("/adm/article/add", this.formValidate).then(resp => {
            if (resp.code == 200) {
              this.$Message.success(resp.msg);
            } else {
              this.$Message.error(resp.msg);
            }
          });
        }
      });
    },
    },
 
};
</script>

<style>
.exchange-container{
  height: 100vh;
  display: flex;
  flex-direction: column;
}
.textview {
  margin-top: 50px;
}

.text1 {
  text-align: center;
  position: relative;
}
.ivu-page{
  position: absolute;
  left: 20px;
}
</style>