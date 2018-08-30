class ItemDataWrap extends Net.Simple.IWrapper {
	
	public ItemID:number; // 道具ID
	public Count:number; // 道具数量
	public LastUpdateTime:Long; // 最后一次更新时间
	public constructor() {
		super();
	}
	Decode(pck:Net.Simple.Packet):ItemDataWrap{
		
		this.ItemID = pck.GetInt(); 
		this.Count = pck.GetInt(); 
		this.LastUpdateTime = pck.GetLong(); 
		return this;
	}
    Encode(pck:Net.Simple.Packet){
		
    	pck.PutInt(this.ItemID); 
    	pck.PutInt(this.Count); 
    	pck.PutLong(this.LastUpdateTime); 
	}

}