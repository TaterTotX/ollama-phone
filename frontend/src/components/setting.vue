<script setup lang="ts">
import {inject} from 'vue';
import {UpdateSetting} from "../../wailsjs/go/main/App.js";

type SettingsType = {
  model: string;
  prompt: string;
  welcomeMessages: string;
  title: string;
  input_text: string;
  speak_text: string;
  resultText: string;
  chat_title:string;
  user1path: string;
  user2path: string;
  chat_ground: string;
  show_settings:boolean;
};

const settings = inject('settings') as SettingsType;
const send_notification_Message = inject<(message: string) => void>('send_notification_Message');


const saveSettings = async () => {
  settings.show_settings=false
  send_notification_Message?.("Let's reboot and refresh our settings")
  const { resultText, chat_title,show_settings, ...restSettings } = settings;
  await UpdateSetting(restSettings);

}
</script>

<template>
  <div class="settings">
    <div v-for="(value, key) in settings" :key="key" class="typing" >
      <h2  v-if="key !== 'user1path' && key !== 'user2path' && key !== 'chat_ground' && key !== 'resultText' && key !== 'chat_title' && key !== 'show_settings'">{{ key }}</h2>
      <input v-model="settings[key]"  v-if="key !== 'user1path' && key !== 'user2path' && key !== 'chat_ground' && key !== 'resultText' && key !== 'chat_title'&& key !== 'show_settings'"/>
    </div>
    <button class="button" @click="saveSettings"> save </button>
  </div>
</template>


<style scoped>
.settings {
  position: fixed;
  color: #cccccc;
  margin-bottom: 1em;
  background-image: linear-gradient(to right, rgba(3, 3, 3, 0.79), rgba(3, 3, 3, 0.84));
  height: 65%;
  width: 65%;
  border-radius: 20px;  /* 添加圆角 */
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.5); /* 添加阴影效果 */

}




.typing {

  margin: 1em 1em;
  text-align: center;
}



input {
  padding: 0.5em;
  background-color: #030303;
  color: #cccccc;
  font-weight: 800;
  border: 1px solid #ccc;
  border-radius: 4px;
  transition: border-color 0.3s ease;
  text-align: center;
}

input:focus {
  border-color: rgba(3, 3, 3, 0.7);
  outline: none;
}

.button {
  display: block;
  margin-left: auto;
  margin-right: auto;
  padding: 10px 20px;
  margin-bottom: 30px;
  background-color: rgba(3, 3, 3, 0.7);
  color: white;
  border: none;

  cursor: pointer;
  transition: background-color 0.3s ease;
  border-radius: 20px;  /* 添加圆角 */
}

.button:hover {
  background-color: rgb(255, 255, 255);
  color: #030303;
}




</style>