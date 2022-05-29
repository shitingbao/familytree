<template>
  <div>
    <h1>家族列表</h1>
    <a-button
      class="editable-add-btn"
      style="margin-bottom: 8px"
      @click="handOK"
    >
      添加一个新家族
    </a-button>
    <a-modal
      v-model:visible="visible"
      title="新家族标示以及祖宗姓名"
      width="250px"
      :footer="null"
    >
      <a-form :model="formState" v-bind="layout" name="nest-messages">
        <a-form-item label="姓名">
          <a-input v-model:value="formState.name" />
        </a-form-item>
        <a-form-item label="家族标志">
          <a-input v-model:value="formState.familySimple" />
        </a-form-item>
        <div class="foot">
          <a-button class="btn" type="primary" @click="handleAdd">
            添加新家族
          </a-button>
          <a-button type="primary" class="cancle" @click="cancle"
            >取消</a-button
          >
        </div>
      </a-form>
    </a-modal>

    <a-table
      class="list"
      bordered
      :data-source="dataSource"
      :columns="columns"
      :pagination="{ pageSize: 3 }"
    >
      <template #bodyCell="{ column, text, record }">
        <template v-if="column.dataIndex === 'name'">
          <div class="editable-cell">
            <div class="editable-cell-text-wrapper">
              {{ text || " " }}
            </div>
          </div>
        </template>
        <template v-else-if="column.dataIndex === 'familySimple'">
          <div class="editable-cell">
            <div class="editable-cell-text-wrapper">
              {{ text || " " }}
            </div>
          </div>
        </template>
        <template v-else-if="column.dataIndex === 'operation'">
          <span class="right">
            <a @click="onDelete(record.memberId)">Delete</a>
          </span>
          <span>
            <a @click="showMemberList(record.memberId)">show</a>
          </span>
        </template>
      </template>
    </a-table>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, Ref } from "vue";
import axios from "axios";
import { Member } from "../model/member";

import bus from "../libs/bus";

interface DataItem {
  memberId: string;
  name: string;
  familySimple: string;
}

const layout = {
  labelCol: { span: 8 },
  wrapperCol: { span: 12 },
};

bus.on("reloadList", MemberList);

const formState = ref<Member>({
  name: "newpeople",
  familySimple: "new family simple",
} as Member); // 弹出框信息

const visible = ref(false);

const columns = [
  {
    title: "祖宗名讳",
    dataIndex: "name",
    width: "30%",
  },
  {
    title: "家族标示",
    dataIndex: "familySimple",
    width: "30%",
  },
  {
    title: "operation",
    dataIndex: "operation",
  },
];

const dataSource: Ref<DataItem[]> = ref([]);

function showMemberList(memberId: string) {
  bus.emit("reloadMember", memberId);
}

function cancle() {
  visible.value = false;
}

function handOK() {
  visible.value = true;
}

function onDelete(memberId: string) {
  const formData = new FormData();
  formData.append("id", memberId);

  axios
    .post("http://localhost:6200/v1/member/delete", formData)
    .then((response) => {
      console.log(response);
      MemberList();
      bus.emit("reloadMember", ""); // 必须传一个空，为了刷新列表
    });
}

function handleAdd() {
  const formData = new FormData();
  formData.append("name", formState.value.name);
  formData.append("familySimple", formState.value.familySimple);
  axios
    .post("http://localhost:6200/v1/member/create", formData)
    .then((response) => {
      MemberList();
      visible.value = false;
      console.log(response);
    });
}

function MemberList() {
  dataSource.value = [];
  const formData = new FormData();
  formData.append("id", "0");
  axios
    .post("http://localhost:6200/v1/member/list", formData)
    .then((response) => {
      response.data.data.forEach((e: any) => {
        dataSource.value.push({
          memberId: e.id,
          name: e.name,
          familySimple: e.familySimple,
        });
      });
      console.log("list:", response.data);
    });
}

onMounted(() => {
  MemberList();
});
</script>

<style scoped>
.right {
  padding-right: 10px;
}

.cancle {
  margin-left: 10px;
}
</style>
