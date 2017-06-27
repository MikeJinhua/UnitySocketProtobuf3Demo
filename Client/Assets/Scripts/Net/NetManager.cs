using UnityEngine;
using Util;
using System.Collections.Generic;
using System;
using Proto;
using System.IO;
using Google.Protobuf;

namespace Net
{
    public class NetManager : SingletonMonoBehaviour<NetManager>
    {
        private Dictionary<Type, TocHandler> _handlerDic;
        private SocketClient _socketClient;
        SocketClient socketClient
        {
            get
            {
                if (_socketClient == null)
                {
                    _socketClient = new SocketClient();
                }
                return _socketClient;
            }
        }

        void Start()
        {
            Init();
        }

        public void Init()
        {
            _handlerDic = new Dictionary<Type, TocHandler>();
            socketClient.OnRegister();
        }

        /// <summary>
        /// 发送链接请求
        /// </summary>
        public void SendConnect()
        {
            socketClient.SendConnect();
        }

        /// <summary>
        /// 关闭网络
        /// </summary>
        public void OnRemove()
        {
            socketClient.OnRemove();
        }

        /// <summary>
        /// 发送SOCKET消息
        /// </summary>
        public void SendMessage(ByteBuffer buffer)
        {
            socketClient.SendMessage(buffer);
        }

        /// <summary>
        /// 发送SOCKET消息
        /// </summary>
        public void SendMessage(IMessage obj)
        {
            if (!ProtoDic.ContainProtoType(obj.GetType()))
            {
                Debug.LogError("不存协议类型");
                return;
            }
            ByteBuffer buff = new ByteBuffer();
            int protoId = ProtoDic.GetProtoIdByProtoType(obj.GetType());

            byte[] result;
            using (MemoryStream ms = new MemoryStream())
            {
                obj.WriteTo(ms);
                result = ms.ToArray();
            }

            UInt16 lengh = (UInt16)(result.Length + 2);
            Debug.Log("lengh" + lengh + ",protoId" + protoId);
            buff.WriteShort((UInt16)lengh);

            buff.WriteShort((UInt16)protoId);
            buff.WriteBytes(result);
            SendMessage(buff);
        }

        /// <summary>
        /// 连接 
        /// </summary>
        public void OnConnect()
        {
            Debug.Log("======连接========");
        }

        /// <summary>
        /// 断开连接
        /// </summary>
        public void OnDisConnect()
        {
            Debug.Log("======断开连接========");
        }

        /// <summary>
        /// 派发协议
        /// </summary>
        /// <param name="protoId"></param>
        /// <param name="buff"></param>
        public void DispatchProto(int protoId, byte[] buff)
        {
            if (!ProtoDic.ContainProtoId(protoId))
            {
                Debug.LogError("未知协议号");
                return;
            }
            Type protoType = ProtoDic.GetProtoTypeByProtoId(protoId);
            try
            {
                MessageParser messageParser = ProtoDic.GetMessageParser(protoType.TypeHandle);
                object toc = messageParser.ParseFrom(buff);
                sEvents.Enqueue(new KeyValuePair<Type, object>(protoType, toc));
            }
            catch
            {
                Debug.Log("DispatchProto Error:" + protoType.ToString());
            }

        }

        static Queue<KeyValuePair<Type, object>> sEvents = new Queue<KeyValuePair<Type, object>>();
        /// <summary>
        /// 交给Command，这里不想关心发给谁。
        /// </summary>
        void Update()
        {
            if (sEvents.Count > 0)
            {
                while (sEvents.Count > 0)
                {
                    KeyValuePair<Type, object> _event = sEvents.Dequeue();
                    if (_handlerDic.ContainsKey(_event.Key))
                    {
                        _handlerDic[_event.Key](_event.Value);
                    }
                }
            }
        }

        public void AddHandler(Type type, TocHandler handler)
        {
            if (_handlerDic.ContainsKey(type))
            {
                _handlerDic[type] += handler;
            }
            else
            {
                _handlerDic.Add(type, handler);
            }
        }
    }

}
