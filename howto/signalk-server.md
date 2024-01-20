# настройка сервера ubuntu 22.04 
- create sudo user
useradd -m <user>
usermod -aG sudo <user>
nano /etc/passwd
- change sh to bash
passwd <user>
exit
- login from terminal

# TODO безопасность
- запретить вход под рутом
- вход по ключу
- сменить порт ssh
- настроить таймауты входа

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
sudo npm install -g signalk-server
- проверка тестовых данных
- в браузере serverip:3000
signalk-server --sample-nmea0183-data
- настройка


