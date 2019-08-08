-- 增加启动时间大小 目前416天就插入失败
ALTER TABLE `cmdb_server_running`
CHANGE COLUMN `boot_time` `boot_time` BIGINT NOT NULL DEFAULT '0' COMMENT '系统重启的时间' ;