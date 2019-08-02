## object_base_info表

| id | asset_tag | object_type | object_id | monitor_switch |
| ------ | ------ | ------ |------ |------ |


## 服务器表 cmdb_server
asset_tag ,monitor_switch ,is_deleted, device_status_id

## 网络设备表  cmdb_network
asset_tag , is_deleted ,device_status_id

## 负载均衡表  cmdb_load_balance
asset_tag ,is_deleted , device_status_id

## 存储设备表 cmdb_storage
asset_tag , is_deleted 

## 刀片机刀箱 cmdb_blade_box 
asset_tag , is_deleted , monitor_switch

## 刀片机刀片 cmdb_physical_server
asset_tag , is_deleted, monitor_switch , 

## 刀片机模块 blade_module
asset_tag , is_deleted, monitor_switch , 

## 自定义对象 cmdb_category_custom_instance
asset_tag , is_deleted

## 应用、数据库、中间件  application_list ?   cmdb_server_monitor_application

## ip表  cmdb_ip  

## 虚拟化平台 cmdb_virtual_platform
name, monitor_switch , object_type->能取到vcenter 

## cmdb_database_backup 对应系统配置-数据备份
## device_status 是记录(1待运营,2运营中,3故障中,4已退役)

## ip网段表  cmdb_ip_network_segment   

### 业务监控表   cmdb_business
name,

### 采集器表   cmdb_terminal 
object_type 为f5的  代表到 cmdb_load_balance表里找

object_type 为haproxy 它到 cmdb_server_monitor_application 去找了

cmdb_server_monitor_application  的类型  它关联了  application_list(), 


f5 属于硬件 ,所以是在cmdb_load_balance表记录着的
harpoxy,weblogic 这种软件层都 在cmdb_server_monitor_application 记录着 


object_info 表里存的object_type 关联的是 application_list表里的app_type 

中间件监控,应用监控, 数据库监控  ,都同属于cmdb_server_monitor_application 表


摄像头监控板块: object_type=  camera(摄像头监控), camera_terminal(终端机监控 )   ,(摄像头地图 )

刀片机监控板块 : object_type= blade_box(刀片机刀箱) , blade_module(刀片机模块), physical_server(刀片机刀片)

机房监控板块 :  object_type= idc(机房) , idc_area(机房区域) , idc_rack(机房的机柜) 

 网络监控板块 : object_type= network , network_segment    (两表有啥区别)

  采集器             object_type= terminal 
  
                 object_type = storage 
                 object_type = vcenter

###  额外
cmdb_attr_ext->自定义模型,里面有个字段是attr_type_id ,用内置属性此id=0,数值是1,文本是2, 是关连   cmdb_attr_type (数值,文本,日期 ), cmdb_attr_type_value(类型的值，例如 float、string、datetime、"a") -> 表里有 attr_type_id 