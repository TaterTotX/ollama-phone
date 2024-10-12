<script setup>
import {ref, nextTick, onMounted, reactive, provide} from 'vue';
import BaseUI from "../components/baseUI.vue";
import user1path from "../assets/images/chat_ui_img/user1.jpg"
import user2path from "../assets/images/chat_ui_img/user2.png"
import iphonepath from "../assets/images/chat_ui_img/iphone-ts.png"
import chat_ground from "../assets/images/chat_ui_img/img.png";
import {Get_settings, Greet} from "../../wailsjs/go/main/App.js";

import Setting from "./setting.vue";
import {LogPrint} from "../../wailsjs/runtime/runtime.js";


const settings = reactive({
  model: '',
  prompt: '',
  welcomeMessages: '',
  title: '',
  input_text: '',
  speak_text: '',
  resultText: "",
  chat_title:"",
  user1path: user1path,
  user2path: user2path,
  chat_ground:chat_ground,
  show_settings: false,
})

onMounted(async () => {
  const data = await Get_settings()
  for (let key in data) {
    if (data[key].length > 4) {
      settings[key] = data[key]
    }

  }

  sendWelcomeMessage()


})



// 定义响应式变量
const messages = ref([]); // 存储聊天消息的数组
const newMessage = ref(""); // 存储新消息的变量
















//发送消息方法
async function greet() {
  const userInput = newMessage.value.trim();

  // 添加用户输入的消息
  messages.value.push({text: userInput, sender: 'user1'});

  // 等待Vue更新DOM后，滚动到底部
  await nextTick();
  scrollToBottom();


  // 清空输入
  newMessage.value = "";
  settings.chat_title  = settings.input_text
  // 执行异步问候操作，并等待结果
  settings.resultText= await Greet(userInput,"user");




  // 添加机器人回复的消息
  const sentences = settings.resultText.split('\n').filter(Boolean);


  const sendMessage = (index) => {
    if (index < sentences.length) {
      settings.chat_title = settings.input_text
      setTimeout(async () => {
        messages.value.push({text: sentences[index], sender: 'user2'});
        await nextTick();
        scrollToBottom();
        settings.chat_title = settings.title

        sendMessage(index + 1); // 递归调用，发送下一条消息
      }, Math.random() * 1500 ); // 随机间隔1到2秒发送下一条消息
    }


  };
  sendMessage(0);

  settings.chat_title  =settings.title
  // 再次等待Vue更新DOM后，滚动到底部
  await nextTick();
  scrollToBottom();
  settings.chat_title = settings.speak_text
  // await playWav(data.resultText)
  settings.chat_title = settings.title
}






// 新增一个函数，用于自动发送欢迎消息
const sendWelcomeMessage = () => {
  settings.chat_title=settings.title
  const welcomeMessages = settings.welcomeMessages.split(',');
  // 定义发送消息的函数
  const sendMessage = (index) => {
    if (index < welcomeMessages.length) {
      settings.chat_title = settings.input_text
      setTimeout(async () => {
        messages.value.push({text: welcomeMessages[index], sender: 'user2'});
        Greet(welcomeMessages[index], "assistant");
        // await playWav(welcomeMessages[index])
        settings.chat_title = settings.title
        scrollToBottom();
        sendMessage(index + 1); // 递归调用，发送下一条消息
      }, Math.random() * 1500 + 500); // 随机间隔1到2秒发送下一条消息
    }


  };

  // 开始发送消息，从第一条消息开始
  sendMessage(0);



};



// 滚动页面方法
const scrollToBottom = () => {
  // 获取 chatPage 元素
  const chatPage = document.getElementById('chat-page');
  if (chatPage) {
    // 如果 chatPage 存在，则设置滚动条到底部
    chatPage.scrollTop = chatPage.scrollHeight;
  } else {
    // 如果 chatPage 不存在，则输出错误信息到控制台
    LogPrint('滚动失败！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！');
  }
};



//播放语音
async function playWav(output) {
  if (output.length>5) {
    return new Promise(async (resolve, reject) => {
      const url = 'http://127.0.0.1:9880';
      const data = {
        refer_wav_path: 'vo_SGLQ103_5_yaeMiko_19.wav',
        prompt_text: '身体当然很重要啦，但我更想知道，你的心情平复了吗？',
        prompt_language: 'all_zh',
        text: output,
        text_language: 'zh',
        cut_punc: "，。,"
      };

      try {
        const response = await fetch(url, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify(data)
        });

        const audioData = await response.arrayBuffer();
        const audioContext = new AudioContext();
        const audioBuffer = await audioContext.decodeAudioData(audioData);
        const source = audioContext.createBufferSource();
        source.buffer = audioBuffer;
        source.connect(audioContext.destination);

        source.onended = () => {
          console.log("播放完成");
          resolve(); // 在这里解决Promise
        };

        source.start(0);
      } catch (error) {
        console.error("播放失败", error);
        reject(error); // 处理错误
        resolve();
      }
    });
  }
}

//发送通知信息
function send_notification_Message(text) {
  LogPrint("消息发送成功")
  messages.value.push({text: text, sender: 'user2'});
}

//上传图片方法
const changeAvatar = (avatar) => {

  const input = document.createElement('input');
  input.type = 'file';
  input.accept = 'image/*';
  input.onchange = (e) => {
    const file = e.target.files[0];
    const reader = new FileReader();
    reader.onload = (e) => {
      settings[avatar] = e.target.result;

    };
    reader.readAsDataURL(file);
  };
  input.click();
};
provide('changeAvatar', changeAvatar);
provide('settings',settings)
provide('send_notification_Message',send_notification_Message)



// const mainSettingPage = ref(false);

</script>

<template>

  <div class="container" style="background: transparent" >

    <img :src="iphonepath" alt="Descriptive text for the image" class="inserted-image">
    <div class="wrapper" :style="{ backgroundImage: `url(${settings.chat_ground})` }">
      <BaseUI :title="settings.chat_title" v-model:setting_page="settings.show_settings" />
      <div class="chat-page" id="chat-page">

        <setting  v-if="settings.show_settings"/>
        <div class="chats">
          <!-- 使用单个循环遍历消息数组，并根据消息的发送者设置不同样式 -->
          <div class="chat-user" v-for="(message, index) in messages" :key="index"
               :class="message.sender === 'user1' ? 'user1' : 'user2'">
            <div class="user-img">
              <!-- 根据消息的发送者显示不同的用户头像 -->
              <img v-if="message.sender === 'user1'" :src="settings.user1path" alt="">
              <img v-else :src="settings.user2path" @click="changeAvatar('user2path')"/>
            </div>
            <div class="user-msg">
              <!-- 使用v-if和v-else来根据消息的发送者添加不同的类名 -->
              <p v-if="message.sender === 'user1'" class="user1-msg">{{ message.text }}</p>
              <p v-else class="user2-msg">{{ message.text }}</p>
            </div>
          </div>
        </div>
      </div>
      <div class="chat-input">
        <img :src="settings.user1path" @click="changeAvatar('user1path')"/>
        <!-- 使用 v-model 双向绑定输入框内容 -->
        <input type="text" placeholder="..." class="my-input" v-model="newMessage"  @keyup.enter="greet" />

        <label for="txt" id="btn" @click="greet">Send</label>
      </div>
    </div>

  </div>




</template>

<style scoped>
/* scss */
.container {
  /* 设置容器宽度为100%，占满父元素宽度 */
  width: 100%;
  /* 设置容器高度为视口高度的100%，占满整个视口 */
  height: 100vh;
  /* 使用flex布局，使子元素在主轴和交叉轴上居中 */
  display: flex;
  justify-content: center;
  align-items: center;
  position: relative; /* 为绝对定位创建参考点 */
  z-index: 0; /* 确保.wrapper在图片之上 */

}

.wrapper {
  width: 82%; /* 或者设置为图片的宽度 */
  height: 93%; /* 保持图片的宽高比 */

  /* 设置包装器的边框圆角为5rem */
  border-radius: 3rem;
  /* 设置包装器的内边距为1.5rem在顶部和底部，2.5rem在左侧和右侧 */
  padding: 1.5rem 2.5rem;


  background-size: cover;
}

.inserted-image {
  position: absolute; /* 设置绝对定位 */
  top: 50%; /* 将图片的顶部定位到.wrapper的垂直中心 */
  left: 50%; /* 将图片的左侧定位到.wrapper的水平中心 */
  transform: translate(-50%, -50%); /* 调整图片位置，使其中心对齐 */
  z-index: 1; /* 确保图片在.wrapper内容之下 */
  /* 根据需要设置图片的宽度和高度 */
  width: 90%; /* 或者设置为图片的宽度 */
  height: auto; /* 保持图片的宽高比 */
  pointer-events: none; /* 使图片不响应鼠标事件 */
}

/* Chat Page Section */

.chat-page {

  position: fixed;

  height: 65%;
  width: 65%;
  overflow-y: scroll;
  margin-top: 5%;


}

.chat-page::-webkit-scrollbar {
  display: none;
}

.chat-user {
  margin-bottom: 1rem;
}

/* User 1 */

.user-img {
  display: block;
  width: 2.5rem;
  float: left;
}

.user-img img {
  width: 100%;
  border-radius: 50%;
}

.user-msg {
  display: inline-block;
  padding: 0 0 0 1rem;
  width: 82%;
  height: auto;

}

.user-msg p {
  width: 100%;
  background: var(--white);
  background-image: linear-gradient(to right, rgb(255, 255, 255), rgba(3, 3, 3, 0.47));

  border-radius: 0 1.4rem 1.4rem 1.4rem;
  font-size: 1rem;
  font-weight: 800;
  padding: 1rem;
  word-wrap: break-word;
  font-family: 'Comic Sans MS', cursive; /* 添加可爱字体 */
}

.user1-msg{
  color: #030303;
}

/* User 2 */





.user2 .user-img {
  float: right;

}

.user2 .user-msg p {
  color: var(--white);
  background: var(--purple);
  border-radius: 1.4rem 0 1.4rem 1.4rem;
  background-image: linear-gradient(to right, rgba(255,0,0,0), rgba(194, 190, 236, 0.51));
}



/* Chat Input Section */
.my-input{


  font-size: 1rem;
  font-weight: 800;
}
.chat-input {
  width: 65%;
  bottom: 8%;
  display: flex;
  position: fixed;
  justify-content: space-between;
  align-items: center;

}

.chat-input img {
  width: 3.5rem;
  border-radius: 50%;
  cursor: pointer;
}

.chat-input input {
  width: 100%;
  background: var(--bg-input);
  border: none;
  line-height: 3;
  border-radius: 3rem;
  margin: 0 1rem;
  padding: 0 1.5rem;
}

.chat-input input::placeholder {
  font-size: 1.2rem;
  font-weight: 500;
  color: var(--grey);
}

.chat-input label {
  font-size: 1.2rem;
  font-weight: 700;
  color: var(--purple);
  cursor: pointer;
}



</style>