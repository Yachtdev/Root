# mqtt+owntracks
https://owntracks.org/
https://github.com/owntracks
install client + brocker
sudo apt install mosquitto-clients
sudo apt-get install mosquitto

test
mosquitto_sub -t owntracks/+/+ 
mosquitto_sub -t test
mosquitto_pub -m test -t test

// set pass & ports
sudo mosquitto_passwd -c /etc/mosquitto/passwd <user>

sudo nano /etc/mosquitto/mosquitto.conf
прописать там нижний файл

sudo nano /etc/mosquitto/conf.d/custom.conf

allow_anonymous false
password_file /etc/mosquitto/passwd
listener 1883 

listener 9001
protocol websockets
allow_anonymous true
connection_messages true

#listener 8883
#certfile /etc/letsencrypt/live/mqtt.example.com/cert.pem
#cafile /etc/letsencrypt/live/mqtt.example.com/chain.pem
#keyfile /etc/letsencrypt/live/mqtt.example.com/privkey.pem
log_type all

- test
mosquitto -c /etc/mosquitto/conf.d/custom.conf

sudo systemctl restart mosquitto
sudo journalctl -xeu mosquitto.service

# web server
sudo apt-add-repository ppa:mosquitto-dev/mosquitto-ppa
sudo apt-get update
sudo apt-get install libmosquitto-dev libcurl4 libcurl4-openssl-dev libconfig-dev liblmdb-dev uuid-dev
sudo apt-get install libsodium-dev libsodium23
wget https://github.com/owntracks/recorder/archive/refs/tags/0.9.6.tar.gz
tar -xvf 0.9.6.tar.gz
cd recorder-0.9.6
cp config.mk.in config.mk
nano config.mk
	WITH_ENCRIPT ?= yes
	WITH_LMDB ?= yes

sudo make 
sudo make install

start
sudo ot-recorder --initialize
sudo nano /etc/default/ot-recorder
//uncomment storage dir
port 1883
host localhost
user recorder
pass test
http host localhost
http port 8083

start systemd
sudo nano /etc/systemd/system/recorder.service

[Unit]
Description=recorder
After=network.target

[Service]
ExecStart=/usr/local/sbin/ot-recorder
Type=simple
Restart=always

[Install]
WantedBy=default.target

sudo systemctl daemon-reload
sudo systemctl start recorder
sudo systemctl enable recorder
sudo systemctl restart recorder
если start too quickly
sudo systemctl reset-failed recorder


MQTT-Explorer.exe

трекинг устройств на моем сервере. как на телефонах, так и на модемах будет работать
https://github.com/owntracks/recorder
web подключение
https://princep.ru

настройка приложения для андроид/иос
приложение называется owntracks
настройки-Подключение к серверу
	протокол 
		MQTT
	адрес 
		адрес princep.ru
		порт 1883
		id клиента - любая строка
		использовать Websocket - отключить
	Идентификация
		Пользователь recorder
		Пароль test
		ID устройства - любая строка
		Метка на карте - любая строка
	Безопасность
		TLS - отключить
В меню статус должно отобразиться подключение ОК

sudo systemctl stop recorder
sudo systemctl disable recorder


# install httpserver
sudo apt install apache2 -y
sudo systemctl start apache2
sudo systemctl enable apache2

- reverse proxy to recorder
sudo a2enmod proxy
sudo a2enmod proxy_http
sudo a2enmod proxy_wstunnel
sudo systemctl restart apache2
sudo nano /etc/apache2/sites-available/000-default.conf

add to host 443 <80>
ProxyPreserveHost On
ProxyPass /owntracks                  http://127.0.0.1:8083/
ProxyPassReverse /owntracks           http://127.0.0.1:8083/
ProxyPass        /owntracks/ws        ws://127.0.0.1:9001/ws keepalive=on retry=60
ProxyPassReverse /owntracks/ws        ws://127.0.0.1:9001/ws keepalive=on

sudo systemctl reload apache2
sudo systemctl restart apache2


# SSL
check
sudo netstat -natp
domain princep.ru
- apache
https://www.digitalocean.com/community/tutorials/how-to-secure-apache-with-let-s-encrypt-on-ubuntu-22-04

sudo apt install certbot python3-certbot-apache
sudo apache2ctl configtest
sudo systemctl reload apache2
sudo certbot --apache
ввести princep.ru 
sudo nano /etc/apache2/sites-available/000-default.conf
sudo apache2ctl configtest
sudo systemctl reload apache2


# install frontend
https://github.com/owntracks/frontend
sudo apt-get install git
sudo apt remove cmdtest
sudo apt remove yarn
curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | sudo apt-key add -
echo "deb https://dl.yarnpkg.com/debian/ stable main" | sudo tee /etc/apt/sources.list.d/yarn.list
sudo apt-get update
sudo apt-get install yarn -y
git clone https://github.com/owntracks/frontend.git
cd frontend
yarn install
yarn build
cp dist ~/server
copy to /var/www/html
Copy public/config/config.example.js to public/config/config.js and make changes as you wish.
See docs/config.md for all available options.
sudo nano /var/www/html/config/config.js
window.owntracks.config = {
  api: {
    baseUrl: "https://princep.ru/owntracks",
  },
};


TODO:
mqtt вебсокет reverse proxy
сертификаты на recorder brocker 
TLS на брокер и на websocket



