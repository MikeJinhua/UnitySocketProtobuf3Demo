
# UnitySocketPtotobuf3Demo
提供一套基础的游戏框架。这套框架不准备支持大型MMORPG。
主要给中小型团队或者个人开发者提供了基础的框架。可以很快速的开发卡牌，棋牌等开房间，进副本战斗的游戏。

用到的开发语言

1,golang开发服务器。

2,C#开发客户端，用unity。

3,python3 开发工具。

用到的框架

https://github.com/caolaoyao/Chat，

https://github.com/name5566/leaf

表示感谢。如有侵权请及时联系。

目前进展:

1.实现了做手游基础网络通讯。一套.proto配置，一键自动生成服务器客户端消息，两边id自动保证一致，上层调用只管new 出消息类交给网络层，收消息注册用消息类名来注册，屏蔽掉了底层id的转换过程。

2.消息工具已经完成。

3.配表工具基础功能已经完成。一键自动把excel文件拆开成服务器配置和客户端配置，并且生成两边配表代码，游戏启动时候自动加载并且解析配置数据。

4.db可以用gorm，很方便。
gorm多说几句，本来想用gorm自动级联跨表查询和保存，结果发现各种问题各种不习惯，最后还是决定，只用orm做基础的结构体和db table的映射，复合结构还是自己写函数处理下比较和自己习惯。

有个demo文件在UnitySocketProtobuf3Demo\Server\src\server\testgorm.go 实现基本的代码类自动生成数据结构，程序方便用的查询接口，保存和替换数据等。


消息批处理在这里UnitySocketProtobuf3Demo\Tools\GenProto.bat

消息配置文件在这里UnitySocketProtobuf3Demo\Tools\Proto\src\*

配表批处理在这里UnitySocketProtobuf3Demo\Tools\GenTableData.bat

配表文件在这里UnitySocketProtobuf3Demo\Tools\TableData\*

消息工具和配表工具，安装说明在这里UnitySocketProtobuf3Demo/Tools/工具安装说明.txt

用到的几个软件

1.Unity2017。

2.Protobuf3。


