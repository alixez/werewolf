package werewolf

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/alixez/werewolf/utils"
	"gopkg.in/yaml.v2"
)

type Env struct {
	Appname     string
	Version     string
	Environment string
	Debug       bool
	HasLoaded   bool
	devConfigs  map[interface{}]interface{}
	prodConfigs map[interface{}]interface{}
}

func (obj *Env) Init(configs map[interface{}]interface{}) {
	obj.Appname = configs["appname"].(string)
	obj.Version = configs["version"].(string)
	obj.Environment = configs["environment"].(string)
	obj.Debug = obj.Environment != "production"
	obj.devConfigs = configs["development"].(map[interface{}]interface{})
	obj.prodConfigs = configs["production"].(map[interface{}]interface{})
	obj.HasLoaded = true
}

// GetConfig like a.b.c
func (obj *Env) GetConfig(query string) interface{} {
	querySlice := strings.Split(query, ".")
	value := obj.prodConfigs[querySlice[0]]
	if obj.Environment == "development" {
		value = obj.devConfigs[querySlice[0]]
	}
	if len(querySlice) == 1 {
		return value
	}
	for _, field := range querySlice[1:] {
		if v, ok := value.(map[interface{}]interface{}); ok {
			value = v[field]
			continue
		}
		value = nil
	}
	return value
}

// SetConfig is a func to set configuration
func (obj *Env) SetConfig(field string, value interface{}) {
	if obj.Environment == "development" {
		obj.devConfigs[field] = value
	} else if obj.Environment == "production" {
		obj.prodConfigs[field] = value
	}
}

var systemEnv = &Env{
	HasLoaded: false,
}

func listDir(path string, suffix string) (files []string, err error) {
	if !utils.IsDirExist(path) {
		err := os.Mkdir(path, 0777)
		if err != nil {
			return nil, err
		}
	}
	files = []string{}
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	suffix = strings.ToUpper(suffix)

	for _, fi := range dir {
		if fi.IsDir() {
			continue
		}

		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, fi.Name())
		}
	}

	return files, nil
}

func LoadApplicationEnv() (env *Env) {
	if systemEnv.HasLoaded {
		env = systemEnv
		return
	}
	fmt.Println("...开始加载配置文件...")
	filepathList, err := listDir("config", "yaml")
	if err != nil {
		panic("获取配置文件列表时发生错误")
	}
	if !utils.ArrayContainer(filepathList, "env.yaml") {
		filepathList = append(filepathList, "env.yaml")
		f, _ := os.Create(path.Join("config", "env.yaml"))
		f.WriteString("appname: arku\r\nversion: v1.0\r\nenvironment: development")
	}
	if !utils.ArrayContainer(filepathList, "default.yaml") {
		filepathList = append(filepathList, "default.yaml")
		f, _ := os.Create(path.Join("config", "default.yaml"))
		f.WriteString("production:\r\ndevelopment:\r\n")
	}

	// filepathList := []string{
	// 	"env.yaml",
	// 	"default.yaml",
	// }
	masterConfigs := map[interface{}]interface{}{
		"environment": "development",
		"version":     "v0.1",
		"appname":     "demo",
		"development": map[interface{}]interface{}{},
		"production":  map[interface{}]interface{}{},
	}
	for _, v := range filepathList {
		readConfigFromFile(v, masterConfigs)
	}
	systemEnv.Init(masterConfigs)
	env = systemEnv
	return
}

func mergeConfig(m1 interface{}, m2 interface{}) interface{} {
	if m1 == nil {
		m1 = map[interface{}]interface{}{}
	}
	if m2 == nil {
		m2 = map[interface{}]interface{}{}
	}
	m3 := m2.(map[interface{}]interface{})
	for k, v := range m1.(map[interface{}]interface{}) {
		m3[k] = v
	}
	return m3
}

func readConfigFromFile(filepath string, out map[interface{}]interface{}) {
	configs := make(map[interface{}]interface{})
	configByte, err := ioutil.ReadFile(path.Join("config", filepath))
	if err != nil {

		panic(err)
	}
	err = yaml.Unmarshal(configByte, &configs)
	if err != nil {
		panic(err)
	}
	if filepath == "env.yaml" {
		out["environment"] = configs["environment"]
		out["version"] = configs["version"]
		out["appname"] = configs["appname"]
	} else {
		out["production"] = mergeConfig(configs["production"], out["production"])
		out["development"] = mergeConfig(configs["development"], out["development"])
	}
}
