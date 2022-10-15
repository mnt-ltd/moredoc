<template>
  <div class="com-form-group">
    <el-form label-position="top" label-width="80px" :model="group">
      <el-form-item
        label="名称"
        prop="title"
        :rules="[{ required: true, trigger: 'blur', message: '请输入名称' }]"
      >
        <el-input
          v-model="group.title"
          placeholder="请输入名称"
          clearable
        ></el-input>
      </el-form-item>
      <el-row :gutter="20">
        <el-col :span="12">
          <el-form-item label="排序(值越大越靠前)">
            <el-input-number
              v-model.number="group.sort"
              clearable
              :min="0"
              :step="1"
              placeholder="请输入排序值"
            ></el-input-number> </el-form-item
        ></el-col>
        <el-col :span="12">
          <el-form-item label="颜色">
            <el-input
              v-model="group.color"
              placeholder="请选择用户颜色"
              clearable
            ></el-input> </el-form-item
        ></el-col>
        <el-col :span="12">
          <el-form-item label="是否为默认用户组">
            <el-input
              v-model="group.is_default"
              placeholder="请选择用户颜色"
              clearable
            ></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="是否在用户名后展示">
            <el-input
              v-model="group.is_display"
              clearable
            ></el-input> </el-form-item
        ></el-col>
      </el-row>

      <el-form-item label="描述">
        <el-input
          v-model="group.description"
          type="textarea"
          rows="3"
          placeholder="请输入分组描述"
        ></el-input>
      </el-form-item>
      <el-form-item>
        <el-button
          type="primary"
          class="btn-block"
          icon="el-icon-check"
          @click="onSubmit"
          >提交</el-button
        >
      </el-form-item>
    </el-form>
  </div>
</template>
<script>
export default {
  name: 'FormGroup',
  props: {
    initGroup: {
      type: Object,
      default: () => {
        return {}
      },
    },
  },
  data() {
    return {
      group: {},
    }
  },
  created() {
    this.group = { ...this.group, ...this.initGroup }
  },
  methods: {
    onSubmit() {
      console.log(this.group)
    },
  },
}
</script>
