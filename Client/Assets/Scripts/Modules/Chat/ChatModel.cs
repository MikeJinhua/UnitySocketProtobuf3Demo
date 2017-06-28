using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using Msg;
using UnityEngine;
using Util;
using Net;

public class ChatModel : BaseModel<ChatModel>
{
    protected override void InitAddTocHandler()
    {
        AddTocHandler(typeof(TocChat), STocChat);
    }

    private void STocChat(object data)
    {
       
        TocChat toc = data as TocChat;
        Debug.Log("STocChat" + toc.Name + toc.Content);
        if (ChatView.Exists)
        {
             string content = toc.Name + ":" + toc.Content;
             Debug.Log(content);
             ChatView.Instance.AddChatItem(content);
        }
    }

    public void CTosChat(string name , string content)
    {
        TosChat tos = new TosChat();
        tos.Name = name;
        tos.Content = content;
        
        SendTos(tos);
    }
}