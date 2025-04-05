# Checktemp
Bare-bones Go program that reads temperature from DS18B20 sensors and serves the
data as JSON over HTTP. I guess you can call this a "temperature API"?

## Compiling
I run this on a Raspberry Pi 2B+, I like to cross-compile from my AMD64
computer:

```
bash
GOARCH=arm go build -ldflags "-s -w" -o checktemp-arm
scp checktemp-arm raspberry-pi2:~/checktemp
```

## Usage
Just run the program. If you need it to keep running, make a Systemd unit.
```
bash
./checktemp
```

## License
This is Free and Open Source Software, made within a Saturday evening and
available for you as BSD-3 Clause license. Check the license file for legal
details.
