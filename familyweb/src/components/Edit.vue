<template>
  <a-form :model="formState" v-bind="layout" name="nest-messages">
    <a-form-item label="姓名">
      <a-input v-model:value="formState.name" />
    </a-form-item>
    <a-form-item label="家族所属">
      <a-input disabled="disabled" v-model:value="formState.familySimple" />
    </a-form-item>
    <a-form-item label="性别">
      <a-switch
        v-model:checked="formState.sex"
        :checkedValue="1"
        :unCheckedValue="2"
        checked-children="男"
        un-checked-children="女"
      />
    </a-form-item>

    <a-form-item label="出生日期">
      <a-space direction="vertical" :size="12">
        <a-date-picker v-model:value="dateBirth" @change="selectDate" />
      </a-space>
    </a-form-item>
    <a-form-item label="结婚日期">
      <a-space direction="vertical" :size="12">
        <a-date-picker v-model:value="dateMarry" />
      </a-space>
    </a-form-item>
    <a-form-item label="出生地">
      <a-input v-model:value="formState.placeBirth" />
    </a-form-item>
    <a-form-item label="死亡时间">
      <a-space direction="vertical" :size="12">
        <a-date-picker v-model:value="dateDeath" />
      </a-space>
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
import dayjs, { Dayjs } from "dayjs";
import bus from "../libs/bus";

const props = defineProps({
  row: {
    type: Object as PropType<Member>,
    default: () => {
      return new Member();
    },
  },
});

const formState = ref<Member>(new Member());

const dateBirth = ref<Dayjs>(dayjs("0000-01-01", "YYYY-MM-DD"));
const dateMarry = ref<Dayjs>(dayjs("0000-01-01", "YYYY-MM-DD"));
const dateDeath = ref<Dayjs>(dayjs("0000-01-01", "YYYY-MM-DD"));

const visible = ref<boolean>(false);

const layout = {
  labelCol: { span: 8 },
  wrapperCol: { span: 12 },
};

const emit = defineEmits(["visible"]);

function selectDate() {
  formState.value.dateBirth = dateBirth.value.format("YYYY-MM-DD");
  formState.value.dateMarry = dateMarry.value.format("YYYY-MM-DD");
  formState.value.dateDeath = dateDeath.value.format("YYYY-MM-DD");
}

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
    if (formState.value.parentId == 0) {
      bus.emit("reloadList");
    }
    // formState.value.getHeader(response.data.data);
    // console.log(" response.data response.data", response.data);
    emit("visible");
  });
}

function addChild() {
  const formData = new FormData();
  formData.append("name", "son");
  formData.append("familySimple", formState.value.familySimple);
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
      // 删除根节点通知列表刷新
      if (formState.value.parentId == 0) {
        bus.emit("reloadList");
        bus.emit("reloadMember", "");
      }
    });
}

watch(
  () => props.row.id,
  () => {
    formState.value.initData(props.row);
    dateBirth.value = dayjs(formState.value.dateBirth, "YYYY-MM-DD");
    dateMarry.value = dayjs(formState.value.dateMarry, "YYYY-MM-DD");
    dateDeath.value = dayjs(formState.value.dateDeath, "YYYY-MM-DD");
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
