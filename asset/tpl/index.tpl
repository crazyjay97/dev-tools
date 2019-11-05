<template>
    <div class="warp-container mod-role">
        <el-form :inline="true" :model="formData" ref="search_from" @keyup.enter.native="getDataList()"
                 label-width="80px">
            <el-row>
                <el-col :span="21">
                    {% for column in searchColumns %}
                    {% if forloop.Counter <= 2 %}
                    <el-form-item label="{{ column.ColumnComment }}">
                        {% if column.ShowMode == 0 %}
                        <el-input v-model="formData.{{ column.FieldName }}" placeholder="{{ column.ColumnComment }}"
                                  clearable></el-input>
                        {% elif column.ShowMode == 1 %}
                        <el-select v-model="formData.{{ column.FieldName }}" placeholder="{{ column.ColumnComment }}" clearable>
                            <el-option v-for="row in queryDictionary('{{column.DictionaryKey}}')" :key="row.codeValue" :value="row.codeValue"
                                       :label="row.codeText">
                            </el-option>
                        </el-select>
                        {% endif %}
                    </el-form-item>
                    {% endif %}
                    {% endfor %}
                    <el-form-item>
                        <el-button @click="getDataList()" type="primary">
                            <icon-svg name="search"></icon-svg>
                            {{"{"}}{ $t("common.search") }{{"}"}}
                        </el-button>
                    </el-form-item>
                    <el-form-item>
                        <el-button @click="resetSearchBox()" type="info">
                            <icon-svg name="reset"></icon-svg>
                            {{"{"}}{ $t("common.reset") }{{"}"}}
                        </el-button>
                    </el-form-item>
                </el-col>
                <el-col :span="3" class="expandsearch">
                    <el-button @click="expandSearch = !expandSearch">
                        <icon-svg :name="expandSearch ? 'up' : 'down'" style="vertical-align: middle;"></icon-svg>
                        <span style="vertical-align: middle;">{{ "{" }}{ $t("common.advancedSearch") }{{ "}" }}</span>
                    </el-button>
                </el-col>
            </el-row>
            <el-row>
                <el-col :span="24">
                    <el-collapse-transition>
                        <div v-show="expandSearch">
                            <el-row>
                                {% for column in searchColumns %}
                                {% if forloop.Counter > 2 %}
                                <el-form-item label="{{ column.ColumnComment }}">{% if column.ShowMode == 0 %}
                                    <el-input v-model="formData.{{ column.FieldName }}" placeholder="{{ column.ColumnComment }}"
                                              clearable></el-input>
                                    {% elif column.ShowMode == 1 %}
                                    <el-select v-model="formData.{{ column.FieldName }}" placeholder="{{ column.ColumnComment }}" clearable>
                                        <el-option v-for="row in queryDictionary('{{column.DictionaryKey}}')" :key="row.codeValue" :value="row.codeValue"
                                                   :label="row.codeText">
                                        </el-option>
                                    </el-select>
                                    {% endif %}
                                </el-form-item>
                                {% endif %}
                                {% endfor %}
                            </el-row>
                        </div>
                    </el-collapse-transition>
                </el-col>
            </el-row>
            <el-card :body-style="{height:'auto'}" class="operate-wrap">
                <el-row>
                    <el-col :span="16">
                        <div class="data_selected">
                            <icon-svg name="tip"></icon-svg>
                            <span class="selectOpBox">{{"{"}}{ $t("common.haveChosen") }{{"}"}}</span>
                            <span class="selectOpBox_number">{{"{"}}{ dataListSelections.length }{{"}"}}</span>
                            <span class="selectOpBox">{{"{"}}{ $t("common.row") }{{"}"}}</span>
                        </div>
                        <span style="margin-left: 30px">
                            <el-button v-if="isAuth('{{ moduleName }}:{{ fileName }}:delete')"
                                       :type="dataListSelections.length == 0 ?'primary' : 'danger'"
                                       plain :disabled="dataListSelections.length == 0" @click="deleteHandle()">
                              <icon-svg name="del"></icon-svg> {{"{"}}{ $t("common.delete") }{{"}"}}
                            </el-button>
                        </span>
                    </el-col>
                    <el-col :span="8" style="text-align:right;" class="operate-wrap-button">
                        <el-form-item>
                            <el-button v-if="isAuth('{{ moduleName }}:{{ fileName }}:save')" type="primary" @click="addOrUpdateHandle()">
                                <icon-svg name="add"></icon-svg>
                                {{"{"}}{ $t("common.add") }{{"}"}}
                            </el-button>
                        </el-form-item>
                        <el-form-item>
                            <el-button  type="info" @click="getDataList()" plain>
                                <icon-svg name="refresh"></icon-svg>
                                {{"{"}}{ $t("common.refresh") }{{"}"}}
                            </el-button>
                        </el-form-item>
                    </el-col>
                </el-row>
            </el-card>
        </el-form>
        <el-table :data="dataList" border v-loading="dataListLoading" @selection-change="selectionChangeHandle"
                  style="width: 100%;">
            <el-table-column type="selection" header-align="center" align="center" width="40">
            </el-table-column>
            {% for column in listColumns %}
            {% if column.ColumnKey != "PRI" %}
            {% if column.NeedShow %}
            <el-table-column prop="{{ column.FieldName }}" header-align="left" align="left" label="{{ column.ColumnComment }}" width="100" show-overflow-tooltip>
            </el-table-column>
            {% endif %}
            {% endif %}
            {% endfor %}
            <el-table-column fixed="right" header-align="center" align="center" width="150"
                             :label="$t('common.operate')">
                <template slot-scope="scope">
                    <el-button v-if="isAuth('{{ moduleName }}:{{ fileName }}:update')" type="warning" size="small"
                               @click="addOrUpdateHandle(scope.row)">
                        {{"{"}}{ $t("common.modify") }{{"}"}}
                    </el-button>
                    <el-button v-if="isAuth('{{ moduleName }}:{{ fileName }}:delete')" type="danger" size="small"
                               @click="deleteHandle(scope.row.id)">
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
        <add-or-update v-model="showMoal" ref="addOrUpdate" @refreshDataList="getDataList"></add-or-update>
    </div>
</template>

<script>
    import AddOrUpdate from './add-or-update'
    import {mapActions} from 'vuex'
    import baseMixin from '_cm/mixin/base'

    export default {
        data() {
            return {
                formData: {
                    {% for column in searchColumns %}
                    {{ column.FieldName }}: '',
                    {% endfor %}
                },
            }
        },
        mixins: [baseMixin],
        components: {
            AddOrUpdate
        },
        methods: {
            ...mapActions({
                listAction: '{{ moduleName }}/{{ fileName }}/list',
                deleteAction: '{{ moduleName }}/{{ fileName }}/delete'
            }),
            // 获取数据列表
            getDataList() {
                this.queryDataList({
                    'page': this.pageIndex,
                    'limit': this.pageSize,
                {% for column in searchColumns %}
                    '{{ column.FieldName }}': this.formData.{{ column.FieldName }},
                {% endfor %}
                })
            },
        }
    }
</script>
