<template>
  <div class="com-form-category">
    <el-form
      ref="formCategory"
      label-position="top"
      label-width="80px"
      :model="category"
    >
      <el-form-item label="上级分类">
        <el-cascader
          v-model="category.parent_id"
          :options="trees"
          :filterable="true"
          :props="{
            checkStrictly: true,
            expandTrigger: 'hover',
            label: 'title',
            value: 'id',
          }"
          clearable
          placeholder="请选择上级分类"
        ></el-cascader>
      </el-form-item>
      <!-- 创建的时候，不支持上传封面。因为创建的时候，支持批量创建 -->
      <el-form-item
        v-if="
          category.id > 0 &&
          (!category.parent_id ||
            category.parent_id === 0 ||
            category.parent_id.length === 0)
        "
        label="分类图标(请上传长宽比为1:1的小图片)"
        class="form-item-cover"
      >
        <UploadImage
          :action="'/api/v1/upload/category'"
          :image="category.icon"
          :width="'48px'"
          @success="successUploadIcon"
        />
      </el-form-item>
      <el-form-item
        v-if="
          category.id > 0 &&
          (!category.parent_id ||
            category.parent_id === 0 ||
            category.parent_id.length === 0)
        "
        label="分类封面(一级分类才需要上传)"
        class="form-item-cover"
      >
        <UploadImage
          :action="'/api/v1/upload/category'"
          :image="category.cover"
          :width="'180px'"
          @success="successUpload"
        />
      </el-form-item>
      <el-row :gutter="20">
        <el-col :span="16">
          <el-form-item
            label="名称"
            prop="title"
            :rules="[
              { required: true, trigger: 'blur', message: '请输入名称' },
            ]"
          >
            <el-input
              v-model="category.title"
              :placeholder="
                category.id > 0
                  ? '请输入分类名称'
                  : '请输入分类名称，多个分类名称换行输入，重复的分类名称自动跳过...'
              "
              :type="category.id > 0 ? 'text' : 'textarea'"
              :rows="5"
              clearable
            ></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="8">
          <el-form-item label="排序(值越大越靠前)">
            <el-input-number
              v-model.number="category.sort"
              clearable
              :min="0"
              :step="1"
              placeholder="请输入排序值"
            ></el-input-number>
          </el-form-item>
        </el-col>
      </el-row>

      <el-form-item label="描述" prop="description">
        <el-input
          v-model="category.description"
          :type="'textarea'"
          description="请输入描述，支持换行"
          :rows="5"
          clearable
        ></el-input>
      </el-form-item>
      <el-row :gutter="20">
        <el-col :span="12">
          <el-form-item label="是否启用">
            <el-switch
              v-model="category.enable"
              style="display: block"
              active-color="#13ce66"
              inactive-color="#ff4949"
              active-text="是"
              inactive-text="否"
            >
            </el-switch> </el-form-item
        ></el-col>
        <el-col :span="12">
          <el-form-item label="显示分类描述">
            <el-switch
              v-model="category.show_description"
              style="display: block"
              active-color="#13ce66"
              inactive-color="#ff4949"
              active-text="是"
              inactive-text="否"
            >
            </el-switch> </el-form-item
        ></el-col>
      </el-row>
      <el-form-item>
        <el-button
          type="primary"
          class="btn-block"
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
import UploadImage from '~/components/UploadImage.vue'
import { createCategory, updateCategory } from '~/api/category'
export default {
  name: 'FormCategory',
  components: {
    UploadImage,
  },
  props: {
    initCategory: {
      type: Object,
      default: () => {
        return {
          id: 0,
          title: '',
          sort: 0,
          enable: false,
          cover: '',
          icon: '',
        }
      },
    },
    trees: {
      type: Array,
      default: () => {
        return []
      },
    },
  },
  data() {
    return {
      loading: false,
      category: {
        id: 0,
        title: '',
        sort: 0,
        enable: false,
        cover: '',
        icon: '',
      },
    }
  },
  watch: {
    initCategory: {
      handler(val) {
        if (!val.sort) val.sort = 0
        if (!val.cover) val.cover = ''
        if (!val.icon) val.icon = ''
        this.category = { ...val }
      },
      immediate: true,
    },
  },
  created() {
    this.category = { title: '', cover: '', sort: 0, ...this.initCategory }
  },
  methods: {
    onSubmit() {
      this.$refs.formCategory.validate(async (valid) => {
        if (!valid) {
          return
        }
        this.loading = true
        const category = { ...this.category }
        if (category.parent_id) {
          if (typeof category.parent_id === 'object') {
            category.parent_id =
              category.parent_id[category.parent_id.length - 1]
          }
        } else {
          category.parent_id = 0
        }

        if (this.category.id > 0) {
          if (
            category.parent_id > 0 ||
            (typeof category.parent_id === 'object' &&
              category.parent_id.length > 0)
          ) {
            category.cover = ''
          }

          const res = await updateCategory(category)
          if (res.status === 200) {
            this.$message.success('修改成功')
            this.resetFields()
            this.$emit('success', res.data)
          } else {
            this.$message.error(res.data.message)
          }
        } else {
          const res = await createCategory(category)
          if (res.status === 200) {
            this.$message.success('新增成功')
            this.resetFields()
            this.$emit('success', res.data)
          } else {
            this.$message.error(res.data.message)
          }
        }
        this.loading = false
      })
    },
    clearValidate() {
      this.$refs.formCategory.clearValidate()
    },
    resetFields() {
      this.category = { title: '', cover: '', sort: 0, ...this.initCategory }
    },
    reset() {
      this.resetFields()
      this.clearValidate()
    },
    successUpload(res) {
      this.category.cover = res.data.path
    },
    successUploadIcon(res) {
      this.category.icon = res.data.path
    },
  },
}
</script>
<style lang="scss">
.com-form-category {
  .form-item-cover {
    .el-form-item__content {
      line-height: 1;
    }
  }
}
</style>
