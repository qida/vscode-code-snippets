
const fs = require('fs');
const path = require('path');
const nodeCmd = require('node-cmd');
const packagePath = path.resolve(__dirname, '../package.json');
const process = require('process');
process.on('uncaughtException', function (err) {
    console.log('Caught Exception:' + err);
});

/**
 * 获取文件内容
 * @param {string} filePath -文件地址
 * @returns {string} -返回的内容
 */
function getContent(filePath) {
    let snippetsFilePath = filePath;
    if (!fs.existsSync(snippetsFilePath)) {
        throw new RangeError("no file exists");
    }
    let contentBuffer = fs.readFileSync(snippetsFilePath, {
        encoding: 'utf8'
    });
    let content = contentBuffer.toString();
    return content;
} 

function runCmd(fileName) {
    nodeCmd.run(
        `code-insiders --enable-proposed-api --disable-workspace-trust --install-extension ${fileName}`,
        function (err, data, stderr) {
            if (err != null) {
                console.log("安装插件出错: ", err);
            } else {
                console.log("插件安装成功");
            }
        }
    );
    nodeCmd.run(
        `code --enable-proposed-api --disable-workspace-trust --install-extension ${fileName}`,
        function (err, data, stderr) {
            if (err != null) {
                console.log("安装插件出错: ", err);
            } else {
                console.log("插件安装成功");
            }
        }
    );
}

function INSTALL() {
    //获取文件名
    let vsix_name = JSON.parse(getContent(packagePath)).name + "-" + JSON.parse(getContent(packagePath)).version + ".vsix";
    runCmd(vsix_name);
}
INSTALL();