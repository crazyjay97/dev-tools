<template>
  <Modal v-model="isShow" title="Code" width="1200px">
    <Tabs @on-click="queryCode" v-model="currentTabName">
      <TabPane v-for="r in tplNames" :name="r" :label="r" :key="r">
        <textarea style="height: 500px;width: 100%;overflow: auto">{{code}}
        </textarea>
      </TabPane>
    </Tabs>
  </Modal>
</template>
<script>
  import {mapState} from 'vuex'

  export default {
    data() {
      return {
        isShow: false,
        table: {},
        tplNames: [],
        currentTabName: '',
        code: ''
      }
    },
    computed: {
      ...mapState(["tables"]),
      ...mapState(["settings"]),
    },
    methods: {
      show(row) {
        this.isShow = true
        this.table = row
        this.queryTplNames()
      },
      queryTplNames() {
        this.$ajax({
          url: "/generator/query/tplNames",
          method: "get"
        }).then(({data}) => {
          this.tplNames = data
          this.queryCode(data[0])
        })
      },
      queryCode(name) {
        let tbs = [...this.tables];
        tbs.filter(r => r.tableName == this.tableName).forEach(t => t.joinTables = t.joinTables.filter(({tableName, selfColumn, joinColumn, alias, description}) =>
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
        );
        let data = {
          mainPath: this.settings.mainPath,
          packageName: this.settings.pkg,
          moduleName: this.settings.moduleName,
          authorName: this.settings.author,
          emailAddress: this.settings.email,
          removePrefix: this.settings.isRemovePrefix,
          autoSettingModuleName: this.settings.autoSettingModuleName,
          templateName: name,
          modules: tbs
        };
        this.$ajax({
          url: "/generator/query/code",
          method: "post",
          data: data
        }).then(({data}) => {
          this.code = data
        })
      }
    }
  }
</script>
