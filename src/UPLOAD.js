/**
 * nodejs 脚本
 * 上传vsix文件
 * @summary 上传vsix文件
 */

const Ftp = require('ftp');
const fs = require('fs');
const path = require('path');
const process = require('process');
const client = new Ftp();

process.on('uncaughtException', function (err) {
    console.log('Caught Exception:' + err);
});


function fromDir(startPath, filter, callback) {
    //console.log('Starting from dir '+startPath+'/');

    if (!fs.existsSync(startPath)) {
        console.log("no dir ", startPath);
        return;
    }

    var files = fs.readdirSync(startPath);
    for (var i = 0; i < files.length; i++) {
        var filename = path.join(startPath, files[i]);
        var stat = fs.lstatSync(filename);
        if (stat.isDirectory()) {
            fromDir(filename, filter, callback); //recurse
        }
        else if (filter.test(filename)) callback(filename);
    };
};
function UPLOAD() {
    fromDir('./', /\.vsix$/, function (filename) {
        console.log('-- found: ', filename);
        client.on('ready', function () {
            client.put(filename, filename, function (err) {
                if (err) throw err;
                console.log("上传成功");
                client.end();
            });
        });
    });
    client.connect({
        host: '10.0.0.2',
        port: '21',
        user: 'vsix',
        password: 'vsixvsix',
        keepalive: 5000
    });
}
UPLOAD();