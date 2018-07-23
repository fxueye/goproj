package main

import (
	"encoding/binary"
	"net"
	"time"

	log "github.com/cihub/seelog"
)

const (
	//	NTP_SERVER_IP = "cn.ntp.org.cn" /*NTP IP*/
	NTP_PORT_STR = "123" /*NTP专用端口号字 符串*/
	NTP_PCK_LEN  = 48
	LI           = 0
	VN           = 3
	MODE         = 3
	STRATUM      = 0
	POLL         = 4
	PREC         = -6
	JAN_1970     = 0x83aa7e80 /* 1900年～1970年之间的时间秒数 */

)

func NTPFRAC(x int64) int64 {
	return (4294*(x) + ((1981 * (x)) >> 11))
}

func USEC(x int64) int64 {
	return (((x) >> 12) - 759*((((x)>>10)+32768)>>16))
}

type ntp_time struct {
	coarse uint32
	fine   uint32
}

type ntp_packet struct {
	leap_ver_mode        byte
	startum              byte
	poll                 byte
	precision            byte
	root_delay           int
	root_dispersion      int
	reference_identifier int
	reference_timestamp  ntp_time
	originage_timestamp  ntp_time
	receive_timestamp    ntp_time
	transmit_timestamp   ntp_time
}

var protocol []byte

func construct_packet() ([]byte, int) {
	reqData := make([]byte, NTP_PCK_LEN)
	//设置16字节的包头
	head := (LI << 30) | (VN << 27) | (MODE << 24) | (STRATUM << 16) | (POLL << 8) | (PREC & 0xff)
	binary.BigEndian.PutUint32(reqData[0:4], uint32(head))
	//设置Root Delay、Root Dispersion和Reference Indentifier
	binary.BigEndian.PutUint32(reqData[4:8], uint32(1<<16))
	binary.BigEndian.PutUint32(reqData[8:12], uint32(1<<16))
	binary.BigEndian.PutUint32(reqData[12:16], uint32(1<<16))

	//设置Timestamp部分
	timeOri := JAN_1970 + time.Now().Unix()

	//设置Transmit Timestamp coarse
	binary.BigEndian.PutUint32(reqData[40:44], uint32(timeOri))
	//设置Transmit Timestamp fine
	binary.BigEndian.PutUint32(reqData[44:48], uint32(NTPFRAC(timeOri)))
	return reqData, NTP_PCK_LEN
}

func GetNtpTime() {
	gotNtp := false
	for {
		if !gotNtp {
			// time.Sleep(time.Second)
		} else {
			// time.Sleep(time.Second * 300)
		}

		gotNtp = false

		protocol = make([]byte, 32)
		// Resolve address
		udpAddr, errData := net.ResolveUDPAddr("udp", "cn.ntp.org.cn:"+NTP_PORT_STR)
		if nil != errData {
			log.Errorf("ntp connect err: %v", errData)
			continue
		}
		log.Debugf("ntp Server: %v", udpAddr)
		conn, err := net.DialUDP("udp", nil, udpAddr)
		defer conn.Close()
		if nil != err {
			log.Errorf("ntp net connect error: %v", err)
			continue
		}

		data, packet_len := construct_packet()
		if packet_len == 0 {
			log.Errorf("ntp packet len is 0")
			continue
		}

		//	log.Debugf("ntp begin send: %v, data: %v", packet_len, data)
		conn.SetWriteDeadline(time.Now().Add(time.Second))
		size, err := conn.Write(data)
		if nil != err {
			log.Errorf("ntp write data error: %v", err)
			continue
		} else {
			log.Debugf("ntp write len: %v", size)
		}

		recvBody := make([]byte, 4096)

		conn.SetReadDeadline(time.Now().Add(time.Second))
		size, remoteAddr, err := conn.ReadFromUDP(recvBody)
		if nil != err {
			log.Errorf("ntp read data error: %v", err)
			continue
		} else {
			log.Debugf("ntp[%v] read len: %v", remoteAddr, size)
		}

		var dataStru ntp_packet
		dataStru.transmit_timestamp.coarse = binary.BigEndian.Uint32(recvBody[40:44]) - JAN_1970
		dataStru.transmit_timestamp.fine = uint32(USEC(int64(binary.BigEndian.Uint32(recvBody[44:48]))))

		localUnix := time.Now().Unix()
		TimeDelta = localUnix - int64(dataStru.transmit_timestamp.coarse)
		log.Infof("Got NTPTime :localTime[%v]  ntpTime[%v] TimeDelta(local-ntp)[%v]", localUnix, dataStru.transmit_timestamp.coarse, TimeDelta)

		gotNtp = true
	}
}
