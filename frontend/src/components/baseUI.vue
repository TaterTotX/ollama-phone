<script setup>
import {defineProps, provide, ref} from 'vue';
import Loader from "../components/Loader.vue";
import { inject } from 'vue';

const changeAvatar = inject('changeAvatar');

const props = defineProps({
  title: String, // 定义一个名为title的prop，类型为String，
  setting_page: Boolean,
});



const emits = defineEmits(['update:setting_page']);

function showSettings() {
  // 更新 props 并触发事件
  emits('update:setting_page', !props.setting_page);
}
</script>

<template>

  <div class="topbar" >
    <div class="topbar-left"></div>
    <div class="topbar-right">

      <Loader class="loader"  />

      <span class="material-symbols-outlined" style="color: #030303; transform: scale(0.8);"> wifi </span>
      <span class="material-symbols-outlined" style="color: #030303; transform: scale(0.8);" > battery_horiz_050 </span>

    </div>
  </div>
  <div class="header">


    <span class="material-symbols-outlined" @click="changeAvatar('chat_ground')"> 🍭</span>
    <p>{{props.title}}</p>

    <span class="material-symbols-outlined" @click="showSettings"> 🍩 </span>

  </div>
</template>

<style scoped>
/* Header Section */


@import url('https://fonts.googleapis.com/css2?family=Montserrat:wght@400;500;600;700&display=swap');
@import url('https://fonts.googleapis.com/css2?family=Material+Symbols+Outlined:opsz,wght,FILL,GRAD@48,400,0,0');




.header {
  width: 100%;
  height: 2rem;
  /* background: green; */
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.header p {
  font-size: 1.2rem;
  font-weight: 600;
  color: var(--grey);
}

.header span {
  font-size: 1.6rem;
}



/* Topbar Section */

.topbar {
  width: 90%;
  height: 3rem;
  /* background: red; */
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.topbar-left {
  font-size: 1.4rem;
  font-weight: 600;
}

.topbar-left,
.topbar-right {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 0.5rem;


  margin-left: 85px;
  margin-bottom: 42px;
}

.topbar-right span {
  font-size: 1.8rem;
}

.topbar-right span:nth-child(2) {
  font-size: 1.6rem;
}

.loader {
  transform: scale(0.25) translate(90px, -15px); /* 缩小并向上和向右移动 */
  transform-origin: center; /* 定义了缩放和旋转的基点，默认是中心 */
}






</style>