# GoTests
Testbed for Go developments.

# Hello
Hello world following Go documentation

# GoDevices
A manual DNS HTTP server for knowing the IP of devices unaccessible through company DNS server.
Listens at port 8080. At /, you'll see a list of the registered devices.
At /device/XXXX, register the device with name XXXX with its current IP.
My use case is to put a wget request on cron of Linux devices to make it easier to SSH to them from Windows machines.
