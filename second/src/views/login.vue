
<template>
<div class="login" @keyup.enter="ada">
<div class="font">
  同学录
</div>
  <Form ref="formInline" :model="formInline" :rules="ruleInline" inline  >
    <br><br><br><br>
    <FormItem prop="username">
      <Input type="text" v-model="formInline.username" placeholder="用户名">
        <Icon type="ios-person-outline" slot="prepend"></Icon>
      </Input>
    </FormItem>
    <br>
    <FormItem prop="password">
      <Input type="password" v-model="formInline.password" placeholder="密码">
        <Icon type="ios-lock-outline" slot="prepend"></Icon>
      </Input>
    </FormItem>
    <br>
    <FormItem>
        <Button type="primary" @click="clicklogin('formInline')">登录</Button>
        <router-link to="/regist"> <Button  >注册</Button></router-link>  
        <router-link to="/HomeCopy"> <Button  >返回</Button></router-link>
          <router-link to="/reset"> <Button  >忘记密码</Button></router-link>
    </FormItem>
  </Form>
  </div>
</template>
<script>
import min from "../api/min";
 import md5 from 'js-md5';
import jwtDecode from 'jwt-decode'

export default {
  data() {
    return {
      formInline: {
        username: "",
        password: ""
      },
      ruleInline: {
        username: [{ required: true, message: "用户名为空", trigger: "blur" }],
        password: [
          { required: true, message: "密码为空", trigger: "blur" },
          { type: "string", min: 6, message: "密码需要6位", trigger: "blur" }
        ]
      }
    };
  },
  methods: {
    clicklogin(name) {
      this.$refs[name].validate(valid => {
        if (valid) {
          this.formInline.password=md5(this.formInline.password)
          min.post("/api/login", this.formInline).then(resp => {
            console.log(resp)
           if(resp.code==200){
              localStorage.setItem('token',resp.data)
               let token = localStorage.getItem('token')
               let newtoken = jwtDecode(token)
              this.formInline.uid = newtoken.id
               this.$Message.success(resp.msg);
              if(this.formInline.uid=='5'){
                  this.$router.push({path:'/Home2'})
              }else{
                 this.$router.push({path:'/HomeCopy'})
              }
            }else{
               this.$Message.error(resp.msg);
            }
          });
        }
      });
      
    },
    ada() {
      console.log(11)
      
      this.clicklogin('formInline')
    }
  }
};
</script>
<style>
  .login{
    height: 100vh;
    width: 100vw;
    background-image: url('../assets/2.png');
  }
</style>

<style>
  .font{
   font-size: 40px;
   
  }
</style>