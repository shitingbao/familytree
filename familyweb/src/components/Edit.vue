<template>
  <a-form :model="formState" v-bind="layout" name="nest-messages">
    <a-form-item
      :name="['user', 'name']"
      label="姓名"
      :rules="[{ required: true }]"
    >
      <a-input v-model:value="formState.name" />
    </a-form-item>

    <a-form-item label="出生日期">
      <a-input v-model:value="formState.dateBirth" />
    </a-form-item>
    <a-form-item label="结婚日期">
      <a-input v-model:value="formState.dateMarry" />
    </a-form-item>
    <a-form-item label="出生地">
      <a-input v-model:value="formState.placeBirth" />
    </a-form-item>
    <a-form-item label="死亡时间">
      <a-input v-model:value="formState.dateDeath" />
    </a-form-item>
    <a-form-item label="死亡地点">
      <a-input v-model:value="formState.placeDeath" />
    </a-form-item>

    <a-form-item label="生平简介">
      <a-textarea v-model:value="formState.content" />
    </a-form-item>
    <a-form-item label="荣誉">
      <a-textarea v-model:value="formState.honor" />
    </a-form-item>
    <div class="foot">
      <a-button class="btn" type="primary" @click="updateMemeber">
        提交修改
      </a-button>
      <a-button class="btn" type="primary" @click="addChild">
        增加子节点
      </a-button>
      <a-button type="primary" @click="deleteMember">删除节点</a-button>
    </div>
  </a-form>
</template>

<script setup lang="ts">
import { ref, onMounted, watch, PropType } from "vue";
import { ManOutlined, WomanOutlined } from "@ant-design/icons-vue";
import axios from "axios";
import { Member } from "../model/member";
import qs from "qs";

const props = defineProps({
  row: {
    type: Object as PropType<Member>,
    default: () => {
      return new Member();
    },
  },
});

const formState = ref<Member>(new Member());

const visible = ref<boolean>(false);

const layout = {
  labelCol: { span: 8 },
  wrapperCol: { span: 12 },
};

const emit = defineEmits(["visible"]);

function updateMemeber() {
  axios({
    method: "POST",
    url: "http://localhost:6200/v1/member/update",
    data: qs.stringify(formState.value),
    headers: {
      "content-type": "application/x-www-form-urlencoded",
    },
  }).then((response) => {
    console.log(response);
    // formState.value.getHeader(response.data.data);
    // console.log(" response.data response.data", response.data);
    emit("visible");
  });
}

function addChild() {
  const formData = new FormData();
  formData.append("name", "son");
  formData.append("parentId", formState.value.id + "");

  // 新的路径等于父节点路径加上父节点id，注意第一个节点为空
  const newPath = formState.value.path == "" ? "" : formState.value.path + ",";
  formData.append("path", newPath + formState.value.id);
  axios
    .post("http://localhost:6200/v1/member/create", formData)
    .then((response) => {
      // formState.value.getHeader(response.data.data);
      console.log(response);
      emit("visible");
    });
}

function deleteMember() {
  const formData = new FormData();
  formData.append("id", formState.value.id + "");

  axios
    .post("http://localhost:6200/v1/member/delete", formData)
    .then((response) => {
      console.log(response);
      // formState.value.getHeader(response.data.data);
      emit("visible");
    });
}

watch(
  () => props.row.id,
  () => {
    formState.value.initData(props.row);
  },
  { deep: true, immediate: true }
);
</script>

<style scoped>
.foot {
  display: flex;
  width: 100%;
  justify-content: center;
  align-items: center;
}

.btn {
  margin-right: 10px;
}
</style>
