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
    console.log("fileNames:" + fileName);
    ////先卸载
    // nodeCmd.run(
    //     `code-insiders --uninstall-extension ${fileName}`,
    //     function (err, data, stderr) {
    //         console.log('examples dir now contains the example file along with : ', data)
    //     }
    // );

    //安装
    nodeCmd.run(
        //需要将code-insiders添加到系统环境变量
        `code-insiders --install-extension ${fileName}`,
        function (err, data, stderr) {
            console.log('Good Job : ', data)
        }
    );
}

function INSTALL() {
    //获取文件
    runCmd("qida-go-snippets-" + JSON.parse(getContent(packagePath)).version + ".vsix");
}
INSTALL();