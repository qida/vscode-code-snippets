
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
    // $env:VSCOE_HOME="$env:S\:\\VSCode-insider";$env:PATH+=";$env:VSCOE_HOME\\bin";
    //安装
    //需要将code-insiders添加到系统环境变量
    nodeCmd.run(
        `code-insiders --enable-proposed-api --disable-workspace-trust --install-extension ${fileName}`,
        function (err, data, stderr) {
            if (err != null) {
                console.log("err: ", err, "stderr: ", stderr)
            }
            console.log("data", data)
        }
    )
}

function INSTALL() {
    //获取文件名
    let vsix_name = JSON.parse(getContent(packagePath)).name + "-" + JSON.parse(getContent(packagePath)).version + ".vsix";
    runCmd(vsix_name);
}
INSTALL();