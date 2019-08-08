INSERT INTO `cmdb_category`(`id`, `name_en`, `name`, `type`, `order`, `is_show`) values('14', 'auto_find', '自动发现', '1', '21', '0');

ALTER TABLE `auto_find` ADD COLUMN `inform_switch` TINYINT(1)  NOT NULL DEFAULT '0' COMMENT '通知开关：0.否 1.是';

ALTER TABLE `user` ADD COLUMN `work_no` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '工号';