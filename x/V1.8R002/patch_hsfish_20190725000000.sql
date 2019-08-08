

ALTER TABLE `monitor_network_custom` ADD COLUMN `is_deleted` TINYINT(1) NOT NULL DEFAULT 0;

UPDATE `monitor_item` SET `value_type`='number', `dependence`='{"type": "system", "name": "new_admin_tag.admin_target_address"}' WHERE `id`=99991;


INSERT INTO `monitor_item`(`id`,`name`, `name_en`, `dependence`, `value_type`, `app_category`, `category`) 
VALUES 
('99966', '整机吞吐', 'throughput', '', 'number', 'network', '网络设备');