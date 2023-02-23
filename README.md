# Advanced Logging Engine (ALE)
A simple yet very powerful logging package great for quick projects and scaling your log collection design.

![image](https://user-images.githubusercontent.com/20157708/220814053-ec10383a-eb9f-433c-95b5-fc49d2c2af90.png)

## How to Install
    go get -u github.com/seanmmitchll/ale

## How to Use
    import (
        "github.com/seanmmitchell/ale"
        "github.com/seanmmitchell/ale/pconsole"
    )

    func main() {
        le := ale.CreateLogEngine("Example")
        le.AddLogPipeline(ale.Info, pconsole.Log)

        le.Log(ale.Critical, "Critical Log")
        le.Log(ale.Error, "Error Log")
        le.Log(ale.Warning, "Warning Log")
        le.Log(ale.Info, "Info Log")
        le.Log(ale.Verbose, "Verbose Log")
        le.Log(ale.Debug, "Debug Log")

        sle := le.CreateSubEngine("Sub-Engine")
        sle.AddLogPipeline(ale.Warning, pconsole.Log)

        sle.Log(ale.Critical, "Critical Log")
        sle.Log(ale.Error, "Error Log")
        sle.Log(ale.Warning, "Warning Log")
        sle.Log(ale.Info, "Info Log")
        sle.Log(ale.Verbose, "Verbose Log")
        sle.Log(ale.Debug, "Debug Log")
    }

![image](https://user-images.githubusercontent.com/20157708/220814094-2e3c66ed-4284-48df-8626-adddcf59e032.png)

## License
This work is licensed under the MIT License.  
Please review [LICENSE](LICENSE.md) (LICENSE.md) for specifics.

## Features in Mind
- Inheritable log pipelines.
- A way to attach to other common packages (syslog, etc)
