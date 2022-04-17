<template>
  <div class="home">
    <nav>
      <router-link to="/home">主页</router-link> |
      <router-link to="/home/message">留言板</router-link> |
      <router-link to="/home/resume">简历</router-link>
    </nav>
    <router-view/>
  </div>
</template>

<script>
// @ is an alias to /src
import HelloWorld from '@/components/HelloWorld.vue'

export default {
  name: 'HomeView',
  components: {
    HelloWorld
  }
}
</script>

<style lang="less">
nav {
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
