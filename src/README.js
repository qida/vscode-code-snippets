/**
 * nodejs 脚本
 *
 * 用于将go.json中的代码片段描述更新README.md文件
 *
 * @summary 更新readme文件
 */

const fs = require('fs');
const path = require('path');
const snippetsDirectoryPath = path.resolve(__dirname, '../snippets/');
const mdFilePath = path.resolve(__dirname, '../README.md');

const process = require('process');
process.on('uncaughtException', function (err) {
    console.log('Caught Exception:' + err);
});


const thanks = [
    "<https://blog.csdn.net/weixin_36719607/article/details/103345353>",
    "<https://github.com/masterZSH/vscode-code-snippets>"
];


//获取目录下文件
function getPathFiles(parentPath, out) {
    try {
        let files = fs.readdirSync(parentPath);
        files.forEach(function (item) {
            let tempPath = path.join(parentPath, item);
            let stats = fs.statSync(tempPath);
            if (stats.isDirectory()) {
                getPathFiles(tempPath, out);
            } else {
                out.push(tempPath);
            }
        });
        return out;
    } catch (e) {
        console.warn("Path Error:" + parentPath);
        return out;
    }
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

function UpdateREADME() {
    //获取文件
    let snippetsFiles = [];
    console.log(getPathFiles(snippetsDirectoryPath, snippetsFiles));
    let writeStream = fs.createWriteStream(mdFilePath, {
        flags: 'w+',
        encoding: "utf8"
    });
    writeStream.on('error', function (err) {
        console.log(err);
    });
    let mdContent = "#\r\n";
    snippetsFiles.forEach(function (item) {
        let content = getContent(item);
        let obj = JSON.parse(content);
        let entries = Object.entries(obj);
        mdContent += getHeader(path.basename(item));
        mdContent += getTableHeader();
        mdContent += getTableContent();
        entries.map(v => {
            mdContent += getTrContent(v[1].prefix, v[1].description);
        });
    })
    mdContent += getFooterContent();
    writeStream.write(mdContent, 'utf8');
    writeStream.end();
}


/**
 * 获取头部
 */
function getHeader(obj) {
    return "\r\n## 用法 Usage <" + obj + ">\r\n\r\n";
}

/**
 * 获取table头部
 */
function getTableHeader() {
    return '|前缀 prefix|说明 description|\r\n';
}

/**
 * 获取table内容
 */
function getTableContent() {
    return "|---|---|\r\n";
}

/**
 * 获取table tr内容
 * @param {string} prefix -前置
 * @param {string} description -描述
 */
function getTrContent(prefix, description) {
    return `|${prefix}|${description}|\r\n`;
}


/**
 * 获取footer内容
 */
function getFooterContent() {
    let content = "\r\n## Thanks  \r\n\r\n"
    thanks.forEach(function (item) {
        content += item + "  \r\n"
    })
    return content;
}

UpdateREADME();