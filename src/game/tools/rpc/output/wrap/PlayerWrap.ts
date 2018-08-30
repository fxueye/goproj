class PlayerWrap extends Net.Simple.IWrapper {
	
	public GUID:string; // 唯一ID
	public CreateTime:Long; // 创建时间
	public constructor() {
		super();
	}
	Decode(pck:Net.Simple.Packet):PlayerWrap{
		
		this.GUID = pck.GetString(); 
		this.CreateTime = pck.GetLong(); 
		return this;
	}
    Encode(pck:Net.Simple.Packet){
		
    	pck.PutString(this.GUID); 
    	pck.PutLong(this.CreateTime); 
	}

}