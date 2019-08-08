/* 12:04:41 PM 开发环境 192.168.1.253 gsn_mcenter */ ALTER TABLE `auth_role_menu` CHANGE `update_at` `update_at` TIMESTAMP  NOT NULL  DEFAULT CURRENT_TIMESTAMP  ON UPDATE CURRENT_TIMESTAMP  COMMENT '更新时间';

delete from auth_role_menu where role_id=1;
INSERT INTO auth_role_menu(role_id, menu_id, data_flag, data_range) SELECT 1, id, data_flag, "all" FROM auth_menu;