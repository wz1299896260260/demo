<template>
  <Form ref="formCustom" :model="formCustom" :rules="ruleCustom" :label-width="80">
    <FormItem label="姓名" prop="name">
      <Input type="text" v-model="formCustom.name"></Input>
    </FormItem>
    <FormItem label="电话" prop="phone">
      <Input type="text" v-model="formCustom.phone"></Input>
    </FormItem>
    <FormItem label="城市" prop="city">
      <Input type="text" v-model="formCustom.city"></Input>
    </FormItem>
    <FormItem label="学校" prop="school">
      <Input type="text" v-model="formCustom.school"></Input>
    </FormItem>
    <FormItem label="班级" prop="class">
      <Input type="text" v-model="formCustom.class"></Input>
    </FormItem>
    <FormItem label="年龄" prop="age">
      <Input type="text" v-model="formCustom.age"></Input>
    </FormItem>
    <FormItem label="企鹅qq" prop="qq">
      <Input type="text" v-model="formCustom.qq"></Input>
    </FormItem>
    <FormItem label="性别" prop="sex">
      <RadioGroup v-model="formCustom.sex">
        <Radio label="male">男</Radio>
        <Radio label="female">女</Radio>
      </RadioGroup>
    </FormItem>
    <FormItem>
      <Button type="primary" @click="handleSubmit('formCustom')">添加</Button>
      <!-- <router-link to="/">
        <Button>返回</Button>
      </router-link>-->
    </FormItem>
  </Form>
</template>
<script>
import jwtDecode from 'jwt-decode'
import min from "../api/min";
export default {
  data() {
    return {
      formCustom: {
        name: "",
        phone: "",
        city: "",
        school: "",
        age: "",
        sex: "",
        qq: "",
        class: ""
      },
      ruleCustom: {
        name: [
          {
            required: true,
            message: "姓名为空",
            trigger: "blur"
          }
        ],
        phone: [
          {
            required: true,
            message: "联系电话",
            trigger: "blur"
          }
        ],
        qq: [
          {
            required: true,
            message: "qq为空",
            trigger: "blur"
          }
        ],
        city: [
          {
            required: true,
            message: "居住城市",
            trigger: "blur"
          }
        ],
        sex: [
          {
            required: true,
            message: "请选择性别",
            trigger: "change"
          }
        ],
        school: [
          {
            required: true,
            message: "毕业学校",
            trigger: "blur"
          }
        ],
        class: [
          {
            required: true,
            message: "班级",
            trigger: "blur"
          }
        ],
        age: [
          {
            required: true,
            message: "年龄",
            trigger: "blur"
          }
        ]
      }
    };
  },
  methods: {
    handleSubmit(name) {
      let token = localStorage.getItem('token')
      let newtoken = jwtDecode(token)
      this.formCustom.uid = newtoken.id
      console.log(this.formCustom)
      this.$refs[name].validate(valid => {
        if (valid) {
          // this.$Message.success("success");
          min.post("/adm/particular/add", this.formCustom).then(resp => {
            if (resp.code == 200) {
              this.$Message.success(resp.msg);
              this.$router.push({ path: "/HomeCopy" });
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