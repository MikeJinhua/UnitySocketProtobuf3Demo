from const import csfile_path


def genCSfile(protos):
    fileContent = ""
    fileContent += ('\n'
                    'using Google.Protobuf;\n'
                    'using Msg;\n'
                    'using System;\n'
                    'using System.Collections.Generic;\n\n'
                    'namespace Proto\n'
                    '{\n'
                    '   public class ProtoDic\n'
                    '   {\n'
                    '       private static List<int> _protoId = new List<int>\n'
                    '       {\n')

    for index in range(len(protos)):
        fileContent = fileContent + '''            ''' + str(index) + ''',\n'''

    fileContent += '        };\n\n      private static List<Type>_protoType = new List<Type>\n      {\n'

    for index in range(len(protos)):
        fileContent = fileContent + '''            typeof(''' + protos[index] + '''),\n'''

    fileContent += ('       };\n\n'
                    '       private static readonly Dictionary<RuntimeTypeHandle, MessageParser> Parsers = new Dictionary<RuntimeTypeHandle, MessageParser>()\n'
                    '       {\n')

    for index in range(len(protos)):
        fileContent = fileContent + '''            {typeof(''' + protos[index] + ''').TypeHandle,''' + protos[index] + '''.Parser },\n'''

    fileContent += ('       };\n\n        public static MessageParser GetMessageParser(RuntimeTypeHandle typeHandle)\n'
                   '        {\n'
                   '            MessageParser messageParser;\n'
                   '            Parsers.TryGetValue(typeHandle, out messageParser);\n'
                   '            return messageParser;\n'
                   '        }\n'
                   '\n'
                   '        public static Type GetProtoTypeByProtoId(int protoId)\n'
                   '        {\n'
                   '            int index = _protoId.IndexOf(protoId);\n'
                   '            return _protoType[index];\n'
                   '        }\n'
                   '\n'
                   '        public static int GetProtoIdByProtoType(Type type)\n'
                   '        {\n'
                   '            int index = _protoType.IndexOf(type);\n'
                   '            return _protoId[index];\n'
                   '        }\n'
                   '\n'
                   '        public static bool ContainProtoId(int protoId)\n'
                   '        {\n'
                   '            if(_protoId.Contains(protoId))\n'
                   '            {\n'
                   '                return true;\n'
                   '            }\n'
                   '            return false;\n'
                   '        }\n'
                   '\n'
                   '        public static bool ContainProtoType(Type type)\n'
                   '        {\n'
                   '            if(_protoType.Contains(type))\n'
                   '            {\n'
                   '                return true;\n'
                   '            }\n'
                   '            return false;\n'
                   '        }\n'
                   '    }\n'
                   '}')

    fo = open(csfile_path, "wb")
    fo.write(fileContent.encode('utf-8'))
    fo.close()
