<template>
  <a-form
    :model="formState"
    v-bind="layout"
    name="nest-messages"
    :validate-messages="validateMessages"
    @finish="onFinish"
  >
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
  </a-form>
</template>

<script setup lang="ts">
import { ref, onMounted, watch, PropType } from "vue";
import { ManOutlined, WomanOutlined } from "@ant-design/icons-vue";
import axios from "axios";
import { Member } from "../model/member";

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

const validateMessages = {
  required: "${label} is required!",
  types: {
    email: "${label} is not a valid email!",
    number: "${label} is not a valid number!",
  },
  number: {
    range: "${label} must be between ${min} and ${max}",
  },
};

const onFinish = (values: any) => {
  console.log("Success:", values);
};

const showModal = () => {
  visible.value = true;
};

const handleOk = (e: MouseEvent) => {
  console.log(e);
  visible.value = false;
};

watch(
  () => props.row.id,
  () => {
    formState.value.initData(props.row);
  },
  { deep: true, immediate: true }
);
</script>

<style scoped>
a {
  color: #42b983;
}

label {
  margin: 0 0.5em;
  font-weight: bold;
}

code {
  background-color: #eee;
  padding: 2px 4px;
  border-radius: 4px;
  color: #304455;
}
</style>
