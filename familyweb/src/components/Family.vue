<template>
  <a-tabs class="tab" v-model:activeKey="activeKey">
    <a-tab-pane key="1" tab="家族信息" force-render>
      <a-button class="save" @click="keepPicture">保存族谱</a-button>
      <div class="body">
        <div class="content">
          <List class="content-list" />
          <div class="content-body" ref="zhupu">
            <div class="tag">
              <Tag />
            </div>
            <Node class="node" :formState="formState" />
          </div>
        </div>
      </div>
    </a-tab-pane>
    <a-tab-pane key="2" tab="人员查询">
      <Search />
    </a-tab-pane>
  </a-tabs>
</template>

<script setup lang="ts">
import { ref, onMounted, getCurrentInstance } from "vue";
import axios from "axios";
import { Member } from "../model/member";
import Node from "./Node.vue";
import List from "./List.vue";
import bus from "../libs/bus";
import Tag from "./Tag.vue";
import html2canvas from "html2canvas";

import Search from "./Search.vue";

const { proxy }: any = getCurrentInstance();

const formState = ref<Member>(new Member());
const activeKey = ref("1");
const selectLastMemberId = ref("");

bus.on("reloadMember", reloadMember as any);

function reloadMember(id: string = selectLastMemberId.value) {
  getMember(id);
}

function getMember(id: string = "") {
  formState.value = new Member();
  const formData = new FormData();
  formData.append("id", id);
  selectLastMemberId.value = id;
  axios
    .post("http://localhost:6200/v1/member/last", formData)
    .then((response) => {
      formState.value.getHeader(response.data.data);
      console.log("getMember start", formState.value);

      if (!(response.data.data as []).length) {
      }
    });
}

//保存图片
function keepPicture() {
  html2canvas(proxy.$refs.zhupu as HTMLElement).then((canvas) => {
    // 获取生成的图片的url,并设置为png格式
    const imgUrl = canvas.toDataURL("png");
    let creatDom = document.createElement("a");
    document.body.appendChild(creatDom);
    creatDom.href = imgUrl;
    //设置生成图片的名称
    creatDom.download = "zhupu";
    creatDom.click();
  });
}

onMounted(() => {
  getMember();
});
</script>

<style scoped>
.body {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
}
.node {
  display: flex;
  /* justify-content: center; */
  align-items: center;
  width: 100%;
  flex: 3;
}

.save {
  position: absolute;
  height: 100px;
  width: 100px;
  bottom: 2%;
  right: 2%;
}

.content {
  display: flex;
  flex-direction: row;
  width: 100%;
}

.content-list {
  width: 200px;
}
.content-body {
  flex: 1;
}

.tab {
  padding: 0 50px;
}

.tag {
  position: absolute;
  top: 10%;
  right: 5%;
}
</style>
