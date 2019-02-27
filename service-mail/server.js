const mail = require('./mail');
const config = require('./config');
const http = require('http');
const format = require('string-template');
const request = require('request');

setInterval(() => {
    let url = config.etcdRegistryHost + ':' + config.etcdRegistryPort;
    request.post({
        url: url,
        body: JSON.stringify({
            name: 'service/mail',
            host: 'mail',
            port: config.port,

        })
    }, function (err, resp, ) {
        if (err) {
            console.log(err);
            return;
        }
        if (resp.statusCode !== 200) {
            console.log('Service registry error.');
        }
    });
}, 2 * 1000);

const server = http.createServer((function (req, res) {
    if (req.method !== 'POST') {
        res.statusCode = 411;
        res.end(JSON.stringify({
            message: 'Method must be post.'
        }));
        return;
    }

    let body = '';
    req.on('data', function (chunk) {
        body = body + chunk;
    });

    req.on('error', function (err) {
        res.statusCode = 411;
        res.end(JSON.stringify({ message: err }));
    });

    req.on('end', function () {
        if (body.toString() === '') {
            body = '{}';
        }
        let bodyJson = JSON.parse(body.toString());
        let message = format(mail.message, {
            address: bodyJson.address,
            message: bodyJson.message
        });
        let messageJson = JSON.parse(message);
        mail.smtTransport.sendMail(messageJson, function (err, ) {
            if (err) {
                res.statusCode = 411;
                res.end(JSON.stringify({ message: err }));
                console.log('MailSendError: ', err);
                return;
            }
            mail.smtTransport.close();
            console.log('MailSendSuccess: send to ', messageJson.to, ', message ', messageJson.html);
            res.statusCode = 200;
            res.end(JSON.stringify({
                message: 'success'
            }));
            return;
        });
    });

}));

server.listen(config.port);