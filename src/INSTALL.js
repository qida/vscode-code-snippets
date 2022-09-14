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
    // var fileName = "ms-ceintl.vscode-language-pack-zh-hans";
    console.log("fileNames:" + fileName);
    // nodeCmd.get(
    //     'code --install-extension ' + fileName ,
    //     function (err, data, stderr) {
    //         console.log(data);
    //     }
    // );
    nodeCmd.run('code --install-extension ' + fileName );
}


function INSTALL() {
    //获取文件
    let content = getContent(packagePath);
    let obj = JSON.parse(content);
    console.log(obj.version);
    console.log(runCmd("qida-go-snippets-"+obj.version+".vsix"));
    // "install": "D:\\software\\VSCode\\bin\\code-insiders.cmd --install-extension qida-go-snippets-0.0.2.vsix",
    // "D:\\software\\VSCode\\bin\\code-insiders.cmd --install-extension qida-go-snippets-0.0.2.vsix"
    // let entries = Object.entries(obj);
    // console.log(entries)
    // entries.map(v => {
    //     mdContent += getTrContent(v[1].prefix, v[1].description);
    // });
}
INSTALL();