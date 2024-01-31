# mqtt+owntracks
https://owntracks.org/
https://github.com/owntracks
install client + brocker
sudo apt install mosquitto-clients
sudo apt-get install mosquitto

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


test
mosquitto_sub -t owntracks/+/+ 
mosquitto_sub -t test
mosquitto_pub -m test -t test

// set pass & portssudo 
новый файл
sudo mosquitto_passwd -c /etc/mosquitto/passwd <user>
добавить юзера
sudo mosquitto_passwd  /etc/mosquitto/passwd <user>

# назначить права
sudo chmod 0700 /etc/mosquitto/passwd
sudo chown mosquitto /etc/mosquitto/passwd
sudo chgrp mosquitto /etc/mosquitto/passwd
cd /etc/letsencrypt
sudo chgrp mosquitto live/ archive/
sudo chmod g+rx  live/ archive/
sudo chgrp mosquitto archive/princep.ru/archive/privkey1.pem
sudo chmod g+r archive/princep.ru/archive/privkey1.pem

sudo nano /etc/mosquitto/mosquitto.conf
прописать там нижний файл

sudo nano /etc/mosquitto/conf.d/custom.conf

allow_anonymous false
password_file /etc/mosquitto/passwd
listener 1883 0.0.0.0

#listener 9001 127.0.0.1
#protocol websockets
#connection_messages true

listener 8883 0.0.0.0
certfile /etc/letsencrypt/live/princep.ru/cert.pem
cafile /etc/letsencrypt/live/princep.ru/chain.pem
keyfile /etc/letsencrypt/live/princep.ru/privkey.pem

listener 9083 0.0.0.0
protocol websockets
connection_messages true
certfile /etc/letsencrypt/live/princep.ru/cert.pem
cafile /etc/letsencrypt/live/princep.ru/chain.pem
keyfile /etc/letsencrypt/live/princep.ru/privkey.pem

log_type all

- test
sudo mosquitto -c /etc/mosquitto/conf.d/custom.conf

sudo systemctl stop mosquitto
sudo systemctl restart mosquitto
sudo journalctl -xeu mosquitto.service

# recorder
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

OTR_HOST 127.0.0.1
OTR_PORT 1883
OTR_USER recorder
OTR_PASS test
OTR_HTTPHOST localhost
OTR_HTTPPORT 8083

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
sudo systemctl status recorder
если start too quickly
sudo systemctl reset-failed recorder


MQTT-Explorer.exe

трекинг устройств на моем сервере. как на телефонах, так и на модемах будет работать
https://github.com/owntracks/recorder
web подключение
https://princep.ru

настройка приложения для андроид/иос нешифрованное MQTT
приложение называется owntracks
настройки-Подключение к серверу
	протокол 
		MQTT
	адрес 
		адрес princep.ru
		порт 8883/9083
		id клиента - любая строка
		использовать Websocket - как угодно, порт для вебсокета 9083
	Идентификация
		Пользователь - свой пользователь, зарегистрированный на сервере 
		Пароль тоже зарегистрированный на сервере
		ID устройства - любая строка
		Метка на карте - любая строка
	Безопасность
		TLS - включить
		остальные поля не заполнять
В меню статус должно отобразиться "Подключено"

sudo systemctl stop recorder
sudo systemctl disable recorder

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

удалить данные сервера
/var/spool/owntracks/recorder/store/last/<user>/<id>/
/var/spool/owntracks/recorder/store/rec/<user>/<id>/

TODO:
убрать логины/пароль recorder вынести в private/ сменить пароль
в файле все заменить на <host> <user> <pass>

mqtt вебсокет reverse proxy
форму регистрации/авторизации для mosquitto
проверить server/owntracks из веба



