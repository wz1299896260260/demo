<template>
<div class="1">
   <Form ref="formValidate" :model="formValidate" :rules="ruleValidate" :label-width="80">
    <FormItem label="修改账号" prop="username">
      <Input type="text" v-model="formValidate.username"></Input>
    </FormItem>
    <!-- <FormItem label="密码" prop="pass">
            <Input  type="password" v-model="formValidate.mail" ></Input>
    </FormItem>-->
    <FormItem>
      <Button type="primary" @click="handleSubmit('formValidate')">修改密码</Button>
      <router-link to="/">
        <Button>返回</Button>
      </router-link>
    </FormItem>
  </Form>
   

</div>
</template>

<script>
import min from "../api/min";
export default {
  data() {
    return {
      formValidate: {
        username: ""
        // pass: '',
      },
      ruleValidate: {
        username: [{ required: true, message: "为空", trigger: "blur" }]
        // pass: [
        //     { required: true, message: '为空', trigger: 'blur' },
        // ],
      },
    };
  },
  methods: {
    handleSubmit(name) {
      this.$refs[name].validate(valid => {
        if (valid) {
             var username=this.formValidate.username
             console.log(username)
          min.get("/api/user/search/"+username, this.formValidate).then(resp => {
              console.log(resp)
            if (resp.code == 200) {  
               
              //this.$router.push({ path: "/password/"+num });
              
              this.$router.push({ path: "/reset1/"+username });
              
            } else {
              this.$Message.error(resp.msg);
            }
          });
        }
      });
    },
    
  }
};
</script>