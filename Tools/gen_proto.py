from MsgID.csfile import genCSfile
from MsgID.proto import loadProto

protos = loadProto()
genCSfile(protos)







