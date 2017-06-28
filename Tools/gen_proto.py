from MsgID.csfile import genCSfile
from MsgID.gofile import genGolangfile
from MsgID.proto import loadProto

protos = loadProto()
genCSfile(protos)
genGolangfile(protos)






