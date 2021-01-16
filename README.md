# Arduino SSD1306 Game of Life

Game of Life running on an Arduino and SSD1306 LCD display.

## Hardware

A regular Arduino Uno, although I suspect many GPIO devices would work here. As well as the SSD1306 LCD display.

You can find these all over. I personally used Alibaba with the following items:

- Arduino Uno kit: https://www.aliexpress.com/item/32543752387.html
- SSD1306: https://www.aliexpress.com/item/32896971385.html

## Getting Started

This is on MacOS, so your milage may vary. I have added the versions to the installations to ensure compatibility.

Install Go: 
```
brew install go@1.15
```

Install TinyGo:
```
brew tap tinygo-org/tools
brew install tinygo@0.16.0
```

Install avr flashing tools:
```
brew tap osx-cross/avr
brew install avr-gcc@9.3.0
brew install avrdude@6.3
```
