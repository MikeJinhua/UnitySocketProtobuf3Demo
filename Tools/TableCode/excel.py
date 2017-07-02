import os

import xlrd

from TableCode.cs_data import GenCSTableData
from TableCode.cs_file import GenCSTableManagerFile, genCSLoadTablesFile
from TableCode.go_data import GenGolangTableData
from TableCode.go_file import GenGoTableManagerFile, genGolangLoadTablesFile
from const import excel_dir

def processExcel(filePath, fileName):
    if "." in fileName:
        fileName = fileName.split('.')
        fileName = fileName[0]
    data = xlrd.open_workbook(filePath)
    table = data.sheets()[0]
    nrows = table.nrows
    ncols = table.ncols

    cs_fields_index = []#filed index
    golang_fields_index = []#filed index
    tableKeysIndex = []

    if table.nrows == 0 or table.ncols == 0:
        print("empty file:" + fileName)

    for index in range(ncols):
        CS_row = table.cell(1, index).value
        if CS_row == "C" or CS_row == "CS":
            cs_fields_index.append(index)

        if CS_row == "S" or CS_row == "CS":
            golang_fields_index.append(index)

    if len(cs_fields_index) > 0:
        cs_files.append(fileName)
        GenCSTableManagerFile(fileName, cs_fields_index, table)
        GenCSTableData(fileName, cs_fields_index, table)

    if len(golang_fields_index) > 0:
        go_files.append(fileName)
        GenGoTableManagerFile(fileName, golang_fields_index, table)
        GenGolangTableData(fileName, golang_fields_index, table)

cs_files = []
go_files = []
def excel_start():
    excels = []
    for dir in os.listdir(excel_dir):  # 遍历当前目录所有文件和目录
        fileName = dir
        child = os.path.join(excel_dir, dir)  # 加上路径，否则找不到
        if os.path.isdir(child):  # 如果是目录，则继续遍历子目录的文件
            for file in os.listdir(child):
                if "~" in child:
                    continue
                if os.path.splitext(file)[1] == '.xlsx':  # 分割文件名和文件扩展名，并且扩展名为'proto'
                    fileName = file
                    processExcel(child, fileName)
                    excels.append(fileName)
        elif os.path.isfile(child):  # 如果是文件，则直接判断扩展名
            if "~" in child:
                continue
            if os.path.splitext(child)[1] == '.xlsx':
                processExcel(child, fileName)
                excels.append(fileName)

    genCSLoadTablesFile(cs_files)
    genGolangLoadTablesFile(go_files)