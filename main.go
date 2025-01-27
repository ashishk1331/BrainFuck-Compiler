package main

import (
	"fmt"
	"os"
)

const (
	Left byte	= 60
	Right		= 62
	Add		= 43
	Sub		= 45
	Out		= 46
	In		= 44
	Start		= 91
	Stop		= 93
)

type Token struct {
	t	byte
	index	int
}

var braceLookUp map[int]int = make(map[int]int);

func parseTokens(content []byte) []Token {
	var res []Token = []Token{};
	for _, char := range content {
		switch char{
		case Left, Right, In, Out, Start, Stop, Add, Sub:
			res = append(res, Token{t:char});
		}
	}

	for index,_ := range res {
		res[index].index = index;
	}
	return res;
}

func fillBraceLookUp(stream []Token) {
	var stack []int;
	for _, token := range stream {
		if token.t == Start {
			stack = append(stack, token.index);
		} else if token.t == Stop {
			braceLookUp[stack[len(stack) - 1]] = token.index;
			braceLookUp[token.index] = stack[len(stack) - 1];
			stack = stack[:len(stack) - 1];
		}
	}
}

func evaluate(stream []Token) {
	var (
		inst	int = 0
		dp	int = 0
		data	[]byte = make([]byte, 3000)
	)

	for inst < len(stream) {
		switch stream[inst].t {
		case Left:
			dp -= 1;
		case Right:
			dp += 1;
		case Add:
			data[dp] += 1;
		case Sub:
			data[dp] -= 1;
		case Out:
			fmt.Print(string(data[dp]));
		case In:
			var target []byte = make([]byte, 1);
			_, err := os.Stdin.Read(target);

			if err != nil {
				panic("inp: invalid input.");
			}
			data[dp] = target[0];
		case Start:
			if data[dp] == 0 {
				inst = braceLookUp[inst];
			}
		case Stop:
			if data[dp] != 0 {
				inst = braceLookUp[inst];
			}
		default:
			fmt.Println(string(stream[inst].t), ": invalid token found.")
		}

		if dp < 0 || dp >= 3000 {
			panic("dp: index out of range.");
		}

		inst += 1;
	}
}



func main() {
	content, err := os.ReadFile("sample.txt");

	if err != nil {
		fmt.Println(err);
	}

	tokens := parseTokens(content);
	fillBraceLookUp(tokens);
	evaluate(tokens);

}
