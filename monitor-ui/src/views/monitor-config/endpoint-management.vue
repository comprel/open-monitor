<template>
<div class="main-content">
  <div v-if="showGroupMsg" style="padding-left:20px">
    <Tag type="border" closable color="primary" @on-close="closeTag">{{$t('field.group')}}:{{groupMsg.name}}</Tag>
  </div>
  <PageTable :pageConfig="pageConfig"></PageTable>
  <ModalComponent :modelConfig="modelConfig">
    <div slot="advancedConfig" class="extentClass">
      <div class="marginbottom params-each">
        <label class="col-md-2 label-name">{{$t('field.endpoint')}}:</label>
        <Select v-model="modelConfig.slotConfig.resourceSelected" multiple filterable style="width:300px">
          <Option v-for="item in modelConfig.slotConfig.resourceOption" :value="item.id" :key="item.id">
            {{ item.guid }}
          </Option>
        </Select>
      </div>
    </div>
  </ModalComponent>
  <ModalComponent :modelConfig="historyAlarmModel">
    <div slot="historyAlarm">
      <tableTemp :table="historyAlarmPageConfig.table" :pageConfig="historyAlarmPageConfig"></tableTemp>
    </div>
  </ModalComponent>
  <ModalComponent :modelConfig="endpointRejectModel">
    <div slot="endpointReject">
      <div class="marginbottom params-each">
        <label class="col-md-2 label-name">{{$t('field.endpoint')}}:</label>
        <Select v-model="endpointRejectModel.addRow.type" style="width:338px" @on-change="typeChange">
          <Option v-for="item in endpointRejectModel.endpointType" :value="item.value" :key="item.value">
            {{item.label}}
          </Option>
        </Select>
      </div>
      <div class="marginbottom params-each" v-if="!(['host','windows'].includes(endpointRejectModel.addRow.type))">
        <label class="col-md-2 label-name">{{$t('field.instance')}}:</label>
        <input v-validate="'required'" v-model="endpointRejectModel.addRow.name" name="name" :class="{ 'red-border': veeErrors.has('name') }" type="text" class="col-md-7 form-control model-input c-dark" />
        <label class="required-tip">*</label>
        <label v-show="veeErrors.has('name')" class="is-danger">{{ veeErrors.first('name')}}</label>
      </div>
      <div class="marginbottom params-each" v-if="['mysql','redis','java'].includes(endpointRejectModel.addRow.type)">
        <label class="col-md-2 label-name">{{$t('button.trusteeship')}}:</label>
        <Checkbox v-model="endpointRejectModel.addRow.agent_manager"></Checkbox>
      </div>
      <section v-if="['mysql','redis','java'].includes(endpointRejectModel.addRow.type) && endpointRejectModel.addRow.agent_manager">
        <div v-if="['mysql','java'].includes(endpointRejectModel.addRow.type)" class="marginbottom params-each">
          <label class="col-md-2 label-name">{{$t('button.username')}}:</label>
          <input v-validate="'required'" v-model="endpointRejectModel.addRow.user" name="user" :class="{ 'red-border': veeErrors.has('user') }" type="text" class="col-md-7 form-control model-input c-dark" />
          <label class="required-tip">*</label>
          <label v-show="veeErrors.has('user')" class="is-danger">{{ veeErrors.first('user')}}</label>
        </div>
        <div class="marginbottom params-each">
          <label class="col-md-2 label-name">{{$t('button.password')}}:</label>
          <input v-validate="'required'" v-model="endpointRejectModel.addRow.password" name="password" :class="{ 'red-border': veeErrors.has('password') }" type="text" class="col-md-7 form-control model-input c-dark" />
          <label class="required-tip">*</label>
          <label v-show="veeErrors.has('password')" class="is-danger">{{ veeErrors.first('password')}}</label>
        </div>
      </section>
      <section v-if="endpointRejectModel.addRow.type === 'http'">
        <div class="marginbottom params-each">
          <label class="col-md-2 label-name">Method:</label>
          <input v-validate="'required'" v-model="endpointRejectModel.addRow.method" name="method" :class="{ 'red-border': veeErrors.has('method') }" type="text" class="col-md-7 form-control model-input c-dark" />
          <label class="required-tip">*</label>
          <label v-show="veeErrors.has('method')" class="is-danger">{{ veeErrors.first('method')}}</label>
        </div>
        <div class="marginbottom params-each">
          <label class="col-md-2 label-name">URL:</label>
          <input v-validate="'required'" v-model="endpointRejectModel.addRow.url" name="url" :class="{ 'red-border': veeErrors.has('url') }" type="text" class="col-md-7 form-control model-input c-dark" />
          <label class="required-tip">*</label>
          <label v-show="veeErrors.has('url')" class="is-danger">{{ veeErrors.first('url')}}</label>
        </div>
      </section>
      <div class="marginbottom params-each" v-if="endpointRejectModel.addRow.type === 'other'">
        <label class="col-md-2 label-name">exporter_type: </label>
        <input v-validate="'required'" v-model="endpointRejectModel.addRow.exporter_type" name="exporter_type" :class="{ 'red-border': veeErrors.has('exporter_type') }" type="text" class="col-md-7 form-control model-input c-dark" />
        <label class="required-tip">*</label>
        <label v-show="veeErrors.has('exporter_type')" class="is-danger">{{ veeErrors.first('exporter_type')}}</label>
      </div>
      <div class="marginbottom params-each" v-if="!(['ping','http'].includes(endpointRejectModel.addRow.type))">
        <label class="col-md-2 label-name">{{$t('button.port')}}:</label>
        <input v-validate="'required|isNumber'" v-model="endpointRejectModel.addRow.port" name="port" :class="{ 'red-border': veeErrors.has('port') }" type="text" class="col-md-7 form-control model-input c-dark" />
        <label class="required-tip">*</label>
        <label v-show="veeErrors.has('port')" class="is-danger">{{ veeErrors.first('port')}}</label>
      </div>
      <div class="marginbottom params-each" v-if="(['ping','http','telnet'].includes(endpointRejectModel.addRow.type))">
        <label class="col-md-2 label-name">{{$t('exporter')}}:</label>
        <Checkbox v-model="endpointRejectModel.addRow.exporter"></Checkbox>
      </div>
      <div v-if="endpointRejectModel.addRow.exporter">
        <label class="col-md-2 label-name">{{$t('exporter_address')}}:</label>
        <input v-validate="'required'" :placeholder="$t('exporter_address_placeholder')" v-model="endpointRejectModel.addRow.export_address" name="export_address" :class="{ 'red-border': veeErrors.has('export_address') }" type="text" class="col-md-7 form-control model-input c-dark" />
        <label class="required-tip">*</label>
        <label v-show="veeErrors.has('export_address')" class="is-danger">{{ veeErrors.first('export_address')}}</label>
      </div>
    </div>
  </ModalComponent>
  <ModalComponent :modelConfig="processConfigModel">
    <div slot="processConfig">
      <div class="marginbottom params-each">
        <div style="color:#fa7821">
          <span>{{$t('button.processConfiguration_tip1')}}</span>
        </div>
      </div>
      <section>
        <div style="display: flex;">
          <div class="port-title">
            <span>{{$t('processName')}}:</span>
          </div>
          <div class="port-title">
            <span>{{$t('processTags')}}:</span>
          </div>
          <div class="port-title">
            <span>{{$t('displayName')}}:</span>
          </div>
          <i class="fa fa-plus-square-o port-config-icon" @click="addProcess" aria-hidden="true"></i>
        </div>
      </section>
      <section v-for="(pl, plIndex) in processConfigModel.process_list" :key="plIndex">
        <div class="port-config">
          <div style="width: 55%">
            <input type="text" v-model.trim="pl.process_name" class="search-input" style="width: 93%" />
            <label class="required-tip">*</label>
          </div>
          <div style="width: 51%">
            <input type="text" v-model.trim="pl.tags" class="search-input" style="width: 93%" />
          </div>
          <div style="width: 47%">
            <input type="text" v-model.trim="pl.display_name" class="search-input" style="width: 93%" />
          </div>
          <span style="float: right" >
            <i class="fa fa-trash-o port-config-icon" @click="delProcess(plIndex)" aria-hidden="true"></i>
          </span>
        </div>
      </section>
    </div>
  </ModalComponent>
  <ModalComponent :modelConfig="portModel">
    <div slot="port">
      <section>
        <div style="display: flex;">
          <div class="port-title">
            <span>{{$t('field.port')}}:</span>
          </div>
          <div class="port-title">
            <span>{{$t('tableKey.description')}}:</span>
          </div>
        </div>
      </section>

      <section v-for="(pm, pmIndex) in portModel.portMsg" :key="pmIndex">
        <div class="port-config">
          <div style="width: 48%">
            <InputNumber v-model.number="pm.port" type="number" :min=1 :max=65535 style="width: 100%" placeholder="Required, type: Number" />
          </div>
          <div style="width: 48%">
            <input type="text" v-model="pm.note" class="search-input" style="width: 100%" />
          </div>
          <i class="fa fa-trash-o port-config-icon" v-if="portModel.portMsg.length > 1" @click="removePort(pmIndex)" aria-hidden="true"></i>
          <i class="fa fa-plus-square-o port-config-icon" @click="addPort" :style="{'visibility': pmIndex+1===portModel.portMsg.length?  'unset' : 'hidden'}" aria-hidden="true"></i>
        </div>
      </section>
    </div>
  </ModalComponent>
  <Modal v-model="isShowWarning" title="Delete confirmation" @on-ok="ok" @on-cancel="cancel">
    <div class="modal-body" style="padding:30px">
      <div style="text-align:center">
        <p style="color: red">Will you delete it?</p>
      </div>
    </div>
  </Modal>
  <Modal v-model="isShowDataMonitor" :title="$t('button.dataMonitoring')" :styles="{top: '100px',width: '1000px'}" footer-hide>
    <DataMonitor :endpointId="dbEndpointId" ref="dataMonitor"></DataMonitor>
  </Modal>

  <ModalComponent :modelConfig="maintenanceWindowModel">
    <template #maintenanceWindow>
      <div style="margin: 4px 12px;padding:8px 12px;border:1px solid #dcdee2;border-radius:4px">
        <template v-for="(item, index) in maintenanceWindowModel.result">
          <p :key="index" style="margin:6px 0">
            <Button
              @click="deleteMaintenanceWindow(index)"
              size="small"
              style="background-color: #ff9900;border-color: #ff9900;"
              type="error"
              icon="md-close"
            ></Button>
            <TimePicker format="HH:mm" type="timerange" v-model="item.time_list" :clearable="false" style="width: 200px"></TimePicker>
            <Select v-model="item.weekday" multiple filterable style="width:200px">
              <Option v-for="cycle in maintenanceWindowModel.cycleOption" :value="cycle" :key="cycle">{{
                cycle
              }}</Option>
            </Select>
          </p>
        </template>
        <Button
          @click="addEmptyMaintenanceWindow"
          type="success"
          size="small"
          style="background-color: #0080FF;border-color: #0080FF;"
          long
          >{{ $t('button.add') }}</Button
        >
      </div>
    </template>
  </ModalComponent>
</div>
</template>

<script>
import tableTemp from '@/components/table-page/table'
import DataMonitor from '@/views/monitor-config/data-monitor'
import {
  interceptParams
} from '@/assets/js/utils'
let tableEle = [{
    title: 'tableKey.endpoint',
    value: 'guid',
    display: true
  },
  {
    title: 'tableKey.group',
    display: true,
    tags: {
      style: 'width: 300px;'
    },
    'render': (item) => {
      let res = item.groups_name.split(',').map((i) => {
        return {
          label: i,
          value: i
        }
      })
      return res
    }
  }
]
let historyAlarmEle = [{
    title: 'tableKey.status',
    value: 'status',
    style: 'min-width:70px',
    display: true
  },
  {
    title: 'tableKey.s_metric',
    value: 's_metric',
    display: true
  },
  {
    title: 'tableKey.start_value',
    value: 'start_value',
    display: true
  },
  {
    title: 'tableKey.s_cond',
    value: 's_cond',
    style: 'min-width:70px',
    display: true
  },
  {
    title: 'tableKey.s_last',
    value: 's_last',
    style: 'min-width:65px',
    display: true
  },
  {
    title: 'tableKey.s_priority',
    value: 's_priority',
    display: true
  },
  {
    title: 'tableKey.start',
    value: 'start_string',
    style: 'min-width:200px',
    display: true
  },
  {
    title: 'tableKey.end',
    value: 'end_string',
    style: 'min-width:200px',
    display: true,
    'render': (item) => {
      if (item.end_string === '0001-01-01 00:00:00') {
        return '-'
      } else {
        return item.end_string
      }
    }
  }
]
const btn = [{
    btn_name: 'button.thresholdManagement',
    btn_func: 'thresholdConfig'
  },
  {
    btn_name: 'button.historicalAlert',
    btn_func: 'historyAlarm'
  },
  {
    btn_name: 'button.remove',
    btn_func: 'deleteConfirmModal'
  },
  {
    btn_name: 'button.logConfiguration',
    btn_func: 'logManagement'
  },
  {
    btn_name: 'button.portConfiguration',
    btn_func: 'portManagement'
  },
  {
    btn_name: 'button.processConfiguration',
    btn_func: 'processManagement'
  },
  {
    btn_name: 'button.businessConfiguration',
    btn_func: 'businessManagement'
  },
  {
    btn_name: 'button.dataMonitoring',
    btn_func: 'dataMonitor'
  },
  {
    btn_name: 'button.maintenanceWindow',
    btn_func: 'maintenanceWindow'
  }
]
export default {
  name: '',
  data() {
    return {
      isShowDataMonitor: false,
      dbEndpointId: '',
      dbMonitorData: [],

      isShowWarning: false,
      pageConfig: {
        CRUD: '',
        researchConfig: {
          input_conditions: [{
            value: 'search',
            type: 'input',
            placeholder: 'placeholder.input',
            style: ''
          }],
          btn_group: [{
            btn_name: 'button.search',
            btn_func: 'search',
            class: 'btn-confirm-f',
            btn_icon: 'fa fa-search'
          }],
          filters: {
            search: ''
          }
        },
        table: {
          tableData: [],
          tableEle: tableEle,
          filterMoreBtn: 'filterMoreBtn',
          primaryKey: 'guid',
          btn: btn,
          pagination: this.pagination,
          handleFloat: true,
        },
        pagination: {
          __orders: '-created_date',
          total: 0,
          page: 1,
          size: 10
        }
      },
      historyAlarmPageConfig: {
        table: {
          tableData: [],
          tableEle: historyAlarmEle,
          btn: [],
        },
      },
      modelConfig: {
        modalId: 'add_object_Modal',
        modalTitle: 'button.add',
        isAdd: true,
        config: [{
          name: 'advancedConfig',
          type: 'slot'
        }],
        addRow: { // [通用]-保存用户新增、编辑时数据
          name: null,
          description: null,
        },
        slotConfig: {
          resourceSelected: [],
          resourceOption: []
        }
      },
      historyAlarmModel: {
        modalId: 'history_alarm_Modal',
        modalTitle: 'button.historicalAlert',
        modalStyle: 'width:930px;max-width: none;',
        noBtn: true,
        isAdd: true,
        config: [{
          name: 'historyAlarm',
          type: 'slot'
        }],
        pageConfig: {
          table: {
            tableData: [],
            tableEle: tableEle
          }
        },
      },
      portModel: {
        modalId: 'port_Modal',
        modalTitle: 'button.portConfiguration',
        saveFunc: 'portSave',
        isAdd: true,
        config: [{
          name: 'port',
          type: 'slot'
        }],
        portMsg: []
      },
      endpointRejectModel: {
        modalId: 'endpoint_reject_model',
        modalTitle: 'title.endpointAdd',
        isAdd: true,
        saveFunc: 'endpointRejectSave',
        config: [{
            name: 'endpointReject',
            type: 'slot'
          },
          {
            label: 'field.ip',
            value: 'ip',
            placeholder: 'tips.required',
            v_validate: 'required:true|isIP',
            disabled: false,
            type: 'text'
          }
        ],
        addRow: {
          name: '',
          type: 'host',
          ip: null,
          port: 9100,
          agent_manager: false,
          user: '',
          password: '',
          method: '',
          url: '',
          exporter_type: '',
          exporter: false,
          export_address: ''
        },
        endpointType: [{
            label: 'host',
            value: 'host'
          },
          {
            label: 'mysql',
            value: 'mysql'
          },
          {
            label: 'redis',
            value: 'redis'
          },
          {
            label: 'java',
            value: 'java'
          },
          {
            label: 'windows',
            value: 'windows'
          },
          {
            label: 'ping',
            value: 'ping'
          },
          {
            label: 'telnet',
            value: 'telnet'
          },
          {
            label: 'http',
            value: 'http'
          },
          {
            label: 'other',
            value: 'other'
          }
        ],
      },
      processConfigModel: {
        modalId: 'process_config_model',
        modalTitle: 'button.processConfiguration',
        isAdd: true,
        saveFunc: 'processConfigSave',
        config: [{
          name: 'processConfig',
          type: 'slot'
        }],
        addRow: {
          businessSet: [],
        },
        process_list: [],
      },

      maintenanceWindowModel: {
        modalId: 'maintenance_window_model',
        modalTitle: 'button.maintenanceWindow',
        isAdd: true,
        saveFunc: 'maintenanceWindowSave',
        config: [{
          name: 'maintenanceWindow',
          type: 'slot'
        }],
        addRow: {
          // businessSet: [],
        },
        result: [],
        cycleOption: ['All', 'Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday', 'Sunday']
      },
      id: null,
      showGroupMsg: false,
      groupMsg: {}
    }
  },
  mounted() {
    this.pageConfig.CRUD = this.$root.apiCenter.endpointManagement.list.api
    if (this.$root.$validate.isEmpty_reset(this.$route.params)) {
      this.groupMsg = {}
      this.showGroupMsg = false
      this.pageConfig.researchConfig.btn_group.push({
        btn_name: 'button.add',
        btn_func: 'endpointReject',
        class: 'btn-cancel-f',
        btn_icon: 'fa fa-plus'
      })
    } else {
      this.$parent.activeTab = '/monitorConfigIndex/endpointManagement'
      if (Object.prototype.hasOwnProperty.call(this.$route.params, 'group')) {
        this.groupMsg = this.$route.params.group
        this.showGroupMsg = true
        this.pageConfig.researchConfig.btn_group.push({
          btn_name: 'button.add',
          btn_func: 'add',
          class: 'btn-cancel-f',
          btn_icon: 'fa fa-plus'
        })
        this.pageConfig.researchConfig.filters.grp = this.groupMsg.id
      }
      if (Object.prototype.hasOwnProperty.call(this.$route.params, 'search')) {
        this.pageConfig.researchConfig.filters.search = this.$route.params.search
      }
    }
    this.initData(this.pageConfig.CRUD, this.pageConfig)
  },
  filters: {
    interceptParams(val) {
      return interceptParams(val, 55)
    }
  },
  methods: {
    maintenanceWindowSave () {
      this.maintenanceWindowModel.result.forEach(item => {
        item.weekday = item.weekday.join(',')
      })
      const params = {
        endpoint: this.id,
        data: this.maintenanceWindowModel.result
      }
      this.$root.$httpRequestEntrance.httpRequestEntrance('POST', this.$root.apiCenter.endpointManagement.maintenanceWindow.update, params, () => {
        this.$root.JQ('#maintenance_window_model').modal('hide')
      })
    },
    maintenanceWindow (rowData) {
      this.id = rowData.guid
      const params = {
        endpoint: rowData.guid
      }
      this.$root.$httpRequestEntrance.httpRequestEntrance('GET', this.$root.apiCenter.endpointManagement.maintenanceWindow.get, params, (responseData) => {
        this.maintenanceWindowModel.result = []
        responseData.forEach(item => {
          this.maintenanceWindowModel.result.push({
            time_list: item.time_list,
            weekday: item.weekday.split(',')
          })
        })
        this.$root.JQ('#maintenance_window_model').modal('show')
      })
    },
    deleteMaintenanceWindow (index) {
      this.maintenanceWindowModel.result.splice(index, 1)
    },
    addEmptyMaintenanceWindow () {
      this.maintenanceWindowModel.result.push({time_list: ['00:00', '00:00'], weekday: ['All']})
    },
    dataMonitor(row) {
      this.dbEndpointId = row.id
      const params = {
        endpoint_id: this.dbEndpointId
      }
      this.$root.$httpRequestEntrance.httpRequestEntrance('GET', this.$root.apiCenter.endpointManagement.db.dbMonitor, params, responseData => {
        this.$refs.dataMonitor.managementData(responseData)
        this.dbMonitorData = responseData
        this.isShowDataMonitor = true
      })

    },
    typeChange(type) {
      this.endpointRejectModel.addRow = Object.assign(this.endpointRejectModel.addRow, {
        name: '',
        type,
        ip: null,
        port: 9100,
        agent_manager: false,
        user: '',
        password: '',
        method: '',
        url: '',
        exporter_type: ''
      })
      const typeToPort = {
        host: 9100,
        mysql: 9104,
        redis: 9121,
        java: 9151,
        windows: 9182
      }
      this.endpointRejectModel.addRow.port = typeToPort[type]
    },
    initData(url = this.pageConfig.CRUD, params) {
      this.$root.$tableUtil.initTable(this, 'GET', url, params)
    },
    filterMoreBtn(rowData) {
      let moreBtnGroup = ['thresholdConfig', 'historyAlarm', 'maintenanceWindow', 'deleteConfirmModal']
      if (rowData.type === 'host') {
        moreBtnGroup.push('processManagement', 'businessManagement', 'logManagement', 'portManagement')
      }
      if (rowData.type === 'mysql') {
        moreBtnGroup.push('dataMonitor')
      }
      // if (this.showGroupMsg) {
      //   moreBtnGroup.push('deleteConfirm')
      // }
      return moreBtnGroup
    },
    add() {
      this.modelConfig.slotConfig.resourceOption = []
      this.modelConfig.slotConfig.resourceSelected = []
      this.$root.$httpRequestEntrance.httpRequestEntrance('GET', '/monitor/api/v1/alarm/endpoint/list', {
        search: '.',
        page: 1,
        size: 300
      }, responseData => {
        responseData.data.forEach((item) => {
          if (item.id !== -1) {
            this.modelConfig.slotConfig.resourceOption.push(item)
          }
        })
        this.$root.JQ('#add_object_Modal').modal('show')
      })
    },
    addPost() {
      let params = {
        grp: this.groupMsg.id,
        endpoints: this.modelConfig.slotConfig.resourceSelected.map(Number),
        operation: 'add'
      }
      this.$root.$httpRequestEntrance.httpRequestEntrance('POST', this.$root.apiCenter.endpointManagement.update.api, params, () => {
        this.$Message.success(this.$t('tips.success'))
        this.$root.JQ('#add_object_Modal').modal('hide')
        this.initData(this.pageConfig.CRUD, this.pageConfig)
      })
    },
    deleteConfirmModal(rowData) {
      this.selectedData = rowData
      this.isShowWarning = true
    },
    ok() {
      this.delF(this.selectedData)
    },
    cancel() {
      this.isShowWarning = false
    },
    deleteConfirm(rowData) {
      this.$delConfirm({
        msg: rowData.guid,
        callback: () => {
          this.delF(rowData)
        }
      })
    },
    delF(rowData) {
      let endpoints = []
      this.pageConfig.table.tableData.forEach((item) => {
        endpoints.push(item.guid.split(':')[0])
      })
      let params = {
        guid: rowData.guid,
        endpoints: [parseInt(rowData.id)]
      }
      let url = this.$root.apiCenter.endpointManagement.deregister.api
      let methodType = 'GET'
      if (this.groupMsg.id) {
        url = this.$root.apiCenter.endpointManagement.update.api
        params.grp = this.groupMsg.id
        params.operation = 'delete'
        methodType = 'POST'
      }
      this.$root.$httpRequestEntrance.httpRequestEntrance(methodType, url, params, () => {
        // this.$root.$eventBus.$emit('hideConfirmModal')
        this.$Message.success(this.$t('tips.success'))
        this.initData(this.pageConfig.CRUD, this.pageConfig)
      })
    },
    thresholdConfig(rowData) {
      this.$router.push({
        name: 'thresholdManagement',
        params: {
          id: rowData.id,
          type: 'endpoint',
          paramsType: rowData.type
        }
      })
    },
    logManagement(rowData) {
      this.$router.push({
        name: 'logManagement',
        params: {
          id: rowData.id,
          type: 'endpoint'
        }
      })
    },
    closeTag() {
      this.groupMsg = {}
      this.showGroupMsg = false
      this.pageConfig.researchConfig.filters.grp = ''
      this.pageConfig.table.btn.splice(this.pageConfig.table.btn.length - 1, 1)
      this.pageConfig.researchConfig.btn_group.splice(this.pageConfig.researchConfig.btn_group.length - 1, 1)
      this.pageConfig.researchConfig.btn_group.push({
        btn_name: 'button.add',
        btn_func: 'endpointReject',
        class: 'btn-cancel-f',
        btn_icon: 'fa fa-plus'
      })
      this.initData(this.pageConfig.CRUD, this.pageConfig)
    },
    historyAlarm(rowData) {
      let params = {
        id: rowData.id
      }
      this.$root.$httpRequestEntrance.httpRequestEntrance('GET', this.$root.apiCenter.alarm.history, params, (responseData) => {
        this.historyAlarmPageConfig.table.tableData = responseData
      })
      this.$root.JQ('#history_alarm_Modal').modal('show')
    },
    endpointReject() {
      this.endpointRejectModel.addRow.type = 'host'
      this.endpointRejectModel.addRow.port = 9100
      this.$root.JQ('#endpoint_reject_model').modal('show')
    },
    endpointRejectSave() {
      this.endpointRejectModel.addRow.port += ''
      let params = this.$root.$validate.isEmptyReturn_JSON(this.endpointRejectModel.addRow)
      this.$validator.validate().then(result => {
        if (!result) return
        if (this.endpointRejectModel.addRow.exporter_type && ['host', 'mysql', 'redis', 'java', 'windows', 'ping', 'telnet', 'http'].includes(this.endpointRejectModel.addRow.exporter_type)) {
          this.$Message.warning('Export port existed!')
          return
        } else {
          if (this.endpointRejectModel.addRow.exporter_type) {
            params.type = this.endpointRejectModel.addRow.exporter_type
          }
        }
        this.$root.$httpRequestEntrance.httpRequestEntrance('POST', this.$root.apiCenter.endpointManagement.register.api, params, () => {
          this.$root.$validate.emptyJson(this.endpointRejectModel.addRow)
          this.$root.JQ('#endpoint_reject_model').modal('hide')
          this.$Message.success(this.$t('tips.success'))
          this.initData(this.pageConfig.CRUD, this.pageConfig)
        })
      })
    },
    processManagement(rowData) {
      this.processConfigModel.processName = ''
      this.id = rowData.id
      this.processConfigModel.addRow.processSet = []
      this.$root.$httpRequestEntrance.httpRequestEntrance('GET', '/monitor/api/v1/alarm/process/list', {
        id: this.id
      }, responseData => {
        if (!responseData.length) {
          responseData.push({
            process_name: '',
            display_name: '',
            tags: ''
          })
        }
        this.processConfigModel.process_list = responseData
        this.$root.JQ('#process_config_model').modal('show')
      })
    },
    processConfigSave() {
      const emptyPath = this.processConfigModel.process_list.some(t => {
        return !t.process_name
      })
      if (emptyPath) {
        this.$Message.warning(this.$t('tableKey.name') + this.$t('tips.required'))
        return
      }
      const params = {
        endpoint_id: +this.id,
        process_list: this.processConfigModel.process_list,
        check: true
      }
      this.$root.$httpRequestEntrance.httpRequestEntrance('POST', '/monitor/api/v1/alarm/process/update', params, () => {
        this.$Message.success(this.$t('tips.success'))
      })
      this.$root.JQ('#process_config_model').modal('hide')
    },
    addProcess() {
      const emptyPath = this.processConfigModel.process_list.some(t => {
        return !t.process_name
      })
      if (emptyPath) {
        this.$Message.warning(this.$t('tableKey.name') + this.$t('tips.required'))
        return
      }
      this.processConfigModel.process_list.push({
        process_name: '',
        display_name: '',
        tags: ''
      })
    },
    delProcess(plIndex) {
      this.processConfigModel.process_list.splice(plIndex, 1)
    },

    businessManagement(rowData) {
      this.$router.push({name: 'businessMonitor', params: rowData})
    },
    addBusiness() {
      const emptyPath = this.businessConfigModel.pathMsg.some(t => {
        return !t.path
      })
      if (emptyPath) {
        this.$Message.warning(this.$t('tableKey.path') + this.$t('tips.required'))
        return
      }
      this.businessConfigModel.pathMsg.push({
        owner_endpoint: null,
        path: null
      })
    },
    delBusiness(pmIndex) {
      this.businessConfigModel.pathMsg.splice(pmIndex, 1)
    },
    portManagement(rowData) {
      this.id = rowData.guid
      let params = {
        guid: rowData.guid
      }
      this.$root.$httpRequestEntrance.httpRequestEntrance('GET', '/monitor/api/v1/agent/export/endpoint/telnet/get', params, (responseData) => {
        if (!responseData.length) {
          responseData.push({
            port: null,
            note: '',
          })
        }
        this.portModel.portMsg = responseData
      })
      this.$root.JQ('#port_Modal').modal('show')
    },
    addPort() {
      const emptyPort = this.portModel.portMsg.some(t => {
        return !t.port === true
      })
      if (emptyPort) {
        this.$Message.warning(this.$t('field.port') + this.$t('tips.required'))
        return
      }
      this.portModel.portMsg.push({
        port: null,
        note: '',
      })
    },
    removePort(pmIndex) {
      this.portModel.portMsg.splice(pmIndex, 1)
    },
    portSave() {
      let temp = JSON.parse(JSON.stringify(this.portModel.portMsg.filter(t => {
        if (!t.port === false) {
          return t
        }
      })))
      temp.map(t => {
        t.port += ''
        return t
      })
      const params = {
        guid: this.id,
        config: temp
      }
      this.$root.$httpRequestEntrance.httpRequestEntrance('POST', '/monitor/api/v1/agent/export/endpoint/telnet/update', params, () => {
        this.$Message.success(this.$t('tips.success'))
        this.$root.JQ('#port_Modal').modal('hide')
      })
    }
  },
  components: {
    tableTemp,
    DataMonitor
  }
}
</script>

<style lang="less" scoped>
.params-each /deep/ .ivu-checkbox {
  top: 3px !important;
}

.red-border {
  border: 1px solid red !important;
}

.is-danger {
  color: red;
  margin-left: 80px;
  margin-bottom: 0px;
}

.search-input {
  display: inline-block;
  height: 32px;
  padding: 4px 7px;
  font-size: 12px;
  border: 1px solid #dcdee2;
  border-radius: 4px;
  color: #515a6e;
  background-color: #fff;
  background-image: none;
  position: relative;
  cursor: text;

  width: 310px;
}

.search-input:focus {
  outline: 0;
  border-color: #57a3f3;
}

.search-input-content {
  display: inline-block;
  vertical-align: middle;
}

.port-title {
  width: 48%;
  font-size: 14px;
  padding: 2px 0 2px 4px;
  // border: 1px solid @blue-2;
}

.port-config {
  display: flex;
  margin-top: 4px;
}

.port-config-icon {
  font-size: 16px;
  margin: 7px 2px;
}

.fa-trash-o {
  color: @color-red;
}

.fa-plus-square-o {
  color: @color-blue;
}
</style>
