#!/bin/bash

if [ ! "$(ls -A /var/lib/mysql)" ]; then
    /usr/bin/mysql_install_db
fi

/etc/init.d/mysql start
/usr/bin/mysqladmin -u root password '123456'

mysql -u root -p123456 -e "CREATE DATABASE gettingLogs CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"
mysql -u root -p123456 -e "CREATE USER 'gettingLogs'@'%' IDENTIFIED BY '123456'"
mysql -u root -p123456 -e "GRANT ALL PRIVILEGES ON gettingLogs . * TO 'gettingLogs'@'%'"
mysql -u root -p123456 -e "FLUSH PRIVILEGES"

mysql -u gettingLogs -p123456 -e "CREATE TABLE gettingLogs.logs (
id int(11) unsigned NOT NULL AUTO_INCREMENT,
uuid char(36) NOT NULL DEFAULT '',
ip char(15) NOT NULL DEFAULT '',
user_uuid char(36) NOT NULL DEFAULT '',
timestamp int(32) NOT NULL,
url varchar(20) NOT NULL DEFAULT '',
dataRequest text NOT NULL,
dataResponse text NOT NULL,
PRIMARY KEY (id)
) ENGINE=InnoDB AUTO_INCREMENT=12061 DEFAULT CHARSET=latin1;"

while true ; do sleep 5; done;