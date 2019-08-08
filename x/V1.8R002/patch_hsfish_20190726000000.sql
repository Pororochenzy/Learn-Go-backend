
-- 同步告警等级为空的问题
UPDATE  `warning_event` SET `warning_level`='严重' where `warning_level`='';
UPDATE `monitor_item_template_item` SET `warning_level_id`=19  WHERE `warning_level_id` NOT IN (SELECT `id` FROM `warning_level`);
UPDATE `monitor_item_obj_relation` SET `warning_level_id`=19  WHERE `warning_level_id` NOT IN (SELECT `id` FROM `warning_level`);
