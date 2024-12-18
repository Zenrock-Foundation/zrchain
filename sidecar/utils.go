package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func (o *Oracle) LoadFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			// Initialize with empty state if file doesn't exist
			o.updateChan <- EmptyOracleState
			o.stateCache = []OracleState{EmptyOracleState}
			return nil
		}
		return err
	}
	defer file.Close()

	var states []OracleState
	if err := json.NewDecoder(file).Decode(&states); err != nil {
		return fmt.Errorf("failed to decode state file: %w", err)
	}

	if len(states) > 0 {
		latestState := &states[len(states)-1]
		o.updateChan <- OracleState{
			Delegations:    latestState.Delegations,
			EthBlockHeight: latestState.EthBlockHeight,
			EthBlockHash:   latestState.EthBlockHash,
			EthGasLimit:    latestState.EthGasLimit,
			EthBaseFee:     latestState.EthBaseFee,
			EthTipCap:      latestState.EthTipCap,
			ETHUSDPrice:    latestState.ETHUSDPrice,
			ROCKUSDPrice:   latestState.ROCKUSDPrice,
		}
		o.stateCache = states
	} else {
		// Initialize with empty state if the file is empty
		o.updateChan <- EmptyOracleState
		o.stateCache = []OracleState{EmptyOracleState}
	}

	return nil
}

func (o *Oracle) SaveToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(o.stateCache)
}

func (o *Oracle) CacheState() {
	currentState := o.currentState.Load().(*OracleState)
	newState := *currentState // Create a copy of the current state

	// Update the ID and store the new state
	o.currentState.Store(&newState)

	// Cache the new state
	o.stateCache = append(o.stateCache, newState)
	if len(o.stateCache) > CacheSize {
		o.stateCache = o.stateCache[1:]
	}

	if err := o.SaveToFile(o.Config.StateFile); err != nil {
		log.Printf("Error saving state to file: %v", err)
	}
}

func (o *Oracle) getStateByEthHeight(height uint64) (*OracleState, error) {
	for _, state := range o.stateCache {
		if state.EthBlockHeight == height {
			return &state, nil
		}
	}
	return nil, fmt.Errorf("state with Ethereum block height %d not found", height)
}

func LoadConfig() Config {
	configFile := getConfigFile()
	cfg, err := readConfig(configFile)
	if err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}
	return cfg
}

func getConfigFile() string {
	configFile := os.Getenv("SIDECAR_CONFIG_FILE")
	if configFile == "" {
		configFile = "config.yaml"
	}
	return configFile
}

func readConfig(configFile string) (Config, error) {
	yamlFile, err := os.ReadFile(configFile)
	if err != nil {
		return Config{}, fmt.Errorf("unable to read config from %s: %v", configFile, err)
	}

	rootConfig := Config{}
	if err = yaml.Unmarshal(yamlFile, &rootConfig); err != nil {
		return Config{}, fmt.Errorf("error unmarshalling config from %s: %v", configFile, err)
	}

	return rootConfig, nil
}
