-- 按钮父菜单ID
set @parentId = (select REPLACE(UUID(),'-',''));

-- 菜单SQL

INSERT INTO `sys_menu`(`MENU_ID`, `PARENT_ID`, `NAME`, `URL`, `PERMS`, `TYPE`, `ICON`, `ORDER_NUM`, `CREATE_BY`, `CREATE_TIME`, `UPDATE_BY`, `UPDATE_TIME`, `DELETED`)
VALUES (@parentId,'1', '{{ table.TableComment }}', '{{moduleName}}/{{fileName}}', NULL, '1', 'config', '6','1',CURRENT_TIMESTAMP,'1',CURRENT_TIMESTAMP,0);

-- 菜单对应按钮SQL
INSERT INTO `sys_menu`(`MENU_ID`, `PARENT_ID`, `NAME`, `URL`, `PERMS`, `TYPE`, `ICON`, `ORDER_NUM`, `CREATE_BY`, `CREATE_TIME`, `UPDATE_BY`, `UPDATE_TIME`, `DELETED`)
VALUES ((select REPLACE(UUID(),'-','')),@parentId, '查看', null, '{{moduleName}}:{{fileName}}:list,{{moduleName}}:{{fileName}}:info', '2', null, '6','1',CURRENT_TIMESTAMP,'1',CURRENT_TIMESTAMP,0);
INSERT INTO `sys_menu`(`MENU_ID`, `PARENT_ID`, `NAME`, `URL`, `PERMS`, `TYPE`, `ICON`, `ORDER_NUM`, `CREATE_BY`, `CREATE_TIME`, `UPDATE_BY`, `UPDATE_TIME`, `DELETED`)
VALUES ((select REPLACE(UUID(),'-','')),@parentId, '新增', null, '{{moduleName}}:{{fileName}}:save', '2', null, '6','1',CURRENT_TIMESTAMP,'1',CURRENT_TIMESTAMP,0);
INSERT INTO `sys_menu`(`MENU_ID`, `PARENT_ID`, `NAME`, `URL`, `PERMS`, `TYPE`, `ICON`, `ORDER_NUM`, `CREATE_BY`, `CREATE_TIME`, `UPDATE_BY`, `UPDATE_TIME`, `DELETED`)
VALUES ((select REPLACE(UUID(),'-','')),@parentId, '修改', null, '{{moduleName}}:{{fileName}}:update', '2', null, '6','1',CURRENT_TIMESTAMP,'1',CURRENT_TIMESTAMP,0);
INSERT INTO `sys_menu`(`MENU_ID`, `PARENT_ID`, `NAME`, `URL`, `PERMS`, `TYPE`, `ICON`, `ORDER_NUM`, `CREATE_BY`, `CREATE_TIME`, `UPDATE_BY`, `UPDATE_TIME`, `DELETED`)
VALUES ((select REPLACE(UUID(),'-','')),@parentId, '删除', null, '{{moduleName}}:{{fileName}}:delete', '2', null, '6','1',CURRENT_TIMESTAMP,'1',CURRENT_TIMESTAMP,0);



