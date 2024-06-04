package replicate

import (
	"errors"
)

var ImageModels = []ReplicateModel{
	{
		Name:     "bytedance/sdxl-lightning-4step",
		Version:  "5f24084160c9089501c1b3545d9be3c27883ae2239b6f412990e82d4a6210f8f",
		Category: "Low",
	},
	{
		Name:     "lucataco/realvisxl-v2.0",
		Version:  "7d6a2f9c4754477b12c14ed2a58f89bb85128edcdd581d24ce58b6926029de08",
		Category: "High",
	},
	{
		Name:     "playgroundai/playground-v2.5-1024px-aesthetic",
		Version:  "a45f82a1382bed5c7aeb861dac7c7d191b0fdf74d8d57c4a0e6ed7d4d0bf7d24",
		Category: "High",
	},
	{
		Name:     "lucataco/dreamshaper-xl-turbo",
		Version:  "0a1710e0187b01a255302738ca0158ff02a22f4638679533e111082f9dd1b615",
		Category: "Low",
	},
	{
		Name:     "lorenzomarines/astra",
		Version:  "6ce68112bcaefc7273692243c933c2dcbe0307757a932fede1ca5e12956e0029",
		Category: "High",
	},
}

func GetModelByName(name string, modelList []ReplicateModel) (*ReplicateModel, error) {
	for _, model := range modelList {
		if model.Name == name {
			return &model, nil
		}
	}
	return nil, errors.New("model not found")
}

func GetModelIndex(modelName string, modelList []ReplicateModel) int {
	for i, model := range modelList {
		if model.Name == modelName {
			return i
		}
	}
	return -1
}
