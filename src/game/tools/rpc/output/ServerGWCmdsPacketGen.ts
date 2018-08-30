class ServerGWCmdsPackGen{
    public static HeartBeatPacket(seqID:number, player:PlayerWrap):Net.Simple.Packet{
        var pack = new Net.Simple.Packet();
		pack.PutShort(seqID);
		pack.PutShort(0);
        player.Encode(pack);
        return pack;
    }
	public static LoginGuestPacket(seqID:number, devID:string, deviceType:string, partnerID:string, version:string):Net.Simple.Packet{
        var pack = new Net.Simple.Packet();
		pack.PutShort(seqID);
		pack.PutShort(10001);
        pack.PutString(devID);pack.PutString(deviceType);pack.PutString(partnerID);pack.PutString(version);
        return pack;
    }
	public static LoginPlatformPacket(seqID:number, ptID:string, account:string, deviceType:string, partnerID:string, version:string, reconnect:boolean, token:string, extension:string):Net.Simple.Packet{
        var pack = new Net.Simple.Packet();
		pack.PutShort(seqID);
		pack.PutShort(10002);
        pack.PutString(ptID);pack.PutString(account);pack.PutString(deviceType);pack.PutString(partnerID);pack.PutString(version);pack.PutBool(reconnect);pack.PutString(token);pack.PutString(extension);
        return pack;
    }
	
}