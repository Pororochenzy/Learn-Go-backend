
-- ipsla状态信息
ALTER TABLE `cmdb_network_ipsla` ADD COLUMN `last_rtt_oper_sense_name` VARCHAR(32) NOT NULL DEFAULT '' COMMENT '状态描述' after`last_rtt_oper_sense`;

-- 对象特殊采集信息
DROP TABLE IF EXISTS `monitor_network_custom`;
CREATE TABLE `monitor_network_custom` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `object_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '对象id',
  `type` varchar(32) NOT NULL DEFAULT '' COMMENT '类型 vs rs vpn',
  `snmp_index` varchar(64) NOT NULL DEFAULT '' COMMENT 'snmp索引',
  `content` json COMMENT '内容',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY`index_object_id_type` (`object_id`, `type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 网络设备端口新增 工作模式
ALTER TABLE `cmdb_network_port` ADD COLUMN `operate_mode` TINYINT(1) UNSIGNED NOT NULL DEFAULT '1' COMMENT '工作模式 1.unknown 2.half-duplex 3.full-duplex';

-- 新增网络设备监控指标
INSERT INTO `monitor_item`(`id`,`name`, `name_en`, `dependence`, `value_type`, `app_category`, `category`) 
VALUES 
('99991', '第三方线路状态', 'ipsla_status', '{"type": "system", "name":"admin_owner"}', 'text', 'network', '网络设备'),
('99990', 'RS状态', 'real_status', '{"type": "system", "name": "real_server_id"}', 'text', 'network', '网络设备'),
('99989', 'RS连接数', 'real_conn_cnt', '{"type": "system", "name": "real_server_id"}', 'number',  'network', '网络设备'),
('99988', 'VS连接数', 'vs_conn_cnt', '{"type": "system", "name": "virt_server_id"}', 'number',  'network', '网络设备'),
('99987', 'VPN Open数', 'vpn_tunnels_open', '{"type": "system", "name": "vpn_id"}', 'number',  'network', '网络设备'),
('99986', 'VPN Establish数', 'vpn_tunnels_est', '{"type": "system", "name": "vpn_id"}', 'number',  'network', '网络设备'),
('99985', 'VPN Rejected数', 'vpn_tunnels_rejected', '{"type": "system", "name": "vpn_id"}', 'number',  'network', '网络设备'),
('99984', 'VPN 入流量', 'vpn_bytes_in', '{"type": "system", "name": "vpn_id"}', 'number',  'network', '网络设备'),
('99983', 'VPN 出流量', 'vpn_bytes_out', '{"type": "system", "name": "vpn_id"}', 'number',  'network', '网络设备'),
('99982', 'SFP Tx Power', 'sfp_tx_power', '{"type": "system", "name": "sfp_name"}', 'number',  'network', '网络设备'),
('99981', 'SFP Rx Power', 'sfp_rx_power', '{"type": "system", "name": "sfp_name"}', 'number',  'network', '网络设备'),
('99979', 'SFP Temperature', 'sft_temperature', '{"type": "system", "name": "sfp_name"}', 'number',  'network', '网络设备'),
('99978', 'SFP Voltage', 'sft_voltage', '{"type": "system", "name": "sfp_name"}', 'number',  'network', '网络设备'),
-- ('99976', 'Temperature', 'temperature_value', '{"type": "system", "name": "temperature_name"}', 'number',  'network', '网络设备'),
('99977', 'SFP Current', 'sft_current', '{"type": "system", "name": "sfp_name"}', 'number',  'network', '网络设备');



INSERT INTO `monitor_item`(`id`,`name`, `name_en`, `dependence`, `value_type`, `app_category`, `category`) 
VALUES 
('99975', '当前会话数', 'sys_current_session', '', 'number', 'network', '网络设备'),
('99974', '会话使用率', 'sys_session_percent', '', 'number', 'network', '网络设备'),
('99973', 'AP连接数', 'ap_connect_number', '', 'number', 'network', '网络设备'),
('99972', '无线用户在线数', 'online_user_number', '', 'number', 'network', '网络设备');
