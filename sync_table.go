package own

import (
	"fmt"
	"timer/global"

	"geesunn.com/define"
	"geesunn.com/lib/mysql"
)

func Init() {
	DeleteRepleteData() //删除所有表重复数据
	AllTableSync()      //同步所有表数据
}
func AllTableSync() {
	ServerTable()
	NetworkTable()
	LoadBalanceTable()
	StorageTable()
	CameraTable()
	CameraTeminalTable()
	BladeBoxTable()
	PhysicalServerTable()
	BladeModuleTable()
	SmartRackHostTable()
	VirtualPlatformTable()
	NetworkSegmenTable()
	IdcTable()
	IdcRackTable()
	IdcAreaTable()
	ApplicationTable()
}

//服务器
func ServerTable() {
	CommomSync(define.OBJECT_TYPE_SERVER)
}

//网络设备
func NetworkTable() {
	CommomSync(define.OBJECT_TYPE_NETWORK)
}

//负载均衡
func LoadBalanceTable() {
	CommomSync(define.OBJECT_TYPE_F5)
}

//存储设备
func StorageTable() {
	CommomSync(define.OBJECT_TYPE_STORAGE)
}

//摄像头
func CameraTable() {
	CommomSync(define.OBJECT_TYPE_CAMERA)
}

//终端机
func CameraTeminalTable() {
	CommomSync(define.OBJECT_TYPE_CAMERA_TERMINAL)
}

//刀片机机箱
func BladeBoxTable() {
	CommomSync(define.OBJECT_TYPE_BLADE_BOX)
}

//刀片机刀片
func PhysicalServerTable() {
	CommomSync(define.OBJECT_TYPE_PHYSICAL_SERVER)
}

//刀片机模块
func BladeModuleTable() {
	CommomSync(define.OBJECT_TYPE_BLADE_MODULE)
}

// 智能机柜主机
func SmartRackHostTable() {
	CommomSync(define.OBJECT_TYPE_UBIT)
}

// vecenter
func VirtualPlatformTable() {
	CommomSync(define.OBJECT_TYPE_VCENTER)
}

//IP网段
func NetworkSegmenTable() {
	CommomSync(define.OBJECT_TYPE_IP_SEGMENT)
}

//机房
func IdcTable() {
	CommomSync(define.OBJECT_TYPE_IDC)
}

//机房机柜
func IdcRackTable() {
	CommomSync(define.OBJECT_TYPE_IDC_RACK)
}

//机房区域
func IdcAreaTable() {
	CommomSync(define.OBJECT_TYPE_IDC_AREA)
}
func ApplicationTable() {
	appList := []string{"redis", "haproxy", "mysql", "mongodb", "iis", "nginx", "apache", "oracle", "sybase", "sqlserver", "weblogic", "tomcat"}
	for _, app := range appList {
		CommomSync(app)
	}
}
func SqlRules(objectTypeDefine string) string {

	tableName := define.OBJECT_TYPE_TABLE[objectTypeDefine]
	switch objectTypeDefine {
	//idc idc_rack idc_area 和智能机柜
	case define.OBJECT_TYPE_IDC, define.OBJECT_TYPE_IDC_RACK, define.OBJECT_TYPE_IDC_AREA, define.OBJECT_TYPE_UBIT:
		sql := ""
		if objectTypeDefine != define.OBJECT_TYPE_IDC_AREA {
			sql = fmt.Sprintf(`SELECT id, asset_tag from %v `, tableName)
		} else {
			sql = fmt.Sprintf(`SELECT id, name AS asset_tag from %v `, tableName)
		}
		return sql
		//应用类
	case "redis", "haproxy", "mysql", "mongodb", "memcache", "iis", "nginx", "apache", "oracle", "sybase", "sqlserver", "weblogic", "oracle_rac", "tomcat":

		sql := fmt.Sprintf(`select 
		csma.id ,csma.aliasname as asset_tag, al.app_type as object_type, csma.monitor_switch 
		from
		cmdb_server_monitor_application as csma 
		inner JOIN 
	application_list as al on csma.application_id=al.id 
		WHERE al.app_type= "%v"
`, objectTypeDefine)

		return sql
		//大表常见类
	case define.OBJECT_TYPE_SERVER, define.OBJECT_TYPE_NETWORK, define.OBJECT_TYPE_F5, define.OBJECT_TYPE_STORAGE, define.OBJECT_TYPE_CAMERA, define.OBJECT_TYPE_CAMERA_TERMINAL,
		define.OBJECT_TYPE_BLADE_BOX, define.OBJECT_TYPE_PHYSICAL_SERVER, define.OBJECT_TYPE_BLADE_MODULE:

		sql := fmt.Sprintf(`SELECT id, asset_tag,device_status_id,is_deleted from %v 
		WHERE is_deleted=0 `, tableName)

		return sql
		//ip网段类
	case define.OBJECT_TYPE_IP_SEGMENT:
		sql := fmt.Sprintf(`SELECT id, network_segment as asset_tag from %v  `, tableName)
		return sql

		//vecenter
	case define.OBJECT_TYPE_VCENTER:
		sql := fmt.Sprintf(`select id , name as asset_tag, monitor_switch from cmdb_virtual_platform  WHERE object_type ="%v" `, objectTypeDefine)
		return sql

	default:
		return ""
	}
}

//. 删除基本信息表多个重复数据
func DeleteRepleteData() {
	repeatSQL := `select count(*),id,object_type,object_id from object_base_info group by object_type,object_id having count(*)>1`
	groupInfos, err := mysql.Query(global.DB, repeatSQL)
	if err != nil {
		global.Logger.Error(err.Error())
		return
	}
	for _, info := range groupInfos {
		limitNum := info["count(*)"].(int64) - 1
		sql := fmt.Sprintf(`delete from object_base_info where id in (select temp.id from (
		select id from  object_base_info where  object_id=%v and object_type ="%v"
		 order by id desc limit %d) as temp )`, info["object_id"].(int64), info["object_type"].(string), limitNum)

		num, err := mysql.Delete(global.DB, sql)
		if err != nil {
			global.Logger.Error(err.Error())
			return
		}
		if num > 0 {
			global.Logger.Info("删除成功")
		}
	}
}
func CommomSync(objectTypeDefine string) {
	sql := SqlRules(objectTypeDefine)
	if sql == "" {
		global.Logger.Error("sql语句为空")
		return
	}
	infos, err := mysql.Query(global.DB, sql)
	if err != nil {
		global.Logger.Error(err.Error())
		return
	}
	mainTable := make(map[int64]map[string]interface{})
	for _, info := range infos {
		mainTable[info["id"].(int64)] = info //key 为 主表的id
	}

	objectBaseInfos, err := mysql.GetTRInfo(global.DB, "object_base_info", map[string]interface{}{
		"object_type": objectTypeDefine,
	})

	if err != nil {
		global.Logger.Error(err.Error())
		return
	}
	for _, obi := range objectBaseInfos {
		//1.如果基本信息表有值, 主表没值
		if _, ok := mainTable[obi["object_id"].(int64)]; !ok {
			//删掉基本信息表
			num, err := mysql.DeleteTR(global.DB, "object_base_info", map[string]interface{}{
				"id": obi["id"],
			})
			if err != nil {
				global.Logger.Error(err.Error())
				return
			}
			if num > 0 {
				global.Logger.Info("删除成功")
			}
		} else {
			mainTableMap := mainTable[obi["object_id"].(int64)]
			flag := false //判断是否需要更新
			//2.判断更新
			if obi["asset_tag"].(string) != mainTableMap["asset_tag"].(string) {
				obi["asset_tag"] = mainTableMap["asset_tag"].(string)
				flag = true
			}

			if v, ok := mainTableMap["device_status_id"].(int64); ok {
				var monitor_switch string
				switch v {
				case 1, 2, 3:
					monitor_switch = "1"

				case 4:
					monitor_switch = "0"

				default:
					global.Logger.Error("设备状态值有误")
					return
				}
				obi["monitor_switch"] = monitor_switch
				flag = true
			}
			if v, ok := mainTableMap["monitor_switch"]; ok {
				if v != obi["monitor_switch"] {
					v = obi["monitor_switch"]
					flag = true
				}
			}

			if flag {
				mysql.UpdateTR(global.DB, "object_base_info", map[string]interface{}{
					"object_type": objectTypeDefine,
					"object_id":   obi["object_id"],
				}, obi)
				if err != nil {
					global.Logger.Error(err.Error())
					return
				}
			}
		}
	}
	//4.基本信息表没值,主表有值
	for _, main := range mainTable {
		flag := false
		for _, obi := range objectBaseInfos {
			if obi["object_id"].(int64) == main["id"].(int64) { //用主表的id 去对比基本信息表的每一行的object_id

				flag = true
				break
			}
		}
		//基本信息表新增 数据
		if !flag {
			num, err := mysql.InsertTR(global.DB, "object_base_info", map[string]interface{}{
				"object_id":      main["id"].(int64),
				"object_type":    objectTypeDefine,
				"asset_tag":      main["asset_tag"],
				"monitor_switch": 1,
			})
			if err != nil {
				global.Logger.Error(err.Error())
				return
			}
			if num > 0 {
				global.Logger.Info("新增成功")
			}
			flag = false
		}
	}

}
