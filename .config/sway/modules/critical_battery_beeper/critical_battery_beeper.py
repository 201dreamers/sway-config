#!/usr/bin/python3

import os
from time import sleep


def get_battery_status():
    return open('/sys/class/power_supply/BAT0/status', 'r').readline().strip()


def get_battery_capacity():
    return int(open('/sys/class/power_supply/BAT0/capacity', 'r').readline().strip())


while True:
    if get_battery_status() == "Discharging":
        battery_capacity = get_battery_capacity()
        if battery_capacity <= 10:
            os.system('mpg123 -q ~/.config/sway/modules/critical_battery_beeper/critical_battery.mp3')
    sleep(60)
