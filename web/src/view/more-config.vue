<template>
  <Modal @on-ok="commitConfig" ok-text="Commit" cancel-text="Cancel" v-model="show" title="Table Config" width="1200px"
         ref="config_modal" @onCancel="show = false">
    <Card class="columnSettingCols" :bordered="false" :dis-hover="true" title="Column Setting">
      <Table :stripe="true" :columns="columnSettingCols" :data="tmpColumnSetting" :draggable="true"
             @on-drag-drop="onRowDrag">
        <template slot-scope="{ row,index }" slot="column">
          <span>{{ columnSetting[index].column }}</span>
        </template>
        <template slot-scope="{ row,index }" slot="comment">
          <Input v-model="columnSetting[index].columnDesc"
                 placeholder="Comment">
          </Input>
        </template>
        <template slot-scope="{ row,index }" slot="needShow">
          <Switch v-model="columnSetting[index].needShow"/>
        </template>
        <template slot-scope="{ row,index }" slot="needAdd">
          <Switch v-model="columnSetting[index].needAdd"/>
        </template>
        <template slot-scope="{ row,index }" slot="needFilter">
          <Switch v-model="columnSetting[index].needFilter"/>
        </template>
        <template slot-scope="{ row,index }" slot="showMode">
          <Select v-model="columnSetting[index].showMode" filterable
                  placeholder="Show Mode">
            <Option :value="0">Input</Option>
            <Option :value="1">Select</Option>
          </Select>
        </template>
        <template slot-scope="{ row,index }" slot="dictionaryLabel">
          <Input v-model="columnSetting[index].dictionaryLabel"
                 placeholder="type:label">
          </Input>
        </template>
        <template slot-scope="{ row,index }" slot="dictionaryValue">
          <Input v-model="columnSetting[index].dictionaryValue"
                 placeholder="v:l,v:l,v:l">
          </Input>
        </template>
      </Table>
    </Card>
    <Card class="multipleTableCols" :bordered="false" :dis-hover="true" title="Multi Table">
      <Table :columns="multipleTableCols" :data="tmpMultipleTables">
        <template slot-scope="{ row,index }" slot="selfColumn">
          <Select v-model="multipleTables[index].selfColumn" filterable placeholder="Column">
            <Option v-for="item in selfCols" :value="item.columnName" :key="item.columnName">{{ item.columnName }}
            </Option>
          </Select>
        </template>
        <template slot-scope="{ row,index }" slot="tableName">
          <Select v-model="multipleTables[index].tableName" filterable @on-change="tableChange(multipleTables[index])"
                  placeholder="Linked Table">
            <Option v-for="item in tbs" :value="item.tableName" :key="item.tableName">{{ item.tableName }}
            </Option>
          </Select>
        </template>
        <template slot-scope="{ row,index }" slot="joinColumn">
          <Select v-model="multipleTables[index].joinColumn" filterable
                  placeholder="Linked Column">
            <Option v-for="item in multipleTables[index].linkedColumns" :value="item.columnName"
                    :key="item.columnName">{{ item.columnName }}
            </Option>
          </Select>
        </template>
        <template slot-scope="{ row,index }" slot="searchColumn">
          <Select v-model="multipleTables[index].searchColumn" filterable
                  placeholder="Search Column">
            <Option v-for="item in multipleTables[index].linkedColumns" :value="item.columnName"
                    :key="item.columnName">{{ item.columnName }}
            </Option>
          </Select>
        </template>
        <template slot-scope="{ row,index }" slot="alias">
          <Input v-model="multipleTables[index].alias" filterable placeholder="Alias">
          </Input>
        </template>
        <template slot-scope="{ row,index }" slot="description">
          <Input v-model="multipleTables[index].description"
                 :disabled="!multipleTables[index].selfColumn || !multipleTables[index].tableName || !multipleTables[index].joinColumn ||!multipleTables[index].alias"
                 @on-blur="addRow(index)" filterable
                 placeholder="Description">
          </Input>
        </template>
        <template slot-scope="{ row,index }" slot="action">
          <Button v-if="(index+1) != multipleTables.length" icon="ios-backspace" type="primary"
                  @click="removeRows(index)">
            Remove
          </Button>
        </template>
      </Table>
    </Card>
  </Modal>
</template>
<script>
  import {mapMutations, mapState} from 'vuex'

  export default {
    data() {
      return {
        show: false,
        selfCols: [],
        tableName: '',
        tbs: [],
        multipleTableCols: [
          {
            title: "Column",
            slot: 'selfColumn'

          },
          {
            title: "Linked Table",
            slot: 'tableName'
          },
          {
            title: "Linked Column",
            slot: 'joinColumn'
          },
          {
            title: "Search Column",
            slot: 'searchColumn'
          },
          {
            title: "Alias",
            slot: 'alias',
            width: 130
          },
          {
            title: "Description",
            slot: 'description',
            width: 130
          },
          {
            title: "Action",
            slot: 'action'
          },
        ],
        multipleTables: [],
        tmpMultipleTables: [],
        columnSettingCols: [
          {
            title: "Column",
            slot: 'column'

          },
          {
            title: "Comment",
            slot: "comment",
          },
          {
            title: "Need Show",
            slot: 'needShow'
          },
          {
            title: "Need Add",
            slot: 'needAdd'
          },
          {
            title: "Need Filter",
            slot: 'needFilter'
          },
          {
            title: "Show Mode",
            slot: 'showMode',
            width: 130
          },
          {
            title: "Dictionary Label",
            slot: 'dictionaryLabel',
            width: 130
          },
          {
            title: "Dictionary Value",
            slot: 'dictionaryValue'
          },
        ],
        columnSetting: [],
        tmpColumnSetting: [],
      }
    },
    computed: {
      ...mapState(["tables"]),
    },
    methods: {
      ...mapMutations(['updateTables']),
      onRowDrag(index1, index2) {
        let tmpRow = this.columnSetting[index1]
        this.columnSetting[index1] = this.columnSetting[index2]
        this.columnSetting[index2] = tmpRow
        this.cloneColumnSetting()
      },
      removeRows(index) {
        this.multipleTables.splice(index, 1)
        this.cloneTables()
      },
      tableChange(row) {
        this.queryColumn(row.tableName, cols => {
          {
            row.linkedColumns = cols
            this.multipleTables.push(-1)
            this.multipleTables.remove(-1)
          }
        })
      },
      addRow(index) {
        let lastRow = this.multipleTables[this.multipleTables.length - 1]
        if (lastRow.tableName && lastRow.selfColumn && lastRow.joinColumn && lastRow.alias && lastRow.description) {
          this.multipleTables.push({})
          this.cloneTables()
        }
      },
      init(row) {
        this.show = true
        this.tableName = row.tableName
        let tmpTable = this.tables.filter(t => t.tableName == row.tableName)[0]
        this.queryColumn(row.tableName, cols => {
          this.selfCols = cols
          //init column setting
          this.columnSetting = tmpTable && tmpTable.columnSetting && tmpTable.columnSetting.length > 0 ? tmpTable.columnSetting : cols.map(c => {
            return {
              column: c.columnName,
              columnDesc: c.columnComment,
              needShow: true,
              needAdd: true,
              needFilter: false,
              showMode: 0,
            }
          })
          this.cloneColumnSetting()
          if (tmpTable) {
            this.addFields = tmpTable.addFields
            this.searchFields = tmpTable.searchFields
            this.multipleTables = tmpTable.joinTables ? tmpTable.joinTables : []
          }
          if (this.multipleTables.length == 0) {
            this.multipleTables.push({
              tableName: '',
              selfColumn: '',
              joinColumn: '',
              searchColumn: '',
              alias: '',
              description: '',
            })
          } else {
            let {alias, description, joinColumn, selfColumn, tableName} = this.multipleTables[this.multipleTables.length - 1]
            if (alias && description && joinColumn && selfColumn && tableName) {
              this.multipleTables.push({
                tableName: '',
                selfColumn: '',
                joinColumn: '',
                searchColumn: '',
                alias: '',
                description: '',
              })
            }
          }
          this.cloneTables()
        })
        this.queryAllTable()
      },
      cloneTables() {
        this.tmpMultipleTables = this.multipleTables.map(() => {
          return {
            tableName: '',
            selfColumn: '',
            joinColumn: '',
            searchColumn: '',
            alias: '',
            description: '',
          }
        });
      },
      cloneColumnSetting() {
        this.tmpColumnSetting = this.columnSetting.map(c => {
          return {
            column: "",
            columnDesc: "",
            needShow: "",
            needAdd: "",
            needFilter: "",
          }
        })
      },
      queryColumn(tableName, rollback) {
        this.$ajax({
          url: "/generator/query/columns",
          method: "get",
          params: {
            tableName: tableName
          }
        }).then(({data}) => {
          rollback(data.list)
        })
      },
      queryAllTable() {
        this.$ajax({
          url: "/generator/query/all",
          method: "get"
        }).then(({data}) => {
          this.tbs = data.list
        })
      },
      commitConfig() {
        this.updateTables([...this.tables.filter(t => t.tableName != this.tableName), {
          tableName: this.tableName,
          joinTables: this.multipleTables.filter(({alias, description, joinColumn, searchColumn, selfColumn, tableName}) => alias && description && joinColumn && searchColumn && selfColumn && tableName),
          columnSetting: this.columnSetting
        }])
      },
    },
  }
</script>
<style>

  .columnSettingCols .ivu-table-body {
    height: 300px;
    overflow: auto;
  }

  .multipleTableCols .ivu-table-wrapper {
    overflow: visible !important;
  }
</style>
