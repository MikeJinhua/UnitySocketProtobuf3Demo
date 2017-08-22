from const import go_table_file_dir


def GenGoTableManagerFile(tableName, fieldsIndex, table):
    filePath = go_table_file_dir + tableName.lower() + ".go"
    fileContent = ""
    fileContent += \
    "package gamedata \n\n"
    fileContent += \
    "type "+ tableName+" struct {\n"
    for index in fieldsIndex:
        fieldtype = table.cell(2, index).value
        fieldtype = fieldtype.lower()
        if fieldtype == "Map[String]Int".lower():
            fieldtype = "map[string]int"

        if fieldtype == "int[]":
            fieldtype = "[]int"

        fieldName = table.cell(3, index).value
        fileContent += \
    "    " + fieldName +" "+fieldtype + "\n"

    fileContent += \
"""
}

var (
"""
    fileContent += "    {0}Data = make(map[int]{1})\n".format(tableName,tableName)
    fileContent += ")\n\n"

    fileContent += "func  "+tableName+"init() {\n"
    fileContent += "	rf := readRf("+tableName+"{})\n"
    fileContent += "	for i := 0; i < rf.NumRecord(); i++ {\n"
    fileContent += "		r := rf.Record(i).(*{0})\n".format(tableName)

    fileContent += "        {0}Data[r.{1}] = *r\n".format(tableName,table.cell(3, 0).value)

    fileContent += "    }\n}\n\n"
    fileContent += "func Get"+tableName+"ByID(id int) ("+tableName+") {\n"
    fileContent += "	return  {0}Data[id]\n".format(tableName)
    fileContent += "}\n"

    fo = open(filePath, "wb")
    fo.write(fileContent.encode('utf-8'))
    fo.close()


def genGolangLoadTablesFile(files):
    filePath = go_table_file_dir + "tables.go"
    fileContent = ""
    fileContent += \
    "package gamedata\
\
\nfunc LoadTables()  {\n"
    for file in files:
        fileContent += "    {0}init()\n".format(file)
    fileContent +="}"


    fo = open(filePath, "wb")
    fo.write(fileContent.encode('utf-8'))
    fo.close()
