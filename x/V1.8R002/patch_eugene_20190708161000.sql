-- application add column

ALTER TABLE `cmdb_server_monitor_application` 
ADD COLUMN `server_id` INT(11) DEFAULT 0 NOT NULL COMMENT '服务器ID' AFTER `attrs`, 
ADD COLUMN `config_path` VARCHAR(256) DEFAULT '' NOT NULL COMMENT '配置文件路径' AFTER `server_id`; 

-- 

INSERT INTO cmdb_category(`id`,`name_en`,`name`,`type`,`order`,`custom`) VALUES(110,'haproxy','Haproxy',6,1,0);

INSERT INTO `monitor_item_template` VALUES (9968, '默认监控模板', '应用', 'haproxy', '多平台', 0, '', '系统默认监控模板不允许删除', 0, '2019-07-05 17:34:07', '2019-07-05 16:49:48', 0, '');

-- monitor_item
INSERT INTO `monitor_item` VALUES (99701, '连接状态', 'app_connect', '', '', 'number', 'haproxy', '', 0, '应用', '', NULL, '', 0, 0, '', 0, '2018-12-26 19:25:50', '', '', '0000-00-00 00:00:00', 0, 0, 0);


-- 99701~99800 monitor_item_template_item
insert into `monitor_item_template_item` (`id`, `template_id`, `item_id`, `monitor_cycle`, `judge_value`, `warning_level_id`, `store`, `judge_opertate`, `warning_switch`, `unit`, `express`, `create_time`) 
                                    values('99701','9968','99701','60','','19','1','','0','','','2019-05-14 10:40:14');
