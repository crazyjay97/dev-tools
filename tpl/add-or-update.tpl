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
                        <el-row>
                            <el-col :span="24">
                                <el-form-item prop="roleName" :label="$t('role.column.roleName')">
                                    <el-input v-model="formData.roleName" :placeholder="$t('role.column.roleName')"
                                              clearable></el-input>
                                </el-form-item>
                            </el-col>
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
                menuList: [],
                tenants: [],
                menuListTreeProps: {
                    label: 'name',
                    children: 'children'
                },
                formData: {
                    id: 0,
                    roleName: '',
                    tenantId: '',
                    remark: ''
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
                infoAction: 'sys/role/info',
                saveOrUpdateAction: 'sys/role/saveOrUpdate',
                menuListAction: 'sys/menu/list',
            }),
            init(id) {
                this.formData.id = id || 0
                this.menuListAction(
                ).then(data => {
                    this.menuList = treeDataTranslate(data, 'menuId')
                }).then(() => {
                    this.visible = true
                    this.$nextTick(() => {
                        this.$refs['form'].resetFields()
                        this.$refs.menuListTree.setCheckedKeys([])
                    })
                }).then(() => {
                    if (this.formData.id) {
                        this.infoAction(this.formData.id)
                            .then(role => {
                                this.formData.roleName = role.roleName
                                this.formData.remark = role.remark
                                this.$refs.menuListTree.setCheckedKeys(role.menuIdList)
                            })
                    }
                })
            },
            // 表单提交
            formDataSubmit() {
                this.$refs['form'].validate((valid) => {
                    if (valid) {
                        this.saveOrUpdateAction({
                            roleId: this.formData.id || undefined,
                            roleName: this.formData.roleName,
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
