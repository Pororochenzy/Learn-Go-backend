```golang
func (p *PublishBase) WebsocketHandler(uid int64, ws *websocket.Conn) revel.Result {
	p.HandlerRecover()

	global.Logger.Debug("uid:%v", uid)

	//查看用户是否存在
	sql := fmt.Sprintf(`SELECT * FROM user WHERE id=%d`, uid)
	userInfo, err := mysql.QueryOne(global.DB, sql)
	if err != nil {
		global.Logger.Error(err.Error())
		p.SetOutSystemErr()
		ws.Close()
		return p.RenderJson(p.Out)
	}
	if userInfo == nil {
		p.Out.Ret = define.RET_DATA_NOT_EXIST_ERR
		p.Out.Msg = fmt.Sprintf("用户[id:%v]不存在", uid)
		ws.Close()
		return p.RenderJson(p.Out)
	}

	//新建ws连接
	uuid := pbls.GenerateConnUuid()
	websocketConn := pbls.NewPubConn(uuid, uid, ws)

	// 读
	for {
		data, err := websocketConn.WaitRead()
		if err != nil {
			global.Logger.Error(err.Error())
			if err == pbls.ErrReqDataError {
				p.SetOutValueErr()
				break
			} else {
				p.SetOutSystemErr()
				break
			}
		}
		if data == "@heart" {
			// 心跳包，不做处理
			//global.Logger.Debug("received websocket heart beat")
			ws.Write([]byte("@heart"))
			continue
		}

		global.Logger.Debug("req: %v", data)

		// 协程处理数据
		go websocketConn.HandleReqFromWeb(data)
	}

	//关闭连接
	websocketConn.Close()

	return p.RenderJson(p.Out)
}
```