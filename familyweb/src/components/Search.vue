<template>
  <div class="search">
    <a-input-search
      v-model:value="memberVal"
      placeholder="请输入名称或者家族标示"
      :loading="loadVisible"
      enter-button
      @search="getMember"
    />

    <a-table
      style="width: 1100px"
      :columns="columns"
      :data-source="data"
      :scroll="{ x: 1500, y: 300 }"
    >
    </a-table>
  </div>
</template>

<script setup lang="ts">
import { ref, Ref, watch, PropType } from "vue";
import axios from "axios";
import { Member } from "../model/member";
import type { TableColumnsType } from "ant-design-vue";

// const data: Member[] = [];

const data: Ref<Member[]> = ref([]);

const loadVisible = ref<boolean>(false);

const columns: TableColumnsType = [
  {
    title: "name",
    width: 100,
    dataIndex: "name",
    fixed: "left",
  },
  {
    title: "家族所属",
    width: 100,
    dataIndex: "familySimple",
    fixed: "left",
  },
  { title: "性别", dataIndex: "sex", width: 150 },
  { title: "出生日期", dataIndex: "dateBirth", width: 150 },
  { title: "结婚日期", dataIndex: "dateMarry", width: 150 },
  { title: "出生地", dataIndex: "placeBirth", width: 150 },
  { title: "死亡时间", dataIndex: "dateDeath", width: 150 },
  { title: "死亡地点", dataIndex: "placeDeath", width: 150 },
  { title: "生平简介", dataIndex: "content", width: 150 },
  { title: "荣誉", dataIndex: "honor" },
];

const memberVal = ref("");

function getMember() {
  data.value = [];

  loadVisible.value = true;
  const formData = new FormData();
  formData.append("name", memberVal.value);
  formData.append("familySimple", memberVal.value);
  axios
    .post("http://localhost:6200/v1/member/search", formData)
    .then((response) => {
      response.data.data.forEach((item: any) => {
        var formState = new Member();
        formState.initData(item);

        data.value.push(formState);
        console.log("item:", data);
      });
    })
    .finally(() => {
      loadVisible.value = false;
    });
}
</script>

<style lang="less" scoped>
.table {
  width: 1500px;
}
.search {
  width: 400px;
}
</style>
