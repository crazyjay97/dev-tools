webpackJsonp([1],{"92qf":function(e,t){},IpcW:function(e,t){},NHnr:function(e,t,a){"use strict";Object.defineProperty(t,"__esModule",{value:!0});var n=a("xd7I"),l={render:function(){var e=this.$createElement,t=this._self._c||e;return t("div",{attrs:{id:"app"}},[t("router-view")],1)},staticRenderFns:[]};var o=a("C7Lr")({name:"App"},l,!1,function(e){a("uPO5")},null,null).exports,i=a("e1F6"),s={render:function(){var e=this.$createElement,t=this._self._c||e;return t("div",[t("Layout",[t("Header",[t("h1",{staticStyle:{color:"#c4c4c4"}},[this._v("\n        Code Generator\n      ")])]),this._v(" "),t("Content",[t("router-view")],1)],1)],1)},staticRenderFns:[]};var r=a("C7Lr")({},s,!1,function(e){a("f7oQ")},null,null).exports,u=a("IHPB"),m=a.n(u),c=a("3cXf"),d=a.n(c),p=a("4YfN"),f=a.n(p),h=a("2bvH"),b={data:function(){return{show:!1,selfCols:[],tableName:"",tbs:[],multipleTableCols:[{title:"Column",slot:"selfColumn"},{title:"Linked Table",slot:"tableName"},{title:"Linked Column",slot:"joinColumn"},{title:"Search Column",slot:"searchColumn"},{title:"Alias",slot:"alias",width:130},{title:"Description",slot:"description",width:130},{title:"Action",slot:"action"}],multipleTables:[],tmpMultipleTables:[],columnSettingCols:[{title:"Column",slot:"column"},{title:"Comment",slot:"comment"},{title:"Need Show",slot:"needShow"},{title:"Need Add",slot:"needAdd"},{title:"Need Filter",slot:"needFilter"},{title:"Show Mode",slot:"showMode",width:130},{title:"Dictionary Label",slot:"dictionaryLabel",width:130},{title:"Dictionary Value",slot:"dictionaryValue"}],columnSetting:[],tmpColumnSetting:[]}},computed:f()({},Object(h.c)(["tables"])),methods:f()({},Object(h.b)(["updateTables"]),{onRowDrag:function(e,t){var a=this.columnSetting[e];this.columnSetting[e]=this.columnSetting[t],this.columnSetting[t]=a,this.cloneColumnSetting()},removeRows:function(e){this.multipleTables.splice(e,1),this.cloneTables()},tableChange:function(e){var t=this;this.queryColumn(e.tableName,function(a){e.linkedColumns=a,t.multipleTables.push(-1),t.multipleTables.remove(-1)})},addRow:function(e){var t=this.multipleTables[this.multipleTables.length-1];t.tableName&&t.selfColumn&&t.joinColumn&&t.alias&&t.description&&(this.multipleTables.push({}),this.cloneTables())},init:function(e){var t=this;this.show=!0,this.tableName=e.tableName;var a=this.tables.filter(function(t){return t.tableName==e.tableName})[0];this.queryColumn(e.tableName,function(e){if(t.selfCols=e,t.columnSetting=a&&a.columnSetting&&a.columnSetting.length>0?a.columnSetting:e.map(function(e){return{column:e.columnName,columnDesc:e.columnComment,needShow:!0,needAdd:!0,needFilter:!1,showMode:0}}),t.cloneColumnSetting(),a&&(t.addFields=a.addFields,t.searchFields=a.searchFields,t.multipleTables=a.joinTables?a.joinTables:[]),0==t.multipleTables.length)t.multipleTables.push({tableName:"",selfColumn:"",joinColumn:"",searchColumn:"",alias:"",description:""});else{var n=t.multipleTables[t.multipleTables.length-1],l=n.alias,o=n.description,i=n.joinColumn,s=n.selfColumn,r=n.tableName;l&&o&&i&&s&&r&&t.multipleTables.push({tableName:"",selfColumn:"",joinColumn:"",searchColumn:"",alias:"",description:""})}t.cloneTables()}),this.queryAllTable()},cloneTables:function(){this.tmpMultipleTables=this.multipleTables.map(function(){return{tableName:"",selfColumn:"",joinColumn:"",searchColumn:"",alias:"",description:""}})},cloneColumnSetting:function(){this.tmpColumnSetting=this.columnSetting.map(function(e){return{column:"",columnDesc:"",needShow:"",needAdd:"",needFilter:""}})},queryColumn:function(e,t){this.$ajax({url:"/generator/query/columns",method:"get",params:{tableName:e}}).then(function(e){var a=e.data;t(a.list)})},queryAllTable:function(){var e=this;this.$ajax({url:"/generator/query/all",method:"get"}).then(function(t){var a=t.data;e.tbs=a.list})},commitConfig:function(){var e=this;this.updateTables([].concat(m()(this.tables.filter(function(t){return t.tableName!=e.tableName})),[{tableName:this.tableName,joinTables:this.multipleTables,columnSetting:this.columnSetting}]))}})},v={render:function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("Modal",{ref:"config_modal",attrs:{"ok-text":"Commit","cancel-text":"Cancel",title:"Table Config",width:"1200px"},on:{"on-ok":e.commitConfig,onCancel:function(t){e.show=!1}},model:{value:e.show,callback:function(t){e.show=t},expression:"show"}},[a("Card",{staticClass:"columnSettingCols",attrs:{bordered:!1,"dis-hover":!0,title:"Column Setting"}},[a("Table",{attrs:{stripe:!0,columns:e.columnSettingCols,data:e.tmpColumnSetting,draggable:!0},on:{"on-drag-drop":e.onRowDrag},scopedSlots:e._u([{key:"column",fn:function(t){t.row;var n=t.index;return[a("span",[e._v(e._s(e.columnSetting[n].column))])]}},{key:"comment",fn:function(t){t.row;var n=t.index;return[a("Input",{attrs:{placeholder:"Dictionary Label"},model:{value:e.columnSetting[n].columnDesc,callback:function(t){e.$set(e.columnSetting[n],"columnDesc",t)},expression:"columnSetting[index].columnDesc"}})]}},{key:"needShow",fn:function(t){t.row;var n=t.index;return[a("i-switch",{model:{value:e.columnSetting[n].needShow,callback:function(t){e.$set(e.columnSetting[n],"needShow",t)},expression:"columnSetting[index].needShow"}})]}},{key:"needAdd",fn:function(t){t.row;var n=t.index;return[a("i-switch",{model:{value:e.columnSetting[n].needAdd,callback:function(t){e.$set(e.columnSetting[n],"needAdd",t)},expression:"columnSetting[index].needAdd"}})]}},{key:"needFilter",fn:function(t){t.row;var n=t.index;return[a("i-switch",{model:{value:e.columnSetting[n].needFilter,callback:function(t){e.$set(e.columnSetting[n],"needFilter",t)},expression:"columnSetting[index].needFilter"}})]}},{key:"showMode",fn:function(t){t.row;var n=t.index;return[a("Select",{attrs:{filterable:"",placeholder:"Show Mode"},model:{value:e.columnSetting[n].showMode,callback:function(t){e.$set(e.columnSetting[n],"showMode",t)},expression:"columnSetting[index].showMode"}},[a("Option",{attrs:{value:0}},[e._v("Input")]),e._v(" "),a("Option",{attrs:{value:1}},[e._v("Select")])],1)]}},{key:"dictionaryLabel",fn:function(t){t.row;var n=t.index;return[a("Input",{attrs:{placeholder:"type:label"},model:{value:e.columnSetting[n].dictionaryLabel,callback:function(t){e.$set(e.columnSetting[n],"dictionaryLabel",t)},expression:"columnSetting[index].dictionaryLabel"}})]}},{key:"dictionaryValue",fn:function(t){t.row;var n=t.index;return[a("Input",{attrs:{placeholder:"v:l,v:l,v:l"},model:{value:e.columnSetting[n].dictionaryValue,callback:function(t){e.$set(e.columnSetting[n],"dictionaryValue",t)},expression:"columnSetting[index].dictionaryValue"}})]}}])})],1),e._v(" "),a("Card",{attrs:{bordered:!1,"dis-hover":!0,title:"Multi Table"}},[a("Table",{attrs:{columns:e.multipleTableCols,data:e.tmpMultipleTables},scopedSlots:e._u([{key:"selfColumn",fn:function(t){t.row;var n=t.index;return[a("Select",{attrs:{filterable:"",placeholder:"Column"},model:{value:e.multipleTables[n].selfColumn,callback:function(t){e.$set(e.multipleTables[n],"selfColumn",t)},expression:"multipleTables[index].selfColumn"}},e._l(e.selfCols,function(t){return a("Option",{key:t.columnName,attrs:{value:t.columnName}},[e._v(e._s(t.columnName)+"\n          ")])}),1)]}},{key:"tableName",fn:function(t){t.row;var n=t.index;return[a("Select",{attrs:{filterable:"",placeholder:"Linked Table"},on:{"on-change":function(t){return e.tableChange(e.multipleTables[n])}},model:{value:e.multipleTables[n].tableName,callback:function(t){e.$set(e.multipleTables[n],"tableName",t)},expression:"multipleTables[index].tableName"}},e._l(e.tbs,function(t){return a("Option",{key:t.tableName,attrs:{value:t.tableName}},[e._v(e._s(t.tableName)+"\n          ")])}),1)]}},{key:"joinColumn",fn:function(t){t.row;var n=t.index;return[a("Select",{attrs:{filterable:"",placeholder:"Linked Column"},model:{value:e.multipleTables[n].joinColumn,callback:function(t){e.$set(e.multipleTables[n],"joinColumn",t)},expression:"multipleTables[index].joinColumn"}},e._l(e.multipleTables[n].linkedColumns,function(t){return a("Option",{key:t.columnName,attrs:{value:t.columnName}},[e._v(e._s(t.columnName)+"\n          ")])}),1)]}},{key:"searchColumn",fn:function(t){t.row;var n=t.index;return[a("Select",{attrs:{filterable:"",placeholder:"Search Column"},model:{value:e.multipleTables[n].searchColumn,callback:function(t){e.$set(e.multipleTables[n],"searchColumn",t)},expression:"multipleTables[index].searchColumn"}},e._l(e.multipleTables[n].linkedColumns,function(t){return a("Option",{key:t.columnName,attrs:{value:t.columnName}},[e._v(e._s(t.columnName)+"\n          ")])}),1)]}},{key:"alias",fn:function(t){t.row;var n=t.index;return[a("Input",{attrs:{filterable:"",placeholder:"Alias"},model:{value:e.multipleTables[n].alias,callback:function(t){e.$set(e.multipleTables[n],"alias",t)},expression:"multipleTables[index].alias"}})]}},{key:"description",fn:function(t){t.row;var n=t.index;return[a("Input",{attrs:{disabled:!(e.multipleTables[n].selfColumn&&e.multipleTables[n].tableName&&e.multipleTables[n].joinColumn&&e.multipleTables[n].alias),filterable:"",placeholder:"Description"},on:{"on-blur":function(t){return e.addRow(n)}},model:{value:e.multipleTables[n].description,callback:function(t){e.$set(e.multipleTables[n],"description",t)},expression:"multipleTables[index].description"}})]}},{key:"action",fn:function(t){t.row;var n=t.index;return[n+1!=e.multipleTables.length?a("Button",{attrs:{icon:"ios-backspace",type:"primary"},on:{click:function(t){return e.removeRows(n)}}},[e._v("\n          Remove\n        ")]):e._e()]}}])})],1)],1)},staticRenderFns:[]};var g={components:{MoreConfig:a("C7Lr")(b,v,!1,function(e){a("IpcW")},null,null).exports},data:function(){return{tableName:"",formData:{mainPath:"com.zyiot.tet",pkg:"com.zyiot.tet.modules",author:"professor X",email:"professorX@mail.com",isRemovePrefix:!0,moduleName:"",autoSettingModuleName:!0},page:1,limit:10,total:0,tbs:[],columns:[{title:"Table Name",key:"tableName"},{title:"Engine",key:"engine"},{title:"Comment",key:"tableComment"},{title:"Create Time",key:"createTime"},{title:"Operator",slot:"operator"}]}},created:function(){errTest=test},computed:f()({},Object(h.c)(["tables"]),{tablesLength:{get:function(){return this.tables.length}}}),mounted:function(){if(this.pageData(),window.localStorage.config){var e=JSON.parse(window.localStorage.config);this.formData.mainPath=e.mainPath,this.formData.pkg=e.pkg,this.formData.moduleName=e.moduleName,this.formData.author=e.author,this.formData.email=e.email,this.formData.isRemovePrefix=e.isRemovePrefix,this.formData.autoSettingModuleName=e.autoSettingModuleName}},methods:f()({},Object(h.b)(["updateTables"]),{removeTable:function(e){this.updateTables(this.tables.filter(function(t){return t.tableName!=e}))},pageSizeHandle:function(e){this.limit=e,this.pageData()},pageHandle:function(e){this.page=e,this.pageData()},saveConfig:function(){window.localStorage.config=d()(this.formData)},openConfig:function(e){this.$refs.configModal.init(e)},gen:function(){var e=[].concat(m()(this.tables));e.forEach(function(e){return e.joinTables=e.joinTables.filter(function(e){var t=e.tableName,a=e.selfColumn,n=e.joinColumn,l=e.alias,o=e.description;return t&&a&&n&&l&&o}).map(function(e){return{tableName:e.tableName,selfColumn:e.selfColumn,joinColumn:e.joinColumn,searchColumn:e.searchColumn,alias:e.alias,description:e.description}})});var t={mainPath:this.formData.mainPath,packageName:this.formData.pkg,moduleName:this.formData.moduleName,authorName:this.formData.author,emailAddress:this.formData.email,removePrefix:this.formData.isRemovePrefix,autoSettingModuleName:this.formData.autoSettingModuleName,modules:e};this.$ajax({url:"/generator/gen",method:"post",responseType:"blob",data:t}).then(function(e){var t=new Blob([e.data]);if(window.navigator.msSaveOrOpenBlob)navigator.msSaveBlob(t,"code.zip");else{var a=document.createElement("a");document.createEvent("HTMLEvents").initEvent("click",!1,!1),a.href=URL.createObjectURL(t),a.download="code.zip",a.style.display="none",document.body.appendChild(a),a.click(),window.URL.revokeObjectURL(a.href)}})},pageData:function(){var e=this;this.$ajax({url:"/generator/list",method:"get",params:{page:this.page,limit:this.limit,tableName:this.tableName}}).then(function(t){var a=t.data;e.total=a.total,e.tbs=a.list})}})},C={render:function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("div",[a("Card",[a("Row",[a("Col",{attrs:{span:"16"}},[e._v("Table Name:\n        "),a("Input",{staticStyle:{width:"300px"},attrs:{placeholder:"Typing Table Name",clearable:!0},model:{value:e.tableName,callback:function(t){e.tableName=t},expression:"tableName"}}),e._v(" "),a("Button",{attrs:{type:"primary",icon:"ios-search"},on:{click:function(t){e.page=1,e.pageData()}}},[e._v("Search")])],1),e._v(" "),a("Col",{attrs:{span:"6"}},[a("Dropdown",[a("a",{attrs:{href:"javascript:void(0)"}},[a("Tag",{attrs:{type:"dot",color:"primary"}},[e._v("Have Been Added "+e._s(e.tablesLength))])],1),e._v(" "),a("DropdownMenu",{attrs:{slot:"list"},slot:"list"},e._l(e.tables,function(t){var n=t.tableName;return a("DropdownItem",{key:n},[a("div",{staticStyle:{height:"25px"}},[a("Row",[a("Col",{attrs:{span:"16"}},[a("span",{staticStyle:{float:"left","line-height":"25px"}},[e._v(e._s(n))])]),e._v(" "),a("Col",{attrs:{span:"8"}},[a("Button",{staticStyle:{position:"relative",bottom:"4px",color:"red"},attrs:{type:"text"},on:{click:function(t){return e.removeTable(n)}}},[e._v("Remove\n                    ")])],1)],1)],1)])}),1)],1),e._v(" "),a("Button",{attrs:{shape:"circle"},on:{click:e.gen}},[e._v("Generate")])],1),e._v(" "),a("Col",{attrs:{span:"2"}},[a("Poptip",{attrs:{title:"Config",placement:"left"},on:{"on-popper-hide":e.saveConfig}},[a("div",{staticStyle:{width:"400px"},attrs:{slot:"content"},slot:"content"},[a("Form",{attrs:{"label-width":100}},[a("FormItem",{attrs:{label:"Main Path"}},[a("Input",{model:{value:e.formData.mainPath,callback:function(t){e.$set(e.formData,"mainPath",t)},expression:"formData.mainPath"}})],1),e._v(" "),a("FormItem",{attrs:{label:"Package Name"}},[a("Input",{model:{value:e.formData.pkg,callback:function(t){e.$set(e.formData,"pkg",t)},expression:"formData.pkg"}})],1),e._v(" "),a("FormItem",{attrs:{label:"Author Name"}},[a("Input",{model:{value:e.formData.author,callback:function(t){e.$set(e.formData,"author",t)},expression:"formData.author"}})],1),e._v(" "),a("FormItem",{attrs:{label:"Email Address"}},[a("Input",{model:{value:e.formData.email,callback:function(t){e.$set(e.formData,"email",t)},expression:"formData.email"}})],1),e._v(" "),a("FormItem",{attrs:{label:"Module Name"}},[a("Checkbox",{model:{value:e.formData.autoSettingModuleName,callback:function(t){e.$set(e.formData,"autoSettingModuleName",t)},expression:"formData.autoSettingModuleName"}},[e._v("Auto Setting")]),e._v(" "),e.formData.autoSettingModuleName?e._e():a("Input",{staticStyle:{width:"204px"},model:{value:e.formData.moduleName,callback:function(t){e.$set(e.formData,"moduleName",t)},expression:"formData.moduleName"}})],1),e._v(" "),a("FormItem",{attrs:{label:"Remove Prefix"}},[a("i-switch",{attrs:{size:"large"},model:{value:e.formData.isRemovePrefix,callback:function(t){e.$set(e.formData,"isRemovePrefix",t)},expression:"formData.isRemovePrefix"}})],1)],1)],1),e._v(" "),a("Button",{staticStyle:{float:"right",color:"blue"},attrs:{type:"text",icon:"ios-cog-outline"}},[e._v("Config")])],1)],1)],1)],1),e._v(" "),a("Table",{attrs:{columns:e.columns,height:"600",data:e.tbs,"no-data-text":"Can Not Find Table"},scopedSlots:e._u([{key:"operator",fn:function(t){var n=t.row;return[a("Button",{attrs:{shape:"circle",icon:"ios-more"},on:{click:function(t){return e.openConfig(n)}}},[e._v("Add")])]}}])}),e._v(" "),a("Page",{staticStyle:{float:"right"},attrs:{total:e.total,"page-size":e.limit,"show-sizer":""},on:{"on-change":e.pageHandle,"on-page-size-change":e.pageSizeHandle}}),e._v(" "),a("more-config",{ref:"configModal"})],1)},staticRenderFns:[]},S=a("C7Lr")(g,C,!1,null,null,null).exports;n.default.use(i.a);var x=new i.a({routes:[{path:"/",name:"Main",component:r,children:[{path:"/generator",name:"Generator",component:S}]}]});x.beforeEach(function(e,t,a){"/"===e.path?x.push({name:"Generator"}):a()});var T=x,w=a("gtAq"),N=a.n(w),k=(a("92qf"),a("84iU")),y=a.n(k);n.default.use(h.a);var _=new h.a.Store({state:{tables:[]},mutations:{updateTables:function(e,t){e.tables=t}}}),D=a("70WB"),j=a("FmMA");D.a({dsn:"https://27f15b2deab14c178e36ff69e44174e1@sentry.io/1515101",integrations:[new j.a({Vue:n.default,attachProps:!0})]}),n.default.prototype.$ajax=y.a,n.default.use(N.a),n.default.config.productionTip=!1,new n.default({el:"#app",router:T,store:_,components:{App:o},template:"<App/>"}),Array.prototype.remove=function(e){var t=this.indexOf(e);t>-1&&this.splice(t,1)}},f7oQ:function(e,t){},uPO5:function(e,t){}},["NHnr"]);
//# sourceMappingURL=app.9a4907a51b3d42915098.js.map