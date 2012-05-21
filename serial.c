#include <fcntl.h>
#include <termios.h>

int serial_open(char *port, int baud) {
	struct termios options;
	int fd;

	fd = open(port, O_RDWR);
	if (fd == -1) {
		return fd;
	}

	tcgetattr(fd, &options);
	cfsetispeed(&options, baud);
	cfsetospeed(&options, baud);
	// CSIZE IS SO FUCKING IMPORTANT!
	// e.g.: Mac OS X seems to set CS5 by default, so only *setting*
	// CS8 doesn’t fix it. You have to *unset* all size bits others as well.
	// Unsetting CSIZE unsets all CS* bits
	options.c_cflag &= ~(PARENB | CSIZE | CSTOPB | CSIZE);
	options.c_cflag |= (CLOCAL | CREAD | CS8);

	options.c_iflag &= ~(ICRNL | INLCR | INPCK | IXON | IXOFF | IXANY | IGNBRK | BRKINT | PARMRK | ISTRIP | IGNCR);

	options.c_oflag &= ~(ONLCR | OCRNL | OPOST);

	options.c_lflag &= ~(ECHO | ECHONL | ICANON | ISIG | IEXTEN);
	tcsetattr(fd, TCSANOW, &options);
	return fd;
}
