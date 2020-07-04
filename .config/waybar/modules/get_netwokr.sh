network=$(ip route get 1.1.1.1 | grep -Po '(?<=dev\s)\w+' | cut -f1 -d ' ')
interface_easyname=$(dmesg | grep $network | grep renamed | awk 'NF>1{print $NF}')

if ! [ $network ]
then
   network_active="⛔"
elif [[ $interface_easyname == *"wlan"* ]]
then
   network_active=""
else
   network_active=""
fi

echo "$network_active $interface_easyname"