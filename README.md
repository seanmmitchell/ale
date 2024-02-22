# Advanced Logging Engine (ALE)
![ale-final-min](https://user-images.githubusercontent.com/20157708/221439553-6aceca40-0063-4493-92d3-944f22a3302d.png)

A simple yet very powerful logging package great for quick projects and scaling your log collection design.

![image](https://user-images.githubusercontent.com/20157708/220814053-ec10383a-eb9f-433c-95b5-fc49d2c2af90.png)

## How to Install
    go get -u github.com/seanmmitchell/ale/v2

## How to Use
```go
import (
    "github.com/seanmmitchell/ale/v2"
    "github.com/seanmmitchell/ale/v2/pconsole"
)

func main() {
    le := ale.CreateLogEngine("Example")
    pCTX, _ := pconsole.New(20, 20)
    le.AddLogPipeline(ale.Info, pCTX.Log)

    le.Log(ale.Critical, "Critical Log")
    le.Log(ale.Error, "Error Log")
    le.Log(ale.Warning, "Warning Log")
    le.Log(ale.Info, "Info Log")
    le.Log(ale.Verbose, "Verbose Log")
    le.Log(ale.Debug, "Debug Log")

    sle := le.CreateSubEngine("Sub-Engine")
    sle.AddLogPipeline(ale.Warning, pCTX.Log)

    sle.Log(ale.Critical, "Critical Log")
    sle.Log(ale.Error, "Error Log")
    sle.Log(ale.Warning, "Warning Log")
    sle.Log(ale.Info, "Info Log")
    sle.Log(ale.Verbose, "Verbose Log")
    sle.Log(ale.Debug, "Debug Log")
}
```

![image](https://user-images.githubusercontent.com/20157708/220814234-054730cb-fc55-468f-bd00-795e050afa93.png)

## License
This work is licensed under the MIT License.  
Please review [LICENSE](LICENSE.md) (LICENSE.md) for specifics.
