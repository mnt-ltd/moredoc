<template>
  <div class="page-admin-dashboard">
    <el-card shadow="never">
      <div slot="header">服务器状态</div>
      <el-row :gutter="20" class="gauges">
        <el-col
          :span="6"
          :xs="12"
          :sm="8"
          :md="6"
          :lg="6"
          v-for="(gauge, index) in gauges"
          :key="'gauge' + index"
        >
          <v-chart class="chart" autoresize :option="gauge" />
          <div class="text-center">
            <ul>
              <li v-for="(item, index) in gauge.labels" :key="'label-' + index">
                <small v-if="item.label">{{ item.label }} : </small
                >{{ item.value }}
              </li>
            </ul>
          </div>
        </el-col>
      </el-row>
    </el-card>
    <el-card shadow="never" class="mgt-20px">
      <div slot="header">
        <span>授权信息</span>
      </div>
      <el-descriptions class="margin-top" :column="2" border>
        <el-descriptions-item>
          <template slot="label">
            <i class="el-icon-tickets"></i>
            最大文档数
          </template>
          不限
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label">
            <i class="el-icon-user"></i>
            最大用户数
          </template>
          不限
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label">
            <i class="el-icon-cpu"></i>
            授权协议
          </template>
          Apache License 2.0
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label">
            <i class="el-icon-time"></i>
            授权截止日期
          </template>
          <span>-</span>
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label">
            <i class="el-icon-guide"></i>
            授权类型
          </template>
          <span class="opensource">
            <span>魔豆文库 · 社区版</span>
            （<a
              href="https://www.bookstack.cn/read/moredoc/price.md"
              target="_blank"
              class="el-link el-link--primary"
              >版本划分与定价策略 <i class="el-icon-top-right"></i> </a
            >）
          </span>
        </el-descriptions-item>
      </el-descriptions>
    </el-card>
    <el-card shadow="never" class="mgt-20px">
      <div slot="header">数据统计</div>
      <el-descriptions class="margin-top" :column="3" border>
        <el-descriptions-item>
          <template slot="label">
            <i class="el-icon-tickets"></i>
            文档
          </template>
          {{ stats.document_count || 0 }}
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label">
            <i class="el-icon-chat-dot-square"></i>
            评论
          </template>
          {{ stats.comment_count || 0 }}
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label">
            <i class="el-icon-user"></i>
            用户
          </template>
          {{ stats.user_count || 0 }}
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label">
            <i class="el-icon-s-grid"></i>
            分类
          </template>
          {{ stats.category_count || 0 }}
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label">
            <i class="el-icon-warning"></i>
            举报
          </template>
          {{ stats.report_count || 0 }}
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label">
            <i class="el-icon-picture-outline"></i>
            横幅
          </template>
          {{ stats.banner_count || 0 }}
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label">
            <i class="el-icon-link"></i>
            友链
          </template>
          {{ stats.friendlink_count || 0 }}
        </el-descriptions-item>
      </el-descriptions>
    </el-card>
    <el-card shadow="never" class="mgt-20px">
      <div slot="header">
        <span>环境依赖</span>
        <a
          href="https://www.bookstack.cn/read/moredoc/install.md"
          target="_blank"
        >
          <el-button
            style="float: right; padding: 3px 0"
            icon="el-icon-tickets"
            type="text"
          >
            依赖安装教程</el-button
          >
        </a>
      </div>
      <el-table
        :data="envs"
        style="width: 100%"
        empty-text="您暂无权限查看环境依赖情况"
      >
        <el-table-column prop="name" label="名称" width="100"> </el-table-column>
        <el-table-column prop="is_required" label="是否必须" width="100">
          <template slot-scope="scope">
            <el-tag
              v-if="scope.row.is_required"
              effect="dark"
              type="danger"
              size="small"
              >必须安装</el-tag
            >
            <el-tag effect="dark" type="info" size="small" v-else
              >非必须</el-tag
            >
          </template>
        </el-table-column>
        <el-table-column prop="is_installed" label="安装" width="100">
          <template slot-scope="scope">
            <el-tag
              v-if="scope.row.is_installed"
              effect="dark"
              type="success"
              size="small"
              >已安装</el-tag
            >
            <el-tag effect="dark" type="warning" size="small" v-else
              >未安装</el-tag
            >
          </template>
        </el-table-column>
        <el-table-column prop="version" label="版本" min-width="100">
          <template slot-scope="scope">
            {{ scope.row.version || '-'  }}
          </template>
        </el-table-column>
        <el-table-column prop="description" min-width="200" label="用途">
        </el-table-column>
        <el-table-column prop="error" label="错误" min-width="100">
          <template slot-scope="scope">
            <span v-if="scope.row.error" size="small">{{
              scope.row.error
            }}</span>
            <span v-else>-</span>
          </template>
        </el-table-column>
        <el-table-column prop="checked_at" width="180" label="检测">
          <template slot-scope="scope">
            {{ formatDatetime(scope.row.checked_at) }}
          </template>
        </el-table-column>
      </el-table>
    </el-card>
    <el-card shadow="never" class="mgt-20px">
      <div slot="header">
        <span>系统信息</span>
        <el-button
          style="float: right; padding: 3px 0"
          @click="updateSitemap"
          :loading="loading"
          icon="el-icon-refresh"
          type="text"
        >
          更新站点地图</el-button
        >
      </div>
      <el-descriptions class="margin-top" :column="1" border>
        <el-descriptions-item>
          <template slot="label"> 操作系统 </template>
          {{ stats.os }}
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label"> 程序名称 </template>
          moredoc · 魔豆文库
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label"> 程序版本 </template>
          {{ stats.version }}
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label"> Git Hash </template>
          {{ stats.hash }}
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label"> 构建时间 </template>
          {{ formatDatetime(stats.build_at) }}
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label"> 程序开发 </template>
          深圳市摩枫网络科技有限公司 <b>M</b>orefun <b>N</b>etwork
          <b>T</b>echnology Co., <b>Ltd</b>.
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label"> 服务邮箱 </template>
          <a href="mailto:truthhun@mnt.ltd" class="el-link el-link--primary"
            >truthhun@mnt.ltd</a
          >
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label"> 服务官网 </template>
          <a
            href="https://mnt.ltd"
            class="el-link el-link--primary"
            target="_blank"
            title="摩枫网络科技 MNT.Ltd"
            >https://mnt.ltd</a
          >
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label"> 使用手册 </template>
          <a
            href="https://www.bookstack.cn/books/moredoc"
            class="el-link el-link--primary"
            target="_blank"
            title="程序使用手册"
            >https://www.bookstack.cn/books/moredoc</a
          >
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label"> 开源地址 </template>
          <ul class="opensource">
            <li>
              MNT：
              <a
                href="https://git.mnt.ltd/mnt/moredoc"
                class="el-link el-link--primary"
                target="_blank"
                title="摩枫Git"
                >https://git.mnt.ltd/mnt/moredoc</a
              >
            </li>
            <li>
              Gitee：
              <a
                href="https://gitee.com/mnt-ltd/moredoc"
                class="el-link el-link--primary"
                target="_blank"
                title="Gitee"
                >https://gitee.com/mnt-ltd/moredoc</a
              >
            </li>
            <li>
              Github：
              <a
                href="https://github.com/mnt-ltd/moredoc"
                class="el-link el-link--primary"
                target="_blank"
                title="Github"
                >https://github.com/mnt-ltd/moredoc</a
              >
            </li>
          </ul>
        </el-descriptions-item>
      </el-descriptions>
    </el-card>
  </div>
</template>

<script>
import { getStats, getEnvs, updateSitemap, getDevice } from '~/api/config'
import { formatDatetime, formatBytes } from '~/utils/utils'
import { use } from 'echarts/core'
import { CanvasRenderer } from 'echarts/renderers'
import { PieChart, GaugeChart } from 'echarts/charts'
import { UniversalTransition } from 'echarts/features'
import {
  TitleComponent,
  TooltipComponent,
  LegendComponent,
  GridComponent,
} from 'echarts/components'
import VChart from 'vue-echarts'

use([
  CanvasRenderer,
  PieChart,
  TitleComponent,
  TooltipComponent,
  LegendComponent,
  GaugeChart,
  UniversalTransition,
  GridComponent,
])

export default {
  layout: 'admin',
  head() {
    return {
      title: `面板 - ${this.settings.system.sitename}`,
    }
  },
  components: {
    VChart,
  },
  data() {
    return {
      stats: {
        admin_count: 0,
        student_count: 0,
        company_count: 0,
        category_count: 0,
        article_count: 0,
        article_pending_count: 0,
        comment_count: 0,
        comment_pending_count: 0,
        banner_count: 0,
        friendlink_count: 0,
        user_pending_count: 0,
        os: '-',
        version: '-',
        hash: '-',
        build_at: '',
      },
      envs: [],
      loading: false,
      gauges: [],
      devices: [],
      timeouter: null,
    }
  },
  computed: {
    settings() {
      return this.$store.state.setting.settings
    },
  },
  created() {
    this.initDevice()
    Promise.all([this.getStats(), this.getEnvs(), this.loopGetDevice()])
  },
  beforeDestroy() {
    clearTimeout(this.timeouter)
  },
  methods: {
    formatDatetime,
    loopGetDevice() {
      this.getDevice()
      clearTimeout(this.timeouter)
      this.timeouter = setTimeout(() => {
        this.loopGetDevice()
      }, 5000)
    },
    async getStats() {
      const res = await getStats()
      if (res.status === 200) {
        this.stats = {
          ...this.stats,
          ...res.data,
        }
      }
    },
    async getEnvs() {
      const res = await getEnvs()
      if (res.status === 200) {
        this.envs = res.data.envs
      }
    },
    async updateSitemap() {
      this.loading = true
      const res = await updateSitemap()
      if (res.status === 200) {
        this.$message.success('更新成功')
        this.loading = false
        return
      }
      this.loading = false
      this.$message.error(res.data.message || '更新失败')
    },
    initDevice() {
      const gauges = [
        {
          ...this.getGaugeOption('CPU', '0.00'),
          labels: [
            {
              label: 'Cores',
              value: '-',
            },
            {
              label: 'Mhz',
              value: '-',
            },
            {
              label: '',
              value: '-',
            },
          ],
        },
        {
          ...this.getGaugeOption('内存', '0.00'),
          labels: [
            {
              label: 'Used',
              value: '-',
            },
            {
              label: 'Free',
              value: '-',
            },
            {
              label: 'Total',
              value: '-',
            },
          ],
        },
      ]
      gauges.push({
        ...this.getGaugeOption('磁盘', '0.00'),
        labels: [
          {
            label: 'Used',
            value: '-',
          },
          {
            label: 'Free',
            value: '-',
          },
          {
            label: 'Total',
            value: '-',
          },
        ],
      })
      this.gauges = gauges
    },
    async getDevice() {
      const res = await getDevice()
      if (res.status === 200) {
        const gauges = [
          {
            ...this.getGaugeOption(
              'CPU',
              (res.data.cpu.percent || 0).toFixed(2) || '0.00'
            ),
            labels: [
              {
                label: 'Cores',
                value: res.data.cpu.cores,
              },
              {
                label: 'Mhz',
                value: (res.data.cpu.mhz || 0).toFixed(0),
              },
              {
                label: '',
                value: res.data.cpu.model_name,
              },
            ],
          },
          {
            ...this.getGaugeOption(
              '内存',
              ((res.data.memory.used / res.data.memory.total) * 100).toFixed(
                2
              ) || '0.00'
            ),
            labels: [
              {
                label: 'Total',
                value: formatBytes(res.data.memory.total),
              },
              {
                label: 'Used',
                value: formatBytes(res.data.memory.used),
              },
              {
                label: 'Free',
                value: formatBytes(
                  res.data.memory.total - res.data.memory.used
                ),
              },
            ],
          },
        ]
        for (let disk of res.data.disk || []) {
          gauges.push({
            ...this.getGaugeOption(
              '磁盘 ' + disk.disk_name,
              (disk.percent || 0).toFixed(2) || '0.00'
            ),
            labels: [
              {
                label: 'Total',
                value: formatBytes(disk.total),
              },
              {
                label: 'Used',
                value: formatBytes(disk.used),
              },
              {
                label: 'Free',
                value: formatBytes(disk.free),
              },
            ],
          })
        }
        this.gauges = gauges
      }
    },
    getGaugeOption(name, percent) {
      return {
        series: [
          {
            type: 'gauge',
            axisLine: {
              lineStyle: {
                width: 10,
                color: [
                  [0.3, '#67e0e3'],
                  [0.7, '#37a2da'],
                  [1, '#fd666d'],
                ],
              },
            },
            pointer: {
              itemStyle: {
                color: 'inherit',
              },
            },
            axisTick: {
              distance: -30,
              length: 8,
              lineStyle: {
                color: '#fff',
                width: 2,
              },
            },
            splitLine: {
              distance: -15,
              length: 20,
              lineStyle: {
                color: '#fff',
                width: 4,
              },
            },
            axisLabel: {
              color: 'inherit',
              distance: 10,
              fontSize: 14,
            },
            detail: {
              valueAnimation: true,
              formatter: '{value} %',
              color: 'inherit',
              fontSize: 20,
              offsetCenter: [0, '85%'],
            },
            data: [
              {
                value: percent,
                name: name,
                fontSize: 14,
              },
            ],
          },
        ],
      }
    },
  },
}
</script>
<style lang="scss">
.page-admin-dashboard {
  .el-descriptions-item__label.is-bordered-label {
    width: 150px;
  }
  .systeminfo {
    b {
      color: crimson;
    }
  }
  .opensource .el-link {
    margin-top: -3px;
  }
  .chart {
    height: 234px;
  }
  .gauges {
    min-height: 294px;
    font-size: 14px;
    ul,
    li {
      list-style: none;
      padding: 0;
      margin: 0;
    }
    small {
      color: #999;
    }
    li {
      white-space: nowrap;
      overflow: hidden;
      text-overflow: ellipsis;
    }
  }
}
</style>
