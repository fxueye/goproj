class ItemDataListWrap extends Net.Simple.IWrapper {
	
	public ItemList:Array<ItemDataWrap>; // 道具列表
	public constructor() {
		super();
	}
	Decode(pck:Net.Simple.Packet):ItemDataListWrap{
		
		this.ItemList = new Array<ItemDataWrap>();
		for (var i = 0,len = this.ItemList.length; i < len; i++)
		{
			this.ItemList[i] = new ItemDataWrap();
			this.ItemList[i].Decode(pck);
			
		}
		
		return this;
	}
    Encode(pck:Net.Simple.Packet){
		
    	if (this.ItemList == null) pck.PutShort(0); 
    	else
    	{
        	pck.PutShort(this.ItemList.length);
        	for(var i = 0,len = this.ItemList.length; i < len; i++)
        	{
        		this.ItemList[i].Encode(pck);
				
        	}
        }
    	
	}

}