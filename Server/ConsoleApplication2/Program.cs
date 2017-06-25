using Net;
using Proto;
using ServerMessage;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Net;
using System.Net.Sockets;
using System.Text;
using System.Threading;
using System.Threading.Tasks;

namespace ConsoleApplication2
{
    class Program
    {
        private static byte[] result = new byte[1024];
        private const int port = 8885;
        private static string ipStr = "127.0.0.1";
        static Socket serverSocket;

        static void Main(string[] args)
        {
            IPAddress ip = IPAddress.Parse(ipStr);
            serverSocket = new Socket(AddressFamily.InterNetwork, SocketType.Stream, ProtocolType.Tcp);
            serverSocket.Bind(new IPEndPoint(ip, port));  //绑定IP地址：端口  
            serverSocket.Listen(10);    //设定最多10个排队连接请求  
            Console.WriteLine("启动监听{0}成功", serverSocket.LocalEndPoint.ToString());
            //通过Clientsoket发送数据  
            Thread myThread = new Thread(ListenClientConnect);
            myThread.Start();
            Console.ReadLine(); 
        }

        private static void ListenClientConnect()
        {
            while(true)
            {
                Socket clientSocket = serverSocket.Accept();
                SignUpResponse sr = new SignUpResponse();
                sr.version = 1f;
                sr.errorCode = 10;
//                 ByteBuffer buff = SendMessage(sr);
//                 clientSocket.Send(WriteMessage(buff.ToBytes()));
                SendMessage(sr, clientSocket);
                Thread receiveThread = new Thread(ReceiveMessage);
                receiveThread.Start(clientSocket);  
            }
        }

        private static void ReceiveMessage(object clientSocket)
        {
            Socket myClientSocket = (Socket)clientSocket;
            while(true)
            {
                try
                {
                    int receiveNumber = myClientSocket.Receive(result);
                    Console.WriteLine("接收客户端{0}消息， 长度为{1}", myClientSocket.RemoteEndPoint.ToString(), receiveNumber);
                    ByteBuffer buff = new ByteBuffer(result);
                    int len = buff.ReadShort();
                    int protoId = buff.ReadShort();
                    if (!ProtoDic.ContainProtoId(protoId))
                    {
                        Console.WriteLine("未知协议号");
                        return;
                    }
                    if (protoId == 1003)
                    {
                        TosChat tos = ProtoBuf.Serializer.Deserialize<TosChat>(new MemoryStream(buff.ReadBytes()));
                        Console.WriteLine(tos.name + "         " + tos.content);
                        TocChat toc = new TocChat();
                        toc.name = "服务端:";
                        toc.content = tos.content;
                        SendMessage(toc, myClientSocket);
                    }
                    else if (protoId == 1002)
                    {

                    }
                }
                catch (Exception ex)
                {
                    Console.WriteLine(ex.Message);
                    myClientSocket.Shutdown(SocketShutdown.Both);
                    myClientSocket.Close();
                    break;
                }
            }
        }

        private static void SendMessage(object obj, Socket clientSocket)
        {
            MemoryStream ms = new MemoryStream();
            ProtoBuf.Serializer.Serialize(ms, obj);
            ByteBuffer buff = new ByteBuffer();
            Type type = obj.GetType();
            int protoId = ProtoDic.GetProtoIdByProtoType(type);
            buff.WriteShort((ushort)protoId);
            buff.WriteBytes(ms.ToArray());
            clientSocket.Send(WriteMessage(buff.ToBytes()));
        }

        private static byte[] WriteMessage(byte[] message)
        {
            MemoryStream ms = null;
            using(ms = new MemoryStream())
            {
                ms.Position = 0;
                BinaryWriter writer = new BinaryWriter(ms);
                ushort msglen = (ushort)message.Length;
                writer.Write(msglen);
                writer.Write(message);
                writer.Flush();
                return ms.ToArray();
            }
        }
    }
}
