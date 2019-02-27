module.exports = (function (env) {
    if (env === 'test') {
        return {
            mailHost: 'smtp.qq.com',
            mailPort: 465,
            mailUser: '1005323491@qq.com',
            mailPass: '',
            mailFrom: '记忆【时光，<1005323491@qq.com>',
            etcdRegistryHost: 'http://registry',
            etcdRegistryPort: 8080,
            port: 7070
        };
    } else if (env === 'product') {
        return {
            etcdRegistryHost: 'service-registry',
            etcdRegistryPort: 8080,
            port: 8080
        };
    }
}(process.env.NODE_ENV || 'test'));