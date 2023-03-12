package connector

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

const filePath = "vmess.txt"

var saver PersistenceStore = NewFilePersistenceStore(filePath)

type connector struct {
	V    string `json:"v"`
	PS   string `json:"ps"`
	Add  string `json:"add"`
	Port string `json:"port"`
	ID   string `json:"id"`
	AID  string `json:"aid"`
	Net  string `json:"net"`
	Type string `json:"type"`
	Host string `json:"host"`
	Path string `json:"path"`
	TLS  string `json:"tls"`
}

func converter(configStr string) (*connector, error) {
	var err error
	var config = &connector{}

	con := strings.Split(configStr, "://")
	if len(con) == 2 && con[0] == "vmess" {
		var tmp []byte
		tmp, err = base64.StdEncoding.DecodeString(con[1])
		if err != nil {
			logrus.Errorf("failed to decode: %v", err)
			return nil, err
		}
		if err = json.Unmarshal(tmp, config); err != nil {
			logrus.Errorf("failed to unmarshal: %v", err)
			return nil, err
		}
	} else {
		err = fmt.Errorf("unknown format of vmess")
		return nil, err
	}

	config.PS = config.Host
	return config, nil
}

func WriteConnector(ctx context.Context, newConnector string) (err error) {
	defer func() {
		if err != nil {
			logrus.Error(err)
		}
	}()

	new, err := converter(newConnector)
	if err != nil {
		return
	}

	toSave, err := json.Marshal(new)
	if err != nil {
		return
	}

	return saver.Save("vmess://" + base64.RawStdEncoding.EncodeToString(toSave) + "\n")
}

func GetConnector(ctx context.Context) (encoded string, err error) {
	defer func() {
		if err != nil {
			logrus.Error(err)
		}
	}()

	connector, err := saver.Load()
	if err != nil {
		return
	}

	// 对文件内容进行base64编码
	encoded = base64.StdEncoding.EncodeToString([]byte(strings.Join(connector, "\n")))
	return
}
