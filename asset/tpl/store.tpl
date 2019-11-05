import {list, deleteList, saveOrUpdate} from "@/api/common";


export default {
    namespaced: true,
    state: {},
    getters: {
        namespace: () => '{{ moduleName }}/{{ fileName }}'
    },
    actions: {
        list({getters}, {page,limit,{% for column in searchColumns %}{{ column.FieldName }},{% endfor %}}) {
            return new Promise((resolve, reject) => {
                list({page,limit,{% for column in searchColumns %} {{ column.FieldName }},{% endfor %}}, getters.namespace).then(({data}) => {
                    if (data && data.code === 200) {
                        let {list, totalCount} = data.page
                        resolve({list, totalCount})
                    } else {
                        reject(data.msg)
                    }
                })
            })
        },
        delete({getters}, ids) {
            return new Promise((resolve, reject) => {
                deleteList(
                    ids, getters.namespace
                ).then(({data}) => {
                    if (data && data.code === 200) {
                        resolve()
                    } else {
                        reject(data.msg)
                    }
                })
            })
        },
        saveOrUpdate({getters}, { {% for column in addColumns %}{{ column.FieldName }},{% endfor %}}) {
            return new Promise((resolve, reject) => {
                saveOrUpdate({ {% for column in addColumns %}{{ column.FieldName }},{% endfor %}}, getters.namespace).then(({data}) => {
                    if (data && data.code === 200) {
                        resolve()
                    } else {
                        reject(data.msg)
                    }
                })
            })
        }
    }
}
