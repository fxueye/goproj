class PlayerWrap extends Net.Simple.IWrapper {
	
	public UID:number; // 游戏的uid
	public GUID:string; // 平台唯一ID,游客登录的话为设备ID
	public PID:string; // 登录平台ID
	public UserName:string; // 用户名（唯一，初始和UID一致）
	public LoginTime:number; // 登录时间戳（秒）
	public OnlineTime:number; // 总在线时长（秒）
	public CreateTime:number; // 创角时间戳（秒）
	public ServerTime:number; // 服务器时间（秒）
	public OnBattleIdx:number; // 上阵索引
	public Items:Array<ItemDataWrap>; // 道具列表
	public Icon:number; // 头像
	public TutorialMask:number; // 新手引导
	public constructor() {
		super();
	}
	Decode(pck:Net.Simple.Packet):PlayerWrap{
		
		this.UID = pck.GetLong(); 
		this.GUID = pck.GetString(); 
		this.PID = pck.GetString(); 
		this.UserName = pck.GetString(); 
		this.LoginTime = pck.GetLong(); 
		this.OnlineTime = pck.GetLong(); 
		this.CreateTime = pck.GetLong(); 
		this.ServerTime = pck.GetLong(); 
		this.OnBattleIdx = pck.GetInt(); 
		this.Items = new Array<ItemDataWrap>();
		for (var i = 0,len = this.Items.length; i < len; i++)
		{
			this.Items[i] = new ItemDataWrap();
			this.Items[i].Decode(pck);
			
		}
		
		this.Icon = pck.GetInt(); 
		this.TutorialMask = pck.GetLong(); 
		return this;
	}
    Encode(pck:Net.Simple.Packet){
		
    	pck.PutLong(this.UID); 
    	pck.PutString(this.GUID); 
    	pck.PutString(this.PID); 
    	pck.PutString(this.UserName); 
    	pck.PutLong(this.LoginTime); 
    	pck.PutLong(this.OnlineTime); 
    	pck.PutLong(this.CreateTime); 
    	pck.PutLong(this.ServerTime); 
    	pck.PutInt(this.OnBattleIdx); 
    	if (this.Items == null) pck.PutShort(0); 
    	else
    	{
        	pck.PutShort(this.Items.length);
        	for(var i = 0,len = this.Items.length; i < len; i++)
        	{
        		this.Items[i].Encode(pck);
				
        	}
        }
    	
    	pck.PutInt(this.Icon); 
    	pck.PutLong(this.TutorialMask); 
	}

}