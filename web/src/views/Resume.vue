<template>
    <h3>简历</h3>
    <a :href="resumePDFLink">简历下载</a>
    <br>
    <img :src="resumePic" class="resumePic">
</template>


<script>
import resumePic  from '../assets/resume.jpg'
import global_msg from '../global.js'
export default {
    name: "Resume",
    data(){
        return{
            resumePic : resumePic,
            resumePDFLink: global_msg.host + "/files/resume.pdf"
        }
    },
}
</script>

<style lang="less">
    .resumePic{
        width: 80%;
        border:1px solid black;
    }
</style>