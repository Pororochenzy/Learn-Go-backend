
INSERT INTO `sys_conf`(`id`, `name`, `remark`) VALUES(21, '公司名称', '特殊处理时指定');

UPDATE `auth_menu` SET `name`='第三方线路' WHERE id=38;
UPDATE `auth_menu` SET `name`='第三方线路列表' WHERE id=239;
INSERT INTO `device_type`(`id`, `name`, `flag`) VALUES('996', '刀片服务器', '0');
UPDATE `cmdb_server` SET `device_type_id`=996 WHERE `is_deleted`=0 AND `device_type_id`=1 AND `physical_server_id`>0;