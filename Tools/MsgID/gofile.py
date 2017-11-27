from const import gofile_path


def genGolangfile(protos):
    fileContent = ""
    fileContent += (
    'package msg'
    '\n'
    'import (\n'
    '	"github.com/name5566/leaf/network/protobuf"\n'
    ')'
    '\n\n'
    'var (\n'
    '	Processor = protobuf.NewProcessor()\n'
    ')\n'
    '\n'
    'func init() {'
     '	// 这里我们注册 protobuf 消息)\n'
     '    Processor.SetByteOrder(true)\n'
    )


    for index in range(len(protos)):
        fileContent += '' \
                       '    Processor.Register(&'
        fileContent += protos[index]
        fileContent += '{})\n'

    fileContent += ("\n}")

    fo = open(gofile_path, "wb")
    fo.write(fileContent.encode('utf-8'))
    fo.close()


