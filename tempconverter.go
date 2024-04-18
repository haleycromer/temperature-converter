Last login: Thu Apr 18 15:17:01 on ttys002
haleycromer@Haleys-MacBook-Pro ~ % mkdir ~/go

haleycromer@Haleys-MacBook-Pro ~ % mkdir -p ~/go/{src,bin,pkg}

haleycromer@Haleys-MacBook-Pro ~ % nano ~/.bash_profile

haleycromer@Haleys-MacBook-Pro ~ % source ~/.zshrc

source: no such file or directory: /Users/haleycromer/.zshrc
haleycromer@Haleys-MacBook-Pro ~ % source ~/.bash_profile

haleycromer@Haleys-MacBook-Pro ~ % source ~/.zshrc

source: no such file or directory: /Users/haleycromer/.zshrc
haleycromer@Haleys-MacBook-Pro ~ % go version

go version go1.22.2 darwin/arm64
haleycromer@Haleys-MacBook-Pro ~ % >....                                        
                os.Exit(1)
        }

        // Determine the unit of temperature
        unit := os.Args[2]

        // Convert the temperature based on the unit
        var convertedTemperature float64
        switch unit {
        case "C":
                convertedTemperature = celsiusToFahrenheit(temperature)
                fmt.Printf("%.2f°C = %.2f°F\n", temperature, convertedTemperature)
        case "F":
                convertedTemperature = fahrenheitToCelsius(temperature)
                fmt.Printf("%.2f°F = %.2f°C\n", temperature, convertedTemperature)
        default:
                fmt.Println("Invalid unit:", unit)
                os.Exit(1)
        }
}

zsh: parse error near `}'
haleycromer@Haleys-MacBook-Pro ~ % 
haleycromer@Haleys-MacBook-Pro ~ % nano ~/.zshrc

haleycromer@Haleys-MacBook-Pro ~ % source ~/.zshrc

source: no such file or directory: /Users/haleycromer/.zshrc
haleycromer@Haleys-MacBook-Pro ~ % >....                                        
                os.Exit(1)
        }

        // Determine the unit of temperature
        unit := os.Args[2]

        // Convert the temperature based on the unit
        var convertedTemperature float64
        switch unit {
        case "C":
                convertedTemperature = celsiusToFahrenheit(temperature)
                fmt.Printf("%.2f°C = %.2f°F\n", temperature, convertedTemperature)
        case "F":
                convertedTemperature = fahrenheitToCelsius(temperature)
                fmt.Printf("%.2f°F = %.2f°C\n", temperature, convertedTemperature)
        default:
                fmt.Println("Invalid unit:", unit)
                os.Exit(1)
        }
}

zsh: parse error near `}'
haleycromer@Haleys-MacBook-Pro ~ % go run program.go

stat program.go: no such file or directory
haleycromer@Haleys-MacBook-Pro ~ % go run program.go

stat program.go: no such file or directory
haleycromer@Haleys-MacBook-Pro ~ % cd path/to/directory

cd: no such file or directory: path/to/directory
haleycromer@Haleys-MacBook-Pro ~ % ls program.go

ls: program.go: No such file or directory
haleycromer@Haleys-MacBook-Pro ~ % package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}

function dquote> cd path/to/directory

function dquote> touch program.go

function dquote> ls program.go

function dquote> go run program.go
function dquote> >....                                                          
                os.Exit(1)
        }

        // Determine the unit of temperature
        unit := os.Args[2]

        // Convert the temperature based on the unit
        var convertedTemperature float64
        switch unit {
        case "C":
                convertedTemperature = celsiusToFahrenheit(temperature)
                fmt.Printf("%.2f°C = %.2f°F\n", temperature, convertedTemperature)
        case "F":
                convertedTemperature = fahrenheitToCelsius(temperature)
                fmt.Printf("%.2f°F = %.2f°C\n", temperature, convertedTemperature)
        default:
                fmt.Println("Invalid unit:", unit)
                os.Exit(1)
        }
}

zsh: parse error near `<'
haleycromer@Haleys-MacBook-Pro ~ % touch program.go

haleycromer@Haleys-MacBook-Pro ~ % 
