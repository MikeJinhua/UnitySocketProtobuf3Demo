using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class TableUtil
{
    
    public static int[] LoadInts(string data)
    {
        int[] decodeData = JsonUtility.FromJson<int[]>(data);
        return decodeData;
    }


    public static Dictionary<string,int> LoadDictString2Int(string data)
    {
        Dictionary<string, int> decodeData = JsonUtility.FromJson<Dictionary<string, int>>(data);
        return decodeData;
    }
}
