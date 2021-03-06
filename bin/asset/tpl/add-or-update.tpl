<template>
    <modal v-model="isShow" width="500px">
        <el-form :model="formData" ref="form" label-width="140px" :rules="dataRule" label-position="left">
            <div>
                <div class="i-box">
                    <div class="i-box-title">
                        <icon-svg name="line"></icon-svg>
                        编辑
                        <h5>{{"{"}}{ $t("common.tip") }{{"}"}}<span>*</span>{{"{"}}{ $t("common.tippp") }{{"}"}}</h5>
                    </div>
                    <div class="i-box-conn">
                        <el-row>
                            {% for column in addColumns %}
                            {% if column.ColumnKey == "PRI" %}
                            <el-col :span="24" v-show="false">
                                <el-form-item prop="{{ column.FieldName }}" label="{{ column.ColumnComment }}">
                                    <el-input v-model="formData.{{ column.FieldName }}" placeholder="{{ column.ColumnComment }}"
                                              clearable></el-input>
                                </el-form-item>
                            </el-col>
                            {% else %}
                            <el-col :span="24">
                                <el-form-item prop="{{ column.FieldName }}" label="{{ column.ColumnComment }}">
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
                            </el-col>
                            {% endif %}
                            {% endfor %}
                        </el-row>
                    </div>
                </div>
            </div>
        </el-form>
        <div slot="footer">
            <el-button type="info" @click="isShow = false">{{"{"}}{ $t('common.cancel') }{{"}"}}</el-button>
            <el-button type="success" @click="formDataSubmit()" :loading="isLoading"
                       :disabled="formData.id != '' && !canUpdate">{{"{"}}{ $t('common.complete') }{{"}"}}</el-button>
        </div>
    </modal>
</template>

<script>
    import {mapActions} from 'vuex'
    import addOrUpdate from '_cm/mixin/add-or-update'

    export default {
        data() {
            return {
                formData: {
                    {% for column in addColumns %}
                    {{ column.FieldName }}: '' ,
                    {% endfor %}
                },
                dataRule: {
                    {% for column in addColumns %}
                    {% if column.ColumnKey != "PRI" %}
                    {{ column.FieldName }}: [{
                        required: true,
                        message: this.$t('common.inputTip'),
                        trigger: 'blur'
                    }],
                    {% endif %}
                    {% endfor %}
                },
            }
        },
        mixins: [addOrUpdate],
        methods: {
            ...mapActions({
                saveOrUpdateAction: '{{ moduleName }}/{{ fileName }}/saveOrUpdate',
            }),
            init(row) {
                {% for column in addColumns %}
                this.formData.{{ column.FieldName }} = row.{{ column.FieldName }}
                {% endfor %}
            },
            // 表单提交
            formDataSubmit() {
                this.$refs['form'].validate((valid) => {
                    if (valid) {
                        this.isLoading = true
                        this.saveOrUpdateAction({
                        {% for column in addColumns %}
                        {% if column.ColumnKey == "PRI" %}
                        {{ column.FieldName }}: this.formData.{{ column.FieldName }} || undefined,
                        {% else %}
                        {{ column.FieldName }}: this.formData.{{ column.FieldName }},
                        {% endif %}
                        {% endfor %}
                        }).then(() => {
                            this.$message({
                                message: this.$t('common.successTip'),
                                type: 'success',
                                duration: 1500,
                                onClose: () => {
                                    this.isShow = false
                                    this.isLoading = false
                                    this.$emit('refreshDataList')
                                }
                            })
                        }).catch(errMsg => {
                            this.$message.error(errMsg)
                            this.isLoading = false
                        })
                    }
                })
            }
        }
    }
</script>
<style>
</style>
