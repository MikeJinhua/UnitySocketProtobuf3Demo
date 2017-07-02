from const import go_table_data_dir


def GenGolangTableData(tableName, fieldsIndex, table):
    filePath = go_table_data_dir + tableName + ".txt"

    fileContent = ""

    #write cols name
    for index in fieldsIndex:
        fileContent += table.cell(3, index).value + "\t"
    fileContent = fileContent.rstrip('\t')
    fileContent = fileContent.rstrip()
    fileContent += "\n"
    for row in range(4,table.nrows):
        for ncols in range(table.ncols):
            if ncols in fieldsIndex:
                fieldType = table.cell(2, ncols).value
                if fieldType == "Int":
                    value = table.cell(row, ncols).value
                    fileContent += ("{0}\t").format(int(table.cell(row, ncols).value))
                else:
                    fileContent += ("{0}\t").format(table.cell(row, ncols).value)
        fileContent = fileContent.rstrip('\t')
        fileContent += "\n"

    fo = open(filePath, "wb")
    fo.write(fileContent.encode('utf-8'))
    fo.close()
