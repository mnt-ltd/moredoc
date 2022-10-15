<template>
  <div v-if="fields.length > 0" class="com-form-search">
    <el-form :inline="true" :model="search">
      <el-form-item
        v-for="field in fields"
        :key="field.name"
        :label="field.label"
      >
        <el-input
          v-if="field.type == 'text'"
          v-model="search[field.name]"
          :placeholder="field.placeholder"
          clearable
          @keydown.native.enter="onSearch"
        ></el-input>
        <el-select
          v-else-if="field.type == 'select'"
          v-model="search[field.name]"
          :placeholder="field.placeholder"
          :multiple="field.multiple"
          clearable
          filterable
        >
          <el-option
            v-for="item in field.options"
            :key="field.name + '_option_' + item.value"
            :label="item.label"
            :value="item.value"
          ></el-option>
        </el-select>
      </el-form-item>

      <el-form-item>
        <el-button
          type="primary"
          icon="el-icon-search"
          :loading="loading"
          @click="onSearch"
          >查询</el-button
        >
      </el-form-item>
      <el-form-item v-if="showCreate">
        <el-button type="primary" icon="el-icon-plus" @click="onCreate"
          >新增</el-button
        >
      </el-form-item>
      <el-form-item v-if="showDelete">
        <el-button
          type="danger"
          icon="el-icon-delete"
          :disabled="disabledDelete"
          @click="onDelete"
          >批量删除</el-button
        >
      </el-form-item>
    </el-form>
  </div>
</template>
<script>
export default {
  name: 'FormSearch',
  props: {
    loading: {
      type: Boolean,
      default: false,
    },

    showCreate: {
      type: Boolean,
      default: true,
    },
    showDelete: {
      type: Boolean,
      default: true,
    },
    disabledDelete: {
      type: Boolean,
      default: true,
    },
    fields: {
      type: Array,
      default: () => {
        return []
      },
    },
  },
  data() {
    return {
      search: {},
    }
  },
  methods: {
    onSearch() {
      this.$emit('onSearch', this.search)
    },
    onCreate() {
      this.$emit('onCreate')
    },
    onDelete() {
      this.$emit('onDelete')
    },
  },
}
</script>
