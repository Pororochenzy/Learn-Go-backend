-- 
DROP TABLE IF EXISTS `cmdb_object_group_relation`;

CREATE TABLE `cmdb_object_group_relation` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `group_id` int(11) NOT NULL DEFAULT '0' COMMENT '小组ID',
  `object_id` int(11) NOT NULL DEFAULT '0' COMMENT '对象ID',
  `object_type` varchar(32) NOT NULL DEFAULT '' COMMENT '对象类型',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `index_union_oid_otype` (`object_id`,`object_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;