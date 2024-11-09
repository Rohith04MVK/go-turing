# go-turing

`go-turing` is a Go implementation of a Turing Machine. This project simulates the behavior of a Turing Machine, allowing you to define custom instructions and input tapes to see how the machine processes them.

## Features

- **Customizable Instructions**: Load instructions from a JSON file to define the behavior of the Turing Machine.
- **Tape Input**: Provide an initial tape input to the Turing Machine.
- **Rendering**: Optionally render the Turing Machine's state and tape at each step.
- **Interactive Mode**: Step through the Turing Machine's execution interactively.
- **Speed Control**: Adjust the speed of the Turing Machine's execution.

## Installation

To install the `go-turing` project, clone the repository and build the project using Go:

```sh
git clone https://github.com/Rohith04MVK/go-turing.git
cd go-turing
go build -o turing-machine
```

## Usage

To run the Turing Machine, use the following command:

```sh
./turing-machine -i <instructions.json> -t <tape> [options]
```

### Command Line Options

- `-b`: Begin state (default: `q0`)
- `-e`: End state (default: `qdone`)
- `-s`: Rendering speed in seconds (default: `0.3`)
- `-r`: Render Turing Machine (default: `false`)
- `-a`: Interactive mode (default: `false`)
- `-i`: Instructions, as JSON file (required)
- `-t`: Input tape (required)

### Example

```sh
./turing-machine -i ./instructions/hello_world.json -t " " -r -s 0.3
```

## Instructions Format
The instructions file should be a JSON file with the following structure:

```json
{
  "state1": {
    "symbol1": {
      "write": "symbol2",
      "move": "left|right",
      "next_state": "state2"
    }
  }
}
```

## License

This project is licensed under the MIT License.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## Reference

This project draws inspiration from the theoretical foundations laid by Alan Turing, one of the most influential figures in computer science. Turing's work on the concept of a universal machine laid the groundwork for modern computing. His pioneering ideas continue to inspire and drive innovation in the field of computer science.

For more information on Turing Machines and Alan Turing's contributions, you can refer to the following resources:

- [Wikipedia: Turing Machine](https://en.wikipedia.org/wiki/Turing_machine)
- [Wikipedia: Alan Turing](https://en.wikipedia.org/wiki/Alan_Turing)
- [Stanford Encyclopedia of Philosophy: Turing Machines](https://plato.stanford.edu/entries/turing-machine/)
- [The Turing Archive for the History of Computing](http://www.alanturing.net/)

Alan Turing's legacy is a testament to the power of human ingenuity and the enduring impact of his work on the world of computing.
