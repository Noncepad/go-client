package test

import "log"

//Holds arguments to be passed to service Arith in RPC call
type Args struct {
	A, B int
}

//Representss service Arith with method Multiply
type Arith int

//Result of RPC call is of this type
type Result int

//This procedure is invoked by rpc and calls rpcexample.Multiply which stores product of args.A and args.B in result pointer
func (t *Arith) Multiply(args Args, result *Result) error {
	log.Print("multiply")
	return Multiply(args, result)
}

//This procedure is invoked by rpc and calls rpcexample.Multiply which stores product of args.A and args.B in result pointer
func (t *Arith) Substract(args Args, result *Result) error {
	log.Print("subtract")
	return Subtract(args, result)
}

//stores product of args.A and args.B in result pointer
func Multiply(args Args, result *Result) error {
	log.Print("multiply")
	//log.Printf("Multiplying %d with %d\n", args.A, args.B)
	*result = Result(args.A * args.B)
	return nil
}

func Subtract(args Args, result *Result) error {
	log.Print("subtract")
	//log.Printf("Multiplying %d with %d\n", args.A, args.B)
	*result = Result(args.A - args.B)
	return nil
}
