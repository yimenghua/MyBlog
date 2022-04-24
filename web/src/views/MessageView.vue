<template>
  <div>
    <h1>留言板</h1>
    <div id="scrolldIV" class="slider">
        <ul class="list">
          <li v-for="(item,key) in messageList" :key="key">
            <p align="left">
              {{item.CreatedAt.slice(0,10)+"  "+item.CreatedAt.slice(11,19)+" "+item.UserName}}
            <br>
              {{item.Msg}}
            </p>
          </li>
        </ul>
    </div>
    名字：<input type="text" v-model="message.UserName">
    <br>
    内容：<input type="text" v-model="message.Msg">
    <br>
    <button @click="send">发送</button>
  </div>
</template>

<script>
import global from '@/global'
import axios from 'axios'
export default {
    name: "Message",
    data(){
        return{
            messageList : [],
            message: {
              "UserName" : "",
              "Msg" : ""
            }
        }
    },
    methods: {
      sliderbottom(){
        var div = document.getElementById('scrolldIV');
        div.scrollTop = div.scrollHeight;
      },
      send(){
        var that = this;
        if (that.message.UserName == "" || that.message.Msg == ""){
            console.log("名字或内容为空")
            return
        }
        var url = global.host+global.msgboardpath+"insert"
        axios.post(url, {
          UserName : that.message.UserName,
          Msg:that.message.Msg
        })
        .then(function (response) {
          console.log(response);
          that.get()
        })
        .catch(function (error) {
          console.log(error);
        });
      },
      get(){
        var that = this
        var url = global.host+global.msgboardpath+"select"
        axios.post(url, {
          UserName : "",
          Msg: ""
        })
        .then(function (response) {
          that.messageList = response.data.data
          setTimeout(() =>{
            that.sliderbottom()
          },500);
          console.log(response);
        })
        .catch(function (error) {
          console.log(error);
        });
      }
    },
    created () {
      this.get()
    },
}
</script>

<style lang="less">
  .slider{
    overflow: auto;
    height: 400px;
    widows: 100%;
    background:#EEEEEE;
  }
</style>