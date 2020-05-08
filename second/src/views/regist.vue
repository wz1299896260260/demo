<template>
<div class="registview">
<div class="regist">
  <!-- modelformCustom 与绑定   rule验证规则-->
  <Form ref="formCustom" :model="formCustom" :rules="ruleCustom" :label-width="80" >
    <FormItem label="账号" prop="username">
      <Input type="text" v-model="formCustom.username" placeholder="账号不能为汉字"></Input>
    </FormItem>
    <FormItem label="密码" prop="password">
      <Input type="password" v-model="formCustom.password"></Input>
    </FormItem>
    <FormItem label="昵称" prop="name">
      <Input type="text" v-model="formCustom.name"></Input>
    </FormItem>
    <FormItem label="邮箱" prop="email">
      <Input type="email" v-model="formCustom.email"></Input>
    </FormItem>
    <FormItem>
      <Button type="primary" @click="regist('formCustom')">注册</Button>
      <router-link to="/">
        <Button @click="handleReset('formCustom')" style="margin-left: 8px">返回</Button>
      </router-link>
    </FormItem>
  </Form>
  </div>
  </div>
</template>
<script>
import min from "../api/min";

export default {
  data() {
    return {
      formCustom: {
        // name:"",
        // desc:""
        username: "",
        password: "",
        name: "",
        email: "",
      },
      //不能为空
      ruleCustom: {
        username: [
          {
            required: true,
            message: "账号不能为空",
            trigger: "blur"
          },
           { type: 'string', min: 4,max:15, message: '账号最少4位,最多15位', trigger: 'blur' }
        ],
        password: [
          {
            required: true,
            message: "密码不能为空",
            trigger: "blur"
          },
          { type: 'string', min: 6, message: '密码需要6位', trigger: 'blur' }
        ],
        name: [
          {
            required: true,
            message: "昵称不能为空",
            trigger: "blur"
          }
        ],
        // phone: [
        //   {
        //     required: true,
        //     message: "电话不能为空",
        //     trigger: "blur"
        //   }
        // ],
        email: [
          {
            required: true,
            message: "邮箱不能为空",
            trigger: "blur"
          },
          { type: "email", message: "正确的邮箱地址", trigger: "blur" }
        ]
      }
    };
  },
  methods: {
    //註冊
    regist(name) {
      this.$refs[name].validate(valid => {
        if (valid) {
          //把数据 发到服务器   formCustom是data数据
         // this.formCustom.pass=md5(this.formCustom.pass)
          min.post("/api/user/add", this.formCustom).then(resp => {
            if(resp.code==200){
              this.$Message.success(resp.msg);
              
               this.$router.push({path:'/login'})
            }else{
               this.$Message.error(resp.msg);
            }
          });
        }
      });
    },
    handleReset(name) {
      this.$refs[name].resetFields();
    }
  }
};
</script>

<style>
  .regist{
    height: 100vh;
    width: 25vw;
    /* background-image: url('../assets/2.png'); */
    position: relative;
    left: 400px;
    top: 20px;
  }
</style>
<style>
  .registview{
    height: 100vh;
    width: 25vw;
    background-image: url('../assets/3.jpg');
    position: relative;
    /* left: 1000px;
    top: 20px; */
  }
</style>
