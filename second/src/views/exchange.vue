
<template>
<div>
    <div class="text1">
    <List>
        <ListItem item-layout="vertical"  v-for="(item, index) in arr" :key="index">
            
           <ListItemMeta :title="item.name"  :description="item.content"  />
        </ListItem>
    </List>
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
        arr: [],

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
    min.get("/api/article/all").then(resp => {
        console.log(resp);
        if (resp.code == 200) {
          this.arr = resp.data;
        }
      });
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
    }
  }
};
</script>

<style>
.textview {
  left: 100px;
  top: 500px;
  width: 500px;
  position: absolute;
}
</style>

<style>
.text1 {
  left: 100px;
  width: 500px;
  position: absolute;
}
</style>