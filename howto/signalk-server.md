# установка и обновление node & npm
sudo apt-get update
curl -fsSL https://deb.nodesource.com/setup_18.x | sudo -E bash -
sudo apt-get install nodejs
node -v
18.19
npm install -g npm@10.3.0
npm -v
10.3.0

# установка сервера
https://github.com/signalk
https://github.com/SignalK/signalk-server

sudo npm install -g signalk-server
- проверка тестовых данных
- в браузере serverip:3000
signalk-server --sample-nmea0183-data
signalk-server --sample-n2k-data
- настройка
sudo signalk-server-setup
port 3000
SSL no

systemctl daemon-reload
systemctl enable signalk.service
systemctl enable signalk.socket
systemctl stop signalk.service
systemctl restart signalk.socket
systemctl restart signalk.service

systemctl status signalk.service
systemctl status signalk.socket

restart system

127.0.0.1:3000
admin admin
settings mdsn on
после настройки data connection - restart
добавим данные с симулятора 
https://github.com/panaaj/nmeasimulator/releases


plugin
add
signalk-mqtt-bridge	
