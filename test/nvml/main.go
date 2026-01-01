package main

import (
	"fmt"
	"time"

	"github.com/NVIDIA/go-nvml/pkg/nvml"
)

func main() {
	ret := nvml.Init()
	if ret != nvml.SUCCESS {
		panic(fmt.Sprintf("NVML 初始化失败: %v", nvml.ErrorString(ret)))
	}
	fmt.Println("Init completed at:", time.Now().Format("2006-01-02 15:04:05"))
	defer nvml.Shutdown()

	count, _ := nvml.DeviceGetCount()
	var healthyIndices []int

	for i := 0; i < count; i++ {
		_, ret := nvml.DeviceGetHandleByIndex(i)
		if ret != nvml.SUCCESS {
			fmt.Printf("GPU %d 初始化失败: %v\n", i, nvml.ErrorString(ret))
			continue
		}
		healthyIndices = append(healthyIndices, i)
	}

	fmt.Printf("健康 GPU 数量: %d / %d\n", len(healthyIndices), count)

	// 无限循环采集状态
	for {
		fmt.Println("------ GPU 状态采集 ------")
		for _, i := range healthyIndices {
			dev, _ := nvml.DeviceGetHandleByIndex(i)
			name, _ := dev.GetName()
			mem, _ := dev.GetMemoryInfo()
			temp, _ := dev.GetTemperature(nvml.TEMPERATURE_GPU)
			power, _ := dev.GetPowerUsage()

			fmt.Printf("GPU %d: %s\n", i, name)
			fmt.Printf("  显存使用: %d / %d MiB\n", mem.Used/1024/1024, mem.Total/1024/1024)
			fmt.Printf("  温度: %d°C\n", temp)
			fmt.Printf("  功耗: %.1f W\n", float64(power)/1000)
		}

		time.Sleep(5 * time.Second) // 每 5 秒采集一次
	}
}
