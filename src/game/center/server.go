package center
import (
	server "game/server/server"
	"os"

	log "github.com/cihub/seelog"
	"github.com/pkg/profile"
)


var (
	prof interface {
		Stop()
	}
)

func Run() {
	defer func() {
		if prof != nil {
			prof.Stop()
		}
		if err := recover(); err != nil {
			log.Critical(err)
			os.Exit(0)
		}
	}()
	prof = profile.Start(profile.MemProfile)
	logger, err := log.LoggerFromConfigAsFile("config/log.xml")
	if err != nil {
		log.Critical("err parsing config log file", err)
		return
	}
	log.ReplaceLogger(logger)
	server.Init()
	server.Instance.Start()

}
