package main

// 断路器模式，就是遇到错误或者障碍就应该及时进行中断
// 场景：路由过滤中间件、api接口参数校验等

// 以下示例（电路与断路器）：
// 断路器是指能够关合、承载和开断正常回路条件下的电流并能在规定的时间内关合、承载和开断异常回路条件下的电流的开关装置。
// 断路器的作用是切断和接通负荷电路，以及切断故障电路，防止事故扩大，保证安全运行


// 定义电路的基本参数
type Circuit struct {
	voltage 	int32 //电压
	electric	int32 //电流
	resistance	int32 //电阻
}

// 定义断路器，如果出现超过阈值则立即抛出异常
// @throw
type Breaker func(Circuit)

// 创建电路
func NewCircuit()  Circuit {
	return Circuit{
		voltage:    12,
		electric:   24,
		resistance: 2,
	}
}

type BreakerController struct {
	counter int32
	breakers []Breaker
}

func NewBreakerController() BreakerController {
	breakers := []Breaker{VoltageBreaker(), ElectricBreaker(), ResistanceBreaker()}
	return BreakerController{
		counter: 0,
		breakers: breakers,
	}
}

// 定义电压断路器
func VoltageBreaker() Breaker {
	failureThreshold := 30

	return func(circuit Circuit) {
		if circuit.voltage >= (int32(failureThreshold)) {
			panic("电压超过阈值，电路中断")
		}
	}
}

// 定义电流断路器
func ElectricBreaker() Breaker {
	failureThreshold := 20

	return func(circuit Circuit) {
		if circuit.electric >= (int32(failureThreshold)) {
			panic("电流超过阈值，电路中断")
		}
	}
}

// 定义电阻断路器
func ResistanceBreaker() Breaker {
	failureThreshold := 5

	return func(circuit Circuit) {
		if circuit.resistance >= (int32(failureThreshold)) {
			panic("电阻超过阈值，电路中断")
		}
	}
}

func main()  {
	circuit := NewCircuit()
	controller := NewBreakerController()

	for _, breakerDealer := range controller.breakers{
		breakerDealer(circuit)
	}
}
