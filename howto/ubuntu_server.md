# настройка сервера ubuntu 22.04 
- create sudo user
useradd -m <user>
usermod -aG sudo <user>
nano /etc/passwd
- change sh to bash
passwd <user>
exit
- login from terminal

# server security 
на юзера установить длинный пароль 
sudo nano /etc/ssh/sshd_config
- PermitRootLogin no
- ssh port <port>
service ssh restart

sudo su
iptables -N ssh_limiter
iptables -A ssh_limiter -p tcp -m tcp --dport <port> -m conntrack --ctstate NEW -m recent --set --name SSH --mask 255.255.255.255 --rsource
iptables -A ssh_limiter -p tcp -m tcp --dport <port>  -m conntrack --ctstate NEW -m recent --update --seconds 180 --hitcount 3 --name SSH --mask 255.255.255.255 --rsource -j DROP
iptables -A ssh_limiter -p tcp -m tcp --dport <port>  -j ACCEPT
iptables -A INPUT -j ssh_limiter

sudo apt-get install iptables-persistent
/etc/iptables/rules.v4
очистка бан-листа вручную:
sudo su
echo / > /proc/net/xt_recent/SSH
"SSH" имя листа из правил выше

netstat -nat

