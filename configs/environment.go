package configs

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
)

//ServerConfiguration configurations for Server
type ServerConfiguration struct {
	APIServer          string
	APIPort            string
	MaxRequestBodySize int
}

//MongoDBConfiguration configurations for MongoDB
type MongoDBConfiguration struct {
	MongoServer string
	User        string
	Pass        string
}

// ServerConfig information about Server
var ServerConfig = CreateServerConfigurationInstance()

// MongoDBConfig information about MongoDB
var MongoDBConfig = CreateMongoDBConfigurationInstance()

//CreateServerConfigurationInstance crate Server instance
func CreateServerConfigurationInstance() *ServerConfiguration {
	servConfig := new(ServerConfiguration)
	return servConfig
}

//CreateMongoDBConfigurationInstance creates empty MongoDB configuration
func CreateMongoDBConfigurationInstance() *MongoDBConfiguration {
	mongoConfig := new(MongoDBConfiguration)
	return mongoConfig
}

//Create prepare all the maps for configurations
func Create(env string) {

	readServConfigFile(env)

	readMongoDBConfigFile(env)

}

func readServConfigFile(env string) {

	servFile := "../configs/files/server/config.server." + env + ".json"

	serverMapped := genericReadFile(servFile)

	ServerConfig.APIServer = serverMapped["apiServer"]
	ServerConfig.APIPort = ":" + serverMapped["apiPort"]
	ServerConfig.MaxRequestBodySize, _ = strconv.Atoi(serverMapped["maxRequestBodySize"])

}

func readMongoDBConfigFile(env string) {

	var mongoFile = "../configs/files/mongodb/config.mongodb." + env + ".json"

	mongoBDMapped := genericReadFile(mongoFile)

	MongoDBConfig.MongoServer = os.Getenv("MONGO_URL") //mongoBDMapped["server"]
	MongoDBConfig.Pass = mongoBDMapped["pass"]
	MongoDBConfig.User = mongoBDMapped["user"]

}

func genericReadFile(path string) map[string]string {

	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	data, _ := ioutil.ReadAll(file)

	var mapped map[string]string

	json.Unmarshal([]byte(data), &mapped)

	return mapped
}
