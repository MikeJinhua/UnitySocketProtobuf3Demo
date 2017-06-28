using UnityEngine;
using System.Collections;
using Msg;

public class LoginModel : BaseModel<LoginModel>
{
    protected override void InitAddTocHandler()
    {
        AddTocHandler(typeof(SignUpResponse), TocLogin);
    }

    private void TocLogin(object data)
    {
        SignUpResponse toc = data as SignUpResponse;
        Debug.Log(toc.Version);
    }
}
