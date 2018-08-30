class ServerCSCmdsPackGen{
    public static GW2CS_PingPacket(seqID:number):Net.Simple.Packet{
        var pack = new Net.Simple.Packet();
		pack.PutShort(seqID);
		pack.PutShort(22001);
        
        return pack;
    }
	public static GW2CS_LoginGuestPacket(seqID:number, deviceID:string, deviceType:string, partnerID:string, ip:string):Net.Simple.Packet{
        var pack = new Net.Simple.Packet();
		pack.PutShort(seqID);
		pack.PutShort(22004);
        pack.PutString(deviceID);pack.PutString(deviceType);pack.PutString(partnerID);pack.PutString(ip);
        return pack;
    }
	
}