const mailer = require('nodemailer');
const config = require('./config');

const smtTransport = mailer.createTransport({
    host: config.mailHost,
    secureConnection: true, // use SSL
    port: config.mailPort,
    secure: true, // secure:true for port 465, secure:false for port 587
    auth: {
        user: config.mailUser,
        pass: config.mailPass//该密码是qq邮箱密码, 不是qq密码, 在邮箱-> 设置 -> pop3... -> 开启服务里面查询.
    }
});

let message = `{ 
    "from": "${config.mailFrom}", 
    "to": "{address}", 
    "subject": "Send Email With Service",
    "text": "Php 是世界上最好的语言",
    "html": "{message}"
}`;

module.exports = {
    mailer: mailer,
    smtTransport: smtTransport,
    message: message
};