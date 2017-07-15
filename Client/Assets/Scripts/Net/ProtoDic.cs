
using Google.Protobuf;
using Msg;
using System;
using System.Collections.Generic;

namespace Proto
{
   public class ProtoDic
   {
       private static List<int> _protoId = new List<int>
       {
            0,
            1,
            2,
            3,
            4,
            5,
            6,
            7,
            8,
            9,
        };

      private static List<Type>_protoType = new List<Type>
      {
            typeof(StartFight),
            typeof(FightResult),
            typeof(EnterFight),
            typeof(SignUpResponse),
            typeof(TosChat),
            typeof(TocChat),
            typeof(Login),
            typeof(PlayerBaseInfo),
            typeof(LoginSuccessfull),
            typeof(LoginFaild),
       };

       private static readonly Dictionary<RuntimeTypeHandle, MessageParser> Parsers = new Dictionary<RuntimeTypeHandle, MessageParser>()
       {
            {typeof(StartFight).TypeHandle,StartFight.Parser },
            {typeof(FightResult).TypeHandle,FightResult.Parser },
            {typeof(EnterFight).TypeHandle,EnterFight.Parser },
            {typeof(SignUpResponse).TypeHandle,SignUpResponse.Parser },
            {typeof(TosChat).TypeHandle,TosChat.Parser },
            {typeof(TocChat).TypeHandle,TocChat.Parser },
            {typeof(Login).TypeHandle,Login.Parser },
            {typeof(PlayerBaseInfo).TypeHandle,PlayerBaseInfo.Parser },
            {typeof(LoginSuccessfull).TypeHandle,LoginSuccessfull.Parser },
            {typeof(LoginFaild).TypeHandle,LoginFaild.Parser },
       };

        public static MessageParser GetMessageParser(RuntimeTypeHandle typeHandle)
        {
            MessageParser messageParser;
            Parsers.TryGetValue(typeHandle, out messageParser);
            return messageParser;
        }

        public static Type GetProtoTypeByProtoId(int protoId)
        {
            int index = _protoId.IndexOf(protoId);
            return _protoType[index];
        }

        public static int GetProtoIdByProtoType(Type type)
        {
            int index = _protoType.IndexOf(type);
            return _protoId[index];
        }

        public static bool ContainProtoId(int protoId)
        {
            if(_protoId.Contains(protoId))
            {
                return true;
            }
            return false;
        }

        public static bool ContainProtoType(Type type)
        {
            if(_protoType.Contains(type))
            {
                return true;
            }
            return false;
        }
    }
}