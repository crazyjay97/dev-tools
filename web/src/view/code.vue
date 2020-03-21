<template>
  <Modal v-model="isShow" title="Code" width="1200px">
    <Tabs @on-click="queryCode" v-model="currentTabName">
      <TabPane v-for="r in tplNames" :name="r" :label="r" :key="r">
        <div style="height: 500px;width: 100%;overflow: auto;position: relative;" v-if="isLoadingCode">
          <Spin fix size="large"></Spin>
        </div>
        <div style="height: 500px;width: 100%;overflow: auto" v-else="!isLoadingCode">
          <pre v-highlightjs="code" contenteditable="true"><code class="plaintext"></code></pre>
        </div>
      </TabPane>
    </Tabs>
    <div slot="footer">
      <Button @click="isShow = false" size="large" type="primary">退出</Button>
    </div>
  </Modal>
</template>
<script>
  import {mapState} from 'vuex'


  export default {
    data() {
      return {
        isLoadingCode: false,
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
          this.currentTabName = data[0]
          this.tplNames = data
          this.queryCode(data[0])
        })
      },
      queryCode(name) {
        this.isLoadingCode = true;
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
          setTimeout(() => {
            this.code = data
            this.isLoadingCode = false;
          }, 1000)
        })
      }
    }
  }
</script>
