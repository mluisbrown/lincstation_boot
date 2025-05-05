#!/bin/bash
# Start the Management Utility
/usr/local/sbin/emhttp &

# Start the LED Service
cp -r /boot/config/led/ /usr/local/sbin/led/
chmod 777 /usr/local/sbin/led/go_8130_led_linux
nohup /usr/local/sbin/led/go_8130_led_linux >/dev/null 2>&1 &

# Start the RESET Service
cp -r /boot/config/reset/ /usr/local/sbin/reset/
chmod 777 /usr/local/sbin/reset/reset
nohup /usr/local/sbin/reset/reset >/dev/null 2>&1 &

# Start the HDMI Service
cp -r /boot/config/logo/ /usr/local/sbin/logo/
chmod 777 /usr/local/sbin/logo/show
nohup /usr/local/sbin/logo/show >/dev/null 2>&1 &
