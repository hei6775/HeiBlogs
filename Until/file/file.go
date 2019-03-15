package file

import "os"

//判断指定目录或文件存不存在
func IsDirOrFileExist(path string) error {
	//Stat，返回指定文件目录的信息
	//如果是软连，返回指向的文件信息
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return err
		}
		panic(err)
	}
	return nil
}

//如果目录不存在创建指定目录
func MustMkdirIfNotExist(path string) {
	if err := IsDirOrFileExist(path); err != nil {
		//MkdirAll 每级目录都会检测创建 权限0777
		os.MkdirAll(path, os.ModePerm)
	}
}
