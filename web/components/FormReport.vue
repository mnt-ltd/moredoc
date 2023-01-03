<template>
  <div class="com-form-report">
    <el-form
      ref="report"
      label-position="top"
      label-width="80px"
      :model="report"
    >
      <el-form-item label="文档">
        <el-input v-model="report.document_title" :disabled="true"></el-input>
      </el-form-item>
      <el-form-item label="举报原因">
        <el-radio-group v-model="report.reason" class="report-reason">
          <el-row>
            <el-col
              :span="8"
              v-for="item in reportOptions"
              :key="'rs' + item.value"
            >
              <el-radio :label="item.value">{{ item.label }}</el-radio>
            </el-col>
          </el-row>
        </el-radio-group>
      </el-form-item>
      <template v-if="isAdmin">
        <el-form-item label="处理状态">
          <el-switch
            v-model="report.status"
            active-text="已处理"
            inactive-text="未处理"
          >
          </el-switch>
        </el-form-item>
        <el-form-item label="处理备注">
          <el-input
            v-model="report.remark"
            placeholder="请输入文档处理相关备注"
            type="textarea"
            rows="3"
          ></el-input>
        </el-form-item>
      </template>
      <el-form-item>
        <el-button
          type="primary"
          icon="el-icon-check"
          class="btn-block"
          @click="setReport"
          >提交</el-button
        >
      </el-form-item>
    </el-form>
  </div>
</template>
<script>
import { createReport, updateReport } from '~/api/report'
import { reportOptions } from '~/utils/enum'
export default {
  name: 'FormReport',
  props: {
    isAdmin: {
      type: Boolean,
      default: false,
    },
    initReport: {
      type: Object,
      default: () => {
        return {
          id: 0,
          report_id: 0,
          document_id: 0,
        }
      },
    },
  },
  data() {
    return {
      report: { id: 0, status: 0 },
      reportOptions,
      statusOptions: [
        { label: '未处理', value: 0 },
        { label: '已处理', value: 1 },
      ],
    }
  },
  watch: {
    initReport: {
      handler(val) {
        this.report = { status: 0, ...val }
      },
      immediate: true,
    },
  },
  created() {
    this.report = this.initReport
  },
  methods: {
    async setReport() {
      if (this.report.id > 0) {
        const res = await updateReport(this.report)
        if (res.status === 200) {
          this.$message.success('更新成功')
          this.$emit('success')
        } else {
          this.$message.error(res.data.message)
        }
      } else {
        const res = await createReport(this.report)
        if (res.status === 200) {
          this.$message.success('提交成功')
          this.$emit('success')
        } else {
          this.$message.error(res.data.message)
        }
      }
    },
    reset() {
      this.report = { id: 0 }
      this.$refs.report.resetFields()
      this.$refs.report.clearValidate()
    },
  },
}
</script>
<style lang="scss">
.com-form-report {
  .el-select {
    width: 100%;
  }
  .report-reason {
    width: 100%;
    .el-radio {
      margin-bottom: 10px;
    }
  }
}
</style>
