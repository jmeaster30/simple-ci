package main

import (
	"bufio"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type ActionConfig struct {
	Action                WorkflowAction `yaml:"action"`
	If                    string         `yaml:"if"`
	BuildCommand          string         `yaml:"build_command"`
	ArtifactUploadCommand string         `yaml:"upload_command"`
}

type ProjectConfig struct {
	RepositoryUrl string         `yaml:"repository_url"`
	Actions       []ActionConfig `yaml:"actions"`
}

type Config struct {
	Projects []ProjectConfig `yaml:"projects"`
}

func LoadConfig(file string) (*Config, error) {
	filedata, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer filedata.Close()

	stat, err := filedata.Stat()
	if err != nil {
		return nil, err
	}

	byteArr := make([]byte, stat.Size())
	_, err = bufio.NewReader(filedata).Read(byteArr)
	if err != nil && err != io.EOF {
		return nil, err
	}

	config := Config{}
	err = yaml.Unmarshal(byteArr, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
