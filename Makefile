run: flash debugger

flash: 
	tinygo flash -target arduino ./cmd/main.go
debugger: 
	# https://github.com/solarwinds/tinygo-lessons
	screen -S arduino_debugger /dev/$(shell ls -l /dev | grep usb | grep tty | awk '{print $$10}')

kill:
	# If screen isn't exited correctly, so here's a quick way to kill it
	 screen -X -S arduino_debugger quit
