DROP TABLE IF EXISTS `cmdb_device_cluster_relation`;
DROP TABLE IF EXISTS `cmdb_object_cluster`;

CREATE TABLE `cmdb_device_cluster_relation` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `object_id` int(11) NOT NULL COMMENT '负载均衡表id',
  `cluster_id` int(11) NOT NULL COMMENT '负载均衡集群表id',
  `create_user_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `index_cluster_id` (`cluster_id`) USING BTREE
) ENGINE=InnoDB  DEFAULT CHARSET=utf8;


CREATE TABLE `cmdb_object_cluster` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `remark` varchar(256) NOT NULL,
  `vip` varchar(64) NOT NULL COMMENT '虚拟ip',
  `create_user_id` int(11) NOT NULL,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `object_type` varchar(32) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8;


