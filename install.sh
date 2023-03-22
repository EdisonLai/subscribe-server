sudo sed -i -e "s|mirrorlist=|#mirrorlist=|g" /etc/yum.repos.d/CentOS-*
sudo sed -i -e "s|#baseurl=http://mirror.centos.org|baseurl=http://vault.centos.org|g" /etc/yum.repos.d/CentOS-*
sudo sed -i 's/^SELINUX=.*/SELINUX=disabled/' /etc/selinux/config
sudo setenforce 0

# install warp ipv6
bash <(curl -fsSL git.io/warp.sh) wg
curl -fsSL git.io/wgcf.sh | sudo bash
echo Y | wgcf register
wgcf generate
sed -i -e '/0.0.0.0/d' wgcf-profile.conf

sudo cp wgcf-profile.conf /etc/wireguard/wgcf.conf
sudo wg-quick up wgcf
sudo wg-quick down wgcf
sudo systemctl start wg-quick@wgcf
sudo systemctl enable wg-quick@wgcf

sudo yum install wget

wget -N --no-check-certificate -q -O install.sh "https://raw.githubusercontent.com/wulabing/V2Ray_ws-tls_bash_onekey/master/install.sh" && chmod +x install.sh && bash install.sh
