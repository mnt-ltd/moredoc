<template>
  <div class="com-form-config">
    <el-form
      ref="formConfig"
      label-position="top"
      label-width="80px"
      :model="configs"
    >
      <el-form-item
        v-for="(item, index) in configs"
        :key="'cfg-' + item.id"
        :label="item.label + '（' + item.placeholder + '）'"
      >
        <el-input-number
          v-if="item.input_type === 'number'"
          v-model="configs[index]['value']"
          clearable
          :min="0"
          :placeholder="item.placeholder"
          :step="1"
        ></el-input-number>
        <el-input
          v-else-if="item.input_type === 'textarea'"
          v-model="configs[index]['value']"
          type="textarea"
          :placeholder="item.placeholder"
          rows="5"
        ></el-input>
        <el-select
          v-else-if="item.input_type === 'select'"
          v-model="configs[index]['value']"
        >
          <el-option
            v-for="option in item.options.split('\n')"
            :key="'option-' + option"
            :label="option.split(':')[1]"
            :value="option.split(':')[0]"
          ></el-option>
        </el-select>
        <el-select
          v-else-if="item.input_type === 'select-multi'"
          v-model="configs[index]['value']"
          multiple
          clearable
        >
          <el-option
            v-for="option in item.options.split('\n')"
            :key="'option-' + option"
            :label="option.split(':')[1]"
            :value="option.split(':')[0]"
          ></el-option>
        </el-select>
        <el-switch
          v-else-if="item.input_type === 'switch'"
          v-model="configs[index]['value']"
          active-color="#13ce66"
          inactive-color="#ff4949"
          active-text="是"
          inactive-text="否"
          :active-value="'true'"
          :inactive-value="'false'"
        >
        </el-switch>
        <UploadImage
          v-else-if="item.input_type === 'image'"
          :action="'/api/v1/upload/config'"
          :image="configs[index]['value']"
          :width="'200px'"
          @success="success($event, index)"
        />
        <el-input
          v-else
          v-model="configs[index]['value']"
          :placeholder="item.placeholder"
          clearable
        ></el-input>
      </el-form-item>
      <el-form-item>
        <el-button
          type="primary"
          icon="el-icon-check"
          :loading="loading"
          @click="onSubmit"
          >提交</el-button
        >
      </el-form-item>
    </el-form>
  </div>
</template>
<script>
import { updateConfig } from '~/api/config'
export default {
  name: 'FormConfig',
  props: {
    initConfigs: {
      type: Array,
      default: () => {
        return []
      },
    },
  },
  data() {
    return {
      loading: false,
      // 转成对象的方式来处理，解决采用数组方式，数据不响应的问题
      configs: {},
    }
  },
  watch: {
    initConfigs: {
      handler(val) {
        let configs = { ...val }
        Object.values(configs).forEach((item) => {
          if (item.input_type === 'select-multi') {
            try {
              item.value = item.value.split(',')
            } catch (error) {}
          }
        })
        this.configs = configs
      },
      immediate: true,
    },
  },
  created() {
    let configs = { ...this.initConfigs }
    Object.values(configs).forEach((item) => {
      if (item.input_type === 'select-multi') {
        try {
          item.value = item.value.split(',')
        } catch (error) {}
      }
    })
    this.configs = configs
  },
  methods: {
    async onSubmit() {
      this.loading = true
      const configs = []
      Object.values(this.configs).forEach((item) => {
        // 注意：value值类型全都是字符串，所以提交上去的value值也要转换成字符串
        let value = ''
        try {
          value = item.value.toString()
        } catch (error) {}
        configs.push({ ...item, value })
      })
      const res = await updateConfig({ config: configs })
      if (res.status === 200) {
        this.$message.success('配置更新成功')
      } else {
        this.$message.error(res.data.message || '配置更新失败')
      }
      this.loading = false
    },
    success(res, index) {
      this.configs[index] = { ...this.configs[index], value: res.data.path }
    },
  },
}
</script>
<style lang="scss">
.com-form-config {
  .el-form-item__label {
    padding-bottom: 0;
  }
}
</style>
