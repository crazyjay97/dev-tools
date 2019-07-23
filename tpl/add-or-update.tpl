<template>
    <modal v-model="isShow" width="500px">
        <el-form :model="formData" ref="form" label-width="140px" :rules="dataRule">
            <div>
                <div class="i-box">
                    <div class="i-box-title">
                        <icon-svg name="line"></icon-svg>
                        {{"{"}}{ $t('role.modal.title') }{{"}"}}
                        <h5>{{"{"}}{ $t("common.tip") }{{"}"}}<span>*</span>{{"{"}}{ $t("common.tippp") }{{"}"}}</h5>
                    </div>
                    <div class="i-box-conn">
                        <el-row>{% for column in addColumns %}{% if column.ColumnKey == "PRI" %}
                            <el-col :span="24" v-show="false">
                                <el-form-item prop="{{ column.FieldName }}" label="{{ column.ColumnComment }}">
                                    <el-input v-model="formData.{{ column.FieldName }}" placeholder="{{ column.ColumnComment }}"
                                              clearable></el-input>
                                </el-form-item>
                            </el-col>{% else %}
                            <el-col :span="24">
                                <el-form-item prop="{{ column.FieldName }}" label="{{ column.ColumnComment }}">
                                    <el-input v-model="formData.{{ column.FieldName }}" placeholder="{{ column.ColumnComment }}"
                                              clearable></el-input>
                                </el-form-item>
                            </el-col>{% endif %}{% endfor %}
                        </el-row>
                    </div>
                </div>
            </div>
        </el-form>
        <div slot="footer">
            <el-button type="info" @click="isShow = false">{{"{"}}{ $t('common.cancel') }{{"}"}}</el-button>
            <el-button type="success" @click="formDataSubmit()">{{"{"}}{ $t('common.complete') }{{"}"}}</el-button>
        </div>
    </modal>
</template>

<script>
    import {treeDataTranslate} from '@/utils'
    import Modal from "@/components/modal/modal"
    import "@/assets/scss/ui.scss"
    import {mapActions} from 'vuex'

    export default {
        props: {
            value: {
                required: true,
                type: Boolean
            }
        },
        watch: {
            value(val) {
                this.isShow = val;
            },
            isShow(val) {
                if (!val) {
                    this.$refs.form.clearValidate();
                    this.$refs.form.resetFields();
                }
                this.$emit("input", val);
            }
        },
        components: {
            Modal
        },
        data() {
            return {
                isShow: false,
                formData: {  {% for column in addColumns %}
                    {{ column.FieldName }}: '' ,{% endfor %}
                },
                dataRule: {
                    roleName: [{
                        required: true,
                        message: this.$t('common.inputTip'),
                        trigger: 'blur'
                    }]
                },
            }
        },
        methods: {
            ...mapActions({
                saveOrUpdateAction: '{{ moduleName }}/{{ fileName }}/saveOrUpdate',
            }),
            init(row) { {% for column in addColumns %}
                this.formData.{{ column.FieldName }} = row.{{ column.FieldName }} {% endfor %}
            },
            // 表单提交
            formDataSubmit() {
                this.$refs['form'].validate((valid) => {
                    if (valid) {
                        this.saveOrUpdateAction({ {% for column in addColumns %}{% if column.ColumnKey == "PRI" %}
                        {{ column.FieldName }}: this.formData.{{ column.FieldName }} || undefined,{% else %}
                        {{ column.FieldName }}: this.formData.{{ column.FieldName }} ,{% endif %}{% endfor %}
                        }).then(() => {
                            this.$message({
                                message: this.$t('common.successTip'),
                                type: 'success',
                                duration: 1500,
                                onClose: () => {
                                    this.isShow = false
                                    this.$emit('refreshDataList')
                                }
                            })
                        }).catch(errMsg => this.$message.error(errMsg))
                    }
                })
            }
        }
    }
</script>
<style>
</style>
