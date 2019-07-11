<template>
    <div class="warp-container mod-role">
        <el-form :inline="true" :model="dataForm" ref="search_from" @keyup.enter.native="getDataList()"
                 label-width="80px">
            <el-row>
                <el-col :span="21">
                    <el-form-item :label="$t('role.column.roleName')">
                        <el-input v-model="dataForm.roleName" :placeholder="$t('role.column.roleName')"
                                  clearable></el-input>
                    </el-form-item>
                    <el-form-item>
                        <el-button @click="getDataList()" type="primary">
                            <icon-svg name="search1"></icon-svg>
                            {{"{"}}{ $t("common.search") }{{"}"}}
                        </el-button>
                    </el-form-item>
                    <el-form-item>
                        <el-button @click="resetSearchBox()" type="info">
                            <icon-svg name="reset"></icon-svg>
                            {{"{"}}{ $t("common.reset") }}" }{{"}"}}
                        </el-button>
                    </el-form-item>
                </el-col>
                <el-col :span="3" class="expandsearch">
                </el-col>
            </el-row>
            <el-collapse-transition>
                <el-card :body-style="{height:'auto'}">
                    <el-row>
                        <el-col :span="18">
                            <div class="data_selected">
                                <icon-svg name="tip"></icon-svg>
                                <span class="selectOpBox">{{"{"}}{ $t("common.haveChosen") }{{"}"}}</span>
                                <span class="selectOpBox_number">{{ dataListSelections.length }}</span>
                                <span class="selectOpBox">{{"{"}}{ $t("common.row") }{{"}"}}</span>
                            </div>
                            <span style="margin-left: 30px">
                                <el-button v-if="isAuth('sys:role:delete')"
                                           :type="dataListSelections.length == 0 ?'primary' : 'danger'"
                                           plain :disabled="dataListSelections.length == 0" @click="deleteHandle()">
                                  <icon-svg name="del"></icon-svg> {{"{"}}{ $t("common.delete") }{{"}"}}
                                </el-button>
                            </span>
                        </el-col>
                        <el-col :span="6" style="text-align:right;">
                            <el-form-item>
                                <el-button v-if="isAuth('sys:role:save')" type="primary" @click="addOrUpdateHandle()">
                                    <icon-svg name="add"></icon-svg>
                                    {{"{"}}{ $t("common.add") }{{"}"}}
                                </el-button>
                            </el-form-item>
                            <el-form-item>
                                <el-button v-if="isAuth('sys:role:save')" type="info" @click="getDataList()" plain>
                                    <icon-svg name="refresh"></icon-svg>
                                    {{"{"}}{ $t("common.refresh") }{{"}"}}
                                </el-button>
                            </el-form-item>
                        </el-col>
                    </el-row>
                </el-card>
            </el-collapse-transition>
        </el-form>
        <el-table :data="dataList" border v-loading="dataListLoading" @selection-change="selectionChangeHandle"
                  style="width: 100%;">
            <el-table-column type="selection" header-align="center" align="center" width="40">
            </el-table-column>
            {% for column in columns %}
            <el-table-column prop="{{ column.FieldName }}" header-align="left" align="left" label="{{ column.ColumnComment }}">
            </el-table-column>
            {% endfor %}
            <el-table-column fixed="right" header-align="center" align="center" width="250"
                             :label="$t('common.operate')">
                <template slot-scope="scope">
                    <el-button v-if="isAuth('sys:role:update')" type="warning" size="small"
                               @click="addOrUpdateHandle(scope.row.roleId)">
                        {{"{"}}{ $t("common.modify") }{{"}"}}
                    </el-button>
                    </el-button>
                    <el-button v-if="isAuth('sys:role:delete')" type="danger" size="small"
                               @click="deleteHandle(scope.row.roleId,scope.row.roleName)">
                        {{"{"}}{ $t("common.delete") }{{"}"}}
                    </el-button>
                </template>
            </el-table-column>
        </el-table>
        <el-pagination @size-change="sizeChangeHandle" @current-change="currentChangeHandle" :current-page="pageIndex"
                       :page-sizes="[10, 20, 50, 100]" :page-size="pageSize" :total="totalPage"
                       layout="total, sizes, prev, pager, next, jumper">
        </el-pagination>
        <!-- 弹窗, 新增 / 修改 -->
        <!-- <add-or-update v-if="addOrUpdateVisible" ref="addOrUpdate" @refreshDataList="getDataList"></add-or-update> -->
        <add-or-update v-model="showMoal" ref="addOrUpdate" @refreshDataList="getDataList"></add-or-update>
    </div>
</template>

<script>
    import AddOrUpdate from './add-or-update'
    import "@/assets/scss/ui.scss"
    import {mapActions} from 'vuex'
    import baseMixin from '_cm/mixin/base'

    export default {
        data() {
            return {
                dataForm: {
                    roleName: ''
                },

            }
        },
        mixins: [baseMixin],
        components: {
            AddOrUpdate
        },
        activated() {
            this.getDataList()
        },
        methods: {
            ...mapActions({
                listAction: 'sys/role/list',
                deleteAction: 'sys/role/delete'
            }),
            //重置搜索框
            resetSearchBox() {
                this.$refs.search_from.resetFields()
                this.getDataList()
            },
            // 获取数据列表
            getDataList() {
                this.dataListLoading = true
                this.listAction({
                    'page': this.pageIndex,
                    'limit': this.pageSize,
                    'roleName': this.dataForm.roleName
                }).then(({list, totalCount}) => {
                    this.dataList = list
                    this.totalPage = totalCount
                    this.dataListLoading = false
                }).catch(errMsg => {
                    this.dataList = []
                    this.totalPage = 0
                    this.dataListLoading = false
                })
            },
        }
    }
</script>
