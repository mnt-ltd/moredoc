<template>
  <div class="com-table-list">
    <el-table
      :data="tableData"
      style="width: 100%"
      @selection-change="selectRow"
    >
      <el-table-column v-if="showSelect" type="selection" width="55">
      </el-table-column>
      <el-table-column
        v-for="item in fields"
        :key="'field-' + item.prop"
        :prop="item.prop"
        :label="item.label"
        :width="item.width"
        :min-width="item.minWidth"
        :fixed="item.fixed"
      >
        <template slot-scope="scope">
          <!-- 头像 -->
          <el-avatar
            v-if="item.type === 'avatar'"
            :size="45"
            :src="scope.row[item.prop]"
          >
            <img src="/static/images/blank.png" />
          </el-avatar>
          <!-- 数字 -->
          <span v-else-if="item.type === 'number'">{{
            scope.row[item.prop] || '0'
          }}</span>
          <el-tag
            v-else-if="item.type === 'bool'"
            :type="scope.row[item.prop] ? 'success' : 'danger'"
            effect="dark"
          >
            {{ scope.row[item.prop] ? '是' : '否' }}</el-tag
          >
          <span v-else-if="item.type === 'datetime'">
            {{ formatDatetime(scope.row[item.prop]) || '0000-00-00 00:00:00' }}
          </span>
          <span v-else-if="item.type === 'color'">
            <span :style="{ color: scope.row[item.prop] }">{{
              scope.row[item.prop] || '-'
            }}</span>
          </span>
          <!-- 字符串。更多，则需要继续扩展 -->
          <span v-else>{{ scope.row[item.prop] || '-' }}</span>
        </template>
      </el-table-column>
      <el-table-column
        v-if="showActions || showView || showDelete || showEdit"
        fixed="right"
        label="操作"
        :min-width="actionsMinWidth"
      >
        <template slot-scope="scope">
          <el-button
            v-if="showView"
            type="text"
            size="small"
            icon="el-icon-view"
            @click="viewRow(scope.row)"
            >查看</el-button
          >
          <el-button
            v-if="showEdit"
            type="text"
            size="small"
            icon="el-icon-edit"
            @click="editRow(row)"
            >编辑</el-button
          >
          <el-button
            v-if="showDelete"
            type="text"
            size="small"
            icon="el-icon-delete"
            @click="deleteRow(scope.row)"
            >删除</el-button
          >
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>
<script>
import { formatDatetime } from '~/utils/utils'
export default {
  name: 'ComTableList',
  props: {
    tableData: {
      type: Array,
      default: () => [],
    },
    // 每个field结构，包括：label、prop、width、min-width、fixed，还有type，其中type枚举包括：text、image、link、button、slot
    fields: {
      type: Array,
      default: () => [],
    },
    actionsMinWidth: {
      type: Number,
      default: 180,
    },
    showActions: {
      type: Boolean,
      default: true,
    },
    showView: {
      type: Boolean,
      default: true,
    },
    showEdit: {
      type: Boolean,
      default: true,
    },
    showDelete: {
      type: Boolean,
      default: true,
    },
    showSelect: {
      type: Boolean,
      default: true,
    },
  },
  methods: {
    formatDatetime,
    viewRow(row) {
      this.$emit('viewRow', row)
    },
    editRow(row) {
      this.$emit('editRow', row)
    },
    deleteRow(row) {
      this.$emit('deleteRoW', row)
    },
    selectRow(val) {
      this.$emit('selectRow', val)
    },
  },
}
</script>
