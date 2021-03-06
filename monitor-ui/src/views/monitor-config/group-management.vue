<template>
  <div class="main-content">
    <PageTable :pageConfig="pageConfig" ref="child">
      <div slot="extraBtn">
        <button type="button" class="btn-cancel-f" @click="exportThreshold">{{$t("button.export")}}</button>
        <div style="display: inline-block;margin-bottom: 3px;vertical-align: bottom;"> 
          <Upload 
          :action="uploadUrl" 
          :show-upload-list="false"
          :max-size="1000"
          with-credentials
          :headers="{'X-Auth-Token': token}"
          :on-success="uploadSucess"
          :on-error="uploadFailed">
            <Button icon="ios-cloud-upload-outline">{{$t('button.upload')}}</Button>
          </Upload>
        </div>
      </div>
    </PageTable>
    <ModalComponent :modelConfig="modelConfig"></ModalComponent>
    <ModalComponent :modelConfig="authorizationModel">
      <div slot="authorization">  
        <div>
          <label class="col-md-2 label-name">{{$t('field.role')}}:</label>
          <Select v-model="authorizationModel.addRow.role" multiple filterable style="width:338px">
              <Option v-for="item in authorizationModel.roleList" :value="item.id" :key="item.id">
              {{item.display_name}}</Option>
          </Select>
        </div>
      </div>
    </ModalComponent>
    <Modal
        v-model="isShowWarning"
        title="Delete confirmation"
        @on-ok="ok"
        @on-cancel="cancel">
        <div class="modal-body" style="padding:30px">
          <div style="text-align:center">
            <p style="color: red">Will you delete it?</p>
          </div>
        </div>
    </Modal>
  </div>
</template>
<script>
  import { getToken} from '@/assets/js/cookies.ts'
  import {baseURL_config} from '@/assets/js/baseURL'
  let tableEle = [
    {title: 'tableKey.name', value: 'name', display: true},
    {title: 'tableKey.description', value: 'description', display: true}
  ]
  const btn = [
    {btn_name: 'field.endpoint', btn_func: 'checkMember'},
    {btn_name: 'field.threshold', btn_func: 'thresholdConfig'},
    {btn_name: 'button.edit', btn_func: 'editF'},
    {btn_name: 'button.remove', btn_func: 'deleteConfirmModal'},
    {btn_name: 'field.log', btn_func: 'logManagement'},
    {btn_name: 'button.authorize', btn_func: 'authorizeF'},
  ]
  import axios from 'axios'
  export default {
    name: '',
    data() {
      return {
        isShowWarning: false,
        token: null,
        uploadUrl: '',
        pageConfig: {
          CRUD: '',
          researchConfig: {
            input_conditions: [
              {value: 'search', type: 'input', placeholder: 'placeholder.input', style: ''}],
            btn_group: [
              {btn_name: 'button.search', btn_func: 'search', class: 'btn-confirm-f', btn_icon: 'fa fa-search'},
              {btn_name: 'button.add', btn_func: 'add', class: 'btn-cancel-f', btn_icon: 'fa fa-plus'}
            ],
            filters: {
              name__icontains: '',
              cmdb_tenant_id__icontains: ''
            }
          },
          table: {
            selection: true,
            tableData: [],
            tableEle: tableEle,
            // filterMoreBtn: 'filterMoreBtn',
            primaryKey: 'id',
            btn: btn,
            pagination: this.pagination,
            handleFloat:true,
          },
          pagination: { // [通用]-分页组件相关配置
            __orders: '-created_date',
            total: 0,
            page: 1,
            size: 10
          }
        },
        modelTip: {
          key: 'name',
          value: null
        },
        modelConfig: {
          modalId: 'add_edit_Modal',
          modalTitle: 'title.groupAdd',
          isAdd: true,
          config: [
            {label: 'tableKey.name', value: 'name', placeholder: 'tips.inputRequired', v_validate: 'required:true|min:2|max:60', disabled: false, type: 'text'},
            {label: 'tableKey.description', value: 'description', placeholder: '', disabled: false, type: 'text'},
          ],
          addRow: { // [通用]-保存用户新增、编辑时数据
            name: null,
            description: null,
          },
        },
        authorizationModel: {
          modalId: 'authorization_model',
          modalTitle: 'button.authorization',
          isAdd: true,
          saveFunc: 'authorizationSave',
          config: [
            {name:'authorization',type:'slot'}
          ],
          addRow: {
            role: []
          },
          roleList: []
        }, 
        id: null, // [通用]-待编辑数据id
        selectedData: '', // 存放选中数据
      }
    },
    mounted() {
      this.pageConfig.CRUD = this.$root.apiCenter.groupManagement.list.api
      this.token = getToken()
      this.initData(this.pageConfig.CRUD, this.pageConfig)
      this.uploadUrl =  baseURL_config + this.$root.apiCenter.groupManagement.upload.api
      this.$refs.child.clearSelectedData()
    },
    methods: {
      initData (url= this.pageConfig.CRUD, params) {
        this.$root.$tableUtil.initTable(this, 'GET', url, params)
      },
      add () {
        this.$root.JQ('#add_edit_Modal').modal('show')
      },
      addPost () {
        let params= this.$root.$validate.isEmptyReturn_JSON(this.modelConfig.addRow)
        this.$root.$httpRequestEntrance.httpRequestEntrance('POST', this.$root.apiCenter.groupManagement.add.api, params, () => {
          this.$root.$validate.emptyJson(this.modelConfig.addRow)
          this.$root.JQ('#add_edit_Modal').modal('hide')
          this.$Message.success(this.$t('tips.success'))
          this.initData(this.pageConfig.CRUD, this.pageConfig)
        })
      },
      editPost () {
        let params= this.$root.$validate.isEmptyReturn_JSON(this.modelConfig.addRow)
        params.id = this.id
        this.$root.$httpRequestEntrance.httpRequestEntrance('POST', this.$root.apiCenter.groupManagement.update.api, params, () => {
          this.$root.$validate.emptyJson(this.modelConfig.addRow)
          this.$root.JQ('#add_edit_Modal').modal('hide')
          this.$Message.success(this.$t('tips.success'))
          this.initData(this.pageConfig.CRUD, this.pageConfig)
        })
      },
      editF (rowData) {
        this.modelConfig.isAdd = false
        this.modelTip.value = rowData[this.modelTip.key]
        this.id = rowData.id
        this.modelConfig.addRow = this.$root.$tableUtil.manageEditParams(this.modelConfig.addRow, rowData)
        this.$root.JQ('#add_edit_Modal').modal('show')
      },
      checkMember (rowData) {
        this.$router.push({name: 'endpointManagement', params: {group: rowData}})
      },
      deleteConfirmModal (rowData) {
        this.selectedData = rowData
        this.isShowWarning = true
      },
      ok () {
        this.delF(this.selectedData)
      },
      cancel () {
        this.isShowWarning = false
      },
      deleteConfirm (rowData) {
        this.$delConfirm({
          msg: rowData.name,
          callback: () => {
            this.delF(rowData)
          }
        })
      },
      delF (rowData) {
        let params = {id: rowData.id}
        this.$root.$httpRequestEntrance.httpRequestEntrance('GET', this.$root.apiCenter.groupManagement.delete.api, params, () => {
          // this.$root.$eventBus.$emit('hideConfirmModal')
          this.$Message.success(this.$t('tips.success'))
          this.initData(this.pageConfig.CRUD, this.pageConfig)
          this.isShowWarning = false
        })
      },
      thresholdConfig (rowData) {
        this.$router.push({name: 'thresholdManagement', params: {id: rowData.id, type: 'grp'}})
      },
      logManagement (rowData) {
        this.$router.push({name: 'logManagement', params: {id: rowData.id, type: 'grp'}})
      },
      exportThreshold () {
        if (!this.$validate.isEmpty(this.selectedData.checkedIds)) {
          this.$Message.warning(this.$t('tips.selectData'))
          return
        }
        const api = this.$root.apiCenter.groupManagement.export.api + '?id=' + this.selectedData.checkedIds.join(',')
        axios({
          method: 'GET',
          url: api,
          headers: {
            'X-Auth-Token': getToken() || null
          }
        }).then((response) => {
          if (response.status < 400) {
           let content = JSON.stringify(response.data)
          let fileName = `grp_strategy_tpl_${new Date().format('yyyyMMddhhmmss')}.json`
          let blob = new Blob([content])
          if('msSaveOrOpenBlob' in navigator){
            // Microsoft Edge and Microsoft Internet Explorer 10-11
            window.navigator.msSaveOrOpenBlob(blob, fileName)
          } else {
            if ('download' in document.createElement('a')) { // 非IE下载
              let elink = document.createElement('a')
              elink.download = fileName
              elink.style.display = 'none'
              elink.href = URL.createObjectURL(blob)  
              document.body.appendChild(elink)
              elink.click()
              URL.revokeObjectURL(elink.href) // 释放URL 对象
              document.body.removeChild(elink)
            } else { // IE10+下载
              navigator.msSaveOrOpenBlob(blob, fileName)
            }
          }
          }
        })
        .catch(() => {
          this.$Message.warning(this.$t('tips.failed'))
        });

      },
      uploadSucess () {
        this.$Message.success(this.$t('tips.success'))
        this.initData(this.pageConfig.CRUD, this.pageConfig)
      },
      uploadFailed () {
        this.$Message.warning(this.$t('tips.failed'))
      },
      authorizeF (rowData) {
        this.id = rowData.id
        this.$root.$httpRequestEntrance.httpRequestEntrance('GET', this.$root.apiCenter.groupManagement.allRoles.api, '', (responseData) => {
          this.authorizationModel.roleList = responseData.data
          this.existRole()
        })
      },
      existRole () {
        this.$root.$httpRequestEntrance.httpRequestEntrance('GET', this.$root.apiCenter.groupManagement.exitRoles.api, {grp_id: this.id}, (responseData) => {
          responseData.forEach((item) => {
            this.authorizationModel.addRow.role.push(item.id)
          })
          this.$root.JQ('#authorization_model').modal('show')
        })
      },
      authorizationSave () {
        let params = {
          grp_id: this.id,
          role_id: this.authorizationModel.addRow.role
        }
        this.$root.$httpRequestEntrance.httpRequestEntrance('POST', this.$root.apiCenter.groupManagement.updateRoles.api, params, () => {
          this.$Message.success(this.$t('tips.success'))
          this.$root.JQ('#authorization_model').modal('hide')
        })
      } 
    },
    components: {
    }
  }
</script>

<style lang="less" scoped>
</style>
