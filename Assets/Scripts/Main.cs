using UnityEngine;
using System.Collections;
using Net;

public class Main : MonoBehaviour 
{
	void Start () 
    {
        InitNet();
        InitManager();
//        ChatView.OpenView("Prefabs/ChatView/ChatView");
	}

    private void InitNet()
    {
        gameObject.AddComponent<NetManager>();
        NetManager.Instance.SendConnect();
    }

    private void InitManager()
    {
        gameObject.AddComponent<LoginModel>();
        gameObject.AddComponent<ChatModel>();
    }
}
