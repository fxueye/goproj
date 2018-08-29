interface IClientCmds
{
	HeartBeat(cmd:Net.Simple.Command, msg:string ); // 心跳
	LoginSuccess(cmd:Net.Simple.Command, player:PlayerWrap , reconnect:boolean , extension:string ); // 登录成功
	LoginFailed(cmd:Net.Simple.Command, errorCode:number , errMsg:string ); // 登录失败（1 用户不存在，2  密码错误， 3 禁止登陆）
	
}
class ClientCmdsInvoker implements Net.Simple.IInvoker {
	private _cmds:IClientCmds = null;
	private _onCmdInvoked:Function = null;
	private _obj:any = null;
	public constructor(cmds:IClientCmds) {
		this._cmds = cmds;
	}
	public SetOnCmdInvoked(func:Function,obj:any){
		this._onCmdInvoked = func;
		this._obj = obj;
	}
	public Invoke(cmd:Net.Simple.Command):void{
		var pack = cmd.Pack;
		switch(cmd.Opcode){
			case 0: _cmds.HeartBeat(cmd, new pack.GetString()); break;
			case 1: _cmds.LoginSuccess(cmd, new PlayerWrap().Decode(pack), new pack.GetBool(), new pack.GetString()); break;
			case 2: _cmds.LoginFailed(cmd, new pack.GetShort(), new pack.GetString()); break;
			
		}
		if(this._onCmdInvoked != null && this._obj != null){
			this._onCmdInvoked.call(this._obj,cmd.Opcode);

		}
	}	

}