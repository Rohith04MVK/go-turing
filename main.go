package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Rohith04MVK/turing-machine/config"
	"github.com/Rohith04MVK/turing-machine/utils"
)

type TuringMachine struct {
	instructions map[string]map[string]map[string]string
	tape         []string
	state        string
	endState     string
	interactive  bool
	render       bool
	speed        float64
}

func NewTuringMachine(instructions map[string]map[string]map[string]string, tape string, startState string, endState string, render bool, speed float64, interactive bool) *TuringMachine {
	tapeSlice := make([]string, len(tape))
	for i, c := range tape {
		tapeSlice[i] = string(c)
	}

	tm := &TuringMachine{
		instructions: instructions,
		tape:         tapeSlice,
		state:        startState,
		endState:     endState,
		interactive:  interactive,
		render:       render,
		speed:        speed,
	}
	tm.validateInstruction()
	return tm
}

func (tm *TuringMachine) Run() string {
	tapeIndex := 0
	stepsCounter := 0
	tm.render_(tapeIndex, stepsCounter, true)

	for tm.state != tm.endState {
		stepsCounter++
		tapeIndex = tm.calculateNextState(tapeIndex)
		tm.render_(tapeIndex, stepsCounter, false)
	}

	tm.render_(tapeIndex, stepsCounter, true)
	return utils.ListToString(utils.RemoveEmptyCharacter(tm.tape))
}

func (tm *TuringMachine) calculateNextState(tapeIndex int) int {
	if tapeIndex == -1 {
		tm.tape = append([]string{config.EmptyCharacter()}, tm.tape...)
		tapeIndex = 0
	}
	if tapeIndex == len(tm.tape) {
		tm.tape = append(tm.tape, config.EmptyCharacter())
	}

	action := tm.instructions[tm.state][tm.tape[tapeIndex]]
	tm.tape[tapeIndex] = action["write"]
	tm.state = action["nextState"]
	return utils.NextIndex(tapeIndex, action["move"])
}

func (tm *TuringMachine) validateInstruction() {
	for instruction := range tm.instructions {
		for case_ := range tm.instructions[instruction] {
			action := tm.instructions[instruction][case_]
			if len(action["write"]) != 1 {
				panic(fmt.Sprintf("Invalid config! Use ONE character, instead of %q!", action["write"]))
			}
			if !config.IsAllowedMovement(action["move"]) {
				panic(fmt.Sprintf("Invalid config! Use \"right\" or \"left\", not %q!", action["move"]))
			}
			if _, exists := tm.instructions[action["nextState"]]; !exists && action["nextState"] != tm.endState {
				panic(fmt.Sprintf("Invalid config! State %q needs to be defined!", action["nextState"]))
			}
		}
	}
}

func (tm *TuringMachine) render_(index int, stepsCounter int, forceRender bool) {
	if tm.shouldRender(forceRender) {
		empty, paddingEnd, paddingStart, visibleLength, visibleTapeSection := tm.tapeFormatCalc(index)
		tm.printSystemClear()
		tm.printStatistics(index, stepsCounter)
		tm.printRenderModeInformation()
		tm.printTape(empty, paddingEnd, paddingStart, visibleLength, visibleTapeSection)
		tm.printSignOccurrences()

		if tm.interactive {
			fmt.Scanln()
		}
		if tm.speed > 0 {
			time.Sleep(time.Duration(tm.speed * float64(time.Second)))
		}
	}
}

func (tm *TuringMachine) shouldRender(forceRender bool) bool {
	return tm.render || tm.interactive || forceRender
}

func (tm *TuringMachine) printSystemClear() {
	fmt.Print("\033[H\033[2J") // ANSI escape codes to clear screen
}

func (tm *TuringMachine) tapeFormatCalc(index int) (string, int, int, int, string) {
	length := len(tm.tape)
	visibleLength := config.VisibleTapeLength()
	empty := config.EmptyCharacter()
	paddingStart := visibleLength - index
	paddingEnd := visibleLength - (length - (index + 1))

	dynamicStart := index - visibleLength
	if dynamicStart < 0 {
		dynamicStart = 0
	}

	dynamicEnd := length - (length - index - visibleLength)
	if length-index <= visibleLength {
		dynamicEnd = length
	}

	visibleTapeSection := strings.Join(tm.tape[dynamicStart:dynamicEnd], "")

	return empty, paddingEnd, paddingStart, visibleLength, visibleTapeSection
}

func (tm *TuringMachine) printSignOccurrences() {
	fmt.Println("Character Counter")
	occurrences := utils.CountOccurrences(tm.tape)
	for char, count := range occurrences {
		fmt.Printf("%dx: %s\n", count, char)
	}
}

func (tm *TuringMachine) printTape(empty string, paddingEnd int, paddingStart int, visibleLength int, visibleTapeSection string) {
	paddingIcons := strings.Repeat("=", visibleLength*2)
	fmt.Printf("%s▼%s\n", paddingIcons, paddingIcons)

	padding := strings.Repeat(empty, paddingStart)
	endPadding := strings.Repeat(empty, paddingEnd)
	fmt.Println(utils.Pipeify(padding + visibleTapeSection + endPadding))

	fmt.Printf("%s▲%s\n", paddingIcons, paddingIcons)
}

func (tm *TuringMachine) printRenderModeInformation() {
	fmt.Println("Render Mode")

	automaticModeText := " "
	if tm.render && !tm.interactive {
		automaticModeText = "X"
	}

	interactiveModeText := " "
	interactiveMessage := " "
	if tm.interactive {
		interactiveModeText = "X"
		interactiveMessage = "(Press enter to render next step...)"
	}

	noneModeText := " "
	noneMessage := " "
	if !tm.interactive && !tm.render {
		noneModeText = "X"
		noneMessage = "(Please wait for results...)"
	}

	fmt.Printf("[%s] Automatic\n", automaticModeText)
	fmt.Printf("[%s] Interactive %s\n", interactiveModeText, interactiveMessage)
	fmt.Printf("[%s] None %s\n", noneModeText, noneMessage)
}

func (tm *TuringMachine) printStatistics(index int, stepsCounter int) {
	fmt.Printf("Steps Counter %7d\n", stepsCounter)
	fmt.Printf("Current State %7s\n", tm.state)
	fmt.Printf("Tape Index %10d\n", index)
}

func loadInstructions(filename string) (map[string]map[string]map[string]string, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading instruction file: %w", err)
	}

	var instructions map[string]map[string]map[string]string
	if err := json.Unmarshal(file, &instructions); err != nil {
		return nil, fmt.Errorf("error parsing JSON instructions: %w", err)
	}

	return instructions, nil
}

func main() {
	beginPtr := flag.String("b", "q0", "Begin state")
	endPtr := flag.String("e", "qdone", "End state")
	speedPtr := flag.Float64("s", 0.3, "Rendering speed in seconds")
	renderPtr := flag.Bool("r", false, "Render turing machine")
	interactivePtr := flag.Bool("a", false, "Interactive mode")
	instructionsPtr := flag.String("i", "", "Instructions, as JSON file")
	tapePtr := flag.String("t", "", "Input tape")

	flag.Parse()

	if *instructionsPtr == "" || *tapePtr == "" {
		fmt.Println("Instructions and tape are required")
		os.Exit(1)
	}

	if *instructionsPtr == "" || *tapePtr == "" {
		fmt.Println("Instructions and tape are required")
		flag.Usage()
		os.Exit(1)
	}

	instructions, err := loadInstructions(*instructionsPtr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading instructions: %v\n", err)
		os.Exit(1)
	}

	tm := NewTuringMachine(instructions, *tapePtr, *beginPtr, *endPtr, *renderPtr, *speedPtr, *interactivePtr)
	result := tm.Run()
	fmt.Println("Final result:", result)
}
