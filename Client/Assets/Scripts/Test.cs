using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class Test : MonoBehaviour
{

    // Use this for initialization
    void Start()
    {
        LoadTables.Init();
        TestTable tt = TestTableManager.Instance.GetDataByID(1);
        Debug.Log(tt.Name);
    }

    // Update is called once per frame
    void Update()
    {

    }
}
