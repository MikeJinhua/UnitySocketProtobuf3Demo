以windows为例，工具代码全是python的，可以自己写对应平台批处理来支持其他平台。
1.python的安装
windows
自己安装python3，去官网自己下载。然后找到安装路径里面的python.exe，复制出来一个改名叫python3.exe。
如果你安装过python2，出现冲突。说明你是老司机，自行解决。
mac & linux 如果不自带python3 自行安装
2第三方库的安装
(理论上讲命令行直接拷贝就行的，真的报错自己谷歌搜吧，环境变量找不到pip之类的错误很简单。)
pip install xlrd // 读取excel用

3使用。
GenProto.bat，生成两端消息文件和对应的id管理类
GenTableData.bat，生成两端配表和初始化统一接口。
