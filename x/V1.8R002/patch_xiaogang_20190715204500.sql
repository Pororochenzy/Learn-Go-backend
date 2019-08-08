
-- 修改隔离搬运content字段类型
ALTER TABLE `ccfacer_c2m`
CHANGE COLUMN `content` `content` MEDIUMTEXT NOT NULL ;

ALTER TABLE `ccfacer_m2c`
CHANGE COLUMN `content` `content` MEDIUMTEXT NOT NULL ;

