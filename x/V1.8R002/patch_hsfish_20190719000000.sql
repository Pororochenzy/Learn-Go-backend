
INSERT INTO `monitor_item`(`id`,`name`, `name_en`, `dependence`, `value_type`, `app_category`, `category`) 
VALUES 
('99971', '1分钟新建立会话数', 'current_session_one_minute', '', 'number', 'network', '网络设备'),
('99970', '5分钟新建立会话数', 'current_session_five_minute', '', 'number', 'network', '网络设备'),
('99969', '路由总数', 'route_total', '', 'number', 'network', '网络设备');

UPDATE `monitor_item` SET `unit`='mA' WHERE `id`=99977;
UPDATE `monitor_item` SET `unit`='V' WHERE `id`=99978;
UPDATE `monitor_item` SET `unit`='dBm' WHERE `id`=99982;
UPDATE `monitor_item` SET `unit`='dBm' WHERE `id`=99981;
UPDATE `monitor_item` SET `unit`='°C' WHERE `id`=99979;

INSERT INTO `monitor_item`(`id`,`name`, `name_en`, `dependence`, `value_type`, `app_category`, `category`) 
VALUES 
('99968', '每秒连接数', 'connectionsPerSec', '', 'number', 'network', '网络设备'),
('99967', '每秒请求数', 'requestsPerSec', '', 'number', 'network', '网络设备');
