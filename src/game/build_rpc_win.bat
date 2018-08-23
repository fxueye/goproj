RD /S /Q tools\rpc\output
tools\rpc -I tools\rpc\input -O tools\rpc\output -T tools\rpc\temp_rpc

RD /S /Q cmds
MD cmds

COPY /Y tools\rpc\output\*.go cmds
MD cmds\wraps
COPY /Y tools\rpc\output\wrap\*.go cmds\wraps

Pause