<template>
  <div class="com-form-punishment">
    <el-form
      ref="formPunishment"
      label-position="top"
      label-width="80px"
      :model="punishment"
    >
      <el-form-item
        label="用户"
        prop="user_id"
        :rules="
          punishment.id === 0
            ? [{ required: true, trigger: 'blur', message: '请选择用户' }]
            : []
        "
      >
        <el-select
          v-if="punishment.id === 0"
          v-model="punishment.user_id"
          filterable
          multiple
          remote
          reserve-keyword
          placeholder="请输入和选择用户"
          :remote-method="remoteSearchUser"
          :loading="loading"
        >
          <el-option
            v-for="user in users"
            :key="'userid' + user.id"
            :label="user.username"
            :value="user.id"
          >
          </el-option>
        </el-select>
        <el-input v-else :disabled="true" v-model="punishment.username" />
      </el-form-item>
      <el-form-item
        prop="type"
        :rules="[
          { required: true, trigger: 'blur', message: '请选择处罚类型' },
        ]"
      >
        <template slot="label">
          处罚类型
          <ToolTip
            content="禁止评论：不允许发表评论；禁止上传：不允许上传文档；禁止收藏：不允许收藏；禁止下载：不允许下载文档；禁用账户：包括上述全部禁用项"
          />
        </template>
        <el-checkbox-group v-if="punishment.id === 0" v-model="punishment.type">
          <el-checkbox
            v-for="item in punishmentTypeOptions"
            :label="item.value"
            :key="'checkbox-pt' + item.value"
            >{{ item.label }}</el-checkbox
          >
        </el-checkbox-group>
        <el-select v-else v-model="punishment.type" :disabled="true">
          <el-option
            v-for="item in punishmentTypeOptions"
            :label="item.label"
            :key="'select-pt-' + item.value"
            :value="item.value"
          >
          </el-option>
        </el-select>
      </el-form-item>
      <el-row :gutter="20">
        <el-col :span="8">
          <el-form-item label="是否启用处罚">
            <el-switch
              v-model="punishment.enable"
              style="display: block"
              active-color="#13ce66"
              inactive-color="#ff4949"
              active-text="是"
              inactive-text="否"
            >
            </el-switch> </el-form-item
        ></el-col>
        <el-col :span="16">
          <el-form-item>
            <template slot="label">
              截止时间
              <ToolTip content="用户被处罚的截止时间，留空则为永久" />
            </template>
            <el-date-picker
              v-model="punishment.end_time"
              type="datetime"
              placeholder="请选择处罚截止时间"
              :picker-options="datetimePickerPunishmentOptions"
            >
            </el-date-picker>
          </el-form-item>
        </el-col>
      </el-row>
      <el-form-item label="处罚原因">
        <el-input
          type="textarea"
          v-model="punishment.reason"
          :rows="3"
          placeholder="请输入处罚原因，被处罚用户可见"
        ></el-input>
      </el-form-item>
      <el-form-item label="处罚备注">
        <el-input
          type="textarea"
          v-model="punishment.remark"
          :rows="3"
          placeholder="请输入处罚备注，管理员可见"
        ></el-input>
      </el-form-item>
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
import { createPunishment, updatePunishment } from '~/api/punishment'
import {
  punishmentTypeOptions,
  datetimePickerPunishmentOptions,
} from '~/utils/enum'
import { listUser } from '~/api/user'
export default {
  name: 'FormPunishment',
  props: {
    initPunishment: {
      type: Object,
      default: () => {
        return {}
      },
    },
  },
  data() {
    return {
      punishmentTypeOptions,
      datetimePickerPunishmentOptions,
      loading: false,
      punishment: {
        id: 0,
        user_id: '',
        remark: '',
        reason: '',
        type: [],
        enable: true,
      },
      users: [],
    }
  },
  watch: {
    initPunishment: {
      handler(val) {
        let enable = val.enable || false
        this.punishment = {
          id: 0,
          user_id: '',
          remark: '',
          reason: '',
          type: [],
          ...val,
          enable: enable,
        }
      },
      immediate: true,
    },
  },
  // created() {
  //   this.punishment = { ...this.initPunishment }
  // },
  methods: {
    onSubmit() {
      this.$refs.formPunishment.validate(async (valid) => {
        if (!valid) {
          return
        }
        this.loading = true
        const punishment = { ...this.punishment }
        if (this.punishment.id > 0) {
          delete punishment.operators
          const res = await updatePunishment(punishment)
          if (res.status === 200) {
            this.$message.success('修改成功')
            this.$emit('success', res.data)
          } else {
            this.$message.error(res.data.message)
          }
        } else {
          const res = await createPunishment(punishment)
          if (res.status === 200) {
            this.$message.success('新增成功')
            this.$emit('success', res.data)
          } else {
            this.$message.error(res.data.message)
          }
        }
        this.loading = false
      })
    },
    async remoteSearchUser(wd) {
      this.searchUser(wd)
    },
    async searchUser(wd, userId = []) {
      const res = await listUser({
        page: 1,
        size: 10,
        wd: wd,
        id: userId || [],
        field: ['id', 'username'],
      })
      if (res.status === 200) {
        this.users = res.data.user || []
      }
    },
    clearValidate() {
      this.$refs.formPunishment.clearValidate()
    },
    resetFields() {
      this.punishment = {
        id: 0,
        title: '',
        link: '',
        sort: 0,
        enable: true,
        description: '',
      }
    },
    reset() {
      this.resetFields()
      this.clearValidate()
    },
  },
}
</script>
