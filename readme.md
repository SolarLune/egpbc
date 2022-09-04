# egpbc :video_game:

egpbc is named constants for Nintendo, Sony, and Xbox game console controllers for [Ebitengine](https://github.com/hajimehoshi/ebiten). 

(And so, egpbc stands for Ebitengine GamePad Button Constants.)

# How to install

`go get github.com/solarlune/egpbc`

# How to use

Each brand is contained within its own package; you can easily import that package to use it.

```go

import "github.com/solarlune/egpbc/microsoft"

func main() {

    // Each brand has its own Constants struct that contains its button and axis 
    // constants and directions; this is done to make syntactic sugar possible / easy to manage
    xbox := egpbc.MicrosoftConstants

    if ebiten.IsStandardGamepadButtonPressed(0, xbox.A) {

        fmt.Println("The A Button was pressed.")

    }

    if ebiten.IsStandardGamepad

}

```

That's it.