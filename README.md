# Dotfiles for sway wm configuration

### Screenshots of my sway configuration
![Screenshot 7](./screenshots/screenshot-7.png "Screenshot 7")

![Screenshot 8](./screenshots/screenshot-8.png "Screenshot 8")

### Used packages:
* wofi - menu launcher
* bmenu - alternative menu
* swaylock-fancy - lockscreen
* grim - screenshot and colorpick utility
* slurp - utility for selecting an area
* brightnessctl - for brightness control
* pulseaudio/pavucontrol - for audio control
* waybar - bar for swaywm
* nm-applet/network-manager-applet-indicator - wifi applet
* mako - nitification manager
* earlyoom - out of memory manager
* redshift/redshift-wayland - blue light filter
* blueman/blueman-applet - bluetooth applet
* parcellite - clipboard manager
* gthumb - image viewer
* nautilus - file manager
* mpg123 - play sound from terminal
* awesome font
* (polkit-gnome-authentication, qt5ct, tlp, tlp-rdw)

### Some features

#### Scratchpad indicator:
Shows how much windows is hidden in scratchpad
* Left mouse click - move to scratchpad
* Right mouse click - show windows from scratchpad

#### Battery:

##### Beeper
Plays sound when capacity <= 10% every 60 seconds

##### Indicator
* green - charging
* default - discharging and capacity > 20%
* yellow - discharging and 10% < capacity < 20%
* red, blinking - discharging and capacity < 10%

#### Cpu indicator
* default - normal load
* yellow - load > 70%
* red - load > 90%

### Shortcuts ($mod = Super/Windows logo key on keyboard)
* $mod+Shift+r - reload
* $mod+Return - start terminal
* $mod+q - kill
* $mod+a - start applications menu
* $mod+b - start alternative terminal menu
* $mod+d - start run menu
* $mod+w - start browser
* $mod+e - start file manager
* $mod+0 - lock screen
* $mod+p - pick color
* Print - Take a screenshot of whole screen and copy it to clipboard
* $mod+Print - Take a screenshot of selected region and copy it to clipboard
* $mod+Shift+Print - Take a screenshot of focused window and copy it to clipboard
* Ctrl+Print - Take a screenshot of whole screen and copy to $HOME/Pictures/screenshots (You should create this folder)
* $mod+Ctrl+Print - Take a screenshot of selected region and copy to $HOME/Pictures/screenshots
* $mod+Ctrl+Shift+Print - Take a screenshot of focused window and copy to $HOME/Pictures/screenshots
* $mod+$left - focus left
* $mod+$down - focus down
* $mod+$up - focus up
* $mod+$right - focus right
* Soon
