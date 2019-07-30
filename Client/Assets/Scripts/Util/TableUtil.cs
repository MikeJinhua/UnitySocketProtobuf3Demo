using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class TableUtil
{
    
    public static int[] LoadInts(string data)
    {
        int[] decodeData = FromJson<int[]>(data);
        return decodeData;
    }


    public static Dictionary<string,int> LoadDictString2Int(string data)
    {
        Dictionary<string, int> decodeData = JsonUtility.FromJson<Dictionary<string, int>>(data);
        return decodeData;
    }
    
    /// <summary> 解析Json </summary>
    /// <typeparam name="T">类型</typeparam>
    /// <param name="json">Json字符串</param>
    public static T FromJson<T>(string json)
    {
        if (json == "null" && typeof(T).IsClass) return default(T);

        if (typeof(T).GetInterface("IList") != null)
        {
            json = "{\"data\":{data}}".Replace("{data}", json);
            Pack<T> Pack = JsonUtility.FromJson<Pack<T>>(json);
            return Pack.data;
        }

        return JsonUtility.FromJson<T>(json);
    }

    /// <summary> 内部包装类 </summary>
    private class Pack<T>
    {
        public T data;
    }
}
