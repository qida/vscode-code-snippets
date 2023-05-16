/**
 * nodejs 脚本
 *
 * 用于将go.json中的代码片段描述更新README.md文件
 *
 * @summary 更新readme文件
 */
const fs = require('fs');
const path = require('path');
const nodeCmd = require('node-cmd');
const packagePath = path.resolve(__dirname, '../package.json');
const process = require('process');
process.on('uncaughtException', function (err) {
    console.log('Caught Exception:' + err);
});

function publishCmd(vsix_name_old) {


    nodeCmd.run(
        `vsce login qida`,
        function (err, data, stderr) {
            if (!err) {
                console.log('success: ', data)
                fs.unlink(vsix_name_old, function (err) {
                    if (err) {
                        return console.error(err);
                    }
                    console.log('删除成功')
                })
            } else {
                console.log('error: ', err)
                console.log('需要安装vsce: npm install --global @vscode/vsce --force')
            }
        }
    );
    nodeCmd.run(
        `vsce publish`,
        function (err, data, stderr) {
            if (!err) {
                console.log('success: ', data)
                fs.unlink(vsix_name_old, function (err) {
                    if (err) {
                        return console.error(err);
                    }
                    console.log('删除成功')
                })
            } else {
                console.log('error: ', err)
                console.log('需要安装vsce: npm install --global @vscode/vsce --force')
            }
        }
    );
    //手动发布 https://marketplace.visualstudio.com/manage/publishers/qida
}

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

function PUBLISH() {
    //获取package.json文件
    let content = getContent(packagePath);
    let obj = JSON.parse(content);
    let vsix_name_old = obj.name + "-" + obj.version + ".vsix";
    publishCmd(vsix_name_old);
}
PUBLISH();