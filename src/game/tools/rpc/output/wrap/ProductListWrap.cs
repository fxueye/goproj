using System;
using Common.Net.Simple;

namespace Game 
{
	public class ProductListWrap : IWrapper
	{	
		public ProductInfoWrap[] ProductList; // 商品列表
		
		public void Decode(Packet pck)
		{	
			ProductList = new ProductInfoWrap[pck.GetShort()];
			for (int i = 0; i < ProductList.Length; i++)
			{
				ProductList[i] = new ProductInfoWrap();
				ProductList[i].Decode(pck);
				
			}
			
		}
		
        public void Encode(Packet pck)
        {	
        	if (ProductList == null) pck.PutShort((short)0); 
        	else
        	{
	        	pck.PutShort((short)ProductList.Length);
	        	for(int i = 0; i < ProductList.Length; i++)
	        	{
	        		ProductList[i].Encode(pck);
					
	        	}
	        }
        	
        }
	}
}