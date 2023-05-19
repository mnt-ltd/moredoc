<template>
  <div class="com-table-list">
    <el-table
      v-loading="loading"
      :data="tableData"
      style="width: 100%"
      row-key="id"
      :tree-props="treeProps"
      @selection-change="selectRow"
    >
      <el-table-column
        v-if="showSelect"
        type="selection"
        width="55"
        :selectable="selectable"
      >
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
          <span v-else-if="item.type === 'bytes'">
            {{ formatBytes(scope.row[item.prop]) }}
          </span>
          <!-- 枚举，键为数字 -->
          <span v-else-if="item.type === 'enum'">
            <el-tag
              v-if="item.enum[scope.row[item.prop] || 0]"
              :type="item.enum[scope.row[item.prop] || 0].type || 'info'"
              :effect="item.enum[scope.row[item.prop] || 0].effect || 'dark'"
            >
              {{ item.enum[scope.row[item.prop] || 0].label }}
            </el-tag>
            <span v-else>-</span>
          </span>
          <span v-else-if="item.type === 'datetime'">
            {{ formatDatetime(scope.row[item.prop]) || '0000-00-00 00:00:00' }}
          </span>
          <span v-else-if="item.type === 'color'">
            <span :style="{ color: scope.row[item.prop] }">{{
              scope.row[item.prop] || '-'
            }}</span>
          </span>
          <span v-else-if="['link', 'url'].includes(item.type)">
            <a :href="scope.row[item.prop]" target="_blank">
              <i class="el-icon-link"></i> {{ scope.row[item.prop] }}</a
            >
          </span>
          <span v-else-if="item.type === 'image'">
            <!-- 因为table cell有个左右的10px内边距，所以需要调整下 -->
            <UploadImage
              v-if="scope.row[item.prop]"
              :disabled="true"
              :image="scope.row[item.prop]"
              :width="item.width ? item.width + 'px' : 'auto'"
              style="margin-left: -10px; margin-right: -10px"
            />
            <span v-else>-</span>
          </span>
          <span v-else-if="item.type === 'array'">
            <template v-if="scope.row[item.prop]">
              <el-tag
                v-for="(value, idx) in scope.row[item.prop]"
                :key="item.prop + idx"
                class="mgr-5px"
                >{{ value }}</el-tag
              >
            </template>
            <template v-else>-</template>
          </span>
          <!-- 有层级的，用breadcrumb -->
          <span v-else-if="item.type === 'breadcrumb'">
            <template v-if="scope.row[item.prop]">
              <el-breadcrumb separator-class="el-icon-arrow-right">
                <el-breadcrumb-item
                  v-for="(value, idx) in scope.row[item.prop]"
                  :key="item.prop + idx"
                  >{{ value }}</el-breadcrumb-item
                >
              </el-breadcrumb>
            </template>
            <template v-else>-</template>
          </span>
          <!-- 字符串。更多，则需要继续扩展 -->
          <span v-else-if="item.type==='html'">
            <span v-html="scope.row[item.prop]"></span>
          </span>
          <span v-else>{{ scope.row[item.prop] || '-' }}</span>
        </template>
      </el-table-column>
      <el-table-column
        v-if="showActions || showView || showDelete || showEdit"
        fixed="right"
        label="操作"
        :min-width="actionsMinWidth"
        class="com-table-list-actions"
      >
        <template slot-scope="scope">
          <slot :row="scope.row" name="actions"></slot>
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
            @click="editRow(scope.row)"
            >编辑</el-button
          >
          <el-button
            v-if="showDelete"
            type="text"
            size="small"
            icon="el-icon-delete"
            :disabled="scope.row.disable_delete"
            @click="deleteRow(scope.row)"
            class="text-danger"
            >删除</el-button
          >
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>
<script>
import UploadImage from './UploadImage.vue'
import { formatDatetime, formatBytes } from '~/utils/utils'
export default {
  name: 'ComTableList',
  components: { UploadImage },
  props: {
    tableData: {
      type: Array,
      default: () => [],
    },
    loading: {
      type: Boolean,
      default: false,
    },
    // 每个field结构，包括：label、prop、width、min-width、fixed，还有type，其中type枚举包括：text、image、link、button、slot
    fields: {
      type: Array,
      default: () => [],
    },
    treeProps: {
      type: Object,
      default: () => ({}),
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
    formatBytes,
    viewRow(row) {
      this.$emit('viewRow', row)
    },
    editRow(row) {
      this.$emit('editRow', row)
    },
    deleteRow(row) {
      this.$emit('deleteRow', row)
    },
    selectRow(rows) {
      this.$emit('selectRow', rows)
    },
    selectable(row) {
      // 取反，禁止删除的行，不可选中
      return !row.disable_delete
    },
  },
}
</script>
<style lang="scss">
.com-table-list {
  .el-button {
    margin-left: 0;
    margin-right: 10px;
  }
  .el-breadcrumb__separator[class*='icon'] {
    margin: 0;
  }
}
</style>
