<template>
  <div>
    <Card>
      <Row>
        <Col span="16">Table Name:
          <Input
            v-model="tableName"
            placeholder="Typing Table Name"
            style="width: 300px"
            :clearable="true"
          />
          <Button type="primary" icon="ios-search" @click="page=1;pageData()">Search</Button>
        </Col>
        <Col span="6">
          <Dropdown>
            <a href="javascript:void(0)">
              <Tag type="dot" color="primary">Have Been Added {{ tablesLength }}</Tag>
            </a>
            <DropdownMenu slot="list">
              <DropdownItem v-for="{tableName} in tables" :key="tableName">
                <div style="height: 25px">
                  <Row>
                    <Col span="16">
                      <span style="float: left;line-height: 25px">{{ tableName }}</span>
                    </Col>
                    <Col span="8">
                      <Button type="text" style="position: relative; bottom:4px;color: red;"
                              @click="removeTable(tableName)">Remove
                      </Button>
                    </Col>
                  </Row>
                </div>
              </DropdownItem>
            </DropdownMenu>
          </Dropdown>
          <Button shape="circle" @click="gen">Generate</Button>
        </Col>
        <Col span="2">
          <Poptip title="Config" placement="left" @on-popper-hide="saveConfig">
            <div slot="content" style="width:400px">
              <Form :label-width="100">
                <FormItem label="Main Path">
                  <Input v-model="formData.mainPath"></Input>
                </FormItem>
                <FormItem label="Package Name">
                  <Input v-model="formData.pkg"></Input>
                </FormItem>
                <FormItem label="Author Name">
                  <Input v-model="formData.author"></Input>
                </FormItem>
                <FormItem label="Email Address">
                  <Input v-model="formData.email"></Input>
                </FormItem>
                <FormItem label="Module Name">
                  <Checkbox v-model="formData.autoSettingModuleName">Auto Setting</Checkbox>
                  <Input
                    v-if="!formData.autoSettingModuleName"
                    v-model="formData.moduleName"
                    style="width:204px"
                  ></Input>
                </FormItem>
                <FormItem label="Remove Prefix">
                  <Switch size="large" v-model="formData.isRemovePrefix"></Switch>
                </FormItem>
              </Form>
            </div>
            <Button type="text" icon="ios-cog-outline" style="float:right;color:blue">Config</Button>
          </Poptip>
        </Col>
      </Row>
    </Card>
    <Table :columns="columns" height="600" :data="tbs" no-data-text="Can Not Find Table">
      <template slot-scope="{ row }" slot="operator">
        <Button shape="circle" icon="ios-more" @click="openConfig(row)">Add</Button>
      </template>
    </Table>
    <Page
      :total="total"
      :page-size="limit"
      show-sizer
      style="float:right"
      @on-change="pageHandle"
      @on-page-size-change="pageSizeHandle"
    />
    <more-config ref="configModal"></more-config>
  </div>
</template>

<script>
  import MoreConfig from './more-config'
  import {mapState, mapMutations} from 'vuex'

  export default {
    components: {
      MoreConfig
    },
    data() {
      return {
        tableName: "",
        formData: {
          mainPath: "com.zyiot.tet",
          pkg: "com.zyiot.tet.modules",
          author: "professor X",
          email: "professorX@mail.com",
          isRemovePrefix: true,
          moduleName: "",
          autoSettingModuleName: true
        },
        page: 1,
        limit: 10,
        total: 0,
        tbs: [],
        columns: [
          {
            title: "Table Name",
            key: "tableName"
          },
          {
            title: "Engine",
            key: "engine"
          },
          {
            title: "Comment",
            key: "tableComment"
          },
          {
            title: "Create Time",
            key: "createTime"
          },
          {
            title: "Operator",
            slot: 'operator',
          }
        ]
      };
    },
    created() {
      errTest = test
    },
    computed: {
      ...mapState(["tables"]),
      tablesLength: {
        get() {
          return this.tables.length
        }
      },
    },
    mounted() {
      this.pageData();
      if (window.localStorage["config"]) {
        let config = JSON.parse(window.localStorage["config"]);
        this.formData.mainPath = config.mainPath;
        this.formData.pkg = config.pkg;
        this.formData.moduleName = config.moduleName;
        this.formData.author = config.author;
        this.formData.email = config.email;
        this.formData.isRemovePrefix = config.isRemovePrefix;
        this.formData.autoSettingModuleName = config.autoSettingModuleName;
      }
    },
    methods: {
      ...mapMutations(['updateTables']),
      removeTable(tableName) {
        this.updateTables(this.tables.filter(t => t.tableName != tableName))
      },
      pageSizeHandle(s) {
        this.limit = s;
        this.pageData();
      },
      pageHandle(p) {
        this.page = p;
        this.pageData();
      },
      saveConfig() {
        window.localStorage["config"] = JSON.stringify(this.formData);
      },
      openConfig(row) {
        this.$refs.configModal.init(row)
      },
      gen() {
        let tbs = [...this.tables]
        tbs.forEach(t => t.joinTables = t.joinTables.filter(({tableName, selfColumn, joinColumn, alias, description}) =>
          tableName && selfColumn && joinColumn && alias && description).map(jt => {
            return {
              tableName: jt.tableName,
              selfColumn: jt.selfColumn,
              joinColumn: jt.joinColumn,
              searchColumn: jt.searchColumn,
              alias: jt.alias,
              description: jt.description,
            }
          })
        )
        let data = {
          mainPath: this.formData.mainPath,
          packageName: this.formData.pkg,
          moduleName: this.formData.moduleName,
          authorName: this.formData.author,
          emailAddress: this.formData.email,
          removePrefix: this.formData.isRemovePrefix,
          autoSettingModuleName: this.formData.autoSettingModuleName,
          modules: tbs
        }
        this.$ajax({
          url: "/generator/gen",
          method: "post",
          responseType: 'blob',
          data: data// JSON.stringify(data)
        }).then((res) => {
          let blob = new Blob([res.data])
          if (window.navigator.msSaveOrOpenBlob) {
            navigator.msSaveBlob(blob, "code.zip")
          } else {
            let link = document.createElement("a")
            let evt = document.createEvent("HTMLEvents")
            evt.initEvent("click", false, false)
            link.href = URL.createObjectURL(blob)
            link.download = "code.zip"
            link.style.display = "none"
            document.body.appendChild(link)
            link.click()
            window.URL.revokeObjectURL(link.href)
          }
        })
      },
      pageData() {
        this.$ajax({
          url: "/generator/list",
          method: "get",
          params: {
            page: this.page,
            limit: this.limit,
            tableName: this.tableName
          }
        }).then(({data}) => {
          this.total = data.total
          this.tbs = data.list
        })
      }
    }
  };
</script>
<style>
  .ivu-table-wrapper {
    overflow: auto;
  }
</style>
