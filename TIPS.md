# 编写提示

---

## 官方模板

``` json
"Print to console": {
 "prefix": "log",
 "body": [
  "console.log('$1');",
  "$2"
 ],
 "description": "Log output to console"
}

```

prefix 前缀，emmet 触发条件，例如上述片段，输入log 时会触发 emmet 提示  
body 片段主体，数组形式  
description 描述，emmet 提示的显示内容  

---

### 基本语法

#### body 内一个数组元素即代表一行代码

#### 使用 $ 插入用户自定义内容

#### $number 数字的大小表示光标的先后顺序，使用 tab键移动光标位置，相同序号的内容会同时更改

#### $0 表示光标最终位置

#### ${number: defaultContent} 为插入内容设置默认值

#### ${number|select1,select2|} 为插入内容提供可选项

#### $name 或 ${name: default} 插入变量值，若变量名未定义则插入变量名

---

### 预定义变量名称

* 文件

#### TM_SELECTED_TEXT：当前选定的文本或空字符串

#### TM_CURRENT_LINE：当前行的内容

#### TM_CURRENT_WORD：光标所处单词或空字符串

#### TM_LINE_INDEX：行号（从零开始）

#### TM_LINE_NUMBER：行号（从一开始）

#### TM_FILENAME：当前文档的文件名

#### TM_FILENAME_BASE：当前文档的文件名（不含后缀名）

#### TM_DIRECTORY：当前文档所在目录

#### TM_FILEPATH：当前文档的完整文件路径

---

* 剪贴板

#### CLIPBOARD：当前剪贴板中内容

---

* 时间

#### CURRENT_YEAR: 当前年份

#### CURRENT_YEAR_SHORT: 当前年份的后两位

#### CURRENT_MONTH: 格式化为两位数字的当前月份，如 02

#### CURRENT_MONTH_NAME: 当前月份的全称，如 July

#### CURRENT_MONTH_NAME_SHORT: 当前月份的简称，如 Jul

#### CURRENT_DATE: 当天月份第几天

#### CURRENT_DAY_NAME: 当天周几，如 Monday

#### CURRENT_DAY_NAME_SHORT: 当天周几的简称，如 Mon

#### CURRENT_HOUR: 当前小时（24 小时制）

#### CURRENT_MINUTE: 当前分钟

#### CURRENT_SECOND: 当前秒数
