<template>
<!-- 主入口 -->
  <div id="app" style="bgImg" >
    <!-- <div id="nav">
      <router-link to="/">首页</router-link> |
      <router-link to="/regist">注册</router-link>
       <router-link to="/login"><button>登录</button></router-link>
    </div> -->
    <router-view />
  </div>
</template>

<style lang="less">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  
}

#nav {
  padding: 30px;

  a {
    font-weight: bold;
    color: #2c3e50;

    &.router-link-exact-active {
      color: #42b983;
    }
  }
}
</style>
