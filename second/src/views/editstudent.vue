<template>
  <Form ref="showData" :model="showData" :rules="ruleCustom" :label-width="80">
    <FormItem label="姓名" prop="name">
      <Input type="text" readonly disabled v-model="showData.name"></Input>
    </FormItem>
    
    <FormItem label="电话" prop="phone">
      <Input type="text" v-model="showData.phone"></Input>
    </FormItem>
    <FormItem label="城市" prop="city">
      <Input type="text" v-model="showData.city"></Input>
    </FormItem>
    <FormItem label="学校" prop="school">
      <Input type="text" v-model="showData.school"></Input>
    </FormItem>
    <FormItem label="班级" prop="class">
      <Input type="text" v-model="showData.class"></Input>
    </FormItem>
    <FormItem label="年龄" prop="age">
      <Input type="text" v-model="showData.age"></Input>
    </FormItem>
    <FormItem label="企鹅qq" prop="qq">
      <Input type="text" v-model="showData.qq"></Input>
    </FormItem>
    <FormItem label="性别" prop="sex">
      <RadioGroup  v-model="showData.sex">
        <Radio  label="male">男</Radio>
        <Radio  label="female">女</Radio>
      </RadioGroup>
      <!-- readonly disabled不允许更改 -->
    </FormItem>
    <FormItem>
      <Button type="primary" @click="handleSubmit('showData')">编辑</Button>
        <Button @click="handleReset('showData')" style="margin-left: 8px">重置</Button>
    </FormItem>
  </Form>
</template>
<script>
import min from "../api/min";
export default {
  data() {
    return {
      getData: {}, // 获取信息副本
      showData: {}, // 显示信息
      ruleCustom: {
        name: [
        
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
          },

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
        ],
      }
    };
  },
  mounted() {
        this.list()
    },
  methods: {
     list() {
         var urlStr = location.href.match('[^/]+(?!.*/)')[0];
    min.get("/adm/particular/get/"+urlStr).then(resp => {
        // console.log(resp);
        if (resp.code == 200) {
           this.showData = resp.data
           this.getData = JSON.parse(JSON.stringify(resp.data))
           console.log(this.getData)
        }
      });
    },
      handleSubmit (name) {
                this.$refs[name].validate((valid) => {
                    if (valid) {
                       min.post("adm/particular/edit",this.showData).then(resp=>{
                           if (resp.code==200){
                                this.$Message.success(resp.msg);
                                this.$router.push({path:'/HomeCopy'})
                           }else{
                               this.$Message.error(resp.msg);
                           }
                       })
                    }
                })
            },
    handleReset (name) {
        console.log(this.showData, this.getData)
        this.showData = JSON.parse(JSON.stringify(this.getData))
    },
  }
};
</script>