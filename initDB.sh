#!/bin/bash

mysqlUser="root"
mysqlHost="mysql"
mysqlPort="3306"
mysqlPassword="123456"

#create database
createDataBase="mysql -h $mysqlHost -u$mysqlUser -p$mysqlPassword"

echo "Create database servicediscovery."
$createDataBase<<EOF
CREATE DATABASE IF NOT EXISTS servicediscovery CHARSET utf8;
EOF

echo "Create table user."
$createDataBase<<EOF
    USE servicediscovery;
    CREATE TABLE user (
        id int(11) NOT NULL AUTO_INCREMENT,
        name char(45) NOT NULL,
        email char(45) DEFAULT NULL,
        PRIMARY KEY (id)
    ) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8;
EOF