package mqttc

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go.ebupt.com/lets/app"
)

var instances map[string]*LetsMQTT = make(map[string]*LetsMQTT)
var once sync.Once

type LetsMQTT struct {
	ConfigID   string
	Opts       *mqtt.ClientOptions
	PahoClient mqtt.Client
}

var r = rand.New(rand.NewSource(time.Now().Unix()))

/*
 生成随机的ClientID
*/
func genRandomClientID(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

func GetInstance(configID string) *LetsMQTT {

	once.Do(func() {
		ins, err := new(configID)
		if err != nil {
			errorMsg := fmt.Sprintf("获取MQTT客户端实例失败[errorMsg: %v]", err)
			app.LLog.Error(errorMsg)
			panic(errorMsg)
		}
		instances[configID] = ins
	})
	return instances[configID]

}

func new(configID string) (*LetsMQTT, error) {

	config := app.LConfig.Get("MQTT." + configID)

	if config == nil {
		return nil, fmt.Errorf(fmt.Sprintf("获取配置文件ID:%v错误", configID))
	}

	configMap, ok := (config).(map[string]interface{})

	if !ok {
		return nil, fmt.Errorf(fmt.Sprintf("转换配置文件为map时发生错误,configID=%v", configID))
	}

	broker, exists := configMap["broker"]
	if !exists || broker == "" {
		return nil, fmt.Errorf(fmt.Sprintf("缺少必填的配置信息Broker[configID=%v]", configID))
	}

	port, exists := configMap["port"]
	if !exists || port == "" {
		return nil, fmt.Errorf(fmt.Sprintf("缺少必填的配置信息Port[configID=%v]", configID))
	}

	userName, exists := configMap["username"]
	if !exists {
		return nil, fmt.Errorf(fmt.Sprintf("缺少必填的配置信息UserName[configID=%v]", configID))
	}

	password, exists := configMap["password"]
	if !exists {
		return nil, fmt.Errorf(fmt.Sprintf("缺少必填的配置信息Password[configID=%v]", configID))
	}

	clientID, exists := configMap["clientid"]
	if !exists || clientID == "" {
		clientID = (clientID).(string)
		clientID = genRandomClientID(15)
	}

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID((clientID).(string))
	opts.SetUsername((userName).(string))
	opts.SetPassword((password).(string))
	opts.SetAutoReconnect(true)

	return &LetsMQTT{
		ConfigID:   configID,
		Opts:       opts,
		PahoClient: nil,
	}, nil

}

/*
	设置连接成功回调函数
*/
func (l *LetsMQTT) Connect() error {
	if l.PahoClient == nil {
		client := mqtt.NewClient(l.Opts)
		if token := client.Connect(); token.Wait() && token.Error() != nil {
			return fmt.Errorf("连接MQTT服务失败[configID=%v]，error:%v", l.ConfigID, token.Error())
		}
		l.PahoClient = client
	} else {
		app.LLog.Debug("连接已建立不需要重连")
	}
	return nil
}
